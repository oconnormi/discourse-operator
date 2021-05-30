/*
Copyright 2021 oconnormi.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DiscourseSpec defines the desired state of Discourse
type DiscourseSpec struct {
	// AdminSecret contains the name of the secret that will be used to store admin credentials for the instance
	AdminSecret string `json:"adminSecret"`
	// AppSecret contains the name of the secret that will be used for any non-admin related secrets used by the instance
	AppSecret string `json:"appSecret"`
	// Registry allows overriding the default registry used for retrieving the discourse image
	Registry string `json:"registry,omitempty"`
	// Image allows overriding the default image used for running discourse
	Image string `json:"image,omitempty"`
	// Plugins holds a list of plugins to add to the instance
	Plugins Plugins `json:"plugins,omitempty"`
	// Database controls the database settings used by the instance. Allows for some prebuilt integrations, or manual connections
	Database Database `json:"database,omitempty"`
}

type Plugins struct {
	Registry  string   `json:"registry,omitempty"`
	Image     string   `json:"image,omitempty"`
	Installed []string `json:"installed,omitempty"`
}

// Database contains details for the specific database instance backing the discourse instance.
type Database struct {
	Enabled      bool     `json:"enabled"`
	Provider     Provider `json:"provider"`
	StorageClass string   `json:"storageClass"`
}

// Provider is a fixed list of supported datase providers. Currently crunchy or custom.
type Provider string

const (
	Crunchy Provider = "crunchy"
	Custom  Provider = "custom"
)

type InstanceState string

const (
	InstanceStateUnknown InstanceState = "Unknown"
	InstanceStateInit    InstanceState = "Initializing"
	InstanceStateReady   InstanceState = "Ready"
	InstanceStateError   InstanceState = "Error"
)

// DiscourseStatus defines the observed state of Discourse
type DiscourseStatus struct {
	State   InstanceState `json:"state"`
	Details string        `json:"details,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Discourse is the Schema for the discourses API
type Discourse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DiscourseSpec   `json:"spec,omitempty"`
	Status DiscourseStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DiscourseList contains a list of Discourse
type DiscourseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Discourse `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Discourse{}, &DiscourseList{})
}
