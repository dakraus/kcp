/*
Copyright 2022 The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package validatingwebhook

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"sync"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/admission/configuration"
	"k8s.io/apiserver/pkg/admission/plugin/webhook/generic"
	"k8s.io/apiserver/pkg/admission/plugin/webhook/validating"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"

	kcpkubernetesinformers "github.com/kcp-dev/client-go/informers"
	kcpkubernetesclientset "github.com/kcp-dev/client-go/kubernetes"
	"github.com/kcp-dev/logicalcluster/v3"

	kcpinitializers "github.com/kcp-dev/kcp/pkg/admission/initializers"
	apisv1alpha1 "github.com/kcp-dev/kcp/sdk/apis/apis/v1alpha1"
	kcpinformers "github.com/kcp-dev/kcp/sdk/client/informers/externalversions"
)

const (
	PluginName = "apis.kcp.io/ValidatingWebhook"
)

type Plugin struct {
	*admission.Handler
	config []byte

	// Injected/set via initializers
	kubeClusterClient               kcpkubernetesclientset.ClusterInterface
	localKubeSharedInformerFactory  kcpkubernetesinformers.SharedInformerFactory
	globalKubeSharedInformerFactory kcpkubernetesinformers.SharedInformerFactory

	getAPIBindings func(clusterName logicalcluster.Name) ([]*apisv1alpha1.APIBinding, error)

	managerLock   sync.Mutex
	managersCache map[logicalcluster.Name]generic.Source
}

var (
	_ = admission.ValidationInterface(&Plugin{})
	_ = admission.InitializationValidator(&Plugin{})
	_ = kcpinitializers.WantsKubeClusterClient(&Plugin{})
	_ = kcpinitializers.WantsKubeInformers(&Plugin{})
	_ = kcpinitializers.WantsKcpInformers(&Plugin{})
)

func NewValidatingAdmissionWebhook(configFile io.Reader) (*Plugin, error) {
	p := &Plugin{
		managerLock:   sync.Mutex{},
		managersCache: make(map[logicalcluster.Name]generic.Source),
		Handler:       admission.NewHandler(admission.Connect, admission.Create, admission.Delete, admission.Update),
	}
	if configFile != nil {
		config, err := io.ReadAll(configFile)
		if err != nil {
			return nil, err
		}
		p.config = config
	}

	return p, nil
}

func Register(plugins *admission.Plugins) {
	plugins.Register(PluginName, func(configFile io.Reader) (admission.Interface, error) {
		return NewValidatingAdmissionWebhook(configFile)
	})
}

func (p *Plugin) Validate(ctx context.Context, attr admission.Attributes, o admission.ObjectInterfaces) error {
	cluster, err := genericapirequest.ValidClusterFrom(ctx)
	if err != nil {
		return err
	}
	clusterName := cluster.Name

	var config io.Reader
	if len(p.config) > 0 {
		config = bytes.NewReader(p.config)
	}

	hookSource, err := p.getHookSource(clusterName, attr.GetResource().GroupResource())
	if err != nil {
		return err
	}

	plugin, err := validating.NewValidatingAdmissionWebhook(config)
	if err != nil {
		return fmt.Errorf("error creating validating admission webhook: %w", err)
	}

	plugin.SetExternalKubeClientSet(p.kubeClusterClient.Cluster(clusterName.Path()))
	plugin.SetNamespaceInformer(p.localKubeSharedInformerFactory.Core().V1().Namespaces().Cluster(clusterName))
	plugin.SetHookSource(hookSource)
	plugin.SetReadyFuncFromKCP(p.localKubeSharedInformerFactory.Core().V1().Namespaces().Cluster(clusterName))

	if err := plugin.ValidateInitialization(); err != nil {
		return fmt.Errorf("error validating ValidatingAdmissionWebhook initialization: %w", err)
	}

	// Add cluster annotation on create
	if attr.GetOperation() == admission.Create {
		u, ok := attr.GetObject().(metav1.Object)
		if !ok {
			return fmt.Errorf("unexpected type %T", attr.GetObject())
		}
		if undo := SetClusterAnnotation(u, clusterName); undo != nil {
			defer undo()
		}
	}

	return plugin.Validate(ctx, attr, o)
}

func (p *Plugin) getHookSource(clusterName logicalcluster.Name, groupResource schema.GroupResource) (generic.Source, error) {
	clusterNameForGroupResource, err := p.getSourceClusterForGroupResource(clusterName, groupResource)
	if err != nil {
		return nil, err
	}

	p.managerLock.Lock()
	defer p.managerLock.Unlock()
	if _, ok := p.managersCache[clusterNameForGroupResource]; !ok {
		p.managersCache[clusterNameForGroupResource] = configuration.NewValidatingWebhookConfigurationManagerForInformer(
			p.globalKubeSharedInformerFactory.Admissionregistration().V1().ValidatingWebhookConfigurations().Cluster(clusterNameForGroupResource),
		)
	}

	return p.managersCache[clusterNameForGroupResource], nil
}

func (p *Plugin) getSourceClusterForGroupResource(clusterName logicalcluster.Name, groupResource schema.GroupResource) (logicalcluster.Name, error) {
	objs, err := p.getAPIBindings(clusterName)
	if err != nil {
		return "", err
	}

	for _, apiBinding := range objs {
		for _, br := range apiBinding.Status.BoundResources {
			if br.Group == groupResource.Group && br.Resource == groupResource.Resource {
				// GroupResource comes from an APIBinding/APIExport
				return logicalcluster.Name(apiBinding.Status.APIExportClusterName), nil
			}
		}
	}

	// GroupResource is local to this cluster
	return clusterName, nil
}

func (p *Plugin) ValidateInitialization() error {
	if p.kubeClusterClient == nil {
		return errors.New("missing kubeClusterClient")
	}
	if p.localKubeSharedInformerFactory == nil {
		return errors.New("missing localKubeSharedInformerFactory")
	}
	if p.globalKubeSharedInformerFactory == nil {
		return errors.New("missing globalKubeSharedInformerFactory")
	}
	return nil
}

func (p *Plugin) SetKubeClusterClient(client kcpkubernetesclientset.ClusterInterface) {
	p.kubeClusterClient = client
}

func (p *Plugin) SetKubeInformers(local, global kcpkubernetesinformers.SharedInformerFactory) {
	p.localKubeSharedInformerFactory = local
	p.globalKubeSharedInformerFactory = global
}

func (p *Plugin) SetKcpInformers(local, global kcpinformers.SharedInformerFactory) {
	p.getAPIBindings = func(clusterName logicalcluster.Name) ([]*apisv1alpha1.APIBinding, error) {
		return local.Apis().V1alpha1().APIBindings().Lister().Cluster(clusterName).List(labels.Everything())
	}
}

// SetClusterAnnotation sets the cluster annotation on the given object to the given clusterName,
// returning an undo function that can be used to revert the change.
func SetClusterAnnotation(obj metav1.Object, clusterName logicalcluster.Name) func() {
	undoFn := func() {
		anns := obj.GetAnnotations()
		delete(anns, logicalcluster.AnnotationKey)
		obj.SetAnnotations(anns)
	}

	anns := obj.GetAnnotations()
	if anns == nil {
		obj.SetAnnotations(map[string]string{logicalcluster.AnnotationKey: clusterName.String()})
		return undoFn
	}

	old, ok := anns[logicalcluster.AnnotationKey]
	if ok && old == clusterName.String() {
		return nil
	}

	anns[logicalcluster.AnnotationKey] = clusterName.String()
	obj.SetAnnotations(anns)
	return undoFn
}
