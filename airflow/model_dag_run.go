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

// DAGRun struct for DAGRun
type DAGRun struct {
	// Run ID.  The value of this field can be set only when creating the object. If you try to modify the field of an existing object, the request fails with an BAD_REQUEST error.  If not provided, a value will be generated based on execution_date.  If the specified dag_run_id is in use, the creation request fails with an ALREADY_EXISTS error.  This together with DAG_ID are a unique key. 
	DagRunId NullableString `json:"dag_run_id,omitempty"`
	DagId *string `json:"dag_id,omitempty"`
	// The logical date (previously called execution date). This is the time or interval covered by this DAG run, according to the DAG definition.  The value of this field can be set only when creating the object. If you try to modify the field of an existing object, the request fails with an BAD_REQUEST error.  This together with DAG_ID are a unique key.  *New in version 2.2.0* 
	LogicalDate NullableTime `json:"logical_date,omitempty"`
	// The execution date. This is the same as logical_date, kept for backwards compatibility. If both this field and logical_date are provided but with different values, the request will fail with an BAD_REQUEST error.  *Changed in version 2.2.0*&#58; Field becomes nullable.  *Deprecated since version 2.2.0*&#58; Use 'logical_date' instead. 
	// Deprecated
	ExecutionDate NullableTime `json:"execution_date,omitempty"`
	// The start time. The time when DAG run was actually created.  *Changed in version 2.1.3*&#58; Field becomes nullable. 
	StartDate NullableTime `json:"start_date,omitempty"`
	EndDate NullableTime `json:"end_date,omitempty"`
	DataIntervalStart NullableTime `json:"data_interval_start,omitempty"`
	DataIntervalEnd NullableTime `json:"data_interval_end,omitempty"`
	LastSchedulingDecision NullableTime `json:"last_scheduling_decision,omitempty"`
	RunType *string `json:"run_type,omitempty"`
	State *DagState `json:"state,omitempty"`
	ExternalTrigger *bool `json:"external_trigger,omitempty"`
	// JSON object describing additional configuration parameters.  The value of this field can be set only when creating the object. If you try to modify the field of an existing object, the request fails with an BAD_REQUEST error. 
	Conf map[string]interface{} `json:"conf,omitempty"`
	// Contains manually entered notes by the user about the DagRun.  *New in version 2.5.0* 
	Note NullableString `json:"note,omitempty"`
}

// NewDAGRun instantiates a new DAGRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDAGRun() *DAGRun {
	this := DAGRun{}
	return &this
}

// NewDAGRunWithDefaults instantiates a new DAGRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDAGRunWithDefaults() *DAGRun {
	this := DAGRun{}
	return &this
}

// GetDagRunId returns the DagRunId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGRun) GetDagRunId() string {
	if o == nil || o.DagRunId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DagRunId.Get()
}

// GetDagRunIdOk returns a tuple with the DagRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGRun) GetDagRunIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DagRunId.Get(), o.DagRunId.IsSet()
}

// HasDagRunId returns a boolean if a field has been set.
func (o *DAGRun) HasDagRunId() bool {
	if o != nil && o.DagRunId.IsSet() {
		return true
	}

	return false
}

// SetDagRunId gets a reference to the given NullableString and assigns it to the DagRunId field.
func (o *DAGRun) SetDagRunId(v string) {
	o.DagRunId.Set(&v)
}
// SetDagRunIdNil sets the value for DagRunId to be an explicit nil
func (o *DAGRun) SetDagRunIdNil() {
	o.DagRunId.Set(nil)
}

// UnsetDagRunId ensures that no value is present for DagRunId, not even an explicit nil
func (o *DAGRun) UnsetDagRunId() {
	o.DagRunId.Unset()
}

// GetDagId returns the DagId field value if set, zero value otherwise.
func (o *DAGRun) GetDagId() string {
	if o == nil || o.DagId == nil {
		var ret string
		return ret
	}
	return *o.DagId
}

// GetDagIdOk returns a tuple with the DagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGRun) GetDagIdOk() (*string, bool) {
	if o == nil || o.DagId == nil {
		return nil, false
	}
	return o.DagId, true
}

// HasDagId returns a boolean if a field has been set.
func (o *DAGRun) HasDagId() bool {
	if o != nil && o.DagId != nil {
		return true
	}

	return false
}

// SetDagId gets a reference to the given string and assigns it to the DagId field.
func (o *DAGRun) SetDagId(v string) {
	o.DagId = &v
}

// GetLogicalDate returns the LogicalDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGRun) GetLogicalDate() time.Time {
	if o == nil || o.LogicalDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LogicalDate.Get()
}

// GetLogicalDateOk returns a tuple with the LogicalDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGRun) GetLogicalDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LogicalDate.Get(), o.LogicalDate.IsSet()
}

// HasLogicalDate returns a boolean if a field has been set.
func (o *DAGRun) HasLogicalDate() bool {
	if o != nil && o.LogicalDate.IsSet() {
		return true
	}

	return false
}

// SetLogicalDate gets a reference to the given NullableTime and assigns it to the LogicalDate field.
func (o *DAGRun) SetLogicalDate(v time.Time) {
	o.LogicalDate.Set(&v)
}
// SetLogicalDateNil sets the value for LogicalDate to be an explicit nil
func (o *DAGRun) SetLogicalDateNil() {
	o.LogicalDate.Set(nil)
}

// UnsetLogicalDate ensures that no value is present for LogicalDate, not even an explicit nil
func (o *DAGRun) UnsetLogicalDate() {
	o.LogicalDate.Unset()
}

// GetExecutionDate returns the ExecutionDate field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *DAGRun) GetExecutionDate() time.Time {
	if o == nil || o.ExecutionDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExecutionDate.Get()
}

// GetExecutionDateOk returns a tuple with the ExecutionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *DAGRun) GetExecutionDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExecutionDate.Get(), o.ExecutionDate.IsSet()
}

// HasExecutionDate returns a boolean if a field has been set.
func (o *DAGRun) HasExecutionDate() bool {
	if o != nil && o.ExecutionDate.IsSet() {
		return true
	}

	return false
}

// SetExecutionDate gets a reference to the given NullableTime and assigns it to the ExecutionDate field.
// Deprecated
func (o *DAGRun) SetExecutionDate(v time.Time) {
	o.ExecutionDate.Set(&v)
}
// SetExecutionDateNil sets the value for ExecutionDate to be an explicit nil
func (o *DAGRun) SetExecutionDateNil() {
	o.ExecutionDate.Set(nil)
}

// UnsetExecutionDate ensures that no value is present for ExecutionDate, not even an explicit nil
func (o *DAGRun) UnsetExecutionDate() {
	o.ExecutionDate.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGRun) GetStartDate() time.Time {
	if o == nil || o.StartDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGRun) GetStartDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *DAGRun) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableTime and assigns it to the StartDate field.
func (o *DAGRun) SetStartDate(v time.Time) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *DAGRun) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *DAGRun) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGRun) GetEndDate() time.Time {
	if o == nil || o.EndDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGRun) GetEndDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *DAGRun) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableTime and assigns it to the EndDate field.
func (o *DAGRun) SetEndDate(v time.Time) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *DAGRun) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *DAGRun) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetDataIntervalStart returns the DataIntervalStart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGRun) GetDataIntervalStart() time.Time {
	if o == nil || o.DataIntervalStart.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DataIntervalStart.Get()
}

// GetDataIntervalStartOk returns a tuple with the DataIntervalStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGRun) GetDataIntervalStartOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DataIntervalStart.Get(), o.DataIntervalStart.IsSet()
}

// HasDataIntervalStart returns a boolean if a field has been set.
func (o *DAGRun) HasDataIntervalStart() bool {
	if o != nil && o.DataIntervalStart.IsSet() {
		return true
	}

	return false
}

// SetDataIntervalStart gets a reference to the given NullableTime and assigns it to the DataIntervalStart field.
func (o *DAGRun) SetDataIntervalStart(v time.Time) {
	o.DataIntervalStart.Set(&v)
}
// SetDataIntervalStartNil sets the value for DataIntervalStart to be an explicit nil
func (o *DAGRun) SetDataIntervalStartNil() {
	o.DataIntervalStart.Set(nil)
}

// UnsetDataIntervalStart ensures that no value is present for DataIntervalStart, not even an explicit nil
func (o *DAGRun) UnsetDataIntervalStart() {
	o.DataIntervalStart.Unset()
}

// GetDataIntervalEnd returns the DataIntervalEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGRun) GetDataIntervalEnd() time.Time {
	if o == nil || o.DataIntervalEnd.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DataIntervalEnd.Get()
}

// GetDataIntervalEndOk returns a tuple with the DataIntervalEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGRun) GetDataIntervalEndOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DataIntervalEnd.Get(), o.DataIntervalEnd.IsSet()
}

// HasDataIntervalEnd returns a boolean if a field has been set.
func (o *DAGRun) HasDataIntervalEnd() bool {
	if o != nil && o.DataIntervalEnd.IsSet() {
		return true
	}

	return false
}

// SetDataIntervalEnd gets a reference to the given NullableTime and assigns it to the DataIntervalEnd field.
func (o *DAGRun) SetDataIntervalEnd(v time.Time) {
	o.DataIntervalEnd.Set(&v)
}
// SetDataIntervalEndNil sets the value for DataIntervalEnd to be an explicit nil
func (o *DAGRun) SetDataIntervalEndNil() {
	o.DataIntervalEnd.Set(nil)
}

// UnsetDataIntervalEnd ensures that no value is present for DataIntervalEnd, not even an explicit nil
func (o *DAGRun) UnsetDataIntervalEnd() {
	o.DataIntervalEnd.Unset()
}

// GetLastSchedulingDecision returns the LastSchedulingDecision field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGRun) GetLastSchedulingDecision() time.Time {
	if o == nil || o.LastSchedulingDecision.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSchedulingDecision.Get()
}

// GetLastSchedulingDecisionOk returns a tuple with the LastSchedulingDecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGRun) GetLastSchedulingDecisionOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastSchedulingDecision.Get(), o.LastSchedulingDecision.IsSet()
}

// HasLastSchedulingDecision returns a boolean if a field has been set.
func (o *DAGRun) HasLastSchedulingDecision() bool {
	if o != nil && o.LastSchedulingDecision.IsSet() {
		return true
	}

	return false
}

// SetLastSchedulingDecision gets a reference to the given NullableTime and assigns it to the LastSchedulingDecision field.
func (o *DAGRun) SetLastSchedulingDecision(v time.Time) {
	o.LastSchedulingDecision.Set(&v)
}
// SetLastSchedulingDecisionNil sets the value for LastSchedulingDecision to be an explicit nil
func (o *DAGRun) SetLastSchedulingDecisionNil() {
	o.LastSchedulingDecision.Set(nil)
}

// UnsetLastSchedulingDecision ensures that no value is present for LastSchedulingDecision, not even an explicit nil
func (o *DAGRun) UnsetLastSchedulingDecision() {
	o.LastSchedulingDecision.Unset()
}

// GetRunType returns the RunType field value if set, zero value otherwise.
func (o *DAGRun) GetRunType() string {
	if o == nil || o.RunType == nil {
		var ret string
		return ret
	}
	return *o.RunType
}

// GetRunTypeOk returns a tuple with the RunType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGRun) GetRunTypeOk() (*string, bool) {
	if o == nil || o.RunType == nil {
		return nil, false
	}
	return o.RunType, true
}

// HasRunType returns a boolean if a field has been set.
func (o *DAGRun) HasRunType() bool {
	if o != nil && o.RunType != nil {
		return true
	}

	return false
}

// SetRunType gets a reference to the given string and assigns it to the RunType field.
func (o *DAGRun) SetRunType(v string) {
	o.RunType = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DAGRun) GetState() DagState {
	if o == nil || o.State == nil {
		var ret DagState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGRun) GetStateOk() (*DagState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DAGRun) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given DagState and assigns it to the State field.
func (o *DAGRun) SetState(v DagState) {
	o.State = &v
}

// GetExternalTrigger returns the ExternalTrigger field value if set, zero value otherwise.
func (o *DAGRun) GetExternalTrigger() bool {
	if o == nil || o.ExternalTrigger == nil {
		var ret bool
		return ret
	}
	return *o.ExternalTrigger
}

// GetExternalTriggerOk returns a tuple with the ExternalTrigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGRun) GetExternalTriggerOk() (*bool, bool) {
	if o == nil || o.ExternalTrigger == nil {
		return nil, false
	}
	return o.ExternalTrigger, true
}

// HasExternalTrigger returns a boolean if a field has been set.
func (o *DAGRun) HasExternalTrigger() bool {
	if o != nil && o.ExternalTrigger != nil {
		return true
	}

	return false
}

// SetExternalTrigger gets a reference to the given bool and assigns it to the ExternalTrigger field.
func (o *DAGRun) SetExternalTrigger(v bool) {
	o.ExternalTrigger = &v
}

// GetConf returns the Conf field value if set, zero value otherwise.
func (o *DAGRun) GetConf() map[string]interface{} {
	if o == nil || o.Conf == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Conf
}

// GetConfOk returns a tuple with the Conf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGRun) GetConfOk() (map[string]interface{}, bool) {
	if o == nil || o.Conf == nil {
		return nil, false
	}
	return o.Conf, true
}

// HasConf returns a boolean if a field has been set.
func (o *DAGRun) HasConf() bool {
	if o != nil && o.Conf != nil {
		return true
	}

	return false
}

// SetConf gets a reference to the given map[string]interface{} and assigns it to the Conf field.
func (o *DAGRun) SetConf(v map[string]interface{}) {
	o.Conf = v
}

// GetNote returns the Note field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGRun) GetNote() string {
	if o == nil || o.Note.Get() == nil {
		var ret string
		return ret
	}
	return *o.Note.Get()
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGRun) GetNoteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Note.Get(), o.Note.IsSet()
}

// HasNote returns a boolean if a field has been set.
func (o *DAGRun) HasNote() bool {
	if o != nil && o.Note.IsSet() {
		return true
	}

	return false
}

// SetNote gets a reference to the given NullableString and assigns it to the Note field.
func (o *DAGRun) SetNote(v string) {
	o.Note.Set(&v)
}
// SetNoteNil sets the value for Note to be an explicit nil
func (o *DAGRun) SetNoteNil() {
	o.Note.Set(nil)
}

// UnsetNote ensures that no value is present for Note, not even an explicit nil
func (o *DAGRun) UnsetNote() {
	o.Note.Unset()
}

func (o DAGRun) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DagRunId.IsSet() {
		toSerialize["dag_run_id"] = o.DagRunId.Get()
	}
	if o.DagId != nil {
		toSerialize["dag_id"] = o.DagId
	}
	if o.LogicalDate.IsSet() {
		toSerialize["logical_date"] = o.LogicalDate.Get()
	}
	if o.ExecutionDate.IsSet() {
		toSerialize["execution_date"] = o.ExecutionDate.Get()
	}
	if o.StartDate.IsSet() {
		toSerialize["start_date"] = o.StartDate.Get()
	}
	if o.EndDate.IsSet() {
		toSerialize["end_date"] = o.EndDate.Get()
	}
	if o.DataIntervalStart.IsSet() {
		toSerialize["data_interval_start"] = o.DataIntervalStart.Get()
	}
	if o.DataIntervalEnd.IsSet() {
		toSerialize["data_interval_end"] = o.DataIntervalEnd.Get()
	}
	if o.LastSchedulingDecision.IsSet() {
		toSerialize["last_scheduling_decision"] = o.LastSchedulingDecision.Get()
	}
	if o.RunType != nil {
		toSerialize["run_type"] = o.RunType
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.ExternalTrigger != nil {
		toSerialize["external_trigger"] = o.ExternalTrigger
	}
	if o.Conf != nil {
		toSerialize["conf"] = o.Conf
	}
	if o.Note.IsSet() {
		toSerialize["note"] = o.Note.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDAGRun struct {
	value *DAGRun
	isSet bool
}

func (v NullableDAGRun) Get() *DAGRun {
	return v.value
}

func (v *NullableDAGRun) Set(val *DAGRun) {
	v.value = val
	v.isSet = true
}

func (v NullableDAGRun) IsSet() bool {
	return v.isSet
}

func (v *NullableDAGRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDAGRun(val *DAGRun) *NullableDAGRun {
	return &NullableDAGRun{value: val, isSet: true}
}

func (v NullableDAGRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDAGRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


