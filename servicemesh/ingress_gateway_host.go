// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IngressGatewayHost Host for the ingress listener.
type IngressGatewayHost struct {

	// A user-friendly name for the host. The name must be unique within the same ingress gateway.
	// This name can be used in the ingress gateway route table resource to attach a route to this host.
	// Example: `MyExampleHost`
	Name *string `mandatory:"true" json:"name"`

	// The listeners for the ingress gateway.
	Listeners []IngressGatewayListener `mandatory:"true" json:"listeners"`

	// Hostnames of the host. Applicable only for HTTP and TLS_PASSTHROUGH listeners.
	// Wildcard hostnames are supported in the prefix form.
	// Examples of valid hostnames are "www.example.com", "*.example.com", "*.com".
	Hostnames []string `mandatory:"false" json:"hostnames"`
}

func (m IngressGatewayHost) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IngressGatewayHost) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
