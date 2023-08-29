/*
Copyright 2021 The Kubernetes Authors.

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

package v1alpha3

import (
	"k8s.io/apimachinery/pkg/conversion"

	"sigs.k8s.io/cluster-api-provider-vsphere/apis/v1beta1"
)

// Convert_v1beta1_VirtualMachineCloneSpec_To_v1alpha3_VirtualMachineCloneSpec is an autogenerated conversion function.
func Convert_v1beta1_VirtualMachineCloneSpec_To_v1alpha3_VirtualMachineCloneSpec(in *v1beta1.VirtualMachineCloneSpec, out *VirtualMachineCloneSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_VirtualMachineCloneSpec_To_v1alpha3_VirtualMachineCloneSpec(in, out, s)
}

// Convert_v1beta1_VSphereVMStatus_To_v1alpha3_VSphereVMStatus is an autogenerated conversion function.
func Convert_v1beta1_VSphereVMStatus_To_v1alpha3_VSphereVMStatus(in *v1beta1.VSphereVMStatus, out *VSphereVMStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_VSphereVMStatus_To_v1alpha3_VSphereVMStatus(in, out, s)
}

func Convert_v1beta1_VSphereClusterStatus_To_v1alpha3_VSphereClusterStatus(in *v1beta1.VSphereClusterStatus, out *VSphereClusterStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_VSphereClusterStatus_To_v1alpha3_VSphereClusterStatus(in, out, s)
}

func Convert_v1beta1_VSphereClusterSpec_To_v1alpha3_VSphereClusterSpec(in *v1beta1.VSphereClusterSpec, out *VSphereClusterSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_VSphereClusterSpec_To_v1alpha3_VSphereClusterSpec(in, out, s)
}

func Convert_v1beta1_VSphereMachineSpec_To_v1alpha3_VSphereMachineSpec(in *v1beta1.VSphereMachineSpec, out *VSphereMachineSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_VSphereMachineSpec_To_v1alpha3_VSphereMachineSpec(in, out, s)
}

func Convert_v1beta1_VSphereVMSpec_To_v1alpha3_VSphereVMSpec(in *v1beta1.VSphereVMSpec, out *VSphereVMSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_VSphereVMSpec_To_v1alpha3_VSphereVMSpec(in, out, s)
}
