package eventhubapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/eventhub/mgmt/2018-01-01-preview/eventhub"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result eventhub.OperationListResultPage, err error)
}

var _ OperationsClientAPI = (*eventhub.OperationsClient)(nil)

// ClustersClientAPI contains the set of methods on the ClustersClient type.
type ClustersClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, clusterName string) (result eventhub.ClustersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string) (result eventhub.Cluster, err error)
	ListAvailableClusters(ctx context.Context) (result eventhub.SetInt32, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result eventhub.ClusterListResultPage, err error)
	ListNamespaces(ctx context.Context, resourceGroupName string, clusterName string) (result eventhub.EHNamespaceIDListResult, err error)
	Patch(ctx context.Context, resourceGroupName string, clusterName string, parameters eventhub.Cluster) (result eventhub.ClustersPatchFuture, err error)
	Put(ctx context.Context, resourceGroupName string, clusterName string) (result eventhub.ClustersPutFuture, err error)
}

var _ ClustersClientAPI = (*eventhub.ClustersClient)(nil)

// ConfigurationClientAPI contains the set of methods on the ConfigurationClient type.
type ConfigurationClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, clusterName string) (result eventhub.ClusterQuotaConfigurationProperties, err error)
	Patch(ctx context.Context, resourceGroupName string, clusterName string, parameters eventhub.ClusterQuotaConfigurationProperties) (result eventhub.ClusterQuotaConfigurationProperties, err error)
}

var _ ConfigurationClientAPI = (*eventhub.ConfigurationClient)(nil)

// NamespacesClientAPI contains the set of methods on the NamespacesClient type.
type NamespacesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, parameters eventhub.EHNamespace) (result eventhub.NamespacesCreateOrUpdateFuture, err error)
	CreateOrUpdateIPFilterRule(ctx context.Context, resourceGroupName string, namespaceName string, IPFilterRuleName string, parameters eventhub.IPFilterRule) (result eventhub.IPFilterRule, err error)
	CreateOrUpdateNetworkRuleSet(ctx context.Context, resourceGroupName string, namespaceName string, parameters eventhub.NetworkRuleSet) (result eventhub.NetworkRuleSet, err error)
	CreateOrUpdateVirtualNetworkRule(ctx context.Context, resourceGroupName string, namespaceName string, virtualNetworkRuleName string, parameters eventhub.VirtualNetworkRule) (result eventhub.VirtualNetworkRule, err error)
	Delete(ctx context.Context, resourceGroupName string, namespaceName string) (result eventhub.NamespacesDeleteFuture, err error)
	DeleteIPFilterRule(ctx context.Context, resourceGroupName string, namespaceName string, IPFilterRuleName string) (result autorest.Response, err error)
	DeleteVirtualNetworkRule(ctx context.Context, resourceGroupName string, namespaceName string, virtualNetworkRuleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, namespaceName string) (result eventhub.EHNamespace, err error)
	GetIPFilterRule(ctx context.Context, resourceGroupName string, namespaceName string, IPFilterRuleName string) (result eventhub.IPFilterRule, err error)
	GetNetworkRuleSet(ctx context.Context, resourceGroupName string, namespaceName string) (result eventhub.NetworkRuleSet, err error)
	GetVirtualNetworkRule(ctx context.Context, resourceGroupName string, namespaceName string, virtualNetworkRuleName string) (result eventhub.VirtualNetworkRule, err error)
	List(ctx context.Context) (result eventhub.EHNamespaceListResultPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result eventhub.EHNamespaceListResultPage, err error)
	ListIPFilterRules(ctx context.Context, resourceGroupName string, namespaceName string) (result eventhub.IPFilterRuleListResultPage, err error)
	ListVirtualNetworkRules(ctx context.Context, resourceGroupName string, namespaceName string) (result eventhub.VirtualNetworkRuleListResultPage, err error)
	Update(ctx context.Context, resourceGroupName string, namespaceName string, parameters eventhub.EHNamespace) (result eventhub.EHNamespace, err error)
}

var _ NamespacesClientAPI = (*eventhub.NamespacesClient)(nil)
