// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osubusage

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListComputedUsageAggregatedsRequest wrapper for the ListComputedUsageAggregateds operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osubusage/ListComputedUsageAggregateds.go.html to see an example of how to use ListComputedUsageAggregatedsRequest.
type ListComputedUsageAggregatedsRequest struct {

	// The OCID of the root compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Subscription Id is an identifier associated to the service used for filter the Computed Usage in SPM.
	SubscriptionId *string `mandatory:"true" contributesTo:"query" name:"subscriptionId"`

	// Initial date to filter Computed Usage data in SPM. In the case of non aggregated data the time period between of fromDate and toDate , expressed in RFC 3339 timestamp format.
	TimeFrom *common.SDKTime `mandatory:"true" contributesTo:"query" name:"timeFrom"`

	// Final date to filter Computed Usage data in SPM, expressed in RFC 3339 timestamp format.
	TimeTo *common.SDKTime `mandatory:"true" contributesTo:"query" name:"timeTo"`

	// Product part number for subscribed service line, called parent product.
	ParentProduct *string `mandatory:"false" contributesTo:"query" name:"parentProduct"`

	// Grouping criteria to use for aggregate the computed Usage, either hourly (`HOURLY`), daily (`DAILY`), monthly(`MONTHLY`) or none (`NONE`) to not follow a grouping criteria by date.
	Grouping ListComputedUsageAggregatedsGroupingEnum `mandatory:"false" contributesTo:"query" name:"grouping" omitEmpty:"true"`

	// The maximum number aggregatedComputedUsages of items to return within the Subscription "List" call, this
	// counts the overall count across all items
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCI home region name in case home region is not us-ashburn-1 (IAD), e.g. ap-mumbai-1, us-phoenix-1 etc.
	XOneOriginRegion *string `mandatory:"false" contributesTo:"header" name:"x-one-origin-region"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListComputedUsageAggregatedsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListComputedUsageAggregatedsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListComputedUsageAggregatedsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListComputedUsageAggregatedsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListComputedUsageAggregatedsResponse wrapper for the ListComputedUsageAggregateds operation
type ListComputedUsageAggregatedsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ComputedUsageAggregatedSummary instances
	Items []ComputedUsageAggregatedSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListComputedUsageAggregatedsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListComputedUsageAggregatedsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListComputedUsageAggregatedsGroupingEnum Enum with underlying type: string
type ListComputedUsageAggregatedsGroupingEnum string

// Set of constants representing the allowable values for ListComputedUsageAggregatedsGroupingEnum
const (
	ListComputedUsageAggregatedsGroupingHourly  ListComputedUsageAggregatedsGroupingEnum = "HOURLY"
	ListComputedUsageAggregatedsGroupingDaily   ListComputedUsageAggregatedsGroupingEnum = "DAILY"
	ListComputedUsageAggregatedsGroupingMonthly ListComputedUsageAggregatedsGroupingEnum = "MONTHLY"
	ListComputedUsageAggregatedsGroupingNone    ListComputedUsageAggregatedsGroupingEnum = "NONE"
)

var mappingListComputedUsageAggregatedsGrouping = map[string]ListComputedUsageAggregatedsGroupingEnum{
	"HOURLY":  ListComputedUsageAggregatedsGroupingHourly,
	"DAILY":   ListComputedUsageAggregatedsGroupingDaily,
	"MONTHLY": ListComputedUsageAggregatedsGroupingMonthly,
	"NONE":    ListComputedUsageAggregatedsGroupingNone,
}

// GetListComputedUsageAggregatedsGroupingEnumValues Enumerates the set of values for ListComputedUsageAggregatedsGroupingEnum
func GetListComputedUsageAggregatedsGroupingEnumValues() []ListComputedUsageAggregatedsGroupingEnum {
	values := make([]ListComputedUsageAggregatedsGroupingEnum, 0)
	for _, v := range mappingListComputedUsageAggregatedsGrouping {
		values = append(values, v)
	}
	return values
}
