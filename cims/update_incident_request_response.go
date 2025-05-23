// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cims

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// UpdateIncidentRequest wrapper for the UpdateIncident operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cims/UpdateIncident.go.html to see an example of how to use UpdateIncidentRequest.
type UpdateIncidentRequest struct {

	// Unique identifier for the support ticket.
	IncidentKey *string `mandatory:"true" contributesTo:"path" name:"incidentKey"`

	// Details about the support ticket being updated.
	UpdateIncidentDetails UpdateIncident `contributesTo:"body"`

	// The OCID of the tenancy.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The Customer Support Identifier (CSI) number associated with the support account.
	// The CSI is optional for all support request types.
	Csi *string `mandatory:"false" contributesTo:"header" name:"csi"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match` parameter to the value of the etag from a previous GET or POST response for that resource. The resource will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// User OCID for Oracle Identity Cloud Service (IDCS) users who also have a federated Oracle Cloud Infrastructure account.
	// User OCID is mandatory for OCI Users and optional for Multicloud users.
	Ocid *string `mandatory:"false" contributesTo:"header" name:"ocid"`

	// The region of the tenancy.
	Homeregion *string `mandatory:"false" contributesTo:"header" name:"homeregion"`

	// Token type that determine which cloud provider the request come from.
	Bearertokentype *string `mandatory:"false" contributesTo:"header" name:"bearertokentype"`

	// Token that provided by multi cloud provider, which help to validate the email.
	Bearertoken *string `mandatory:"false" contributesTo:"header" name:"bearertoken"`

	// IdToken that provided by multi cloud provider, which help to validate the email.
	Idtoken *string `mandatory:"false" contributesTo:"header" name:"idtoken"`

	// The OCID of identity domain.
	// DomainID is mandatory if the user is part of Non Default Identity domain.
	Domainid *string `mandatory:"false" contributesTo:"header" name:"domainid"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateIncidentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateIncidentRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request UpdateIncidentRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateIncidentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request UpdateIncidentRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateIncidentResponse wrapper for the UpdateIncident operation
type UpdateIncidentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Incident instance
	Incident `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response UpdateIncidentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateIncidentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
