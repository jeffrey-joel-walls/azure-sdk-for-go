// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder
// commit ID: 2014fbbf031942474ad27a5a66dffaed5347f3fb

package advisor

import original "github.com/Azure/azure-sdk-for-go/services/advisor/mgmt/2017-04-19/advisor"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type ConfigurationsClient = original.ConfigurationsClient
type Category = original.Category

const (
	Cost             Category = original.Cost
	HighAvailability Category = original.HighAvailability
	Performance      Category = original.Performance
	Security         Category = original.Security
)

type Impact = original.Impact

const (
	High   Impact = original.High
	Low    Impact = original.Low
	Medium Impact = original.Medium
)

type Risk = original.Risk

const (
	Error   Risk = original.Error
	None    Risk = original.None
	Warning Risk = original.Warning
)

type ARMErrorResponseBody = original.ARMErrorResponseBody
type ConfigData = original.ConfigData
type ConfigDataProperties = original.ConfigDataProperties
type ConfigurationListResult = original.ConfigurationListResult
type OperationDisplayInfo = original.OperationDisplayInfo
type OperationEntity = original.OperationEntity
type OperationEntityListResult = original.OperationEntityListResult
type OperationEntityListResultIterator = original.OperationEntityListResultIterator
type OperationEntityListResultPage = original.OperationEntityListResultPage
type RecommendationProperties = original.RecommendationProperties
type RecommendationsGetGenerateStatusFuture = original.RecommendationsGetGenerateStatusFuture
type Resource = original.Resource
type ResourceRecommendationBase = original.ResourceRecommendationBase
type ResourceRecommendationBaseListResult = original.ResourceRecommendationBaseListResult
type ResourceRecommendationBaseListResultIterator = original.ResourceRecommendationBaseListResultIterator
type ResourceRecommendationBaseListResultPage = original.ResourceRecommendationBaseListResultPage
type ShortDescription = original.ShortDescription
type SuppressionContract = original.SuppressionContract
type SuppressionContractListResult = original.SuppressionContractListResult
type SuppressionContractListResultIterator = original.SuppressionContractListResultIterator
type SuppressionContractListResultPage = original.SuppressionContractListResultPage
type SuppressionProperties = original.SuppressionProperties
type OperationsClient = original.OperationsClient
type RecommendationsClient = original.RecommendationsClient
type SuppressionsClient = original.SuppressionsClient

func NewRecommendationsClient(subscriptionID string) RecommendationsClient {
	return original.NewRecommendationsClient(subscriptionID)
}
func NewRecommendationsClientWithBaseURI(baseURI string, subscriptionID string) RecommendationsClient {
	return original.NewRecommendationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSuppressionsClient(subscriptionID string) SuppressionsClient {
	return original.NewSuppressionsClient(subscriptionID)
}
func NewSuppressionsClientWithBaseURI(baseURI string, subscriptionID string) SuppressionsClient {
	return original.NewSuppressionsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewConfigurationsClient(subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClient(subscriptionID)
}
func NewConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
