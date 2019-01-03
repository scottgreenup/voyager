// Generated code
// run `make generate` to update

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/atlassian/voyager/pkg/apis/formation/v1"
	scheme "github.com/atlassian/voyager/pkg/formation/client/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LocationDescriptorsGetter has a method to return a LocationDescriptorInterface.
// A group's client should implement this interface.
type LocationDescriptorsGetter interface {
	LocationDescriptors(namespace string) LocationDescriptorInterface
}

// LocationDescriptorInterface has methods to work with LocationDescriptor resources.
type LocationDescriptorInterface interface {
	Create(*v1.LocationDescriptor) (*v1.LocationDescriptor, error)
	Update(*v1.LocationDescriptor) (*v1.LocationDescriptor, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.LocationDescriptor, error)
	List(opts metav1.ListOptions) (*v1.LocationDescriptorList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.LocationDescriptor, err error)
	LocationDescriptorExpansion
}

// locationDescriptors implements LocationDescriptorInterface
type locationDescriptors struct {
	client rest.Interface
	ns     string
}

// newLocationDescriptors returns a LocationDescriptors
func newLocationDescriptors(c *FormationV1Client, namespace string) *locationDescriptors {
	return &locationDescriptors{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the locationDescriptor, and returns the corresponding locationDescriptor object, and an error if there is any.
func (c *locationDescriptors) Get(name string, options metav1.GetOptions) (result *v1.LocationDescriptor, err error) {
	result = &v1.LocationDescriptor{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("locationdescriptors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LocationDescriptors that match those selectors.
func (c *locationDescriptors) List(opts metav1.ListOptions) (result *v1.LocationDescriptorList, err error) {
	result = &v1.LocationDescriptorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("locationdescriptors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested locationDescriptors.
func (c *locationDescriptors) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("locationdescriptors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a locationDescriptor and creates it.  Returns the server's representation of the locationDescriptor, and an error, if there is any.
func (c *locationDescriptors) Create(locationDescriptor *v1.LocationDescriptor) (result *v1.LocationDescriptor, err error) {
	result = &v1.LocationDescriptor{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("locationdescriptors").
		Body(locationDescriptor).
		Do().
		Into(result)
	return
}

// Update takes the representation of a locationDescriptor and updates it. Returns the server's representation of the locationDescriptor, and an error, if there is any.
func (c *locationDescriptors) Update(locationDescriptor *v1.LocationDescriptor) (result *v1.LocationDescriptor, err error) {
	result = &v1.LocationDescriptor{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("locationdescriptors").
		Name(locationDescriptor.Name).
		Body(locationDescriptor).
		Do().
		Into(result)
	return
}

// Delete takes name of the locationDescriptor and deletes it. Returns an error if one occurs.
func (c *locationDescriptors) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("locationdescriptors").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *locationDescriptors) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("locationdescriptors").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched locationDescriptor.
func (c *locationDescriptors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.LocationDescriptor, err error) {
	result = &v1.LocationDescriptor{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("locationdescriptors").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}