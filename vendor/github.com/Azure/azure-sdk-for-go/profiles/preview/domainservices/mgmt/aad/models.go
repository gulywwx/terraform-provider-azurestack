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

package aad

import original "github.com/Azure/azure-sdk-for-go/services/domainservices/mgmt/2017-06-01/aad"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type DomainServiceOperationsClient = original.DomainServiceOperationsClient
type DomainServicesClient = original.DomainServicesClient
type ExternalAccess = original.ExternalAccess

const (
	Disabled ExternalAccess = original.Disabled
	Enabled  ExternalAccess = original.Enabled
)

type Ldaps = original.Ldaps

const (
	LdapsDisabled Ldaps = original.LdapsDisabled
	LdapsEnabled  Ldaps = original.LdapsEnabled
)

type DomainService = original.DomainService
type DomainServiceListResult = original.DomainServiceListResult
type DomainServicePatchProperties = original.DomainServicePatchProperties
type DomainServiceProperties = original.DomainServiceProperties
type DomainServicesCreateOrUpdateFuture = original.DomainServicesCreateOrUpdateFuture
type DomainServicesDeleteFuture = original.DomainServicesDeleteFuture
type DomainServicesUpdateFuture = original.DomainServicesUpdateFuture
type LdapsSettings = original.LdapsSettings
type OperationDisplayInfo = original.OperationDisplayInfo
type OperationEntity = original.OperationEntity
type OperationEntityListResult = original.OperationEntityListResult
type Resource = original.Resource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewDomainServiceOperationsClient(subscriptionID string) DomainServiceOperationsClient {
	return original.NewDomainServiceOperationsClient(subscriptionID)
}
func NewDomainServiceOperationsClientWithBaseURI(baseURI string, subscriptionID string) DomainServiceOperationsClient {
	return original.NewDomainServiceOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDomainServicesClient(subscriptionID string) DomainServicesClient {
	return original.NewDomainServicesClient(subscriptionID)
}
func NewDomainServicesClientWithBaseURI(baseURI string, subscriptionID string) DomainServicesClient {
	return original.NewDomainServicesClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleExternalAccessValues() []ExternalAccess {
	return original.PossibleExternalAccessValues()
}
func PossibleLdapsValues() []Ldaps {
	return original.PossibleLdapsValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
