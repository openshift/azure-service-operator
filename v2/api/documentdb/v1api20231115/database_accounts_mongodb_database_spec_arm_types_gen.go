// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20231115

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DatabaseAccounts_MongodbDatabase_Spec_ARM struct {
	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties to create and update Azure Cosmos DB MongoDB database.
	Properties *MongoDBDatabaseCreateUpdateProperties_ARM `json:"properties,omitempty"`
	Tags       map[string]string                          `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccounts_MongodbDatabase_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-11-15"
func (database DatabaseAccounts_MongodbDatabase_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (database *DatabaseAccounts_MongodbDatabase_Spec_ARM) GetName() string {
	return database.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases"
func (database *DatabaseAccounts_MongodbDatabase_Spec_ARM) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases"
}

// Properties to create and update Azure Cosmos DB MongoDB database.
type MongoDBDatabaseCreateUpdateProperties_ARM struct {
	// Options: A key-value pair of options to be applied for the request. This corresponds to the headers sent with the
	// request.
	Options *CreateUpdateOptions_ARM `json:"options,omitempty"`

	// Resource: The standard JSON format of a MongoDB database
	Resource *MongoDBDatabaseResource_ARM `json:"resource,omitempty"`
}

// CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are "If-Match",
// "If-None-Match", "Session-Token" and "Throughput"
type CreateUpdateOptions_ARM struct {
	// AutoscaleSettings: Specifies the Autoscale settings. Note: Either throughput or autoscaleSettings is required, but not
	// both.
	AutoscaleSettings *AutoscaleSettings_ARM `json:"autoscaleSettings,omitempty"`

	// Throughput: Request Units per second. For example, "throughput": 10000.
	Throughput *int `json:"throughput,omitempty"`
}

// Cosmos DB MongoDB database resource object
type MongoDBDatabaseResource_ARM struct {
	// CreateMode: Enum to indicate the mode of resource creation.
	CreateMode *CreateMode `json:"createMode,omitempty"`

	// Id: Name of the Cosmos DB MongoDB database
	Id *string `json:"id,omitempty"`

	// RestoreParameters: Parameters to indicate the information about the restore
	RestoreParameters *RestoreParametersBase_ARM `json:"restoreParameters,omitempty"`
}

type AutoscaleSettings_ARM struct {
	// MaxThroughput: Represents maximum throughput, the resource can scale up to.
	MaxThroughput *int `json:"maxThroughput,omitempty"`
}

// Parameters to indicate the information about the restore.
type RestoreParametersBase_ARM struct {
	// RestoreSource: The id of the restorable database account from which the restore has to be initiated. For example:
	// /subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{restorableDatabaseAccountName}
	RestoreSource *string `json:"restoreSource,omitempty"`

	// RestoreTimestampInUtc: Time to which the account has to be restored (ISO-8601 format).
	RestoreTimestampInUtc *string `json:"restoreTimestampInUtc,omitempty"`
}