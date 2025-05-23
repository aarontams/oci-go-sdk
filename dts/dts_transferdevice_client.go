// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Transfer Service API
//
// Data Transfer Service API Specification
//

package dts

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// TransferDeviceClient a client for TransferDevice
type TransferDeviceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewTransferDeviceClientWithConfigurationProvider Creates a new default TransferDevice client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewTransferDeviceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client TransferDeviceClient, err error) {
	if enabled := common.CheckForEnabledServices("dts"); !enabled {
		return client, fmt.Errorf("the Developer Tool configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local developer-tool-configuration.json file configured the service you're targeting or contact the cloud provider on the availability of this service")
	}
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newTransferDeviceClientFromBaseClient(baseClient, provider)
}

// NewTransferDeviceClientWithOboToken Creates a new default TransferDevice client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewTransferDeviceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client TransferDeviceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newTransferDeviceClientFromBaseClient(baseClient, configProvider)
}

func newTransferDeviceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client TransferDeviceClient, err error) {
	// TransferDevice service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("TransferDevice"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = TransferDeviceClient{BaseClient: baseClient}
	client.BasePath = "20171001"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *TransferDeviceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dts", "https://datatransfer.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *TransferDeviceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *TransferDeviceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateTransferDevice Create a new Transfer Device
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/CreateTransferDevice.go.html to see an example of how to use CreateTransferDevice API.
func (client TransferDeviceClient) CreateTransferDevice(ctx context.Context, request CreateTransferDeviceRequest) (response CreateTransferDeviceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createTransferDevice, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTransferDeviceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTransferDeviceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTransferDeviceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTransferDeviceResponse")
	}
	return
}

// createTransferDevice implements the OCIOperation interface (enables retrying operations)
func (client TransferDeviceClient) createTransferDevice(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/transferJobs/{id}/transferDevices", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTransferDeviceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "TransferDevice", "CreateTransferDevice", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTransferDevice deletes a transfer Device
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/DeleteTransferDevice.go.html to see an example of how to use DeleteTransferDevice API.
func (client TransferDeviceClient) DeleteTransferDevice(ctx context.Context, request DeleteTransferDeviceRequest) (response DeleteTransferDeviceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.deleteTransferDevice, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTransferDeviceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTransferDeviceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTransferDeviceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTransferDeviceResponse")
	}
	return
}

// deleteTransferDevice implements the OCIOperation interface (enables retrying operations)
func (client TransferDeviceClient) deleteTransferDevice(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/transferJobs/{id}/transferDevices/{transferDeviceLabel}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTransferDeviceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "TransferDevice", "DeleteTransferDevice", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTransferDevice Describes a transfer package in detail
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/GetTransferDevice.go.html to see an example of how to use GetTransferDevice API.
func (client TransferDeviceClient) GetTransferDevice(ctx context.Context, request GetTransferDeviceRequest) (response GetTransferDeviceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTransferDevice, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTransferDeviceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTransferDeviceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTransferDeviceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTransferDeviceResponse")
	}
	return
}

// getTransferDevice implements the OCIOperation interface (enables retrying operations)
func (client TransferDeviceClient) getTransferDevice(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/transferJobs/{id}/transferDevices/{transferDeviceLabel}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTransferDeviceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "TransferDevice", "GetTransferDevice", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTransferDevices Lists Transfer Devices associated with a transferJob
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/ListTransferDevices.go.html to see an example of how to use ListTransferDevices API.
func (client TransferDeviceClient) ListTransferDevices(ctx context.Context, request ListTransferDevicesRequest) (response ListTransferDevicesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTransferDevices, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTransferDevicesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTransferDevicesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTransferDevicesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTransferDevicesResponse")
	}
	return
}

// listTransferDevices implements the OCIOperation interface (enables retrying operations)
func (client TransferDeviceClient) listTransferDevices(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/transferJobs/{id}/transferDevices", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTransferDevicesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "TransferDevice", "ListTransferDevices", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTransferDevice Updates a Transfer Device
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/UpdateTransferDevice.go.html to see an example of how to use UpdateTransferDevice API.
func (client TransferDeviceClient) UpdateTransferDevice(ctx context.Context, request UpdateTransferDeviceRequest) (response UpdateTransferDeviceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTransferDevice, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTransferDeviceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTransferDeviceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTransferDeviceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTransferDeviceResponse")
	}
	return
}

// updateTransferDevice implements the OCIOperation interface (enables retrying operations)
func (client TransferDeviceClient) updateTransferDevice(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/transferJobs/{id}/transferDevices/{transferDeviceLabel}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTransferDeviceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "TransferDevice", "UpdateTransferDevice", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
