// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Work Requests API
//
// Many of the API operations that you use to create and configure Compute resources do not take effect
// immediately. In these cases, the operation spawns an asynchronous workflow to fulfill the request.
// Work requests provide visibility into the status of these in-progress, long-running workflows.
// For more information about work requests and the operations that spawn work requests, see
// Viewing the State of a Compute Work Request (https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/viewingworkrequestcompute.htm).
//

package workrequests

import (
	"github.com/oracle/oci-go-sdk/v26/common"
)

// WorkRequestError An error encountered while executing an operation that is tracked by a work request.
type WorkRequestError struct {

	// A machine-usable code for the error that occured.
	Code *string `mandatory:"true" json:"code"`

	// A human-readable error string.
	Message *string `mandatory:"true" json:"message"`

	// The date and time the error occurred.
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`
}

func (m WorkRequestError) String() string {
	return common.PointerString(m)
}
