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
 * Airflow API (Stable)
 *
 * Apache Airflow management API.
 *
 * API version: 1.0.0
 * Contact: dev@airflow.apache.org
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package airflow

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"github.com/antihax/optional"
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

// TaskInstanceApiService TaskInstanceApi service
type TaskInstanceApiService service

/*
GetExtraLinks Get extra links for task instance
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dagId The DAG ID.
 * @param dagRunId The DAG Run ID.
 * @param taskId The Task ID.
@return ExtraLinkCollection
*/
func (a *TaskInstanceApiService) GetExtraLinks(ctx _context.Context, dagId string, dagRunId string, taskId string) (ExtraLinkCollection, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ExtraLinkCollection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/links"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", _neturl.QueryEscape(parameterToString(dagId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", _neturl.QueryEscape(parameterToString(dagRunId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", _neturl.QueryEscape(parameterToString(taskId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GetLogOpts Optional parameters for the method 'GetLog'
type GetLogOpts struct {
    FullContent optional.Bool
    Token optional.String
}

/*
GetLog Get logs for a task instance
Get logs for a specific task instance and its try number
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dagId The DAG ID.
 * @param dagRunId The DAG Run ID.
 * @param taskId The Task ID.
 * @param taskTryNumber The Task Try Number.
 * @param optional nil or *GetLogOpts - Optional Parameters:
 * @param "FullContent" (optional.Bool) -  A full content will be returned. By default, only the first fragment will be returned. 
 * @param "Token" (optional.String) -  A token that allows you to continue fetching logs. If passed, it will specify the location from which the download should be continued. 
@return InlineResponse200
*/
func (a *TaskInstanceApiService) GetLog(ctx _context.Context, dagId string, dagRunId string, taskId string, taskTryNumber int32, localVarOptionals *GetLogOpts) (InlineResponse200, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse200
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/logs/{task_try_number}"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", _neturl.QueryEscape(parameterToString(dagId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", _neturl.QueryEscape(parameterToString(dagRunId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", _neturl.QueryEscape(parameterToString(taskId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"task_try_number"+"}", _neturl.QueryEscape(parameterToString(taskTryNumber, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.FullContent.IsSet() {
		localVarQueryParams.Add("full_content", parameterToString(localVarOptionals.FullContent.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Token.IsSet() {
		localVarQueryParams.Add("token", parameterToString(localVarOptionals.Token.Value(), ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
GetTaskInstance Get a task instance
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dagId The DAG ID.
 * @param dagRunId The DAG Run ID.
 * @param taskId The Task ID.
@return TaskInstance
*/
func (a *TaskInstanceApiService) GetTaskInstance(ctx _context.Context, dagId string, dagRunId string, taskId string) (TaskInstance, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TaskInstance
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", _neturl.QueryEscape(parameterToString(dagId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", _neturl.QueryEscape(parameterToString(dagRunId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", _neturl.QueryEscape(parameterToString(taskId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GetTaskInstancesOpts Optional parameters for the method 'GetTaskInstances'
type GetTaskInstancesOpts struct {
    ExecutionDateGte optional.Time
    ExecutionDateLte optional.Time
    StartDateGte optional.Time
    StartDateLte optional.Time
    EndDateGte optional.Time
    EndDateLte optional.Time
    DurationGte optional.Float32
    DurationLte optional.Float32
    State optional.Interface
    Pool optional.Interface
    Queue optional.Interface
    Limit optional.Int32
    Offset optional.Int32
}

/*
GetTaskInstances Get a list of task instance of DAG
This endpoint allows specifying &#x60;~&#x60; as the dag_id, dag_run_id to retrieve DAG Runs for all DAGs and DAG Runs. 
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dagId The DAG ID.
 * @param dagRunId The DAG Run ID.
 * @param optional nil or *GetTaskInstancesOpts - Optional Parameters:
 * @param "ExecutionDateGte" (optional.Time) -  Returns objects greater or equal to the specified date. This can be combined with execution_date_lte parameter to receive only the selected period. 
 * @param "ExecutionDateLte" (optional.Time) -  Returns objects less than or equal to the specified date. This can be combined with execution_date_gte parameter to receive only the selected period. 
 * @param "StartDateGte" (optional.Time) -  Returns objects greater or equal the specified date. This can be combined with startd_ate_lte parameter to receive only the selected period. 
 * @param "StartDateLte" (optional.Time) -  Returns objects less or equal the specified date. This can be combined with start_date_gte parameter to receive only the selected period. 
 * @param "EndDateGte" (optional.Time) -  Returns objects greater or equal the specified date. This can be combined with start_date_lte parameter to receive only the selected period. 
 * @param "EndDateLte" (optional.Time) -  Returns objects less than or equal to the specified date. This can be combined with start_date_gte parameter to receive only the selected period. 
 * @param "DurationGte" (optional.Float32) -  Returns objects greater than or equal to the specified values. This can be combined with duration_lte parameter to receive only the selected period. 
 * @param "DurationLte" (optional.Float32) -  Returns objects less than or equal to the specified values. This can be combined with duration_gte parameter to receive only the selected range. 
 * @param "State" (optional.Interface of []string) -  The value can be repeated to retrieve multiple matching values (OR condition).
 * @param "Pool" (optional.Interface of []string) -  The value can be repeated to retrieve multiple matching values (OR condition).
 * @param "Queue" (optional.Interface of []string) -  The value can be repeated to retrieve multiple matching values (OR condition).
 * @param "Limit" (optional.Int32) -  The numbers of items to return.
 * @param "Offset" (optional.Int32) -  The number of items to skip before starting to collect the result set.
@return TaskInstanceCollection
*/
func (a *TaskInstanceApiService) GetTaskInstances(ctx _context.Context, dagId string, dagRunId string, localVarOptionals *GetTaskInstancesOpts) (TaskInstanceCollection, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TaskInstanceCollection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", _neturl.QueryEscape(parameterToString(dagId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", _neturl.QueryEscape(parameterToString(dagRunId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.ExecutionDateGte.IsSet() {
		localVarQueryParams.Add("execution_date_gte", parameterToString(localVarOptionals.ExecutionDateGte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ExecutionDateLte.IsSet() {
		localVarQueryParams.Add("execution_date_lte", parameterToString(localVarOptionals.ExecutionDateLte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartDateGte.IsSet() {
		localVarQueryParams.Add("start_date_gte", parameterToString(localVarOptionals.StartDateGte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartDateLte.IsSet() {
		localVarQueryParams.Add("start_date_lte", parameterToString(localVarOptionals.StartDateLte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndDateGte.IsSet() {
		localVarQueryParams.Add("end_date_gte", parameterToString(localVarOptionals.EndDateGte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndDateLte.IsSet() {
		localVarQueryParams.Add("end_date_lte", parameterToString(localVarOptionals.EndDateLte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DurationGte.IsSet() {
		localVarQueryParams.Add("duration_gte", parameterToString(localVarOptionals.DurationGte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DurationLte.IsSet() {
		localVarQueryParams.Add("duration_lte", parameterToString(localVarOptionals.DurationLte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.State.IsSet() {
		t:=localVarOptionals.State.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("state", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("state", parameterToString(t, "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.Pool.IsSet() {
		t:=localVarOptionals.Pool.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("pool", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("pool", parameterToString(t, "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.Queue.IsSet() {
		t:=localVarOptionals.Queue.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("queue", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("queue", parameterToString(t, "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
GetTaskInstancesBatch Get list of task instances from all DAGs and DAG Runs.
This endpoint is a POST to allow filtering across a large number of DAG IDs, where as a GET it would run in to maximum HTTP request URL length limits. 
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param listTaskInstanceForm
@return TaskInstanceCollection
*/
func (a *TaskInstanceApiService) GetTaskInstancesBatch(ctx _context.Context, listTaskInstanceForm ListTaskInstanceForm) (TaskInstanceCollection, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TaskInstanceCollection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/dags/~/dagRuns/~/taskInstances/list"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarPostBody = &listTaskInstanceForm
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
