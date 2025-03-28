/*
Copyright The KCP Authors.

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


// Code generated by kcp code-generator. DO NOT EDIT.

package fake

import (
	"github.com/kcp-dev/logicalcluster/v3"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apisv1alpha1 "github.com/kcp-dev/kcp/sdk/apis/apis/v1alpha1"

	"k8s.io/apimachinery/pkg/watch"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/testing"

	"fmt"
	"encoding/json"

	"k8s.io/apimachinery/pkg/types"

	applyconfigurationsapisv1alpha1 "github.com/kcp-dev/kcp/sdk/client/applyconfiguration/apis/v1alpha1"

	apisv1alpha1client "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/typed/apis/v1alpha1"
)

var aPIBindingsResource = schema.GroupVersionResource{Group: "apis.kcp.io", Version: "v1alpha1", Resource: "apibindings"}
var aPIBindingsKind = schema.GroupVersionKind{Group: "apis.kcp.io", Version: "v1alpha1", Kind: "APIBinding"}

type aPIBindingsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *aPIBindingsClusterClient) Cluster(clusterPath logicalcluster.Path) apisv1alpha1client.APIBindingInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &aPIBindingsClient{Fake: c.Fake, ClusterPath: clusterPath}
}


// List takes label and field selectors, and returns the list of APIBindings that match those selectors across all clusters.
func (c *aPIBindingsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*apisv1alpha1.APIBindingList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(aPIBindingsResource, aPIBindingsKind, logicalcluster.Wildcard, opts), &apisv1alpha1.APIBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &apisv1alpha1.APIBindingList{ListMeta: obj.(*apisv1alpha1.APIBindingList).ListMeta}
	for _, item := range obj.(*apisv1alpha1.APIBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested APIBindings across all clusters.
func (c *aPIBindingsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(aPIBindingsResource, logicalcluster.Wildcard, opts))
}
type aPIBindingsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	
}


func (c *aPIBindingsClient) Create(ctx context.Context, aPIBinding *apisv1alpha1.APIBinding, opts metav1.CreateOptions) (*apisv1alpha1.APIBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(aPIBindingsResource, c.ClusterPath, aPIBinding), &apisv1alpha1.APIBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1alpha1.APIBinding), err
}

func (c *aPIBindingsClient) Update(ctx context.Context, aPIBinding *apisv1alpha1.APIBinding, opts metav1.UpdateOptions) (*apisv1alpha1.APIBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(aPIBindingsResource, c.ClusterPath, aPIBinding), &apisv1alpha1.APIBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1alpha1.APIBinding), err
}

func (c *aPIBindingsClient) UpdateStatus(ctx context.Context, aPIBinding *apisv1alpha1.APIBinding, opts metav1.UpdateOptions) (*apisv1alpha1.APIBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(aPIBindingsResource, c.ClusterPath, "status", aPIBinding), &apisv1alpha1.APIBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1alpha1.APIBinding), err
}

func (c *aPIBindingsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(aPIBindingsResource, c.ClusterPath, name, opts), &apisv1alpha1.APIBinding{})
	return err
}

func (c *aPIBindingsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(aPIBindingsResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &apisv1alpha1.APIBindingList{})
	return err
}

func (c *aPIBindingsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*apisv1alpha1.APIBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(aPIBindingsResource, c.ClusterPath, name), &apisv1alpha1.APIBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1alpha1.APIBinding), err
}

// List takes label and field selectors, and returns the list of APIBindings that match those selectors.
func (c *aPIBindingsClient) List(ctx context.Context, opts metav1.ListOptions) (*apisv1alpha1.APIBindingList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(aPIBindingsResource, aPIBindingsKind, c.ClusterPath, opts), &apisv1alpha1.APIBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &apisv1alpha1.APIBindingList{ListMeta: obj.(*apisv1alpha1.APIBindingList).ListMeta}
	for _, item := range obj.(*apisv1alpha1.APIBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *aPIBindingsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(aPIBindingsResource, c.ClusterPath, opts))
}

func (c *aPIBindingsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*apisv1alpha1.APIBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(aPIBindingsResource, c.ClusterPath, name, pt, data, subresources...), &apisv1alpha1.APIBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1alpha1.APIBinding), err
}

func (c *aPIBindingsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsapisv1alpha1.APIBindingApplyConfiguration, opts metav1.ApplyOptions) (*apisv1alpha1.APIBinding, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(aPIBindingsResource, c.ClusterPath, *name, types.ApplyPatchType, data), &apisv1alpha1.APIBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1alpha1.APIBinding), err
}

func (c *aPIBindingsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsapisv1alpha1.APIBindingApplyConfiguration, opts metav1.ApplyOptions) (*apisv1alpha1.APIBinding, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(aPIBindingsResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &apisv1alpha1.APIBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1alpha1.APIBinding), err
}
