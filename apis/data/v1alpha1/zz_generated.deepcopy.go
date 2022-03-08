//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONDataObservation) DeepCopyInto(out *JSONDataObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONDataObservation.
func (in *JSONDataObservation) DeepCopy() *JSONDataObservation {
	if in == nil {
		return nil
	}
	out := new(JSONDataObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONDataParameters) DeepCopyInto(out *JSONDataParameters) {
	*out = *in
	if in.AssumeRoleArn != nil {
		in, out := &in.AssumeRoleArn, &out.AssumeRoleArn
		*out = new(string)
		**out = **in
	}
	if in.AuthType != nil {
		in, out := &in.AuthType, &out.AuthType
		*out = new(string)
		**out = **in
	}
	if in.AuthenticationType != nil {
		in, out := &in.AuthenticationType, &out.AuthenticationType
		*out = new(string)
		**out = **in
	}
	if in.ClientEmail != nil {
		in, out := &in.ClientEmail, &out.ClientEmail
		*out = new(string)
		**out = **in
	}
	if in.ConnMaxLifetime != nil {
		in, out := &in.ConnMaxLifetime, &out.ConnMaxLifetime
		*out = new(float64)
		**out = **in
	}
	if in.CustomMetricsNamespaces != nil {
		in, out := &in.CustomMetricsNamespaces, &out.CustomMetricsNamespaces
		*out = new(string)
		**out = **in
	}
	if in.DefaultProject != nil {
		in, out := &in.DefaultProject, &out.DefaultProject
		*out = new(string)
		**out = **in
	}
	if in.DefaultRegion != nil {
		in, out := &in.DefaultRegion, &out.DefaultRegion
		*out = new(string)
		**out = **in
	}
	if in.Encrypt != nil {
		in, out := &in.Encrypt, &out.Encrypt
		*out = new(string)
		**out = **in
	}
	if in.EsVersion != nil {
		in, out := &in.EsVersion, &out.EsVersion
		*out = new(string)
		**out = **in
	}
	if in.GraphiteVersion != nil {
		in, out := &in.GraphiteVersion, &out.GraphiteVersion
		*out = new(string)
		**out = **in
	}
	if in.HTTPMethod != nil {
		in, out := &in.HTTPMethod, &out.HTTPMethod
		*out = new(string)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(string)
		**out = **in
	}
	if in.LogLevelField != nil {
		in, out := &in.LogLevelField, &out.LogLevelField
		*out = new(string)
		**out = **in
	}
	if in.LogMessageField != nil {
		in, out := &in.LogMessageField, &out.LogMessageField
		*out = new(string)
		**out = **in
	}
	if in.MaxConcurrentShardRequests != nil {
		in, out := &in.MaxConcurrentShardRequests, &out.MaxConcurrentShardRequests
		*out = new(float64)
		**out = **in
	}
	if in.MaxIdleConns != nil {
		in, out := &in.MaxIdleConns, &out.MaxIdleConns
		*out = new(float64)
		**out = **in
	}
	if in.MaxOpenConns != nil {
		in, out := &in.MaxOpenConns, &out.MaxOpenConns
		*out = new(float64)
		**out = **in
	}
	if in.PostgresVersion != nil {
		in, out := &in.PostgresVersion, &out.PostgresVersion
		*out = new(float64)
		**out = **in
	}
	if in.Profile != nil {
		in, out := &in.Profile, &out.Profile
		*out = new(string)
		**out = **in
	}
	if in.QueryTimeout != nil {
		in, out := &in.QueryTimeout, &out.QueryTimeout
		*out = new(string)
		**out = **in
	}
	if in.SSLMode != nil {
		in, out := &in.SSLMode, &out.SSLMode
		*out = new(string)
		**out = **in
	}
	if in.Sigv4AssumeRoleArn != nil {
		in, out := &in.Sigv4AssumeRoleArn, &out.Sigv4AssumeRoleArn
		*out = new(string)
		**out = **in
	}
	if in.Sigv4Auth != nil {
		in, out := &in.Sigv4Auth, &out.Sigv4Auth
		*out = new(bool)
		**out = **in
	}
	if in.Sigv4AuthType != nil {
		in, out := &in.Sigv4AuthType, &out.Sigv4AuthType
		*out = new(string)
		**out = **in
	}
	if in.Sigv4ExternalID != nil {
		in, out := &in.Sigv4ExternalID, &out.Sigv4ExternalID
		*out = new(string)
		**out = **in
	}
	if in.Sigv4Profile != nil {
		in, out := &in.Sigv4Profile, &out.Sigv4Profile
		*out = new(string)
		**out = **in
	}
	if in.Sigv4Region != nil {
		in, out := &in.Sigv4Region, &out.Sigv4Region
		*out = new(string)
		**out = **in
	}
	if in.TLSAuth != nil {
		in, out := &in.TLSAuth, &out.TLSAuth
		*out = new(bool)
		**out = **in
	}
	if in.TLSAuthWithCACert != nil {
		in, out := &in.TLSAuthWithCACert, &out.TLSAuthWithCACert
		*out = new(bool)
		**out = **in
	}
	if in.TLSSkipVerify != nil {
		in, out := &in.TLSSkipVerify, &out.TLSSkipVerify
		*out = new(bool)
		**out = **in
	}
	if in.TimeField != nil {
		in, out := &in.TimeField, &out.TimeField
		*out = new(string)
		**out = **in
	}
	if in.TimeInterval != nil {
		in, out := &in.TimeInterval, &out.TimeInterval
		*out = new(string)
		**out = **in
	}
	if in.Timescaledb != nil {
		in, out := &in.Timescaledb, &out.Timescaledb
		*out = new(bool)
		**out = **in
	}
	if in.TokenURI != nil {
		in, out := &in.TokenURI, &out.TokenURI
		*out = new(string)
		**out = **in
	}
	if in.TsdbResolution != nil {
		in, out := &in.TsdbResolution, &out.TsdbResolution
		*out = new(string)
		**out = **in
	}
	if in.TsdbVersion != nil {
		in, out := &in.TsdbVersion, &out.TsdbVersion
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONDataParameters.
func (in *JSONDataParameters) DeepCopy() *JSONDataParameters {
	if in == nil {
		return nil
	}
	out := new(JSONDataParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionsObservation) DeepCopyInto(out *PermissionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionsObservation.
func (in *PermissionsObservation) DeepCopy() *PermissionsObservation {
	if in == nil {
		return nil
	}
	out := new(PermissionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionsParameters) DeepCopyInto(out *PermissionsParameters) {
	*out = *in
	if in.Permission != nil {
		in, out := &in.Permission, &out.Permission
		*out = new(string)
		**out = **in
	}
	if in.TeamID != nil {
		in, out := &in.TeamID, &out.TeamID
		*out = new(float64)
		**out = **in
	}
	if in.UserID != nil {
		in, out := &in.UserID, &out.UserID
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionsParameters.
func (in *PermissionsParameters) DeepCopy() *PermissionsParameters {
	if in == nil {
		return nil
	}
	out := new(PermissionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecureJSONDataObservation) DeepCopyInto(out *SecureJSONDataObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecureJSONDataObservation.
func (in *SecureJSONDataObservation) DeepCopy() *SecureJSONDataObservation {
	if in == nil {
		return nil
	}
	out := new(SecureJSONDataObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecureJSONDataParameters) DeepCopyInto(out *SecureJSONDataParameters) {
	*out = *in
	if in.AccessKeySecretRef != nil {
		in, out := &in.AccessKeySecretRef, &out.AccessKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.BasicAuthPasswordSecretRef != nil {
		in, out := &in.BasicAuthPasswordSecretRef, &out.BasicAuthPasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.PasswordSecretRef != nil {
		in, out := &in.PasswordSecretRef, &out.PasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.PrivateKeySecretRef != nil {
		in, out := &in.PrivateKeySecretRef, &out.PrivateKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.SecretKeySecretRef != nil {
		in, out := &in.SecretKeySecretRef, &out.SecretKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Sigv4AccessKeySecretRef != nil {
		in, out := &in.Sigv4AccessKeySecretRef, &out.Sigv4AccessKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Sigv4SecretKeySecretRef != nil {
		in, out := &in.Sigv4SecretKeySecretRef, &out.Sigv4SecretKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.TLSCACertSecretRef != nil {
		in, out := &in.TLSCACertSecretRef, &out.TLSCACertSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.TLSClientCertSecretRef != nil {
		in, out := &in.TLSClientCertSecretRef, &out.TLSClientCertSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.TLSClientKeySecretRef != nil {
		in, out := &in.TLSClientKeySecretRef, &out.TLSClientKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecureJSONDataParameters.
func (in *SecureJSONDataParameters) DeepCopy() *SecureJSONDataParameters {
	if in == nil {
		return nil
	}
	out := new(SecureJSONDataParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Source) DeepCopyInto(out *Source) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Source.
func (in *Source) DeepCopy() *Source {
	if in == nil {
		return nil
	}
	out := new(Source)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Source) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceList) DeepCopyInto(out *SourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Source, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceList.
func (in *SourceList) DeepCopy() *SourceList {
	if in == nil {
		return nil
	}
	out := new(SourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceObservation) DeepCopyInto(out *SourceObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceObservation.
func (in *SourceObservation) DeepCopy() *SourceObservation {
	if in == nil {
		return nil
	}
	out := new(SourceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceParameters) DeepCopyInto(out *SourceParameters) {
	*out = *in
	if in.AccessMode != nil {
		in, out := &in.AccessMode, &out.AccessMode
		*out = new(string)
		**out = **in
	}
	if in.BasicAuthEnabled != nil {
		in, out := &in.BasicAuthEnabled, &out.BasicAuthEnabled
		*out = new(bool)
		**out = **in
	}
	if in.BasicAuthPasswordSecretRef != nil {
		in, out := &in.BasicAuthPasswordSecretRef, &out.BasicAuthPasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.BasicAuthUsername != nil {
		in, out := &in.BasicAuthUsername, &out.BasicAuthUsername
		*out = new(string)
		**out = **in
	}
	if in.DatabaseName != nil {
		in, out := &in.DatabaseName, &out.DatabaseName
		*out = new(string)
		**out = **in
	}
	if in.HTTPHeadersSecretRef != nil {
		in, out := &in.HTTPHeadersSecretRef, &out.HTTPHeadersSecretRef
		*out = make(map[string]v1.SecretKeySelector, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.IsDefault != nil {
		in, out := &in.IsDefault, &out.IsDefault
		*out = new(bool)
		**out = **in
	}
	if in.JSONData != nil {
		in, out := &in.JSONData, &out.JSONData
		*out = make([]JSONDataParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PasswordSecretRef != nil {
		in, out := &in.PasswordSecretRef, &out.PasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.SecureJSONData != nil {
		in, out := &in.SecureJSONData, &out.SecureJSONData
		*out = make([]SecureJSONDataParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.UID != nil {
		in, out := &in.UID, &out.UID
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceParameters.
func (in *SourceParameters) DeepCopy() *SourceParameters {
	if in == nil {
		return nil
	}
	out := new(SourceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourcePermission) DeepCopyInto(out *SourcePermission) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourcePermission.
func (in *SourcePermission) DeepCopy() *SourcePermission {
	if in == nil {
		return nil
	}
	out := new(SourcePermission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SourcePermission) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourcePermissionList) DeepCopyInto(out *SourcePermissionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SourcePermission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourcePermissionList.
func (in *SourcePermissionList) DeepCopy() *SourcePermissionList {
	if in == nil {
		return nil
	}
	out := new(SourcePermissionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SourcePermissionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourcePermissionObservation) DeepCopyInto(out *SourcePermissionObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourcePermissionObservation.
func (in *SourcePermissionObservation) DeepCopy() *SourcePermissionObservation {
	if in == nil {
		return nil
	}
	out := new(SourcePermissionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourcePermissionParameters) DeepCopyInto(out *SourcePermissionParameters) {
	*out = *in
	if in.DatasourceID != nil {
		in, out := &in.DatasourceID, &out.DatasourceID
		*out = new(float64)
		**out = **in
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]PermissionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourcePermissionParameters.
func (in *SourcePermissionParameters) DeepCopy() *SourcePermissionParameters {
	if in == nil {
		return nil
	}
	out := new(SourcePermissionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourcePermissionSpec) DeepCopyInto(out *SourcePermissionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourcePermissionSpec.
func (in *SourcePermissionSpec) DeepCopy() *SourcePermissionSpec {
	if in == nil {
		return nil
	}
	out := new(SourcePermissionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourcePermissionStatus) DeepCopyInto(out *SourcePermissionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourcePermissionStatus.
func (in *SourcePermissionStatus) DeepCopy() *SourcePermissionStatus {
	if in == nil {
		return nil
	}
	out := new(SourcePermissionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceSpec) DeepCopyInto(out *SourceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceSpec.
func (in *SourceSpec) DeepCopy() *SourceSpec {
	if in == nil {
		return nil
	}
	out := new(SourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceStatus) DeepCopyInto(out *SourceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceStatus.
func (in *SourceStatus) DeepCopy() *SourceStatus {
	if in == nil {
		return nil
	}
	out := new(SourceStatus)
	in.DeepCopyInto(out)
	return out
}