/*
author: nizhenyang
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "cloudedgenetwork/crd/api/cloudedgeservice/v1alpha1"
	scheme "cloudedgenetwork/crd/generated/clientset/versioned/scheme"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CloudEdgeServicesGetter has a method to return a CloudEdgeServiceInterface.
// A group's client should implement this interface.
type CloudEdgeServicesGetter interface {
	CloudEdgeServices(namespace string) CloudEdgeServiceInterface
}

// CloudEdgeServiceInterface has methods to work with CloudEdgeService resources.
type CloudEdgeServiceInterface interface {
	Create(*v1alpha1.CloudEdgeService) (*v1alpha1.CloudEdgeService, error)
	Update(*v1alpha1.CloudEdgeService) (*v1alpha1.CloudEdgeService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CloudEdgeService, error)
	List(opts v1.ListOptions) (*v1alpha1.CloudEdgeServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudEdgeService, err error)
	CloudEdgeServiceExpansion
}

// cloudEdgeServices implements CloudEdgeServiceInterface
type cloudEdgeServices struct {
	client rest.Interface
	ns     string
}

// newCloudEdgeServices returns a CloudEdgeServices
func newCloudEdgeServices(c *CloudedgeserviceV1alpha1Client, namespace string) *cloudEdgeServices {
	return &cloudEdgeServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cloudEdgeService, and returns the corresponding cloudEdgeService object, and an error if there is any.
func (c *cloudEdgeServices) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudEdgeService, err error) {
	result = &v1alpha1.CloudEdgeService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudedgeservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudEdgeServices that match those selectors.
func (c *cloudEdgeServices) List(opts v1.ListOptions) (result *v1alpha1.CloudEdgeServiceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CloudEdgeServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudedgeservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudEdgeServices.
func (c *cloudEdgeServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cloudedgeservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cloudEdgeService and creates it.  Returns the server's representation of the cloudEdgeService, and an error, if there is any.
func (c *cloudEdgeServices) Create(cloudEdgeService *v1alpha1.CloudEdgeService) (result *v1alpha1.CloudEdgeService, err error) {
	result = &v1alpha1.CloudEdgeService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cloudedgeservices").
		Body(cloudEdgeService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cloudEdgeService and updates it. Returns the server's representation of the cloudEdgeService, and an error, if there is any.
func (c *cloudEdgeServices) Update(cloudEdgeService *v1alpha1.CloudEdgeService) (result *v1alpha1.CloudEdgeService, err error) {
	result = &v1alpha1.CloudEdgeService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudedgeservices").
		Name(cloudEdgeService.Name).
		Body(cloudEdgeService).
		Do().
		Into(result)
	return
}

// Delete takes name of the cloudEdgeService and deletes it. Returns an error if one occurs.
func (c *cloudEdgeServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudedgeservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudEdgeServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudedgeservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cloudEdgeService.
func (c *cloudEdgeServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudEdgeService, err error) {
	result = &v1alpha1.CloudEdgeService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cloudedgeservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}