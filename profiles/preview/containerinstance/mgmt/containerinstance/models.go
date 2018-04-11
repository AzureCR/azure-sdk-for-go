// +build go1.9

// Copyright 2018 Microsoft Corporation
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

package containerinstance

import original "github.com/Azure/azure-sdk-for-go/services/containerinstance/mgmt/2018-02-01-preview/containerinstance"

type OperationsClient = original.OperationsClient

func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}

type ContainerLogsClient = original.ContainerLogsClient

func NewContainerLogsClient(subscriptionID string) ContainerLogsClient {
	return original.NewContainerLogsClient(subscriptionID)
}
func NewContainerLogsClientWithBaseURI(baseURI string, subscriptionID string) ContainerLogsClient {
	return original.NewContainerLogsClientWithBaseURI(baseURI, subscriptionID)
}

type ContainerGroupUsageClient = original.ContainerGroupUsageClient

func NewContainerGroupUsageClient(subscriptionID string) ContainerGroupUsageClient {
	return original.NewContainerGroupUsageClient(subscriptionID)
}
func NewContainerGroupUsageClientWithBaseURI(baseURI string, subscriptionID string) ContainerGroupUsageClient {
	return original.NewContainerGroupUsageClientWithBaseURI(baseURI, subscriptionID)
}

type ContainerGroupNetworkProtocol = original.ContainerGroupNetworkProtocol

const (
	TCP ContainerGroupNetworkProtocol = original.TCP
	UDP ContainerGroupNetworkProtocol = original.UDP
)

func PossibleContainerGroupNetworkProtocolValues() []ContainerGroupNetworkProtocol {
	return original.PossibleContainerGroupNetworkProtocolValues()
}

type ContainerGroupRestartPolicy = original.ContainerGroupRestartPolicy

const (
	Always    ContainerGroupRestartPolicy = original.Always
	Never     ContainerGroupRestartPolicy = original.Never
	OnFailure ContainerGroupRestartPolicy = original.OnFailure
)

func PossibleContainerGroupRestartPolicyValues() []ContainerGroupRestartPolicy {
	return original.PossibleContainerGroupRestartPolicyValues()
}

type ContainerNetworkProtocol = original.ContainerNetworkProtocol

const (
	ContainerNetworkProtocolTCP ContainerNetworkProtocol = original.ContainerNetworkProtocolTCP
	ContainerNetworkProtocolUDP ContainerNetworkProtocol = original.ContainerNetworkProtocolUDP
)

func PossibleContainerNetworkProtocolValues() []ContainerNetworkProtocol {
	return original.PossibleContainerNetworkProtocolValues()
}

type OperatingSystemTypes = original.OperatingSystemTypes

const (
	Linux   OperatingSystemTypes = original.Linux
	Windows OperatingSystemTypes = original.Windows
)

func PossibleOperatingSystemTypesValues() []OperatingSystemTypes {
	return original.PossibleOperatingSystemTypesValues()
}

type OperationsOrigin = original.OperationsOrigin

const (
	System OperationsOrigin = original.System
	User   OperationsOrigin = original.User
)

func PossibleOperationsOriginValues() []OperationsOrigin {
	return original.PossibleOperationsOriginValues()
}

type AzureFileVolume = original.AzureFileVolume
type Container = original.Container
type ContainerExecRequest = original.ContainerExecRequest
type ContainerExecRequestTerminalSize = original.ContainerExecRequestTerminalSize
type ContainerExecResponse = original.ContainerExecResponse
type ContainerGroup = original.ContainerGroup
type ContainerGroupListResult = original.ContainerGroupListResult
type ContainerGroupListResultIterator = original.ContainerGroupListResultIterator
type ContainerGroupListResultPage = original.ContainerGroupListResultPage
type ContainerGroupProperties = original.ContainerGroupProperties
type ContainerGroupPropertiesInstanceView = original.ContainerGroupPropertiesInstanceView
type ContainerGroupsCreateOrUpdateFuture = original.ContainerGroupsCreateOrUpdateFuture
type ContainerPort = original.ContainerPort
type ContainerProperties = original.ContainerProperties
type ContainerPropertiesInstanceView = original.ContainerPropertiesInstanceView
type ContainerState = original.ContainerState
type EnvironmentVariable = original.EnvironmentVariable
type Event = original.Event
type GitRepoVolume = original.GitRepoVolume
type ImageRegistryCredential = original.ImageRegistryCredential
type IPAddress = original.IPAddress
type Logs = original.Logs
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type Port = original.Port
type Resource = original.Resource
type ResourceLimits = original.ResourceLimits
type ResourceRequests = original.ResourceRequests
type ResourceRequirements = original.ResourceRequirements
type Usage = original.Usage
type UsageListResult = original.UsageListResult
type UsageName = original.UsageName
type Volume = original.Volume
type VolumeMount = original.VolumeMount

func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}

type ContainerGroupsClient = original.ContainerGroupsClient

func NewContainerGroupsClient(subscriptionID string) ContainerGroupsClient {
	return original.NewContainerGroupsClient(subscriptionID)
}
func NewContainerGroupsClientWithBaseURI(baseURI string, subscriptionID string) ContainerGroupsClient {
	return original.NewContainerGroupsClientWithBaseURI(baseURI, subscriptionID)
}

type StartContainerClient = original.StartContainerClient

func NewStartContainerClient(subscriptionID string) StartContainerClient {
	return original.NewStartContainerClient(subscriptionID)
}
func NewStartContainerClientWithBaseURI(baseURI string, subscriptionID string) StartContainerClient {
	return original.NewStartContainerClientWithBaseURI(baseURI, subscriptionID)
}

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
