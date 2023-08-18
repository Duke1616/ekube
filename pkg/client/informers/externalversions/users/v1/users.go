/*
Copyright 2020 The KubeSphere Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	usersv1 "ekube/api/k8s/users/v1"
	versioned "ekube/pkg/client/clientset/versioned"
	internalinterfaces "ekube/pkg/client/informers/externalversions/internalinterfaces"
	v1 "ekube/pkg/client/listers/users/v1"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// UsersInformer provides access to a shared informer and lister for
// Userses.
type UsersInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.UsersLister
}

type usersInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewUsersInformer constructs a new informer for Users type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewUsersInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredUsersInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredUsersInformer constructs a new informer for Users type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredUsersInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.UsersV1().Userses().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.UsersV1().Userses().Watch(context.TODO(), options)
			},
		},
		&usersv1.Users{},
		resyncPeriod,
		indexers,
	)
}

func (f *usersInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredUsersInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *usersInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&usersv1.Users{}, f.defaultInformer)
}

func (f *usersInformer) Lister() v1.UsersLister {
	return v1.NewUsersLister(f.Informer().GetIndexer())
}
