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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1alpha1 "github.com/kcp-dev/kcp/sdk/apis/apis/v1alpha1"
	corev1alpha1 "github.com/kcp-dev/kcp/sdk/apis/core/v1alpha1"
	tenancyv1alpha1 "github.com/kcp-dev/kcp/sdk/apis/tenancy/v1alpha1"
	conditionsv1alpha1 "github.com/kcp-dev/kcp/sdk/apis/third_party/conditions/apis/conditions/v1alpha1"
	topologyv1alpha1 "github.com/kcp-dev/kcp/sdk/apis/topology/v1alpha1"
	apisv1alpha1 "github.com/kcp-dev/kcp/sdk/client/applyconfiguration/apis/v1alpha1"
	applyconfigurationconditionsv1alpha1 "github.com/kcp-dev/kcp/sdk/client/applyconfiguration/conditions/v1alpha1"
	applyconfigurationcorev1alpha1 "github.com/kcp-dev/kcp/sdk/client/applyconfiguration/core/v1alpha1"
	internal "github.com/kcp-dev/kcp/sdk/client/applyconfiguration/internal"
	applyconfigurationmetav1 "github.com/kcp-dev/kcp/sdk/client/applyconfiguration/meta/v1"
	applyconfigurationtenancyv1alpha1 "github.com/kcp-dev/kcp/sdk/client/applyconfiguration/tenancy/v1alpha1"
	applyconfigurationtopologyv1alpha1 "github.com/kcp-dev/kcp/sdk/client/applyconfiguration/topology/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	metav1 "k8s.io/client-go/applyconfigurations/meta/v1"
	testing "k8s.io/client-go/testing"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=apis.kcp.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("AcceptablePermissionClaim"):
		return &apisv1alpha1.AcceptablePermissionClaimApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("APIBinding"):
		return &apisv1alpha1.APIBindingApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("APIBindingSpec"):
		return &apisv1alpha1.APIBindingSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("APIBindingStatus"):
		return &apisv1alpha1.APIBindingStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("APIConversion"):
		return &apisv1alpha1.APIConversionApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("APIConversionRule"):
		return &apisv1alpha1.APIConversionRuleApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("APIConversionSpec"):
		return &apisv1alpha1.APIConversionSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("APIExport"):
		return &apisv1alpha1.APIExportApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("APIExportEndpoint"):
		return &apisv1alpha1.APIExportEndpointApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("APIExportEndpointSlice"):
		return &apisv1alpha1.APIExportEndpointSliceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("APIExportEndpointSliceSpec"):
		return &apisv1alpha1.APIExportEndpointSliceSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("APIExportEndpointSliceStatus"):
		return &apisv1alpha1.APIExportEndpointSliceStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("APIExportSpec"):
		return &apisv1alpha1.APIExportSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("APIExportStatus"):
		return &apisv1alpha1.APIExportStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("APIResourceSchema"):
		return &apisv1alpha1.APIResourceSchemaApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("APIResourceSchemaSpec"):
		return &apisv1alpha1.APIResourceSchemaSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("APIResourceVersion"):
		return &apisv1alpha1.APIResourceVersionApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("APIVersionConversion"):
		return &apisv1alpha1.APIVersionConversionApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("BindingReference"):
		return &apisv1alpha1.BindingReferenceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("BoundAPIResource"):
		return &apisv1alpha1.BoundAPIResourceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("BoundAPIResourceSchema"):
		return &apisv1alpha1.BoundAPIResourceSchemaApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("CustomResourceConversion"):
		return &apisv1alpha1.CustomResourceConversionApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ExportBindingReference"):
		return &apisv1alpha1.ExportBindingReferenceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("GroupResource"):
		return &apisv1alpha1.GroupResourceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Identity"):
		return &apisv1alpha1.IdentityApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("MaximalPermissionPolicy"):
		return &apisv1alpha1.MaximalPermissionPolicyApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("PermissionClaim"):
		return &apisv1alpha1.PermissionClaimApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ResourceSelector"):
		return &apisv1alpha1.ResourceSelectorApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("VirtualWorkspace"):
		return &apisv1alpha1.VirtualWorkspaceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("WebhookClientConfig"):
		return &apisv1alpha1.WebhookClientConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("WebhookConversion"):
		return &apisv1alpha1.WebhookConversionApplyConfiguration{}

		// Group=conditions, Version=v1alpha1
	case conditionsv1alpha1.SchemeGroupVersion.WithKind("Condition"):
		return &applyconfigurationconditionsv1alpha1.ConditionApplyConfiguration{}

		// Group=core.kcp.io, Version=v1alpha1
	case corev1alpha1.SchemeGroupVersion.WithKind("LogicalCluster"):
		return &applyconfigurationcorev1alpha1.LogicalClusterApplyConfiguration{}
	case corev1alpha1.SchemeGroupVersion.WithKind("LogicalClusterOwner"):
		return &applyconfigurationcorev1alpha1.LogicalClusterOwnerApplyConfiguration{}
	case corev1alpha1.SchemeGroupVersion.WithKind("LogicalClusterSpec"):
		return &applyconfigurationcorev1alpha1.LogicalClusterSpecApplyConfiguration{}
	case corev1alpha1.SchemeGroupVersion.WithKind("LogicalClusterStatus"):
		return &applyconfigurationcorev1alpha1.LogicalClusterStatusApplyConfiguration{}
	case corev1alpha1.SchemeGroupVersion.WithKind("Shard"):
		return &applyconfigurationcorev1alpha1.ShardApplyConfiguration{}
	case corev1alpha1.SchemeGroupVersion.WithKind("ShardSpec"):
		return &applyconfigurationcorev1alpha1.ShardSpecApplyConfiguration{}
	case corev1alpha1.SchemeGroupVersion.WithKind("ShardStatus"):
		return &applyconfigurationcorev1alpha1.ShardStatusApplyConfiguration{}

		// Group=meta.k8s.io, Version=v1
	case v1.SchemeGroupVersion.WithKind("Condition"):
		return &metav1.ConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DeleteOptions"):
		return &metav1.DeleteOptionsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("LabelSelector"):
		return &applyconfigurationmetav1.LabelSelectorApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("LabelSelectorRequirement"):
		return &applyconfigurationmetav1.LabelSelectorRequirementApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ManagedFieldsEntry"):
		return &applyconfigurationmetav1.ManagedFieldsEntryApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ObjectMeta"):
		return &applyconfigurationmetav1.ObjectMetaApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OwnerReference"):
		return &applyconfigurationmetav1.OwnerReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("TypeMeta"):
		return &applyconfigurationmetav1.TypeMetaApplyConfiguration{}

		// Group=tenancy.kcp.io, Version=v1alpha1
	case tenancyv1alpha1.SchemeGroupVersion.WithKind("APIExportReference"):
		return &applyconfigurationtenancyv1alpha1.APIExportReferenceApplyConfiguration{}
	case tenancyv1alpha1.SchemeGroupVersion.WithKind("VirtualWorkspace"):
		return &applyconfigurationtenancyv1alpha1.VirtualWorkspaceApplyConfiguration{}
	case tenancyv1alpha1.SchemeGroupVersion.WithKind("Workspace"):
		return &applyconfigurationtenancyv1alpha1.WorkspaceApplyConfiguration{}
	case tenancyv1alpha1.SchemeGroupVersion.WithKind("WorkspaceLocation"):
		return &applyconfigurationtenancyv1alpha1.WorkspaceLocationApplyConfiguration{}
	case tenancyv1alpha1.SchemeGroupVersion.WithKind("WorkspaceSpec"):
		return &applyconfigurationtenancyv1alpha1.WorkspaceSpecApplyConfiguration{}
	case tenancyv1alpha1.SchemeGroupVersion.WithKind("WorkspaceStatus"):
		return &applyconfigurationtenancyv1alpha1.WorkspaceStatusApplyConfiguration{}
	case tenancyv1alpha1.SchemeGroupVersion.WithKind("WorkspaceType"):
		return &applyconfigurationtenancyv1alpha1.WorkspaceTypeApplyConfiguration{}
	case tenancyv1alpha1.SchemeGroupVersion.WithKind("WorkspaceTypeExtension"):
		return &applyconfigurationtenancyv1alpha1.WorkspaceTypeExtensionApplyConfiguration{}
	case tenancyv1alpha1.SchemeGroupVersion.WithKind("WorkspaceTypeReference"):
		return &applyconfigurationtenancyv1alpha1.WorkspaceTypeReferenceApplyConfiguration{}
	case tenancyv1alpha1.SchemeGroupVersion.WithKind("WorkspaceTypeSelector"):
		return &applyconfigurationtenancyv1alpha1.WorkspaceTypeSelectorApplyConfiguration{}
	case tenancyv1alpha1.SchemeGroupVersion.WithKind("WorkspaceTypeSpec"):
		return &applyconfigurationtenancyv1alpha1.WorkspaceTypeSpecApplyConfiguration{}
	case tenancyv1alpha1.SchemeGroupVersion.WithKind("WorkspaceTypeStatus"):
		return &applyconfigurationtenancyv1alpha1.WorkspaceTypeStatusApplyConfiguration{}

		// Group=topology.kcp.io, Version=v1alpha1
	case topologyv1alpha1.SchemeGroupVersion.WithKind("Partition"):
		return &applyconfigurationtopologyv1alpha1.PartitionApplyConfiguration{}
	case topologyv1alpha1.SchemeGroupVersion.WithKind("PartitionSet"):
		return &applyconfigurationtopologyv1alpha1.PartitionSetApplyConfiguration{}
	case topologyv1alpha1.SchemeGroupVersion.WithKind("PartitionSetSpec"):
		return &applyconfigurationtopologyv1alpha1.PartitionSetSpecApplyConfiguration{}
	case topologyv1alpha1.SchemeGroupVersion.WithKind("PartitionSetStatus"):
		return &applyconfigurationtopologyv1alpha1.PartitionSetStatusApplyConfiguration{}
	case topologyv1alpha1.SchemeGroupVersion.WithKind("PartitionSpec"):
		return &applyconfigurationtopologyv1alpha1.PartitionSpecApplyConfiguration{}

	}
	return nil
}

func NewTypeConverter(scheme *runtime.Scheme) *testing.TypeConverter {
	return &testing.TypeConverter{Scheme: scheme, TypeResolver: internal.Parser()}
}
