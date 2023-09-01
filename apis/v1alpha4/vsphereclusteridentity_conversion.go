/*
Copyright 2022 The Kubernetes Authors.

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

package v1alpha4

import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	infrav1 "sigs.k8s.io/cluster-api-provider-vsphere/apis/v1beta1"
)

// ConvertTo converts this VSphereClusterIdentity to the Hub version (v1beta1).
func (src *VSphereClusterIdentity) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1.VSphereClusterIdentity)
	return Convert_v1alpha4_VSphereClusterIdentity_To_v1beta1_VSphereClusterIdentity(src, dst, nil)
}

// ConvertFrom converts from the Hub version (v1beta1) to this VSphereClusterIdentity.
func (dst *VSphereClusterIdentity) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1.VSphereClusterIdentity)
	return Convert_v1beta1_VSphereClusterIdentity_To_v1alpha4_VSphereClusterIdentity(src, dst, nil)
}

// ConvertTo converts this VSphereClusterIdentityList to the Hub version (v1beta1).
func (src *VSphereClusterIdentityList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1.VSphereClusterIdentityList)
	return Convert_v1alpha4_VSphereClusterIdentityList_To_v1beta1_VSphereClusterIdentityList(src, dst, nil)
}

// ConvertFrom converts this VSphereClusterIdentityList to the Hub version (v1beta1).
func (dst *VSphereClusterIdentityList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1.VSphereClusterIdentityList)
	return Convert_v1beta1_VSphereClusterIdentityList_To_v1alpha4_VSphereClusterIdentityList(src, dst, nil)
}
