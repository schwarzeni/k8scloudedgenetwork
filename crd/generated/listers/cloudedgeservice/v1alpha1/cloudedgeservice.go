/*
author: nizhenyang
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "cloudedgenetwork/crd/api/cloudedgeservice/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CloudEdgeServiceLister helps list CloudEdgeServices.
type CloudEdgeServiceLister interface {
	// List lists all CloudEdgeServices in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CloudEdgeService, err error)
	// CloudEdgeServices returns an object that can list and get CloudEdgeServices.
	CloudEdgeServices(namespace string) CloudEdgeServiceNamespaceLister
	CloudEdgeServiceListerExpansion
}

// cloudEdgeServiceLister implements the CloudEdgeServiceLister interface.
type cloudEdgeServiceLister struct {
	indexer cache.Indexer
}

// NewCloudEdgeServiceLister returns a new CloudEdgeServiceLister.
func NewCloudEdgeServiceLister(indexer cache.Indexer) CloudEdgeServiceLister {
	return &cloudEdgeServiceLister{indexer: indexer}
}

// List lists all CloudEdgeServices in the indexer.
func (s *cloudEdgeServiceLister) List(selector labels.Selector) (ret []*v1alpha1.CloudEdgeService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudEdgeService))
	})
	return ret, err
}

// CloudEdgeServices returns an object that can list and get CloudEdgeServices.
func (s *cloudEdgeServiceLister) CloudEdgeServices(namespace string) CloudEdgeServiceNamespaceLister {
	return cloudEdgeServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CloudEdgeServiceNamespaceLister helps list and get CloudEdgeServices.
type CloudEdgeServiceNamespaceLister interface {
	// List lists all CloudEdgeServices in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CloudEdgeService, err error)
	// Get retrieves the CloudEdgeService from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CloudEdgeService, error)
	CloudEdgeServiceNamespaceListerExpansion
}

// cloudEdgeServiceNamespaceLister implements the CloudEdgeServiceNamespaceLister
// interface.
type cloudEdgeServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CloudEdgeServices in the indexer for a given namespace.
func (s cloudEdgeServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CloudEdgeService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudEdgeService))
	})
	return ret, err
}

// Get retrieves the CloudEdgeService from the indexer for a given namespace and name.
func (s cloudEdgeServiceNamespaceLister) Get(name string) (*v1alpha1.CloudEdgeService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cloudedgeservice"), name)
	}
	return obj.(*v1alpha1.CloudEdgeService), nil
}