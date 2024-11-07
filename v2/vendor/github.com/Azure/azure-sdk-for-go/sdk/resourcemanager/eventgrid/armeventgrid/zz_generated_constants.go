//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

const (
	moduleName    = "armeventgrid"
	moduleVersion = "v1.0.0"
)

// AdvancedFilterOperatorType - The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
type AdvancedFilterOperatorType string

const (
	AdvancedFilterOperatorTypeBoolEquals                AdvancedFilterOperatorType = "BoolEquals"
	AdvancedFilterOperatorTypeIsNotNull                 AdvancedFilterOperatorType = "IsNotNull"
	AdvancedFilterOperatorTypeIsNullOrUndefined         AdvancedFilterOperatorType = "IsNullOrUndefined"
	AdvancedFilterOperatorTypeNumberGreaterThan         AdvancedFilterOperatorType = "NumberGreaterThan"
	AdvancedFilterOperatorTypeNumberGreaterThanOrEquals AdvancedFilterOperatorType = "NumberGreaterThanOrEquals"
	AdvancedFilterOperatorTypeNumberIn                  AdvancedFilterOperatorType = "NumberIn"
	AdvancedFilterOperatorTypeNumberInRange             AdvancedFilterOperatorType = "NumberInRange"
	AdvancedFilterOperatorTypeNumberLessThan            AdvancedFilterOperatorType = "NumberLessThan"
	AdvancedFilterOperatorTypeNumberLessThanOrEquals    AdvancedFilterOperatorType = "NumberLessThanOrEquals"
	AdvancedFilterOperatorTypeNumberNotIn               AdvancedFilterOperatorType = "NumberNotIn"
	AdvancedFilterOperatorTypeNumberNotInRange          AdvancedFilterOperatorType = "NumberNotInRange"
	AdvancedFilterOperatorTypeStringBeginsWith          AdvancedFilterOperatorType = "StringBeginsWith"
	AdvancedFilterOperatorTypeStringContains            AdvancedFilterOperatorType = "StringContains"
	AdvancedFilterOperatorTypeStringEndsWith            AdvancedFilterOperatorType = "StringEndsWith"
	AdvancedFilterOperatorTypeStringIn                  AdvancedFilterOperatorType = "StringIn"
	AdvancedFilterOperatorTypeStringNotBeginsWith       AdvancedFilterOperatorType = "StringNotBeginsWith"
	AdvancedFilterOperatorTypeStringNotContains         AdvancedFilterOperatorType = "StringNotContains"
	AdvancedFilterOperatorTypeStringNotEndsWith         AdvancedFilterOperatorType = "StringNotEndsWith"
	AdvancedFilterOperatorTypeStringNotIn               AdvancedFilterOperatorType = "StringNotIn"
)

// PossibleAdvancedFilterOperatorTypeValues returns the possible values for the AdvancedFilterOperatorType const type.
func PossibleAdvancedFilterOperatorTypeValues() []AdvancedFilterOperatorType {
	return []AdvancedFilterOperatorType{
		AdvancedFilterOperatorTypeBoolEquals,
		AdvancedFilterOperatorTypeIsNotNull,
		AdvancedFilterOperatorTypeIsNullOrUndefined,
		AdvancedFilterOperatorTypeNumberGreaterThan,
		AdvancedFilterOperatorTypeNumberGreaterThanOrEquals,
		AdvancedFilterOperatorTypeNumberIn,
		AdvancedFilterOperatorTypeNumberInRange,
		AdvancedFilterOperatorTypeNumberLessThan,
		AdvancedFilterOperatorTypeNumberLessThanOrEquals,
		AdvancedFilterOperatorTypeNumberNotIn,
		AdvancedFilterOperatorTypeNumberNotInRange,
		AdvancedFilterOperatorTypeStringBeginsWith,
		AdvancedFilterOperatorTypeStringContains,
		AdvancedFilterOperatorTypeStringEndsWith,
		AdvancedFilterOperatorTypeStringIn,
		AdvancedFilterOperatorTypeStringNotBeginsWith,
		AdvancedFilterOperatorTypeStringNotContains,
		AdvancedFilterOperatorTypeStringNotEndsWith,
		AdvancedFilterOperatorTypeStringNotIn,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// DeadLetterEndPointType - Type of the endpoint for the dead letter destination
type DeadLetterEndPointType string

const (
	DeadLetterEndPointTypeStorageBlob DeadLetterEndPointType = "StorageBlob"
)

// PossibleDeadLetterEndPointTypeValues returns the possible values for the DeadLetterEndPointType const type.
func PossibleDeadLetterEndPointTypeValues() []DeadLetterEndPointType {
	return []DeadLetterEndPointType{
		DeadLetterEndPointTypeStorageBlob,
	}
}

// DeliveryAttributeMappingType - Type of the delivery attribute or header name.
type DeliveryAttributeMappingType string

const (
	DeliveryAttributeMappingTypeDynamic DeliveryAttributeMappingType = "Dynamic"
	DeliveryAttributeMappingTypeStatic  DeliveryAttributeMappingType = "Static"
)

// PossibleDeliveryAttributeMappingTypeValues returns the possible values for the DeliveryAttributeMappingType const type.
func PossibleDeliveryAttributeMappingTypeValues() []DeliveryAttributeMappingType {
	return []DeliveryAttributeMappingType{
		DeliveryAttributeMappingTypeDynamic,
		DeliveryAttributeMappingTypeStatic,
	}
}

// DomainProvisioningState - Provisioning state of the Event Grid Domain Resource.
type DomainProvisioningState string

const (
	DomainProvisioningStateCanceled  DomainProvisioningState = "Canceled"
	DomainProvisioningStateCreating  DomainProvisioningState = "Creating"
	DomainProvisioningStateDeleting  DomainProvisioningState = "Deleting"
	DomainProvisioningStateFailed    DomainProvisioningState = "Failed"
	DomainProvisioningStateSucceeded DomainProvisioningState = "Succeeded"
	DomainProvisioningStateUpdating  DomainProvisioningState = "Updating"
)

// PossibleDomainProvisioningStateValues returns the possible values for the DomainProvisioningState const type.
func PossibleDomainProvisioningStateValues() []DomainProvisioningState {
	return []DomainProvisioningState{
		DomainProvisioningStateCanceled,
		DomainProvisioningStateCreating,
		DomainProvisioningStateDeleting,
		DomainProvisioningStateFailed,
		DomainProvisioningStateSucceeded,
		DomainProvisioningStateUpdating,
	}
}

// DomainTopicProvisioningState - Provisioning state of the domain topic.
type DomainTopicProvisioningState string

const (
	DomainTopicProvisioningStateCanceled  DomainTopicProvisioningState = "Canceled"
	DomainTopicProvisioningStateCreating  DomainTopicProvisioningState = "Creating"
	DomainTopicProvisioningStateDeleting  DomainTopicProvisioningState = "Deleting"
	DomainTopicProvisioningStateFailed    DomainTopicProvisioningState = "Failed"
	DomainTopicProvisioningStateSucceeded DomainTopicProvisioningState = "Succeeded"
	DomainTopicProvisioningStateUpdating  DomainTopicProvisioningState = "Updating"
)

// PossibleDomainTopicProvisioningStateValues returns the possible values for the DomainTopicProvisioningState const type.
func PossibleDomainTopicProvisioningStateValues() []DomainTopicProvisioningState {
	return []DomainTopicProvisioningState{
		DomainTopicProvisioningStateCanceled,
		DomainTopicProvisioningStateCreating,
		DomainTopicProvisioningStateDeleting,
		DomainTopicProvisioningStateFailed,
		DomainTopicProvisioningStateSucceeded,
		DomainTopicProvisioningStateUpdating,
	}
}

// EndpointType - Type of the endpoint for the event subscription destination.
type EndpointType string

const (
	EndpointTypeAzureFunction    EndpointType = "AzureFunction"
	EndpointTypeEventHub         EndpointType = "EventHub"
	EndpointTypeHybridConnection EndpointType = "HybridConnection"
	EndpointTypeServiceBusQueue  EndpointType = "ServiceBusQueue"
	EndpointTypeServiceBusTopic  EndpointType = "ServiceBusTopic"
	EndpointTypeStorageQueue     EndpointType = "StorageQueue"
	EndpointTypeWebHook          EndpointType = "WebHook"
)

// PossibleEndpointTypeValues returns the possible values for the EndpointType const type.
func PossibleEndpointTypeValues() []EndpointType {
	return []EndpointType{
		EndpointTypeAzureFunction,
		EndpointTypeEventHub,
		EndpointTypeHybridConnection,
		EndpointTypeServiceBusQueue,
		EndpointTypeServiceBusTopic,
		EndpointTypeStorageQueue,
		EndpointTypeWebHook,
	}
}

// EventDeliverySchema - The event delivery schema for the event subscription.
type EventDeliverySchema string

const (
	EventDeliverySchemaCloudEventSchemaV10 EventDeliverySchema = "CloudEventSchemaV1_0"
	EventDeliverySchemaCustomInputSchema   EventDeliverySchema = "CustomInputSchema"
	EventDeliverySchemaEventGridSchema     EventDeliverySchema = "EventGridSchema"
)

// PossibleEventDeliverySchemaValues returns the possible values for the EventDeliverySchema const type.
func PossibleEventDeliverySchemaValues() []EventDeliverySchema {
	return []EventDeliverySchema{
		EventDeliverySchemaCloudEventSchemaV10,
		EventDeliverySchemaCustomInputSchema,
		EventDeliverySchemaEventGridSchema,
	}
}

// EventSubscriptionIdentityType - The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both
// an implicitly created identity and a set of user-assigned identities. The type 'None' will remove any identity.
type EventSubscriptionIdentityType string

const (
	EventSubscriptionIdentityTypeSystemAssigned EventSubscriptionIdentityType = "SystemAssigned"
	EventSubscriptionIdentityTypeUserAssigned   EventSubscriptionIdentityType = "UserAssigned"
)

// PossibleEventSubscriptionIdentityTypeValues returns the possible values for the EventSubscriptionIdentityType const type.
func PossibleEventSubscriptionIdentityTypeValues() []EventSubscriptionIdentityType {
	return []EventSubscriptionIdentityType{
		EventSubscriptionIdentityTypeSystemAssigned,
		EventSubscriptionIdentityTypeUserAssigned,
	}
}

// EventSubscriptionProvisioningState - Provisioning state of the event subscription.
type EventSubscriptionProvisioningState string

const (
	EventSubscriptionProvisioningStateAwaitingManualAction EventSubscriptionProvisioningState = "AwaitingManualAction"
	EventSubscriptionProvisioningStateCanceled             EventSubscriptionProvisioningState = "Canceled"
	EventSubscriptionProvisioningStateCreating             EventSubscriptionProvisioningState = "Creating"
	EventSubscriptionProvisioningStateDeleting             EventSubscriptionProvisioningState = "Deleting"
	EventSubscriptionProvisioningStateFailed               EventSubscriptionProvisioningState = "Failed"
	EventSubscriptionProvisioningStateSucceeded            EventSubscriptionProvisioningState = "Succeeded"
	EventSubscriptionProvisioningStateUpdating             EventSubscriptionProvisioningState = "Updating"
)

// PossibleEventSubscriptionProvisioningStateValues returns the possible values for the EventSubscriptionProvisioningState const type.
func PossibleEventSubscriptionProvisioningStateValues() []EventSubscriptionProvisioningState {
	return []EventSubscriptionProvisioningState{
		EventSubscriptionProvisioningStateAwaitingManualAction,
		EventSubscriptionProvisioningStateCanceled,
		EventSubscriptionProvisioningStateCreating,
		EventSubscriptionProvisioningStateDeleting,
		EventSubscriptionProvisioningStateFailed,
		EventSubscriptionProvisioningStateSucceeded,
		EventSubscriptionProvisioningStateUpdating,
	}
}

// IPActionType - Action to perform based on the match or no match of the IpMask.
type IPActionType string

const (
	IPActionTypeAllow IPActionType = "Allow"
)

// PossibleIPActionTypeValues returns the possible values for the IPActionType const type.
func PossibleIPActionTypeValues() []IPActionType {
	return []IPActionType{
		IPActionTypeAllow,
	}
}

// IdentityType - The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created
// identity and a set of user-assigned identities. The type 'None' will remove any identity.
type IdentityType string

const (
	IdentityTypeNone                       IdentityType = "None"
	IdentityTypeSystemAssigned             IdentityType = "SystemAssigned"
	IdentityTypeSystemAssignedUserAssigned IdentityType = "SystemAssigned, UserAssigned"
	IdentityTypeUserAssigned               IdentityType = "UserAssigned"
)

// PossibleIdentityTypeValues returns the possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{
		IdentityTypeNone,
		IdentityTypeSystemAssigned,
		IdentityTypeSystemAssignedUserAssigned,
		IdentityTypeUserAssigned,
	}
}

// InputSchema - This determines the format that Event Grid should expect for incoming events published to the domain.
type InputSchema string

const (
	InputSchemaCloudEventSchemaV10 InputSchema = "CloudEventSchemaV1_0"
	InputSchemaCustomEventSchema   InputSchema = "CustomEventSchema"
	InputSchemaEventGridSchema     InputSchema = "EventGridSchema"
)

// PossibleInputSchemaValues returns the possible values for the InputSchema const type.
func PossibleInputSchemaValues() []InputSchema {
	return []InputSchema{
		InputSchemaCloudEventSchemaV10,
		InputSchemaCustomEventSchema,
		InputSchemaEventGridSchema,
	}
}

// InputSchemaMappingType - Type of the custom mapping
type InputSchemaMappingType string

const (
	InputSchemaMappingTypeJSON InputSchemaMappingType = "Json"
)

// PossibleInputSchemaMappingTypeValues returns the possible values for the InputSchemaMappingType const type.
func PossibleInputSchemaMappingTypeValues() []InputSchemaMappingType {
	return []InputSchemaMappingType{
		InputSchemaMappingTypeJSON,
	}
}

// PersistedConnectionStatus - Status of the connection.
type PersistedConnectionStatus string

const (
	PersistedConnectionStatusApproved     PersistedConnectionStatus = "Approved"
	PersistedConnectionStatusDisconnected PersistedConnectionStatus = "Disconnected"
	PersistedConnectionStatusPending      PersistedConnectionStatus = "Pending"
	PersistedConnectionStatusRejected     PersistedConnectionStatus = "Rejected"
)

// PossiblePersistedConnectionStatusValues returns the possible values for the PersistedConnectionStatus const type.
func PossiblePersistedConnectionStatusValues() []PersistedConnectionStatus {
	return []PersistedConnectionStatus{
		PersistedConnectionStatusApproved,
		PersistedConnectionStatusDisconnected,
		PersistedConnectionStatusPending,
		PersistedConnectionStatusRejected,
	}
}

type PrivateEndpointConnectionsParentType string

const (
	PrivateEndpointConnectionsParentTypeDomains PrivateEndpointConnectionsParentType = "domains"
	PrivateEndpointConnectionsParentTypeTopics  PrivateEndpointConnectionsParentType = "topics"
)

// PossiblePrivateEndpointConnectionsParentTypeValues returns the possible values for the PrivateEndpointConnectionsParentType const type.
func PossiblePrivateEndpointConnectionsParentTypeValues() []PrivateEndpointConnectionsParentType {
	return []PrivateEndpointConnectionsParentType{
		PrivateEndpointConnectionsParentTypeDomains,
		PrivateEndpointConnectionsParentTypeTopics,
	}
}

// PublicNetworkAccess - This determines if traffic is allowed over public network. By default it is enabled. You can further
// restrict to specific IPs by configuring
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
	}
}

// ResourceProvisioningState - Provisioning state of the Private Endpoint Connection.
type ResourceProvisioningState string

const (
	ResourceProvisioningStateCanceled  ResourceProvisioningState = "Canceled"
	ResourceProvisioningStateCreating  ResourceProvisioningState = "Creating"
	ResourceProvisioningStateDeleting  ResourceProvisioningState = "Deleting"
	ResourceProvisioningStateFailed    ResourceProvisioningState = "Failed"
	ResourceProvisioningStateSucceeded ResourceProvisioningState = "Succeeded"
	ResourceProvisioningStateUpdating  ResourceProvisioningState = "Updating"
)

// PossibleResourceProvisioningStateValues returns the possible values for the ResourceProvisioningState const type.
func PossibleResourceProvisioningStateValues() []ResourceProvisioningState {
	return []ResourceProvisioningState{
		ResourceProvisioningStateCanceled,
		ResourceProvisioningStateCreating,
		ResourceProvisioningStateDeleting,
		ResourceProvisioningStateFailed,
		ResourceProvisioningStateSucceeded,
		ResourceProvisioningStateUpdating,
	}
}

// ResourceRegionType - Region type of the resource.
type ResourceRegionType string

const (
	ResourceRegionTypeGlobalResource   ResourceRegionType = "GlobalResource"
	ResourceRegionTypeRegionalResource ResourceRegionType = "RegionalResource"
)

// PossibleResourceRegionTypeValues returns the possible values for the ResourceRegionType const type.
func PossibleResourceRegionTypeValues() []ResourceRegionType {
	return []ResourceRegionType{
		ResourceRegionTypeGlobalResource,
		ResourceRegionTypeRegionalResource,
	}
}

// TopicProvisioningState - Provisioning state of the topic.
type TopicProvisioningState string

const (
	TopicProvisioningStateCanceled  TopicProvisioningState = "Canceled"
	TopicProvisioningStateCreating  TopicProvisioningState = "Creating"
	TopicProvisioningStateDeleting  TopicProvisioningState = "Deleting"
	TopicProvisioningStateFailed    TopicProvisioningState = "Failed"
	TopicProvisioningStateSucceeded TopicProvisioningState = "Succeeded"
	TopicProvisioningStateUpdating  TopicProvisioningState = "Updating"
)

// PossibleTopicProvisioningStateValues returns the possible values for the TopicProvisioningState const type.
func PossibleTopicProvisioningStateValues() []TopicProvisioningState {
	return []TopicProvisioningState{
		TopicProvisioningStateCanceled,
		TopicProvisioningStateCreating,
		TopicProvisioningStateDeleting,
		TopicProvisioningStateFailed,
		TopicProvisioningStateSucceeded,
		TopicProvisioningStateUpdating,
	}
}

type TopicTypePropertiesSupportedScopesForSourceItem string

const (
	TopicTypePropertiesSupportedScopesForSourceItemAzureSubscription TopicTypePropertiesSupportedScopesForSourceItem = "AzureSubscription"
	TopicTypePropertiesSupportedScopesForSourceItemResource          TopicTypePropertiesSupportedScopesForSourceItem = "Resource"
	TopicTypePropertiesSupportedScopesForSourceItemResourceGroup     TopicTypePropertiesSupportedScopesForSourceItem = "ResourceGroup"
)

// PossibleTopicTypePropertiesSupportedScopesForSourceItemValues returns the possible values for the TopicTypePropertiesSupportedScopesForSourceItem const type.
func PossibleTopicTypePropertiesSupportedScopesForSourceItemValues() []TopicTypePropertiesSupportedScopesForSourceItem {
	return []TopicTypePropertiesSupportedScopesForSourceItem{
		TopicTypePropertiesSupportedScopesForSourceItemAzureSubscription,
		TopicTypePropertiesSupportedScopesForSourceItemResource,
		TopicTypePropertiesSupportedScopesForSourceItemResourceGroup,
	}
}

// TopicTypeProvisioningState - Provisioning state of the topic type
type TopicTypeProvisioningState string

const (
	TopicTypeProvisioningStateCanceled  TopicTypeProvisioningState = "Canceled"
	TopicTypeProvisioningStateCreating  TopicTypeProvisioningState = "Creating"
	TopicTypeProvisioningStateDeleting  TopicTypeProvisioningState = "Deleting"
	TopicTypeProvisioningStateFailed    TopicTypeProvisioningState = "Failed"
	TopicTypeProvisioningStateSucceeded TopicTypeProvisioningState = "Succeeded"
	TopicTypeProvisioningStateUpdating  TopicTypeProvisioningState = "Updating"
)

// PossibleTopicTypeProvisioningStateValues returns the possible values for the TopicTypeProvisioningState const type.
func PossibleTopicTypeProvisioningStateValues() []TopicTypeProvisioningState {
	return []TopicTypeProvisioningState{
		TopicTypeProvisioningStateCanceled,
		TopicTypeProvisioningStateCreating,
		TopicTypeProvisioningStateDeleting,
		TopicTypeProvisioningStateFailed,
		TopicTypeProvisioningStateSucceeded,
		TopicTypeProvisioningStateUpdating,
	}
}
