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

package subscription

import original "github.com/Azure/azure-sdk-for-go/services/subscription/mgmt/2017-11-01-preview/subscription"

type DefinitionsOperationMetadataClient = original.DefinitionsOperationMetadataClient
type Definition = original.Definition
type DefinitionList = original.DefinitionList
type DefinitionListIterator = original.DefinitionListIterator
type DefinitionListPage = original.DefinitionListPage
type DefinitionProperties = original.DefinitionProperties
type DefinitionsCreateFuture = original.DefinitionsCreateFuture
type ErrorResponse = original.ErrorResponse
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type DefinitionsClient = original.DefinitionsClient

func New() BaseClient {
	return original.New()
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func NewDefinitionsClient() DefinitionsClient {
	return original.NewDefinitionsClient()
}
func NewDefinitionsClientWithBaseURI(baseURI string) DefinitionsClient {
	return original.NewDefinitionsClientWithBaseURI(baseURI)
}
func NewDefinitionsOperationMetadataClient() DefinitionsOperationMetadataClient {
	return original.NewDefinitionsOperationMetadataClient()
}
func NewDefinitionsOperationMetadataClientWithBaseURI(baseURI string) DefinitionsOperationMetadataClient {
	return original.NewDefinitionsOperationMetadataClientWithBaseURI(baseURI)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
