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

// UpdatePrivateIpDetails The representation of UpdatePrivateIpDetails
type UpdatePrivateIpDetails struct {

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The hostname for the private IP. Used for DNS. The value
	// is the hostname portion of the private IP's fully qualified domain name (FQDN)
	// (for example, `bminstance1` in FQDN `bminstance1.subnet123.vcn1.oraclevcn.com`).
	// Must be unique across all VNICs in the subnet and comply with
	// RFC 952 (https://tools.ietf.org/html/rfc952) and
	// RFC 1123 (https://tools.ietf.org/html/rfc1123).
	// For more information, see
	// DNS in Your Virtual Cloud Network (https://docs.oracle.com/iaas/Content/Network/Concepts/dns.htm).
	// Example: `bminstance1`
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC to reassign the private IP to. The VNIC must
	// be in the same subnet as the current VNIC.
	VnicId *string `mandatory:"false" json:"vnicId"`

	// Lifetime of the IP address.
	// There are two types of IPv6 IPs:
	//  - Ephemeral
	//  - Reserved
	Lifetime UpdatePrivateIpDetailsLifetimeEnum `mandatory:"false" json:"lifetime,omitempty"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the IP address or VNIC will use. For more information, see
	// Source Based Routing (https://docs.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm#Overview_of_Routing_for_Your_VCN__source_routing).
	RouteTableId *string `mandatory:"false" json:"routeTableId"`
}

func (m UpdatePrivateIpDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdatePrivateIpDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdatePrivateIpDetailsLifetimeEnum(string(m.Lifetime)); !ok && m.Lifetime != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Lifetime: %s. Supported values are: %s.", m.Lifetime, strings.Join(GetUpdatePrivateIpDetailsLifetimeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdatePrivateIpDetailsLifetimeEnum Enum with underlying type: string
type UpdatePrivateIpDetailsLifetimeEnum string

// Set of constants representing the allowable values for UpdatePrivateIpDetailsLifetimeEnum
const (
	UpdatePrivateIpDetailsLifetimeEphemeral UpdatePrivateIpDetailsLifetimeEnum = "EPHEMERAL"
	UpdatePrivateIpDetailsLifetimeReserved  UpdatePrivateIpDetailsLifetimeEnum = "RESERVED"
)

var mappingUpdatePrivateIpDetailsLifetimeEnum = map[string]UpdatePrivateIpDetailsLifetimeEnum{
	"EPHEMERAL": UpdatePrivateIpDetailsLifetimeEphemeral,
	"RESERVED":  UpdatePrivateIpDetailsLifetimeReserved,
}

var mappingUpdatePrivateIpDetailsLifetimeEnumLowerCase = map[string]UpdatePrivateIpDetailsLifetimeEnum{
	"ephemeral": UpdatePrivateIpDetailsLifetimeEphemeral,
	"reserved":  UpdatePrivateIpDetailsLifetimeReserved,
}

// GetUpdatePrivateIpDetailsLifetimeEnumValues Enumerates the set of values for UpdatePrivateIpDetailsLifetimeEnum
func GetUpdatePrivateIpDetailsLifetimeEnumValues() []UpdatePrivateIpDetailsLifetimeEnum {
	values := make([]UpdatePrivateIpDetailsLifetimeEnum, 0)
	for _, v := range mappingUpdatePrivateIpDetailsLifetimeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdatePrivateIpDetailsLifetimeEnumStringValues Enumerates the set of values in String for UpdatePrivateIpDetailsLifetimeEnum
func GetUpdatePrivateIpDetailsLifetimeEnumStringValues() []string {
	return []string{
		"EPHEMERAL",
		"RESERVED",
	}
}

// GetMappingUpdatePrivateIpDetailsLifetimeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdatePrivateIpDetailsLifetimeEnum(val string) (UpdatePrivateIpDetailsLifetimeEnum, bool) {
	enum, ok := mappingUpdatePrivateIpDetailsLifetimeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
