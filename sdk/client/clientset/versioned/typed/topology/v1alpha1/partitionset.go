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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/kcp-dev/kcp/sdk/apis/topology/v1alpha1"
	topologyv1alpha1 "github.com/kcp-dev/kcp/sdk/client/applyconfiguration/topology/v1alpha1"
	scheme "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// PartitionSetsGetter has a method to return a PartitionSetInterface.
// A group's client should implement this interface.
type PartitionSetsGetter interface {
	PartitionSets() PartitionSetInterface
}

// PartitionSetInterface has methods to work with PartitionSet resources.
type PartitionSetInterface interface {
	Create(ctx context.Context, partitionSet *v1alpha1.PartitionSet, opts v1.CreateOptions) (*v1alpha1.PartitionSet, error)
	Update(ctx context.Context, partitionSet *v1alpha1.PartitionSet, opts v1.UpdateOptions) (*v1alpha1.PartitionSet, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, partitionSet *v1alpha1.PartitionSet, opts v1.UpdateOptions) (*v1alpha1.PartitionSet, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.PartitionSet, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.PartitionSetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PartitionSet, err error)
	Apply(ctx context.Context, partitionSet *topologyv1alpha1.PartitionSetApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PartitionSet, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, partitionSet *topologyv1alpha1.PartitionSetApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PartitionSet, err error)
	PartitionSetExpansion
}

// partitionSets implements PartitionSetInterface
type partitionSets struct {
	*gentype.ClientWithListAndApply[*v1alpha1.PartitionSet, *v1alpha1.PartitionSetList, *topologyv1alpha1.PartitionSetApplyConfiguration]
}

// newPartitionSets returns a PartitionSets
func newPartitionSets(c *TopologyV1alpha1Client) *partitionSets {
	return &partitionSets{
		gentype.NewClientWithListAndApply[*v1alpha1.PartitionSet, *v1alpha1.PartitionSetList, *topologyv1alpha1.PartitionSetApplyConfiguration](
			"partitionsets",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1alpha1.PartitionSet { return &v1alpha1.PartitionSet{} },
			func() *v1alpha1.PartitionSetList { return &v1alpha1.PartitionSetList{} }),
	}
}
