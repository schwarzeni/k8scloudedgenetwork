/*
author: nizhenyang
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	cloudedgeservicev1alpha1 "cloudedgenetwork/crd/api/cloudedgeservice/v1alpha1"
	versioned "cloudedgenetwork/crd/generated/clientset/versioned"
	internalinterfaces "cloudedgenetwork/crd/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "cloudedgenetwork/crd/generated/listers/cloudedgeservice/v1alpha1"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CloudEdgeServiceInformer provides access to a shared informer and lister for
// CloudEdgeServices.
type CloudEdgeServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CloudEdgeServiceLister
}

type cloudEdgeServiceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCloudEdgeServiceInformer constructs a new informer for CloudEdgeService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCloudEdgeServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCloudEdgeServiceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCloudEdgeServiceInformer constructs a new informer for CloudEdgeService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCloudEdgeServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CloudedgeserviceV1alpha1().CloudEdgeServices(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CloudedgeserviceV1alpha1().CloudEdgeServices(namespace).Watch(options)
			},
		},
		&cloudedgeservicev1alpha1.CloudEdgeService{},
		resyncPeriod,
		indexers,
	)
}

func (f *cloudEdgeServiceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCloudEdgeServiceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cloudEdgeServiceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cloudedgeservicev1alpha1.CloudEdgeService{}, f.defaultInformer)
}

func (f *cloudEdgeServiceInformer) Lister() v1alpha1.CloudEdgeServiceLister {
	return v1alpha1.NewCloudEdgeServiceLister(f.Informer().GetIndexer())
}
