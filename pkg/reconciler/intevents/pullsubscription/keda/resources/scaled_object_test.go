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

package resources

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/knative-gcp/pkg/apis/duck"
	"github.com/google/knative-gcp/pkg/apis/intevents/v1beta1"
	"github.com/google/knative-gcp/pkg/reconciler/intevents/pullsubscription/resources"
	. "github.com/google/knative-gcp/pkg/reconciler/testing"

	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"knative.dev/pkg/apis"
)

func newAnnotations() map[string]string {
	return map[string]string{
		duck.AutoscalingClassAnnotation:                duck.KEDA,
		duck.AutoscalingMinScaleAnnotation:             "0",
		duck.AutoscalingMaxScaleAnnotation:             "3",
		duck.KedaAutoscalingSubscriptionSizeAnnotation: "5",
		duck.KedaAutoscalingCooldownPeriodAnnotation:   "60",
		duck.KedaAutoscalingPollingIntervalAnnotation:  "30",
	}
}

func newPullSubscription() *v1beta1.PullSubscription {
	return NewPullSubscription("psname", "psnamespace",
		WithPullSubscriptionUID("psuid"),
		WithPullSubscriptionAnnotations(newAnnotations()),
		WithPullSubscriptionSubscriptionID("subscriptionId"),
		WithPullSubscriptionSetDefaults,
	)
}

func newReceiveAdapter(ps *v1beta1.PullSubscription) *v1.Deployment {
	raArgs := &resources.ReceiveAdapterArgs{
		Image:            "image",
		PullSubscription: ps,
		Labels:           resources.GetLabels("agentName", "psName"),
		SubscriptionID:   "subscriptionId",
		SinkURI:          apis.HTTP("sinkURI"),
	}
	return resources.MakeReceiveAdapter(context.Background(), raArgs)
}

func TestMakeScaledObject(t *testing.T) {

	ps := newPullSubscription()
	ra := newReceiveAdapter(ps)

	want := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "keda.k8s.io/v1beta1",
			"kind":       "ScaledObject",
			"metadata": map[string]interface{}{
				"namespace": "psnamespace",
				"name":      GenerateScaledObjectName(ps),
				"labels": map[string]interface{}{
					"deploymentName":                  ra.Name,
					"events.cloud.google.com/ps-name": ps.Name,
				},
				"ownerReferences": []interface{}{
					map[string]interface{}{
						"apiVersion":         "internal.events.cloud.google.com/v1beta1",
						"kind":               "PullSubscription",
						"blockOwnerDeletion": true,
						"controller":         true,
						"name":               ps.Name,
						"uid":                string(ps.UID),
					}},
			},
			"spec": map[string]interface{}{
				"scaleTargetRef": map[string]interface{}{
					"deploymentName": ra.Name,
				},
				"minReplicaCount": int64(0),
				"maxReplicaCount": int64(3),
				"cooldownPeriod":  int64(60),
				"pollingInterval": int64(30),
				"triggers": []interface{}{
					map[string]interface{}{
						"type": "gcp-pubsub",
						"metadata": map[string]interface{}{
							"subscriptionSize": "5",
							"subscriptionName": "subscriptionId",
							"credentials":      "GOOGLE_APPLICATION_CREDENTIALS_JSON",
						},
					}},
			},
		},
	}

	got := MakeScaledObject(context.Background(), ra, ps)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("unexpected (-want, +got) = %v", diff)
	}
}
