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

type TeamPreferencesObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TeamPreferencesParameters struct {

	// The numeric ID of the dashboard to display when a team member logs in.
	// +kubebuilder:validation:Optional
	HomeDashboardID *int64 `json:"homeDashboardId,omitempty" tf:"home_dashboard_id,omitempty"`

	// The numeric team ID.
	// +kubebuilder:validation:Required
	TeamID *int64 `json:"teamId" tf:"team_id,omitempty"`

	// The theme for the specified team. Available themes are `light`, `dark`, or an empty string for the default theme.
	// +kubebuilder:validation:Optional
	Theme *string `json:"theme,omitempty" tf:"theme,omitempty"`

	// The timezone for the specified team. Available values are `utc`, `browser`, or an empty string for the default.
	// +kubebuilder:validation:Optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

// TeamPreferencesSpec defines the desired state of TeamPreferences
type TeamPreferencesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TeamPreferencesParameters `json:"forProvider"`
}

// TeamPreferencesStatus defines the observed state of TeamPreferences.
type TeamPreferencesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TeamPreferencesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TeamPreferences is the Schema for the TeamPreferencess API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafanajet}
type TeamPreferences struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TeamPreferencesSpec   `json:"spec"`
	Status            TeamPreferencesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TeamPreferencesList contains a list of TeamPreferencess
type TeamPreferencesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TeamPreferences `json:"items"`
}

// Repository type metadata.
var (
	TeamPreferences_Kind             = "TeamPreferences"
	TeamPreferences_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TeamPreferences_Kind}.String()
	TeamPreferences_KindAPIVersion   = TeamPreferences_Kind + "." + CRDGroupVersion.String()
	TeamPreferences_GroupVersionKind = CRDGroupVersion.WithKind(TeamPreferences_Kind)
)

func init() {
	SchemeBuilder.Register(&TeamPreferences{}, &TeamPreferencesList{})
}