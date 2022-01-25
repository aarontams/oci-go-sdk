// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osubsubscription

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListRateCardsRequest wrapper for the ListRateCards operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osubsubscription/ListRateCards.go.html to see an example of how to use ListRateCardsRequest.
type ListRateCardsRequest struct {

	// Line level Subscription Id
	SubscriptionId *string `mandatory:"true" contributesTo:"query" name:"subscriptionId"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// This param is used to get the rate card(s) whose effective start date starts on or after a particular date
	TimeFrom *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeFrom"`

	// This param is used to get the rate card(s) whose effective end date ends on or before a particular date
	TimeTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeTo"`

	// This param is used to get the rate card(s) filterd by the partNumber
	PartNumber *string `mandatory:"false" contributesTo:"query" name:"partNumber"`

	// The maximum number of items to return in a paginated "List" call. Default: (`50`)
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListRateCardsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`).
	SortBy ListRateCardsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCI home region name in case home region is not us-ashburn-1 (IAD), e.g. ap-mumbai-1, us-phoenix-1 etc.
	XOneOriginRegion *string `mandatory:"false" contributesTo:"header" name:"x-one-origin-region"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRateCardsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRateCardsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRateCardsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRateCardsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListRateCardsResponse wrapper for the ListRateCards operation
type ListRateCardsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []RateCardSummary instances
	Items []RateCardSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListRateCardsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRateCardsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRateCardsSortOrderEnum Enum with underlying type: string
type ListRateCardsSortOrderEnum string

// Set of constants representing the allowable values for ListRateCardsSortOrderEnum
const (
	ListRateCardsSortOrderAsc  ListRateCardsSortOrderEnum = "ASC"
	ListRateCardsSortOrderDesc ListRateCardsSortOrderEnum = "DESC"
)

var mappingListRateCardsSortOrder = map[string]ListRateCardsSortOrderEnum{
	"ASC":  ListRateCardsSortOrderAsc,
	"DESC": ListRateCardsSortOrderDesc,
}

// GetListRateCardsSortOrderEnumValues Enumerates the set of values for ListRateCardsSortOrderEnum
func GetListRateCardsSortOrderEnumValues() []ListRateCardsSortOrderEnum {
	values := make([]ListRateCardsSortOrderEnum, 0)
	for _, v := range mappingListRateCardsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListRateCardsSortByEnum Enum with underlying type: string
type ListRateCardsSortByEnum string

// Set of constants representing the allowable values for ListRateCardsSortByEnum
const (
	ListRateCardsSortByTimecreated ListRateCardsSortByEnum = "TIMECREATED"
	ListRateCardsSortByTimestart   ListRateCardsSortByEnum = "TIMESTART"
)

var mappingListRateCardsSortBy = map[string]ListRateCardsSortByEnum{
	"TIMECREATED": ListRateCardsSortByTimecreated,
	"TIMESTART":   ListRateCardsSortByTimestart,
}

// GetListRateCardsSortByEnumValues Enumerates the set of values for ListRateCardsSortByEnum
func GetListRateCardsSortByEnumValues() []ListRateCardsSortByEnum {
	values := make([]ListRateCardsSortByEnum, 0)
	for _, v := range mappingListRateCardsSortBy {
		values = append(values, v)
	}
	return values
}
