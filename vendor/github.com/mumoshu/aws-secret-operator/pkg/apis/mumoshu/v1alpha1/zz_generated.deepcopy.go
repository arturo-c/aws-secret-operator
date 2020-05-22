// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSSecret) DeepCopyInto(out *AWSSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSSecret.
func (in *AWSSecret) DeepCopy() *AWSSecret {
	if in == nil {
		return nil
	}
	out := new(AWSSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSSecretList) DeepCopyInto(out *AWSSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSSecretList.
func (in *AWSSecretList) DeepCopy() *AWSSecretList {
	if in == nil {
		return nil
	}
	out := new(AWSSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSSecretSpec) DeepCopyInto(out *AWSSecretSpec) {
	*out = *in
	out.StringDataFrom = in.StringDataFrom
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSSecretSpec.
func (in *AWSSecretSpec) DeepCopy() *AWSSecretSpec {
	if in == nil {
		return nil
	}
	out := new(AWSSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSSecretStatus) DeepCopyInto(out *AWSSecretStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSSecretStatus.
func (in *AWSSecretStatus) DeepCopy() *AWSSecretStatus {
	if in == nil {
		return nil
	}
	out := new(AWSSecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretsManagerSecretRef) DeepCopyInto(out *SecretsManagerSecretRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretsManagerSecretRef.
func (in *SecretsManagerSecretRef) DeepCopy() *SecretsManagerSecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretsManagerSecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StringDataFrom) DeepCopyInto(out *StringDataFrom) {
	*out = *in
	out.SecretsManagerSecretRef = in.SecretsManagerSecretRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StringDataFrom.
func (in *StringDataFrom) DeepCopy() *StringDataFrom {
	if in == nil {
		return nil
	}
	out := new(StringDataFrom)
	in.DeepCopyInto(out)
	return out
}
