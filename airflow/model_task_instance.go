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
	"encoding/json"
)

// TaskInstance struct for TaskInstance
type TaskInstance struct {
	TaskId *string `json:"task_id,omitempty"`
	DagId *string `json:"dag_id,omitempty"`
	// The DagRun ID for this task instance  *New in version 2.3.0* 
	DagRunId *string `json:"dag_run_id,omitempty"`
	ExecutionDate *string `json:"execution_date,omitempty"`
	StartDate NullableString `json:"start_date,omitempty"`
	EndDate NullableString `json:"end_date,omitempty"`
	Duration NullableFloat32 `json:"duration,omitempty"`
	State NullableTaskState `json:"state,omitempty"`
	TryNumber *int32 `json:"try_number,omitempty"`
	MapIndex *int32 `json:"map_index,omitempty"`
	MaxTries *int32 `json:"max_tries,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	Unixname *string `json:"unixname,omitempty"`
	Pool *string `json:"pool,omitempty"`
	PoolSlots *int32 `json:"pool_slots,omitempty"`
	Queue NullableString `json:"queue,omitempty"`
	PriorityWeight NullableInt32 `json:"priority_weight,omitempty"`
	// *Changed in version 2.1.1*&#58; Field becomes nullable. 
	Operator NullableString `json:"operator,omitempty"`
	QueuedWhen NullableString `json:"queued_when,omitempty"`
	Pid NullableInt32 `json:"pid,omitempty"`
	ExecutorConfig *string `json:"executor_config,omitempty"`
	SlaMiss NullableSLAMiss `json:"sla_miss,omitempty"`
	// JSON object describing rendered fields.  *New in version 2.3.0* 
	RenderedFields map[string]interface{} `json:"rendered_fields,omitempty"`
	Trigger NullableTrigger `json:"trigger,omitempty"`
	TriggererJob NullableJob `json:"triggerer_job,omitempty"`
	// Contains manually entered notes by the user about the TaskInstance.  *New in version 2.5.0* 
	Note NullableString `json:"note,omitempty"`
}

// NewTaskInstance instantiates a new TaskInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskInstance() *TaskInstance {
	this := TaskInstance{}
	return &this
}

// NewTaskInstanceWithDefaults instantiates a new TaskInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskInstanceWithDefaults() *TaskInstance {
	this := TaskInstance{}
	return &this
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *TaskInstance) GetTaskId() string {
	if o == nil || o.TaskId == nil {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstance) GetTaskIdOk() (*string, bool) {
	if o == nil || o.TaskId == nil {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *TaskInstance) HasTaskId() bool {
	if o != nil && o.TaskId != nil {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *TaskInstance) SetTaskId(v string) {
	o.TaskId = &v
}

// GetDagId returns the DagId field value if set, zero value otherwise.
func (o *TaskInstance) GetDagId() string {
	if o == nil || o.DagId == nil {
		var ret string
		return ret
	}
	return *o.DagId
}

// GetDagIdOk returns a tuple with the DagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstance) GetDagIdOk() (*string, bool) {
	if o == nil || o.DagId == nil {
		return nil, false
	}
	return o.DagId, true
}

// HasDagId returns a boolean if a field has been set.
func (o *TaskInstance) HasDagId() bool {
	if o != nil && o.DagId != nil {
		return true
	}

	return false
}

// SetDagId gets a reference to the given string and assigns it to the DagId field.
func (o *TaskInstance) SetDagId(v string) {
	o.DagId = &v
}

// GetDagRunId returns the DagRunId field value if set, zero value otherwise.
func (o *TaskInstance) GetDagRunId() string {
	if o == nil || o.DagRunId == nil {
		var ret string
		return ret
	}
	return *o.DagRunId
}

// GetDagRunIdOk returns a tuple with the DagRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstance) GetDagRunIdOk() (*string, bool) {
	if o == nil || o.DagRunId == nil {
		return nil, false
	}
	return o.DagRunId, true
}

// HasDagRunId returns a boolean if a field has been set.
func (o *TaskInstance) HasDagRunId() bool {
	if o != nil && o.DagRunId != nil {
		return true
	}

	return false
}

// SetDagRunId gets a reference to the given string and assigns it to the DagRunId field.
func (o *TaskInstance) SetDagRunId(v string) {
	o.DagRunId = &v
}

// GetExecutionDate returns the ExecutionDate field value if set, zero value otherwise.
func (o *TaskInstance) GetExecutionDate() string {
	if o == nil || o.ExecutionDate == nil {
		var ret string
		return ret
	}
	return *o.ExecutionDate
}

// GetExecutionDateOk returns a tuple with the ExecutionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstance) GetExecutionDateOk() (*string, bool) {
	if o == nil || o.ExecutionDate == nil {
		return nil, false
	}
	return o.ExecutionDate, true
}

// HasExecutionDate returns a boolean if a field has been set.
func (o *TaskInstance) HasExecutionDate() bool {
	if o != nil && o.ExecutionDate != nil {
		return true
	}

	return false
}

// SetExecutionDate gets a reference to the given string and assigns it to the ExecutionDate field.
func (o *TaskInstance) SetExecutionDate(v string) {
	o.ExecutionDate = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInstance) GetStartDate() string {
	if o == nil || o.StartDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInstance) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *TaskInstance) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableString and assigns it to the StartDate field.
func (o *TaskInstance) SetStartDate(v string) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *TaskInstance) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *TaskInstance) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInstance) GetEndDate() string {
	if o == nil || o.EndDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInstance) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *TaskInstance) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableString and assigns it to the EndDate field.
func (o *TaskInstance) SetEndDate(v string) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *TaskInstance) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *TaskInstance) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInstance) GetDuration() float32 {
	if o == nil || o.Duration.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInstance) GetDurationOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *TaskInstance) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableFloat32 and assigns it to the Duration field.
func (o *TaskInstance) SetDuration(v float32) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *TaskInstance) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *TaskInstance) UnsetDuration() {
	o.Duration.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInstance) GetState() TaskState {
	if o == nil || o.State.Get() == nil {
		var ret TaskState
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInstance) GetStateOk() (*TaskState, bool) {
	if o == nil  {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *TaskInstance) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableTaskState and assigns it to the State field.
func (o *TaskInstance) SetState(v TaskState) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *TaskInstance) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *TaskInstance) UnsetState() {
	o.State.Unset()
}

// GetTryNumber returns the TryNumber field value if set, zero value otherwise.
func (o *TaskInstance) GetTryNumber() int32 {
	if o == nil || o.TryNumber == nil {
		var ret int32
		return ret
	}
	return *o.TryNumber
}

// GetTryNumberOk returns a tuple with the TryNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstance) GetTryNumberOk() (*int32, bool) {
	if o == nil || o.TryNumber == nil {
		return nil, false
	}
	return o.TryNumber, true
}

// HasTryNumber returns a boolean if a field has been set.
func (o *TaskInstance) HasTryNumber() bool {
	if o != nil && o.TryNumber != nil {
		return true
	}

	return false
}

// SetTryNumber gets a reference to the given int32 and assigns it to the TryNumber field.
func (o *TaskInstance) SetTryNumber(v int32) {
	o.TryNumber = &v
}

// GetMapIndex returns the MapIndex field value if set, zero value otherwise.
func (o *TaskInstance) GetMapIndex() int32 {
	if o == nil || o.MapIndex == nil {
		var ret int32
		return ret
	}
	return *o.MapIndex
}

// GetMapIndexOk returns a tuple with the MapIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstance) GetMapIndexOk() (*int32, bool) {
	if o == nil || o.MapIndex == nil {
		return nil, false
	}
	return o.MapIndex, true
}

// HasMapIndex returns a boolean if a field has been set.
func (o *TaskInstance) HasMapIndex() bool {
	if o != nil && o.MapIndex != nil {
		return true
	}

	return false
}

// SetMapIndex gets a reference to the given int32 and assigns it to the MapIndex field.
func (o *TaskInstance) SetMapIndex(v int32) {
	o.MapIndex = &v
}

// GetMaxTries returns the MaxTries field value if set, zero value otherwise.
func (o *TaskInstance) GetMaxTries() int32 {
	if o == nil || o.MaxTries == nil {
		var ret int32
		return ret
	}
	return *o.MaxTries
}

// GetMaxTriesOk returns a tuple with the MaxTries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstance) GetMaxTriesOk() (*int32, bool) {
	if o == nil || o.MaxTries == nil {
		return nil, false
	}
	return o.MaxTries, true
}

// HasMaxTries returns a boolean if a field has been set.
func (o *TaskInstance) HasMaxTries() bool {
	if o != nil && o.MaxTries != nil {
		return true
	}

	return false
}

// SetMaxTries gets a reference to the given int32 and assigns it to the MaxTries field.
func (o *TaskInstance) SetMaxTries(v int32) {
	o.MaxTries = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *TaskInstance) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstance) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *TaskInstance) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *TaskInstance) SetHostname(v string) {
	o.Hostname = &v
}

// GetUnixname returns the Unixname field value if set, zero value otherwise.
func (o *TaskInstance) GetUnixname() string {
	if o == nil || o.Unixname == nil {
		var ret string
		return ret
	}
	return *o.Unixname
}

// GetUnixnameOk returns a tuple with the Unixname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstance) GetUnixnameOk() (*string, bool) {
	if o == nil || o.Unixname == nil {
		return nil, false
	}
	return o.Unixname, true
}

// HasUnixname returns a boolean if a field has been set.
func (o *TaskInstance) HasUnixname() bool {
	if o != nil && o.Unixname != nil {
		return true
	}

	return false
}

// SetUnixname gets a reference to the given string and assigns it to the Unixname field.
func (o *TaskInstance) SetUnixname(v string) {
	o.Unixname = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *TaskInstance) GetPool() string {
	if o == nil || o.Pool == nil {
		var ret string
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstance) GetPoolOk() (*string, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *TaskInstance) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given string and assigns it to the Pool field.
func (o *TaskInstance) SetPool(v string) {
	o.Pool = &v
}

// GetPoolSlots returns the PoolSlots field value if set, zero value otherwise.
func (o *TaskInstance) GetPoolSlots() int32 {
	if o == nil || o.PoolSlots == nil {
		var ret int32
		return ret
	}
	return *o.PoolSlots
}

// GetPoolSlotsOk returns a tuple with the PoolSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstance) GetPoolSlotsOk() (*int32, bool) {
	if o == nil || o.PoolSlots == nil {
		return nil, false
	}
	return o.PoolSlots, true
}

// HasPoolSlots returns a boolean if a field has been set.
func (o *TaskInstance) HasPoolSlots() bool {
	if o != nil && o.PoolSlots != nil {
		return true
	}

	return false
}

// SetPoolSlots gets a reference to the given int32 and assigns it to the PoolSlots field.
func (o *TaskInstance) SetPoolSlots(v int32) {
	o.PoolSlots = &v
}

// GetQueue returns the Queue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInstance) GetQueue() string {
	if o == nil || o.Queue.Get() == nil {
		var ret string
		return ret
	}
	return *o.Queue.Get()
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInstance) GetQueueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Queue.Get(), o.Queue.IsSet()
}

// HasQueue returns a boolean if a field has been set.
func (o *TaskInstance) HasQueue() bool {
	if o != nil && o.Queue.IsSet() {
		return true
	}

	return false
}

// SetQueue gets a reference to the given NullableString and assigns it to the Queue field.
func (o *TaskInstance) SetQueue(v string) {
	o.Queue.Set(&v)
}
// SetQueueNil sets the value for Queue to be an explicit nil
func (o *TaskInstance) SetQueueNil() {
	o.Queue.Set(nil)
}

// UnsetQueue ensures that no value is present for Queue, not even an explicit nil
func (o *TaskInstance) UnsetQueue() {
	o.Queue.Unset()
}

// GetPriorityWeight returns the PriorityWeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInstance) GetPriorityWeight() int32 {
	if o == nil || o.PriorityWeight.Get() == nil {
		var ret int32
		return ret
	}
	return *o.PriorityWeight.Get()
}

// GetPriorityWeightOk returns a tuple with the PriorityWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInstance) GetPriorityWeightOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PriorityWeight.Get(), o.PriorityWeight.IsSet()
}

// HasPriorityWeight returns a boolean if a field has been set.
func (o *TaskInstance) HasPriorityWeight() bool {
	if o != nil && o.PriorityWeight.IsSet() {
		return true
	}

	return false
}

// SetPriorityWeight gets a reference to the given NullableInt32 and assigns it to the PriorityWeight field.
func (o *TaskInstance) SetPriorityWeight(v int32) {
	o.PriorityWeight.Set(&v)
}
// SetPriorityWeightNil sets the value for PriorityWeight to be an explicit nil
func (o *TaskInstance) SetPriorityWeightNil() {
	o.PriorityWeight.Set(nil)
}

// UnsetPriorityWeight ensures that no value is present for PriorityWeight, not even an explicit nil
func (o *TaskInstance) UnsetPriorityWeight() {
	o.PriorityWeight.Unset()
}

// GetOperator returns the Operator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInstance) GetOperator() string {
	if o == nil || o.Operator.Get() == nil {
		var ret string
		return ret
	}
	return *o.Operator.Get()
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInstance) GetOperatorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Operator.Get(), o.Operator.IsSet()
}

// HasOperator returns a boolean if a field has been set.
func (o *TaskInstance) HasOperator() bool {
	if o != nil && o.Operator.IsSet() {
		return true
	}

	return false
}

// SetOperator gets a reference to the given NullableString and assigns it to the Operator field.
func (o *TaskInstance) SetOperator(v string) {
	o.Operator.Set(&v)
}
// SetOperatorNil sets the value for Operator to be an explicit nil
func (o *TaskInstance) SetOperatorNil() {
	o.Operator.Set(nil)
}

// UnsetOperator ensures that no value is present for Operator, not even an explicit nil
func (o *TaskInstance) UnsetOperator() {
	o.Operator.Unset()
}

// GetQueuedWhen returns the QueuedWhen field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInstance) GetQueuedWhen() string {
	if o == nil || o.QueuedWhen.Get() == nil {
		var ret string
		return ret
	}
	return *o.QueuedWhen.Get()
}

// GetQueuedWhenOk returns a tuple with the QueuedWhen field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInstance) GetQueuedWhenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.QueuedWhen.Get(), o.QueuedWhen.IsSet()
}

// HasQueuedWhen returns a boolean if a field has been set.
func (o *TaskInstance) HasQueuedWhen() bool {
	if o != nil && o.QueuedWhen.IsSet() {
		return true
	}

	return false
}

// SetQueuedWhen gets a reference to the given NullableString and assigns it to the QueuedWhen field.
func (o *TaskInstance) SetQueuedWhen(v string) {
	o.QueuedWhen.Set(&v)
}
// SetQueuedWhenNil sets the value for QueuedWhen to be an explicit nil
func (o *TaskInstance) SetQueuedWhenNil() {
	o.QueuedWhen.Set(nil)
}

// UnsetQueuedWhen ensures that no value is present for QueuedWhen, not even an explicit nil
func (o *TaskInstance) UnsetQueuedWhen() {
	o.QueuedWhen.Unset()
}

// GetPid returns the Pid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInstance) GetPid() int32 {
	if o == nil || o.Pid.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Pid.Get()
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInstance) GetPidOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Pid.Get(), o.Pid.IsSet()
}

// HasPid returns a boolean if a field has been set.
func (o *TaskInstance) HasPid() bool {
	if o != nil && o.Pid.IsSet() {
		return true
	}

	return false
}

// SetPid gets a reference to the given NullableInt32 and assigns it to the Pid field.
func (o *TaskInstance) SetPid(v int32) {
	o.Pid.Set(&v)
}
// SetPidNil sets the value for Pid to be an explicit nil
func (o *TaskInstance) SetPidNil() {
	o.Pid.Set(nil)
}

// UnsetPid ensures that no value is present for Pid, not even an explicit nil
func (o *TaskInstance) UnsetPid() {
	o.Pid.Unset()
}

// GetExecutorConfig returns the ExecutorConfig field value if set, zero value otherwise.
func (o *TaskInstance) GetExecutorConfig() string {
	if o == nil || o.ExecutorConfig == nil {
		var ret string
		return ret
	}
	return *o.ExecutorConfig
}

// GetExecutorConfigOk returns a tuple with the ExecutorConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstance) GetExecutorConfigOk() (*string, bool) {
	if o == nil || o.ExecutorConfig == nil {
		return nil, false
	}
	return o.ExecutorConfig, true
}

// HasExecutorConfig returns a boolean if a field has been set.
func (o *TaskInstance) HasExecutorConfig() bool {
	if o != nil && o.ExecutorConfig != nil {
		return true
	}

	return false
}

// SetExecutorConfig gets a reference to the given string and assigns it to the ExecutorConfig field.
func (o *TaskInstance) SetExecutorConfig(v string) {
	o.ExecutorConfig = &v
}

// GetSlaMiss returns the SlaMiss field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInstance) GetSlaMiss() SLAMiss {
	if o == nil || o.SlaMiss.Get() == nil {
		var ret SLAMiss
		return ret
	}
	return *o.SlaMiss.Get()
}

// GetSlaMissOk returns a tuple with the SlaMiss field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInstance) GetSlaMissOk() (*SLAMiss, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SlaMiss.Get(), o.SlaMiss.IsSet()
}

// HasSlaMiss returns a boolean if a field has been set.
func (o *TaskInstance) HasSlaMiss() bool {
	if o != nil && o.SlaMiss.IsSet() {
		return true
	}

	return false
}

// SetSlaMiss gets a reference to the given NullableSLAMiss and assigns it to the SlaMiss field.
func (o *TaskInstance) SetSlaMiss(v SLAMiss) {
	o.SlaMiss.Set(&v)
}
// SetSlaMissNil sets the value for SlaMiss to be an explicit nil
func (o *TaskInstance) SetSlaMissNil() {
	o.SlaMiss.Set(nil)
}

// UnsetSlaMiss ensures that no value is present for SlaMiss, not even an explicit nil
func (o *TaskInstance) UnsetSlaMiss() {
	o.SlaMiss.Unset()
}

// GetRenderedFields returns the RenderedFields field value if set, zero value otherwise.
func (o *TaskInstance) GetRenderedFields() map[string]interface{} {
	if o == nil || o.RenderedFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.RenderedFields
}

// GetRenderedFieldsOk returns a tuple with the RenderedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstance) GetRenderedFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.RenderedFields == nil {
		return nil, false
	}
	return o.RenderedFields, true
}

// HasRenderedFields returns a boolean if a field has been set.
func (o *TaskInstance) HasRenderedFields() bool {
	if o != nil && o.RenderedFields != nil {
		return true
	}

	return false
}

// SetRenderedFields gets a reference to the given map[string]interface{} and assigns it to the RenderedFields field.
func (o *TaskInstance) SetRenderedFields(v map[string]interface{}) {
	o.RenderedFields = v
}

// GetTrigger returns the Trigger field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInstance) GetTrigger() Trigger {
	if o == nil || o.Trigger.Get() == nil {
		var ret Trigger
		return ret
	}
	return *o.Trigger.Get()
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInstance) GetTriggerOk() (*Trigger, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Trigger.Get(), o.Trigger.IsSet()
}

// HasTrigger returns a boolean if a field has been set.
func (o *TaskInstance) HasTrigger() bool {
	if o != nil && o.Trigger.IsSet() {
		return true
	}

	return false
}

// SetTrigger gets a reference to the given NullableTrigger and assigns it to the Trigger field.
func (o *TaskInstance) SetTrigger(v Trigger) {
	o.Trigger.Set(&v)
}
// SetTriggerNil sets the value for Trigger to be an explicit nil
func (o *TaskInstance) SetTriggerNil() {
	o.Trigger.Set(nil)
}

// UnsetTrigger ensures that no value is present for Trigger, not even an explicit nil
func (o *TaskInstance) UnsetTrigger() {
	o.Trigger.Unset()
}

// GetTriggererJob returns the TriggererJob field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInstance) GetTriggererJob() Job {
	if o == nil || o.TriggererJob.Get() == nil {
		var ret Job
		return ret
	}
	return *o.TriggererJob.Get()
}

// GetTriggererJobOk returns a tuple with the TriggererJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInstance) GetTriggererJobOk() (*Job, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TriggererJob.Get(), o.TriggererJob.IsSet()
}

// HasTriggererJob returns a boolean if a field has been set.
func (o *TaskInstance) HasTriggererJob() bool {
	if o != nil && o.TriggererJob.IsSet() {
		return true
	}

	return false
}

// SetTriggererJob gets a reference to the given NullableJob and assigns it to the TriggererJob field.
func (o *TaskInstance) SetTriggererJob(v Job) {
	o.TriggererJob.Set(&v)
}
// SetTriggererJobNil sets the value for TriggererJob to be an explicit nil
func (o *TaskInstance) SetTriggererJobNil() {
	o.TriggererJob.Set(nil)
}

// UnsetTriggererJob ensures that no value is present for TriggererJob, not even an explicit nil
func (o *TaskInstance) UnsetTriggererJob() {
	o.TriggererJob.Unset()
}

// GetNote returns the Note field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInstance) GetNote() string {
	if o == nil || o.Note.Get() == nil {
		var ret string
		return ret
	}
	return *o.Note.Get()
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInstance) GetNoteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Note.Get(), o.Note.IsSet()
}

// HasNote returns a boolean if a field has been set.
func (o *TaskInstance) HasNote() bool {
	if o != nil && o.Note.IsSet() {
		return true
	}

	return false
}

// SetNote gets a reference to the given NullableString and assigns it to the Note field.
func (o *TaskInstance) SetNote(v string) {
	o.Note.Set(&v)
}
// SetNoteNil sets the value for Note to be an explicit nil
func (o *TaskInstance) SetNoteNil() {
	o.Note.Set(nil)
}

// UnsetNote ensures that no value is present for Note, not even an explicit nil
func (o *TaskInstance) UnsetNote() {
	o.Note.Unset()
}

func (o TaskInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TaskId != nil {
		toSerialize["task_id"] = o.TaskId
	}
	if o.DagId != nil {
		toSerialize["dag_id"] = o.DagId
	}
	if o.DagRunId != nil {
		toSerialize["dag_run_id"] = o.DagRunId
	}
	if o.ExecutionDate != nil {
		toSerialize["execution_date"] = o.ExecutionDate
	}
	if o.StartDate.IsSet() {
		toSerialize["start_date"] = o.StartDate.Get()
	}
	if o.EndDate.IsSet() {
		toSerialize["end_date"] = o.EndDate.Get()
	}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.TryNumber != nil {
		toSerialize["try_number"] = o.TryNumber
	}
	if o.MapIndex != nil {
		toSerialize["map_index"] = o.MapIndex
	}
	if o.MaxTries != nil {
		toSerialize["max_tries"] = o.MaxTries
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Unixname != nil {
		toSerialize["unixname"] = o.Unixname
	}
	if o.Pool != nil {
		toSerialize["pool"] = o.Pool
	}
	if o.PoolSlots != nil {
		toSerialize["pool_slots"] = o.PoolSlots
	}
	if o.Queue.IsSet() {
		toSerialize["queue"] = o.Queue.Get()
	}
	if o.PriorityWeight.IsSet() {
		toSerialize["priority_weight"] = o.PriorityWeight.Get()
	}
	if o.Operator.IsSet() {
		toSerialize["operator"] = o.Operator.Get()
	}
	if o.QueuedWhen.IsSet() {
		toSerialize["queued_when"] = o.QueuedWhen.Get()
	}
	if o.Pid.IsSet() {
		toSerialize["pid"] = o.Pid.Get()
	}
	if o.ExecutorConfig != nil {
		toSerialize["executor_config"] = o.ExecutorConfig
	}
	if o.SlaMiss.IsSet() {
		toSerialize["sla_miss"] = o.SlaMiss.Get()
	}
	if o.RenderedFields != nil {
		toSerialize["rendered_fields"] = o.RenderedFields
	}
	if o.Trigger.IsSet() {
		toSerialize["trigger"] = o.Trigger.Get()
	}
	if o.TriggererJob.IsSet() {
		toSerialize["triggerer_job"] = o.TriggererJob.Get()
	}
	if o.Note.IsSet() {
		toSerialize["note"] = o.Note.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTaskInstance struct {
	value *TaskInstance
	isSet bool
}

func (v NullableTaskInstance) Get() *TaskInstance {
	return v.value
}

func (v *NullableTaskInstance) Set(val *TaskInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskInstance(val *TaskInstance) *NullableTaskInstance {
	return &NullableTaskInstance{value: val, isSet: true}
}

func (v NullableTaskInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


