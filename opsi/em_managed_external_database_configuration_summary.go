// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v49/common"
)

// EmManagedExternalDatabaseConfigurationSummary Configuration summary of a EM Managed External database.
type EmManagedExternalDatabaseConfigurationSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database insight resource.
	DatabaseInsightId *string `mandatory:"true" json:"databaseInsightId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The database name. The database name is unique within the tenancy.
	DatabaseName *string `mandatory:"true" json:"databaseName"`

	// The user-friendly name for the database. The name does not have to be unique.
	DatabaseDisplayName *string `mandatory:"true" json:"databaseDisplayName"`

	// Operations Insights internal representation of the database type.
	DatabaseType *string `mandatory:"true" json:"databaseType"`

	// The version of the database.
	DatabaseVersion *string `mandatory:"true" json:"databaseVersion"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Enterprise Manager Unique Identifier
	EnterpriseManagerIdentifier *string `mandatory:"true" json:"enterpriseManagerIdentifier"`

	// OPSI Enterprise Manager Bridge OCID
	EnterpriseManagerBridgeId *string `mandatory:"true" json:"enterpriseManagerBridgeId"`

	// Array of hostname and instance name.
	Instances []HostInstanceMap `mandatory:"true" json:"instances"`

	// Processor count. This is the OCPU count for Autonomous Database and CPU core count for other database types.
	ProcessorCount *int `mandatory:"false" json:"processorCount"`
}

//GetDatabaseInsightId returns DatabaseInsightId
func (m EmManagedExternalDatabaseConfigurationSummary) GetDatabaseInsightId() *string {
	return m.DatabaseInsightId
}

//GetCompartmentId returns CompartmentId
func (m EmManagedExternalDatabaseConfigurationSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDatabaseName returns DatabaseName
func (m EmManagedExternalDatabaseConfigurationSummary) GetDatabaseName() *string {
	return m.DatabaseName
}

//GetDatabaseDisplayName returns DatabaseDisplayName
func (m EmManagedExternalDatabaseConfigurationSummary) GetDatabaseDisplayName() *string {
	return m.DatabaseDisplayName
}

//GetDatabaseType returns DatabaseType
func (m EmManagedExternalDatabaseConfigurationSummary) GetDatabaseType() *string {
	return m.DatabaseType
}

//GetDatabaseVersion returns DatabaseVersion
func (m EmManagedExternalDatabaseConfigurationSummary) GetDatabaseVersion() *string {
	return m.DatabaseVersion
}

//GetDefinedTags returns DefinedTags
func (m EmManagedExternalDatabaseConfigurationSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetFreeformTags returns FreeformTags
func (m EmManagedExternalDatabaseConfigurationSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetProcessorCount returns ProcessorCount
func (m EmManagedExternalDatabaseConfigurationSummary) GetProcessorCount() *int {
	return m.ProcessorCount
}

func (m EmManagedExternalDatabaseConfigurationSummary) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m EmManagedExternalDatabaseConfigurationSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeEmManagedExternalDatabaseConfigurationSummary EmManagedExternalDatabaseConfigurationSummary
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeEmManagedExternalDatabaseConfigurationSummary
	}{
		"EM_MANAGED_EXTERNAL_DATABASE",
		(MarshalTypeEmManagedExternalDatabaseConfigurationSummary)(m),
	}

	return json.Marshal(&s)
}
