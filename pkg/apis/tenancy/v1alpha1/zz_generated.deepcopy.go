//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"

	conditionsv1alpha1 "github.com/kcp-dev/kcp/third_party/conditions/apis/conditions/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkspace) DeepCopyInto(out *ClusterWorkspace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkspace.
func (in *ClusterWorkspace) DeepCopy() *ClusterWorkspace {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterWorkspace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkspaceList) DeepCopyInto(out *ClusterWorkspaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterWorkspace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkspaceList.
func (in *ClusterWorkspaceList) DeepCopy() *ClusterWorkspaceList {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkspaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterWorkspaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkspaceLocation) DeepCopyInto(out *ClusterWorkspaceLocation) {
	*out = *in
	if in.History != nil {
		in, out := &in.History, &out.History
		*out = make([]ShardStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkspaceLocation.
func (in *ClusterWorkspaceLocation) DeepCopy() *ClusterWorkspaceLocation {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkspaceLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkspaceSpec) DeepCopyInto(out *ClusterWorkspaceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkspaceSpec.
func (in *ClusterWorkspaceSpec) DeepCopy() *ClusterWorkspaceSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkspaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkspaceStatus) DeepCopyInto(out *ClusterWorkspaceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(conditionsv1alpha1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Location.DeepCopyInto(&out.Location)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkspaceStatus.
func (in *ClusterWorkspaceStatus) DeepCopy() *ClusterWorkspaceStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkspaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionInfo) DeepCopyInto(out *ConnectionInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionInfo.
func (in *ConnectionInfo) DeepCopy() *ConnectionInfo {
	if in == nil {
		return nil
	}
	out := new(ConnectionInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShardStatus) DeepCopyInto(out *ShardStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShardStatus.
func (in *ShardStatus) DeepCopy() *ShardStatus {
	if in == nil {
		return nil
	}
	out := new(ShardStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceShard) DeepCopyInto(out *WorkspaceShard) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceShard.
func (in *WorkspaceShard) DeepCopy() *WorkspaceShard {
	if in == nil {
		return nil
	}
	out := new(WorkspaceShard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceShard) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceShardList) DeepCopyInto(out *WorkspaceShardList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkspaceShard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceShardList.
func (in *WorkspaceShardList) DeepCopy() *WorkspaceShardList {
	if in == nil {
		return nil
	}
	out := new(WorkspaceShardList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceShardList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceShardSpec) DeepCopyInto(out *WorkspaceShardSpec) {
	*out = *in
	out.Credentials = in.Credentials
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceShardSpec.
func (in *WorkspaceShardSpec) DeepCopy() *WorkspaceShardSpec {
	if in == nil {
		return nil
	}
	out := new(WorkspaceShardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceShardStatus) DeepCopyInto(out *WorkspaceShardStatus) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(conditionsv1alpha1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ConnectionInfo != nil {
		in, out := &in.ConnectionInfo, &out.ConnectionInfo
		*out = new(ConnectionInfo)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceShardStatus.
func (in *WorkspaceShardStatus) DeepCopy() *WorkspaceShardStatus {
	if in == nil {
		return nil
	}
	out := new(WorkspaceShardStatus)
	in.DeepCopyInto(out)
	return out
}
