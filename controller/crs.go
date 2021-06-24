// Copyright 2019 HAProxy Technologies LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"k8s.io/client-go/tools/cache"

	"github.com/haproxytech/kubernetes-ingress/controller/store"
	corev1alpha1 "github.com/haproxytech/kubernetes-ingress/crs/api/core/v1alpha1"
	informers "github.com/haproxytech/kubernetes-ingress/crs/generated/informers/externalversions"
)

type GlobalCR struct {
}

func NewGlobalCR() GlobalCR {
	return GlobalCR{}
}

func (c GlobalCR) GetKind() string {
	return "Global"
}

func (c GlobalCR) GetInformer(eventChan chan SyncDataEvent, factory informers.SharedInformerFactory) cache.SharedIndexInformer {
	informer := factory.Core().V1alpha1().Globals().Informer()

	sendToChannel := func(eventChan chan SyncDataEvent, object interface{}, status store.Status) {
		data := object.(*corev1alpha1.Global)
		logger.Debugf("%s %s: %s", data.GetNamespace(), status, data.GetName())
		if status == DELETED {
			eventChan <- SyncDataEvent{SyncType: CUSTOM_RESOURCE, CRKind: c.GetKind(), Namespace: data.GetNamespace(), Name: data.GetName(), Data: nil}
			return
		}
		eventChan <- SyncDataEvent{SyncType: CUSTOM_RESOURCE, CRKind: c.GetKind(), Namespace: data.GetNamespace(), Name: data.GetName(), Data: data}
	}

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			sendToChannel(eventChan, obj, store.ADDED)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			sendToChannel(eventChan, newObj, store.MODIFIED)
		},
		DeleteFunc: func(obj interface{}) {
			sendToChannel(eventChan, obj, store.DELETED)
		},
	})
	return informer
}

func (c GlobalCR) ProcessEvent(s *store.K8s, job SyncDataEvent) bool {
	if job.Data == nil {
		s.CR.Global = nil
		return true
	}
	data, ok := job.Data.(*corev1alpha1.Global)
	if !ok {
		logger.Warning(CoreGroupVersion + ": type mismatch with Global kind")
		return false
	}
	s.CR.Global = data.Spec.Config
	return true
}
