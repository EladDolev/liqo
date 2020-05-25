// +build !ignore_autogenerated

/*

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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Advertisement) DeepCopyInto(out *Advertisement) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Advertisement.
func (in *Advertisement) DeepCopy() *Advertisement {
	if in == nil {
		return nil
	}
	out := new(Advertisement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Advertisement) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvertisementList) DeepCopyInto(out *AdvertisementList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Advertisement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvertisementList.
func (in *AdvertisementList) DeepCopy() *AdvertisementList {
	if in == nil {
		return nil
	}
	out := new(AdvertisementList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdvertisementList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvertisementSpec) DeepCopyInto(out *AdvertisementSpec) {
	*out = *in
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]corev1.ContainerImage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.LimitRange.DeepCopyInto(&out.LimitRange)
	in.ResourceQuota.DeepCopyInto(&out.ResourceQuota)
	if in.Neighbors != nil {
		in, out := &in.Neighbors, &out.Neighbors
		*out = make(map[corev1.ResourceName]corev1.ResourceList, len(*in))
		for key, val := range *in {
			var outVal map[corev1.ResourceName]resource.Quantity
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(corev1.ResourceList, len(*in))
				for key, val := range *in {
					(*out)[key] = val.DeepCopy()
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[corev1.ResourceName]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Prices != nil {
		in, out := &in.Prices, &out.Prices
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	in.Network.DeepCopyInto(&out.Network)
	in.Timestamp.DeepCopyInto(&out.Timestamp)
	in.TimeToLive.DeepCopyInto(&out.TimeToLive)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvertisementSpec.
func (in *AdvertisementSpec) DeepCopy() *AdvertisementSpec {
	if in == nil {
		return nil
	}
	out := new(AdvertisementSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvertisementStatus) DeepCopyInto(out *AdvertisementStatus) {
	*out = *in
	in.ForeignNetwork.DeepCopyInto(&out.ForeignNetwork)
	out.TunnelEndpointKey = in.TunnelEndpointKey
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvertisementStatus.
func (in *AdvertisementStatus) DeepCopy() *AdvertisementStatus {
	if in == nil {
		return nil
	}
	out := new(AdvertisementStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedName) DeepCopyInto(out *NamespacedName) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedName.
func (in *NamespacedName) DeepCopy() *NamespacedName {
	if in == nil {
		return nil
	}
	out := new(NamespacedName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInfo) DeepCopyInto(out *NetworkInfo) {
	*out = *in
	if in.SupportedProtocols != nil {
		in, out := &in.SupportedProtocols, &out.SupportedProtocols
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInfo.
func (in *NetworkInfo) DeepCopy() *NetworkInfo {
	if in == nil {
		return nil
	}
	out := new(NetworkInfo)
	in.DeepCopyInto(out)
	return out
}
