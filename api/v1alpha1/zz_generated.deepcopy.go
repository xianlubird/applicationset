// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSet) DeepCopyInto(out *ApplicationSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSet.
func (in *ApplicationSet) DeepCopy() *ApplicationSet {
	if in == nil {
		return nil
	}
	out := new(ApplicationSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSetGenerator) DeepCopyInto(out *ApplicationSetGenerator) {
	*out = *in
	if in.List != nil {
		in, out := &in.List, &out.List
		*out = new(ListGenerator)
		(*in).DeepCopyInto(*out)
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = new(ClusterGenerator)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSetGenerator.
func (in *ApplicationSetGenerator) DeepCopy() *ApplicationSetGenerator {
	if in == nil {
		return nil
	}
	out := new(ApplicationSetGenerator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSetList) DeepCopyInto(out *ApplicationSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApplicationSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSetList.
func (in *ApplicationSetList) DeepCopy() *ApplicationSetList {
	if in == nil {
		return nil
	}
	out := new(ApplicationSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSetSpec) DeepCopyInto(out *ApplicationSetSpec) {
	*out = *in
	if in.Generators != nil {
		in, out := &in.Generators, &out.Generators
		*out = make([]ApplicationSetGenerator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Template.DeepCopyInto(&out.Template)
	if in.SyncPolicy != nil {
		in, out := &in.SyncPolicy, &out.SyncPolicy
		*out = new(ApplicationSetSyncPolicy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSetSpec.
func (in *ApplicationSetSpec) DeepCopy() *ApplicationSetSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSetSyncPolicy) DeepCopyInto(out *ApplicationSetSyncPolicy) {
	*out = *in
	if in.Automated != nil {
		in, out := &in.Automated, &out.Automated
		*out = new(SyncPolicyAutomated)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSetSyncPolicy.
func (in *ApplicationSetSyncPolicy) DeepCopy() *ApplicationSetSyncPolicy {
	if in == nil {
		return nil
	}
	out := new(ApplicationSetSyncPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSetTemplate) DeepCopyInto(out *ApplicationSetTemplate) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSetTemplate.
func (in *ApplicationSetTemplate) DeepCopy() *ApplicationSetTemplate {
	if in == nil {
		return nil
	}
	out := new(ApplicationSetTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterGenerator) DeepCopyInto(out *ClusterGenerator) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterGenerator.
func (in *ClusterGenerator) DeepCopy() *ClusterGenerator {
	if in == nil {
		return nil
	}
	out := new(ClusterGenerator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListGenerator) DeepCopyInto(out *ListGenerator) {
	*out = *in
	if in.Elements != nil {
		in, out := &in.Elements, &out.Elements
		*out = make([]ListGeneratorElement, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListGenerator.
func (in *ListGenerator) DeepCopy() *ListGenerator {
	if in == nil {
		return nil
	}
	out := new(ListGenerator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListGeneratorElement) DeepCopyInto(out *ListGeneratorElement) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListGeneratorElement.
func (in *ListGeneratorElement) DeepCopy() *ListGeneratorElement {
	if in == nil {
		return nil
	}
	out := new(ListGeneratorElement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncPolicyAutomated) DeepCopyInto(out *SyncPolicyAutomated) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncPolicyAutomated.
func (in *SyncPolicyAutomated) DeepCopy() *SyncPolicyAutomated {
	if in == nil {
		return nil
	}
	out := new(SyncPolicyAutomated)
	in.DeepCopyInto(out)
	return out
}
