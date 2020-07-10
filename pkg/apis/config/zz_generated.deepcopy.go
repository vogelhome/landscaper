// +build !ignore_autogenerated

/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

package config

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Component) DeepCopyInto(out *Component) {
	*out = *in
	out.DependencyMeta = in.DependencyMeta
	in.Dependencies.DeepCopyInto(&out.Dependencies)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Component.
func (in *Component) DeepCopy() *Component {
	if in == nil {
		return nil
	}
	out := new(Component)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentDependency) DeepCopyInto(out *ComponentDependency) {
	*out = *in
	out.DependencyMeta = in.DependencyMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDependency.
func (in *ComponentDependency) DeepCopy() *ComponentDependency {
	if in == nil {
		return nil
	}
	out := new(ComponentDependency)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentList) DeepCopyInto(out *ComponentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]*Component, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Component)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentList.
func (in *ComponentList) DeepCopy() *ComponentList {
	if in == nil {
		return nil
	}
	out := new(ComponentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComponentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentOverwrite) DeepCopyInto(out *ComponentOverwrite) {
	*out = *in
	out.DeclaringComponent = in.DeclaringComponent
	in.DependencyOverwrites.DeepCopyInto(&out.DependencyOverwrites)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentOverwrite.
func (in *ComponentOverwrite) DeepCopy() *ComponentOverwrite {
	if in == nil {
		return nil
	}
	out := new(ComponentOverwrite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dependencies) DeepCopyInto(out *Dependencies) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]ComponentDependency, len(*in))
		copy(*out, *in)
	}
	if in.ContainerImages != nil {
		in, out := &in.ContainerImages, &out.ContainerImages
		*out = make([]ImageDependency, len(*in))
		copy(*out, *in)
	}
	if in.HelmCharts != nil {
		in, out := &in.HelmCharts, &out.HelmCharts
		*out = make([]HelmChartDependency, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dependencies.
func (in *Dependencies) DeepCopy() *Dependencies {
	if in == nil {
		return nil
	}
	out := new(Dependencies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DependencyMeta) DeepCopyInto(out *DependencyMeta) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DependencyMeta.
func (in *DependencyMeta) DeepCopy() *DependencyMeta {
	if in == nil {
		return nil
	}
	out := new(DependencyMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmChartDependency) DeepCopyInto(out *HelmChartDependency) {
	*out = *in
	out.DependencyMeta = in.DependencyMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmChartDependency.
func (in *HelmChartDependency) DeepCopy() *HelmChartDependency {
	if in == nil {
		return nil
	}
	out := new(HelmChartDependency)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageDependency) DeepCopyInto(out *ImageDependency) {
	*out = *in
	out.DependencyMeta = in.DependencyMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageDependency.
func (in *ImageDependency) DeepCopy() *ImageDependency {
	if in == nil {
		return nil
	}
	out := new(ImageDependency)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LandscaperConfiguration) DeepCopyInto(out *LandscaperConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Registry.DeepCopyInto(&out.Registry)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LandscaperConfiguration.
func (in *LandscaperConfiguration) DeepCopy() *LandscaperConfiguration {
	if in == nil {
		return nil
	}
	out := new(LandscaperConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LandscaperConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalRegistryConfiguration) DeepCopyInto(out *LocalRegistryConfiguration) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalRegistryConfiguration.
func (in *LocalRegistryConfiguration) DeepCopy() *LocalRegistryConfiguration {
	if in == nil {
		return nil
	}
	out := new(LocalRegistryConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryConfiguration) DeepCopyInto(out *RegistryConfiguration) {
	*out = *in
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(LocalRegistryConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryConfiguration.
func (in *RegistryConfiguration) DeepCopy() *RegistryConfiguration {
	if in == nil {
		return nil
	}
	out := new(RegistryConfiguration)
	in.DeepCopyInto(out)
	return out
}
