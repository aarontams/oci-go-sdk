// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GovernanceRulesControlPlane API
//
// A description of the GovernanceRulesControlPlane API
//

package governancerulescontrolplane

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateGovernanceRuleDetails Request object for CreateGovernanceRule operation.
type CreateGovernanceRuleDetails struct {

	// The Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the root compartment containing the governance rule.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Display name of the governance rule.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Type of the governance rule, can be one of QUOTA, TAG, ALLOWED_REGIONS.
	// Example: `QUOTA`
	Type GovernanceRuleTypeEnum `mandatory:"true" json:"type"`

	// The type of option used to create the governance rule, could be one of TEMPLATE or CLONE.
	// Example: `TEMPLATE`
	CreationOption CreationOptionEnum `mandatory:"true" json:"creationOption"`

	Template Template `mandatory:"true" json:"template"`

	// Description of the governance rule.
	Description *string `mandatory:"false" json:"description"`

	// The Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the resource, which was used as a template to create this governance rule.
	RelatedResourceId *string `mandatory:"false" json:"relatedResourceId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateGovernanceRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateGovernanceRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGovernanceRuleTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetGovernanceRuleTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreationOptionEnum(string(m.CreationOption)); !ok && m.CreationOption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CreationOption: %s. Supported values are: %s.", m.CreationOption, strings.Join(GetCreationOptionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateGovernanceRuleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description       *string                           `json:"description"`
		RelatedResourceId *string                           `json:"relatedResourceId"`
		FreeformTags      map[string]string                 `json:"freeformTags"`
		DefinedTags       map[string]map[string]interface{} `json:"definedTags"`
		CompartmentId     *string                           `json:"compartmentId"`
		DisplayName       *string                           `json:"displayName"`
		Type              GovernanceRuleTypeEnum            `json:"type"`
		CreationOption    CreationOptionEnum                `json:"creationOption"`
		Template          template                          `json:"template"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.RelatedResourceId = model.RelatedResourceId

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.Type = model.Type

	m.CreationOption = model.CreationOption

	nn, e = model.Template.UnmarshalPolymorphicJSON(model.Template.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Template = nn.(Template)
	} else {
		m.Template = nil
	}

	return
}
