// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v26/common"
	"net/http"
)

// ExtractStructuredLogHeaderPathsRequest wrapper for the ExtractStructuredLogHeaderPaths operation
type ExtractStructuredLogHeaderPathsRequest struct {

	// The Log Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// parser definition
	LoganParserDetails LogAnalyticsParser `contributesTo:"body"`

	// type - possible values are xml or json
	ParserType ExtractStructuredLogHeaderPathsParserTypeEnum `mandatory:"false" contributesTo:"query" name:"parserType" omitEmpty:"true"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ExtractStructuredLogHeaderPathsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ExtractStructuredLogHeaderPathsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ExtractStructuredLogHeaderPathsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ExtractStructuredLogHeaderPathsResponse wrapper for the ExtractStructuredLogHeaderPaths operation
type ExtractStructuredLogHeaderPathsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ExtractLogHeaderResults instance
	ExtractLogHeaderResults `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ExtractStructuredLogHeaderPathsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ExtractStructuredLogHeaderPathsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ExtractStructuredLogHeaderPathsParserTypeEnum Enum with underlying type: string
type ExtractStructuredLogHeaderPathsParserTypeEnum string

// Set of constants representing the allowable values for ExtractStructuredLogHeaderPathsParserTypeEnum
const (
	ExtractStructuredLogHeaderPathsParserTypeXml  ExtractStructuredLogHeaderPathsParserTypeEnum = "XML"
	ExtractStructuredLogHeaderPathsParserTypeJson ExtractStructuredLogHeaderPathsParserTypeEnum = "JSON"
)

var mappingExtractStructuredLogHeaderPathsParserType = map[string]ExtractStructuredLogHeaderPathsParserTypeEnum{
	"XML":  ExtractStructuredLogHeaderPathsParserTypeXml,
	"JSON": ExtractStructuredLogHeaderPathsParserTypeJson,
}

// GetExtractStructuredLogHeaderPathsParserTypeEnumValues Enumerates the set of values for ExtractStructuredLogHeaderPathsParserTypeEnum
func GetExtractStructuredLogHeaderPathsParserTypeEnumValues() []ExtractStructuredLogHeaderPathsParserTypeEnum {
	values := make([]ExtractStructuredLogHeaderPathsParserTypeEnum, 0)
	for _, v := range mappingExtractStructuredLogHeaderPathsParserType {
		values = append(values, v)
	}
	return values
}
