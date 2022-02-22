// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Budgets API
//
// Use the Budgets API to manage budgets and budget alerts.
//

package budget

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v59/common"
	"strings"
)

// UpdateBudgetDetails The update budget details.
type UpdateBudgetDetails struct {

	// The displayName of the budget.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the budget.
	Description *string `mandatory:"false" json:"description"`

	// The amount of the budget expressed as a whole number in the currency of the customer's rate card.
	Amount *float32 `mandatory:"false" json:"amount"`

	// The number of days offset from the first day of the month, at which the budget processing period starts. In months that have fewer days than this value, processing will begin on the last day of that month. For example, for a value of 12, processing starts every month on the 12th at midnight.
	BudgetProcessingPeriodStartOffset *int `mandatory:"false" json:"budgetProcessingPeriodStartOffset"`

	// The reset period for the budget.
	ResetPeriod ResetPeriodEnum `mandatory:"false" json:"resetPeriod,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateBudgetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateBudgetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingResetPeriodEnum(string(m.ResetPeriod)); !ok && m.ResetPeriod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResetPeriod: %s. Supported values are: %s.", m.ResetPeriod, strings.Join(GetResetPeriodEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateBudgetDetailsResetPeriodEnum is an alias to type: ResetPeriodEnum
// Consider using ResetPeriodEnum instead
// Deprecated
type UpdateBudgetDetailsResetPeriodEnum = ResetPeriodEnum

// Set of constants representing the allowable values for ResetPeriodEnum
// Deprecated
const (
	UpdateBudgetDetailsResetPeriodMonthly ResetPeriodEnum = "MONTHLY"
)

// GetUpdateBudgetDetailsResetPeriodEnumValues Enumerates the set of values for ResetPeriodEnum
// Consider using GetResetPeriodEnumValue
// Deprecated
var GetUpdateBudgetDetailsResetPeriodEnumValues = GetResetPeriodEnumValues
