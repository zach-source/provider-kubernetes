//go:build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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

package config

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identity) DeepCopyInto(out *Identity) {
	*out = *in
	in.ProviderCredentials.DeepCopyInto(&out.ProviderCredentials)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identity.
func (in *Identity) DeepCopy() *Identity {
	if in == nil {
		return nil
	}
	out := new(Identity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderConfigSpec) DeepCopyInto(out *ProviderConfigSpec) {
	*out = *in
	in.Credentials.DeepCopyInto(&out.Credentials)
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(Identity)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderConfigSpec.
func (in *ProviderConfigSpec) DeepCopy() *ProviderConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ProviderConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderCredentials) DeepCopyInto(out *ProviderCredentials) {
	*out = *in
	in.CommonCredentialSelectors.DeepCopyInto(&out.CommonCredentialSelectors)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderCredentials.
func (in *ProviderCredentials) DeepCopy() *ProviderCredentials {
	if in == nil {
		return nil
	}
	out := new(ProviderCredentials)
	in.DeepCopyInto(out)
	return out
}