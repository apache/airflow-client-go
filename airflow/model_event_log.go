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

// EventLog Log of user operations via CLI or Web UI.
type EventLog struct {
	// The event log ID
	EventLogId *int32 `json:"event_log_id,omitempty"`
	// The time when these events happened.
	When *time.Time `json:"when,omitempty"`
	// The DAG ID
	DagId NullableString `json:"dag_id,omitempty"`
	// The DAG ID
	TaskId NullableString `json:"task_id,omitempty"`
	// A key describing the type of event.
	Event *string `json:"event,omitempty"`
	// When the event was dispatched for an object having execution_date, the value of this field. 
	ExecutionDate NullableTime `json:"execution_date,omitempty"`
	// Name of the user who triggered these events a.
	Owner *string `json:"owner,omitempty"`
	// Other information that was not included in the other fields, e.g. the complete CLI command. 
	Extra NullableString `json:"extra,omitempty"`
}

// NewEventLog instantiates a new EventLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventLog() *EventLog {
	this := EventLog{}
	return &this
}

// NewEventLogWithDefaults instantiates a new EventLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventLogWithDefaults() *EventLog {
	this := EventLog{}
	return &this
}

// GetEventLogId returns the EventLogId field value if set, zero value otherwise.
func (o *EventLog) GetEventLogId() int32 {
	if o == nil || o.EventLogId == nil {
		var ret int32
		return ret
	}
	return *o.EventLogId
}

// GetEventLogIdOk returns a tuple with the EventLogId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventLog) GetEventLogIdOk() (*int32, bool) {
	if o == nil || o.EventLogId == nil {
		return nil, false
	}
	return o.EventLogId, true
}

// HasEventLogId returns a boolean if a field has been set.
func (o *EventLog) HasEventLogId() bool {
	if o != nil && o.EventLogId != nil {
		return true
	}

	return false
}

// SetEventLogId gets a reference to the given int32 and assigns it to the EventLogId field.
func (o *EventLog) SetEventLogId(v int32) {
	o.EventLogId = &v
}

// GetWhen returns the When field value if set, zero value otherwise.
func (o *EventLog) GetWhen() time.Time {
	if o == nil || o.When == nil {
		var ret time.Time
		return ret
	}
	return *o.When
}

// GetWhenOk returns a tuple with the When field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventLog) GetWhenOk() (*time.Time, bool) {
	if o == nil || o.When == nil {
		return nil, false
	}
	return o.When, true
}

// HasWhen returns a boolean if a field has been set.
func (o *EventLog) HasWhen() bool {
	if o != nil && o.When != nil {
		return true
	}

	return false
}

// SetWhen gets a reference to the given time.Time and assigns it to the When field.
func (o *EventLog) SetWhen(v time.Time) {
	o.When = &v
}

// GetDagId returns the DagId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventLog) GetDagId() string {
	if o == nil || o.DagId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DagId.Get()
}

// GetDagIdOk returns a tuple with the DagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventLog) GetDagIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DagId.Get(), o.DagId.IsSet()
}

// HasDagId returns a boolean if a field has been set.
func (o *EventLog) HasDagId() bool {
	if o != nil && o.DagId.IsSet() {
		return true
	}

	return false
}

// SetDagId gets a reference to the given NullableString and assigns it to the DagId field.
func (o *EventLog) SetDagId(v string) {
	o.DagId.Set(&v)
}
// SetDagIdNil sets the value for DagId to be an explicit nil
func (o *EventLog) SetDagIdNil() {
	o.DagId.Set(nil)
}

// UnsetDagId ensures that no value is present for DagId, not even an explicit nil
func (o *EventLog) UnsetDagId() {
	o.DagId.Unset()
}

// GetTaskId returns the TaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventLog) GetTaskId() string {
	if o == nil || o.TaskId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TaskId.Get()
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventLog) GetTaskIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TaskId.Get(), o.TaskId.IsSet()
}

// HasTaskId returns a boolean if a field has been set.
func (o *EventLog) HasTaskId() bool {
	if o != nil && o.TaskId.IsSet() {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given NullableString and assigns it to the TaskId field.
func (o *EventLog) SetTaskId(v string) {
	o.TaskId.Set(&v)
}
// SetTaskIdNil sets the value for TaskId to be an explicit nil
func (o *EventLog) SetTaskIdNil() {
	o.TaskId.Set(nil)
}

// UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
func (o *EventLog) UnsetTaskId() {
	o.TaskId.Unset()
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *EventLog) GetEvent() string {
	if o == nil || o.Event == nil {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventLog) GetEventOk() (*string, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *EventLog) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *EventLog) SetEvent(v string) {
	o.Event = &v
}

// GetExecutionDate returns the ExecutionDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventLog) GetExecutionDate() time.Time {
	if o == nil || o.ExecutionDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExecutionDate.Get()
}

// GetExecutionDateOk returns a tuple with the ExecutionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventLog) GetExecutionDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExecutionDate.Get(), o.ExecutionDate.IsSet()
}

// HasExecutionDate returns a boolean if a field has been set.
func (o *EventLog) HasExecutionDate() bool {
	if o != nil && o.ExecutionDate.IsSet() {
		return true
	}

	return false
}

// SetExecutionDate gets a reference to the given NullableTime and assigns it to the ExecutionDate field.
func (o *EventLog) SetExecutionDate(v time.Time) {
	o.ExecutionDate.Set(&v)
}
// SetExecutionDateNil sets the value for ExecutionDate to be an explicit nil
func (o *EventLog) SetExecutionDateNil() {
	o.ExecutionDate.Set(nil)
}

// UnsetExecutionDate ensures that no value is present for ExecutionDate, not even an explicit nil
func (o *EventLog) UnsetExecutionDate() {
	o.ExecutionDate.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *EventLog) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventLog) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *EventLog) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *EventLog) SetOwner(v string) {
	o.Owner = &v
}

// GetExtra returns the Extra field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventLog) GetExtra() string {
	if o == nil || o.Extra.Get() == nil {
		var ret string
		return ret
	}
	return *o.Extra.Get()
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventLog) GetExtraOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Extra.Get(), o.Extra.IsSet()
}

// HasExtra returns a boolean if a field has been set.
func (o *EventLog) HasExtra() bool {
	if o != nil && o.Extra.IsSet() {
		return true
	}

	return false
}

// SetExtra gets a reference to the given NullableString and assigns it to the Extra field.
func (o *EventLog) SetExtra(v string) {
	o.Extra.Set(&v)
}
// SetExtraNil sets the value for Extra to be an explicit nil
func (o *EventLog) SetExtraNil() {
	o.Extra.Set(nil)
}

// UnsetExtra ensures that no value is present for Extra, not even an explicit nil
func (o *EventLog) UnsetExtra() {
	o.Extra.Unset()
}

func (o EventLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EventLogId != nil {
		toSerialize["event_log_id"] = o.EventLogId
	}
	if o.When != nil {
		toSerialize["when"] = o.When
	}
	if o.DagId.IsSet() {
		toSerialize["dag_id"] = o.DagId.Get()
	}
	if o.TaskId.IsSet() {
		toSerialize["task_id"] = o.TaskId.Get()
	}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if o.ExecutionDate.IsSet() {
		toSerialize["execution_date"] = o.ExecutionDate.Get()
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Extra.IsSet() {
		toSerialize["extra"] = o.Extra.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEventLog struct {
	value *EventLog
	isSet bool
}

func (v NullableEventLog) Get() *EventLog {
	return v.value
}

func (v *NullableEventLog) Set(val *EventLog) {
	v.value = val
	v.isSet = true
}

func (v NullableEventLog) IsSet() bool {
	return v.isSet
}

func (v *NullableEventLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventLog(val *EventLog) *NullableEventLog {
	return &NullableEventLog{value: val, isSet: true}
}

func (v NullableEventLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


