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

package v1alpha1

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v3"

	apisv1alpha1 "github.com/kcp-dev/kcp/sdk/apis/apis/v1alpha1"
	apisv1alpha1listers "github.com/kcp-dev/kcp/sdk/client/listers/apis/v1alpha1"
	clientset "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/cluster"
	scopedclientset "github.com/kcp-dev/kcp/sdk/client/clientset/versioned"
	"github.com/kcp-dev/kcp/sdk/client/informers/externalversions/internalinterfaces"
)

// APIConversionClusterInformer provides access to a shared informer and lister for
// APIConversions.
type APIConversionClusterInformer interface {
	Cluster(logicalcluster.Name) APIConversionInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() apisv1alpha1listers.APIConversionClusterLister
}

type aPIConversionClusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewAPIConversionClusterInformer constructs a new informer for APIConversion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAPIConversionClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredAPIConversionClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredAPIConversionClusterInformer constructs a new informer for APIConversion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAPIConversionClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApisV1alpha1().APIConversions().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApisV1alpha1().APIConversions().Watch(context.TODO(), options)
			},
		},
		&apisv1alpha1.APIConversion{},
		resyncPeriod,
		indexers,
	)
}

func (f *aPIConversionClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredAPIConversionClusterInformer(client, resyncPeriod, cache.Indexers{
			kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
			}, 
		f.tweakListOptions,
	)
}

func (f *aPIConversionClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&apisv1alpha1.APIConversion{}, f.defaultInformer)
}

func (f *aPIConversionClusterInformer) Lister() apisv1alpha1listers.APIConversionClusterLister {
	return apisv1alpha1listers.NewAPIConversionClusterLister(f.Informer().GetIndexer())
}


// APIConversionInformer provides access to a shared informer and lister for
// APIConversions.
type APIConversionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() apisv1alpha1listers.APIConversionLister
}
func (f *aPIConversionClusterInformer) Cluster(clusterName logicalcluster.Name) APIConversionInformer {
	return &aPIConversionInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().Cluster(clusterName),
	}
}

type aPIConversionInformer struct {
	informer cache.SharedIndexInformer
	lister apisv1alpha1listers.APIConversionLister
}

func (f *aPIConversionInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *aPIConversionInformer) Lister() apisv1alpha1listers.APIConversionLister {
	return f.lister
}

type aPIConversionScopedInformer struct {
	factory internalinterfaces.SharedScopedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	}

func (f *aPIConversionScopedInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisv1alpha1.APIConversion{}, f.defaultInformer)
}

func (f *aPIConversionScopedInformer) Lister() apisv1alpha1listers.APIConversionLister {
	return apisv1alpha1listers.NewAPIConversionLister(f.Informer().GetIndexer())
}

// NewAPIConversionInformer constructs a new informer for APIConversion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAPIConversionInformer(client scopedclientset.Interface, resyncPeriod time.Duration,indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAPIConversionInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredAPIConversionInformer constructs a new informer for APIConversion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAPIConversionInformer(client scopedclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApisV1alpha1().APIConversions().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApisV1alpha1().APIConversions().Watch(context.TODO(), options)
			},
		},
		&apisv1alpha1.APIConversion{},
		resyncPeriod,
		indexers,
	)
}

func (f *aPIConversionScopedInformer) defaultInformer(client scopedclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAPIConversionInformer(client, resyncPeriod,cache.Indexers{ }, f.tweakListOptions)
}

