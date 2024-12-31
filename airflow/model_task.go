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
	"time"
)

// Task For details see: [airflow.models.baseoperator.BaseOperator](https://airflow.apache.org/docs/apache-airflow/stable/_api/airflow/models/baseoperator/index.html#airflow.models.baseoperator.BaseOperator) 
type Task struct {
	ClassRef *ClassReference `json:"class_ref,omitempty"`
	TaskId *string `json:"task_id,omitempty"`
	Owner *string `json:"owner,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate NullableTime `json:"end_date,omitempty"`
	TriggerRule *TriggerRule `json:"trigger_rule,omitempty"`
	ExtraLinks []TaskExtraLinks `json:"extra_links,omitempty"`
	DependsOnPast *bool `json:"depends_on_past,omitempty"`
	IsMapped *bool `json:"is_mapped,omitempty"`
	WaitForDownstream *bool `json:"wait_for_downstream,omitempty"`
	Retries *float32 `json:"retries,omitempty"`
	Queue NullableString `json:"queue,omitempty"`
	Pool *string `json:"pool,omitempty"`
	PoolSlots *float32 `json:"pool_slots,omitempty"`
	ExecutionTimeout NullableTimeDelta `json:"execution_timeout,omitempty"`
	RetryDelay NullableTimeDelta `json:"retry_delay,omitempty"`
	RetryExponentialBackoff *bool `json:"retry_exponential_backoff,omitempty"`
	PriorityWeight *float32 `json:"priority_weight,omitempty"`
	WeightRule *WeightRule `json:"weight_rule,omitempty"`
	// Color in hexadecimal notation.
	UiColor *string `json:"ui_color,omitempty"`
	// Color in hexadecimal notation.
	UiFgcolor *string `json:"ui_fgcolor,omitempty"`
	TemplateFields []string `json:"template_fields,omitempty"`
	SubDag *DAG `json:"sub_dag,omitempty"`
	DownstreamTaskIds []string `json:"downstream_task_ids,omitempty"`
}

// NewTask instantiates a new Task object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTask() *Task {
	this := Task{}
	return &this
}

// NewTaskWithDefaults instantiates a new Task object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskWithDefaults() *Task {
	this := Task{}
	return &this
}

// GetClassRef returns the ClassRef field value if set, zero value otherwise.
func (o *Task) GetClassRef() ClassReference {
	if o == nil || o.ClassRef == nil {
		var ret ClassReference
		return ret
	}
	return *o.ClassRef
}

// GetClassRefOk returns a tuple with the ClassRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetClassRefOk() (*ClassReference, bool) {
	if o == nil || o.ClassRef == nil {
		return nil, false
	}
	return o.ClassRef, true
}

// HasClassRef returns a boolean if a field has been set.
func (o *Task) HasClassRef() bool {
	if o != nil && o.ClassRef != nil {
		return true
	}

	return false
}

// SetClassRef gets a reference to the given ClassReference and assigns it to the ClassRef field.
func (o *Task) SetClassRef(v ClassReference) {
	o.ClassRef = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *Task) GetTaskId() string {
	if o == nil || o.TaskId == nil {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetTaskIdOk() (*string, bool) {
	if o == nil || o.TaskId == nil {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *Task) HasTaskId() bool {
	if o != nil && o.TaskId != nil {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *Task) SetTaskId(v string) {
	o.TaskId = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Task) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Task) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *Task) SetOwner(v string) {
	o.Owner = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *Task) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Task) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *Task) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetEndDate() time.Time {
	if o == nil || o.EndDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetEndDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *Task) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableTime and assigns it to the EndDate field.
func (o *Task) SetEndDate(v time.Time) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *Task) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *Task) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetTriggerRule returns the TriggerRule field value if set, zero value otherwise.
func (o *Task) GetTriggerRule() TriggerRule {
	if o == nil || o.TriggerRule == nil {
		var ret TriggerRule
		return ret
	}
	return *o.TriggerRule
}

// GetTriggerRuleOk returns a tuple with the TriggerRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetTriggerRuleOk() (*TriggerRule, bool) {
	if o == nil || o.TriggerRule == nil {
		return nil, false
	}
	return o.TriggerRule, true
}

// HasTriggerRule returns a boolean if a field has been set.
func (o *Task) HasTriggerRule() bool {
	if o != nil && o.TriggerRule != nil {
		return true
	}

	return false
}

// SetTriggerRule gets a reference to the given TriggerRule and assigns it to the TriggerRule field.
func (o *Task) SetTriggerRule(v TriggerRule) {
	o.TriggerRule = &v
}

// GetExtraLinks returns the ExtraLinks field value if set, zero value otherwise.
func (o *Task) GetExtraLinks() []TaskExtraLinks {
	if o == nil || o.ExtraLinks == nil {
		var ret []TaskExtraLinks
		return ret
	}
	return o.ExtraLinks
}

// GetExtraLinksOk returns a tuple with the ExtraLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetExtraLinksOk() ([]TaskExtraLinks, bool) {
	if o == nil || o.ExtraLinks == nil {
		return nil, false
	}
	return o.ExtraLinks, true
}

// HasExtraLinks returns a boolean if a field has been set.
func (o *Task) HasExtraLinks() bool {
	if o != nil && o.ExtraLinks != nil {
		return true
	}

	return false
}

// SetExtraLinks gets a reference to the given []TaskExtraLinks and assigns it to the ExtraLinks field.
func (o *Task) SetExtraLinks(v []TaskExtraLinks) {
	o.ExtraLinks = v
}

// GetDependsOnPast returns the DependsOnPast field value if set, zero value otherwise.
func (o *Task) GetDependsOnPast() bool {
	if o == nil || o.DependsOnPast == nil {
		var ret bool
		return ret
	}
	return *o.DependsOnPast
}

// GetDependsOnPastOk returns a tuple with the DependsOnPast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetDependsOnPastOk() (*bool, bool) {
	if o == nil || o.DependsOnPast == nil {
		return nil, false
	}
	return o.DependsOnPast, true
}

// HasDependsOnPast returns a boolean if a field has been set.
func (o *Task) HasDependsOnPast() bool {
	if o != nil && o.DependsOnPast != nil {
		return true
	}

	return false
}

// SetDependsOnPast gets a reference to the given bool and assigns it to the DependsOnPast field.
func (o *Task) SetDependsOnPast(v bool) {
	o.DependsOnPast = &v
}

// GetIsMapped returns the IsMapped field value if set, zero value otherwise.
func (o *Task) GetIsMapped() bool {
	if o == nil || o.IsMapped == nil {
		var ret bool
		return ret
	}
	return *o.IsMapped
}

// GetIsMappedOk returns a tuple with the IsMapped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetIsMappedOk() (*bool, bool) {
	if o == nil || o.IsMapped == nil {
		return nil, false
	}
	return o.IsMapped, true
}

// HasIsMapped returns a boolean if a field has been set.
func (o *Task) HasIsMapped() bool {
	if o != nil && o.IsMapped != nil {
		return true
	}

	return false
}

// SetIsMapped gets a reference to the given bool and assigns it to the IsMapped field.
func (o *Task) SetIsMapped(v bool) {
	o.IsMapped = &v
}

// GetWaitForDownstream returns the WaitForDownstream field value if set, zero value otherwise.
func (o *Task) GetWaitForDownstream() bool {
	if o == nil || o.WaitForDownstream == nil {
		var ret bool
		return ret
	}
	return *o.WaitForDownstream
}

// GetWaitForDownstreamOk returns a tuple with the WaitForDownstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetWaitForDownstreamOk() (*bool, bool) {
	if o == nil || o.WaitForDownstream == nil {
		return nil, false
	}
	return o.WaitForDownstream, true
}

// HasWaitForDownstream returns a boolean if a field has been set.
func (o *Task) HasWaitForDownstream() bool {
	if o != nil && o.WaitForDownstream != nil {
		return true
	}

	return false
}

// SetWaitForDownstream gets a reference to the given bool and assigns it to the WaitForDownstream field.
func (o *Task) SetWaitForDownstream(v bool) {
	o.WaitForDownstream = &v
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *Task) GetRetries() float32 {
	if o == nil || o.Retries == nil {
		var ret float32
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetRetriesOk() (*float32, bool) {
	if o == nil || o.Retries == nil {
		return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *Task) HasRetries() bool {
	if o != nil && o.Retries != nil {
		return true
	}

	return false
}

// SetRetries gets a reference to the given float32 and assigns it to the Retries field.
func (o *Task) SetRetries(v float32) {
	o.Retries = &v
}

// GetQueue returns the Queue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetQueue() string {
	if o == nil || o.Queue.Get() == nil {
		var ret string
		return ret
	}
	return *o.Queue.Get()
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetQueueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Queue.Get(), o.Queue.IsSet()
}

// HasQueue returns a boolean if a field has been set.
func (o *Task) HasQueue() bool {
	if o != nil && o.Queue.IsSet() {
		return true
	}

	return false
}

// SetQueue gets a reference to the given NullableString and assigns it to the Queue field.
func (o *Task) SetQueue(v string) {
	o.Queue.Set(&v)
}
// SetQueueNil sets the value for Queue to be an explicit nil
func (o *Task) SetQueueNil() {
	o.Queue.Set(nil)
}

// UnsetQueue ensures that no value is present for Queue, not even an explicit nil
func (o *Task) UnsetQueue() {
	o.Queue.Unset()
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *Task) GetPool() string {
	if o == nil || o.Pool == nil {
		var ret string
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetPoolOk() (*string, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *Task) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given string and assigns it to the Pool field.
func (o *Task) SetPool(v string) {
	o.Pool = &v
}

// GetPoolSlots returns the PoolSlots field value if set, zero value otherwise.
func (o *Task) GetPoolSlots() float32 {
	if o == nil || o.PoolSlots == nil {
		var ret float32
		return ret
	}
	return *o.PoolSlots
}

// GetPoolSlotsOk returns a tuple with the PoolSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetPoolSlotsOk() (*float32, bool) {
	if o == nil || o.PoolSlots == nil {
		return nil, false
	}
	return o.PoolSlots, true
}

// HasPoolSlots returns a boolean if a field has been set.
func (o *Task) HasPoolSlots() bool {
	if o != nil && o.PoolSlots != nil {
		return true
	}

	return false
}

// SetPoolSlots gets a reference to the given float32 and assigns it to the PoolSlots field.
func (o *Task) SetPoolSlots(v float32) {
	o.PoolSlots = &v
}

// GetExecutionTimeout returns the ExecutionTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetExecutionTimeout() TimeDelta {
	if o == nil || o.ExecutionTimeout.Get() == nil {
		var ret TimeDelta
		return ret
	}
	return *o.ExecutionTimeout.Get()
}

// GetExecutionTimeoutOk returns a tuple with the ExecutionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetExecutionTimeoutOk() (*TimeDelta, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExecutionTimeout.Get(), o.ExecutionTimeout.IsSet()
}

// HasExecutionTimeout returns a boolean if a field has been set.
func (o *Task) HasExecutionTimeout() bool {
	if o != nil && o.ExecutionTimeout.IsSet() {
		return true
	}

	return false
}

// SetExecutionTimeout gets a reference to the given NullableTimeDelta and assigns it to the ExecutionTimeout field.
func (o *Task) SetExecutionTimeout(v TimeDelta) {
	o.ExecutionTimeout.Set(&v)
}
// SetExecutionTimeoutNil sets the value for ExecutionTimeout to be an explicit nil
func (o *Task) SetExecutionTimeoutNil() {
	o.ExecutionTimeout.Set(nil)
}

// UnsetExecutionTimeout ensures that no value is present for ExecutionTimeout, not even an explicit nil
func (o *Task) UnsetExecutionTimeout() {
	o.ExecutionTimeout.Unset()
}

// GetRetryDelay returns the RetryDelay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetRetryDelay() TimeDelta {
	if o == nil || o.RetryDelay.Get() == nil {
		var ret TimeDelta
		return ret
	}
	return *o.RetryDelay.Get()
}

// GetRetryDelayOk returns a tuple with the RetryDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetRetryDelayOk() (*TimeDelta, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RetryDelay.Get(), o.RetryDelay.IsSet()
}

// HasRetryDelay returns a boolean if a field has been set.
func (o *Task) HasRetryDelay() bool {
	if o != nil && o.RetryDelay.IsSet() {
		return true
	}

	return false
}

// SetRetryDelay gets a reference to the given NullableTimeDelta and assigns it to the RetryDelay field.
func (o *Task) SetRetryDelay(v TimeDelta) {
	o.RetryDelay.Set(&v)
}
// SetRetryDelayNil sets the value for RetryDelay to be an explicit nil
func (o *Task) SetRetryDelayNil() {
	o.RetryDelay.Set(nil)
}

// UnsetRetryDelay ensures that no value is present for RetryDelay, not even an explicit nil
func (o *Task) UnsetRetryDelay() {
	o.RetryDelay.Unset()
}

// GetRetryExponentialBackoff returns the RetryExponentialBackoff field value if set, zero value otherwise.
func (o *Task) GetRetryExponentialBackoff() bool {
	if o == nil || o.RetryExponentialBackoff == nil {
		var ret bool
		return ret
	}
	return *o.RetryExponentialBackoff
}

// GetRetryExponentialBackoffOk returns a tuple with the RetryExponentialBackoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetRetryExponentialBackoffOk() (*bool, bool) {
	if o == nil || o.RetryExponentialBackoff == nil {
		return nil, false
	}
	return o.RetryExponentialBackoff, true
}

// HasRetryExponentialBackoff returns a boolean if a field has been set.
func (o *Task) HasRetryExponentialBackoff() bool {
	if o != nil && o.RetryExponentialBackoff != nil {
		return true
	}

	return false
}

// SetRetryExponentialBackoff gets a reference to the given bool and assigns it to the RetryExponentialBackoff field.
func (o *Task) SetRetryExponentialBackoff(v bool) {
	o.RetryExponentialBackoff = &v
}

// GetPriorityWeight returns the PriorityWeight field value if set, zero value otherwise.
func (o *Task) GetPriorityWeight() float32 {
	if o == nil || o.PriorityWeight == nil {
		var ret float32
		return ret
	}
	return *o.PriorityWeight
}

// GetPriorityWeightOk returns a tuple with the PriorityWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetPriorityWeightOk() (*float32, bool) {
	if o == nil || o.PriorityWeight == nil {
		return nil, false
	}
	return o.PriorityWeight, true
}

// HasPriorityWeight returns a boolean if a field has been set.
func (o *Task) HasPriorityWeight() bool {
	if o != nil && o.PriorityWeight != nil {
		return true
	}

	return false
}

// SetPriorityWeight gets a reference to the given float32 and assigns it to the PriorityWeight field.
func (o *Task) SetPriorityWeight(v float32) {
	o.PriorityWeight = &v
}

// GetWeightRule returns the WeightRule field value if set, zero value otherwise.
func (o *Task) GetWeightRule() WeightRule {
	if o == nil || o.WeightRule == nil {
		var ret WeightRule
		return ret
	}
	return *o.WeightRule
}

// GetWeightRuleOk returns a tuple with the WeightRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetWeightRuleOk() (*WeightRule, bool) {
	if o == nil || o.WeightRule == nil {
		return nil, false
	}
	return o.WeightRule, true
}

// HasWeightRule returns a boolean if a field has been set.
func (o *Task) HasWeightRule() bool {
	if o != nil && o.WeightRule != nil {
		return true
	}

	return false
}

// SetWeightRule gets a reference to the given WeightRule and assigns it to the WeightRule field.
func (o *Task) SetWeightRule(v WeightRule) {
	o.WeightRule = &v
}

// GetUiColor returns the UiColor field value if set, zero value otherwise.
func (o *Task) GetUiColor() string {
	if o == nil || o.UiColor == nil {
		var ret string
		return ret
	}
	return *o.UiColor
}

// GetUiColorOk returns a tuple with the UiColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetUiColorOk() (*string, bool) {
	if o == nil || o.UiColor == nil {
		return nil, false
	}
	return o.UiColor, true
}

// HasUiColor returns a boolean if a field has been set.
func (o *Task) HasUiColor() bool {
	if o != nil && o.UiColor != nil {
		return true
	}

	return false
}

// SetUiColor gets a reference to the given string and assigns it to the UiColor field.
func (o *Task) SetUiColor(v string) {
	o.UiColor = &v
}

// GetUiFgcolor returns the UiFgcolor field value if set, zero value otherwise.
func (o *Task) GetUiFgcolor() string {
	if o == nil || o.UiFgcolor == nil {
		var ret string
		return ret
	}
	return *o.UiFgcolor
}

// GetUiFgcolorOk returns a tuple with the UiFgcolor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetUiFgcolorOk() (*string, bool) {
	if o == nil || o.UiFgcolor == nil {
		return nil, false
	}
	return o.UiFgcolor, true
}

// HasUiFgcolor returns a boolean if a field has been set.
func (o *Task) HasUiFgcolor() bool {
	if o != nil && o.UiFgcolor != nil {
		return true
	}

	return false
}

// SetUiFgcolor gets a reference to the given string and assigns it to the UiFgcolor field.
func (o *Task) SetUiFgcolor(v string) {
	o.UiFgcolor = &v
}

// GetTemplateFields returns the TemplateFields field value if set, zero value otherwise.
func (o *Task) GetTemplateFields() []string {
	if o == nil || o.TemplateFields == nil {
		var ret []string
		return ret
	}
	return o.TemplateFields
}

// GetTemplateFieldsOk returns a tuple with the TemplateFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetTemplateFieldsOk() ([]string, bool) {
	if o == nil || o.TemplateFields == nil {
		return nil, false
	}
	return o.TemplateFields, true
}

// HasTemplateFields returns a boolean if a field has been set.
func (o *Task) HasTemplateFields() bool {
	if o != nil && o.TemplateFields != nil {
		return true
	}

	return false
}

// SetTemplateFields gets a reference to the given []string and assigns it to the TemplateFields field.
func (o *Task) SetTemplateFields(v []string) {
	o.TemplateFields = v
}

// GetSubDag returns the SubDag field value if set, zero value otherwise.
func (o *Task) GetSubDag() DAG {
	if o == nil || o.SubDag == nil {
		var ret DAG
		return ret
	}
	return *o.SubDag
}

// GetSubDagOk returns a tuple with the SubDag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetSubDagOk() (*DAG, bool) {
	if o == nil || o.SubDag == nil {
		return nil, false
	}
	return o.SubDag, true
}

// HasSubDag returns a boolean if a field has been set.
func (o *Task) HasSubDag() bool {
	if o != nil && o.SubDag != nil {
		return true
	}

	return false
}

// SetSubDag gets a reference to the given DAG and assigns it to the SubDag field.
func (o *Task) SetSubDag(v DAG) {
	o.SubDag = &v
}

// GetDownstreamTaskIds returns the DownstreamTaskIds field value if set, zero value otherwise.
func (o *Task) GetDownstreamTaskIds() []string {
	if o == nil || o.DownstreamTaskIds == nil {
		var ret []string
		return ret
	}
	return o.DownstreamTaskIds
}

// GetDownstreamTaskIdsOk returns a tuple with the DownstreamTaskIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetDownstreamTaskIdsOk() ([]string, bool) {
	if o == nil || o.DownstreamTaskIds == nil {
		return nil, false
	}
	return o.DownstreamTaskIds, true
}

// HasDownstreamTaskIds returns a boolean if a field has been set.
func (o *Task) HasDownstreamTaskIds() bool {
	if o != nil && o.DownstreamTaskIds != nil {
		return true
	}

	return false
}

// SetDownstreamTaskIds gets a reference to the given []string and assigns it to the DownstreamTaskIds field.
func (o *Task) SetDownstreamTaskIds(v []string) {
	o.DownstreamTaskIds = v
}

func (o Task) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClassRef != nil {
		toSerialize["class_ref"] = o.ClassRef
	}
	if o.TaskId != nil {
		toSerialize["task_id"] = o.TaskId
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	if o.EndDate.IsSet() {
		toSerialize["end_date"] = o.EndDate.Get()
	}
	if o.TriggerRule != nil {
		toSerialize["trigger_rule"] = o.TriggerRule
	}
	if o.ExtraLinks != nil {
		toSerialize["extra_links"] = o.ExtraLinks
	}
	if o.DependsOnPast != nil {
		toSerialize["depends_on_past"] = o.DependsOnPast
	}
	if o.IsMapped != nil {
		toSerialize["is_mapped"] = o.IsMapped
	}
	if o.WaitForDownstream != nil {
		toSerialize["wait_for_downstream"] = o.WaitForDownstream
	}
	if o.Retries != nil {
		toSerialize["retries"] = o.Retries
	}
	if o.Queue.IsSet() {
		toSerialize["queue"] = o.Queue.Get()
	}
	if o.Pool != nil {
		toSerialize["pool"] = o.Pool
	}
	if o.PoolSlots != nil {
		toSerialize["pool_slots"] = o.PoolSlots
	}
	if o.ExecutionTimeout.IsSet() {
		toSerialize["execution_timeout"] = o.ExecutionTimeout.Get()
	}
	if o.RetryDelay.IsSet() {
		toSerialize["retry_delay"] = o.RetryDelay.Get()
	}
	if o.RetryExponentialBackoff != nil {
		toSerialize["retry_exponential_backoff"] = o.RetryExponentialBackoff
	}
	if o.PriorityWeight != nil {
		toSerialize["priority_weight"] = o.PriorityWeight
	}
	if o.WeightRule != nil {
		toSerialize["weight_rule"] = o.WeightRule
	}
	if o.UiColor != nil {
		toSerialize["ui_color"] = o.UiColor
	}
	if o.UiFgcolor != nil {
		toSerialize["ui_fgcolor"] = o.UiFgcolor
	}
	if o.TemplateFields != nil {
		toSerialize["template_fields"] = o.TemplateFields
	}
	if o.SubDag != nil {
		toSerialize["sub_dag"] = o.SubDag
	}
	if o.DownstreamTaskIds != nil {
		toSerialize["downstream_task_ids"] = o.DownstreamTaskIds
	}
	return json.Marshal(toSerialize)
}

type NullableTask struct {
	value *Task
	isSet bool
}

func (v NullableTask) Get() *Task {
	return v.value
}

func (v *NullableTask) Set(val *Task) {
	v.value = val
	v.isSet = true
}

func (v NullableTask) IsSet() bool {
	return v.isSet
}

func (v *NullableTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTask(val *Task) *NullableTask {
	return &NullableTask{value: val, isSet: true}
}

func (v NullableTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


