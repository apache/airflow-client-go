// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

/*
Airflow API (Stable)

# Overview  To facilitate management, Apache Airflow supports a range of REST API endpoints across its objects. This section provides an overview of the API design, methods, and supported use cases.  Most of the endpoints accept `JSON` as input and return `JSON` responses. This means that you must usually add the following headers to your request: ``` Content-type: application/json Accept: application/json ```  ## Resources  The term `resource` refers to a single type of object in the Airflow metadata. An API is broken up by its endpoint's corresponding resource. The name of a resource is typically plural and expressed in camelCase. Example: `dagRuns`.  Resource names are used as part of endpoint URLs, as well as in API parameters and responses.  ## CRUD Operations  The platform supports **C**reate, **R**ead, **U**pdate, and **D**elete operations on most resources. You can review the standards for these operations and their standard parameters below.  Some endpoints have special behavior as exceptions.  ### Create  To create a resource, you typically submit an HTTP `POST` request with the resource's required metadata in the request body. The response returns a `201 Created` response code upon success with the resource's metadata, including its internal `id`, in the response body.  ### Read  The HTTP `GET` request can be used to read a resource or to list a number of resources.  A resource's `id` can be submitted in the request parameters to read a specific resource. The response usually returns a `200 OK` response code upon success, with the resource's metadata in the response body.  If a `GET` request does not include a specific resource `id`, it is treated as a list request. The response usually returns a `200 OK` response code upon success, with an object containing a list of resources' metadata in the response body.  When reading resources, some common query parameters are usually available. e.g.: ``` v1/connections?limit=25&offset=25 ```  |Query Parameter|Type|Description| |---------------|----|-----------| |limit|integer|Maximum number of objects to fetch. Usually 25 by default| |offset|integer|Offset after which to start returning objects. For use with limit query parameter.|  ### Update  Updating a resource requires the resource `id`, and is typically done using an HTTP `PATCH` request, with the fields to modify in the request body. The response usually returns a `200 OK` response code upon success, with information about the modified resource in the response body.  ### Delete  Deleting a resource requires the resource `id` and is typically executed via an HTTP `DELETE` request. The response usually returns a `204 No Content` response code upon success.  ## Conventions  - Resource names are plural and expressed in camelCase. - Names are consistent between URL parameter name and field name.  - Field names are in snake_case. ```json {     \"description\": \"string\",     \"name\": \"string\",     \"occupied_slots\": 0,     \"open_slots\": 0     \"queued_slots\": 0,     \"running_slots\": 0,     \"scheduled_slots\": 0,     \"slots\": 0, } ```  ### Update Mask  Update mask is available as a query parameter in patch endpoints. It is used to notify the API which fields you want to update. Using `update_mask` makes it easier to update objects by helping the server know which fields to update in an object instead of updating all fields. The update request ignores any fields that aren't specified in the field mask, leaving them with their current values.  Example: ```   resource = request.get('/resource/my-id').json()   resource['my_field'] = 'new-value'   request.patch('/resource/my-id?update_mask=my_field', data=json.dumps(resource)) ```  ## Versioning and Endpoint Lifecycle  - API versioning is not synchronized to specific releases of the Apache Airflow. - APIs are designed to be backward compatible. - Any changes to the API will first go through a deprecation phase.  # Trying the API  You can use a third party client, such as [curl](https://curl.haxx.se/), [HTTPie](https://httpie.org/), [Postman](https://www.postman.com/) or [the Insomnia rest client](https://insomnia.rest/) to test the Apache Airflow API.  Note that you will need to pass credentials data.  For e.g., here is how to pause a DAG with [curl](https://curl.haxx.se/), when basic authorization is used: ```bash curl -X PATCH 'https://example.com/api/v1/dags/{dag_id}?update_mask=is_paused' \\ -H 'Content-Type: application/json' \\ --user \"username:password\" \\ -d '{     \"is_paused\": true }' ```  Using a graphical tool such as [Postman](https://www.postman.com/) or [Insomnia](https://insomnia.rest/), it is possible to import the API specifications directly:  1. Download the API specification by clicking the **Download** button at the top of this document 2. Import the JSON specification in the graphical tool of your choice.   - In *Postman*, you can click the **import** button at the top   - With *Insomnia*, you can just drag-and-drop the file on the UI  Note that with *Postman*, you can also generate code snippets by selecting a request and clicking on the **Code** button.  ## Enabling CORS  [Cross-origin resource sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) is a browser security feature that restricts HTTP requests that are initiated from scripts running in the browser.  For details on enabling/configuring CORS, see [Enabling CORS](https://airflow.apache.org/docs/apache-airflow/stable/security/api.html).  # Authentication  To be able to meet the requirements of many organizations, Airflow supports many authentication methods, and it is even possible to add your own method.  If you want to check which auth backend is currently set, you can use `airflow config get-value api auth_backends` command as in the example below. ```bash $ airflow config get-value api auth_backends airflow.api.auth.backend.basic_auth ``` The default is to deny all requests.  For details on configuring the authentication, see [API Authorization](https://airflow.apache.org/docs/apache-airflow/stable/security/api.html).  # Errors  We follow the error response format proposed in [RFC 7807](https://tools.ietf.org/html/rfc7807) also known as Problem Details for HTTP APIs. As with our normal API responses, your client must be prepared to gracefully handle additional members of the response.  ## Unauthenticated  This indicates that the request has not been applied because it lacks valid authentication credentials for the target resource. Please check that you have valid credentials.  ## PermissionDenied  This response means that the server understood the request but refuses to authorize it because it lacks sufficient rights to the resource. It happens when you do not have the necessary permission to execute the action you performed. You need to get the appropriate permissions in other to resolve this error.  ## BadRequest  This response means that the server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing). To resolve this, please ensure that your syntax is correct.  ## NotFound  This client error response indicates that the server cannot find the requested resource.  ## MethodNotAllowed  Indicates that the request method is known by the server but is not supported by the target resource.  ## NotAcceptable  The target resource does not have a current representation that would be acceptable to the user agent, according to the proactive negotiation header fields received in the request, and the server is unwilling to supply a default representation.  ## AlreadyExists  The request could not be completed due to a conflict with the current state of the target resource, e.g. the resource it tries to create already exists.  ## Unknown  This means that the server encountered an unexpected condition that prevented it from fulfilling the request. 

API version: 2.7.0
Contact: dev@airflow.apache.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package airflow

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"reflect"
)

// Linger please
var (
	_ context.Context
)

// DAGRunApiService DAGRunApi service
type DAGRunApiService service

type DAGRunApiApiClearDagRunRequest struct {
	ctx context.Context
	ApiService *DAGRunApiService
	dagId string
	dagRunId string
	clearDagRun *ClearDagRun
}

func (r DAGRunApiApiClearDagRunRequest) ClearDagRun(clearDagRun ClearDagRun) DAGRunApiApiClearDagRunRequest {
	r.clearDagRun = &clearDagRun
	return r
}

func (r DAGRunApiApiClearDagRunRequest) Execute() (*AnyOfDAGRunTaskInstanceCollection, *http.Response, error) {
	return r.ApiService.ClearDagRunExecute(r)
}

/*
ClearDagRun Clear a DAG run

Clear a DAG run.

*New in version 2.4.0*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @return DAGRunApiApiClearDagRunRequest
*/
func (a *DAGRunApiService) ClearDagRun(ctx context.Context, dagId string, dagRunId string) DAGRunApiApiClearDagRunRequest {
	return DAGRunApiApiClearDagRunRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
	}
}

// Execute executes the request
//  @return AnyOfDAGRunTaskInstanceCollection
func (a *DAGRunApiService) ClearDagRunExecute(r DAGRunApiApiClearDagRunRequest) (*AnyOfDAGRunTaskInstanceCollection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AnyOfDAGRunTaskInstanceCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DAGRunApiService.ClearDagRun")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/clear"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterToString(r.dagId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterToString(r.dagRunId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clearDagRun == nil {
		return localVarReturnValue, nil, reportError("clearDagRun is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.clearDagRun
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type DAGRunApiApiDeleteDagRunRequest struct {
	ctx context.Context
	ApiService *DAGRunApiService
	dagId string
	dagRunId string
}


func (r DAGRunApiApiDeleteDagRunRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteDagRunExecute(r)
}

/*
DeleteDagRun Delete a DAG run

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @return DAGRunApiApiDeleteDagRunRequest
*/
func (a *DAGRunApiService) DeleteDagRun(ctx context.Context, dagId string, dagRunId string) DAGRunApiApiDeleteDagRunRequest {
	return DAGRunApiApiDeleteDagRunRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
	}
}

// Execute executes the request
func (a *DAGRunApiService) DeleteDagRunExecute(r DAGRunApiApiDeleteDagRunRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DAGRunApiService.DeleteDagRun")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterToString(r.dagId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterToString(r.dagRunId, "")), -1)

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type DAGRunApiApiGetDagRunRequest struct {
	ctx context.Context
	ApiService *DAGRunApiService
	dagId string
	dagRunId string
}


func (r DAGRunApiApiGetDagRunRequest) Execute() (*DAGRun, *http.Response, error) {
	return r.ApiService.GetDagRunExecute(r)
}

/*
GetDagRun Get a DAG run

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @return DAGRunApiApiGetDagRunRequest
*/
func (a *DAGRunApiService) GetDagRun(ctx context.Context, dagId string, dagRunId string) DAGRunApiApiGetDagRunRequest {
	return DAGRunApiApiGetDagRunRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
	}
}

// Execute executes the request
//  @return DAGRun
func (a *DAGRunApiService) GetDagRunExecute(r DAGRunApiApiGetDagRunRequest) (*DAGRun, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DAGRun
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DAGRunApiService.GetDagRun")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterToString(r.dagId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterToString(r.dagRunId, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type DAGRunApiApiGetDagRunsRequest struct {
	ctx context.Context
	ApiService *DAGRunApiService
	dagId string
	limit *int32
	offset *int32
	executionDateGte *time.Time
	executionDateLte *time.Time
	startDateGte *time.Time
	startDateLte *time.Time
	endDateGte *time.Time
	endDateLte *time.Time
	updatedAtGte *time.Time
	updatedAtLte *time.Time
	state *[]string
	orderBy *string
}

// The numbers of items to return.
func (r DAGRunApiApiGetDagRunsRequest) Limit(limit int32) DAGRunApiApiGetDagRunsRequest {
	r.limit = &limit
	return r
}
// The number of items to skip before starting to collect the result set.
func (r DAGRunApiApiGetDagRunsRequest) Offset(offset int32) DAGRunApiApiGetDagRunsRequest {
	r.offset = &offset
	return r
}
// Returns objects greater or equal to the specified date.  This can be combined with execution_date_lte parameter to receive only the selected period. 
func (r DAGRunApiApiGetDagRunsRequest) ExecutionDateGte(executionDateGte time.Time) DAGRunApiApiGetDagRunsRequest {
	r.executionDateGte = &executionDateGte
	return r
}
// Returns objects less than or equal to the specified date.  This can be combined with execution_date_gte parameter to receive only the selected period. 
func (r DAGRunApiApiGetDagRunsRequest) ExecutionDateLte(executionDateLte time.Time) DAGRunApiApiGetDagRunsRequest {
	r.executionDateLte = &executionDateLte
	return r
}
// Returns objects greater or equal the specified date.  This can be combined with start_date_lte parameter to receive only the selected period. 
func (r DAGRunApiApiGetDagRunsRequest) StartDateGte(startDateGte time.Time) DAGRunApiApiGetDagRunsRequest {
	r.startDateGte = &startDateGte
	return r
}
// Returns objects less or equal the specified date.  This can be combined with start_date_gte parameter to receive only the selected period. 
func (r DAGRunApiApiGetDagRunsRequest) StartDateLte(startDateLte time.Time) DAGRunApiApiGetDagRunsRequest {
	r.startDateLte = &startDateLte
	return r
}
// Returns objects greater or equal the specified date.  This can be combined with start_date_lte parameter to receive only the selected period. 
func (r DAGRunApiApiGetDagRunsRequest) EndDateGte(endDateGte time.Time) DAGRunApiApiGetDagRunsRequest {
	r.endDateGte = &endDateGte
	return r
}
// Returns objects less than or equal to the specified date.  This can be combined with start_date_gte parameter to receive only the selected period. 
func (r DAGRunApiApiGetDagRunsRequest) EndDateLte(endDateLte time.Time) DAGRunApiApiGetDagRunsRequest {
	r.endDateLte = &endDateLte
	return r
}
// Returns objects greater or equal the specified date.  This can be combined with updated_at_lte parameter to receive only the selected period.  *New in version 2.6.0* 
func (r DAGRunApiApiGetDagRunsRequest) UpdatedAtGte(updatedAtGte time.Time) DAGRunApiApiGetDagRunsRequest {
	r.updatedAtGte = &updatedAtGte
	return r
}
// Returns objects less or equal the specified date.  This can be combined with updated_at_gte parameter to receive only the selected period.  *New in version 2.6.0* 
func (r DAGRunApiApiGetDagRunsRequest) UpdatedAtLte(updatedAtLte time.Time) DAGRunApiApiGetDagRunsRequest {
	r.updatedAtLte = &updatedAtLte
	return r
}
// The value can be repeated to retrieve multiple matching values (OR condition).
func (r DAGRunApiApiGetDagRunsRequest) State(state []string) DAGRunApiApiGetDagRunsRequest {
	r.state = &state
	return r
}
// The name of the field to order the results by. Prefix a field name with &#x60;-&#x60; to reverse the sort order.  *New in version 2.1.0* 
func (r DAGRunApiApiGetDagRunsRequest) OrderBy(orderBy string) DAGRunApiApiGetDagRunsRequest {
	r.orderBy = &orderBy
	return r
}

func (r DAGRunApiApiGetDagRunsRequest) Execute() (*DAGRunCollection, *http.Response, error) {
	return r.ApiService.GetDagRunsExecute(r)
}

/*
GetDagRuns List DAG runs

This endpoint allows specifying `~` as the dag_id to retrieve DAG runs for all DAGs.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @return DAGRunApiApiGetDagRunsRequest
*/
func (a *DAGRunApiService) GetDagRuns(ctx context.Context, dagId string) DAGRunApiApiGetDagRunsRequest {
	return DAGRunApiApiGetDagRunsRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
	}
}

// Execute executes the request
//  @return DAGRunCollection
func (a *DAGRunApiService) GetDagRunsExecute(r DAGRunApiApiGetDagRunsRequest) (*DAGRunCollection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DAGRunCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DAGRunApiService.GetDagRuns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterToString(r.dagId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.executionDateGte != nil {
		localVarQueryParams.Add("execution_date_gte", parameterToString(*r.executionDateGte, ""))
	}
	if r.executionDateLte != nil {
		localVarQueryParams.Add("execution_date_lte", parameterToString(*r.executionDateLte, ""))
	}
	if r.startDateGte != nil {
		localVarQueryParams.Add("start_date_gte", parameterToString(*r.startDateGte, ""))
	}
	if r.startDateLte != nil {
		localVarQueryParams.Add("start_date_lte", parameterToString(*r.startDateLte, ""))
	}
	if r.endDateGte != nil {
		localVarQueryParams.Add("end_date_gte", parameterToString(*r.endDateGte, ""))
	}
	if r.endDateLte != nil {
		localVarQueryParams.Add("end_date_lte", parameterToString(*r.endDateLte, ""))
	}
	if r.updatedAtGte != nil {
		localVarQueryParams.Add("updated_at_gte", parameterToString(*r.updatedAtGte, ""))
	}
	if r.updatedAtLte != nil {
		localVarQueryParams.Add("updated_at_lte", parameterToString(*r.updatedAtLte, ""))
	}
	if r.state != nil {
		t := *r.state
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("state", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("state", parameterToString(t, "multi"))
		}
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("order_by", parameterToString(*r.orderBy, ""))
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type DAGRunApiApiGetDagRunsBatchRequest struct {
	ctx context.Context
	ApiService *DAGRunApiService
	listDagRunsForm *ListDagRunsForm
}

func (r DAGRunApiApiGetDagRunsBatchRequest) ListDagRunsForm(listDagRunsForm ListDagRunsForm) DAGRunApiApiGetDagRunsBatchRequest {
	r.listDagRunsForm = &listDagRunsForm
	return r
}

func (r DAGRunApiApiGetDagRunsBatchRequest) Execute() (*DAGRunCollection, *http.Response, error) {
	return r.ApiService.GetDagRunsBatchExecute(r)
}

/*
GetDagRunsBatch List DAG runs (batch)

This endpoint is a POST to allow filtering across a large number of DAG IDs, where as a GET it would run in to maximum HTTP request URL length limit.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DAGRunApiApiGetDagRunsBatchRequest
*/
func (a *DAGRunApiService) GetDagRunsBatch(ctx context.Context) DAGRunApiApiGetDagRunsBatchRequest {
	return DAGRunApiApiGetDagRunsBatchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DAGRunCollection
func (a *DAGRunApiService) GetDagRunsBatchExecute(r DAGRunApiApiGetDagRunsBatchRequest) (*DAGRunCollection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DAGRunCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DAGRunApiService.GetDagRunsBatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/~/dagRuns/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.listDagRunsForm == nil {
		return localVarReturnValue, nil, reportError("listDagRunsForm is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.listDagRunsForm
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type DAGRunApiApiGetUpstreamDatasetEventsRequest struct {
	ctx context.Context
	ApiService *DAGRunApiService
	dagId string
	dagRunId string
}


func (r DAGRunApiApiGetUpstreamDatasetEventsRequest) Execute() (*DatasetEventCollection, *http.Response, error) {
	return r.ApiService.GetUpstreamDatasetEventsExecute(r)
}

/*
GetUpstreamDatasetEvents Get dataset events for a DAG run

Get datasets for a dag run.

*New in version 2.4.0*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @return DAGRunApiApiGetUpstreamDatasetEventsRequest
*/
func (a *DAGRunApiService) GetUpstreamDatasetEvents(ctx context.Context, dagId string, dagRunId string) DAGRunApiApiGetUpstreamDatasetEventsRequest {
	return DAGRunApiApiGetUpstreamDatasetEventsRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
	}
}

// Execute executes the request
//  @return DatasetEventCollection
func (a *DAGRunApiService) GetUpstreamDatasetEventsExecute(r DAGRunApiApiGetUpstreamDatasetEventsRequest) (*DatasetEventCollection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DatasetEventCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DAGRunApiService.GetUpstreamDatasetEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/upstreamDatasetEvents"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterToString(r.dagId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterToString(r.dagRunId, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type DAGRunApiApiPostDagRunRequest struct {
	ctx context.Context
	ApiService *DAGRunApiService
	dagId string
	dAGRun *DAGRun
}

func (r DAGRunApiApiPostDagRunRequest) DAGRun(dAGRun DAGRun) DAGRunApiApiPostDagRunRequest {
	r.dAGRun = &dAGRun
	return r
}

func (r DAGRunApiApiPostDagRunRequest) Execute() (*DAGRun, *http.Response, error) {
	return r.ApiService.PostDagRunExecute(r)
}

/*
PostDagRun Trigger a new DAG run.

This will initiate a dagrun. If DAG is paused then dagrun state will remain queued, and the task won't run.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @return DAGRunApiApiPostDagRunRequest
*/
func (a *DAGRunApiService) PostDagRun(ctx context.Context, dagId string) DAGRunApiApiPostDagRunRequest {
	return DAGRunApiApiPostDagRunRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
	}
}

// Execute executes the request
//  @return DAGRun
func (a *DAGRunApiService) PostDagRunExecute(r DAGRunApiApiPostDagRunRequest) (*DAGRun, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DAGRun
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DAGRunApiService.PostDagRun")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterToString(r.dagId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dAGRun == nil {
		return localVarReturnValue, nil, reportError("dAGRun is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.dAGRun
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type DAGRunApiApiSetDagRunNoteRequest struct {
	ctx context.Context
	ApiService *DAGRunApiService
	dagId string
	dagRunId string
	setDagRunNote *SetDagRunNote
}

// Parameters of set DagRun note.
func (r DAGRunApiApiSetDagRunNoteRequest) SetDagRunNote(setDagRunNote SetDagRunNote) DAGRunApiApiSetDagRunNoteRequest {
	r.setDagRunNote = &setDagRunNote
	return r
}

func (r DAGRunApiApiSetDagRunNoteRequest) Execute() (*DAGRun, *http.Response, error) {
	return r.ApiService.SetDagRunNoteExecute(r)
}

/*
SetDagRunNote Update the DagRun note.

Update the manual user note of a DagRun.

*New in version 2.5.0*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @return DAGRunApiApiSetDagRunNoteRequest
*/
func (a *DAGRunApiService) SetDagRunNote(ctx context.Context, dagId string, dagRunId string) DAGRunApiApiSetDagRunNoteRequest {
	return DAGRunApiApiSetDagRunNoteRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
	}
}

// Execute executes the request
//  @return DAGRun
func (a *DAGRunApiService) SetDagRunNoteExecute(r DAGRunApiApiSetDagRunNoteRequest) (*DAGRun, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DAGRun
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DAGRunApiService.SetDagRunNote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/setNote"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterToString(r.dagId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterToString(r.dagRunId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.setDagRunNote == nil {
		return localVarReturnValue, nil, reportError("setDagRunNote is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.setDagRunNote
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type DAGRunApiApiUpdateDagRunStateRequest struct {
	ctx context.Context
	ApiService *DAGRunApiService
	dagId string
	dagRunId string
	updateDagRunState *UpdateDagRunState
}

func (r DAGRunApiApiUpdateDagRunStateRequest) UpdateDagRunState(updateDagRunState UpdateDagRunState) DAGRunApiApiUpdateDagRunStateRequest {
	r.updateDagRunState = &updateDagRunState
	return r
}

func (r DAGRunApiApiUpdateDagRunStateRequest) Execute() (*DAGRun, *http.Response, error) {
	return r.ApiService.UpdateDagRunStateExecute(r)
}

/*
UpdateDagRunState Modify a DAG run

Modify a DAG run.

*New in version 2.2.0*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @return DAGRunApiApiUpdateDagRunStateRequest
*/
func (a *DAGRunApiService) UpdateDagRunState(ctx context.Context, dagId string, dagRunId string) DAGRunApiApiUpdateDagRunStateRequest {
	return DAGRunApiApiUpdateDagRunStateRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
	}
}

// Execute executes the request
//  @return DAGRun
func (a *DAGRunApiService) UpdateDagRunStateExecute(r DAGRunApiApiUpdateDagRunStateRequest) (*DAGRun, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DAGRun
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DAGRunApiService.UpdateDagRunState")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterToString(r.dagId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterToString(r.dagRunId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateDagRunState == nil {
		return localVarReturnValue, nil, reportError("updateDagRunState is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.updateDagRunState
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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
