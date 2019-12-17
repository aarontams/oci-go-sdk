// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// OracleFunctionBackend Send the request to an Oracle Functions function.
type OracleFunctionBackend struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Functions function resource.
	FunctionId *string `mandatory:"true" json:"functionId"`
}

func (m OracleFunctionBackend) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m OracleFunctionBackend) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOracleFunctionBackend OracleFunctionBackend
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeOracleFunctionBackend
	}{
		"ORACLE_FUNCTIONS_BACKEND",
		(MarshalTypeOracleFunctionBackend)(m),
	}

	return json.Marshal(&s)
}
