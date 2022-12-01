//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessGroup) DeepCopyInto(out *AccessGroup) {
	*out = *in
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(AccessGroupReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessGroup.
func (in *AccessGroup) DeepCopy() *AccessGroup {
	if in == nil {
		return nil
	}
	out := new(AccessGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessGroupReference) DeepCopyInto(out *AccessGroupReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessGroupReference.
func (in *AccessGroupReference) DeepCopy() *AccessGroupReference {
	if in == nil {
		return nil
	}
	out := new(AccessGroupReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudFlareAccessGroupRule) DeepCopyInto(out *CloudFlareAccessGroupRule) {
	*out = *in
	if in.Emails != nil {
		in, out := &in.Emails, &out.Emails
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EmailDomains != nil {
		in, out := &in.EmailDomains, &out.EmailDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IPRanges != nil {
		in, out := &in.IPRanges, &out.IPRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AccessGroups != nil {
		in, out := &in.AccessGroups, &out.AccessGroups
		*out = make([]AccessGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceToken != nil {
		in, out := &in.ServiceToken, &out.ServiceToken
		*out = make([]ServiceToken, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AnyAccessServiceToken != nil {
		in, out := &in.AnyAccessServiceToken, &out.AnyAccessServiceToken
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudFlareAccessGroupRule.
func (in *CloudFlareAccessGroupRule) DeepCopy() *CloudFlareAccessGroupRule {
	if in == nil {
		return nil
	}
	out := new(CloudFlareAccessGroupRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in CloudFlareAccessGroupRuleGroups) DeepCopyInto(out *CloudFlareAccessGroupRuleGroups) {
	{
		in := &in
		*out = make(CloudFlareAccessGroupRuleGroups, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make([]CloudFlareAccessGroupRule, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudFlareAccessGroupRuleGroups.
func (in CloudFlareAccessGroupRuleGroups) DeepCopy() CloudFlareAccessGroupRuleGroups {
	if in == nil {
		return nil
	}
	out := new(CloudFlareAccessGroupRuleGroups)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudflareAccessApplication) DeepCopyInto(out *CloudflareAccessApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudflareAccessApplication.
func (in *CloudflareAccessApplication) DeepCopy() *CloudflareAccessApplication {
	if in == nil {
		return nil
	}
	out := new(CloudflareAccessApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudflareAccessApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudflareAccessApplicationList) DeepCopyInto(out *CloudflareAccessApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudflareAccessApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudflareAccessApplicationList.
func (in *CloudflareAccessApplicationList) DeepCopy() *CloudflareAccessApplicationList {
	if in == nil {
		return nil
	}
	out := new(CloudflareAccessApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudflareAccessApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudflareAccessApplicationSpec) DeepCopyInto(out *CloudflareAccessApplicationSpec) {
	*out = *in
	if in.AppLauncherVisible != nil {
		in, out := &in.AppLauncherVisible, &out.AppLauncherVisible
		*out = new(bool)
		**out = **in
	}
	if in.AllowedIdps != nil {
		in, out := &in.AllowedIdps, &out.AllowedIdps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AutoRedirectToIdentity != nil {
		in, out := &in.AutoRedirectToIdentity, &out.AutoRedirectToIdentity
		*out = new(bool)
		**out = **in
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make(CloudflareAccessPolicyList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnableBindingCookie != nil {
		in, out := &in.EnableBindingCookie, &out.EnableBindingCookie
		*out = new(bool)
		**out = **in
	}
	if in.HTTPOnlyCookieAttribute != nil {
		in, out := &in.HTTPOnlyCookieAttribute, &out.HTTPOnlyCookieAttribute
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudflareAccessApplicationSpec.
func (in *CloudflareAccessApplicationSpec) DeepCopy() *CloudflareAccessApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(CloudflareAccessApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudflareAccessApplicationStatus) DeepCopyInto(out *CloudflareAccessApplicationStatus) {
	*out = *in
	in.CreatedAt.DeepCopyInto(&out.CreatedAt)
	in.UpdatedAt.DeepCopyInto(&out.UpdatedAt)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudflareAccessApplicationStatus.
func (in *CloudflareAccessApplicationStatus) DeepCopy() *CloudflareAccessApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(CloudflareAccessApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudflareAccessGroup) DeepCopyInto(out *CloudflareAccessGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudflareAccessGroup.
func (in *CloudflareAccessGroup) DeepCopy() *CloudflareAccessGroup {
	if in == nil {
		return nil
	}
	out := new(CloudflareAccessGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudflareAccessGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudflareAccessGroupList) DeepCopyInto(out *CloudflareAccessGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudflareAccessGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudflareAccessGroupList.
func (in *CloudflareAccessGroupList) DeepCopy() *CloudflareAccessGroupList {
	if in == nil {
		return nil
	}
	out := new(CloudflareAccessGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudflareAccessGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudflareAccessGroupSpec) DeepCopyInto(out *CloudflareAccessGroupSpec) {
	*out = *in
	if in.Include != nil {
		in, out := &in.Include, &out.Include
		*out = make([]CloudFlareAccessGroupRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Require != nil {
		in, out := &in.Require, &out.Require
		*out = make([]CloudFlareAccessGroupRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Exclude != nil {
		in, out := &in.Exclude, &out.Exclude
		*out = make([]CloudFlareAccessGroupRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudflareAccessGroupSpec.
func (in *CloudflareAccessGroupSpec) DeepCopy() *CloudflareAccessGroupSpec {
	if in == nil {
		return nil
	}
	out := new(CloudflareAccessGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudflareAccessGroupStatus) DeepCopyInto(out *CloudflareAccessGroupStatus) {
	*out = *in
	in.CreatedAt.DeepCopyInto(&out.CreatedAt)
	in.UpdatedAt.DeepCopyInto(&out.UpdatedAt)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudflareAccessGroupStatus.
func (in *CloudflareAccessGroupStatus) DeepCopy() *CloudflareAccessGroupStatus {
	if in == nil {
		return nil
	}
	out := new(CloudflareAccessGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudflareAccessPolicy) DeepCopyInto(out *CloudflareAccessPolicy) {
	*out = *in
	if in.Include != nil {
		in, out := &in.Include, &out.Include
		*out = make([]CloudFlareAccessGroupRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Require != nil {
		in, out := &in.Require, &out.Require
		*out = make([]CloudFlareAccessGroupRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Exclude != nil {
		in, out := &in.Exclude, &out.Exclude
		*out = make([]CloudFlareAccessGroupRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudflareAccessPolicy.
func (in *CloudflareAccessPolicy) DeepCopy() *CloudflareAccessPolicy {
	if in == nil {
		return nil
	}
	out := new(CloudflareAccessPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in CloudflareAccessPolicyList) DeepCopyInto(out *CloudflareAccessPolicyList) {
	{
		in := &in
		*out = make(CloudflareAccessPolicyList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudflareAccessPolicyList.
func (in CloudflareAccessPolicyList) DeepCopy() CloudflareAccessPolicyList {
	if in == nil {
		return nil
	}
	out := new(CloudflareAccessPolicyList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudflareServiceToken) DeepCopyInto(out *CloudflareServiceToken) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudflareServiceToken.
func (in *CloudflareServiceToken) DeepCopy() *CloudflareServiceToken {
	if in == nil {
		return nil
	}
	out := new(CloudflareServiceToken)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudflareServiceToken) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudflareServiceTokenList) DeepCopyInto(out *CloudflareServiceTokenList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudflareServiceToken, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudflareServiceTokenList.
func (in *CloudflareServiceTokenList) DeepCopy() *CloudflareServiceTokenList {
	if in == nil {
		return nil
	}
	out := new(CloudflareServiceTokenList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudflareServiceTokenList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudflareServiceTokenSpec) DeepCopyInto(out *CloudflareServiceTokenSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudflareServiceTokenSpec.
func (in *CloudflareServiceTokenSpec) DeepCopy() *CloudflareServiceTokenSpec {
	if in == nil {
		return nil
	}
	out := new(CloudflareServiceTokenSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudflareServiceTokenStatus) DeepCopyInto(out *CloudflareServiceTokenStatus) {
	*out = *in
	in.CreatedAt.DeepCopyInto(&out.CreatedAt)
	in.UpdatedAt.DeepCopyInto(&out.UpdatedAt)
	in.ExpiresAt.DeepCopyInto(&out.ExpiresAt)
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(SecretRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudflareServiceTokenStatus.
func (in *CloudflareServiceTokenStatus) DeepCopy() *CloudflareServiceTokenStatus {
	if in == nil {
		return nil
	}
	out := new(CloudflareServiceTokenStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretTemplateSpec) DeepCopyInto(out *SecretTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretTemplateSpec.
func (in *SecretTemplateSpec) DeepCopy() *SecretTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(SecretTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceToken) DeepCopyInto(out *ServiceToken) {
	*out = *in
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(ServiceTokenReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceToken.
func (in *ServiceToken) DeepCopy() *ServiceToken {
	if in == nil {
		return nil
	}
	out := new(ServiceToken)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceTokenReference) DeepCopyInto(out *ServiceTokenReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceTokenReference.
func (in *ServiceTokenReference) DeepCopy() *ServiceTokenReference {
	if in == nil {
		return nil
	}
	out := new(ServiceTokenReference)
	in.DeepCopyInto(out)
	return out
}
