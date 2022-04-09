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

# Overview  To facilitate management, Apache Airflow supports a range of REST API endpoints across its objects. This section provides an overview of the API design, methods, and supported use cases.  Most of the endpoints accept `JSON` as input and return `JSON` responses. This means that you must usually add the following headers to your request: ``` Content-type: application/json Accept: application/json ```  ## Resources  The term `resource` refers to a single type of object in the Airflow metadata. An API is broken up by its endpoint's corresponding resource. The name of a resource is typically plural and expressed in camelCase. Example: `dagRuns`.  Resource names are used as part of endpoint URLs, as well as in API parameters and responses.  ## CRUD Operations  The platform supports **C**reate, **R**ead, **U**pdate, and **D**elete operations on most resources. You can review the standards for these operations and their standard parameters below.  Some endpoints have special behavior as exceptions.  ### Create  To create a resource, you typically submit an HTTP `POST` request with the resource's required metadata in the request body. The response returns a `201 Created` response code upon success with the resource's metadata, including its internal `id`, in the response body.  ### Read  The HTTP `GET` request can be used to read a resource or to list a number of resources.  A resource's `id` can be submitted in the request parameters to read a specific resource. The response usually returns a `200 OK` response code upon success, with the resource's metadata in the response body.  If a `GET` request does not include a specific resource `id`, it is treated as a list request. The response usually returns a `200 OK` response code upon success, with an object containing a list of resources' metadata in the response body.  When reading resources, some common query parameters are usually available. e.g.: ``` v1/connections?limit=25&offset=25 ```  |Query Parameter|Type|Description| |---------------|----|-----------| |limit|integer|Maximum number of objects to fetch. Usually 25 by default| |offset|integer|Offset after which to start returning objects. For use with limit query parameter.|  ### Update  Updating a resource requires the resource `id`, and is typically done using an HTTP `PATCH` request, with the fields to modify in the request body. The response usually returns a `200 OK` response code upon success, with information about the modified resource in the response body.  ### Delete  Deleting a resource requires the resource `id` and is typically executing via an HTTP `DELETE` request. The response usually returns a `204 No Content` response code upon success.  ## Conventions  - Resource names are plural and expressed in camelCase. - Names are consistent between URL parameter name and field name.  - Field names are in snake_case. ```json {     \"name\": \"string\",     \"slots\": 0,     \"occupied_slots\": 0,     \"used_slots\": 0,     \"queued_slots\": 0,     \"open_slots\": 0 } ```  ### Update Mask  Update mask is available as a query parameter in patch endpoints. It is used to notify the API which fields you want to update. Using `update_mask` makes it easier to update objects by helping the server know which fields to update in an object instead of updating all fields. The update request ignores any fields that aren't specified in the field mask, leaving them with their current values.  Example: ```   resource = request.get('/resource/my-id').json()   resource['my_field'] = 'new-value'   request.patch('/resource/my-id?update_mask=my_field', data=json.dumps(resource)) ```  ## Versioning and Endpoint Lifecycle  - API versioning is not synchronized to specific releases of the Apache Airflow. - APIs are designed to be backward compatible. - Any changes to the API will first go through a deprecation phase.  # Trying the API  You can use a third party client, such as [curl](https://curl.haxx.se/), [HTTPie](https://httpie.org/), [Postman](https://www.postman.com/) or [the Insomnia rest client](https://insomnia.rest/) to test the Apache Airflow API.  Note that you will need to pass credentials data.  For e.g., here is how to pause a DAG with [curl](https://curl.haxx.se/), when basic authorization is used: ```bash curl -X PATCH 'https://example.com/api/v1/dags/{dag_id}?update_mask=is_paused' \\ -H 'Content-Type: application/json' \\ --user \"username:password\" \\ -d '{     \"is_paused\": true }' ```  Using a graphical tool such as [Postman](https://www.postman.com/) or [Insomnia](https://insomnia.rest/), it is possible to import the API specifications directly:  1. Download the API specification by clicking the **Download** button at top of this document 2. Import the JSON specification in the graphical tool of your choice.   - In *Postman*, you can click the **import** button at the top   - With *Insomnia*, you can just drag-and-drop the file on the UI  Note that with *Postman*, you can also generate code snippets by selecting a request and clicking on the **Code** button.  ## Enabling CORS  [Cross-origin resource sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) is a browser security feature that restricts HTTP requests that are initiated from scripts running in the browser.  For details on enabling/configuring CORS, see [Enabling CORS](https://airflow.apache.org/docs/apache-airflow/stable/security/api.html).  # Authentication  To be able to meet the requirements of many organizations, Airflow supports many authentication methods, and it is even possible to add your own method.  If you want to check which auth backend is currently set, you can use `airflow config get-value api auth_backends` command as in the example below. ```bash $ airflow config get-value api auth_backends airflow.api.auth.backend.basic_auth ``` The default is to deny all requests.  For details on configuring the authentication, see [API Authorization](https://airflow.apache.org/docs/apache-airflow/stable/security/api.html).  # Errors  We follow the error response format proposed in [RFC 7807](https://tools.ietf.org/html/rfc7807) also known as Problem Details for HTTP APIs. As with our normal API responses, your client must be prepared to gracefully handle additional members of the response.  ## Unauthenticated  This indicates that the request has not been applied because it lacks valid authentication credentials for the target resource. Please check that you have valid credentials.  ## PermissionDenied  This response means that the server understood the request but refuses to authorize it because it lacks sufficient rights to the resource. It happens when you do not have the necessary permission to execute the action you performed. You need to get the appropriate permissions in other to resolve this error.  ## BadRequest  This response means that the server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing). To resolve this, please ensure that your syntax is correct.  ## NotFound  This client error response indicates that the server cannot find the requested resource.  ## MethodNotAllowed  Indicates that the request method is known by the server but is not supported by the target resource.  ## NotAcceptable  The target resource does not have a current representation that would be acceptable to the user agent, according to the proactive negotiation header fields received in the request, and the server is unwilling to supply a default representation.  ## AlreadyExists  The request could not be completed due to a conflict with the current state of the target resource, e.g. the resource it tries to create already exists.  ## Unknown  This means that the server encountered an unexpected condition that prevented it from fulfilling the request. 

API version: 1.0.0
Contact: dev@airflow.apache.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package airflow

import (
	"encoding/json"
)

// UpdateTaskInstancesState struct for UpdateTaskInstancesState
type UpdateTaskInstancesState struct {
	// If set, don't actually run this operation. The response will contain a list of task instances planned to be affected, but won't be modified in any way. 
	DryRun *bool `json:"dry_run,omitempty"`
	// The task ID.
	TaskId *string `json:"task_id,omitempty"`
	// The execution date. Either set this or dag_run_id but not both.
	ExecutionDate *string `json:"execution_date,omitempty"`
	// The task instance's DAG run ID. Either set this or execution_date but not both.  *New in version 2.3.0* 
	DagRunId *string `json:"dag_run_id,omitempty"`
	// If set to true, upstream tasks are also affected.
	IncludeUpstream *bool `json:"include_upstream,omitempty"`
	// If set to true, downstream tasks are also affected.
	IncludeDownstream *bool `json:"include_downstream,omitempty"`
	// If set to True, also tasks from future DAG Runs are affected.
	IncludeFuture *bool `json:"include_future,omitempty"`
	// If set to True, also tasks from past DAG Runs are affected.
	IncludePast *bool `json:"include_past,omitempty"`
	// Expected new state.
	NewState *string `json:"new_state,omitempty"`
}

// NewUpdateTaskInstancesState instantiates a new UpdateTaskInstancesState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTaskInstancesState() *UpdateTaskInstancesState {
	this := UpdateTaskInstancesState{}
	var dryRun bool = true
	this.DryRun = &dryRun
	return &this
}

// NewUpdateTaskInstancesStateWithDefaults instantiates a new UpdateTaskInstancesState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTaskInstancesStateWithDefaults() *UpdateTaskInstancesState {
	this := UpdateTaskInstancesState{}
	var dryRun bool = true
	this.DryRun = &dryRun
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateTaskInstancesState) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTaskInstancesState) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateTaskInstancesState) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateTaskInstancesState) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *UpdateTaskInstancesState) GetTaskId() string {
	if o == nil || o.TaskId == nil {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTaskInstancesState) GetTaskIdOk() (*string, bool) {
	if o == nil || o.TaskId == nil {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *UpdateTaskInstancesState) HasTaskId() bool {
	if o != nil && o.TaskId != nil {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *UpdateTaskInstancesState) SetTaskId(v string) {
	o.TaskId = &v
}

// GetExecutionDate returns the ExecutionDate field value if set, zero value otherwise.
func (o *UpdateTaskInstancesState) GetExecutionDate() string {
	if o == nil || o.ExecutionDate == nil {
		var ret string
		return ret
	}
	return *o.ExecutionDate
}

// GetExecutionDateOk returns a tuple with the ExecutionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTaskInstancesState) GetExecutionDateOk() (*string, bool) {
	if o == nil || o.ExecutionDate == nil {
		return nil, false
	}
	return o.ExecutionDate, true
}

// HasExecutionDate returns a boolean if a field has been set.
func (o *UpdateTaskInstancesState) HasExecutionDate() bool {
	if o != nil && o.ExecutionDate != nil {
		return true
	}

	return false
}

// SetExecutionDate gets a reference to the given string and assigns it to the ExecutionDate field.
func (o *UpdateTaskInstancesState) SetExecutionDate(v string) {
	o.ExecutionDate = &v
}

// GetDagRunId returns the DagRunId field value if set, zero value otherwise.
func (o *UpdateTaskInstancesState) GetDagRunId() string {
	if o == nil || o.DagRunId == nil {
		var ret string
		return ret
	}
	return *o.DagRunId
}

// GetDagRunIdOk returns a tuple with the DagRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTaskInstancesState) GetDagRunIdOk() (*string, bool) {
	if o == nil || o.DagRunId == nil {
		return nil, false
	}
	return o.DagRunId, true
}

// HasDagRunId returns a boolean if a field has been set.
func (o *UpdateTaskInstancesState) HasDagRunId() bool {
	if o != nil && o.DagRunId != nil {
		return true
	}

	return false
}

// SetDagRunId gets a reference to the given string and assigns it to the DagRunId field.
func (o *UpdateTaskInstancesState) SetDagRunId(v string) {
	o.DagRunId = &v
}

// GetIncludeUpstream returns the IncludeUpstream field value if set, zero value otherwise.
func (o *UpdateTaskInstancesState) GetIncludeUpstream() bool {
	if o == nil || o.IncludeUpstream == nil {
		var ret bool
		return ret
	}
	return *o.IncludeUpstream
}

// GetIncludeUpstreamOk returns a tuple with the IncludeUpstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTaskInstancesState) GetIncludeUpstreamOk() (*bool, bool) {
	if o == nil || o.IncludeUpstream == nil {
		return nil, false
	}
	return o.IncludeUpstream, true
}

// HasIncludeUpstream returns a boolean if a field has been set.
func (o *UpdateTaskInstancesState) HasIncludeUpstream() bool {
	if o != nil && o.IncludeUpstream != nil {
		return true
	}

	return false
}

// SetIncludeUpstream gets a reference to the given bool and assigns it to the IncludeUpstream field.
func (o *UpdateTaskInstancesState) SetIncludeUpstream(v bool) {
	o.IncludeUpstream = &v
}

// GetIncludeDownstream returns the IncludeDownstream field value if set, zero value otherwise.
func (o *UpdateTaskInstancesState) GetIncludeDownstream() bool {
	if o == nil || o.IncludeDownstream == nil {
		var ret bool
		return ret
	}
	return *o.IncludeDownstream
}

// GetIncludeDownstreamOk returns a tuple with the IncludeDownstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTaskInstancesState) GetIncludeDownstreamOk() (*bool, bool) {
	if o == nil || o.IncludeDownstream == nil {
		return nil, false
	}
	return o.IncludeDownstream, true
}

// HasIncludeDownstream returns a boolean if a field has been set.
func (o *UpdateTaskInstancesState) HasIncludeDownstream() bool {
	if o != nil && o.IncludeDownstream != nil {
		return true
	}

	return false
}

// SetIncludeDownstream gets a reference to the given bool and assigns it to the IncludeDownstream field.
func (o *UpdateTaskInstancesState) SetIncludeDownstream(v bool) {
	o.IncludeDownstream = &v
}

// GetIncludeFuture returns the IncludeFuture field value if set, zero value otherwise.
func (o *UpdateTaskInstancesState) GetIncludeFuture() bool {
	if o == nil || o.IncludeFuture == nil {
		var ret bool
		return ret
	}
	return *o.IncludeFuture
}

// GetIncludeFutureOk returns a tuple with the IncludeFuture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTaskInstancesState) GetIncludeFutureOk() (*bool, bool) {
	if o == nil || o.IncludeFuture == nil {
		return nil, false
	}
	return o.IncludeFuture, true
}

// HasIncludeFuture returns a boolean if a field has been set.
func (o *UpdateTaskInstancesState) HasIncludeFuture() bool {
	if o != nil && o.IncludeFuture != nil {
		return true
	}

	return false
}

// SetIncludeFuture gets a reference to the given bool and assigns it to the IncludeFuture field.
func (o *UpdateTaskInstancesState) SetIncludeFuture(v bool) {
	o.IncludeFuture = &v
}

// GetIncludePast returns the IncludePast field value if set, zero value otherwise.
func (o *UpdateTaskInstancesState) GetIncludePast() bool {
	if o == nil || o.IncludePast == nil {
		var ret bool
		return ret
	}
	return *o.IncludePast
}

// GetIncludePastOk returns a tuple with the IncludePast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTaskInstancesState) GetIncludePastOk() (*bool, bool) {
	if o == nil || o.IncludePast == nil {
		return nil, false
	}
	return o.IncludePast, true
}

// HasIncludePast returns a boolean if a field has been set.
func (o *UpdateTaskInstancesState) HasIncludePast() bool {
	if o != nil && o.IncludePast != nil {
		return true
	}

	return false
}

// SetIncludePast gets a reference to the given bool and assigns it to the IncludePast field.
func (o *UpdateTaskInstancesState) SetIncludePast(v bool) {
	o.IncludePast = &v
}

// GetNewState returns the NewState field value if set, zero value otherwise.
func (o *UpdateTaskInstancesState) GetNewState() string {
	if o == nil || o.NewState == nil {
		var ret string
		return ret
	}
	return *o.NewState
}

// GetNewStateOk returns a tuple with the NewState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTaskInstancesState) GetNewStateOk() (*string, bool) {
	if o == nil || o.NewState == nil {
		return nil, false
	}
	return o.NewState, true
}

// HasNewState returns a boolean if a field has been set.
func (o *UpdateTaskInstancesState) HasNewState() bool {
	if o != nil && o.NewState != nil {
		return true
	}

	return false
}

// SetNewState gets a reference to the given string and assigns it to the NewState field.
func (o *UpdateTaskInstancesState) SetNewState(v string) {
	o.NewState = &v
}

func (o UpdateTaskInstancesState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["dry_run"] = o.DryRun
	}
	if o.TaskId != nil {
		toSerialize["task_id"] = o.TaskId
	}
	if o.ExecutionDate != nil {
		toSerialize["execution_date"] = o.ExecutionDate
	}
	if o.DagRunId != nil {
		toSerialize["dag_run_id"] = o.DagRunId
	}
	if o.IncludeUpstream != nil {
		toSerialize["include_upstream"] = o.IncludeUpstream
	}
	if o.IncludeDownstream != nil {
		toSerialize["include_downstream"] = o.IncludeDownstream
	}
	if o.IncludeFuture != nil {
		toSerialize["include_future"] = o.IncludeFuture
	}
	if o.IncludePast != nil {
		toSerialize["include_past"] = o.IncludePast
	}
	if o.NewState != nil {
		toSerialize["new_state"] = o.NewState
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateTaskInstancesState struct {
	value *UpdateTaskInstancesState
	isSet bool
}

func (v NullableUpdateTaskInstancesState) Get() *UpdateTaskInstancesState {
	return v.value
}

func (v *NullableUpdateTaskInstancesState) Set(val *UpdateTaskInstancesState) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTaskInstancesState) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTaskInstancesState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTaskInstancesState(val *UpdateTaskInstancesState) *NullableUpdateTaskInstancesState {
	return &NullableUpdateTaskInstancesState{value: val, isSet: true}
}

func (v NullableUpdateTaskInstancesState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTaskInstancesState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


