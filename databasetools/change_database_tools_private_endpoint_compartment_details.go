// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Database Tools APIs to manage Connections and Private Endpoints.
//

package databasetools

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v59/common"
	"strings"
)

// ChangeDatabaseToolsPrivateEndpointCompartmentDetails Contains the details for the compartment to move the DatabaseToolsPrivateEndpoint to.
type ChangeDatabaseToolsPrivateEndpointCompartmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment to move the DatabaseConnectionProfile to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeDatabaseToolsPrivateEndpointCompartmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChangeDatabaseToolsPrivateEndpointCompartmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
