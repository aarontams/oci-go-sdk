// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MaskingPolicyHealthReport The masking policy health report.
type MaskingPolicyHealthReport struct {

	// The OCID of the health report.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the masking policy.
	MaskingPolicyId *string `mandatory:"true" json:"maskingPolicyId"`

	// The OCID of the target database for which this report was created.
	TargetId *string `mandatory:"true" json:"targetId"`

	// The OCID of the compartment that contains the health report.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the health report.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time the report was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the health report.
	LifecycleState MaskingPolicyHealthReportLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The date and time the report was last updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339)
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The description of the masking health report.
	Description *string `mandatory:"false" json:"description"`

	// The count of errors in the masking health report.
	ErrorCount *int64 `mandatory:"false" json:"errorCount"`

	// The count of warnings in the masking health report.
	WarningCount *int64 `mandatory:"false" json:"warningCount"`
}

func (m MaskingPolicyHealthReport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaskingPolicyHealthReport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMaskingPolicyHealthReportLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMaskingPolicyHealthReportLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MaskingPolicyHealthReportLifecycleStateEnum Enum with underlying type: string
type MaskingPolicyHealthReportLifecycleStateEnum string

// Set of constants representing the allowable values for MaskingPolicyHealthReportLifecycleStateEnum
const (
	MaskingPolicyHealthReportLifecycleStateCreating MaskingPolicyHealthReportLifecycleStateEnum = "CREATING"
	MaskingPolicyHealthReportLifecycleStateActive   MaskingPolicyHealthReportLifecycleStateEnum = "ACTIVE"
	MaskingPolicyHealthReportLifecycleStateUpdating MaskingPolicyHealthReportLifecycleStateEnum = "UPDATING"
	MaskingPolicyHealthReportLifecycleStateDeleting MaskingPolicyHealthReportLifecycleStateEnum = "DELETING"
	MaskingPolicyHealthReportLifecycleStateFailed   MaskingPolicyHealthReportLifecycleStateEnum = "FAILED"
)

var mappingMaskingPolicyHealthReportLifecycleStateEnum = map[string]MaskingPolicyHealthReportLifecycleStateEnum{
	"CREATING": MaskingPolicyHealthReportLifecycleStateCreating,
	"ACTIVE":   MaskingPolicyHealthReportLifecycleStateActive,
	"UPDATING": MaskingPolicyHealthReportLifecycleStateUpdating,
	"DELETING": MaskingPolicyHealthReportLifecycleStateDeleting,
	"FAILED":   MaskingPolicyHealthReportLifecycleStateFailed,
}

var mappingMaskingPolicyHealthReportLifecycleStateEnumLowerCase = map[string]MaskingPolicyHealthReportLifecycleStateEnum{
	"creating": MaskingPolicyHealthReportLifecycleStateCreating,
	"active":   MaskingPolicyHealthReportLifecycleStateActive,
	"updating": MaskingPolicyHealthReportLifecycleStateUpdating,
	"deleting": MaskingPolicyHealthReportLifecycleStateDeleting,
	"failed":   MaskingPolicyHealthReportLifecycleStateFailed,
}

// GetMaskingPolicyHealthReportLifecycleStateEnumValues Enumerates the set of values for MaskingPolicyHealthReportLifecycleStateEnum
func GetMaskingPolicyHealthReportLifecycleStateEnumValues() []MaskingPolicyHealthReportLifecycleStateEnum {
	values := make([]MaskingPolicyHealthReportLifecycleStateEnum, 0)
	for _, v := range mappingMaskingPolicyHealthReportLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMaskingPolicyHealthReportLifecycleStateEnumStringValues Enumerates the set of values in String for MaskingPolicyHealthReportLifecycleStateEnum
func GetMaskingPolicyHealthReportLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"FAILED",
	}
}

// GetMappingMaskingPolicyHealthReportLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaskingPolicyHealthReportLifecycleStateEnum(val string) (MaskingPolicyHealthReportLifecycleStateEnum, bool) {
	enum, ok := mappingMaskingPolicyHealthReportLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
