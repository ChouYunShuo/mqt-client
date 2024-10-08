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
)


// ProgramFundingAPIService ProgramFundingAPI service
type ProgramFundingAPIService service

type ProgramFundingAPIGetProgramFundingsRequest struct {
	ctx context.Context
	ApiService *ProgramFundingAPIService
	count *int32
	startIndex *int32
	startDate *string
	endDate *string
	shortCode *string
}

// Number of program funding resources to retrieve.
func (r ProgramFundingAPIGetProgramFundingsRequest) Count(count int32) ProgramFundingAPIGetProgramFundingsRequest {
	r.count = &count
	return r
}

// Sort order index of the first resource in the returned array.
func (r ProgramFundingAPIGetProgramFundingsRequest) StartIndex(startIndex int32) ProgramFundingAPIGetProgramFundingsRequest {
	r.startIndex = &startIndex
	return r
}

// Start date for filtering program funding entries.
func (r ProgramFundingAPIGetProgramFundingsRequest) StartDate(startDate string) ProgramFundingAPIGetProgramFundingsRequest {
	r.startDate = &startDate
	return r
}

// End date for filtering program funding entries.
func (r ProgramFundingAPIGetProgramFundingsRequest) EndDate(endDate string) ProgramFundingAPIGetProgramFundingsRequest {
	r.endDate = &endDate
	return r
}

// Short code for filtering program funding entries.
func (r ProgramFundingAPIGetProgramFundingsRequest) ShortCode(shortCode string) ProgramFundingAPIGetProgramFundingsRequest {
	r.shortCode = &shortCode
	return r
}

func (r ProgramFundingAPIGetProgramFundingsRequest) Execute() (*ProgramFundingPage, *http.Response, error) {
	return r.ApiService.GetProgramFundingsExecute(r)
}

/*
GetProgramFundings List program fundings

Retrieve an array of program funding entries.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ProgramFundingAPIGetProgramFundingsRequest
*/
func (a *ProgramFundingAPIService) GetProgramFundings(ctx context.Context) ProgramFundingAPIGetProgramFundingsRequest {
	return ProgramFundingAPIGetProgramFundingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ProgramFundingPage
func (a *ProgramFundingAPIService) GetProgramFundingsExecute(r ProgramFundingAPIGetProgramFundingsRequest) (*ProgramFundingPage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProgramFundingPage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProgramFundingAPIService.GetProgramFundings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/programs/funding"

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
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "")
	}
	if r.shortCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "short_code", r.shortCode, "")
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
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ProgramFundingAPIGetProgramFundingsByShortCodeRequest struct {
	ctx context.Context
	ApiService *ProgramFundingAPIService
	count *int32
	startIndex *int32
	startDate *string
	endDate *string
}

// Number of program funding resources to retrieve.
func (r ProgramFundingAPIGetProgramFundingsByShortCodeRequest) Count(count int32) ProgramFundingAPIGetProgramFundingsByShortCodeRequest {
	r.count = &count
	return r
}

// Sort order index of the first resource in the returned array.
func (r ProgramFundingAPIGetProgramFundingsByShortCodeRequest) StartIndex(startIndex int32) ProgramFundingAPIGetProgramFundingsByShortCodeRequest {
	r.startIndex = &startIndex
	return r
}

// Start date for filtering program funding entries.
func (r ProgramFundingAPIGetProgramFundingsByShortCodeRequest) StartDate(startDate string) ProgramFundingAPIGetProgramFundingsByShortCodeRequest {
	r.startDate = &startDate
	return r
}

// End date for filtering program funding entries.
func (r ProgramFundingAPIGetProgramFundingsByShortCodeRequest) EndDate(endDate string) ProgramFundingAPIGetProgramFundingsByShortCodeRequest {
	r.endDate = &endDate
	return r
}

func (r ProgramFundingAPIGetProgramFundingsByShortCodeRequest) Execute() (*ProgramFundingPage, *http.Response, error) {
	return r.ApiService.GetProgramFundingsByShortCodeExecute(r)
}

/*
GetProgramFundingsByShortCode List program fundings

Retrieve an array of program funding entries.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ProgramFundingAPIGetProgramFundingsByShortCodeRequest
*/
func (a *ProgramFundingAPIService) GetProgramFundingsByShortCode(ctx context.Context) ProgramFundingAPIGetProgramFundingsByShortCodeRequest {
	return ProgramFundingAPIGetProgramFundingsByShortCodeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ProgramFundingPage
func (a *ProgramFundingAPIService) GetProgramFundingsByShortCodeExecute(r ProgramFundingAPIGetProgramFundingsByShortCodeRequest) (*ProgramFundingPage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProgramFundingPage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProgramFundingAPIService.GetProgramFundingsByShortCode")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/programs/funding"

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
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "")
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
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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
