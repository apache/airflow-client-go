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

// TaskInstanceApiService TaskInstanceApi service
type TaskInstanceApiService service

type TaskInstanceApiApiGetExtraLinksRequest struct {
	ctx context.Context
	ApiService *TaskInstanceApiService
	dagId string
	dagRunId string
	taskId string
}


func (r TaskInstanceApiApiGetExtraLinksRequest) Execute() (*ExtraLinkCollection, *http.Response, error) {
	return r.ApiService.GetExtraLinksExecute(r)
}

/*
GetExtraLinks List extra links

List extra links for task instance.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @return TaskInstanceApiApiGetExtraLinksRequest
*/
func (a *TaskInstanceApiService) GetExtraLinks(ctx context.Context, dagId string, dagRunId string, taskId string) TaskInstanceApiApiGetExtraLinksRequest {
	return TaskInstanceApiApiGetExtraLinksRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return ExtraLinkCollection
func (a *TaskInstanceApiService) GetExtraLinksExecute(r TaskInstanceApiApiGetExtraLinksRequest) (*ExtraLinkCollection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExtraLinkCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceApiService.GetExtraLinks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/links"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterToString(r.dagId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterToString(r.dagRunId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterToString(r.taskId, "")), -1)

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

type TaskInstanceApiApiGetLogRequest struct {
	ctx context.Context
	ApiService *TaskInstanceApiService
	dagId string
	dagRunId string
	taskId string
	taskTryNumber int32
	fullContent *bool
	mapIndex *int32
	token *string
}

// A full content will be returned. By default, only the first fragment will be returned. 
func (r TaskInstanceApiApiGetLogRequest) FullContent(fullContent bool) TaskInstanceApiApiGetLogRequest {
	r.fullContent = &fullContent
	return r
}
// Filter on map index for mapped task.
func (r TaskInstanceApiApiGetLogRequest) MapIndex(mapIndex int32) TaskInstanceApiApiGetLogRequest {
	r.mapIndex = &mapIndex
	return r
}
// A token that allows you to continue fetching logs. If passed, it will specify the location from which the download should be continued. 
func (r TaskInstanceApiApiGetLogRequest) Token(token string) TaskInstanceApiApiGetLogRequest {
	r.token = &token
	return r
}

func (r TaskInstanceApiApiGetLogRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.ApiService.GetLogExecute(r)
}

/*
GetLog Get logs

Get logs for a specific task instance and its try number.
To get log from specific character position, following way of using
URLSafeSerializer can be used.

Example:
```
from itsdangerous.url_safe import URLSafeSerializer

request_url = f"api/v1/dags/{DAG_ID}/dagRuns/{RUN_ID}/taskInstances/{TASK_ID}/logs/1"
key = app.config["SECRET_KEY"]
serializer = URLSafeSerializer(key)
token = serializer.dumps({"log_pos": 10000})

response = self.client.get(
    request_url,
    query_string={"token": token},
    headers={"Accept": "text/plain"},
    environ_overrides={"REMOTE_USER": "test"},
)
continuation_token = response.json["continuation_token"]
    metadata = URLSafeSerializer(key).loads(continuation_token)
    log_pos = metadata["log_pos"]
    end_of_log = metadata["end_of_log"]
```
If log_pos is passed as 10000 like the above example, it renders the logs starting
from char position 10000 to last (not the end as the logs may be tailing behind in
running state). This way pagination can be done with metadata as part of the token.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @param taskTryNumber The task try number.
 @return TaskInstanceApiApiGetLogRequest
*/
func (a *TaskInstanceApiService) GetLog(ctx context.Context, dagId string, dagRunId string, taskId string, taskTryNumber int32) TaskInstanceApiApiGetLogRequest {
	return TaskInstanceApiApiGetLogRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
		taskTryNumber: taskTryNumber,
	}
}

// Execute executes the request
//  @return InlineResponse200
func (a *TaskInstanceApiService) GetLogExecute(r TaskInstanceApiApiGetLogRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceApiService.GetLog")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/logs/{task_try_number}"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterToString(r.dagId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterToString(r.dagRunId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterToString(r.taskId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_try_number"+"}", url.PathEscape(parameterToString(r.taskTryNumber, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fullContent != nil {
		localVarQueryParams.Add("full_content", parameterToString(*r.fullContent, ""))
	}
	if r.mapIndex != nil {
		localVarQueryParams.Add("map_index", parameterToString(*r.mapIndex, ""))
	}
	if r.token != nil {
		localVarQueryParams.Add("token", parameterToString(*r.token, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/plain"}

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

type TaskInstanceApiApiGetMappedTaskInstanceRequest struct {
	ctx context.Context
	ApiService *TaskInstanceApiService
	dagId string
	dagRunId string
	taskId string
	mapIndex int32
}


func (r TaskInstanceApiApiGetMappedTaskInstanceRequest) Execute() (*TaskInstance, *http.Response, error) {
	return r.ApiService.GetMappedTaskInstanceExecute(r)
}

/*
GetMappedTaskInstance Get a mapped task instance

Get details of a mapped task instance.

*New in version 2.3.0*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @param mapIndex The map index.
 @return TaskInstanceApiApiGetMappedTaskInstanceRequest
*/
func (a *TaskInstanceApiService) GetMappedTaskInstance(ctx context.Context, dagId string, dagRunId string, taskId string, mapIndex int32) TaskInstanceApiApiGetMappedTaskInstanceRequest {
	return TaskInstanceApiApiGetMappedTaskInstanceRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
		mapIndex: mapIndex,
	}
}

// Execute executes the request
//  @return TaskInstance
func (a *TaskInstanceApiService) GetMappedTaskInstanceExecute(r TaskInstanceApiApiGetMappedTaskInstanceRequest) (*TaskInstance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceApiService.GetMappedTaskInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/{map_index}"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterToString(r.dagId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterToString(r.dagRunId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterToString(r.taskId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"map_index"+"}", url.PathEscape(parameterToString(r.mapIndex, "")), -1)

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

type TaskInstanceApiApiGetMappedTaskInstancesRequest struct {
	ctx context.Context
	ApiService *TaskInstanceApiService
	dagId string
	dagRunId string
	taskId string
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
	durationGte *float32
	durationLte *float32
	state *[]string
	pool *[]string
	queue *[]string
	orderBy *string
}

// The numbers of items to return.
func (r TaskInstanceApiApiGetMappedTaskInstancesRequest) Limit(limit int32) TaskInstanceApiApiGetMappedTaskInstancesRequest {
	r.limit = &limit
	return r
}
// The number of items to skip before starting to collect the result set.
func (r TaskInstanceApiApiGetMappedTaskInstancesRequest) Offset(offset int32) TaskInstanceApiApiGetMappedTaskInstancesRequest {
	r.offset = &offset
	return r
}
// Returns objects greater or equal to the specified date.  This can be combined with execution_date_lte parameter to receive only the selected period. 
func (r TaskInstanceApiApiGetMappedTaskInstancesRequest) ExecutionDateGte(executionDateGte time.Time) TaskInstanceApiApiGetMappedTaskInstancesRequest {
	r.executionDateGte = &executionDateGte
	return r
}
// Returns objects less than or equal to the specified date.  This can be combined with execution_date_gte parameter to receive only the selected period. 
func (r TaskInstanceApiApiGetMappedTaskInstancesRequest) ExecutionDateLte(executionDateLte time.Time) TaskInstanceApiApiGetMappedTaskInstancesRequest {
	r.executionDateLte = &executionDateLte
	return r
}
// Returns objects greater or equal the specified date.  This can be combined with start_date_lte parameter to receive only the selected period. 
func (r TaskInstanceApiApiGetMappedTaskInstancesRequest) StartDateGte(startDateGte time.Time) TaskInstanceApiApiGetMappedTaskInstancesRequest {
	r.startDateGte = &startDateGte
	return r
}
// Returns objects less or equal the specified date.  This can be combined with start_date_gte parameter to receive only the selected period. 
func (r TaskInstanceApiApiGetMappedTaskInstancesRequest) StartDateLte(startDateLte time.Time) TaskInstanceApiApiGetMappedTaskInstancesRequest {
	r.startDateLte = &startDateLte
	return r
}
// Returns objects greater or equal the specified date.  This can be combined with start_date_lte parameter to receive only the selected period. 
func (r TaskInstanceApiApiGetMappedTaskInstancesRequest) EndDateGte(endDateGte time.Time) TaskInstanceApiApiGetMappedTaskInstancesRequest {
	r.endDateGte = &endDateGte
	return r
}
// Returns objects less than or equal to the specified date.  This can be combined with start_date_gte parameter to receive only the selected period. 
func (r TaskInstanceApiApiGetMappedTaskInstancesRequest) EndDateLte(endDateLte time.Time) TaskInstanceApiApiGetMappedTaskInstancesRequest {
	r.endDateLte = &endDateLte
	return r
}
// Returns objects greater or equal the specified date.  This can be combined with updated_at_lte parameter to receive only the selected period.  *New in version 2.6.0* 
func (r TaskInstanceApiApiGetMappedTaskInstancesRequest) UpdatedAtGte(updatedAtGte time.Time) TaskInstanceApiApiGetMappedTaskInstancesRequest {
	r.updatedAtGte = &updatedAtGte
	return r
}
// Returns objects less or equal the specified date.  This can be combined with updated_at_gte parameter to receive only the selected period.  *New in version 2.6.0* 
func (r TaskInstanceApiApiGetMappedTaskInstancesRequest) UpdatedAtLte(updatedAtLte time.Time) TaskInstanceApiApiGetMappedTaskInstancesRequest {
	r.updatedAtLte = &updatedAtLte
	return r
}
// Returns objects greater than or equal to the specified values.  This can be combined with duration_lte parameter to receive only the selected period. 
func (r TaskInstanceApiApiGetMappedTaskInstancesRequest) DurationGte(durationGte float32) TaskInstanceApiApiGetMappedTaskInstancesRequest {
	r.durationGte = &durationGte
	return r
}
// Returns objects less than or equal to the specified values.  This can be combined with duration_gte parameter to receive only the selected range. 
func (r TaskInstanceApiApiGetMappedTaskInstancesRequest) DurationLte(durationLte float32) TaskInstanceApiApiGetMappedTaskInstancesRequest {
	r.durationLte = &durationLte
	return r
}
// The value can be repeated to retrieve multiple matching values (OR condition).
func (r TaskInstanceApiApiGetMappedTaskInstancesRequest) State(state []string) TaskInstanceApiApiGetMappedTaskInstancesRequest {
	r.state = &state
	return r
}
// The value can be repeated to retrieve multiple matching values (OR condition).
func (r TaskInstanceApiApiGetMappedTaskInstancesRequest) Pool(pool []string) TaskInstanceApiApiGetMappedTaskInstancesRequest {
	r.pool = &pool
	return r
}
// The value can be repeated to retrieve multiple matching values (OR condition).
func (r TaskInstanceApiApiGetMappedTaskInstancesRequest) Queue(queue []string) TaskInstanceApiApiGetMappedTaskInstancesRequest {
	r.queue = &queue
	return r
}
// The name of the field to order the results by. Prefix a field name with &#x60;-&#x60; to reverse the sort order.  *New in version 2.1.0* 
func (r TaskInstanceApiApiGetMappedTaskInstancesRequest) OrderBy(orderBy string) TaskInstanceApiApiGetMappedTaskInstancesRequest {
	r.orderBy = &orderBy
	return r
}

func (r TaskInstanceApiApiGetMappedTaskInstancesRequest) Execute() (*TaskInstanceCollection, *http.Response, error) {
	return r.ApiService.GetMappedTaskInstancesExecute(r)
}

/*
GetMappedTaskInstances List mapped task instances

Get details of all mapped task instances.

*New in version 2.3.0*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @return TaskInstanceApiApiGetMappedTaskInstancesRequest
*/
func (a *TaskInstanceApiService) GetMappedTaskInstances(ctx context.Context, dagId string, dagRunId string, taskId string) TaskInstanceApiApiGetMappedTaskInstancesRequest {
	return TaskInstanceApiApiGetMappedTaskInstancesRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return TaskInstanceCollection
func (a *TaskInstanceApiService) GetMappedTaskInstancesExecute(r TaskInstanceApiApiGetMappedTaskInstancesRequest) (*TaskInstanceCollection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstanceCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceApiService.GetMappedTaskInstances")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/listMapped"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterToString(r.dagId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterToString(r.dagRunId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterToString(r.taskId, "")), -1)

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
	if r.durationGte != nil {
		localVarQueryParams.Add("duration_gte", parameterToString(*r.durationGte, ""))
	}
	if r.durationLte != nil {
		localVarQueryParams.Add("duration_lte", parameterToString(*r.durationLte, ""))
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
	if r.pool != nil {
		t := *r.pool
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("pool", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("pool", parameterToString(t, "multi"))
		}
	}
	if r.queue != nil {
		t := *r.queue
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("queue", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("queue", parameterToString(t, "multi"))
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

type TaskInstanceApiApiGetTaskInstanceRequest struct {
	ctx context.Context
	ApiService *TaskInstanceApiService
	dagId string
	dagRunId string
	taskId string
}


func (r TaskInstanceApiApiGetTaskInstanceRequest) Execute() (*TaskInstance, *http.Response, error) {
	return r.ApiService.GetTaskInstanceExecute(r)
}

/*
GetTaskInstance Get a task instance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @return TaskInstanceApiApiGetTaskInstanceRequest
*/
func (a *TaskInstanceApiService) GetTaskInstance(ctx context.Context, dagId string, dagRunId string, taskId string) TaskInstanceApiApiGetTaskInstanceRequest {
	return TaskInstanceApiApiGetTaskInstanceRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return TaskInstance
func (a *TaskInstanceApiService) GetTaskInstanceExecute(r TaskInstanceApiApiGetTaskInstanceRequest) (*TaskInstance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceApiService.GetTaskInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterToString(r.dagId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterToString(r.dagRunId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterToString(r.taskId, "")), -1)

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

type TaskInstanceApiApiGetTaskInstancesRequest struct {
	ctx context.Context
	ApiService *TaskInstanceApiService
	dagId string
	dagRunId string
	executionDateGte *time.Time
	executionDateLte *time.Time
	startDateGte *time.Time
	startDateLte *time.Time
	endDateGte *time.Time
	endDateLte *time.Time
	updatedAtGte *time.Time
	updatedAtLte *time.Time
	durationGte *float32
	durationLte *float32
	state *[]string
	pool *[]string
	queue *[]string
	limit *int32
	offset *int32
}

// Returns objects greater or equal to the specified date.  This can be combined with execution_date_lte parameter to receive only the selected period. 
func (r TaskInstanceApiApiGetTaskInstancesRequest) ExecutionDateGte(executionDateGte time.Time) TaskInstanceApiApiGetTaskInstancesRequest {
	r.executionDateGte = &executionDateGte
	return r
}
// Returns objects less than or equal to the specified date.  This can be combined with execution_date_gte parameter to receive only the selected period. 
func (r TaskInstanceApiApiGetTaskInstancesRequest) ExecutionDateLte(executionDateLte time.Time) TaskInstanceApiApiGetTaskInstancesRequest {
	r.executionDateLte = &executionDateLte
	return r
}
// Returns objects greater or equal the specified date.  This can be combined with start_date_lte parameter to receive only the selected period. 
func (r TaskInstanceApiApiGetTaskInstancesRequest) StartDateGte(startDateGte time.Time) TaskInstanceApiApiGetTaskInstancesRequest {
	r.startDateGte = &startDateGte
	return r
}
// Returns objects less or equal the specified date.  This can be combined with start_date_gte parameter to receive only the selected period. 
func (r TaskInstanceApiApiGetTaskInstancesRequest) StartDateLte(startDateLte time.Time) TaskInstanceApiApiGetTaskInstancesRequest {
	r.startDateLte = &startDateLte
	return r
}
// Returns objects greater or equal the specified date.  This can be combined with start_date_lte parameter to receive only the selected period. 
func (r TaskInstanceApiApiGetTaskInstancesRequest) EndDateGte(endDateGte time.Time) TaskInstanceApiApiGetTaskInstancesRequest {
	r.endDateGte = &endDateGte
	return r
}
// Returns objects less than or equal to the specified date.  This can be combined with start_date_gte parameter to receive only the selected period. 
func (r TaskInstanceApiApiGetTaskInstancesRequest) EndDateLte(endDateLte time.Time) TaskInstanceApiApiGetTaskInstancesRequest {
	r.endDateLte = &endDateLte
	return r
}
// Returns objects greater or equal the specified date.  This can be combined with updated_at_lte parameter to receive only the selected period.  *New in version 2.6.0* 
func (r TaskInstanceApiApiGetTaskInstancesRequest) UpdatedAtGte(updatedAtGte time.Time) TaskInstanceApiApiGetTaskInstancesRequest {
	r.updatedAtGte = &updatedAtGte
	return r
}
// Returns objects less or equal the specified date.  This can be combined with updated_at_gte parameter to receive only the selected period.  *New in version 2.6.0* 
func (r TaskInstanceApiApiGetTaskInstancesRequest) UpdatedAtLte(updatedAtLte time.Time) TaskInstanceApiApiGetTaskInstancesRequest {
	r.updatedAtLte = &updatedAtLte
	return r
}
// Returns objects greater than or equal to the specified values.  This can be combined with duration_lte parameter to receive only the selected period. 
func (r TaskInstanceApiApiGetTaskInstancesRequest) DurationGte(durationGte float32) TaskInstanceApiApiGetTaskInstancesRequest {
	r.durationGte = &durationGte
	return r
}
// Returns objects less than or equal to the specified values.  This can be combined with duration_gte parameter to receive only the selected range. 
func (r TaskInstanceApiApiGetTaskInstancesRequest) DurationLte(durationLte float32) TaskInstanceApiApiGetTaskInstancesRequest {
	r.durationLte = &durationLte
	return r
}
// The value can be repeated to retrieve multiple matching values (OR condition).
func (r TaskInstanceApiApiGetTaskInstancesRequest) State(state []string) TaskInstanceApiApiGetTaskInstancesRequest {
	r.state = &state
	return r
}
// The value can be repeated to retrieve multiple matching values (OR condition).
func (r TaskInstanceApiApiGetTaskInstancesRequest) Pool(pool []string) TaskInstanceApiApiGetTaskInstancesRequest {
	r.pool = &pool
	return r
}
// The value can be repeated to retrieve multiple matching values (OR condition).
func (r TaskInstanceApiApiGetTaskInstancesRequest) Queue(queue []string) TaskInstanceApiApiGetTaskInstancesRequest {
	r.queue = &queue
	return r
}
// The numbers of items to return.
func (r TaskInstanceApiApiGetTaskInstancesRequest) Limit(limit int32) TaskInstanceApiApiGetTaskInstancesRequest {
	r.limit = &limit
	return r
}
// The number of items to skip before starting to collect the result set.
func (r TaskInstanceApiApiGetTaskInstancesRequest) Offset(offset int32) TaskInstanceApiApiGetTaskInstancesRequest {
	r.offset = &offset
	return r
}

func (r TaskInstanceApiApiGetTaskInstancesRequest) Execute() (*TaskInstanceCollection, *http.Response, error) {
	return r.ApiService.GetTaskInstancesExecute(r)
}

/*
GetTaskInstances List task instances

This endpoint allows specifying `~` as the dag_id, dag_run_id to retrieve DAG runs for all DAGs and DAG runs.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @return TaskInstanceApiApiGetTaskInstancesRequest
*/
func (a *TaskInstanceApiService) GetTaskInstances(ctx context.Context, dagId string, dagRunId string) TaskInstanceApiApiGetTaskInstancesRequest {
	return TaskInstanceApiApiGetTaskInstancesRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
	}
}

// Execute executes the request
//  @return TaskInstanceCollection
func (a *TaskInstanceApiService) GetTaskInstancesExecute(r TaskInstanceApiApiGetTaskInstancesRequest) (*TaskInstanceCollection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstanceCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceApiService.GetTaskInstances")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterToString(r.dagId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterToString(r.dagRunId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.durationGte != nil {
		localVarQueryParams.Add("duration_gte", parameterToString(*r.durationGte, ""))
	}
	if r.durationLte != nil {
		localVarQueryParams.Add("duration_lte", parameterToString(*r.durationLte, ""))
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
	if r.pool != nil {
		t := *r.pool
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("pool", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("pool", parameterToString(t, "multi"))
		}
	}
	if r.queue != nil {
		t := *r.queue
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("queue", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("queue", parameterToString(t, "multi"))
		}
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
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

type TaskInstanceApiApiGetTaskInstancesBatchRequest struct {
	ctx context.Context
	ApiService *TaskInstanceApiService
	listTaskInstanceForm *ListTaskInstanceForm
}

func (r TaskInstanceApiApiGetTaskInstancesBatchRequest) ListTaskInstanceForm(listTaskInstanceForm ListTaskInstanceForm) TaskInstanceApiApiGetTaskInstancesBatchRequest {
	r.listTaskInstanceForm = &listTaskInstanceForm
	return r
}

func (r TaskInstanceApiApiGetTaskInstancesBatchRequest) Execute() (*TaskInstanceCollection, *http.Response, error) {
	return r.ApiService.GetTaskInstancesBatchExecute(r)
}

/*
GetTaskInstancesBatch List task instances (batch)

List task instances from all DAGs and DAG runs.
This endpoint is a POST to allow filtering across a large number of DAG IDs, where as a GET it would run in to maximum HTTP request URL length limits.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TaskInstanceApiApiGetTaskInstancesBatchRequest
*/
func (a *TaskInstanceApiService) GetTaskInstancesBatch(ctx context.Context) TaskInstanceApiApiGetTaskInstancesBatchRequest {
	return TaskInstanceApiApiGetTaskInstancesBatchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TaskInstanceCollection
func (a *TaskInstanceApiService) GetTaskInstancesBatchExecute(r TaskInstanceApiApiGetTaskInstancesBatchRequest) (*TaskInstanceCollection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstanceCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceApiService.GetTaskInstancesBatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/~/dagRuns/~/taskInstances/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.listTaskInstanceForm == nil {
		return localVarReturnValue, nil, reportError("listTaskInstanceForm is required and must be specified")
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
	localVarPostBody = r.listTaskInstanceForm
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

type TaskInstanceApiApiPatchMappedTaskInstanceRequest struct {
	ctx context.Context
	ApiService *TaskInstanceApiService
	dagId string
	dagRunId string
	taskId string
	mapIndex int32
	updateTaskInstance *UpdateTaskInstance
}

// Parameters of action
func (r TaskInstanceApiApiPatchMappedTaskInstanceRequest) UpdateTaskInstance(updateTaskInstance UpdateTaskInstance) TaskInstanceApiApiPatchMappedTaskInstanceRequest {
	r.updateTaskInstance = &updateTaskInstance
	return r
}

func (r TaskInstanceApiApiPatchMappedTaskInstanceRequest) Execute() (*TaskInstanceReference, *http.Response, error) {
	return r.ApiService.PatchMappedTaskInstanceExecute(r)
}

/*
PatchMappedTaskInstance Updates the state of a mapped task instance

Updates the state for single mapped task instance.
*New in version 2.5.0*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @param mapIndex The map index.
 @return TaskInstanceApiApiPatchMappedTaskInstanceRequest
*/
func (a *TaskInstanceApiService) PatchMappedTaskInstance(ctx context.Context, dagId string, dagRunId string, taskId string, mapIndex int32) TaskInstanceApiApiPatchMappedTaskInstanceRequest {
	return TaskInstanceApiApiPatchMappedTaskInstanceRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
		mapIndex: mapIndex,
	}
}

// Execute executes the request
//  @return TaskInstanceReference
func (a *TaskInstanceApiService) PatchMappedTaskInstanceExecute(r TaskInstanceApiApiPatchMappedTaskInstanceRequest) (*TaskInstanceReference, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstanceReference
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceApiService.PatchMappedTaskInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/{map_index}"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterToString(r.dagId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterToString(r.dagRunId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterToString(r.taskId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"map_index"+"}", url.PathEscape(parameterToString(r.mapIndex, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.updateTaskInstance
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

type TaskInstanceApiApiPatchTaskInstanceRequest struct {
	ctx context.Context
	ApiService *TaskInstanceApiService
	dagId string
	dagRunId string
	taskId string
	updateTaskInstance *UpdateTaskInstance
}

// Parameters of action
func (r TaskInstanceApiApiPatchTaskInstanceRequest) UpdateTaskInstance(updateTaskInstance UpdateTaskInstance) TaskInstanceApiApiPatchTaskInstanceRequest {
	r.updateTaskInstance = &updateTaskInstance
	return r
}

func (r TaskInstanceApiApiPatchTaskInstanceRequest) Execute() (*TaskInstanceReference, *http.Response, error) {
	return r.ApiService.PatchTaskInstanceExecute(r)
}

/*
PatchTaskInstance Updates the state of a task instance

Updates the state for single task instance.
*New in version 2.5.0*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @return TaskInstanceApiApiPatchTaskInstanceRequest
*/
func (a *TaskInstanceApiService) PatchTaskInstance(ctx context.Context, dagId string, dagRunId string, taskId string) TaskInstanceApiApiPatchTaskInstanceRequest {
	return TaskInstanceApiApiPatchTaskInstanceRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return TaskInstanceReference
func (a *TaskInstanceApiService) PatchTaskInstanceExecute(r TaskInstanceApiApiPatchTaskInstanceRequest) (*TaskInstanceReference, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstanceReference
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceApiService.PatchTaskInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterToString(r.dagId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterToString(r.dagRunId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterToString(r.taskId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateTaskInstance == nil {
		return localVarReturnValue, nil, reportError("updateTaskInstance is required and must be specified")
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
	localVarPostBody = r.updateTaskInstance
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

type TaskInstanceApiApiSetMappedTaskInstanceNoteRequest struct {
	ctx context.Context
	ApiService *TaskInstanceApiService
	dagId string
	dagRunId string
	taskId string
	mapIndex int32
	setTaskInstanceNote *SetTaskInstanceNote
}

// Parameters of set Task Instance note.
func (r TaskInstanceApiApiSetMappedTaskInstanceNoteRequest) SetTaskInstanceNote(setTaskInstanceNote SetTaskInstanceNote) TaskInstanceApiApiSetMappedTaskInstanceNoteRequest {
	r.setTaskInstanceNote = &setTaskInstanceNote
	return r
}

func (r TaskInstanceApiApiSetMappedTaskInstanceNoteRequest) Execute() (*TaskInstance, *http.Response, error) {
	return r.ApiService.SetMappedTaskInstanceNoteExecute(r)
}

/*
SetMappedTaskInstanceNote Update the TaskInstance note.

Update the manual user note of a mapped Task Instance.

*New in version 2.5.0*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @param mapIndex The map index.
 @return TaskInstanceApiApiSetMappedTaskInstanceNoteRequest
*/
func (a *TaskInstanceApiService) SetMappedTaskInstanceNote(ctx context.Context, dagId string, dagRunId string, taskId string, mapIndex int32) TaskInstanceApiApiSetMappedTaskInstanceNoteRequest {
	return TaskInstanceApiApiSetMappedTaskInstanceNoteRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
		mapIndex: mapIndex,
	}
}

// Execute executes the request
//  @return TaskInstance
func (a *TaskInstanceApiService) SetMappedTaskInstanceNoteExecute(r TaskInstanceApiApiSetMappedTaskInstanceNoteRequest) (*TaskInstance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceApiService.SetMappedTaskInstanceNote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/{map_index}/setNote"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterToString(r.dagId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterToString(r.dagRunId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterToString(r.taskId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"map_index"+"}", url.PathEscape(parameterToString(r.mapIndex, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.setTaskInstanceNote == nil {
		return localVarReturnValue, nil, reportError("setTaskInstanceNote is required and must be specified")
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
	localVarPostBody = r.setTaskInstanceNote
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

type TaskInstanceApiApiSetTaskInstanceNoteRequest struct {
	ctx context.Context
	ApiService *TaskInstanceApiService
	dagId string
	dagRunId string
	taskId string
	setTaskInstanceNote *SetTaskInstanceNote
}

// Parameters of set Task Instance note.
func (r TaskInstanceApiApiSetTaskInstanceNoteRequest) SetTaskInstanceNote(setTaskInstanceNote SetTaskInstanceNote) TaskInstanceApiApiSetTaskInstanceNoteRequest {
	r.setTaskInstanceNote = &setTaskInstanceNote
	return r
}

func (r TaskInstanceApiApiSetTaskInstanceNoteRequest) Execute() (*TaskInstance, *http.Response, error) {
	return r.ApiService.SetTaskInstanceNoteExecute(r)
}

/*
SetTaskInstanceNote Update the TaskInstance note.

Update the manual user note of a non-mapped Task Instance.

*New in version 2.5.0*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @return TaskInstanceApiApiSetTaskInstanceNoteRequest
*/
func (a *TaskInstanceApiService) SetTaskInstanceNote(ctx context.Context, dagId string, dagRunId string, taskId string) TaskInstanceApiApiSetTaskInstanceNoteRequest {
	return TaskInstanceApiApiSetTaskInstanceNoteRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return TaskInstance
func (a *TaskInstanceApiService) SetTaskInstanceNoteExecute(r TaskInstanceApiApiSetTaskInstanceNoteRequest) (*TaskInstance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceApiService.SetTaskInstanceNote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/setNote"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterToString(r.dagId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterToString(r.dagRunId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterToString(r.taskId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.setTaskInstanceNote == nil {
		return localVarReturnValue, nil, reportError("setTaskInstanceNote is required and must be specified")
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
	localVarPostBody = r.setTaskInstanceNote
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
