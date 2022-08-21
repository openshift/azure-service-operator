// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211001

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_SubscriptionAliasResponse_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubscriptionAliasResponse_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubscriptionAliasResponseSTATUSARM, SubscriptionAliasResponseSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubscriptionAliasResponseSTATUSARM runs a test to see if a specific instance of SubscriptionAliasResponse_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubscriptionAliasResponseSTATUSARM(subject SubscriptionAliasResponse_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubscriptionAliasResponse_STATUSARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SubscriptionAliasResponse_STATUSARM instances for property testing - lazily instantiated by
// SubscriptionAliasResponseSTATUSARMGenerator()
var subscriptionAliasResponseSTATUSARMGenerator gopter.Gen

// SubscriptionAliasResponseSTATUSARMGenerator returns a generator of SubscriptionAliasResponse_STATUSARM instances for property testing.
// We first initialize subscriptionAliasResponseSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SubscriptionAliasResponseSTATUSARMGenerator() gopter.Gen {
	if subscriptionAliasResponseSTATUSARMGenerator != nil {
		return subscriptionAliasResponseSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubscriptionAliasResponseSTATUSARM(generators)
	subscriptionAliasResponseSTATUSARMGenerator = gen.Struct(reflect.TypeOf(SubscriptionAliasResponse_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubscriptionAliasResponseSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForSubscriptionAliasResponseSTATUSARM(generators)
	subscriptionAliasResponseSTATUSARMGenerator = gen.Struct(reflect.TypeOf(SubscriptionAliasResponse_STATUSARM{}), generators)

	return subscriptionAliasResponseSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSubscriptionAliasResponseSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubscriptionAliasResponseSTATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSubscriptionAliasResponseSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSubscriptionAliasResponseSTATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SubscriptionAliasResponsePropertiesSTATUSARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataSTATUSARMGenerator())
}

func Test_SubscriptionAliasResponseProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubscriptionAliasResponseProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubscriptionAliasResponsePropertiesSTATUSARM, SubscriptionAliasResponsePropertiesSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubscriptionAliasResponsePropertiesSTATUSARM runs a test to see if a specific instance of SubscriptionAliasResponseProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubscriptionAliasResponsePropertiesSTATUSARM(subject SubscriptionAliasResponseProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubscriptionAliasResponseProperties_STATUSARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SubscriptionAliasResponseProperties_STATUSARM instances for property testing - lazily instantiated by
// SubscriptionAliasResponsePropertiesSTATUSARMGenerator()
var subscriptionAliasResponsePropertiesSTATUSARMGenerator gopter.Gen

// SubscriptionAliasResponsePropertiesSTATUSARMGenerator returns a generator of SubscriptionAliasResponseProperties_STATUSARM instances for property testing.
func SubscriptionAliasResponsePropertiesSTATUSARMGenerator() gopter.Gen {
	if subscriptionAliasResponsePropertiesSTATUSARMGenerator != nil {
		return subscriptionAliasResponsePropertiesSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubscriptionAliasResponsePropertiesSTATUSARM(generators)
	subscriptionAliasResponsePropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(SubscriptionAliasResponseProperties_STATUSARM{}), generators)

	return subscriptionAliasResponsePropertiesSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSubscriptionAliasResponsePropertiesSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubscriptionAliasResponsePropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["AcceptOwnershipState"] = gen.PtrOf(gen.OneConstOf(AcceptOwnershipState_STATUS_Completed, AcceptOwnershipState_STATUS_Expired, AcceptOwnershipState_STATUS_Pending))
	gens["AcceptOwnershipUrl"] = gen.PtrOf(gen.AlphaString())
	gens["BillingScope"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedTime"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["ManagementGroupId"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(SubscriptionAliasResponsePropertiesSTATUSProvisioningState_Accepted, SubscriptionAliasResponsePropertiesSTATUSProvisioningState_Failed, SubscriptionAliasResponsePropertiesSTATUSProvisioningState_Succeeded))
	gens["ResellerId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionOwnerId"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Workload"] = gen.PtrOf(gen.OneConstOf(Workload_STATUS_DevTest, Workload_STATUS_Production))
}

func Test_SystemData_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemDataSTATUSARM, SystemDataSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemDataSTATUSARM runs a test to see if a specific instance of SystemData_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemDataSTATUSARM(subject SystemData_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_STATUSARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SystemData_STATUSARM instances for property testing - lazily instantiated by
// SystemDataSTATUSARMGenerator()
var systemDataSTATUSARMGenerator gopter.Gen

// SystemDataSTATUSARMGenerator returns a generator of SystemData_STATUSARM instances for property testing.
func SystemDataSTATUSARMGenerator() gopter.Gen {
	if systemDataSTATUSARMGenerator != nil {
		return systemDataSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemDataSTATUSARM(generators)
	systemDataSTATUSARMGenerator = gen.Struct(reflect.TypeOf(SystemData_STATUSARM{}), generators)

	return systemDataSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSystemDataSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemDataSTATUSARM(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemDataSTATUSCreatedByType_Application,
		SystemDataSTATUSCreatedByType_Key,
		SystemDataSTATUSCreatedByType_ManagedIdentity,
		SystemDataSTATUSCreatedByType_User))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemDataSTATUSLastModifiedByType_Application,
		SystemDataSTATUSLastModifiedByType_Key,
		SystemDataSTATUSLastModifiedByType_ManagedIdentity,
		SystemDataSTATUSLastModifiedByType_User))
}