/*
Copyright 2023.

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

package v1

import (
	kbatch "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// +kubebuilder:docs-gen:collapse=Imports
// +kubebuilder:validation:Enum=Allow;Forbid;Replace
type ConcurrencyPolicy string

const (
	// AllowConcurrent allows CronJobs to run concurrently.
	AllowConcurrent ConcurrencyPolicy = "Allow"

	// ForbidConcurrent forbids concurrent runs, skipping next run if previous
	// hasn't finished yet.
	ForbidConcurrent ConcurrencyPolicy = "Forbid"

	// ReplaceConcurrent cancels currently running job and replaces it with a new one.
	ReplaceConcurrent ConcurrencyPolicy = "Replace"
)

// CronJobSpec defines the desired state of CronJob
type CronJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Suspend                    *bool                  `json:"suspend,omitempty"`
	Schedule                   string                 `json:"schedule,omitempty"`
	StartingDeadlineSeconds    *int64                 `json:"startingDeadlineSeconds,omitempty"`
	FailedJobsHistoryLimit     *int32                 `json:"failedJobsHistoryLimit,omitempty"`
	SuccessfulJobsHistoryLimit *int32                 `json:"successfulJobsHistoryLimit,omitempty"`
	ConcurrencyPolicy          ConcurrencyPolicy      `json:"concurrencyPolicy,omitempty"`
	JobTemplate                kbatch.JobTemplateSpec `json:"jobTemplate,omitempty"`
	// Foo is an example field of CronJob. Edit cronjob_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// CronJobStatus defines the observed state of CronJob
type CronJobStatus struct {
	LastScheduleTime *metav1.Time             `json:"last_schedule_time,omitempty"`
	Active           []corev1.ObjectReference `json:"active,omitempty"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:docs-gen:collapse=old stuff

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CronJob is the Schema for the cronjobs API
type CronJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CronJobSpec   `json:"spec,omitempty"`
	Status CronJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CronJobList contains a list of CronJob
type CronJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CronJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CronJob{}, &CronJobList{})
}

// +kubebuilder:docs-gen:collapse=old stuff
