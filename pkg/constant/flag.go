/*
Copyright (C) 2022-2024 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package constant

const (
	EnableRBACManager = "EnableRBACManager"

	ManagedNamespacesFlag = "managed-namespaces"
)

const (
	// ShardSvcAnnotationKey defines the feature gate of creating service for each shard.
	// Sharding name defined in the annotation value, a set of Service defined in Cluster.Spec.Services with the ShardingSelector will be automatically generated for each shard when Cluster.Spec.ShardingSpecs[x].shards is not nil.
	// Multiple sharding names are separated by ','. for example: "kubeblocks.io/enabled-shard-svc: proxy-shard,db-shard"
	ShardSvcAnnotationKey = "kubeblocks.io/enabled-shard-svc"

	// HostNetworkAnnotationKey defines the feature gate to enable the host-network for specified components or shardings.
	HostNetworkAnnotationKey = "kubeblocks.io/host-network"

	// FeatureReconciliationInCompactModeAnnotationKey indicates that the controller should run in compact mode,
	// means to try the best to cutoff useless objects.
	FeatureReconciliationInCompactModeAnnotationKey = "kubeblocks.io/compact-mode"

	// FeatureGateInPlacePodVerticalScaling specifies to enable in-place pod vertical scaling
	// NOTE: This feature depends on the InPlacePodVerticalScaling feature of the K8s cluster in which the KubeBlocks runs.
	FeatureGateInPlacePodVerticalScaling = "IN_PLACE_POD_VERTICAL_SCALING"

	// FeatureGateNoRSMEnv specifies not to create the '$(instanceset.name)-rsm-env' ConfigMap object.
	FeatureGateNoRSMEnv = "NO_RSM_ENV"
)
