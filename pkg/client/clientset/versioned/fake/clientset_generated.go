/*
Copyright 2020 Google LLC

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "github.com/google/knative-gcp/pkg/client/clientset/versioned"
	eventingv1beta1 "github.com/google/knative-gcp/pkg/client/clientset/versioned/typed/broker/v1beta1"
	fakeeventingv1beta1 "github.com/google/knative-gcp/pkg/client/clientset/versioned/typed/broker/v1beta1/fake"
	eventsv1 "github.com/google/knative-gcp/pkg/client/clientset/versioned/typed/events/v1"
	fakeeventsv1 "github.com/google/knative-gcp/pkg/client/clientset/versioned/typed/events/v1/fake"
	eventsv1alpha1 "github.com/google/knative-gcp/pkg/client/clientset/versioned/typed/events/v1alpha1"
	fakeeventsv1alpha1 "github.com/google/knative-gcp/pkg/client/clientset/versioned/typed/events/v1alpha1/fake"
	eventsv1beta1 "github.com/google/knative-gcp/pkg/client/clientset/versioned/typed/events/v1beta1"
	fakeeventsv1beta1 "github.com/google/knative-gcp/pkg/client/clientset/versioned/typed/events/v1beta1/fake"
	internalv1 "github.com/google/knative-gcp/pkg/client/clientset/versioned/typed/intevents/v1"
	fakeinternalv1 "github.com/google/knative-gcp/pkg/client/clientset/versioned/typed/intevents/v1/fake"
	internalv1alpha1 "github.com/google/knative-gcp/pkg/client/clientset/versioned/typed/intevents/v1alpha1"
	fakeinternalv1alpha1 "github.com/google/knative-gcp/pkg/client/clientset/versioned/typed/intevents/v1alpha1/fake"
	internalv1beta1 "github.com/google/knative-gcp/pkg/client/clientset/versioned/typed/intevents/v1beta1"
	fakeinternalv1beta1 "github.com/google/knative-gcp/pkg/client/clientset/versioned/typed/intevents/v1beta1/fake"
	messagingv1alpha1 "github.com/google/knative-gcp/pkg/client/clientset/versioned/typed/messaging/v1alpha1"
	fakemessagingv1alpha1 "github.com/google/knative-gcp/pkg/client/clientset/versioned/typed/messaging/v1alpha1/fake"
	messagingv1beta1 "github.com/google/knative-gcp/pkg/client/clientset/versioned/typed/messaging/v1beta1"
	fakemessagingv1beta1 "github.com/google/knative-gcp/pkg/client/clientset/versioned/typed/messaging/v1beta1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var _ clientset.Interface = &Clientset{}

// EventingV1beta1 retrieves the EventingV1beta1Client
func (c *Clientset) EventingV1beta1() eventingv1beta1.EventingV1beta1Interface {
	return &fakeeventingv1beta1.FakeEventingV1beta1{Fake: &c.Fake}
}

// EventsV1alpha1 retrieves the EventsV1alpha1Client
func (c *Clientset) EventsV1alpha1() eventsv1alpha1.EventsV1alpha1Interface {
	return &fakeeventsv1alpha1.FakeEventsV1alpha1{Fake: &c.Fake}
}

// EventsV1beta1 retrieves the EventsV1beta1Client
func (c *Clientset) EventsV1beta1() eventsv1beta1.EventsV1beta1Interface {
	return &fakeeventsv1beta1.FakeEventsV1beta1{Fake: &c.Fake}
}

// EventsV1 retrieves the EventsV1Client
func (c *Clientset) EventsV1() eventsv1.EventsV1Interface {
	return &fakeeventsv1.FakeEventsV1{Fake: &c.Fake}
}

// InternalV1alpha1 retrieves the InternalV1alpha1Client
func (c *Clientset) InternalV1alpha1() internalv1alpha1.InternalV1alpha1Interface {
	return &fakeinternalv1alpha1.FakeInternalV1alpha1{Fake: &c.Fake}
}

// InternalV1beta1 retrieves the InternalV1beta1Client
func (c *Clientset) InternalV1beta1() internalv1beta1.InternalV1beta1Interface {
	return &fakeinternalv1beta1.FakeInternalV1beta1{Fake: &c.Fake}
}

// InternalV1 retrieves the InternalV1Client
func (c *Clientset) InternalV1() internalv1.InternalV1Interface {
	return &fakeinternalv1.FakeInternalV1{Fake: &c.Fake}
}

// MessagingV1alpha1 retrieves the MessagingV1alpha1Client
func (c *Clientset) MessagingV1alpha1() messagingv1alpha1.MessagingV1alpha1Interface {
	return &fakemessagingv1alpha1.FakeMessagingV1alpha1{Fake: &c.Fake}
}

// MessagingV1beta1 retrieves the MessagingV1beta1Client
func (c *Clientset) MessagingV1beta1() messagingv1beta1.MessagingV1beta1Interface {
	return &fakemessagingv1beta1.FakeMessagingV1beta1{Fake: &c.Fake}
}
