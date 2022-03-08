/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type MonitoringProbeObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TenantID *float64 `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type MonitoringProbeParameters struct {

	// Custom labels to be included with collected metrics and logs.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Latitude coordinates.
	// +kubebuilder:validation:Required
	Latitude *float64 `json:"latitude" tf:"latitude,omitempty"`

	// Longitude coordinates.
	// +kubebuilder:validation:Required
	Longitude *float64 `json:"longitude" tf:"longitude,omitempty"`

	// Public probes are run by Grafana Labs and can be used by all users. Only Grafana Labs managed public probes will be set to `true`. Defaults to `false`.
	// +kubebuilder:validation:Optional
	Public *bool `json:"public,omitempty" tf:"public,omitempty"`

	// Region of the probe.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`
}

// MonitoringProbeSpec defines the desired state of MonitoringProbe
type MonitoringProbeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitoringProbeParameters `json:"forProvider"`
}

// MonitoringProbeStatus defines the observed state of MonitoringProbe.
type MonitoringProbeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitoringProbeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MonitoringProbe is the Schema for the MonitoringProbes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafanajet}
type MonitoringProbe struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitoringProbeSpec   `json:"spec"`
	Status            MonitoringProbeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitoringProbeList contains a list of MonitoringProbes
type MonitoringProbeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringProbe `json:"items"`
}

// Repository type metadata.
var (
	MonitoringProbe_Kind             = "MonitoringProbe"
	MonitoringProbe_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitoringProbe_Kind}.String()
	MonitoringProbe_KindAPIVersion   = MonitoringProbe_Kind + "." + CRDGroupVersion.String()
	MonitoringProbe_GroupVersionKind = CRDGroupVersion.WithKind(MonitoringProbe_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitoringProbe{}, &MonitoringProbeList{})
}