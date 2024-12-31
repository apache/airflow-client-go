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

// DAGDetail DAG details.  For details see: [airflow.models.dag.DAG](https://airflow.apache.org/docs/apache-airflow/stable/_api/airflow/models/dag/index.html#airflow.models.dag.DAG) 
type DAGDetail struct {
	// The ID of the DAG.
	DagId *string `json:"dag_id,omitempty"`
	// If the DAG is SubDAG then it is the top level DAG identifier. Otherwise, null.
	RootDagId NullableString `json:"root_dag_id,omitempty"`
	// Whether the DAG is paused.
	IsPaused NullableBool `json:"is_paused,omitempty"`
	// Whether the DAG is currently seen by the scheduler(s).  *New in version 2.1.1*  *Changed in version 2.2.0*&#58; Field is read-only. 
	IsActive NullableBool `json:"is_active,omitempty"`
	// Whether the DAG is SubDAG.
	IsSubdag *bool `json:"is_subdag,omitempty"`
	// The last time the DAG was parsed.  *New in version 2.3.0* 
	LastParsedTime NullableTime `json:"last_parsed_time,omitempty"`
	// The last time the DAG was pickled.  *New in version 2.3.0* 
	LastPickled NullableTime `json:"last_pickled,omitempty"`
	// Time when the DAG last received a refresh signal (e.g. the DAG's \"refresh\" button was clicked in the web UI)  *New in version 2.3.0* 
	LastExpired NullableTime `json:"last_expired,omitempty"`
	// Whether (one of) the scheduler is scheduling this DAG at the moment  *New in version 2.3.0* 
	SchedulerLock NullableBool `json:"scheduler_lock,omitempty"`
	// Foreign key to the latest pickle_id  *New in version 2.3.0* 
	PickleId NullableString `json:"pickle_id,omitempty"`
	DefaultView *string `json:"default_view,omitempty"`
	// The absolute path to the file.
	Fileloc *string `json:"fileloc,omitempty"`
	// The key containing the encrypted path to the file. Encryption and decryption take place only on the server. This prevents the client from reading an non-DAG file. This also ensures API extensibility, because the format of encrypted data may change. 
	FileToken *string `json:"file_token,omitempty"`
	Owners []string `json:"owners,omitempty"`
	// User-provided DAG description, which can consist of several sentences or paragraphs that describe DAG contents. 
	Description NullableString `json:"description,omitempty"`
	ScheduleInterval NullableScheduleInterval `json:"schedule_interval,omitempty"`
	// Timetable/Schedule Interval description.  *New in version 2.3.0* 
	TimetableDescription NullableString `json:"timetable_description,omitempty"`
	// List of tags.
	Tags []Tag `json:"tags,omitempty"`
	// Maximum number of active tasks that can be run on the DAG  *New in version 2.3.0* 
	MaxActiveTasks NullableInt32 `json:"max_active_tasks,omitempty"`
	// Maximum number of active DAG runs for the DAG  *New in version 2.3.0* 
	MaxActiveRuns NullableInt32 `json:"max_active_runs,omitempty"`
	// Whether the DAG has task concurrency limits  *New in version 2.3.0* 
	HasTaskConcurrencyLimits NullableBool `json:"has_task_concurrency_limits,omitempty"`
	// Whether the DAG has import errors  *New in version 2.3.0* 
	HasImportErrors NullableBool `json:"has_import_errors,omitempty"`
	// The logical date of the next dag run.  *New in version 2.3.0* 
	NextDagrun NullableTime `json:"next_dagrun,omitempty"`
	// The start of the interval of the next dag run.  *New in version 2.3.0* 
	NextDagrunDataIntervalStart NullableTime `json:"next_dagrun_data_interval_start,omitempty"`
	// The end of the interval of the next dag run.  *New in version 2.3.0* 
	NextDagrunDataIntervalEnd NullableTime `json:"next_dagrun_data_interval_end,omitempty"`
	// Earliest time at which this ``next_dagrun`` can be created.  *New in version 2.3.0* 
	NextDagrunCreateAfter NullableTime `json:"next_dagrun_create_after,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
	Catchup *bool `json:"catchup,omitempty"`
	Orientation *string `json:"orientation,omitempty"`
	Concurrency *float32 `json:"concurrency,omitempty"`
	// The DAG's start date.  *Changed in version 2.0.1*&#58; Field becomes nullable. 
	StartDate NullableTime `json:"start_date,omitempty"`
	DagRunTimeout NullableTimeDelta `json:"dag_run_timeout,omitempty"`
	DocMd NullableString `json:"doc_md,omitempty"`
	// User-specified DAG params.  *New in version 2.0.1* 
	Params map[string]interface{} `json:"params,omitempty"`
	// The DAG's end date.  *New in version 2.3.0*. 
	EndDate NullableTime `json:"end_date,omitempty"`
	// Whether the DAG is paused upon creation.  *New in version 2.3.0* 
	IsPausedUponCreation NullableBool `json:"is_paused_upon_creation,omitempty"`
	// The last time the DAG was parsed.  *New in version 2.3.0* 
	LastParsed NullableTime `json:"last_parsed,omitempty"`
	// The template search path.  *New in version 2.3.0* 
	TemplateSearchPath []string `json:"template_search_path,omitempty"`
	// Whether to render templates as native Python objects.  *New in version 2.3.0* 
	RenderTemplateAsNativeObj NullableBool `json:"render_template_as_native_obj,omitempty"`
}

// NewDAGDetail instantiates a new DAGDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDAGDetail() *DAGDetail {
	this := DAGDetail{}
	return &this
}

// NewDAGDetailWithDefaults instantiates a new DAGDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDAGDetailWithDefaults() *DAGDetail {
	this := DAGDetail{}
	return &this
}

// GetDagId returns the DagId field value if set, zero value otherwise.
func (o *DAGDetail) GetDagId() string {
	if o == nil || o.DagId == nil {
		var ret string
		return ret
	}
	return *o.DagId
}

// GetDagIdOk returns a tuple with the DagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGDetail) GetDagIdOk() (*string, bool) {
	if o == nil || o.DagId == nil {
		return nil, false
	}
	return o.DagId, true
}

// HasDagId returns a boolean if a field has been set.
func (o *DAGDetail) HasDagId() bool {
	if o != nil && o.DagId != nil {
		return true
	}

	return false
}

// SetDagId gets a reference to the given string and assigns it to the DagId field.
func (o *DAGDetail) SetDagId(v string) {
	o.DagId = &v
}

// GetRootDagId returns the RootDagId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetRootDagId() string {
	if o == nil || o.RootDagId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RootDagId.Get()
}

// GetRootDagIdOk returns a tuple with the RootDagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetRootDagIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RootDagId.Get(), o.RootDagId.IsSet()
}

// HasRootDagId returns a boolean if a field has been set.
func (o *DAGDetail) HasRootDagId() bool {
	if o != nil && o.RootDagId.IsSet() {
		return true
	}

	return false
}

// SetRootDagId gets a reference to the given NullableString and assigns it to the RootDagId field.
func (o *DAGDetail) SetRootDagId(v string) {
	o.RootDagId.Set(&v)
}
// SetRootDagIdNil sets the value for RootDagId to be an explicit nil
func (o *DAGDetail) SetRootDagIdNil() {
	o.RootDagId.Set(nil)
}

// UnsetRootDagId ensures that no value is present for RootDagId, not even an explicit nil
func (o *DAGDetail) UnsetRootDagId() {
	o.RootDagId.Unset()
}

// GetIsPaused returns the IsPaused field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetIsPaused() bool {
	if o == nil || o.IsPaused.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsPaused.Get()
}

// GetIsPausedOk returns a tuple with the IsPaused field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetIsPausedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsPaused.Get(), o.IsPaused.IsSet()
}

// HasIsPaused returns a boolean if a field has been set.
func (o *DAGDetail) HasIsPaused() bool {
	if o != nil && o.IsPaused.IsSet() {
		return true
	}

	return false
}

// SetIsPaused gets a reference to the given NullableBool and assigns it to the IsPaused field.
func (o *DAGDetail) SetIsPaused(v bool) {
	o.IsPaused.Set(&v)
}
// SetIsPausedNil sets the value for IsPaused to be an explicit nil
func (o *DAGDetail) SetIsPausedNil() {
	o.IsPaused.Set(nil)
}

// UnsetIsPaused ensures that no value is present for IsPaused, not even an explicit nil
func (o *DAGDetail) UnsetIsPaused() {
	o.IsPaused.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetIsActive() bool {
	if o == nil || o.IsActive.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsActive.Get()
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetIsActiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsActive.Get(), o.IsActive.IsSet()
}

// HasIsActive returns a boolean if a field has been set.
func (o *DAGDetail) HasIsActive() bool {
	if o != nil && o.IsActive.IsSet() {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given NullableBool and assigns it to the IsActive field.
func (o *DAGDetail) SetIsActive(v bool) {
	o.IsActive.Set(&v)
}
// SetIsActiveNil sets the value for IsActive to be an explicit nil
func (o *DAGDetail) SetIsActiveNil() {
	o.IsActive.Set(nil)
}

// UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
func (o *DAGDetail) UnsetIsActive() {
	o.IsActive.Unset()
}

// GetIsSubdag returns the IsSubdag field value if set, zero value otherwise.
func (o *DAGDetail) GetIsSubdag() bool {
	if o == nil || o.IsSubdag == nil {
		var ret bool
		return ret
	}
	return *o.IsSubdag
}

// GetIsSubdagOk returns a tuple with the IsSubdag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGDetail) GetIsSubdagOk() (*bool, bool) {
	if o == nil || o.IsSubdag == nil {
		return nil, false
	}
	return o.IsSubdag, true
}

// HasIsSubdag returns a boolean if a field has been set.
func (o *DAGDetail) HasIsSubdag() bool {
	if o != nil && o.IsSubdag != nil {
		return true
	}

	return false
}

// SetIsSubdag gets a reference to the given bool and assigns it to the IsSubdag field.
func (o *DAGDetail) SetIsSubdag(v bool) {
	o.IsSubdag = &v
}

// GetLastParsedTime returns the LastParsedTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetLastParsedTime() time.Time {
	if o == nil || o.LastParsedTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastParsedTime.Get()
}

// GetLastParsedTimeOk returns a tuple with the LastParsedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetLastParsedTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastParsedTime.Get(), o.LastParsedTime.IsSet()
}

// HasLastParsedTime returns a boolean if a field has been set.
func (o *DAGDetail) HasLastParsedTime() bool {
	if o != nil && o.LastParsedTime.IsSet() {
		return true
	}

	return false
}

// SetLastParsedTime gets a reference to the given NullableTime and assigns it to the LastParsedTime field.
func (o *DAGDetail) SetLastParsedTime(v time.Time) {
	o.LastParsedTime.Set(&v)
}
// SetLastParsedTimeNil sets the value for LastParsedTime to be an explicit nil
func (o *DAGDetail) SetLastParsedTimeNil() {
	o.LastParsedTime.Set(nil)
}

// UnsetLastParsedTime ensures that no value is present for LastParsedTime, not even an explicit nil
func (o *DAGDetail) UnsetLastParsedTime() {
	o.LastParsedTime.Unset()
}

// GetLastPickled returns the LastPickled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetLastPickled() time.Time {
	if o == nil || o.LastPickled.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastPickled.Get()
}

// GetLastPickledOk returns a tuple with the LastPickled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetLastPickledOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastPickled.Get(), o.LastPickled.IsSet()
}

// HasLastPickled returns a boolean if a field has been set.
func (o *DAGDetail) HasLastPickled() bool {
	if o != nil && o.LastPickled.IsSet() {
		return true
	}

	return false
}

// SetLastPickled gets a reference to the given NullableTime and assigns it to the LastPickled field.
func (o *DAGDetail) SetLastPickled(v time.Time) {
	o.LastPickled.Set(&v)
}
// SetLastPickledNil sets the value for LastPickled to be an explicit nil
func (o *DAGDetail) SetLastPickledNil() {
	o.LastPickled.Set(nil)
}

// UnsetLastPickled ensures that no value is present for LastPickled, not even an explicit nil
func (o *DAGDetail) UnsetLastPickled() {
	o.LastPickled.Unset()
}

// GetLastExpired returns the LastExpired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetLastExpired() time.Time {
	if o == nil || o.LastExpired.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastExpired.Get()
}

// GetLastExpiredOk returns a tuple with the LastExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetLastExpiredOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastExpired.Get(), o.LastExpired.IsSet()
}

// HasLastExpired returns a boolean if a field has been set.
func (o *DAGDetail) HasLastExpired() bool {
	if o != nil && o.LastExpired.IsSet() {
		return true
	}

	return false
}

// SetLastExpired gets a reference to the given NullableTime and assigns it to the LastExpired field.
func (o *DAGDetail) SetLastExpired(v time.Time) {
	o.LastExpired.Set(&v)
}
// SetLastExpiredNil sets the value for LastExpired to be an explicit nil
func (o *DAGDetail) SetLastExpiredNil() {
	o.LastExpired.Set(nil)
}

// UnsetLastExpired ensures that no value is present for LastExpired, not even an explicit nil
func (o *DAGDetail) UnsetLastExpired() {
	o.LastExpired.Unset()
}

// GetSchedulerLock returns the SchedulerLock field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetSchedulerLock() bool {
	if o == nil || o.SchedulerLock.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SchedulerLock.Get()
}

// GetSchedulerLockOk returns a tuple with the SchedulerLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetSchedulerLockOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SchedulerLock.Get(), o.SchedulerLock.IsSet()
}

// HasSchedulerLock returns a boolean if a field has been set.
func (o *DAGDetail) HasSchedulerLock() bool {
	if o != nil && o.SchedulerLock.IsSet() {
		return true
	}

	return false
}

// SetSchedulerLock gets a reference to the given NullableBool and assigns it to the SchedulerLock field.
func (o *DAGDetail) SetSchedulerLock(v bool) {
	o.SchedulerLock.Set(&v)
}
// SetSchedulerLockNil sets the value for SchedulerLock to be an explicit nil
func (o *DAGDetail) SetSchedulerLockNil() {
	o.SchedulerLock.Set(nil)
}

// UnsetSchedulerLock ensures that no value is present for SchedulerLock, not even an explicit nil
func (o *DAGDetail) UnsetSchedulerLock() {
	o.SchedulerLock.Unset()
}

// GetPickleId returns the PickleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetPickleId() string {
	if o == nil || o.PickleId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PickleId.Get()
}

// GetPickleIdOk returns a tuple with the PickleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetPickleIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PickleId.Get(), o.PickleId.IsSet()
}

// HasPickleId returns a boolean if a field has been set.
func (o *DAGDetail) HasPickleId() bool {
	if o != nil && o.PickleId.IsSet() {
		return true
	}

	return false
}

// SetPickleId gets a reference to the given NullableString and assigns it to the PickleId field.
func (o *DAGDetail) SetPickleId(v string) {
	o.PickleId.Set(&v)
}
// SetPickleIdNil sets the value for PickleId to be an explicit nil
func (o *DAGDetail) SetPickleIdNil() {
	o.PickleId.Set(nil)
}

// UnsetPickleId ensures that no value is present for PickleId, not even an explicit nil
func (o *DAGDetail) UnsetPickleId() {
	o.PickleId.Unset()
}

// GetDefaultView returns the DefaultView field value if set, zero value otherwise.
func (o *DAGDetail) GetDefaultView() string {
	if o == nil || o.DefaultView == nil {
		var ret string
		return ret
	}
	return *o.DefaultView
}

// GetDefaultViewOk returns a tuple with the DefaultView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGDetail) GetDefaultViewOk() (*string, bool) {
	if o == nil || o.DefaultView == nil {
		return nil, false
	}
	return o.DefaultView, true
}

// HasDefaultView returns a boolean if a field has been set.
func (o *DAGDetail) HasDefaultView() bool {
	if o != nil && o.DefaultView != nil {
		return true
	}

	return false
}

// SetDefaultView gets a reference to the given string and assigns it to the DefaultView field.
func (o *DAGDetail) SetDefaultView(v string) {
	o.DefaultView = &v
}

// GetFileloc returns the Fileloc field value if set, zero value otherwise.
func (o *DAGDetail) GetFileloc() string {
	if o == nil || o.Fileloc == nil {
		var ret string
		return ret
	}
	return *o.Fileloc
}

// GetFilelocOk returns a tuple with the Fileloc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGDetail) GetFilelocOk() (*string, bool) {
	if o == nil || o.Fileloc == nil {
		return nil, false
	}
	return o.Fileloc, true
}

// HasFileloc returns a boolean if a field has been set.
func (o *DAGDetail) HasFileloc() bool {
	if o != nil && o.Fileloc != nil {
		return true
	}

	return false
}

// SetFileloc gets a reference to the given string and assigns it to the Fileloc field.
func (o *DAGDetail) SetFileloc(v string) {
	o.Fileloc = &v
}

// GetFileToken returns the FileToken field value if set, zero value otherwise.
func (o *DAGDetail) GetFileToken() string {
	if o == nil || o.FileToken == nil {
		var ret string
		return ret
	}
	return *o.FileToken
}

// GetFileTokenOk returns a tuple with the FileToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGDetail) GetFileTokenOk() (*string, bool) {
	if o == nil || o.FileToken == nil {
		return nil, false
	}
	return o.FileToken, true
}

// HasFileToken returns a boolean if a field has been set.
func (o *DAGDetail) HasFileToken() bool {
	if o != nil && o.FileToken != nil {
		return true
	}

	return false
}

// SetFileToken gets a reference to the given string and assigns it to the FileToken field.
func (o *DAGDetail) SetFileToken(v string) {
	o.FileToken = &v
}

// GetOwners returns the Owners field value if set, zero value otherwise.
func (o *DAGDetail) GetOwners() []string {
	if o == nil || o.Owners == nil {
		var ret []string
		return ret
	}
	return o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGDetail) GetOwnersOk() ([]string, bool) {
	if o == nil || o.Owners == nil {
		return nil, false
	}
	return o.Owners, true
}

// HasOwners returns a boolean if a field has been set.
func (o *DAGDetail) HasOwners() bool {
	if o != nil && o.Owners != nil {
		return true
	}

	return false
}

// SetOwners gets a reference to the given []string and assigns it to the Owners field.
func (o *DAGDetail) SetOwners(v []string) {
	o.Owners = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DAGDetail) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *DAGDetail) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *DAGDetail) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *DAGDetail) UnsetDescription() {
	o.Description.Unset()
}

// GetScheduleInterval returns the ScheduleInterval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetScheduleInterval() ScheduleInterval {
	if o == nil || o.ScheduleInterval.Get() == nil {
		var ret ScheduleInterval
		return ret
	}
	return *o.ScheduleInterval.Get()
}

// GetScheduleIntervalOk returns a tuple with the ScheduleInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetScheduleIntervalOk() (*ScheduleInterval, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ScheduleInterval.Get(), o.ScheduleInterval.IsSet()
}

// HasScheduleInterval returns a boolean if a field has been set.
func (o *DAGDetail) HasScheduleInterval() bool {
	if o != nil && o.ScheduleInterval.IsSet() {
		return true
	}

	return false
}

// SetScheduleInterval gets a reference to the given NullableScheduleInterval and assigns it to the ScheduleInterval field.
func (o *DAGDetail) SetScheduleInterval(v ScheduleInterval) {
	o.ScheduleInterval.Set(&v)
}
// SetScheduleIntervalNil sets the value for ScheduleInterval to be an explicit nil
func (o *DAGDetail) SetScheduleIntervalNil() {
	o.ScheduleInterval.Set(nil)
}

// UnsetScheduleInterval ensures that no value is present for ScheduleInterval, not even an explicit nil
func (o *DAGDetail) UnsetScheduleInterval() {
	o.ScheduleInterval.Unset()
}

// GetTimetableDescription returns the TimetableDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetTimetableDescription() string {
	if o == nil || o.TimetableDescription.Get() == nil {
		var ret string
		return ret
	}
	return *o.TimetableDescription.Get()
}

// GetTimetableDescriptionOk returns a tuple with the TimetableDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetTimetableDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TimetableDescription.Get(), o.TimetableDescription.IsSet()
}

// HasTimetableDescription returns a boolean if a field has been set.
func (o *DAGDetail) HasTimetableDescription() bool {
	if o != nil && o.TimetableDescription.IsSet() {
		return true
	}

	return false
}

// SetTimetableDescription gets a reference to the given NullableString and assigns it to the TimetableDescription field.
func (o *DAGDetail) SetTimetableDescription(v string) {
	o.TimetableDescription.Set(&v)
}
// SetTimetableDescriptionNil sets the value for TimetableDescription to be an explicit nil
func (o *DAGDetail) SetTimetableDescriptionNil() {
	o.TimetableDescription.Set(nil)
}

// UnsetTimetableDescription ensures that no value is present for TimetableDescription, not even an explicit nil
func (o *DAGDetail) UnsetTimetableDescription() {
	o.TimetableDescription.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetTags() []Tag {
	if o == nil  {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetTagsOk() ([]Tag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DAGDetail) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *DAGDetail) SetTags(v []Tag) {
	o.Tags = v
}

// GetMaxActiveTasks returns the MaxActiveTasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetMaxActiveTasks() int32 {
	if o == nil || o.MaxActiveTasks.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaxActiveTasks.Get()
}

// GetMaxActiveTasksOk returns a tuple with the MaxActiveTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetMaxActiveTasksOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxActiveTasks.Get(), o.MaxActiveTasks.IsSet()
}

// HasMaxActiveTasks returns a boolean if a field has been set.
func (o *DAGDetail) HasMaxActiveTasks() bool {
	if o != nil && o.MaxActiveTasks.IsSet() {
		return true
	}

	return false
}

// SetMaxActiveTasks gets a reference to the given NullableInt32 and assigns it to the MaxActiveTasks field.
func (o *DAGDetail) SetMaxActiveTasks(v int32) {
	o.MaxActiveTasks.Set(&v)
}
// SetMaxActiveTasksNil sets the value for MaxActiveTasks to be an explicit nil
func (o *DAGDetail) SetMaxActiveTasksNil() {
	o.MaxActiveTasks.Set(nil)
}

// UnsetMaxActiveTasks ensures that no value is present for MaxActiveTasks, not even an explicit nil
func (o *DAGDetail) UnsetMaxActiveTasks() {
	o.MaxActiveTasks.Unset()
}

// GetMaxActiveRuns returns the MaxActiveRuns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetMaxActiveRuns() int32 {
	if o == nil || o.MaxActiveRuns.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaxActiveRuns.Get()
}

// GetMaxActiveRunsOk returns a tuple with the MaxActiveRuns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetMaxActiveRunsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxActiveRuns.Get(), o.MaxActiveRuns.IsSet()
}

// HasMaxActiveRuns returns a boolean if a field has been set.
func (o *DAGDetail) HasMaxActiveRuns() bool {
	if o != nil && o.MaxActiveRuns.IsSet() {
		return true
	}

	return false
}

// SetMaxActiveRuns gets a reference to the given NullableInt32 and assigns it to the MaxActiveRuns field.
func (o *DAGDetail) SetMaxActiveRuns(v int32) {
	o.MaxActiveRuns.Set(&v)
}
// SetMaxActiveRunsNil sets the value for MaxActiveRuns to be an explicit nil
func (o *DAGDetail) SetMaxActiveRunsNil() {
	o.MaxActiveRuns.Set(nil)
}

// UnsetMaxActiveRuns ensures that no value is present for MaxActiveRuns, not even an explicit nil
func (o *DAGDetail) UnsetMaxActiveRuns() {
	o.MaxActiveRuns.Unset()
}

// GetHasTaskConcurrencyLimits returns the HasTaskConcurrencyLimits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetHasTaskConcurrencyLimits() bool {
	if o == nil || o.HasTaskConcurrencyLimits.Get() == nil {
		var ret bool
		return ret
	}
	return *o.HasTaskConcurrencyLimits.Get()
}

// GetHasTaskConcurrencyLimitsOk returns a tuple with the HasTaskConcurrencyLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetHasTaskConcurrencyLimitsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HasTaskConcurrencyLimits.Get(), o.HasTaskConcurrencyLimits.IsSet()
}

// HasHasTaskConcurrencyLimits returns a boolean if a field has been set.
func (o *DAGDetail) HasHasTaskConcurrencyLimits() bool {
	if o != nil && o.HasTaskConcurrencyLimits.IsSet() {
		return true
	}

	return false
}

// SetHasTaskConcurrencyLimits gets a reference to the given NullableBool and assigns it to the HasTaskConcurrencyLimits field.
func (o *DAGDetail) SetHasTaskConcurrencyLimits(v bool) {
	o.HasTaskConcurrencyLimits.Set(&v)
}
// SetHasTaskConcurrencyLimitsNil sets the value for HasTaskConcurrencyLimits to be an explicit nil
func (o *DAGDetail) SetHasTaskConcurrencyLimitsNil() {
	o.HasTaskConcurrencyLimits.Set(nil)
}

// UnsetHasTaskConcurrencyLimits ensures that no value is present for HasTaskConcurrencyLimits, not even an explicit nil
func (o *DAGDetail) UnsetHasTaskConcurrencyLimits() {
	o.HasTaskConcurrencyLimits.Unset()
}

// GetHasImportErrors returns the HasImportErrors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetHasImportErrors() bool {
	if o == nil || o.HasImportErrors.Get() == nil {
		var ret bool
		return ret
	}
	return *o.HasImportErrors.Get()
}

// GetHasImportErrorsOk returns a tuple with the HasImportErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetHasImportErrorsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HasImportErrors.Get(), o.HasImportErrors.IsSet()
}

// HasHasImportErrors returns a boolean if a field has been set.
func (o *DAGDetail) HasHasImportErrors() bool {
	if o != nil && o.HasImportErrors.IsSet() {
		return true
	}

	return false
}

// SetHasImportErrors gets a reference to the given NullableBool and assigns it to the HasImportErrors field.
func (o *DAGDetail) SetHasImportErrors(v bool) {
	o.HasImportErrors.Set(&v)
}
// SetHasImportErrorsNil sets the value for HasImportErrors to be an explicit nil
func (o *DAGDetail) SetHasImportErrorsNil() {
	o.HasImportErrors.Set(nil)
}

// UnsetHasImportErrors ensures that no value is present for HasImportErrors, not even an explicit nil
func (o *DAGDetail) UnsetHasImportErrors() {
	o.HasImportErrors.Unset()
}

// GetNextDagrun returns the NextDagrun field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetNextDagrun() time.Time {
	if o == nil || o.NextDagrun.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.NextDagrun.Get()
}

// GetNextDagrunOk returns a tuple with the NextDagrun field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetNextDagrunOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextDagrun.Get(), o.NextDagrun.IsSet()
}

// HasNextDagrun returns a boolean if a field has been set.
func (o *DAGDetail) HasNextDagrun() bool {
	if o != nil && o.NextDagrun.IsSet() {
		return true
	}

	return false
}

// SetNextDagrun gets a reference to the given NullableTime and assigns it to the NextDagrun field.
func (o *DAGDetail) SetNextDagrun(v time.Time) {
	o.NextDagrun.Set(&v)
}
// SetNextDagrunNil sets the value for NextDagrun to be an explicit nil
func (o *DAGDetail) SetNextDagrunNil() {
	o.NextDagrun.Set(nil)
}

// UnsetNextDagrun ensures that no value is present for NextDagrun, not even an explicit nil
func (o *DAGDetail) UnsetNextDagrun() {
	o.NextDagrun.Unset()
}

// GetNextDagrunDataIntervalStart returns the NextDagrunDataIntervalStart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetNextDagrunDataIntervalStart() time.Time {
	if o == nil || o.NextDagrunDataIntervalStart.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.NextDagrunDataIntervalStart.Get()
}

// GetNextDagrunDataIntervalStartOk returns a tuple with the NextDagrunDataIntervalStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetNextDagrunDataIntervalStartOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextDagrunDataIntervalStart.Get(), o.NextDagrunDataIntervalStart.IsSet()
}

// HasNextDagrunDataIntervalStart returns a boolean if a field has been set.
func (o *DAGDetail) HasNextDagrunDataIntervalStart() bool {
	if o != nil && o.NextDagrunDataIntervalStart.IsSet() {
		return true
	}

	return false
}

// SetNextDagrunDataIntervalStart gets a reference to the given NullableTime and assigns it to the NextDagrunDataIntervalStart field.
func (o *DAGDetail) SetNextDagrunDataIntervalStart(v time.Time) {
	o.NextDagrunDataIntervalStart.Set(&v)
}
// SetNextDagrunDataIntervalStartNil sets the value for NextDagrunDataIntervalStart to be an explicit nil
func (o *DAGDetail) SetNextDagrunDataIntervalStartNil() {
	o.NextDagrunDataIntervalStart.Set(nil)
}

// UnsetNextDagrunDataIntervalStart ensures that no value is present for NextDagrunDataIntervalStart, not even an explicit nil
func (o *DAGDetail) UnsetNextDagrunDataIntervalStart() {
	o.NextDagrunDataIntervalStart.Unset()
}

// GetNextDagrunDataIntervalEnd returns the NextDagrunDataIntervalEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetNextDagrunDataIntervalEnd() time.Time {
	if o == nil || o.NextDagrunDataIntervalEnd.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.NextDagrunDataIntervalEnd.Get()
}

// GetNextDagrunDataIntervalEndOk returns a tuple with the NextDagrunDataIntervalEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetNextDagrunDataIntervalEndOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextDagrunDataIntervalEnd.Get(), o.NextDagrunDataIntervalEnd.IsSet()
}

// HasNextDagrunDataIntervalEnd returns a boolean if a field has been set.
func (o *DAGDetail) HasNextDagrunDataIntervalEnd() bool {
	if o != nil && o.NextDagrunDataIntervalEnd.IsSet() {
		return true
	}

	return false
}

// SetNextDagrunDataIntervalEnd gets a reference to the given NullableTime and assigns it to the NextDagrunDataIntervalEnd field.
func (o *DAGDetail) SetNextDagrunDataIntervalEnd(v time.Time) {
	o.NextDagrunDataIntervalEnd.Set(&v)
}
// SetNextDagrunDataIntervalEndNil sets the value for NextDagrunDataIntervalEnd to be an explicit nil
func (o *DAGDetail) SetNextDagrunDataIntervalEndNil() {
	o.NextDagrunDataIntervalEnd.Set(nil)
}

// UnsetNextDagrunDataIntervalEnd ensures that no value is present for NextDagrunDataIntervalEnd, not even an explicit nil
func (o *DAGDetail) UnsetNextDagrunDataIntervalEnd() {
	o.NextDagrunDataIntervalEnd.Unset()
}

// GetNextDagrunCreateAfter returns the NextDagrunCreateAfter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetNextDagrunCreateAfter() time.Time {
	if o == nil || o.NextDagrunCreateAfter.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.NextDagrunCreateAfter.Get()
}

// GetNextDagrunCreateAfterOk returns a tuple with the NextDagrunCreateAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetNextDagrunCreateAfterOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextDagrunCreateAfter.Get(), o.NextDagrunCreateAfter.IsSet()
}

// HasNextDagrunCreateAfter returns a boolean if a field has been set.
func (o *DAGDetail) HasNextDagrunCreateAfter() bool {
	if o != nil && o.NextDagrunCreateAfter.IsSet() {
		return true
	}

	return false
}

// SetNextDagrunCreateAfter gets a reference to the given NullableTime and assigns it to the NextDagrunCreateAfter field.
func (o *DAGDetail) SetNextDagrunCreateAfter(v time.Time) {
	o.NextDagrunCreateAfter.Set(&v)
}
// SetNextDagrunCreateAfterNil sets the value for NextDagrunCreateAfter to be an explicit nil
func (o *DAGDetail) SetNextDagrunCreateAfterNil() {
	o.NextDagrunCreateAfter.Set(nil)
}

// UnsetNextDagrunCreateAfter ensures that no value is present for NextDagrunCreateAfter, not even an explicit nil
func (o *DAGDetail) UnsetNextDagrunCreateAfter() {
	o.NextDagrunCreateAfter.Unset()
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *DAGDetail) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGDetail) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *DAGDetail) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *DAGDetail) SetTimezone(v string) {
	o.Timezone = &v
}

// GetCatchup returns the Catchup field value if set, zero value otherwise.
func (o *DAGDetail) GetCatchup() bool {
	if o == nil || o.Catchup == nil {
		var ret bool
		return ret
	}
	return *o.Catchup
}

// GetCatchupOk returns a tuple with the Catchup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGDetail) GetCatchupOk() (*bool, bool) {
	if o == nil || o.Catchup == nil {
		return nil, false
	}
	return o.Catchup, true
}

// HasCatchup returns a boolean if a field has been set.
func (o *DAGDetail) HasCatchup() bool {
	if o != nil && o.Catchup != nil {
		return true
	}

	return false
}

// SetCatchup gets a reference to the given bool and assigns it to the Catchup field.
func (o *DAGDetail) SetCatchup(v bool) {
	o.Catchup = &v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *DAGDetail) GetOrientation() string {
	if o == nil || o.Orientation == nil {
		var ret string
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGDetail) GetOrientationOk() (*string, bool) {
	if o == nil || o.Orientation == nil {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *DAGDetail) HasOrientation() bool {
	if o != nil && o.Orientation != nil {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given string and assigns it to the Orientation field.
func (o *DAGDetail) SetOrientation(v string) {
	o.Orientation = &v
}

// GetConcurrency returns the Concurrency field value if set, zero value otherwise.
func (o *DAGDetail) GetConcurrency() float32 {
	if o == nil || o.Concurrency == nil {
		var ret float32
		return ret
	}
	return *o.Concurrency
}

// GetConcurrencyOk returns a tuple with the Concurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGDetail) GetConcurrencyOk() (*float32, bool) {
	if o == nil || o.Concurrency == nil {
		return nil, false
	}
	return o.Concurrency, true
}

// HasConcurrency returns a boolean if a field has been set.
func (o *DAGDetail) HasConcurrency() bool {
	if o != nil && o.Concurrency != nil {
		return true
	}

	return false
}

// SetConcurrency gets a reference to the given float32 and assigns it to the Concurrency field.
func (o *DAGDetail) SetConcurrency(v float32) {
	o.Concurrency = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetStartDate() time.Time {
	if o == nil || o.StartDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetStartDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *DAGDetail) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableTime and assigns it to the StartDate field.
func (o *DAGDetail) SetStartDate(v time.Time) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *DAGDetail) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *DAGDetail) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetDagRunTimeout returns the DagRunTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetDagRunTimeout() TimeDelta {
	if o == nil || o.DagRunTimeout.Get() == nil {
		var ret TimeDelta
		return ret
	}
	return *o.DagRunTimeout.Get()
}

// GetDagRunTimeoutOk returns a tuple with the DagRunTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetDagRunTimeoutOk() (*TimeDelta, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DagRunTimeout.Get(), o.DagRunTimeout.IsSet()
}

// HasDagRunTimeout returns a boolean if a field has been set.
func (o *DAGDetail) HasDagRunTimeout() bool {
	if o != nil && o.DagRunTimeout.IsSet() {
		return true
	}

	return false
}

// SetDagRunTimeout gets a reference to the given NullableTimeDelta and assigns it to the DagRunTimeout field.
func (o *DAGDetail) SetDagRunTimeout(v TimeDelta) {
	o.DagRunTimeout.Set(&v)
}
// SetDagRunTimeoutNil sets the value for DagRunTimeout to be an explicit nil
func (o *DAGDetail) SetDagRunTimeoutNil() {
	o.DagRunTimeout.Set(nil)
}

// UnsetDagRunTimeout ensures that no value is present for DagRunTimeout, not even an explicit nil
func (o *DAGDetail) UnsetDagRunTimeout() {
	o.DagRunTimeout.Unset()
}

// GetDocMd returns the DocMd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetDocMd() string {
	if o == nil || o.DocMd.Get() == nil {
		var ret string
		return ret
	}
	return *o.DocMd.Get()
}

// GetDocMdOk returns a tuple with the DocMd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetDocMdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DocMd.Get(), o.DocMd.IsSet()
}

// HasDocMd returns a boolean if a field has been set.
func (o *DAGDetail) HasDocMd() bool {
	if o != nil && o.DocMd.IsSet() {
		return true
	}

	return false
}

// SetDocMd gets a reference to the given NullableString and assigns it to the DocMd field.
func (o *DAGDetail) SetDocMd(v string) {
	o.DocMd.Set(&v)
}
// SetDocMdNil sets the value for DocMd to be an explicit nil
func (o *DAGDetail) SetDocMdNil() {
	o.DocMd.Set(nil)
}

// UnsetDocMd ensures that no value is present for DocMd, not even an explicit nil
func (o *DAGDetail) UnsetDocMd() {
	o.DocMd.Unset()
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *DAGDetail) GetParams() map[string]interface{} {
	if o == nil || o.Params == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGDetail) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *DAGDetail) HasParams() bool {
	if o != nil && o.Params != nil {
		return true
	}

	return false
}

// SetParams gets a reference to the given map[string]interface{} and assigns it to the Params field.
func (o *DAGDetail) SetParams(v map[string]interface{}) {
	o.Params = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetEndDate() time.Time {
	if o == nil || o.EndDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetEndDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *DAGDetail) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableTime and assigns it to the EndDate field.
func (o *DAGDetail) SetEndDate(v time.Time) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *DAGDetail) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *DAGDetail) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetIsPausedUponCreation returns the IsPausedUponCreation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetIsPausedUponCreation() bool {
	if o == nil || o.IsPausedUponCreation.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsPausedUponCreation.Get()
}

// GetIsPausedUponCreationOk returns a tuple with the IsPausedUponCreation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetIsPausedUponCreationOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsPausedUponCreation.Get(), o.IsPausedUponCreation.IsSet()
}

// HasIsPausedUponCreation returns a boolean if a field has been set.
func (o *DAGDetail) HasIsPausedUponCreation() bool {
	if o != nil && o.IsPausedUponCreation.IsSet() {
		return true
	}

	return false
}

// SetIsPausedUponCreation gets a reference to the given NullableBool and assigns it to the IsPausedUponCreation field.
func (o *DAGDetail) SetIsPausedUponCreation(v bool) {
	o.IsPausedUponCreation.Set(&v)
}
// SetIsPausedUponCreationNil sets the value for IsPausedUponCreation to be an explicit nil
func (o *DAGDetail) SetIsPausedUponCreationNil() {
	o.IsPausedUponCreation.Set(nil)
}

// UnsetIsPausedUponCreation ensures that no value is present for IsPausedUponCreation, not even an explicit nil
func (o *DAGDetail) UnsetIsPausedUponCreation() {
	o.IsPausedUponCreation.Unset()
}

// GetLastParsed returns the LastParsed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetLastParsed() time.Time {
	if o == nil || o.LastParsed.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastParsed.Get()
}

// GetLastParsedOk returns a tuple with the LastParsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetLastParsedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastParsed.Get(), o.LastParsed.IsSet()
}

// HasLastParsed returns a boolean if a field has been set.
func (o *DAGDetail) HasLastParsed() bool {
	if o != nil && o.LastParsed.IsSet() {
		return true
	}

	return false
}

// SetLastParsed gets a reference to the given NullableTime and assigns it to the LastParsed field.
func (o *DAGDetail) SetLastParsed(v time.Time) {
	o.LastParsed.Set(&v)
}
// SetLastParsedNil sets the value for LastParsed to be an explicit nil
func (o *DAGDetail) SetLastParsedNil() {
	o.LastParsed.Set(nil)
}

// UnsetLastParsed ensures that no value is present for LastParsed, not even an explicit nil
func (o *DAGDetail) UnsetLastParsed() {
	o.LastParsed.Unset()
}

// GetTemplateSearchPath returns the TemplateSearchPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetTemplateSearchPath() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.TemplateSearchPath
}

// GetTemplateSearchPathOk returns a tuple with the TemplateSearchPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetTemplateSearchPathOk() ([]string, bool) {
	if o == nil || o.TemplateSearchPath == nil {
		return nil, false
	}
	return o.TemplateSearchPath, true
}

// HasTemplateSearchPath returns a boolean if a field has been set.
func (o *DAGDetail) HasTemplateSearchPath() bool {
	if o != nil && o.TemplateSearchPath != nil {
		return true
	}

	return false
}

// SetTemplateSearchPath gets a reference to the given []string and assigns it to the TemplateSearchPath field.
func (o *DAGDetail) SetTemplateSearchPath(v []string) {
	o.TemplateSearchPath = v
}

// GetRenderTemplateAsNativeObj returns the RenderTemplateAsNativeObj field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetail) GetRenderTemplateAsNativeObj() bool {
	if o == nil || o.RenderTemplateAsNativeObj.Get() == nil {
		var ret bool
		return ret
	}
	return *o.RenderTemplateAsNativeObj.Get()
}

// GetRenderTemplateAsNativeObjOk returns a tuple with the RenderTemplateAsNativeObj field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetail) GetRenderTemplateAsNativeObjOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RenderTemplateAsNativeObj.Get(), o.RenderTemplateAsNativeObj.IsSet()
}

// HasRenderTemplateAsNativeObj returns a boolean if a field has been set.
func (o *DAGDetail) HasRenderTemplateAsNativeObj() bool {
	if o != nil && o.RenderTemplateAsNativeObj.IsSet() {
		return true
	}

	return false
}

// SetRenderTemplateAsNativeObj gets a reference to the given NullableBool and assigns it to the RenderTemplateAsNativeObj field.
func (o *DAGDetail) SetRenderTemplateAsNativeObj(v bool) {
	o.RenderTemplateAsNativeObj.Set(&v)
}
// SetRenderTemplateAsNativeObjNil sets the value for RenderTemplateAsNativeObj to be an explicit nil
func (o *DAGDetail) SetRenderTemplateAsNativeObjNil() {
	o.RenderTemplateAsNativeObj.Set(nil)
}

// UnsetRenderTemplateAsNativeObj ensures that no value is present for RenderTemplateAsNativeObj, not even an explicit nil
func (o *DAGDetail) UnsetRenderTemplateAsNativeObj() {
	o.RenderTemplateAsNativeObj.Unset()
}

func (o DAGDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DagId != nil {
		toSerialize["dag_id"] = o.DagId
	}
	if o.RootDagId.IsSet() {
		toSerialize["root_dag_id"] = o.RootDagId.Get()
	}
	if o.IsPaused.IsSet() {
		toSerialize["is_paused"] = o.IsPaused.Get()
	}
	if o.IsActive.IsSet() {
		toSerialize["is_active"] = o.IsActive.Get()
	}
	if o.IsSubdag != nil {
		toSerialize["is_subdag"] = o.IsSubdag
	}
	if o.LastParsedTime.IsSet() {
		toSerialize["last_parsed_time"] = o.LastParsedTime.Get()
	}
	if o.LastPickled.IsSet() {
		toSerialize["last_pickled"] = o.LastPickled.Get()
	}
	if o.LastExpired.IsSet() {
		toSerialize["last_expired"] = o.LastExpired.Get()
	}
	if o.SchedulerLock.IsSet() {
		toSerialize["scheduler_lock"] = o.SchedulerLock.Get()
	}
	if o.PickleId.IsSet() {
		toSerialize["pickle_id"] = o.PickleId.Get()
	}
	if o.DefaultView != nil {
		toSerialize["default_view"] = o.DefaultView
	}
	if o.Fileloc != nil {
		toSerialize["fileloc"] = o.Fileloc
	}
	if o.FileToken != nil {
		toSerialize["file_token"] = o.FileToken
	}
	if o.Owners != nil {
		toSerialize["owners"] = o.Owners
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.ScheduleInterval.IsSet() {
		toSerialize["schedule_interval"] = o.ScheduleInterval.Get()
	}
	if o.TimetableDescription.IsSet() {
		toSerialize["timetable_description"] = o.TimetableDescription.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.MaxActiveTasks.IsSet() {
		toSerialize["max_active_tasks"] = o.MaxActiveTasks.Get()
	}
	if o.MaxActiveRuns.IsSet() {
		toSerialize["max_active_runs"] = o.MaxActiveRuns.Get()
	}
	if o.HasTaskConcurrencyLimits.IsSet() {
		toSerialize["has_task_concurrency_limits"] = o.HasTaskConcurrencyLimits.Get()
	}
	if o.HasImportErrors.IsSet() {
		toSerialize["has_import_errors"] = o.HasImportErrors.Get()
	}
	if o.NextDagrun.IsSet() {
		toSerialize["next_dagrun"] = o.NextDagrun.Get()
	}
	if o.NextDagrunDataIntervalStart.IsSet() {
		toSerialize["next_dagrun_data_interval_start"] = o.NextDagrunDataIntervalStart.Get()
	}
	if o.NextDagrunDataIntervalEnd.IsSet() {
		toSerialize["next_dagrun_data_interval_end"] = o.NextDagrunDataIntervalEnd.Get()
	}
	if o.NextDagrunCreateAfter.IsSet() {
		toSerialize["next_dagrun_create_after"] = o.NextDagrunCreateAfter.Get()
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	if o.Catchup != nil {
		toSerialize["catchup"] = o.Catchup
	}
	if o.Orientation != nil {
		toSerialize["orientation"] = o.Orientation
	}
	if o.Concurrency != nil {
		toSerialize["concurrency"] = o.Concurrency
	}
	if o.StartDate.IsSet() {
		toSerialize["start_date"] = o.StartDate.Get()
	}
	if o.DagRunTimeout.IsSet() {
		toSerialize["dag_run_timeout"] = o.DagRunTimeout.Get()
	}
	if o.DocMd.IsSet() {
		toSerialize["doc_md"] = o.DocMd.Get()
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	if o.EndDate.IsSet() {
		toSerialize["end_date"] = o.EndDate.Get()
	}
	if o.IsPausedUponCreation.IsSet() {
		toSerialize["is_paused_upon_creation"] = o.IsPausedUponCreation.Get()
	}
	if o.LastParsed.IsSet() {
		toSerialize["last_parsed"] = o.LastParsed.Get()
	}
	if o.TemplateSearchPath != nil {
		toSerialize["template_search_path"] = o.TemplateSearchPath
	}
	if o.RenderTemplateAsNativeObj.IsSet() {
		toSerialize["render_template_as_native_obj"] = o.RenderTemplateAsNativeObj.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDAGDetail struct {
	value *DAGDetail
	isSet bool
}

func (v NullableDAGDetail) Get() *DAGDetail {
	return v.value
}

func (v *NullableDAGDetail) Set(val *DAGDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableDAGDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableDAGDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDAGDetail(val *DAGDetail) *NullableDAGDetail {
	return &NullableDAGDetail{value: val, isSet: true}
}

func (v NullableDAGDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDAGDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


