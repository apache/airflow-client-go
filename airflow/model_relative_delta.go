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

// RelativeDelta Relative delta
type RelativeDelta struct {
	Type string `json:"__type"`
	Years int32 `json:"years"`
	Months int32 `json:"months"`
	Days int32 `json:"days"`
	Leapdays int32 `json:"leapdays"`
	Hours int32 `json:"hours"`
	Minutes int32 `json:"minutes"`
	Seconds int32 `json:"seconds"`
	Microseconds int32 `json:"microseconds"`
	Year int32 `json:"year"`
	Month int32 `json:"month"`
	Day int32 `json:"day"`
	Hour int32 `json:"hour"`
	Minute int32 `json:"minute"`
	Second int32 `json:"second"`
	Microsecond int32 `json:"microsecond"`
}

// NewRelativeDelta instantiates a new RelativeDelta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelativeDelta(type_ string, years int32, months int32, days int32, leapdays int32, hours int32, minutes int32, seconds int32, microseconds int32, year int32, month int32, day int32, hour int32, minute int32, second int32, microsecond int32) *RelativeDelta {
	this := RelativeDelta{}
	this.Type = type_
	this.Years = years
	this.Months = months
	this.Days = days
	this.Leapdays = leapdays
	this.Hours = hours
	this.Minutes = minutes
	this.Seconds = seconds
	this.Microseconds = microseconds
	this.Year = year
	this.Month = month
	this.Day = day
	this.Hour = hour
	this.Minute = minute
	this.Second = second
	this.Microsecond = microsecond
	return &this
}

// NewRelativeDeltaWithDefaults instantiates a new RelativeDelta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelativeDeltaWithDefaults() *RelativeDelta {
	this := RelativeDelta{}
	return &this
}

// GetType returns the Type field value
func (o *RelativeDelta) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RelativeDelta) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RelativeDelta) SetType(v string) {
	o.Type = v
}

// GetYears returns the Years field value
func (o *RelativeDelta) GetYears() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Years
}

// GetYearsOk returns a tuple with the Years field value
// and a boolean to check if the value has been set.
func (o *RelativeDelta) GetYearsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Years, true
}

// SetYears sets field value
func (o *RelativeDelta) SetYears(v int32) {
	o.Years = v
}

// GetMonths returns the Months field value
func (o *RelativeDelta) GetMonths() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Months
}

// GetMonthsOk returns a tuple with the Months field value
// and a boolean to check if the value has been set.
func (o *RelativeDelta) GetMonthsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Months, true
}

// SetMonths sets field value
func (o *RelativeDelta) SetMonths(v int32) {
	o.Months = v
}

// GetDays returns the Days field value
func (o *RelativeDelta) GetDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Days
}

// GetDaysOk returns a tuple with the Days field value
// and a boolean to check if the value has been set.
func (o *RelativeDelta) GetDaysOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Days, true
}

// SetDays sets field value
func (o *RelativeDelta) SetDays(v int32) {
	o.Days = v
}

// GetLeapdays returns the Leapdays field value
func (o *RelativeDelta) GetLeapdays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Leapdays
}

// GetLeapdaysOk returns a tuple with the Leapdays field value
// and a boolean to check if the value has been set.
func (o *RelativeDelta) GetLeapdaysOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Leapdays, true
}

// SetLeapdays sets field value
func (o *RelativeDelta) SetLeapdays(v int32) {
	o.Leapdays = v
}

// GetHours returns the Hours field value
func (o *RelativeDelta) GetHours() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Hours
}

// GetHoursOk returns a tuple with the Hours field value
// and a boolean to check if the value has been set.
func (o *RelativeDelta) GetHoursOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Hours, true
}

// SetHours sets field value
func (o *RelativeDelta) SetHours(v int32) {
	o.Hours = v
}

// GetMinutes returns the Minutes field value
func (o *RelativeDelta) GetMinutes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Minutes
}

// GetMinutesOk returns a tuple with the Minutes field value
// and a boolean to check if the value has been set.
func (o *RelativeDelta) GetMinutesOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Minutes, true
}

// SetMinutes sets field value
func (o *RelativeDelta) SetMinutes(v int32) {
	o.Minutes = v
}

// GetSeconds returns the Seconds field value
func (o *RelativeDelta) GetSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Seconds
}

// GetSecondsOk returns a tuple with the Seconds field value
// and a boolean to check if the value has been set.
func (o *RelativeDelta) GetSecondsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Seconds, true
}

// SetSeconds sets field value
func (o *RelativeDelta) SetSeconds(v int32) {
	o.Seconds = v
}

// GetMicroseconds returns the Microseconds field value
func (o *RelativeDelta) GetMicroseconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Microseconds
}

// GetMicrosecondsOk returns a tuple with the Microseconds field value
// and a boolean to check if the value has been set.
func (o *RelativeDelta) GetMicrosecondsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Microseconds, true
}

// SetMicroseconds sets field value
func (o *RelativeDelta) SetMicroseconds(v int32) {
	o.Microseconds = v
}

// GetYear returns the Year field value
func (o *RelativeDelta) GetYear() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Year
}

// GetYearOk returns a tuple with the Year field value
// and a boolean to check if the value has been set.
func (o *RelativeDelta) GetYearOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Year, true
}

// SetYear sets field value
func (o *RelativeDelta) SetYear(v int32) {
	o.Year = v
}

// GetMonth returns the Month field value
func (o *RelativeDelta) GetMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Month
}

// GetMonthOk returns a tuple with the Month field value
// and a boolean to check if the value has been set.
func (o *RelativeDelta) GetMonthOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Month, true
}

// SetMonth sets field value
func (o *RelativeDelta) SetMonth(v int32) {
	o.Month = v
}

// GetDay returns the Day field value
func (o *RelativeDelta) GetDay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Day
}

// GetDayOk returns a tuple with the Day field value
// and a boolean to check if the value has been set.
func (o *RelativeDelta) GetDayOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Day, true
}

// SetDay sets field value
func (o *RelativeDelta) SetDay(v int32) {
	o.Day = v
}

// GetHour returns the Hour field value
func (o *RelativeDelta) GetHour() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Hour
}

// GetHourOk returns a tuple with the Hour field value
// and a boolean to check if the value has been set.
func (o *RelativeDelta) GetHourOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Hour, true
}

// SetHour sets field value
func (o *RelativeDelta) SetHour(v int32) {
	o.Hour = v
}

// GetMinute returns the Minute field value
func (o *RelativeDelta) GetMinute() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Minute
}

// GetMinuteOk returns a tuple with the Minute field value
// and a boolean to check if the value has been set.
func (o *RelativeDelta) GetMinuteOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Minute, true
}

// SetMinute sets field value
func (o *RelativeDelta) SetMinute(v int32) {
	o.Minute = v
}

// GetSecond returns the Second field value
func (o *RelativeDelta) GetSecond() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Second
}

// GetSecondOk returns a tuple with the Second field value
// and a boolean to check if the value has been set.
func (o *RelativeDelta) GetSecondOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Second, true
}

// SetSecond sets field value
func (o *RelativeDelta) SetSecond(v int32) {
	o.Second = v
}

// GetMicrosecond returns the Microsecond field value
func (o *RelativeDelta) GetMicrosecond() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Microsecond
}

// GetMicrosecondOk returns a tuple with the Microsecond field value
// and a boolean to check if the value has been set.
func (o *RelativeDelta) GetMicrosecondOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Microsecond, true
}

// SetMicrosecond sets field value
func (o *RelativeDelta) SetMicrosecond(v int32) {
	o.Microsecond = v
}

func (o RelativeDelta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["__type"] = o.Type
	}
	if true {
		toSerialize["years"] = o.Years
	}
	if true {
		toSerialize["months"] = o.Months
	}
	if true {
		toSerialize["days"] = o.Days
	}
	if true {
		toSerialize["leapdays"] = o.Leapdays
	}
	if true {
		toSerialize["hours"] = o.Hours
	}
	if true {
		toSerialize["minutes"] = o.Minutes
	}
	if true {
		toSerialize["seconds"] = o.Seconds
	}
	if true {
		toSerialize["microseconds"] = o.Microseconds
	}
	if true {
		toSerialize["year"] = o.Year
	}
	if true {
		toSerialize["month"] = o.Month
	}
	if true {
		toSerialize["day"] = o.Day
	}
	if true {
		toSerialize["hour"] = o.Hour
	}
	if true {
		toSerialize["minute"] = o.Minute
	}
	if true {
		toSerialize["second"] = o.Second
	}
	if true {
		toSerialize["microsecond"] = o.Microsecond
	}
	return json.Marshal(toSerialize)
}

type NullableRelativeDelta struct {
	value *RelativeDelta
	isSet bool
}

func (v NullableRelativeDelta) Get() *RelativeDelta {
	return v.value
}

func (v *NullableRelativeDelta) Set(val *RelativeDelta) {
	v.value = val
	v.isSet = true
}

func (v NullableRelativeDelta) IsSet() bool {
	return v.isSet
}

func (v *NullableRelativeDelta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelativeDelta(val *RelativeDelta) *NullableRelativeDelta {
	return &NullableRelativeDelta{value: val, isSet: true}
}

func (v NullableRelativeDelta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelativeDelta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


