// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"fmt"
	"github.com/oracle/oci-go-sdk/v59/common"
	"strings"
)

// SummarizeOperationsInsightsWarehouseResourceUsageAggregation Details of resource usage by an Operations Insights Warehouse resource.
type SummarizeOperationsInsightsWarehouseResourceUsageAggregation struct {

	// OPSI Warehouse OCID
	Id *string `mandatory:"true" json:"id"`

	// Possible lifecycle states
	LifecycleState OperationsInsightsWarehouseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Number of OCPUs used by OPSI Warehouse ADW. Can be fractional.
	CpuUsed *float64 `mandatory:"false" json:"cpuUsed"`

	// Storage by OPSI Warehouse ADW in GB.
	StorageUsedInGBs *float64 `mandatory:"false" json:"storageUsedInGBs"`
}

func (m SummarizeOperationsInsightsWarehouseResourceUsageAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeOperationsInsightsWarehouseResourceUsageAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOperationsInsightsWarehouseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOperationsInsightsWarehouseLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
