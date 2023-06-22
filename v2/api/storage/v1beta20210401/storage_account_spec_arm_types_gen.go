// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of StorageAccount_Spec. Use v1api20210401.StorageAccount_Spec instead
type StorageAccount_Spec_ARM struct {
	ExtendedLocation *ExtendedLocation_ARM                         `json:"extendedLocation,omitempty"`
	Identity         *Identity_ARM                                 `json:"identity,omitempty"`
	Kind             *StorageAccount_Kind_Spec                     `json:"kind,omitempty"`
	Location         *string                                       `json:"location,omitempty"`
	Name             string                                        `json:"name,omitempty"`
	Properties       *StorageAccountPropertiesCreateParameters_ARM `json:"properties,omitempty"`
	Sku              *Sku_ARM                                      `json:"sku,omitempty"`
	Tags             map[string]string                             `json:"tags"`
}

var _ genruntime.ARMResourceSpec = &StorageAccount_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (account StorageAccount_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (account *StorageAccount_Spec_ARM) GetName() string {
	return account.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts"
func (account *StorageAccount_Spec_ARM) GetType() string {
	return "Microsoft.Storage/storageAccounts"
}

// Deprecated version of ExtendedLocation. Use v1api20210401.ExtendedLocation instead
type ExtendedLocation_ARM struct {
	Name *string               `json:"name,omitempty"`
	Type *ExtendedLocationType `json:"type,omitempty"`
}

// Deprecated version of Identity. Use v1api20210401.Identity instead
type Identity_ARM struct {
	Type                   *Identity_Type                             `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails_ARM `json:"userAssignedIdentities,omitempty"`
}

// Deprecated version of Sku. Use v1api20210401.Sku instead
type Sku_ARM struct {
	Name *SkuName `json:"name,omitempty"`
	Tier *Tier    `json:"tier,omitempty"`
}

// Deprecated version of StorageAccount_Kind_Spec. Use v1api20210401.StorageAccount_Kind_Spec instead
// +kubebuilder:validation:Enum={"BlobStorage","BlockBlobStorage","FileStorage","Storage","StorageV2"}
type StorageAccount_Kind_Spec string

const (
	StorageAccount_Kind_Spec_BlobStorage      = StorageAccount_Kind_Spec("BlobStorage")
	StorageAccount_Kind_Spec_BlockBlobStorage = StorageAccount_Kind_Spec("BlockBlobStorage")
	StorageAccount_Kind_Spec_FileStorage      = StorageAccount_Kind_Spec("FileStorage")
	StorageAccount_Kind_Spec_Storage          = StorageAccount_Kind_Spec("Storage")
	StorageAccount_Kind_Spec_StorageV2        = StorageAccount_Kind_Spec("StorageV2")
)

// Deprecated version of StorageAccountPropertiesCreateParameters. Use v1api20210401.StorageAccountPropertiesCreateParameters instead
type StorageAccountPropertiesCreateParameters_ARM struct {
	AccessTier                            *StorageAccountPropertiesCreateParameters_AccessTier           `json:"accessTier,omitempty"`
	AllowBlobPublicAccess                 *bool                                                          `json:"allowBlobPublicAccess,omitempty"`
	AllowCrossTenantReplication           *bool                                                          `json:"allowCrossTenantReplication,omitempty"`
	AllowSharedKeyAccess                  *bool                                                          `json:"allowSharedKeyAccess,omitempty"`
	AzureFilesIdentityBasedAuthentication *AzureFilesIdentityBasedAuthentication_ARM                     `json:"azureFilesIdentityBasedAuthentication,omitempty"`
	CustomDomain                          *CustomDomain_ARM                                              `json:"customDomain,omitempty"`
	Encryption                            *Encryption_ARM                                                `json:"encryption,omitempty"`
	IsHnsEnabled                          *bool                                                          `json:"isHnsEnabled,omitempty"`
	IsNfsV3Enabled                        *bool                                                          `json:"isNfsV3Enabled,omitempty"`
	KeyPolicy                             *KeyPolicy_ARM                                                 `json:"keyPolicy,omitempty"`
	LargeFileSharesState                  *StorageAccountPropertiesCreateParameters_LargeFileSharesState `json:"largeFileSharesState,omitempty"`
	MinimumTlsVersion                     *StorageAccountPropertiesCreateParameters_MinimumTlsVersion    `json:"minimumTlsVersion,omitempty"`
	NetworkAcls                           *NetworkRuleSet_ARM                                            `json:"networkAcls,omitempty"`
	RoutingPreference                     *RoutingPreference_ARM                                         `json:"routingPreference,omitempty"`
	SasPolicy                             *SasPolicy_ARM                                                 `json:"sasPolicy,omitempty"`
	SupportsHttpsTrafficOnly              *bool                                                          `json:"supportsHttpsTrafficOnly,omitempty"`
}

// Deprecated version of AzureFilesIdentityBasedAuthentication. Use v1api20210401.AzureFilesIdentityBasedAuthentication instead
type AzureFilesIdentityBasedAuthentication_ARM struct {
	ActiveDirectoryProperties *ActiveDirectoryProperties_ARM                                 `json:"activeDirectoryProperties,omitempty"`
	DefaultSharePermission    *AzureFilesIdentityBasedAuthentication_DefaultSharePermission  `json:"defaultSharePermission,omitempty"`
	DirectoryServiceOptions   *AzureFilesIdentityBasedAuthentication_DirectoryServiceOptions `json:"directoryServiceOptions,omitempty"`
}

// Deprecated version of CustomDomain. Use v1api20210401.CustomDomain instead
type CustomDomain_ARM struct {
	Name             *string `json:"name,omitempty"`
	UseSubDomainName *bool   `json:"useSubDomainName,omitempty"`
}

// Deprecated version of Encryption. Use v1api20210401.Encryption instead
type Encryption_ARM struct {
	Identity                        *EncryptionIdentity_ARM `json:"identity,omitempty"`
	KeySource                       *Encryption_KeySource   `json:"keySource,omitempty"`
	Keyvaultproperties              *KeyVaultProperties_ARM `json:"keyvaultproperties,omitempty"`
	RequireInfrastructureEncryption *bool                   `json:"requireInfrastructureEncryption,omitempty"`
	Services                        *EncryptionServices_ARM `json:"services,omitempty"`
}

// Deprecated version of ExtendedLocationType. Use v1api20210401.ExtendedLocationType instead
// +kubebuilder:validation:Enum={"EdgeZone"}
type ExtendedLocationType string

const ExtendedLocationType_EdgeZone = ExtendedLocationType("EdgeZone")

// Deprecated version of Identity_Type. Use v1api20210401.Identity_Type instead
// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned,UserAssigned","UserAssigned"}
type Identity_Type string

const (
	Identity_Type_None                       = Identity_Type("None")
	Identity_Type_SystemAssigned             = Identity_Type("SystemAssigned")
	Identity_Type_SystemAssignedUserAssigned = Identity_Type("SystemAssigned,UserAssigned")
	Identity_Type_UserAssigned               = Identity_Type("UserAssigned")
)

// Deprecated version of KeyPolicy. Use v1api20210401.KeyPolicy instead
type KeyPolicy_ARM struct {
	KeyExpirationPeriodInDays *int `json:"keyExpirationPeriodInDays,omitempty"`
}

// Deprecated version of NetworkRuleSet. Use v1api20210401.NetworkRuleSet instead
type NetworkRuleSet_ARM struct {
	Bypass              *NetworkRuleSet_Bypass        `json:"bypass,omitempty"`
	DefaultAction       *NetworkRuleSet_DefaultAction `json:"defaultAction,omitempty"`
	IpRules             []IPRule_ARM                  `json:"ipRules"`
	ResourceAccessRules []ResourceAccessRule_ARM      `json:"resourceAccessRules"`
	VirtualNetworkRules []VirtualNetworkRule_ARM      `json:"virtualNetworkRules"`
}

// Deprecated version of RoutingPreference. Use v1api20210401.RoutingPreference instead
type RoutingPreference_ARM struct {
	PublishInternetEndpoints  *bool                            `json:"publishInternetEndpoints,omitempty"`
	PublishMicrosoftEndpoints *bool                            `json:"publishMicrosoftEndpoints,omitempty"`
	RoutingChoice             *RoutingPreference_RoutingChoice `json:"routingChoice,omitempty"`
}

// Deprecated version of SasPolicy. Use v1api20210401.SasPolicy instead
type SasPolicy_ARM struct {
	ExpirationAction    *SasPolicy_ExpirationAction `json:"expirationAction,omitempty"`
	SasExpirationPeriod *string                     `json:"sasExpirationPeriod,omitempty"`
}

// Deprecated version of SkuName. Use v1api20210401.SkuName instead
// +kubebuilder:validation:Enum={"Premium_LRS","Premium_ZRS","Standard_GRS","Standard_GZRS","Standard_LRS","Standard_RAGRS","Standard_RAGZRS","Standard_ZRS"}
type SkuName string

const (
	SkuName_Premium_LRS     = SkuName("Premium_LRS")
	SkuName_Premium_ZRS     = SkuName("Premium_ZRS")
	SkuName_Standard_GRS    = SkuName("Standard_GRS")
	SkuName_Standard_GZRS   = SkuName("Standard_GZRS")
	SkuName_Standard_LRS    = SkuName("Standard_LRS")
	SkuName_Standard_RAGRS  = SkuName("Standard_RAGRS")
	SkuName_Standard_RAGZRS = SkuName("Standard_RAGZRS")
	SkuName_Standard_ZRS    = SkuName("Standard_ZRS")
)

// Deprecated version of Tier. Use v1api20210401.Tier instead
// +kubebuilder:validation:Enum={"Premium","Standard"}
type Tier string

const (
	Tier_Premium  = Tier("Premium")
	Tier_Standard = Tier("Standard")
)

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails_ARM struct {
}

// Deprecated version of ActiveDirectoryProperties. Use v1api20210401.ActiveDirectoryProperties instead
type ActiveDirectoryProperties_ARM struct {
	AzureStorageSid   *string `json:"azureStorageSid,omitempty"`
	DomainGuid        *string `json:"domainGuid,omitempty"`
	DomainName        *string `json:"domainName,omitempty"`
	DomainSid         *string `json:"domainSid,omitempty"`
	ForestName        *string `json:"forestName,omitempty"`
	NetBiosDomainName *string `json:"netBiosDomainName,omitempty"`
}

// Deprecated version of EncryptionIdentity. Use v1api20210401.EncryptionIdentity instead
type EncryptionIdentity_ARM struct {
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}

// Deprecated version of EncryptionServices. Use v1api20210401.EncryptionServices instead
type EncryptionServices_ARM struct {
	Blob  *EncryptionService_ARM `json:"blob,omitempty"`
	File  *EncryptionService_ARM `json:"file,omitempty"`
	Queue *EncryptionService_ARM `json:"queue,omitempty"`
	Table *EncryptionService_ARM `json:"table,omitempty"`
}

// Deprecated version of IPRule. Use v1api20210401.IPRule instead
type IPRule_ARM struct {
	Action *IPRule_Action `json:"action,omitempty"`
	Value  *string        `json:"value,omitempty"`
}

// Deprecated version of KeyVaultProperties. Use v1api20210401.KeyVaultProperties instead
type KeyVaultProperties_ARM struct {
	Keyname     *string `json:"keyname,omitempty"`
	Keyvaulturi *string `json:"keyvaulturi,omitempty"`
	Keyversion  *string `json:"keyversion,omitempty"`
}

// Deprecated version of ResourceAccessRule. Use v1api20210401.ResourceAccessRule instead
type ResourceAccessRule_ARM struct {
	ResourceId *string `json:"resourceId,omitempty"`
	TenantId   *string `json:"tenantId,omitempty"`
}

// Deprecated version of VirtualNetworkRule. Use v1api20210401.VirtualNetworkRule instead
type VirtualNetworkRule_ARM struct {
	Action *VirtualNetworkRule_Action `json:"action,omitempty"`
	Id     *string                    `json:"id,omitempty"`
	State  *VirtualNetworkRule_State  `json:"state,omitempty"`
}

// Deprecated version of EncryptionService. Use v1api20210401.EncryptionService instead
type EncryptionService_ARM struct {
	Enabled *bool                      `json:"enabled,omitempty"`
	KeyType *EncryptionService_KeyType `json:"keyType,omitempty"`
}
