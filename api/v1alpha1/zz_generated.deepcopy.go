//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDeletion) DeepCopyInto(out *MachineDeletion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDeletion.
func (in *MachineDeletion) DeepCopy() *MachineDeletion {
	if in == nil {
		return nil
	}
	out := new(MachineDeletion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineDeletion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDeletionList) DeepCopyInto(out *MachineDeletionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineDeletion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDeletionList.
func (in *MachineDeletionList) DeepCopy() *MachineDeletionList {
	if in == nil {
		return nil
	}
	out := new(MachineDeletionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineDeletionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDeletionSpec) DeepCopyInto(out *MachineDeletionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDeletionSpec.
func (in *MachineDeletionSpec) DeepCopy() *MachineDeletionSpec {
	if in == nil {
		return nil
	}
	out := new(MachineDeletionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDeletionStatus) DeepCopyInto(out *MachineDeletionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDeletionStatus.
func (in *MachineDeletionStatus) DeepCopy() *MachineDeletionStatus {
	if in == nil {
		return nil
	}
	out := new(MachineDeletionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDeletionTemplate) DeepCopyInto(out *MachineDeletionTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDeletionTemplate.
func (in *MachineDeletionTemplate) DeepCopy() *MachineDeletionTemplate {
	if in == nil {
		return nil
	}
	out := new(MachineDeletionTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineDeletionTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDeletionTemplateList) DeepCopyInto(out *MachineDeletionTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineDeletionTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDeletionTemplateList.
func (in *MachineDeletionTemplateList) DeepCopy() *MachineDeletionTemplateList {
	if in == nil {
		return nil
	}
	out := new(MachineDeletionTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineDeletionTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDeletionTemplateResource) DeepCopyInto(out *MachineDeletionTemplateResource) {
	*out = *in
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDeletionTemplateResource.
func (in *MachineDeletionTemplateResource) DeepCopy() *MachineDeletionTemplateResource {
	if in == nil {
		return nil
	}
	out := new(MachineDeletionTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDeletionTemplateSpec) DeepCopyInto(out *MachineDeletionTemplateSpec) {
	*out = *in
	out.Template = in.Template
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDeletionTemplateSpec.
func (in *MachineDeletionTemplateSpec) DeepCopy() *MachineDeletionTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(MachineDeletionTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDeletionTemplateStatus) DeepCopyInto(out *MachineDeletionTemplateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDeletionTemplateStatus.
func (in *MachineDeletionTemplateStatus) DeepCopy() *MachineDeletionTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(MachineDeletionTemplateStatus)
	in.DeepCopyInto(out)
	return out
}
