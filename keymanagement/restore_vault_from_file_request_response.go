// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package keymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"io"
	"net/http"
	"strings"
)

// RestoreVaultFromFileRequest wrapper for the RestoreVaultFromFile operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/keymanagement/RestoreVaultFromFile.go.html to see an example of how to use RestoreVaultFromFileRequest.
type RestoreVaultFromFileRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The content length of the body.
	ContentLength *int64 `mandatory:"false" contributesTo:"header" name:"content-length"`

	// The encrypted backup file to upload to restore the vault.
	RestoreVaultFromFileDetails io.ReadCloser `mandatory:"true" contributesTo:"body" encoding:"binary"`

	// For optimistic concurrency control. In the PUT or DELETE call for a
	// resource, set the `if-match` parameter to the value of the etag from a
	// previous GET or POST response for that resource. The resource will be
	// updated or deleted only if the etag you provide matches the resource's
	// current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The base64-encoded MD5 hash value of the body, as described in RFC 2616 (https://tools.ietf.org/rfc/rfc2616), section 14.15.
	// If the Content-MD5 header is present, Key Management performs an integrity check on the body of the HTTP request by computing the MD5
	// hash for the body and comparing it to the MD5 hash supplied in the header. If the two hashes don't match, the object is rejected and
	// a response with 400 Unmatched Content MD5 error is returned, along with the message: "The computed MD5 of the request body (ACTUAL_MD5)
	// does not match the Content-MD5 header (HEADER_MD5)."
	ContentMd5 *string `mandatory:"false" contributesTo:"header" name:"content-md5"`

	// Unique identifier for the request. If provided, the returned request ID
	// will include this value. Otherwise, a random request ID will be
	// generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case
	// of a timeout or server error without risk of executing that same action
	// again. Retry tokens expire after 24 hours, but can be invalidated
	// before then due to conflicting operations (e.g., if a resource has been
	// deleted and purged from the system, then a retry of the original
	// creation request may be rejected).
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request RestoreVaultFromFileRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request RestoreVaultFromFileRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
	if err == nil && binaryRequestBody.Seekable() {
		common.UpdateRequestBinaryBody(&httpRequest, binaryRequestBody)
	}
	return httpRequest, err
}

// BinaryRequestBody implements the OCIRequest interface
func (request RestoreVaultFromFileRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {
	rsc := common.NewOCIReadSeekCloser(request.RestoreVaultFromFileDetails)
	if rsc.Seekable() {
		return rsc, true
	}
	return nil, true

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request RestoreVaultFromFileRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request RestoreVaultFromFileRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RestoreVaultFromFileResponse wrapper for the RestoreVaultFromFile operation
type RestoreVaultFromFileResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Vault instance
	Vault `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// The base64-encoded MD5 hash value of the request body, as computed
	// by the server.
	OpcContentMd5 *string `presentIn:"header" name:"opc-content-md5"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Unique Oracle-assigned identifier for the work request, used to track the progress of the
	// restore operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`
}

func (response RestoreVaultFromFileResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response RestoreVaultFromFileResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
