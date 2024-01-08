// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateLifecycleProfileDetails Description of a lifecycle registration profile to be created.
type CreateLifecycleProfileDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the tenancy containing the registration profile.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the lifecycle stage from which the registration profile will inherit its software source.
	LifecycleStageId *string `mandatory:"true" json:"lifecycleStageId"`

	// The description of the registration profile.
	Description *string `mandatory:"false" json:"description"`

	// The OCID of the management station.
	ManagementStationId *string `mandatory:"false" json:"managementStationId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

// GetDisplayName returns DisplayName
func (m CreateLifecycleProfileDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m CreateLifecycleProfileDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDescription returns Description
func (m CreateLifecycleProfileDetails) GetDescription() *string {
	return m.Description
}

// GetManagementStationId returns ManagementStationId
func (m CreateLifecycleProfileDetails) GetManagementStationId() *string {
	return m.ManagementStationId
}

// GetFreeformTags returns FreeformTags
func (m CreateLifecycleProfileDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateLifecycleProfileDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateLifecycleProfileDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateLifecycleProfileDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateLifecycleProfileDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateLifecycleProfileDetails CreateLifecycleProfileDetails
	s := struct {
		DiscriminatorParam string `json:"profileType"`
		MarshalTypeCreateLifecycleProfileDetails
	}{
		"LIFECYCLE",
		(MarshalTypeCreateLifecycleProfileDetails)(m),
	}

	return json.Marshal(&s)
}
