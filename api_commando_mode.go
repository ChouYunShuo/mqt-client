/*
Core API

Marqeta's Core API endpoints, conveniently annotated to enable code generation (including SDKs), test cases, and documentation. Currently in beta.

API version: 3.0.19
Contact: support@marqeta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// CommandoModeAPIService CommandoModeAPI service
type CommandoModeAPIService service

type CommandoModeAPIGetCommandomodesRequest struct {
	ctx context.Context
	ApiService *CommandoModeAPIService
	count *int32
	startIndex *int32
	sortBy *string
}

// Number of Commando Mode control sets to retrieve.
func (r CommandoModeAPIGetCommandomodesRequest) Count(count int32) CommandoModeAPIGetCommandomodesRequest {
	r.count = &count
	return r
}

// Sort order index of the first resource in the returned array.
func (r CommandoModeAPIGetCommandomodesRequest) StartIndex(startIndex int32) CommandoModeAPIGetCommandomodesRequest {
	r.startIndex = &startIndex
	return r
}

// Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.
func (r CommandoModeAPIGetCommandomodesRequest) SortBy(sortBy string) CommandoModeAPIGetCommandomodesRequest {
	r.sortBy = &sortBy
	return r
}

func (r CommandoModeAPIGetCommandomodesRequest) Execute() (*CommandoModeListResponse, *http.Response, error) {
	return r.ApiService.GetCommandomodesExecute(r)
}

/*
GetCommandomodes List Commando Mode control sets

Retrieve a list of Commando Mode control sets.

This endpoint supports <</core-api/field-filtering, field filtering>> and <</core-api/sorting-and-pagination, pagination>>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CommandoModeAPIGetCommandomodesRequest
*/
func (a *CommandoModeAPIService) GetCommandomodes(ctx context.Context) CommandoModeAPIGetCommandomodesRequest {
	return CommandoModeAPIGetCommandomodesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CommandoModeListResponse
func (a *CommandoModeAPIService) GetCommandomodesExecute(r CommandoModeAPIGetCommandomodesRequest) (*CommandoModeListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CommandoModeListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommandoModeAPIService.GetCommandomodes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/commandomodes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	} else {
		var defaultValue int32 = 5
		r.count = &defaultValue
	}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_index", r.startIndex, "")
	} else {
		var defaultValue int32 = 0
		r.startIndex = &defaultValue
	}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_by", r.sortBy, "")
	} else {
		var defaultValue string = "-lastModifiedTime"
		r.sortBy = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CommandoModeAPIGetCommandomodesCommandomodetokenTransitionsRequest struct {
	ctx context.Context
	ApiService *CommandoModeAPIService
	commandomodeToken string
	count *int32
	startIndex *int32
	sortBy *string
}

// Number of Commando Mode control set transitions to retrieve.
func (r CommandoModeAPIGetCommandomodesCommandomodetokenTransitionsRequest) Count(count int32) CommandoModeAPIGetCommandomodesCommandomodetokenTransitionsRequest {
	r.count = &count
	return r
}

// Sort order index of the first resource in the returned array.
func (r CommandoModeAPIGetCommandomodesCommandomodetokenTransitionsRequest) StartIndex(startIndex int32) CommandoModeAPIGetCommandomodesCommandomodetokenTransitionsRequest {
	r.startIndex = &startIndex
	return r
}

// Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.
func (r CommandoModeAPIGetCommandomodesCommandomodetokenTransitionsRequest) SortBy(sortBy string) CommandoModeAPIGetCommandomodesCommandomodetokenTransitionsRequest {
	r.sortBy = &sortBy
	return r
}

func (r CommandoModeAPIGetCommandomodesCommandomodetokenTransitionsRequest) Execute() (*CommandoModeTransitionListResponse, *http.Response, error) {
	return r.ApiService.GetCommandomodesCommandomodetokenTransitionsExecute(r)
}

/*
GetCommandomodesCommandomodetokenTransitions List Commando Mode transitions set

Retrieve a list of Commando Mode transitions for a specific control set.

This endpoint supports <</core-api/field-filtering, field filtering>> and <</core-api/sorting-and-pagination, pagination>>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param commandomodeToken Unique identifier of the Commando Mode control set.
 @return CommandoModeAPIGetCommandomodesCommandomodetokenTransitionsRequest
*/
func (a *CommandoModeAPIService) GetCommandomodesCommandomodetokenTransitions(ctx context.Context, commandomodeToken string) CommandoModeAPIGetCommandomodesCommandomodetokenTransitionsRequest {
	return CommandoModeAPIGetCommandomodesCommandomodetokenTransitionsRequest{
		ApiService: a,
		ctx: ctx,
		commandomodeToken: commandomodeToken,
	}
}

// Execute executes the request
//  @return CommandoModeTransitionListResponse
func (a *CommandoModeAPIService) GetCommandomodesCommandomodetokenTransitionsExecute(r CommandoModeAPIGetCommandomodesCommandomodetokenTransitionsRequest) (*CommandoModeTransitionListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CommandoModeTransitionListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommandoModeAPIService.GetCommandomodesCommandomodetokenTransitions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/commandomodes/{commandomode_token}/transitions"
	localVarPath = strings.Replace(localVarPath, "{"+"commandomode_token"+"}", url.PathEscape(parameterValueToString(r.commandomodeToken, "commandomodeToken")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	} else {
		var defaultValue int32 = 5
		r.count = &defaultValue
	}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_index", r.startIndex, "")
	} else {
		var defaultValue int32 = 0
		r.startIndex = &defaultValue
	}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_by", r.sortBy, "")
	} else {
		var defaultValue string = "-createdTime"
		r.sortBy = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CommandoModeAPIGetCommandomodesTokenRequest struct {
	ctx context.Context
	ApiService *CommandoModeAPIService
	token string
}

func (r CommandoModeAPIGetCommandomodesTokenRequest) Execute() (*CommandoModeResponse, *http.Response, error) {
	return r.ApiService.GetCommandomodesTokenExecute(r)
}

/*
GetCommandomodesToken Retrieve Commando Mode control set

Retrieve a specific Commando Mode control set.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token Unique identifier of the Commando Mode control set you want to retrieve.
 @return CommandoModeAPIGetCommandomodesTokenRequest
*/
func (a *CommandoModeAPIService) GetCommandomodesToken(ctx context.Context, token string) CommandoModeAPIGetCommandomodesTokenRequest {
	return CommandoModeAPIGetCommandomodesTokenRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return CommandoModeResponse
func (a *CommandoModeAPIService) GetCommandomodesTokenExecute(r CommandoModeAPIGetCommandomodesTokenRequest) (*CommandoModeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CommandoModeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommandoModeAPIService.GetCommandomodesToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/commandomodes/{token}"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CommandoModeAPIGetCommandomodesTransitionsTokenRequest struct {
	ctx context.Context
	ApiService *CommandoModeAPIService
	token string
}

func (r CommandoModeAPIGetCommandomodesTransitionsTokenRequest) Execute() (*CommandoModeTransitionResponse, *http.Response, error) {
	return r.ApiService.GetCommandomodesTransitionsTokenExecute(r)
}

/*
GetCommandomodesTransitionsToken Retrieve Commando Mode transition

Retrieve a specific Commando Mode control set transition.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token Unique identifier of the Commando Mode control set transition.
 @return CommandoModeAPIGetCommandomodesTransitionsTokenRequest
*/
func (a *CommandoModeAPIService) GetCommandomodesTransitionsToken(ctx context.Context, token string) CommandoModeAPIGetCommandomodesTransitionsTokenRequest {
	return CommandoModeAPIGetCommandomodesTransitionsTokenRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return CommandoModeTransitionResponse
func (a *CommandoModeAPIService) GetCommandomodesTransitionsTokenExecute(r CommandoModeAPIGetCommandomodesTransitionsTokenRequest) (*CommandoModeTransitionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CommandoModeTransitionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommandoModeAPIService.GetCommandomodesTransitionsToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/commandomodes/transitions/{token}"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
