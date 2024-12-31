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

// Job struct for Job
type Job struct {
	Id *int32 `json:"id,omitempty"`
	DagId NullableString `json:"dag_id,omitempty"`
	State NullableString `json:"state,omitempty"`
	JobType NullableString `json:"job_type,omitempty"`
	StartDate NullableString `json:"start_date,omitempty"`
	EndDate NullableString `json:"end_date,omitempty"`
	LatestHeartbeat NullableString `json:"latest_heartbeat,omitempty"`
	ExecutorClass NullableString `json:"executor_class,omitempty"`
	Hostname NullableString `json:"hostname,omitempty"`
	Unixname NullableString `json:"unixname,omitempty"`
}

// NewJob instantiates a new Job object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJob() *Job {
	this := Job{}
	return &this
}

// NewJobWithDefaults instantiates a new Job object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobWithDefaults() *Job {
	this := Job{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Job) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Job) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Job) SetId(v int32) {
	o.Id = &v
}

// GetDagId returns the DagId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetDagId() string {
	if o == nil || o.DagId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DagId.Get()
}

// GetDagIdOk returns a tuple with the DagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetDagIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DagId.Get(), o.DagId.IsSet()
}

// HasDagId returns a boolean if a field has been set.
func (o *Job) HasDagId() bool {
	if o != nil && o.DagId.IsSet() {
		return true
	}

	return false
}

// SetDagId gets a reference to the given NullableString and assigns it to the DagId field.
func (o *Job) SetDagId(v string) {
	o.DagId.Set(&v)
}
// SetDagIdNil sets the value for DagId to be an explicit nil
func (o *Job) SetDagIdNil() {
	o.DagId.Set(nil)
}

// UnsetDagId ensures that no value is present for DagId, not even an explicit nil
func (o *Job) UnsetDagId() {
	o.DagId.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetState() string {
	if o == nil || o.State.Get() == nil {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *Job) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *Job) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *Job) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *Job) UnsetState() {
	o.State.Unset()
}

// GetJobType returns the JobType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetJobType() string {
	if o == nil || o.JobType.Get() == nil {
		var ret string
		return ret
	}
	return *o.JobType.Get()
}

// GetJobTypeOk returns a tuple with the JobType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetJobTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.JobType.Get(), o.JobType.IsSet()
}

// HasJobType returns a boolean if a field has been set.
func (o *Job) HasJobType() bool {
	if o != nil && o.JobType.IsSet() {
		return true
	}

	return false
}

// SetJobType gets a reference to the given NullableString and assigns it to the JobType field.
func (o *Job) SetJobType(v string) {
	o.JobType.Set(&v)
}
// SetJobTypeNil sets the value for JobType to be an explicit nil
func (o *Job) SetJobTypeNil() {
	o.JobType.Set(nil)
}

// UnsetJobType ensures that no value is present for JobType, not even an explicit nil
func (o *Job) UnsetJobType() {
	o.JobType.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetStartDate() string {
	if o == nil || o.StartDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *Job) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableString and assigns it to the StartDate field.
func (o *Job) SetStartDate(v string) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *Job) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *Job) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetEndDate() string {
	if o == nil || o.EndDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *Job) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableString and assigns it to the EndDate field.
func (o *Job) SetEndDate(v string) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *Job) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *Job) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetLatestHeartbeat returns the LatestHeartbeat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetLatestHeartbeat() string {
	if o == nil || o.LatestHeartbeat.Get() == nil {
		var ret string
		return ret
	}
	return *o.LatestHeartbeat.Get()
}

// GetLatestHeartbeatOk returns a tuple with the LatestHeartbeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetLatestHeartbeatOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LatestHeartbeat.Get(), o.LatestHeartbeat.IsSet()
}

// HasLatestHeartbeat returns a boolean if a field has been set.
func (o *Job) HasLatestHeartbeat() bool {
	if o != nil && o.LatestHeartbeat.IsSet() {
		return true
	}

	return false
}

// SetLatestHeartbeat gets a reference to the given NullableString and assigns it to the LatestHeartbeat field.
func (o *Job) SetLatestHeartbeat(v string) {
	o.LatestHeartbeat.Set(&v)
}
// SetLatestHeartbeatNil sets the value for LatestHeartbeat to be an explicit nil
func (o *Job) SetLatestHeartbeatNil() {
	o.LatestHeartbeat.Set(nil)
}

// UnsetLatestHeartbeat ensures that no value is present for LatestHeartbeat, not even an explicit nil
func (o *Job) UnsetLatestHeartbeat() {
	o.LatestHeartbeat.Unset()
}

// GetExecutorClass returns the ExecutorClass field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetExecutorClass() string {
	if o == nil || o.ExecutorClass.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExecutorClass.Get()
}

// GetExecutorClassOk returns a tuple with the ExecutorClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetExecutorClassOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExecutorClass.Get(), o.ExecutorClass.IsSet()
}

// HasExecutorClass returns a boolean if a field has been set.
func (o *Job) HasExecutorClass() bool {
	if o != nil && o.ExecutorClass.IsSet() {
		return true
	}

	return false
}

// SetExecutorClass gets a reference to the given NullableString and assigns it to the ExecutorClass field.
func (o *Job) SetExecutorClass(v string) {
	o.ExecutorClass.Set(&v)
}
// SetExecutorClassNil sets the value for ExecutorClass to be an explicit nil
func (o *Job) SetExecutorClassNil() {
	o.ExecutorClass.Set(nil)
}

// UnsetExecutorClass ensures that no value is present for ExecutorClass, not even an explicit nil
func (o *Job) UnsetExecutorClass() {
	o.ExecutorClass.Unset()
}

// GetHostname returns the Hostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetHostname() string {
	if o == nil || o.Hostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.Hostname.Get()
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Hostname.Get(), o.Hostname.IsSet()
}

// HasHostname returns a boolean if a field has been set.
func (o *Job) HasHostname() bool {
	if o != nil && o.Hostname.IsSet() {
		return true
	}

	return false
}

// SetHostname gets a reference to the given NullableString and assigns it to the Hostname field.
func (o *Job) SetHostname(v string) {
	o.Hostname.Set(&v)
}
// SetHostnameNil sets the value for Hostname to be an explicit nil
func (o *Job) SetHostnameNil() {
	o.Hostname.Set(nil)
}

// UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
func (o *Job) UnsetHostname() {
	o.Hostname.Unset()
}

// GetUnixname returns the Unixname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetUnixname() string {
	if o == nil || o.Unixname.Get() == nil {
		var ret string
		return ret
	}
	return *o.Unixname.Get()
}

// GetUnixnameOk returns a tuple with the Unixname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetUnixnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Unixname.Get(), o.Unixname.IsSet()
}

// HasUnixname returns a boolean if a field has been set.
func (o *Job) HasUnixname() bool {
	if o != nil && o.Unixname.IsSet() {
		return true
	}

	return false
}

// SetUnixname gets a reference to the given NullableString and assigns it to the Unixname field.
func (o *Job) SetUnixname(v string) {
	o.Unixname.Set(&v)
}
// SetUnixnameNil sets the value for Unixname to be an explicit nil
func (o *Job) SetUnixnameNil() {
	o.Unixname.Set(nil)
}

// UnsetUnixname ensures that no value is present for Unixname, not even an explicit nil
func (o *Job) UnsetUnixname() {
	o.Unixname.Unset()
}

func (o Job) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DagId.IsSet() {
		toSerialize["dag_id"] = o.DagId.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.JobType.IsSet() {
		toSerialize["job_type"] = o.JobType.Get()
	}
	if o.StartDate.IsSet() {
		toSerialize["start_date"] = o.StartDate.Get()
	}
	if o.EndDate.IsSet() {
		toSerialize["end_date"] = o.EndDate.Get()
	}
	if o.LatestHeartbeat.IsSet() {
		toSerialize["latest_heartbeat"] = o.LatestHeartbeat.Get()
	}
	if o.ExecutorClass.IsSet() {
		toSerialize["executor_class"] = o.ExecutorClass.Get()
	}
	if o.Hostname.IsSet() {
		toSerialize["hostname"] = o.Hostname.Get()
	}
	if o.Unixname.IsSet() {
		toSerialize["unixname"] = o.Unixname.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableJob struct {
	value *Job
	isSet bool
}

func (v NullableJob) Get() *Job {
	return v.value
}

func (v *NullableJob) Set(val *Job) {
	v.value = val
	v.isSet = true
}

func (v NullableJob) IsSet() bool {
	return v.isSet
}

func (v *NullableJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJob(val *Job) *NullableJob {
	return &NullableJob{value: val, isSet: true}
}

func (v NullableJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


