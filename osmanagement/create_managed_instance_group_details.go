// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/v49/common"
)

// CreateManagedInstanceGroupDetails Detail information for creating a managed instance group
type CreateManagedInstanceGroupDetails struct {

	// Managed Instance Group identifier
	DisplayName *string `mandatory:"true" json:"displayName"`

	// OCID for the Compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Information specified by the user about the managed instance group
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The Operating System type of the managed instance(s) on which this scheduled job will operate.
	// If not specified, this defaults to Linux.
	OsFamily OsFamiliesEnum `mandatory:"false" json:"osFamily,omitempty"`
}

func (m CreateManagedInstanceGroupDetails) String() string {
	return common.PointerString(m)
}
