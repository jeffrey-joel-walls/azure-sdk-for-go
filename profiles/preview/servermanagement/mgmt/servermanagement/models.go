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

package servermanagement

import original "github.com/Azure/azure-sdk-for-go/services/servermanagement/mgmt/2016-07-01-preview/servermanagement"

type CredentialDataFormat = original.CredentialDataFormat

const (
	RsaEncrypted CredentialDataFormat = original.RsaEncrypted
)

type GatewayExpandOption = original.GatewayExpandOption

const (
	Download GatewayExpandOption = original.Download
	Status   GatewayExpandOption = original.Status
)

type PowerShellExpandOption = original.PowerShellExpandOption

const (
	Output PowerShellExpandOption = original.Output
)

type PromptFieldType = original.PromptFieldType

const (
	Credential   PromptFieldType = original.Credential
	SecureString PromptFieldType = original.SecureString
	String       PromptFieldType = original.String
)

type RetentionPeriod = original.RetentionPeriod

const (
	Persistent RetentionPeriod = original.Persistent
	Session    RetentionPeriod = original.Session
)

type UpgradeMode = original.UpgradeMode

const (
	Automatic UpgradeMode = original.Automatic
	Manual    UpgradeMode = original.Manual
)

type EncryptionJwkResource = original.EncryptionJwkResource
type Error = original.Error
type GatewayCreateFuture = original.GatewayCreateFuture
type GatewayGetProfileFuture = original.GatewayGetProfileFuture
type GatewayParameters = original.GatewayParameters
type GatewayParametersProperties = original.GatewayParametersProperties
type GatewayProfile = original.GatewayProfile
type GatewayRegenerateProfileFuture = original.GatewayRegenerateProfileFuture
type GatewayResource = original.GatewayResource
type GatewayResourceProperties = original.GatewayResourceProperties
type GatewayResources = original.GatewayResources
type GatewayResourcesIterator = original.GatewayResourcesIterator
type GatewayResourcesPage = original.GatewayResourcesPage
type GatewayStatus = original.GatewayStatus
type GatewayUpdateFuture = original.GatewayUpdateFuture
type GatewayUpgradeFuture = original.GatewayUpgradeFuture
type NodeCreateFuture = original.NodeCreateFuture
type NodeParameters = original.NodeParameters
type NodeParametersProperties = original.NodeParametersProperties
type NodeResource = original.NodeResource
type NodeResourceProperties = original.NodeResourceProperties
type NodeResources = original.NodeResources
type NodeResourcesIterator = original.NodeResourcesIterator
type NodeResourcesPage = original.NodeResourcesPage
type NodeUpdateFuture = original.NodeUpdateFuture
type PowerShellCancelCommandFuture = original.PowerShellCancelCommandFuture
type PowerShellCommandParameters = original.PowerShellCommandParameters
type PowerShellCommandParametersProperties = original.PowerShellCommandParametersProperties
type PowerShellCommandResult = original.PowerShellCommandResult
type PowerShellCommandResults = original.PowerShellCommandResults
type PowerShellCommandStatus = original.PowerShellCommandStatus
type PowerShellCreateSessionFuture = original.PowerShellCreateSessionFuture
type PowerShellInvokeCommandFuture = original.PowerShellInvokeCommandFuture
type PowerShellSessionResource = original.PowerShellSessionResource
type PowerShellSessionResourceProperties = original.PowerShellSessionResourceProperties
type PowerShellSessionResources = original.PowerShellSessionResources
type PowerShellTabCompletionParameters = original.PowerShellTabCompletionParameters
type PowerShellTabCompletionResults = original.PowerShellTabCompletionResults
type PowerShellUpdateCommandFuture = original.PowerShellUpdateCommandFuture
type PromptFieldDescription = original.PromptFieldDescription
type PromptMessageResponse = original.PromptMessageResponse
type Resource = original.Resource
type SessionCreateFuture = original.SessionCreateFuture
type SessionParameters = original.SessionParameters
type SessionParametersProperties = original.SessionParametersProperties
type SessionResource = original.SessionResource
type SessionResourceProperties = original.SessionResourceProperties
type VersionServermanagement = original.VersionServermanagement
type NodeClient = original.NodeClient
type PowerShellClient = original.PowerShellClient
type SessionClient = original.SessionClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type GatewayClient = original.GatewayClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewGatewayClient(subscriptionID string) GatewayClient {
	return original.NewGatewayClient(subscriptionID)
}
func NewGatewayClientWithBaseURI(baseURI string, subscriptionID string) GatewayClient {
	return original.NewGatewayClientWithBaseURI(baseURI, subscriptionID)
}
func NewNodeClient(subscriptionID string) NodeClient {
	return original.NewNodeClient(subscriptionID)
}
func NewNodeClientWithBaseURI(baseURI string, subscriptionID string) NodeClient {
	return original.NewNodeClientWithBaseURI(baseURI, subscriptionID)
}
func NewPowerShellClient(subscriptionID string) PowerShellClient {
	return original.NewPowerShellClient(subscriptionID)
}
func NewPowerShellClientWithBaseURI(baseURI string, subscriptionID string) PowerShellClient {
	return original.NewPowerShellClientWithBaseURI(baseURI, subscriptionID)
}
func NewSessionClient(subscriptionID string) SessionClient {
	return original.NewSessionClient(subscriptionID)
}
func NewSessionClientWithBaseURI(baseURI string, subscriptionID string) SessionClient {
	return original.NewSessionClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
