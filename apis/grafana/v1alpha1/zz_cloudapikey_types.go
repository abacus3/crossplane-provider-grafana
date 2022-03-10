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

type CloudAPIKeyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CloudAPIKeyParameters struct {

	// The slug of the organization to create the API key in. This is the same slug as the organization name in the URL.
	// +kubebuilder:validation:Required
	CloudOrgSlug *string `json:"cloudOrgSlug" tf:"cloud_org_slug,omitempty"`

	// Name of the API key.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Role of the API key. Should be one of [Viewer Editor Admin MetricsPublisher PluginPublisher]. See https://grafana.com/docs/grafana-cloud/api/#create-api-key for details.
	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// CloudAPIKeySpec defines the desired state of CloudAPIKey
type CloudAPIKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CloudAPIKeyParameters `json:"forProvider"`
}

// CloudAPIKeyStatus defines the observed state of CloudAPIKey.
type CloudAPIKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CloudAPIKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudAPIKey is the Schema for the CloudAPIKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafanajet}
type CloudAPIKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudAPIKeySpec   `json:"spec"`
	Status            CloudAPIKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudAPIKeyList contains a list of CloudAPIKeys
type CloudAPIKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudAPIKey `json:"items"`
}

// Repository type metadata.
var (
	CloudAPIKey_Kind             = "CloudAPIKey"
	CloudAPIKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CloudAPIKey_Kind}.String()
	CloudAPIKey_KindAPIVersion   = CloudAPIKey_Kind + "." + CRDGroupVersion.String()
	CloudAPIKey_GroupVersionKind = CRDGroupVersion.WithKind(CloudAPIKey_Kind)
)

func init() {
	SchemeBuilder.Register(&CloudAPIKey{}, &CloudAPIKeyList{})
}
