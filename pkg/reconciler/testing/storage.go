/*
Copyright 2019 The Knative Authors

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

package testing

import (
	"time"

	gcpauthtesthelper "github.com/google/knative-gcp/pkg/apis/configs/gcpauth/testhelper"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/google/knative-gcp/pkg/apis/duck"
	"github.com/google/knative-gcp/pkg/apis/events/v1beta1"
	"github.com/google/knative-gcp/pkg/gclient/metadata/testing"
)

// CloudStorageSourceOption enables further configuration of a CloudStorageSource.
type CloudStorageSourceOption func(*v1beta1.CloudStorageSource)

// NewCloudStorageSource creates a CloudStorageSource with CloudStorageSourceOptions
func NewCloudStorageSource(name, namespace string, so ...CloudStorageSourceOption) *v1beta1.CloudStorageSource {
	s := &v1beta1.CloudStorageSource{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			UID:       "test-storage-uid",
			Annotations: map[string]string{
				duck.ClusterNameAnnotation: testing.FakeClusterName,
			},
		},
	}
	for _, opt := range so {
		opt(s)
	}
	return s
}

func WithCloudStorageSourceBucket(bucket string) CloudStorageSourceOption {
	return func(s *v1beta1.CloudStorageSource) {
		s.Spec.Bucket = bucket
	}
}

func WithCloudStorageSourceProject(project string) CloudStorageSourceOption {
	return func(s *v1beta1.CloudStorageSource) {
		s.Spec.Project = project
	}
}

func WithCloudStorageSourceEventTypes(eventTypes []string) CloudStorageSourceOption {
	return func(s *v1beta1.CloudStorageSource) {
		s.Spec.EventTypes = eventTypes
	}
}

func WithCloudStorageSourceSink(gvk metav1.GroupVersionKind, name string) CloudStorageSourceOption {
	return func(s *v1beta1.CloudStorageSource) {
		s.Spec.Sink = duckv1.Destination{
			Ref: &duckv1.KReference{
				APIVersion: apiVersion(gvk),
				Kind:       gvk.Kind,
				Name:       name,
			},
		}
	}
}

func WithCloudStorageSourceSinkDestination(sink duckv1.Destination) CloudStorageSourceOption {
	return func(s *v1beta1.CloudStorageSource) {
		s.Spec.Sink = sink
	}
}

// WithInitCloudStorageSourceConditions initializes the CloudStorageSources's conditions.
func WithInitCloudStorageSourceConditions(s *v1beta1.CloudStorageSource) {
	s.Status.InitializeConditions()
}

// WithCloudStorageSourceServiceAccountName will give status.ServiceAccountName a k8s service account name, which is related on Workload Identity's Google service account.
func WithCloudStorageSourceServiceAccountName(name string) CloudStorageSourceOption {
	return func(s *v1beta1.CloudStorageSource) {
		s.Status.ServiceAccountName = name
	}
}

func WithCloudStorageSourceWorkloadIdentityFailed(reason, message string) CloudStorageSourceOption {
	return func(s *v1beta1.CloudStorageSource) {
		s.Status.MarkWorkloadIdentityFailed(s.ConditionSet(), reason, message)
	}
}

func WithCloudStorageSourceServiceAccount(kServiceAccount string) CloudStorageSourceOption {
	return func(ps *v1beta1.CloudStorageSource) {
		ps.Spec.ServiceAccountName = kServiceAccount
	}
}

// WithCloudStorageSourceTopicFailed marks the condition that the
// topic is False
func WithCloudStorageSourceTopicFailed(reason, message string) CloudStorageSourceOption {
	return func(s *v1beta1.CloudStorageSource) {
		s.Status.MarkTopicFailed(s.ConditionSet(), reason, message)
	}
}

// WithCloudStorageSourceTopicUnknown marks the condition that the
// topic is False
func WithCloudStorageSourceTopicUnknown(reason, message string) CloudStorageSourceOption {
	return func(s *v1beta1.CloudStorageSource) {
		s.Status.MarkTopicUnknown(s.ConditionSet(), reason, message)
	}
}

// WithCloudStorageSourceTopicNotReady marks the condition that the
// topic is not ready
func WithCloudStorageSourceTopicReady(topicID string) CloudStorageSourceOption {
	return func(s *v1beta1.CloudStorageSource) {
		s.Status.MarkTopicReady(s.ConditionSet())
		s.Status.TopicID = topicID
	}
}

func WithCloudStorageSourceTopicID(topicID string) CloudStorageSourceOption {
	return func(s *v1beta1.CloudStorageSource) {
		s.Status.TopicID = topicID
	}
}

// WithCloudStorageSourcePullSubscriptionFailed marks the condition that the
// status of topic is False
func WithCloudStorageSourcePullSubscriptionFailed(reason, message string) CloudStorageSourceOption {
	return func(s *v1beta1.CloudStorageSource) {
		s.Status.MarkPullSubscriptionFailed(s.ConditionSet(), reason, message)
	}
}

// WithCloudStorageSourcePullSubscriptionUnknown marks the condition that the
// status of topic is Unknown.
func WithCloudStorageSourcePullSubscriptionUnknown(reason, message string) CloudStorageSourceOption {
	return func(s *v1beta1.CloudStorageSource) {
		s.Status.MarkPullSubscriptionUnknown(s.ConditionSet(), reason, message)
	}
}

// WithCloudStorageSourcePullSubscriptionReady marks the condition that the
// topic is ready.
func WithCloudStorageSourcePullSubscriptionReady() CloudStorageSourceOption {
	return func(s *v1beta1.CloudStorageSource) {
		s.Status.MarkPullSubscriptionReady(s.ConditionSet())
	}
}

// WithCloudStorageSourceNotificationNotReady marks the condition that the
// GCS Notification is not ready.
func WithCloudStorageSourceNotificationNotReady(reason, message string) CloudStorageSourceOption {
	return func(s *v1beta1.CloudStorageSource) {
		s.Status.MarkNotificationNotReady(reason, message)
	}
}

// WithCloudStorageSourceNotificationReady marks the condition that the GCS
// Notification is ready.
func WithCloudStorageSourceNotificationReady(notificationID string) CloudStorageSourceOption {
	return func(s *v1beta1.CloudStorageSource) {
		s.Status.MarkNotificationReady(notificationID)
	}
}

// WithCloudStorageSourceSinkURI sets the status for sink URI
func WithCloudStorageSourceSinkURI(url *apis.URL) CloudStorageSourceOption {
	return func(s *v1beta1.CloudStorageSource) {
		s.Status.SinkURI = url
	}
}

// WithCloudStorageSourceNotificationId sets the status for Notification ID
func WithCloudStorageSourceNotificationID(notificationID string) CloudStorageSourceOption {
	return func(s *v1beta1.CloudStorageSource) {
		s.Status.NotificationID = notificationID
	}
}

// WithCloudStorageSourceProjectId sets the status for Project ID
func WithCloudStorageSourceProjectID(projectID string) CloudStorageSourceOption {
	return func(s *v1beta1.CloudStorageSource) {
		s.Status.ProjectID = projectID
	}
}

func WithCloudStorageSourceSubscriptionID(subscriptionID string) CloudStorageSourceOption {
	return func(s *v1beta1.CloudStorageSource) {
		s.Status.SubscriptionID = subscriptionID
	}
}

func WithCloudStorageSourceStatusObservedGeneration(generation int64) CloudStorageSourceOption {
	return func(s *v1beta1.CloudStorageSource) {
		s.Status.Status.ObservedGeneration = generation
	}
}

func WithCloudStorageSourceObjectMetaGeneration(generation int64) CloudStorageSourceOption {
	return func(s *v1beta1.CloudStorageSource) {
		s.ObjectMeta.Generation = generation
	}
}

func WithDeletionTimestamp(s *v1beta1.CloudStorageSource) {
	ts := metav1.NewTime(time.Unix(1e9, 0))
	s.DeletionTimestamp = &ts
}

func WithCloudStorageSourceAnnotations(Annotations map[string]string) CloudStorageSourceOption {
	return func(s *v1beta1.CloudStorageSource) {
		s.ObjectMeta.Annotations = Annotations
	}
}

func WithCloudStorageSourceSetDefaults(s *v1beta1.CloudStorageSource) {
	s.SetDefaults(gcpauthtesthelper.ContextWithDefaults())
}
