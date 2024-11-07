//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

import "encoding/json"

func unmarshalBackupPolicyClassification(rawMsg json.RawMessage) (BackupPolicyClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b BackupPolicyClassification
	switch m["type"] {
	case string(BackupPolicyTypeContinuous):
		b = &ContinuousModeBackupPolicy{}
	case string(BackupPolicyTypePeriodic):
		b = &PeriodicModeBackupPolicy{}
	default:
		b = &BackupPolicy{}
	}
	return b, json.Unmarshal(rawMsg, b)
}
