// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

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

func Test_VirtualNetwork_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetwork_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkSTATUSARM, VirtualNetworkSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkSTATUSARM runs a test to see if a specific instance of VirtualNetwork_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkSTATUSARM(subject VirtualNetwork_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetwork_STATUSARM
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

// Generator of VirtualNetwork_STATUSARM instances for property testing - lazily instantiated by
// VirtualNetworkSTATUSARMGenerator()
var virtualNetworkSTATUSARMGenerator gopter.Gen

// VirtualNetworkSTATUSARMGenerator returns a generator of VirtualNetwork_STATUSARM instances for property testing.
// We first initialize virtualNetworkSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkSTATUSARMGenerator() gopter.Gen {
	if virtualNetworkSTATUSARMGenerator != nil {
		return virtualNetworkSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkSTATUSARM(generators)
	virtualNetworkSTATUSARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkSTATUSARM(generators)
	virtualNetworkSTATUSARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_STATUSARM{}), generators)

	return virtualNetworkSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkSTATUSARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkSTATUSARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationSTATUSARMGenerator())
	gens["Properties"] = gen.PtrOf(VirtualNetworkPropertiesFormatSTATUSARMGenerator())
}

func Test_VirtualNetworkPropertiesFormat_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkPropertiesFormat_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkPropertiesFormatSTATUSARM, VirtualNetworkPropertiesFormatSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkPropertiesFormatSTATUSARM runs a test to see if a specific instance of VirtualNetworkPropertiesFormat_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkPropertiesFormatSTATUSARM(subject VirtualNetworkPropertiesFormat_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkPropertiesFormat_STATUSARM
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

// Generator of VirtualNetworkPropertiesFormat_STATUSARM instances for property testing - lazily instantiated by
// VirtualNetworkPropertiesFormatSTATUSARMGenerator()
var virtualNetworkPropertiesFormatSTATUSARMGenerator gopter.Gen

// VirtualNetworkPropertiesFormatSTATUSARMGenerator returns a generator of VirtualNetworkPropertiesFormat_STATUSARM instances for property testing.
// We first initialize virtualNetworkPropertiesFormatSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkPropertiesFormatSTATUSARMGenerator() gopter.Gen {
	if virtualNetworkPropertiesFormatSTATUSARMGenerator != nil {
		return virtualNetworkPropertiesFormatSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPropertiesFormatSTATUSARM(generators)
	virtualNetworkPropertiesFormatSTATUSARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPropertiesFormat_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPropertiesFormatSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkPropertiesFormatSTATUSARM(generators)
	virtualNetworkPropertiesFormatSTATUSARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPropertiesFormat_STATUSARM{}), generators)

	return virtualNetworkPropertiesFormatSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkPropertiesFormatSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkPropertiesFormatSTATUSARM(gens map[string]gopter.Gen) {
	gens["EnableDdosProtection"] = gen.PtrOf(gen.Bool())
	gens["EnableVmProtection"] = gen.PtrOf(gen.Bool())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkPropertiesFormatSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkPropertiesFormatSTATUSARM(gens map[string]gopter.Gen) {
	gens["AddressSpace"] = gen.PtrOf(AddressSpaceSTATUSARMGenerator())
	gens["BgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunitiesSTATUSARMGenerator())
	gens["DdosProtectionPlan"] = gen.PtrOf(SubResourceSTATUSARMGenerator())
	gens["DhcpOptions"] = gen.PtrOf(DhcpOptionsSTATUSARMGenerator())
	gens["IpAllocations"] = gen.SliceOf(SubResourceSTATUSARMGenerator())
}

func Test_DhcpOptions_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DhcpOptions_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDhcpOptionsSTATUSARM, DhcpOptionsSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDhcpOptionsSTATUSARM runs a test to see if a specific instance of DhcpOptions_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDhcpOptionsSTATUSARM(subject DhcpOptions_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DhcpOptions_STATUSARM
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

// Generator of DhcpOptions_STATUSARM instances for property testing - lazily instantiated by
// DhcpOptionsSTATUSARMGenerator()
var dhcpOptionsSTATUSARMGenerator gopter.Gen

// DhcpOptionsSTATUSARMGenerator returns a generator of DhcpOptions_STATUSARM instances for property testing.
func DhcpOptionsSTATUSARMGenerator() gopter.Gen {
	if dhcpOptionsSTATUSARMGenerator != nil {
		return dhcpOptionsSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDhcpOptionsSTATUSARM(generators)
	dhcpOptionsSTATUSARMGenerator = gen.Struct(reflect.TypeOf(DhcpOptions_STATUSARM{}), generators)

	return dhcpOptionsSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForDhcpOptionsSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDhcpOptionsSTATUSARM(gens map[string]gopter.Gen) {
	gens["DnsServers"] = gen.SliceOf(gen.AlphaString())
}