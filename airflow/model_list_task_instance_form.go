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

// ListTaskInstanceForm struct for ListTaskInstanceForm
type ListTaskInstanceForm struct {
	// Return objects with specific DAG IDs. The value can be repeated to retrieve multiple matching values (OR condition).
	DagIds []string `json:"dag_ids,omitempty"`
	// Return objects with specific DAG Run IDs. The value can be repeated to retrieve multiple matching values (OR condition). *New in version 2.7.1*
	DagRunIds []string `json:"dag_run_ids,omitempty"`
	// Return objects with specific task IDs. The value can be repeated to retrieve multiple matching values (OR condition). *New in version 2.7.1*
	TaskIds []string `json:"task_ids,omitempty"`
	// Returns objects greater or equal to the specified date.  This can be combined with execution_date_lte parameter to receive only the selected period. 
	ExecutionDateGte *time.Time `json:"execution_date_gte,omitempty"`
	// Returns objects less than or equal to the specified date.  This can be combined with execution_date_gte parameter to receive only the selected period. 
	ExecutionDateLte *time.Time `json:"execution_date_lte,omitempty"`
	// Returns objects greater or equal the specified date.  This can be combined with start_date_lte parameter to receive only the selected period. 
	StartDateGte *time.Time `json:"start_date_gte,omitempty"`
	// Returns objects less or equal the specified date.  This can be combined with start_date_gte parameter to receive only the selected period. 
	StartDateLte *time.Time `json:"start_date_lte,omitempty"`
	// Returns objects greater or equal the specified date.  This can be combined with start_date_lte parameter to receive only the selected period. 
	EndDateGte *time.Time `json:"end_date_gte,omitempty"`
	// Returns objects less than or equal to the specified date.  This can be combined with start_date_gte parameter to receive only the selected period. 
	EndDateLte *time.Time `json:"end_date_lte,omitempty"`
	// Returns objects greater than or equal to the specified values.  This can be combined with duration_lte parameter to receive only the selected period. 
	DurationGte *float32 `json:"duration_gte,omitempty"`
	// Returns objects less than or equal to the specified values.  This can be combined with duration_gte parameter to receive only the selected range. 
	DurationLte *float32 `json:"duration_lte,omitempty"`
	// The value can be repeated to retrieve multiple matching values (OR condition).
	State []TaskState `json:"state,omitempty"`
	// The value can be repeated to retrieve multiple matching values (OR condition).
	Pool []string `json:"pool,omitempty"`
	// The value can be repeated to retrieve multiple matching values (OR condition).
	Queue []string `json:"queue,omitempty"`
}

// NewListTaskInstanceForm instantiates a new ListTaskInstanceForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTaskInstanceForm() *ListTaskInstanceForm {
	this := ListTaskInstanceForm{}
	return &this
}

// NewListTaskInstanceFormWithDefaults instantiates a new ListTaskInstanceForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTaskInstanceFormWithDefaults() *ListTaskInstanceForm {
	this := ListTaskInstanceForm{}
	return &this
}

// GetDagIds returns the DagIds field value if set, zero value otherwise.
func (o *ListTaskInstanceForm) GetDagIds() []string {
	if o == nil || o.DagIds == nil {
		var ret []string
		return ret
	}
	return o.DagIds
}

// GetDagIdsOk returns a tuple with the DagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskInstanceForm) GetDagIdsOk() ([]string, bool) {
	if o == nil || o.DagIds == nil {
		return nil, false
	}
	return o.DagIds, true
}

// HasDagIds returns a boolean if a field has been set.
func (o *ListTaskInstanceForm) HasDagIds() bool {
	if o != nil && o.DagIds != nil {
		return true
	}

	return false
}

// SetDagIds gets a reference to the given []string and assigns it to the DagIds field.
func (o *ListTaskInstanceForm) SetDagIds(v []string) {
	o.DagIds = v
}

// GetDagRunIds returns the DagRunIds field value if set, zero value otherwise.
func (o *ListTaskInstanceForm) GetDagRunIds() []string {
	if o == nil || o.DagRunIds == nil {
		var ret []string
		return ret
	}
	return o.DagRunIds
}

// GetDagRunIdsOk returns a tuple with the DagRunIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskInstanceForm) GetDagRunIdsOk() ([]string, bool) {
	if o == nil || o.DagRunIds == nil {
		return nil, false
	}
	return o.DagRunIds, true
}

// HasDagRunIds returns a boolean if a field has been set.
func (o *ListTaskInstanceForm) HasDagRunIds() bool {
	if o != nil && o.DagRunIds != nil {
		return true
	}

	return false
}

// SetDagRunIds gets a reference to the given []string and assigns it to the DagRunIds field.
func (o *ListTaskInstanceForm) SetDagRunIds(v []string) {
	o.DagRunIds = v
}

// GetTaskIds returns the TaskIds field value if set, zero value otherwise.
func (o *ListTaskInstanceForm) GetTaskIds() []string {
	if o == nil || o.TaskIds == nil {
		var ret []string
		return ret
	}
	return o.TaskIds
}

// GetTaskIdsOk returns a tuple with the TaskIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskInstanceForm) GetTaskIdsOk() ([]string, bool) {
	if o == nil || o.TaskIds == nil {
		return nil, false
	}
	return o.TaskIds, true
}

// HasTaskIds returns a boolean if a field has been set.
func (o *ListTaskInstanceForm) HasTaskIds() bool {
	if o != nil && o.TaskIds != nil {
		return true
	}

	return false
}

// SetTaskIds gets a reference to the given []string and assigns it to the TaskIds field.
func (o *ListTaskInstanceForm) SetTaskIds(v []string) {
	o.TaskIds = v
}

// GetExecutionDateGte returns the ExecutionDateGte field value if set, zero value otherwise.
func (o *ListTaskInstanceForm) GetExecutionDateGte() time.Time {
	if o == nil || o.ExecutionDateGte == nil {
		var ret time.Time
		return ret
	}
	return *o.ExecutionDateGte
}

// GetExecutionDateGteOk returns a tuple with the ExecutionDateGte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskInstanceForm) GetExecutionDateGteOk() (*time.Time, bool) {
	if o == nil || o.ExecutionDateGte == nil {
		return nil, false
	}
	return o.ExecutionDateGte, true
}

// HasExecutionDateGte returns a boolean if a field has been set.
func (o *ListTaskInstanceForm) HasExecutionDateGte() bool {
	if o != nil && o.ExecutionDateGte != nil {
		return true
	}

	return false
}

// SetExecutionDateGte gets a reference to the given time.Time and assigns it to the ExecutionDateGte field.
func (o *ListTaskInstanceForm) SetExecutionDateGte(v time.Time) {
	o.ExecutionDateGte = &v
}

// GetExecutionDateLte returns the ExecutionDateLte field value if set, zero value otherwise.
func (o *ListTaskInstanceForm) GetExecutionDateLte() time.Time {
	if o == nil || o.ExecutionDateLte == nil {
		var ret time.Time
		return ret
	}
	return *o.ExecutionDateLte
}

// GetExecutionDateLteOk returns a tuple with the ExecutionDateLte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskInstanceForm) GetExecutionDateLteOk() (*time.Time, bool) {
	if o == nil || o.ExecutionDateLte == nil {
		return nil, false
	}
	return o.ExecutionDateLte, true
}

// HasExecutionDateLte returns a boolean if a field has been set.
func (o *ListTaskInstanceForm) HasExecutionDateLte() bool {
	if o != nil && o.ExecutionDateLte != nil {
		return true
	}

	return false
}

// SetExecutionDateLte gets a reference to the given time.Time and assigns it to the ExecutionDateLte field.
func (o *ListTaskInstanceForm) SetExecutionDateLte(v time.Time) {
	o.ExecutionDateLte = &v
}

// GetStartDateGte returns the StartDateGte field value if set, zero value otherwise.
func (o *ListTaskInstanceForm) GetStartDateGte() time.Time {
	if o == nil || o.StartDateGte == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDateGte
}

// GetStartDateGteOk returns a tuple with the StartDateGte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskInstanceForm) GetStartDateGteOk() (*time.Time, bool) {
	if o == nil || o.StartDateGte == nil {
		return nil, false
	}
	return o.StartDateGte, true
}

// HasStartDateGte returns a boolean if a field has been set.
func (o *ListTaskInstanceForm) HasStartDateGte() bool {
	if o != nil && o.StartDateGte != nil {
		return true
	}

	return false
}

// SetStartDateGte gets a reference to the given time.Time and assigns it to the StartDateGte field.
func (o *ListTaskInstanceForm) SetStartDateGte(v time.Time) {
	o.StartDateGte = &v
}

// GetStartDateLte returns the StartDateLte field value if set, zero value otherwise.
func (o *ListTaskInstanceForm) GetStartDateLte() time.Time {
	if o == nil || o.StartDateLte == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDateLte
}

// GetStartDateLteOk returns a tuple with the StartDateLte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskInstanceForm) GetStartDateLteOk() (*time.Time, bool) {
	if o == nil || o.StartDateLte == nil {
		return nil, false
	}
	return o.StartDateLte, true
}

// HasStartDateLte returns a boolean if a field has been set.
func (o *ListTaskInstanceForm) HasStartDateLte() bool {
	if o != nil && o.StartDateLte != nil {
		return true
	}

	return false
}

// SetStartDateLte gets a reference to the given time.Time and assigns it to the StartDateLte field.
func (o *ListTaskInstanceForm) SetStartDateLte(v time.Time) {
	o.StartDateLte = &v
}

// GetEndDateGte returns the EndDateGte field value if set, zero value otherwise.
func (o *ListTaskInstanceForm) GetEndDateGte() time.Time {
	if o == nil || o.EndDateGte == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDateGte
}

// GetEndDateGteOk returns a tuple with the EndDateGte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskInstanceForm) GetEndDateGteOk() (*time.Time, bool) {
	if o == nil || o.EndDateGte == nil {
		return nil, false
	}
	return o.EndDateGte, true
}

// HasEndDateGte returns a boolean if a field has been set.
func (o *ListTaskInstanceForm) HasEndDateGte() bool {
	if o != nil && o.EndDateGte != nil {
		return true
	}

	return false
}

// SetEndDateGte gets a reference to the given time.Time and assigns it to the EndDateGte field.
func (o *ListTaskInstanceForm) SetEndDateGte(v time.Time) {
	o.EndDateGte = &v
}

// GetEndDateLte returns the EndDateLte field value if set, zero value otherwise.
func (o *ListTaskInstanceForm) GetEndDateLte() time.Time {
	if o == nil || o.EndDateLte == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDateLte
}

// GetEndDateLteOk returns a tuple with the EndDateLte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskInstanceForm) GetEndDateLteOk() (*time.Time, bool) {
	if o == nil || o.EndDateLte == nil {
		return nil, false
	}
	return o.EndDateLte, true
}

// HasEndDateLte returns a boolean if a field has been set.
func (o *ListTaskInstanceForm) HasEndDateLte() bool {
	if o != nil && o.EndDateLte != nil {
		return true
	}

	return false
}

// SetEndDateLte gets a reference to the given time.Time and assigns it to the EndDateLte field.
func (o *ListTaskInstanceForm) SetEndDateLte(v time.Time) {
	o.EndDateLte = &v
}

// GetDurationGte returns the DurationGte field value if set, zero value otherwise.
func (o *ListTaskInstanceForm) GetDurationGte() float32 {
	if o == nil || o.DurationGte == nil {
		var ret float32
		return ret
	}
	return *o.DurationGte
}

// GetDurationGteOk returns a tuple with the DurationGte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskInstanceForm) GetDurationGteOk() (*float32, bool) {
	if o == nil || o.DurationGte == nil {
		return nil, false
	}
	return o.DurationGte, true
}

// HasDurationGte returns a boolean if a field has been set.
func (o *ListTaskInstanceForm) HasDurationGte() bool {
	if o != nil && o.DurationGte != nil {
		return true
	}

	return false
}

// SetDurationGte gets a reference to the given float32 and assigns it to the DurationGte field.
func (o *ListTaskInstanceForm) SetDurationGte(v float32) {
	o.DurationGte = &v
}

// GetDurationLte returns the DurationLte field value if set, zero value otherwise.
func (o *ListTaskInstanceForm) GetDurationLte() float32 {
	if o == nil || o.DurationLte == nil {
		var ret float32
		return ret
	}
	return *o.DurationLte
}

// GetDurationLteOk returns a tuple with the DurationLte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskInstanceForm) GetDurationLteOk() (*float32, bool) {
	if o == nil || o.DurationLte == nil {
		return nil, false
	}
	return o.DurationLte, true
}

// HasDurationLte returns a boolean if a field has been set.
func (o *ListTaskInstanceForm) HasDurationLte() bool {
	if o != nil && o.DurationLte != nil {
		return true
	}

	return false
}

// SetDurationLte gets a reference to the given float32 and assigns it to the DurationLte field.
func (o *ListTaskInstanceForm) SetDurationLte(v float32) {
	o.DurationLte = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ListTaskInstanceForm) GetState() []TaskState {
	if o == nil || o.State == nil {
		var ret []TaskState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskInstanceForm) GetStateOk() ([]TaskState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ListTaskInstanceForm) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given []TaskState and assigns it to the State field.
func (o *ListTaskInstanceForm) SetState(v []TaskState) {
	o.State = v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *ListTaskInstanceForm) GetPool() []string {
	if o == nil || o.Pool == nil {
		var ret []string
		return ret
	}
	return o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskInstanceForm) GetPoolOk() ([]string, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *ListTaskInstanceForm) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given []string and assigns it to the Pool field.
func (o *ListTaskInstanceForm) SetPool(v []string) {
	o.Pool = v
}

// GetQueue returns the Queue field value if set, zero value otherwise.
func (o *ListTaskInstanceForm) GetQueue() []string {
	if o == nil || o.Queue == nil {
		var ret []string
		return ret
	}
	return o.Queue
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskInstanceForm) GetQueueOk() ([]string, bool) {
	if o == nil || o.Queue == nil {
		return nil, false
	}
	return o.Queue, true
}

// HasQueue returns a boolean if a field has been set.
func (o *ListTaskInstanceForm) HasQueue() bool {
	if o != nil && o.Queue != nil {
		return true
	}

	return false
}

// SetQueue gets a reference to the given []string and assigns it to the Queue field.
func (o *ListTaskInstanceForm) SetQueue(v []string) {
	o.Queue = v
}

func (o ListTaskInstanceForm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DagIds != nil {
		toSerialize["dag_ids"] = o.DagIds
	}
	if o.DagRunIds != nil {
		toSerialize["dag_run_ids"] = o.DagRunIds
	}
	if o.TaskIds != nil {
		toSerialize["task_ids"] = o.TaskIds
	}
	if o.ExecutionDateGte != nil {
		toSerialize["execution_date_gte"] = o.ExecutionDateGte
	}
	if o.ExecutionDateLte != nil {
		toSerialize["execution_date_lte"] = o.ExecutionDateLte
	}
	if o.StartDateGte != nil {
		toSerialize["start_date_gte"] = o.StartDateGte
	}
	if o.StartDateLte != nil {
		toSerialize["start_date_lte"] = o.StartDateLte
	}
	if o.EndDateGte != nil {
		toSerialize["end_date_gte"] = o.EndDateGte
	}
	if o.EndDateLte != nil {
		toSerialize["end_date_lte"] = o.EndDateLte
	}
	if o.DurationGte != nil {
		toSerialize["duration_gte"] = o.DurationGte
	}
	if o.DurationLte != nil {
		toSerialize["duration_lte"] = o.DurationLte
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Pool != nil {
		toSerialize["pool"] = o.Pool
	}
	if o.Queue != nil {
		toSerialize["queue"] = o.Queue
	}
	return json.Marshal(toSerialize)
}

type NullableListTaskInstanceForm struct {
	value *ListTaskInstanceForm
	isSet bool
}

func (v NullableListTaskInstanceForm) Get() *ListTaskInstanceForm {
	return v.value
}

func (v *NullableListTaskInstanceForm) Set(val *ListTaskInstanceForm) {
	v.value = val
	v.isSet = true
}

func (v NullableListTaskInstanceForm) IsSet() bool {
	return v.isSet
}

func (v *NullableListTaskInstanceForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTaskInstanceForm(val *ListTaskInstanceForm) *NullableListTaskInstanceForm {
	return &NullableListTaskInstanceForm{value: val, isSet: true}
}

func (v NullableListTaskInstanceForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTaskInstanceForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


