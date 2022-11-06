// Copyright 2019-2022 The Liqo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha1 "github.com/akaSomix/liqo/apis/virtualkubelet/v1alpha1"
	scheme "github.com/akaSomix/liqo/pkg/client/clientset/versioned/scheme"
)

// ShadowPodsGetter has a method to return a ShadowPodInterface.
// A group's client should implement this interface.
type ShadowPodsGetter interface {
	ShadowPods(namespace string) ShadowPodInterface
}

// ShadowPodInterface has methods to work with ShadowPod resources.
type ShadowPodInterface interface {
	Create(ctx context.Context, shadowPod *v1alpha1.ShadowPod, opts v1.CreateOptions) (*v1alpha1.ShadowPod, error)
	Update(ctx context.Context, shadowPod *v1alpha1.ShadowPod, opts v1.UpdateOptions) (*v1alpha1.ShadowPod, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ShadowPod, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ShadowPodList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ShadowPod, err error)
	ShadowPodExpansion
}

// shadowPods implements ShadowPodInterface
type shadowPods struct {
	client rest.Interface
	ns     string
}

// newShadowPods returns a ShadowPods
func newShadowPods(c *VirtualkubeletV1alpha1Client, namespace string) *shadowPods {
	return &shadowPods{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the shadowPod, and returns the corresponding shadowPod object, and an error if there is any.
func (c *shadowPods) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ShadowPod, err error) {
	result = &v1alpha1.ShadowPod{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("shadowpods").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ShadowPods that match those selectors.
func (c *shadowPods) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ShadowPodList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ShadowPodList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("shadowpods").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested shadowPods.
func (c *shadowPods) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("shadowpods").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a shadowPod and creates it.  Returns the server's representation of the shadowPod, and an error, if there is any.
func (c *shadowPods) Create(ctx context.Context, shadowPod *v1alpha1.ShadowPod, opts v1.CreateOptions) (result *v1alpha1.ShadowPod, err error) {
	result = &v1alpha1.ShadowPod{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("shadowpods").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(shadowPod).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a shadowPod and updates it. Returns the server's representation of the shadowPod, and an error, if there is any.
func (c *shadowPods) Update(ctx context.Context, shadowPod *v1alpha1.ShadowPod, opts v1.UpdateOptions) (result *v1alpha1.ShadowPod, err error) {
	result = &v1alpha1.ShadowPod{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("shadowpods").
		Name(shadowPod.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(shadowPod).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the shadowPod and deletes it. Returns an error if one occurs.
func (c *shadowPods) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("shadowpods").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *shadowPods) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("shadowpods").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched shadowPod.
func (c *shadowPods) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ShadowPod, err error) {
	result = &v1alpha1.ShadowPod{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("shadowpods").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
