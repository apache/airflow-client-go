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

// DatasetEvent A dataset event.  *New in version 2.4.0* 
type DatasetEvent struct {
	// The dataset id
	DatasetId *int32 `json:"dataset_id,omitempty"`
	// The URI of the dataset
	DatasetUri *string `json:"dataset_uri,omitempty"`
	// The dataset event extra
	Extra map[string]interface{} `json:"extra,omitempty"`
	// The DAG ID that updated the dataset.
	SourceDagId NullableString `json:"source_dag_id,omitempty"`
	// The task ID that updated the dataset.
	SourceTaskId NullableString `json:"source_task_id,omitempty"`
	// The DAG run ID that updated the dataset.
	SourceRunId NullableString `json:"source_run_id,omitempty"`
	// The task map index that updated the dataset.
	SourceMapIndex NullableInt32 `json:"source_map_index,omitempty"`
	CreatedDagruns []BasicDAGRun `json:"created_dagruns,omitempty"`
	// The dataset event creation time
	Timestamp *string `json:"timestamp,omitempty"`
}

// NewDatasetEvent instantiates a new DatasetEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasetEvent() *DatasetEvent {
	this := DatasetEvent{}
	return &this
}

// NewDatasetEventWithDefaults instantiates a new DatasetEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasetEventWithDefaults() *DatasetEvent {
	this := DatasetEvent{}
	return &this
}

// GetDatasetId returns the DatasetId field value if set, zero value otherwise.
func (o *DatasetEvent) GetDatasetId() int32 {
	if o == nil || o.DatasetId == nil {
		var ret int32
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetEvent) GetDatasetIdOk() (*int32, bool) {
	if o == nil || o.DatasetId == nil {
		return nil, false
	}
	return o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *DatasetEvent) HasDatasetId() bool {
	if o != nil && o.DatasetId != nil {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given int32 and assigns it to the DatasetId field.
func (o *DatasetEvent) SetDatasetId(v int32) {
	o.DatasetId = &v
}

// GetDatasetUri returns the DatasetUri field value if set, zero value otherwise.
func (o *DatasetEvent) GetDatasetUri() string {
	if o == nil || o.DatasetUri == nil {
		var ret string
		return ret
	}
	return *o.DatasetUri
}

// GetDatasetUriOk returns a tuple with the DatasetUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetEvent) GetDatasetUriOk() (*string, bool) {
	if o == nil || o.DatasetUri == nil {
		return nil, false
	}
	return o.DatasetUri, true
}

// HasDatasetUri returns a boolean if a field has been set.
func (o *DatasetEvent) HasDatasetUri() bool {
	if o != nil && o.DatasetUri != nil {
		return true
	}

	return false
}

// SetDatasetUri gets a reference to the given string and assigns it to the DatasetUri field.
func (o *DatasetEvent) SetDatasetUri(v string) {
	o.DatasetUri = &v
}

// GetExtra returns the Extra field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatasetEvent) GetExtra() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatasetEvent) GetExtraOk() (map[string]interface{}, bool) {
	if o == nil || o.Extra == nil {
		return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *DatasetEvent) HasExtra() bool {
	if o != nil && o.Extra != nil {
		return true
	}

	return false
}

// SetExtra gets a reference to the given map[string]interface{} and assigns it to the Extra field.
func (o *DatasetEvent) SetExtra(v map[string]interface{}) {
	o.Extra = v
}

// GetSourceDagId returns the SourceDagId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatasetEvent) GetSourceDagId() string {
	if o == nil || o.SourceDagId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceDagId.Get()
}

// GetSourceDagIdOk returns a tuple with the SourceDagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatasetEvent) GetSourceDagIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceDagId.Get(), o.SourceDagId.IsSet()
}

// HasSourceDagId returns a boolean if a field has been set.
func (o *DatasetEvent) HasSourceDagId() bool {
	if o != nil && o.SourceDagId.IsSet() {
		return true
	}

	return false
}

// SetSourceDagId gets a reference to the given NullableString and assigns it to the SourceDagId field.
func (o *DatasetEvent) SetSourceDagId(v string) {
	o.SourceDagId.Set(&v)
}
// SetSourceDagIdNil sets the value for SourceDagId to be an explicit nil
func (o *DatasetEvent) SetSourceDagIdNil() {
	o.SourceDagId.Set(nil)
}

// UnsetSourceDagId ensures that no value is present for SourceDagId, not even an explicit nil
func (o *DatasetEvent) UnsetSourceDagId() {
	o.SourceDagId.Unset()
}

// GetSourceTaskId returns the SourceTaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatasetEvent) GetSourceTaskId() string {
	if o == nil || o.SourceTaskId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceTaskId.Get()
}

// GetSourceTaskIdOk returns a tuple with the SourceTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatasetEvent) GetSourceTaskIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceTaskId.Get(), o.SourceTaskId.IsSet()
}

// HasSourceTaskId returns a boolean if a field has been set.
func (o *DatasetEvent) HasSourceTaskId() bool {
	if o != nil && o.SourceTaskId.IsSet() {
		return true
	}

	return false
}

// SetSourceTaskId gets a reference to the given NullableString and assigns it to the SourceTaskId field.
func (o *DatasetEvent) SetSourceTaskId(v string) {
	o.SourceTaskId.Set(&v)
}
// SetSourceTaskIdNil sets the value for SourceTaskId to be an explicit nil
func (o *DatasetEvent) SetSourceTaskIdNil() {
	o.SourceTaskId.Set(nil)
}

// UnsetSourceTaskId ensures that no value is present for SourceTaskId, not even an explicit nil
func (o *DatasetEvent) UnsetSourceTaskId() {
	o.SourceTaskId.Unset()
}

// GetSourceRunId returns the SourceRunId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatasetEvent) GetSourceRunId() string {
	if o == nil || o.SourceRunId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceRunId.Get()
}

// GetSourceRunIdOk returns a tuple with the SourceRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatasetEvent) GetSourceRunIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceRunId.Get(), o.SourceRunId.IsSet()
}

// HasSourceRunId returns a boolean if a field has been set.
func (o *DatasetEvent) HasSourceRunId() bool {
	if o != nil && o.SourceRunId.IsSet() {
		return true
	}

	return false
}

// SetSourceRunId gets a reference to the given NullableString and assigns it to the SourceRunId field.
func (o *DatasetEvent) SetSourceRunId(v string) {
	o.SourceRunId.Set(&v)
}
// SetSourceRunIdNil sets the value for SourceRunId to be an explicit nil
func (o *DatasetEvent) SetSourceRunIdNil() {
	o.SourceRunId.Set(nil)
}

// UnsetSourceRunId ensures that no value is present for SourceRunId, not even an explicit nil
func (o *DatasetEvent) UnsetSourceRunId() {
	o.SourceRunId.Unset()
}

// GetSourceMapIndex returns the SourceMapIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatasetEvent) GetSourceMapIndex() int32 {
	if o == nil || o.SourceMapIndex.Get() == nil {
		var ret int32
		return ret
	}
	return *o.SourceMapIndex.Get()
}

// GetSourceMapIndexOk returns a tuple with the SourceMapIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatasetEvent) GetSourceMapIndexOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceMapIndex.Get(), o.SourceMapIndex.IsSet()
}

// HasSourceMapIndex returns a boolean if a field has been set.
func (o *DatasetEvent) HasSourceMapIndex() bool {
	if o != nil && o.SourceMapIndex.IsSet() {
		return true
	}

	return false
}

// SetSourceMapIndex gets a reference to the given NullableInt32 and assigns it to the SourceMapIndex field.
func (o *DatasetEvent) SetSourceMapIndex(v int32) {
	o.SourceMapIndex.Set(&v)
}
// SetSourceMapIndexNil sets the value for SourceMapIndex to be an explicit nil
func (o *DatasetEvent) SetSourceMapIndexNil() {
	o.SourceMapIndex.Set(nil)
}

// UnsetSourceMapIndex ensures that no value is present for SourceMapIndex, not even an explicit nil
func (o *DatasetEvent) UnsetSourceMapIndex() {
	o.SourceMapIndex.Unset()
}

// GetCreatedDagruns returns the CreatedDagruns field value if set, zero value otherwise.
func (o *DatasetEvent) GetCreatedDagruns() []BasicDAGRun {
	if o == nil || o.CreatedDagruns == nil {
		var ret []BasicDAGRun
		return ret
	}
	return o.CreatedDagruns
}

// GetCreatedDagrunsOk returns a tuple with the CreatedDagruns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetEvent) GetCreatedDagrunsOk() ([]BasicDAGRun, bool) {
	if o == nil || o.CreatedDagruns == nil {
		return nil, false
	}
	return o.CreatedDagruns, true
}

// HasCreatedDagruns returns a boolean if a field has been set.
func (o *DatasetEvent) HasCreatedDagruns() bool {
	if o != nil && o.CreatedDagruns != nil {
		return true
	}

	return false
}

// SetCreatedDagruns gets a reference to the given []BasicDAGRun and assigns it to the CreatedDagruns field.
func (o *DatasetEvent) SetCreatedDagruns(v []BasicDAGRun) {
	o.CreatedDagruns = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *DatasetEvent) GetTimestamp() string {
	if o == nil || o.Timestamp == nil {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetEvent) GetTimestampOk() (*string, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *DatasetEvent) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *DatasetEvent) SetTimestamp(v string) {
	o.Timestamp = &v
}

func (o DatasetEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DatasetId != nil {
		toSerialize["dataset_id"] = o.DatasetId
	}
	if o.DatasetUri != nil {
		toSerialize["dataset_uri"] = o.DatasetUri
	}
	if o.Extra != nil {
		toSerialize["extra"] = o.Extra
	}
	if o.SourceDagId.IsSet() {
		toSerialize["source_dag_id"] = o.SourceDagId.Get()
	}
	if o.SourceTaskId.IsSet() {
		toSerialize["source_task_id"] = o.SourceTaskId.Get()
	}
	if o.SourceRunId.IsSet() {
		toSerialize["source_run_id"] = o.SourceRunId.Get()
	}
	if o.SourceMapIndex.IsSet() {
		toSerialize["source_map_index"] = o.SourceMapIndex.Get()
	}
	if o.CreatedDagruns != nil {
		toSerialize["created_dagruns"] = o.CreatedDagruns
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableDatasetEvent struct {
	value *DatasetEvent
	isSet bool
}

func (v NullableDatasetEvent) Get() *DatasetEvent {
	return v.value
}

func (v *NullableDatasetEvent) Set(val *DatasetEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetEvent(val *DatasetEvent) *NullableDatasetEvent {
	return &NullableDatasetEvent{value: val, isSet: true}
}

func (v NullableDatasetEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


