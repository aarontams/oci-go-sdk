// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Speech API
//
// The OCI Speech Service harnesses the power of spoken language by allowing developers to easily convert file-based data containing human speech into highly accurate text transcriptions.
//

package aispeech

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TranscriptionJobSummary Summary of the Transcription Job.
type TranscriptionJobSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly display name for the job.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the job.
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// How much progress the operation has made, vs the total amount of work that must be performed.
	PercentComplete *int `mandatory:"false" json:"percentComplete"`

	// Job accepted time.
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// Job started time.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Job finished time.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// Total number of tasks in a job.
	TotalTasks *int `mandatory:"false" json:"totalTasks"`

	// Total outstanding tasks in a job.
	OutstandingTasks *int `mandatory:"false" json:"outstandingTasks"`

	// Total successful tasks in a job.
	SuccessfulTasks *int `mandatory:"false" json:"successfulTasks"`

	// The current state of the Speech Job.
	LifecycleState TranscriptionJobLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace-1": {"bar-key-1": "value-1", "bar-key-2": "value-2"}, "foo-namespace-2": {"bar-key-1": "value-1", "bar-key-2": "value-2"}}`.
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`.
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m TranscriptionJobSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TranscriptionJobSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTranscriptionJobLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTranscriptionJobLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
