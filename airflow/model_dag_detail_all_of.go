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

// DAGDetailAllOf struct for DAGDetailAllOf
type DAGDetailAllOf struct {
	Timezone *string `json:"timezone,omitempty"`
	Catchup *bool `json:"catchup,omitempty"`
	Orientation *string `json:"orientation,omitempty"`
	Concurrency *float32 `json:"concurrency,omitempty"`
	// The DAG's start date.  *Changed in version 2.0.1*&#58; Field becomes nullable. 
	StartDate NullableTime `json:"start_date,omitempty"`
	DagRunTimeout NullableTimeDelta `json:"dag_run_timeout,omitempty"`
	DocMd NullableString `json:"doc_md,omitempty"`
	DefaultView *string `json:"default_view,omitempty"`
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

// NewDAGDetailAllOf instantiates a new DAGDetailAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDAGDetailAllOf() *DAGDetailAllOf {
	this := DAGDetailAllOf{}
	return &this
}

// NewDAGDetailAllOfWithDefaults instantiates a new DAGDetailAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDAGDetailAllOfWithDefaults() *DAGDetailAllOf {
	this := DAGDetailAllOf{}
	return &this
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *DAGDetailAllOf) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGDetailAllOf) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *DAGDetailAllOf) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *DAGDetailAllOf) SetTimezone(v string) {
	o.Timezone = &v
}

// GetCatchup returns the Catchup field value if set, zero value otherwise.
func (o *DAGDetailAllOf) GetCatchup() bool {
	if o == nil || o.Catchup == nil {
		var ret bool
		return ret
	}
	return *o.Catchup
}

// GetCatchupOk returns a tuple with the Catchup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGDetailAllOf) GetCatchupOk() (*bool, bool) {
	if o == nil || o.Catchup == nil {
		return nil, false
	}
	return o.Catchup, true
}

// HasCatchup returns a boolean if a field has been set.
func (o *DAGDetailAllOf) HasCatchup() bool {
	if o != nil && o.Catchup != nil {
		return true
	}

	return false
}

// SetCatchup gets a reference to the given bool and assigns it to the Catchup field.
func (o *DAGDetailAllOf) SetCatchup(v bool) {
	o.Catchup = &v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *DAGDetailAllOf) GetOrientation() string {
	if o == nil || o.Orientation == nil {
		var ret string
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGDetailAllOf) GetOrientationOk() (*string, bool) {
	if o == nil || o.Orientation == nil {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *DAGDetailAllOf) HasOrientation() bool {
	if o != nil && o.Orientation != nil {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given string and assigns it to the Orientation field.
func (o *DAGDetailAllOf) SetOrientation(v string) {
	o.Orientation = &v
}

// GetConcurrency returns the Concurrency field value if set, zero value otherwise.
func (o *DAGDetailAllOf) GetConcurrency() float32 {
	if o == nil || o.Concurrency == nil {
		var ret float32
		return ret
	}
	return *o.Concurrency
}

// GetConcurrencyOk returns a tuple with the Concurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGDetailAllOf) GetConcurrencyOk() (*float32, bool) {
	if o == nil || o.Concurrency == nil {
		return nil, false
	}
	return o.Concurrency, true
}

// HasConcurrency returns a boolean if a field has been set.
func (o *DAGDetailAllOf) HasConcurrency() bool {
	if o != nil && o.Concurrency != nil {
		return true
	}

	return false
}

// SetConcurrency gets a reference to the given float32 and assigns it to the Concurrency field.
func (o *DAGDetailAllOf) SetConcurrency(v float32) {
	o.Concurrency = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetailAllOf) GetStartDate() time.Time {
	if o == nil || o.StartDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetailAllOf) GetStartDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *DAGDetailAllOf) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableTime and assigns it to the StartDate field.
func (o *DAGDetailAllOf) SetStartDate(v time.Time) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *DAGDetailAllOf) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *DAGDetailAllOf) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetDagRunTimeout returns the DagRunTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetailAllOf) GetDagRunTimeout() TimeDelta {
	if o == nil || o.DagRunTimeout.Get() == nil {
		var ret TimeDelta
		return ret
	}
	return *o.DagRunTimeout.Get()
}

// GetDagRunTimeoutOk returns a tuple with the DagRunTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetailAllOf) GetDagRunTimeoutOk() (*TimeDelta, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DagRunTimeout.Get(), o.DagRunTimeout.IsSet()
}

// HasDagRunTimeout returns a boolean if a field has been set.
func (o *DAGDetailAllOf) HasDagRunTimeout() bool {
	if o != nil && o.DagRunTimeout.IsSet() {
		return true
	}

	return false
}

// SetDagRunTimeout gets a reference to the given NullableTimeDelta and assigns it to the DagRunTimeout field.
func (o *DAGDetailAllOf) SetDagRunTimeout(v TimeDelta) {
	o.DagRunTimeout.Set(&v)
}
// SetDagRunTimeoutNil sets the value for DagRunTimeout to be an explicit nil
func (o *DAGDetailAllOf) SetDagRunTimeoutNil() {
	o.DagRunTimeout.Set(nil)
}

// UnsetDagRunTimeout ensures that no value is present for DagRunTimeout, not even an explicit nil
func (o *DAGDetailAllOf) UnsetDagRunTimeout() {
	o.DagRunTimeout.Unset()
}

// GetDocMd returns the DocMd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetailAllOf) GetDocMd() string {
	if o == nil || o.DocMd.Get() == nil {
		var ret string
		return ret
	}
	return *o.DocMd.Get()
}

// GetDocMdOk returns a tuple with the DocMd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetailAllOf) GetDocMdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DocMd.Get(), o.DocMd.IsSet()
}

// HasDocMd returns a boolean if a field has been set.
func (o *DAGDetailAllOf) HasDocMd() bool {
	if o != nil && o.DocMd.IsSet() {
		return true
	}

	return false
}

// SetDocMd gets a reference to the given NullableString and assigns it to the DocMd field.
func (o *DAGDetailAllOf) SetDocMd(v string) {
	o.DocMd.Set(&v)
}
// SetDocMdNil sets the value for DocMd to be an explicit nil
func (o *DAGDetailAllOf) SetDocMdNil() {
	o.DocMd.Set(nil)
}

// UnsetDocMd ensures that no value is present for DocMd, not even an explicit nil
func (o *DAGDetailAllOf) UnsetDocMd() {
	o.DocMd.Unset()
}

// GetDefaultView returns the DefaultView field value if set, zero value otherwise.
func (o *DAGDetailAllOf) GetDefaultView() string {
	if o == nil || o.DefaultView == nil {
		var ret string
		return ret
	}
	return *o.DefaultView
}

// GetDefaultViewOk returns a tuple with the DefaultView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGDetailAllOf) GetDefaultViewOk() (*string, bool) {
	if o == nil || o.DefaultView == nil {
		return nil, false
	}
	return o.DefaultView, true
}

// HasDefaultView returns a boolean if a field has been set.
func (o *DAGDetailAllOf) HasDefaultView() bool {
	if o != nil && o.DefaultView != nil {
		return true
	}

	return false
}

// SetDefaultView gets a reference to the given string and assigns it to the DefaultView field.
func (o *DAGDetailAllOf) SetDefaultView(v string) {
	o.DefaultView = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *DAGDetailAllOf) GetParams() map[string]interface{} {
	if o == nil || o.Params == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGDetailAllOf) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *DAGDetailAllOf) HasParams() bool {
	if o != nil && o.Params != nil {
		return true
	}

	return false
}

// SetParams gets a reference to the given map[string]interface{} and assigns it to the Params field.
func (o *DAGDetailAllOf) SetParams(v map[string]interface{}) {
	o.Params = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetailAllOf) GetEndDate() time.Time {
	if o == nil || o.EndDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetailAllOf) GetEndDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *DAGDetailAllOf) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableTime and assigns it to the EndDate field.
func (o *DAGDetailAllOf) SetEndDate(v time.Time) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *DAGDetailAllOf) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *DAGDetailAllOf) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetIsPausedUponCreation returns the IsPausedUponCreation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetailAllOf) GetIsPausedUponCreation() bool {
	if o == nil || o.IsPausedUponCreation.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsPausedUponCreation.Get()
}

// GetIsPausedUponCreationOk returns a tuple with the IsPausedUponCreation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetailAllOf) GetIsPausedUponCreationOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsPausedUponCreation.Get(), o.IsPausedUponCreation.IsSet()
}

// HasIsPausedUponCreation returns a boolean if a field has been set.
func (o *DAGDetailAllOf) HasIsPausedUponCreation() bool {
	if o != nil && o.IsPausedUponCreation.IsSet() {
		return true
	}

	return false
}

// SetIsPausedUponCreation gets a reference to the given NullableBool and assigns it to the IsPausedUponCreation field.
func (o *DAGDetailAllOf) SetIsPausedUponCreation(v bool) {
	o.IsPausedUponCreation.Set(&v)
}
// SetIsPausedUponCreationNil sets the value for IsPausedUponCreation to be an explicit nil
func (o *DAGDetailAllOf) SetIsPausedUponCreationNil() {
	o.IsPausedUponCreation.Set(nil)
}

// UnsetIsPausedUponCreation ensures that no value is present for IsPausedUponCreation, not even an explicit nil
func (o *DAGDetailAllOf) UnsetIsPausedUponCreation() {
	o.IsPausedUponCreation.Unset()
}

// GetLastParsed returns the LastParsed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetailAllOf) GetLastParsed() time.Time {
	if o == nil || o.LastParsed.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastParsed.Get()
}

// GetLastParsedOk returns a tuple with the LastParsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetailAllOf) GetLastParsedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastParsed.Get(), o.LastParsed.IsSet()
}

// HasLastParsed returns a boolean if a field has been set.
func (o *DAGDetailAllOf) HasLastParsed() bool {
	if o != nil && o.LastParsed.IsSet() {
		return true
	}

	return false
}

// SetLastParsed gets a reference to the given NullableTime and assigns it to the LastParsed field.
func (o *DAGDetailAllOf) SetLastParsed(v time.Time) {
	o.LastParsed.Set(&v)
}
// SetLastParsedNil sets the value for LastParsed to be an explicit nil
func (o *DAGDetailAllOf) SetLastParsedNil() {
	o.LastParsed.Set(nil)
}

// UnsetLastParsed ensures that no value is present for LastParsed, not even an explicit nil
func (o *DAGDetailAllOf) UnsetLastParsed() {
	o.LastParsed.Unset()
}

// GetTemplateSearchPath returns the TemplateSearchPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetailAllOf) GetTemplateSearchPath() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.TemplateSearchPath
}

// GetTemplateSearchPathOk returns a tuple with the TemplateSearchPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetailAllOf) GetTemplateSearchPathOk() ([]string, bool) {
	if o == nil || o.TemplateSearchPath == nil {
		return nil, false
	}
	return o.TemplateSearchPath, true
}

// HasTemplateSearchPath returns a boolean if a field has been set.
func (o *DAGDetailAllOf) HasTemplateSearchPath() bool {
	if o != nil && o.TemplateSearchPath != nil {
		return true
	}

	return false
}

// SetTemplateSearchPath gets a reference to the given []string and assigns it to the TemplateSearchPath field.
func (o *DAGDetailAllOf) SetTemplateSearchPath(v []string) {
	o.TemplateSearchPath = v
}

// GetRenderTemplateAsNativeObj returns the RenderTemplateAsNativeObj field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAGDetailAllOf) GetRenderTemplateAsNativeObj() bool {
	if o == nil || o.RenderTemplateAsNativeObj.Get() == nil {
		var ret bool
		return ret
	}
	return *o.RenderTemplateAsNativeObj.Get()
}

// GetRenderTemplateAsNativeObjOk returns a tuple with the RenderTemplateAsNativeObj field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAGDetailAllOf) GetRenderTemplateAsNativeObjOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RenderTemplateAsNativeObj.Get(), o.RenderTemplateAsNativeObj.IsSet()
}

// HasRenderTemplateAsNativeObj returns a boolean if a field has been set.
func (o *DAGDetailAllOf) HasRenderTemplateAsNativeObj() bool {
	if o != nil && o.RenderTemplateAsNativeObj.IsSet() {
		return true
	}

	return false
}

// SetRenderTemplateAsNativeObj gets a reference to the given NullableBool and assigns it to the RenderTemplateAsNativeObj field.
func (o *DAGDetailAllOf) SetRenderTemplateAsNativeObj(v bool) {
	o.RenderTemplateAsNativeObj.Set(&v)
}
// SetRenderTemplateAsNativeObjNil sets the value for RenderTemplateAsNativeObj to be an explicit nil
func (o *DAGDetailAllOf) SetRenderTemplateAsNativeObjNil() {
	o.RenderTemplateAsNativeObj.Set(nil)
}

// UnsetRenderTemplateAsNativeObj ensures that no value is present for RenderTemplateAsNativeObj, not even an explicit nil
func (o *DAGDetailAllOf) UnsetRenderTemplateAsNativeObj() {
	o.RenderTemplateAsNativeObj.Unset()
}

func (o DAGDetailAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.DefaultView != nil {
		toSerialize["default_view"] = o.DefaultView
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

type NullableDAGDetailAllOf struct {
	value *DAGDetailAllOf
	isSet bool
}

func (v NullableDAGDetailAllOf) Get() *DAGDetailAllOf {
	return v.value
}

func (v *NullableDAGDetailAllOf) Set(val *DAGDetailAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDAGDetailAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDAGDetailAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDAGDetailAllOf(val *DAGDetailAllOf) *NullableDAGDetailAllOf {
	return &NullableDAGDetailAllOf{value: val, isSet: true}
}

func (v NullableDAGDetailAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDAGDetailAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


