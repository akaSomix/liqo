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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v1alpha1 "github.com/akaSomix/liqo/apis/virtualkubelet/v1alpha1"
)

// ShadowPodLister helps list ShadowPods.
// All objects returned here must be treated as read-only.
type ShadowPodLister interface {
	// List lists all ShadowPods in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ShadowPod, err error)
	// ShadowPods returns an object that can list and get ShadowPods.
	ShadowPods(namespace string) ShadowPodNamespaceLister
	ShadowPodListerExpansion
}

// shadowPodLister implements the ShadowPodLister interface.
type shadowPodLister struct {
	indexer cache.Indexer
}

// NewShadowPodLister returns a new ShadowPodLister.
func NewShadowPodLister(indexer cache.Indexer) ShadowPodLister {
	return &shadowPodLister{indexer: indexer}
}

// List lists all ShadowPods in the indexer.
func (s *shadowPodLister) List(selector labels.Selector) (ret []*v1alpha1.ShadowPod, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ShadowPod))
	})
	return ret, err
}

// ShadowPods returns an object that can list and get ShadowPods.
func (s *shadowPodLister) ShadowPods(namespace string) ShadowPodNamespaceLister {
	return shadowPodNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ShadowPodNamespaceLister helps list and get ShadowPods.
// All objects returned here must be treated as read-only.
type ShadowPodNamespaceLister interface {
	// List lists all ShadowPods in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ShadowPod, err error)
	// Get retrieves the ShadowPod from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ShadowPod, error)
	ShadowPodNamespaceListerExpansion
}

// shadowPodNamespaceLister implements the ShadowPodNamespaceLister
// interface.
type shadowPodNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ShadowPods in the indexer for a given namespace.
func (s shadowPodNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ShadowPod, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ShadowPod))
	})
	return ret, err
}

// Get retrieves the ShadowPod from the indexer for a given namespace and name.
func (s shadowPodNamespaceLister) Get(name string) (*v1alpha1.ShadowPod, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("shadowpod"), name)
	}
	return obj.(*v1alpha1.ShadowPod), nil
}
