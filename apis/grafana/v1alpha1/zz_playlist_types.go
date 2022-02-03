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

type ItemObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ItemParameters struct {

	// +kubebuilder:validation:Required
	Order *int64 `json:"order" tf:"order,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PlaylistObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`
}

type PlaylistParameters struct {

	// +kubebuilder:validation:Required
	Interval *string `json:"interval" tf:"interval,omitempty"`

	// +kubebuilder:validation:Required
	Item []ItemParameters `json:"item" tf:"item,omitempty"`
}

// PlaylistSpec defines the desired state of Playlist
type PlaylistSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PlaylistParameters `json:"forProvider"`
}

// PlaylistStatus defines the observed state of Playlist.
type PlaylistStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PlaylistObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Playlist is the Schema for the Playlists API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafanajet}
type Playlist struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PlaylistSpec   `json:"spec"`
	Status            PlaylistStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PlaylistList contains a list of Playlists
type PlaylistList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Playlist `json:"items"`
}

// Repository type metadata.
var (
	Playlist_Kind             = "Playlist"
	Playlist_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Playlist_Kind}.String()
	Playlist_KindAPIVersion   = Playlist_Kind + "." + CRDGroupVersion.String()
	Playlist_GroupVersionKind = CRDGroupVersion.WithKind(Playlist_Kind)
)

func init() {
	SchemeBuilder.Register(&Playlist{}, &PlaylistList{})
}