// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AttachLoadBalancerDetails Represents a load balancer that is to be attached to an instance pool.
type AttachLoadBalancerDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer to attach to the instance pool.
	LoadBalancerId *string `mandatory:"true" json:"loadBalancerId"`

	// The name of the backend set on the load balancer to add instances to.
	BackendSetName *string `mandatory:"true" json:"backendSetName"`

	// The port value to use when creating the backend set.
	Port *int `mandatory:"true" json:"port"`

	// Indicates which VNIC on each instance in the pool should be used to associate with the load balancer.
	// Possible values are "PrimaryVnic" or the displayName of one of the secondary VNICs on the instance configuration
	// that is associated with the instance pool.
	VnicSelection *string `mandatory:"true" json:"vnicSelection"`
}

func (m AttachLoadBalancerDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AttachLoadBalancerDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
