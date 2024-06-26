//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RoleAssignmentScheduleRequestsClient contains the methods for the RoleAssignmentScheduleRequests group.
// Don't use this type directly, use NewRoleAssignmentScheduleRequestsClient() instead.
type RoleAssignmentScheduleRequestsClient struct {
	internal *arm.Client
}

// NewRoleAssignmentScheduleRequestsClient creates a new instance of RoleAssignmentScheduleRequestsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRoleAssignmentScheduleRequestsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*RoleAssignmentScheduleRequestsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RoleAssignmentScheduleRequestsClient{
		internal: cl,
	}
	return client, nil
}

// Cancel - Cancels a pending role assignment schedule request.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-01
//   - scope - The scope of the role assignment request to cancel.
//   - roleAssignmentScheduleRequestName - The name of the role assignment request to cancel.
//   - options - RoleAssignmentScheduleRequestsClientCancelOptions contains the optional parameters for the RoleAssignmentScheduleRequestsClient.Cancel
//     method.
func (client *RoleAssignmentScheduleRequestsClient) Cancel(ctx context.Context, scope string, roleAssignmentScheduleRequestName string, options *RoleAssignmentScheduleRequestsClientCancelOptions) (RoleAssignmentScheduleRequestsClientCancelResponse, error) {
	var err error
	const operationName = "RoleAssignmentScheduleRequestsClient.Cancel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.cancelCreateRequest(ctx, scope, roleAssignmentScheduleRequestName, options)
	if err != nil {
		return RoleAssignmentScheduleRequestsClientCancelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoleAssignmentScheduleRequestsClientCancelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RoleAssignmentScheduleRequestsClientCancelResponse{}, err
	}
	return RoleAssignmentScheduleRequestsClientCancelResponse{}, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *RoleAssignmentScheduleRequestsClient) cancelCreateRequest(ctx context.Context, scope string, roleAssignmentScheduleRequestName string, options *RoleAssignmentScheduleRequestsClientCancelOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignmentScheduleRequests/{roleAssignmentScheduleRequestName}/cancel"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentScheduleRequestName == "" {
		return nil, errors.New("parameter roleAssignmentScheduleRequestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentScheduleRequestName}", url.PathEscape(roleAssignmentScheduleRequestName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Create - Creates a role assignment schedule request.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-01
//   - scope - The scope of the role assignment schedule request to create. The scope can be any REST resource instance. For example,
//     use '/subscriptions/{subscription-id}/' for a subscription,
//     '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}' for a resource group, and
//     '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider}/{resource-type}/{resource-name}'
//     for a resource.
//   - roleAssignmentScheduleRequestName - A GUID for the role assignment to create. The name must be unique and different for
//     each role assignment.
//   - parameters - Parameters for the role assignment schedule request.
//   - options - RoleAssignmentScheduleRequestsClientCreateOptions contains the optional parameters for the RoleAssignmentScheduleRequestsClient.Create
//     method.
func (client *RoleAssignmentScheduleRequestsClient) Create(ctx context.Context, scope string, roleAssignmentScheduleRequestName string, parameters RoleAssignmentScheduleRequest, options *RoleAssignmentScheduleRequestsClientCreateOptions) (RoleAssignmentScheduleRequestsClientCreateResponse, error) {
	var err error
	const operationName = "RoleAssignmentScheduleRequestsClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, scope, roleAssignmentScheduleRequestName, parameters, options)
	if err != nil {
		return RoleAssignmentScheduleRequestsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoleAssignmentScheduleRequestsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return RoleAssignmentScheduleRequestsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *RoleAssignmentScheduleRequestsClient) createCreateRequest(ctx context.Context, scope string, roleAssignmentScheduleRequestName string, parameters RoleAssignmentScheduleRequest, options *RoleAssignmentScheduleRequestsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignmentScheduleRequests/{roleAssignmentScheduleRequestName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentScheduleRequestName == "" {
		return nil, errors.New("parameter roleAssignmentScheduleRequestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentScheduleRequestName}", url.PathEscape(roleAssignmentScheduleRequestName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *RoleAssignmentScheduleRequestsClient) createHandleResponse(resp *http.Response) (RoleAssignmentScheduleRequestsClientCreateResponse, error) {
	result := RoleAssignmentScheduleRequestsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentScheduleRequest); err != nil {
		return RoleAssignmentScheduleRequestsClientCreateResponse{}, err
	}
	return result, nil
}

// Get - Get the specified role assignment schedule request.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-01
//   - scope - The scope of the role assignment schedule request.
//   - roleAssignmentScheduleRequestName - The name (guid) of the role assignment schedule request to get.
//   - options - RoleAssignmentScheduleRequestsClientGetOptions contains the optional parameters for the RoleAssignmentScheduleRequestsClient.Get
//     method.
func (client *RoleAssignmentScheduleRequestsClient) Get(ctx context.Context, scope string, roleAssignmentScheduleRequestName string, options *RoleAssignmentScheduleRequestsClientGetOptions) (RoleAssignmentScheduleRequestsClientGetResponse, error) {
	var err error
	const operationName = "RoleAssignmentScheduleRequestsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, scope, roleAssignmentScheduleRequestName, options)
	if err != nil {
		return RoleAssignmentScheduleRequestsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoleAssignmentScheduleRequestsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RoleAssignmentScheduleRequestsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RoleAssignmentScheduleRequestsClient) getCreateRequest(ctx context.Context, scope string, roleAssignmentScheduleRequestName string, options *RoleAssignmentScheduleRequestsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignmentScheduleRequests/{roleAssignmentScheduleRequestName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentScheduleRequestName == "" {
		return nil, errors.New("parameter roleAssignmentScheduleRequestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentScheduleRequestName}", url.PathEscape(roleAssignmentScheduleRequestName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoleAssignmentScheduleRequestsClient) getHandleResponse(resp *http.Response) (RoleAssignmentScheduleRequestsClientGetResponse, error) {
	result := RoleAssignmentScheduleRequestsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentScheduleRequest); err != nil {
		return RoleAssignmentScheduleRequestsClientGetResponse{}, err
	}
	return result, nil
}

// NewListForScopePager - Gets role assignment schedule requests for a scope.
//
// Generated from API version 2020-10-01
//   - scope - The scope of the role assignments schedule requests.
//   - options - RoleAssignmentScheduleRequestsClientListForScopeOptions contains the optional parameters for the RoleAssignmentScheduleRequestsClient.NewListForScopePager
//     method.
func (client *RoleAssignmentScheduleRequestsClient) NewListForScopePager(scope string, options *RoleAssignmentScheduleRequestsClientListForScopeOptions) *runtime.Pager[RoleAssignmentScheduleRequestsClientListForScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[RoleAssignmentScheduleRequestsClientListForScopeResponse]{
		More: func(page RoleAssignmentScheduleRequestsClientListForScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RoleAssignmentScheduleRequestsClientListForScopeResponse) (RoleAssignmentScheduleRequestsClientListForScopeResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RoleAssignmentScheduleRequestsClient.NewListForScopePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listForScopeCreateRequest(ctx, scope, options)
			}, nil)
			if err != nil {
				return RoleAssignmentScheduleRequestsClientListForScopeResponse{}, err
			}
			return client.listForScopeHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *RoleAssignmentScheduleRequestsClient) listForScopeCreateRequest(ctx context.Context, scope string, options *RoleAssignmentScheduleRequestsClientListForScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignmentScheduleRequests"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForScopeHandleResponse handles the ListForScope response.
func (client *RoleAssignmentScheduleRequestsClient) listForScopeHandleResponse(resp *http.Response) (RoleAssignmentScheduleRequestsClientListForScopeResponse, error) {
	result := RoleAssignmentScheduleRequestsClientListForScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentScheduleRequestListResult); err != nil {
		return RoleAssignmentScheduleRequestsClientListForScopeResponse{}, err
	}
	return result, nil
}

// Validate - Validates a new role assignment schedule request.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-01
//   - scope - The scope of the role assignment request to validate.
//   - roleAssignmentScheduleRequestName - The name of the role assignment request to validate.
//   - parameters - Parameters for the role assignment schedule request.
//   - options - RoleAssignmentScheduleRequestsClientValidateOptions contains the optional parameters for the RoleAssignmentScheduleRequestsClient.Validate
//     method.
func (client *RoleAssignmentScheduleRequestsClient) Validate(ctx context.Context, scope string, roleAssignmentScheduleRequestName string, parameters RoleAssignmentScheduleRequest, options *RoleAssignmentScheduleRequestsClientValidateOptions) (RoleAssignmentScheduleRequestsClientValidateResponse, error) {
	var err error
	const operationName = "RoleAssignmentScheduleRequestsClient.Validate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.validateCreateRequest(ctx, scope, roleAssignmentScheduleRequestName, parameters, options)
	if err != nil {
		return RoleAssignmentScheduleRequestsClientValidateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoleAssignmentScheduleRequestsClientValidateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RoleAssignmentScheduleRequestsClientValidateResponse{}, err
	}
	resp, err := client.validateHandleResponse(httpResp)
	return resp, err
}

// validateCreateRequest creates the Validate request.
func (client *RoleAssignmentScheduleRequestsClient) validateCreateRequest(ctx context.Context, scope string, roleAssignmentScheduleRequestName string, parameters RoleAssignmentScheduleRequest, options *RoleAssignmentScheduleRequestsClientValidateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignmentScheduleRequests/{roleAssignmentScheduleRequestName}/validate"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentScheduleRequestName == "" {
		return nil, errors.New("parameter roleAssignmentScheduleRequestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentScheduleRequestName}", url.PathEscape(roleAssignmentScheduleRequestName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// validateHandleResponse handles the Validate response.
func (client *RoleAssignmentScheduleRequestsClient) validateHandleResponse(resp *http.Response) (RoleAssignmentScheduleRequestsClientValidateResponse, error) {
	result := RoleAssignmentScheduleRequestsClientValidateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentScheduleRequest); err != nil {
		return RoleAssignmentScheduleRequestsClientValidateResponse{}, err
	}
	return result, nil
}
