// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ServiceSummary Summary of services that are integrated with public logging.
type ServiceSummary struct {

	// Tenant OCID.
	TenantId *string `mandatory:"true" json:"tenantId"`

	// Service ID as set in Service Principal.
	ServicePrincipalName *string `mandatory:"true" json:"servicePrincipalName"`

	// Service endpoint.
	Endpoint *string `mandatory:"true" json:"endpoint"`

	// User-friendly service name.
	Name *string `mandatory:"true" json:"name"`

	// Type of resource that a service provides.
	ResourceTypes []ResourceType `mandatory:"true" json:"resourceTypes"`

	// Apollo project namespace, if any.
	Namespace *string `mandatory:"false" json:"namespace"`

	// Service ID.
	Id *string `mandatory:"false" json:"id"`
}

func (m ServiceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ServiceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
