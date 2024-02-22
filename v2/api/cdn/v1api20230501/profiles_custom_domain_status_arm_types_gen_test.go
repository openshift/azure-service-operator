// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501

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

func Test_Profiles_CustomDomain_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Profiles_CustomDomain_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfiles_CustomDomain_STATUS_ARM, Profiles_CustomDomain_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfiles_CustomDomain_STATUS_ARM runs a test to see if a specific instance of Profiles_CustomDomain_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForProfiles_CustomDomain_STATUS_ARM(subject Profiles_CustomDomain_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Profiles_CustomDomain_STATUS_ARM
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

// Generator of Profiles_CustomDomain_STATUS_ARM instances for property testing - lazily instantiated by
// Profiles_CustomDomain_STATUS_ARMGenerator()
var profiles_CustomDomain_STATUS_ARMGenerator gopter.Gen

// Profiles_CustomDomain_STATUS_ARMGenerator returns a generator of Profiles_CustomDomain_STATUS_ARM instances for property testing.
// We first initialize profiles_CustomDomain_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Profiles_CustomDomain_STATUS_ARMGenerator() gopter.Gen {
	if profiles_CustomDomain_STATUS_ARMGenerator != nil {
		return profiles_CustomDomain_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_CustomDomain_STATUS_ARM(generators)
	profiles_CustomDomain_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Profiles_CustomDomain_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_CustomDomain_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForProfiles_CustomDomain_STATUS_ARM(generators)
	profiles_CustomDomain_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Profiles_CustomDomain_STATUS_ARM{}), generators)

	return profiles_CustomDomain_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForProfiles_CustomDomain_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfiles_CustomDomain_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForProfiles_CustomDomain_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProfiles_CustomDomain_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(AFDDomainProperties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}

func Test_AFDDomainProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AFDDomainProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAFDDomainProperties_STATUS_ARM, AFDDomainProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAFDDomainProperties_STATUS_ARM runs a test to see if a specific instance of AFDDomainProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAFDDomainProperties_STATUS_ARM(subject AFDDomainProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AFDDomainProperties_STATUS_ARM
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

// Generator of AFDDomainProperties_STATUS_ARM instances for property testing - lazily instantiated by
// AFDDomainProperties_STATUS_ARMGenerator()
var afdDomainProperties_STATUS_ARMGenerator gopter.Gen

// AFDDomainProperties_STATUS_ARMGenerator returns a generator of AFDDomainProperties_STATUS_ARM instances for property testing.
// We first initialize afdDomainProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AFDDomainProperties_STATUS_ARMGenerator() gopter.Gen {
	if afdDomainProperties_STATUS_ARMGenerator != nil {
		return afdDomainProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAFDDomainProperties_STATUS_ARM(generators)
	afdDomainProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AFDDomainProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAFDDomainProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForAFDDomainProperties_STATUS_ARM(generators)
	afdDomainProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AFDDomainProperties_STATUS_ARM{}), generators)

	return afdDomainProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAFDDomainProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAFDDomainProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["DeploymentStatus"] = gen.PtrOf(gen.OneConstOf(
		AFDDomainProperties_DeploymentStatus_STATUS_Failed,
		AFDDomainProperties_DeploymentStatus_STATUS_InProgress,
		AFDDomainProperties_DeploymentStatus_STATUS_NotStarted,
		AFDDomainProperties_DeploymentStatus_STATUS_Succeeded))
	gens["DomainValidationState"] = gen.PtrOf(gen.OneConstOf(
		AFDDomainProperties_DomainValidationState_STATUS_Approved,
		AFDDomainProperties_DomainValidationState_STATUS_InternalError,
		AFDDomainProperties_DomainValidationState_STATUS_Pending,
		AFDDomainProperties_DomainValidationState_STATUS_PendingRevalidation,
		AFDDomainProperties_DomainValidationState_STATUS_RefreshingValidationToken,
		AFDDomainProperties_DomainValidationState_STATUS_Rejected,
		AFDDomainProperties_DomainValidationState_STATUS_Submitting,
		AFDDomainProperties_DomainValidationState_STATUS_TimedOut,
		AFDDomainProperties_DomainValidationState_STATUS_Unknown))
	gens["ExtendedProperties"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
	gens["ProfileName"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		AFDDomainProperties_ProvisioningState_STATUS_Creating,
		AFDDomainProperties_ProvisioningState_STATUS_Deleting,
		AFDDomainProperties_ProvisioningState_STATUS_Failed,
		AFDDomainProperties_ProvisioningState_STATUS_Succeeded,
		AFDDomainProperties_ProvisioningState_STATUS_Updating))
}

// AddRelatedPropertyGeneratorsForAFDDomainProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAFDDomainProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AzureDnsZone"] = gen.PtrOf(ResourceReference_STATUS_ARMGenerator())
	gens["PreValidatedCustomDomainResourceId"] = gen.PtrOf(ResourceReference_STATUS_ARMGenerator())
	gens["TlsSettings"] = gen.PtrOf(AFDDomainHttpsParameters_STATUS_ARMGenerator())
	gens["ValidationProperties"] = gen.PtrOf(DomainValidationProperties_STATUS_ARMGenerator())
}

func Test_AFDDomainHttpsParameters_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AFDDomainHttpsParameters_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAFDDomainHttpsParameters_STATUS_ARM, AFDDomainHttpsParameters_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAFDDomainHttpsParameters_STATUS_ARM runs a test to see if a specific instance of AFDDomainHttpsParameters_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAFDDomainHttpsParameters_STATUS_ARM(subject AFDDomainHttpsParameters_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AFDDomainHttpsParameters_STATUS_ARM
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

// Generator of AFDDomainHttpsParameters_STATUS_ARM instances for property testing - lazily instantiated by
// AFDDomainHttpsParameters_STATUS_ARMGenerator()
var afdDomainHttpsParameters_STATUS_ARMGenerator gopter.Gen

// AFDDomainHttpsParameters_STATUS_ARMGenerator returns a generator of AFDDomainHttpsParameters_STATUS_ARM instances for property testing.
// We first initialize afdDomainHttpsParameters_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AFDDomainHttpsParameters_STATUS_ARMGenerator() gopter.Gen {
	if afdDomainHttpsParameters_STATUS_ARMGenerator != nil {
		return afdDomainHttpsParameters_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAFDDomainHttpsParameters_STATUS_ARM(generators)
	afdDomainHttpsParameters_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AFDDomainHttpsParameters_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAFDDomainHttpsParameters_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForAFDDomainHttpsParameters_STATUS_ARM(generators)
	afdDomainHttpsParameters_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AFDDomainHttpsParameters_STATUS_ARM{}), generators)

	return afdDomainHttpsParameters_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAFDDomainHttpsParameters_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAFDDomainHttpsParameters_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CertificateType"] = gen.PtrOf(gen.OneConstOf(AFDDomainHttpsParameters_CertificateType_STATUS_AzureFirstPartyManagedCertificate, AFDDomainHttpsParameters_CertificateType_STATUS_CustomerCertificate, AFDDomainHttpsParameters_CertificateType_STATUS_ManagedCertificate))
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.OneConstOf(AFDDomainHttpsParameters_MinimumTlsVersion_STATUS_TLS10, AFDDomainHttpsParameters_MinimumTlsVersion_STATUS_TLS12))
}

// AddRelatedPropertyGeneratorsForAFDDomainHttpsParameters_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAFDDomainHttpsParameters_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Secret"] = gen.PtrOf(ResourceReference_STATUS_ARMGenerator())
}

func Test_DomainValidationProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DomainValidationProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainValidationProperties_STATUS_ARM, DomainValidationProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainValidationProperties_STATUS_ARM runs a test to see if a specific instance of DomainValidationProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainValidationProperties_STATUS_ARM(subject DomainValidationProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DomainValidationProperties_STATUS_ARM
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

// Generator of DomainValidationProperties_STATUS_ARM instances for property testing - lazily instantiated by
// DomainValidationProperties_STATUS_ARMGenerator()
var domainValidationProperties_STATUS_ARMGenerator gopter.Gen

// DomainValidationProperties_STATUS_ARMGenerator returns a generator of DomainValidationProperties_STATUS_ARM instances for property testing.
func DomainValidationProperties_STATUS_ARMGenerator() gopter.Gen {
	if domainValidationProperties_STATUS_ARMGenerator != nil {
		return domainValidationProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainValidationProperties_STATUS_ARM(generators)
	domainValidationProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DomainValidationProperties_STATUS_ARM{}), generators)

	return domainValidationProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDomainValidationProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomainValidationProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ExpirationDate"] = gen.PtrOf(gen.AlphaString())
	gens["ValidationToken"] = gen.PtrOf(gen.AlphaString())
}