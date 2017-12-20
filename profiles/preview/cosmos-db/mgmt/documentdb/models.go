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

package documentdb

import original "github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2015-04-08/documentdb"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type CollectionRegionClient = original.CollectionRegionClient
type DatabaseAccountKind = original.DatabaseAccountKind

const (
	GlobalDocumentDB DatabaseAccountKind = original.GlobalDocumentDB
	MongoDB          DatabaseAccountKind = original.MongoDB
	Parse            DatabaseAccountKind = original.Parse
)

type DatabaseAccountOfferType = original.DatabaseAccountOfferType

const (
	Standard DatabaseAccountOfferType = original.Standard
)

type DefaultConsistencyLevel = original.DefaultConsistencyLevel

const (
	BoundedStaleness DefaultConsistencyLevel = original.BoundedStaleness
	ConsistentPrefix DefaultConsistencyLevel = original.ConsistentPrefix
	Eventual         DefaultConsistencyLevel = original.Eventual
	Session          DefaultConsistencyLevel = original.Session
	Strong           DefaultConsistencyLevel = original.Strong
)

type KeyKind = original.KeyKind

const (
	Primary           KeyKind = original.Primary
	PrimaryReadonly   KeyKind = original.PrimaryReadonly
	Secondary         KeyKind = original.Secondary
	SecondaryReadonly KeyKind = original.SecondaryReadonly
)

type PrimaryAggregationType = original.PrimaryAggregationType

const (
	Average   PrimaryAggregationType = original.Average
	Last      PrimaryAggregationType = original.Last
	Maximum   PrimaryAggregationType = original.Maximum
	Minimimum PrimaryAggregationType = original.Minimimum
	None      PrimaryAggregationType = original.None
	Total     PrimaryAggregationType = original.Total
)

type UnitType = original.UnitType

const (
	Bytes          UnitType = original.Bytes
	BytesPerSecond UnitType = original.BytesPerSecond
	Count          UnitType = original.Count
	CountPerSecond UnitType = original.CountPerSecond
	Milliseconds   UnitType = original.Milliseconds
	Percent        UnitType = original.Percent
	Seconds        UnitType = original.Seconds
)

type ConsistencyPolicy = original.ConsistencyPolicy
type DatabaseAccount = original.DatabaseAccount
type DatabaseAccountConnectionString = original.DatabaseAccountConnectionString
type DatabaseAccountCreateUpdateParameters = original.DatabaseAccountCreateUpdateParameters
type DatabaseAccountCreateUpdateProperties = original.DatabaseAccountCreateUpdateProperties
type DatabaseAccountListConnectionStringsResult = original.DatabaseAccountListConnectionStringsResult
type DatabaseAccountListKeysResult = original.DatabaseAccountListKeysResult
type DatabaseAccountListReadOnlyKeysResult = original.DatabaseAccountListReadOnlyKeysResult
type DatabaseAccountPatchParameters = original.DatabaseAccountPatchParameters
type DatabaseAccountProperties = original.DatabaseAccountProperties
type DatabaseAccountRegenerateKeyParameters = original.DatabaseAccountRegenerateKeyParameters
type DatabaseAccountsCreateOrUpdateFuture = original.DatabaseAccountsCreateOrUpdateFuture
type DatabaseAccountsDeleteFuture = original.DatabaseAccountsDeleteFuture
type DatabaseAccountsFailoverPriorityChangeFuture = original.DatabaseAccountsFailoverPriorityChangeFuture
type DatabaseAccountsListResult = original.DatabaseAccountsListResult
type DatabaseAccountsPatchFuture = original.DatabaseAccountsPatchFuture
type DatabaseAccountsRegenerateKeyFuture = original.DatabaseAccountsRegenerateKeyFuture
type FailoverPolicies = original.FailoverPolicies
type FailoverPolicy = original.FailoverPolicy
type Location = original.Location
type Metric = original.Metric
type MetricAvailability = original.MetricAvailability
type MetricDefinition = original.MetricDefinition
type MetricDefinitionsListResult = original.MetricDefinitionsListResult
type MetricListResult = original.MetricListResult
type MetricName = original.MetricName
type MetricValue = original.MetricValue
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type Resource = original.Resource
type Usage = original.Usage
type UsagesResult = original.UsagesResult
type OperationsClient = original.OperationsClient
type CollectionClient = original.CollectionClient
type DatabaseClient = original.DatabaseClient
type DatabaseAccountRegionClient = original.DatabaseAccountRegionClient
type DatabaseAccountsClient = original.DatabaseAccountsClient

func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
func New(subscriptionID string, filter string, filter1 string, databaseRid string, collectionRid string, region string) BaseClient {
	return original.New(subscriptionID, filter, filter1, databaseRid, collectionRid, region)
}
func NewWithBaseURI(baseURI string, subscriptionID string, filter string, filter1 string, databaseRid string, collectionRid string, region string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID, filter, filter1, databaseRid, collectionRid, region)
}
func NewCollectionRegionClient(subscriptionID string, filter string, filter1 string, databaseRid string, collectionRid string, region string) CollectionRegionClient {
	return original.NewCollectionRegionClient(subscriptionID, filter, filter1, databaseRid, collectionRid, region)
}
func NewCollectionRegionClientWithBaseURI(baseURI string, subscriptionID string, filter string, filter1 string, databaseRid string, collectionRid string, region string) CollectionRegionClient {
	return original.NewCollectionRegionClientWithBaseURI(baseURI, subscriptionID, filter, filter1, databaseRid, collectionRid, region)
}
func NewOperationsClient(subscriptionID string, filter string, filter1 string, databaseRid string, collectionRid string, region string) OperationsClient {
	return original.NewOperationsClient(subscriptionID, filter, filter1, databaseRid, collectionRid, region)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string, filter string, filter1 string, databaseRid string, collectionRid string, region string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID, filter, filter1, databaseRid, collectionRid, region)
}
func NewCollectionClient(subscriptionID string, filter string, filter1 string, databaseRid string, collectionRid string, region string) CollectionClient {
	return original.NewCollectionClient(subscriptionID, filter, filter1, databaseRid, collectionRid, region)
}
func NewCollectionClientWithBaseURI(baseURI string, subscriptionID string, filter string, filter1 string, databaseRid string, collectionRid string, region string) CollectionClient {
	return original.NewCollectionClientWithBaseURI(baseURI, subscriptionID, filter, filter1, databaseRid, collectionRid, region)
}
func NewDatabaseClient(subscriptionID string, filter string, filter1 string, databaseRid string, collectionRid string, region string) DatabaseClient {
	return original.NewDatabaseClient(subscriptionID, filter, filter1, databaseRid, collectionRid, region)
}
func NewDatabaseClientWithBaseURI(baseURI string, subscriptionID string, filter string, filter1 string, databaseRid string, collectionRid string, region string) DatabaseClient {
	return original.NewDatabaseClientWithBaseURI(baseURI, subscriptionID, filter, filter1, databaseRid, collectionRid, region)
}
func NewDatabaseAccountRegionClient(subscriptionID string, filter string, filter1 string, databaseRid string, collectionRid string, region string) DatabaseAccountRegionClient {
	return original.NewDatabaseAccountRegionClient(subscriptionID, filter, filter1, databaseRid, collectionRid, region)
}
func NewDatabaseAccountRegionClientWithBaseURI(baseURI string, subscriptionID string, filter string, filter1 string, databaseRid string, collectionRid string, region string) DatabaseAccountRegionClient {
	return original.NewDatabaseAccountRegionClientWithBaseURI(baseURI, subscriptionID, filter, filter1, databaseRid, collectionRid, region)
}
func NewDatabaseAccountsClient(subscriptionID string, filter string, filter1 string, databaseRid string, collectionRid string, region string) DatabaseAccountsClient {
	return original.NewDatabaseAccountsClient(subscriptionID, filter, filter1, databaseRid, collectionRid, region)
}
func NewDatabaseAccountsClientWithBaseURI(baseURI string, subscriptionID string, filter string, filter1 string, databaseRid string, collectionRid string, region string) DatabaseAccountsClient {
	return original.NewDatabaseAccountsClientWithBaseURI(baseURI, subscriptionID, filter, filter1, databaseRid, collectionRid, region)
}
