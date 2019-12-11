// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MustGather) DeepCopyInto(out *MustGather) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MustGather.
func (in *MustGather) DeepCopy() *MustGather {
	if in == nil {
		return nil
	}
	out := new(MustGather)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MustGather) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MustGatherList) DeepCopyInto(out *MustGatherList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MustGather, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MustGatherList.
func (in *MustGatherList) DeepCopy() *MustGatherList {
	if in == nil {
		return nil
	}
	out := new(MustGatherList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MustGatherList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MustGatherSpec) DeepCopyInto(out *MustGatherSpec) {
	*out = *in
	out.CaseManagementAccountSecretRef = in.CaseManagementAccountSecretRef
	out.ServiceAccountRef = in.ServiceAccountRef
	if in.MustGatherImages != nil {
		in, out := &in.MustGatherImages, &out.MustGatherImages
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.ProxyConfig = in.ProxyConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MustGatherSpec.
func (in *MustGatherSpec) DeepCopy() *MustGatherSpec {
	if in == nil {
		return nil
	}
	out := new(MustGatherSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MustGatherStatus) DeepCopyInto(out *MustGatherStatus) {
	*out = *in
	in.LastUpdate.DeepCopyInto(&out.LastUpdate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MustGatherStatus.
func (in *MustGatherStatus) DeepCopy() *MustGatherStatus {
	if in == nil {
		return nil
	}
	out := new(MustGatherStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxySpec) DeepCopyInto(out *ProxySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxySpec.
func (in *ProxySpec) DeepCopy() *ProxySpec {
	if in == nil {
		return nil
	}
	out := new(ProxySpec)
	in.DeepCopyInto(out)
	return out
}
