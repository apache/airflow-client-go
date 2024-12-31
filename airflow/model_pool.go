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

// Pool The pool
type Pool struct {
	// The name of pool.
	Name *string `json:"name,omitempty"`
	// The maximum number of slots that can be assigned to tasks. One job may occupy one or more slots. 
	Slots *int32 `json:"slots,omitempty"`
	// The number of slots used by running/queued tasks at the moment. May include deferred tasks if 'include_deferred' is set to true.
	OccupiedSlots *int32 `json:"occupied_slots,omitempty"`
	// The number of slots used by running tasks at the moment.
	RunningSlots *int32 `json:"running_slots,omitempty"`
	// The number of slots used by queued tasks at the moment.
	QueuedSlots *int32 `json:"queued_slots,omitempty"`
	// The number of free slots at the moment.
	OpenSlots *int32 `json:"open_slots,omitempty"`
	// The number of slots used by scheduled tasks at the moment.
	ScheduledSlots *int32 `json:"scheduled_slots,omitempty"`
	// The number of slots used by deferred tasks at the moment. Relevant if 'include_deferred' is set to true.  *New in version 2.7.0* 
	DeferredSlots *int32 `json:"deferred_slots,omitempty"`
	// The description of the pool.  *New in version 2.3.0* 
	Description NullableString `json:"description,omitempty"`
	// If set to true, deferred tasks are considered when calculating open pool slots.  *New in version 2.7.0* 
	IncludeDeferred *bool `json:"include_deferred,omitempty"`
}

// NewPool instantiates a new Pool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPool() *Pool {
	this := Pool{}
	return &this
}

// NewPoolWithDefaults instantiates a new Pool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolWithDefaults() *Pool {
	this := Pool{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Pool) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pool) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Pool) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Pool) SetName(v string) {
	o.Name = &v
}

// GetSlots returns the Slots field value if set, zero value otherwise.
func (o *Pool) GetSlots() int32 {
	if o == nil || o.Slots == nil {
		var ret int32
		return ret
	}
	return *o.Slots
}

// GetSlotsOk returns a tuple with the Slots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pool) GetSlotsOk() (*int32, bool) {
	if o == nil || o.Slots == nil {
		return nil, false
	}
	return o.Slots, true
}

// HasSlots returns a boolean if a field has been set.
func (o *Pool) HasSlots() bool {
	if o != nil && o.Slots != nil {
		return true
	}

	return false
}

// SetSlots gets a reference to the given int32 and assigns it to the Slots field.
func (o *Pool) SetSlots(v int32) {
	o.Slots = &v
}

// GetOccupiedSlots returns the OccupiedSlots field value if set, zero value otherwise.
func (o *Pool) GetOccupiedSlots() int32 {
	if o == nil || o.OccupiedSlots == nil {
		var ret int32
		return ret
	}
	return *o.OccupiedSlots
}

// GetOccupiedSlotsOk returns a tuple with the OccupiedSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pool) GetOccupiedSlotsOk() (*int32, bool) {
	if o == nil || o.OccupiedSlots == nil {
		return nil, false
	}
	return o.OccupiedSlots, true
}

// HasOccupiedSlots returns a boolean if a field has been set.
func (o *Pool) HasOccupiedSlots() bool {
	if o != nil && o.OccupiedSlots != nil {
		return true
	}

	return false
}

// SetOccupiedSlots gets a reference to the given int32 and assigns it to the OccupiedSlots field.
func (o *Pool) SetOccupiedSlots(v int32) {
	o.OccupiedSlots = &v
}

// GetRunningSlots returns the RunningSlots field value if set, zero value otherwise.
func (o *Pool) GetRunningSlots() int32 {
	if o == nil || o.RunningSlots == nil {
		var ret int32
		return ret
	}
	return *o.RunningSlots
}

// GetRunningSlotsOk returns a tuple with the RunningSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pool) GetRunningSlotsOk() (*int32, bool) {
	if o == nil || o.RunningSlots == nil {
		return nil, false
	}
	return o.RunningSlots, true
}

// HasRunningSlots returns a boolean if a field has been set.
func (o *Pool) HasRunningSlots() bool {
	if o != nil && o.RunningSlots != nil {
		return true
	}

	return false
}

// SetRunningSlots gets a reference to the given int32 and assigns it to the RunningSlots field.
func (o *Pool) SetRunningSlots(v int32) {
	o.RunningSlots = &v
}

// GetQueuedSlots returns the QueuedSlots field value if set, zero value otherwise.
func (o *Pool) GetQueuedSlots() int32 {
	if o == nil || o.QueuedSlots == nil {
		var ret int32
		return ret
	}
	return *o.QueuedSlots
}

// GetQueuedSlotsOk returns a tuple with the QueuedSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pool) GetQueuedSlotsOk() (*int32, bool) {
	if o == nil || o.QueuedSlots == nil {
		return nil, false
	}
	return o.QueuedSlots, true
}

// HasQueuedSlots returns a boolean if a field has been set.
func (o *Pool) HasQueuedSlots() bool {
	if o != nil && o.QueuedSlots != nil {
		return true
	}

	return false
}

// SetQueuedSlots gets a reference to the given int32 and assigns it to the QueuedSlots field.
func (o *Pool) SetQueuedSlots(v int32) {
	o.QueuedSlots = &v
}

// GetOpenSlots returns the OpenSlots field value if set, zero value otherwise.
func (o *Pool) GetOpenSlots() int32 {
	if o == nil || o.OpenSlots == nil {
		var ret int32
		return ret
	}
	return *o.OpenSlots
}

// GetOpenSlotsOk returns a tuple with the OpenSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pool) GetOpenSlotsOk() (*int32, bool) {
	if o == nil || o.OpenSlots == nil {
		return nil, false
	}
	return o.OpenSlots, true
}

// HasOpenSlots returns a boolean if a field has been set.
func (o *Pool) HasOpenSlots() bool {
	if o != nil && o.OpenSlots != nil {
		return true
	}

	return false
}

// SetOpenSlots gets a reference to the given int32 and assigns it to the OpenSlots field.
func (o *Pool) SetOpenSlots(v int32) {
	o.OpenSlots = &v
}

// GetScheduledSlots returns the ScheduledSlots field value if set, zero value otherwise.
func (o *Pool) GetScheduledSlots() int32 {
	if o == nil || o.ScheduledSlots == nil {
		var ret int32
		return ret
	}
	return *o.ScheduledSlots
}

// GetScheduledSlotsOk returns a tuple with the ScheduledSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pool) GetScheduledSlotsOk() (*int32, bool) {
	if o == nil || o.ScheduledSlots == nil {
		return nil, false
	}
	return o.ScheduledSlots, true
}

// HasScheduledSlots returns a boolean if a field has been set.
func (o *Pool) HasScheduledSlots() bool {
	if o != nil && o.ScheduledSlots != nil {
		return true
	}

	return false
}

// SetScheduledSlots gets a reference to the given int32 and assigns it to the ScheduledSlots field.
func (o *Pool) SetScheduledSlots(v int32) {
	o.ScheduledSlots = &v
}

// GetDeferredSlots returns the DeferredSlots field value if set, zero value otherwise.
func (o *Pool) GetDeferredSlots() int32 {
	if o == nil || o.DeferredSlots == nil {
		var ret int32
		return ret
	}
	return *o.DeferredSlots
}

// GetDeferredSlotsOk returns a tuple with the DeferredSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pool) GetDeferredSlotsOk() (*int32, bool) {
	if o == nil || o.DeferredSlots == nil {
		return nil, false
	}
	return o.DeferredSlots, true
}

// HasDeferredSlots returns a boolean if a field has been set.
func (o *Pool) HasDeferredSlots() bool {
	if o != nil && o.DeferredSlots != nil {
		return true
	}

	return false
}

// SetDeferredSlots gets a reference to the given int32 and assigns it to the DeferredSlots field.
func (o *Pool) SetDeferredSlots(v int32) {
	o.DeferredSlots = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Pool) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pool) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Pool) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Pool) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Pool) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Pool) UnsetDescription() {
	o.Description.Unset()
}

// GetIncludeDeferred returns the IncludeDeferred field value if set, zero value otherwise.
func (o *Pool) GetIncludeDeferred() bool {
	if o == nil || o.IncludeDeferred == nil {
		var ret bool
		return ret
	}
	return *o.IncludeDeferred
}

// GetIncludeDeferredOk returns a tuple with the IncludeDeferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pool) GetIncludeDeferredOk() (*bool, bool) {
	if o == nil || o.IncludeDeferred == nil {
		return nil, false
	}
	return o.IncludeDeferred, true
}

// HasIncludeDeferred returns a boolean if a field has been set.
func (o *Pool) HasIncludeDeferred() bool {
	if o != nil && o.IncludeDeferred != nil {
		return true
	}

	return false
}

// SetIncludeDeferred gets a reference to the given bool and assigns it to the IncludeDeferred field.
func (o *Pool) SetIncludeDeferred(v bool) {
	o.IncludeDeferred = &v
}

func (o Pool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Slots != nil {
		toSerialize["slots"] = o.Slots
	}
	if o.OccupiedSlots != nil {
		toSerialize["occupied_slots"] = o.OccupiedSlots
	}
	if o.RunningSlots != nil {
		toSerialize["running_slots"] = o.RunningSlots
	}
	if o.QueuedSlots != nil {
		toSerialize["queued_slots"] = o.QueuedSlots
	}
	if o.OpenSlots != nil {
		toSerialize["open_slots"] = o.OpenSlots
	}
	if o.ScheduledSlots != nil {
		toSerialize["scheduled_slots"] = o.ScheduledSlots
	}
	if o.DeferredSlots != nil {
		toSerialize["deferred_slots"] = o.DeferredSlots
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.IncludeDeferred != nil {
		toSerialize["include_deferred"] = o.IncludeDeferred
	}
	return json.Marshal(toSerialize)
}

type NullablePool struct {
	value *Pool
	isSet bool
}

func (v NullablePool) Get() *Pool {
	return v.value
}

func (v *NullablePool) Set(val *Pool) {
	v.value = val
	v.isSet = true
}

func (v NullablePool) IsSet() bool {
	return v.isSet
}

func (v *NullablePool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePool(val *Pool) *NullablePool {
	return &NullablePool{value: val, isSet: true}
}

func (v NullablePool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


