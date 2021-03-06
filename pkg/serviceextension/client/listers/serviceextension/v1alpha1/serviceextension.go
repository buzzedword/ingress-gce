/*
Copyright 2018 The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "k8s.io/ingress-gce/pkg/apis/serviceextension/v1alpha1"
)

// ServiceExtensionLister helps list ServiceExtensions.
type ServiceExtensionLister interface {
	// List lists all ServiceExtensions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceExtension, err error)
	// ServiceExtensions returns an object that can list and get ServiceExtensions.
	ServiceExtensions(namespace string) ServiceExtensionNamespaceLister
	ServiceExtensionListerExpansion
}

// serviceExtensionLister implements the ServiceExtensionLister interface.
type serviceExtensionLister struct {
	indexer cache.Indexer
}

// NewServiceExtensionLister returns a new ServiceExtensionLister.
func NewServiceExtensionLister(indexer cache.Indexer) ServiceExtensionLister {
	return &serviceExtensionLister{indexer: indexer}
}

// List lists all ServiceExtensions in the indexer.
func (s *serviceExtensionLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceExtension, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceExtension))
	})
	return ret, err
}

// ServiceExtensions returns an object that can list and get ServiceExtensions.
func (s *serviceExtensionLister) ServiceExtensions(namespace string) ServiceExtensionNamespaceLister {
	return serviceExtensionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServiceExtensionNamespaceLister helps list and get ServiceExtensions.
type ServiceExtensionNamespaceLister interface {
	// List lists all ServiceExtensions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceExtension, err error)
	// Get retrieves the ServiceExtension from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ServiceExtension, error)
	ServiceExtensionNamespaceListerExpansion
}

// serviceExtensionNamespaceLister implements the ServiceExtensionNamespaceLister
// interface.
type serviceExtensionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServiceExtensions in the indexer for a given namespace.
func (s serviceExtensionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceExtension, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceExtension))
	})
	return ret, err
}

// Get retrieves the ServiceExtension from the indexer for a given namespace and name.
func (s serviceExtensionNamespaceLister) Get(name string) (*v1alpha1.ServiceExtension, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("serviceextension"), name)
	}
	return obj.(*v1alpha1.ServiceExtension), nil
}
