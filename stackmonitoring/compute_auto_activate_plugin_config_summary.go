// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputeAutoActivatePluginConfigSummary Summary of a COMPUTE_AUTO_ACTIVATE_PLUGIN config.
type ComputeAutoActivatePluginConfigSummary struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// True if automatic activation of the Management Agent plugin is enabled, false if it is not enabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// Config Identifier, can be renamed.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time the the configuration was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the configuration was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The current state of the configuration.
	LifecycleState ConfigLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m ComputeAutoActivatePluginConfigSummary) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m ComputeAutoActivatePluginConfigSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m ComputeAutoActivatePluginConfigSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetTimeCreated returns TimeCreated
func (m ComputeAutoActivatePluginConfigSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m ComputeAutoActivatePluginConfigSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m ComputeAutoActivatePluginConfigSummary) GetLifecycleState() ConfigLifecycleStateEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m ComputeAutoActivatePluginConfigSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m ComputeAutoActivatePluginConfigSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m ComputeAutoActivatePluginConfigSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m ComputeAutoActivatePluginConfigSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeAutoActivatePluginConfigSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingConfigLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConfigLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ComputeAutoActivatePluginConfigSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeComputeAutoActivatePluginConfigSummary ComputeAutoActivatePluginConfigSummary
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeComputeAutoActivatePluginConfigSummary
	}{
		"COMPUTE_AUTO_ACTIVATE_PLUGIN",
		(MarshalTypeComputeAutoActivatePluginConfigSummary)(m),
	}

	return json.Marshal(&s)
}
