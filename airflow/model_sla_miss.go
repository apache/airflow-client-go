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
 * Airflow API (Stable)
 *
 * # Overview  To facilitate management, Apache Airflow supports a range of REST API endpoints across its objects. This section provides an overview of the API design, methods, and supported use cases.  Most of the endpoints accept `JSON` as input and return `JSON` responses. This means that you must usually add the following headers to your request: ``` Content-type: application/json Accept: application/json ```  ## Resources  The term `resource` refers to a single type of object in the Airflow metadata. An API is broken up by its endpoint's corresponding resource. The name of a resource is typically plural and expressed in camelCase. Example: `dagRuns`.  Resource names are used as part of endpoint URLs, as well as in API parameters and responses.  ## CRUD Operations  The platform supports **C**reate, **R**ead, **U**pdate, and **D**elete operations on most resources. You can review the standards for these operations and their standard parameters below.  Some endpoints have special behavior as exceptions.  ### Create  To create a resource, you typically submit an HTTP `POST` request with the resource's required metadata in the request body. The response returns a `201 Created` response code upon success with the resource's metadata, including its internal `id`, in the response body.  ### Read  The HTTP `GET` request can be used to read a resource or to list a number of resources.  A resource's `id` can be submitted in the request parameters to read a specific resource. The response usually returns a `200 OK` response code upon success, with the resource's metadata in the response body.  If a `GET` request does not include a specific resource `id`, it is treated as a list request. The response usually returns a `200 OK` response code upon success, with an object containing a list of resources' metadata in the response body.  When reading resources, some common query parameters are usually available. e.g.: ``` v1/connections?limit=25&offset=25 ```  |Query Parameter|Type|Description| |---------------|----|-----------| |limit|integer|Maximum number of objects to fetch. Usually 25 by default| |offset|integer|Offset after which to start returning objects. For use with limit query parameter.|  ### Update  Updating a resource requires the resource `id`, and is typically done using an HTTP `PATCH` request, with the fields to modify in the request body. The response usually returns a `200 OK` response code upon success, with information about the modified resource in the response body.  ### Delete  Deleting a resource requires the resource `id` and is typically executing via an HTTP `DELETE` request. The response usually returns a `204 No Content` response code upon success.  ## Conventions  - Resource names are plural and expressed in camelCase. - Names are consistent between URL parameter name and field name.  - Field names are in snake_case. ```json {     \"name\": \"string\",     \"slots\": 0,     \"occupied_slots\": 0,     \"used_slots\": 0,     \"queued_slots\": 0,     \"open_slots\": 0 } ```  ### Update Mask  Update mask is available as a query parameter in patch endpoints. It is used to notify the API which fields you want to update. Using `update_mask` makes it easier to update objects by helping the server know which fields to update in an object instead of updating all fields. The update request ignores any fields that aren't specified in the field mask, leaving them with their current values.  Example: ```   resource = request.get('/resource/my-id').json()   resource['my_field'] = 'new-value'   request.patch('/resource/my-id?update_mask=my_field', data=json.dumps(resource)) ```  ## Versioning and Endpoint Lifecycle  - API versioning is not synchronized to specific releases of the Apache Airflow. - APIs are designed to be backward compatible. - Any changes to the API will first go through a deprecation phase.  # Summary of Changes  | Airflow version | Description | |-|-| | v2.0 | Initial release | | v2.0.2    | Added /plugins endpoint | | v2.1 | New providers endpoint |  # Trying the API  You can use a third party client, such as [curl](https://curl.haxx.se/), [HTTPie](https://httpie.org/), [Postman](https://www.postman.com/) or [the Insomnia rest client](https://insomnia.rest/) to test the Apache Airflow API.  Note that you will need to pass credentials data.  For e.g., here is how to pause a DAG with [curl](https://curl.haxx.se/), when basic authorization is used: ```bash curl -X POST 'https://example.com/api/v1/dags/{dag_id}?update_mask=is_paused' \\ -H 'Content-Type: application/json' \\ --user \"username:password\" \\ -d '{     \"is_paused\": true }' ```  Using a graphical tool such as [Postman](https://www.postman.com/) or [Insomnia](https://insomnia.rest/), it is possible to import the API specifications directly:  1. Download the API specification by clicking the **Download** button at top of this document 2. Import the JSON specification in the graphical tool of your choice.   - In *Postman*, you can click the **import** button at the top   - With *Insomnia*, you can just drag-and-drop the file on the UI  Note that with *Postman*, you can also generate code snippets by selecting a request and clicking on the **Code** button.  ## Enabling CORS  [Cross-origin resource sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) is a browser security feature that restricts HTTP requests that are initiated from scripts running in the browser.  For details on enabling/configuring CORS, see [Enabling CORS](https://airflow.apache.org/docs/apache-airflow/stable/security/api.html).  # Authentication  To be able to meet the requirements of many organizations, Airflow supports many authentication methods, and it is even possible to add your own method.  If you want to check which auth backend is currently set, you can use `airflow config get-value api auth_backend` command as in the example below. ```bash $ airflow config get-value api auth_backend airflow.api.auth.backend.basic_auth ``` The default is to deny all requests.  For details on configuring the authentication, see [API Authorization](https://airflow.apache.org/docs/apache-airflow/stable/security/api.html).  # Errors  We follow the error response format proposed in [RFC 7807](https://tools.ietf.org/html/rfc7807) also known as Problem Details for HTTP APIs. As with our normal API responses, your client must be prepared to gracefully handle additional members of the response.  ## Unauthenticated  This indicates that the request has not been applied because it lacks valid authentication credentials for the target resource. Please check that you have valid credentials.  ## PermissionDenied  This response means that the server understood the request but refuses to authorize it because it lacks sufficient rights to the resource. It happens when you do not have the necessary permission to execute the action you performed. You need to get the appropriate permissions in other to resolve this error.  ## BadRequest  This response means that the server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing). To resolve this, please ensure that your syntax is correct.  ## NotFound  This client error response indicates that the server cannot find the requested resource.  ## MethodNotAllowed  Indicates that the request method is known by the server but is not supported by the target resource.  ## NotAcceptable  The target resource does not have a current representation that would be acceptable to the user agent, according to the proactive negotiation header fields received in the request, and the server is unwilling to supply a default representation.  ## AlreadyExists  The request could not be completed due to a conflict with the current state of the target resource, meaning that the resource already exists  ## Unknown  This means that the server encountered an unexpected condition that prevented it from fulfilling the request. 
 *
 * API version: 1.0.0
 * Contact: dev@airflow.apache.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package airflow

import (
	"encoding/json"
)

// SLAMiss struct for SLAMiss
type SLAMiss struct {
	// The task ID.
	TaskId *string `json:"task_id,omitempty"`
	// The DAG ID.
	DagId *string `json:"dag_id,omitempty"`
	ExecutionDate *string `json:"execution_date,omitempty"`
	EmailSent *bool `json:"email_sent,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
	Description NullableString `json:"description,omitempty"`
	NotificationSent *bool `json:"notification_sent,omitempty"`
}

// NewSLAMiss instantiates a new SLAMiss object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSLAMiss() *SLAMiss {
	this := SLAMiss{}
	return &this
}

// NewSLAMissWithDefaults instantiates a new SLAMiss object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSLAMissWithDefaults() *SLAMiss {
	this := SLAMiss{}
	return &this
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *SLAMiss) GetTaskId() string {
	if o == nil || o.TaskId == nil {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLAMiss) GetTaskIdOk() (*string, bool) {
	if o == nil || o.TaskId == nil {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *SLAMiss) HasTaskId() bool {
	if o != nil && o.TaskId != nil {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *SLAMiss) SetTaskId(v string) {
	o.TaskId = &v
}

// GetDagId returns the DagId field value if set, zero value otherwise.
func (o *SLAMiss) GetDagId() string {
	if o == nil || o.DagId == nil {
		var ret string
		return ret
	}
	return *o.DagId
}

// GetDagIdOk returns a tuple with the DagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLAMiss) GetDagIdOk() (*string, bool) {
	if o == nil || o.DagId == nil {
		return nil, false
	}
	return o.DagId, true
}

// HasDagId returns a boolean if a field has been set.
func (o *SLAMiss) HasDagId() bool {
	if o != nil && o.DagId != nil {
		return true
	}

	return false
}

// SetDagId gets a reference to the given string and assigns it to the DagId field.
func (o *SLAMiss) SetDagId(v string) {
	o.DagId = &v
}

// GetExecutionDate returns the ExecutionDate field value if set, zero value otherwise.
func (o *SLAMiss) GetExecutionDate() string {
	if o == nil || o.ExecutionDate == nil {
		var ret string
		return ret
	}
	return *o.ExecutionDate
}

// GetExecutionDateOk returns a tuple with the ExecutionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLAMiss) GetExecutionDateOk() (*string, bool) {
	if o == nil || o.ExecutionDate == nil {
		return nil, false
	}
	return o.ExecutionDate, true
}

// HasExecutionDate returns a boolean if a field has been set.
func (o *SLAMiss) HasExecutionDate() bool {
	if o != nil && o.ExecutionDate != nil {
		return true
	}

	return false
}

// SetExecutionDate gets a reference to the given string and assigns it to the ExecutionDate field.
func (o *SLAMiss) SetExecutionDate(v string) {
	o.ExecutionDate = &v
}

// GetEmailSent returns the EmailSent field value if set, zero value otherwise.
func (o *SLAMiss) GetEmailSent() bool {
	if o == nil || o.EmailSent == nil {
		var ret bool
		return ret
	}
	return *o.EmailSent
}

// GetEmailSentOk returns a tuple with the EmailSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLAMiss) GetEmailSentOk() (*bool, bool) {
	if o == nil || o.EmailSent == nil {
		return nil, false
	}
	return o.EmailSent, true
}

// HasEmailSent returns a boolean if a field has been set.
func (o *SLAMiss) HasEmailSent() bool {
	if o != nil && o.EmailSent != nil {
		return true
	}

	return false
}

// SetEmailSent gets a reference to the given bool and assigns it to the EmailSent field.
func (o *SLAMiss) SetEmailSent(v bool) {
	o.EmailSent = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *SLAMiss) GetTimestamp() string {
	if o == nil || o.Timestamp == nil {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLAMiss) GetTimestampOk() (*string, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SLAMiss) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *SLAMiss) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SLAMiss) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SLAMiss) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SLAMiss) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SLAMiss) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SLAMiss) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SLAMiss) UnsetDescription() {
	o.Description.Unset()
}

// GetNotificationSent returns the NotificationSent field value if set, zero value otherwise.
func (o *SLAMiss) GetNotificationSent() bool {
	if o == nil || o.NotificationSent == nil {
		var ret bool
		return ret
	}
	return *o.NotificationSent
}

// GetNotificationSentOk returns a tuple with the NotificationSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLAMiss) GetNotificationSentOk() (*bool, bool) {
	if o == nil || o.NotificationSent == nil {
		return nil, false
	}
	return o.NotificationSent, true
}

// HasNotificationSent returns a boolean if a field has been set.
func (o *SLAMiss) HasNotificationSent() bool {
	if o != nil && o.NotificationSent != nil {
		return true
	}

	return false
}

// SetNotificationSent gets a reference to the given bool and assigns it to the NotificationSent field.
func (o *SLAMiss) SetNotificationSent(v bool) {
	o.NotificationSent = &v
}

func (o SLAMiss) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TaskId != nil {
		toSerialize["task_id"] = o.TaskId
	}
	if o.DagId != nil {
		toSerialize["dag_id"] = o.DagId
	}
	if o.ExecutionDate != nil {
		toSerialize["execution_date"] = o.ExecutionDate
	}
	if o.EmailSent != nil {
		toSerialize["email_sent"] = o.EmailSent
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.NotificationSent != nil {
		toSerialize["notification_sent"] = o.NotificationSent
	}
	return json.Marshal(toSerialize)
}

type NullableSLAMiss struct {
	value *SLAMiss
	isSet bool
}

func (v NullableSLAMiss) Get() *SLAMiss {
	return v.value
}

func (v *NullableSLAMiss) Set(val *SLAMiss) {
	v.value = val
	v.isSet = true
}

func (v NullableSLAMiss) IsSet() bool {
	return v.isSet
}

func (v *NullableSLAMiss) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSLAMiss(val *SLAMiss) *NullableSLAMiss {
	return &NullableSLAMiss{value: val, isSet: true}
}

func (v NullableSLAMiss) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSLAMiss) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


