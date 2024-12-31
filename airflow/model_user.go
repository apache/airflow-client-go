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

// User A user object with sensitive data.  *New in version 2.1.0* 
type User struct {
	// The user's first name.  *Changed in version 2.4.0*&#58; The requirement for this to be non-empty was removed. 
	FirstName *string `json:"first_name,omitempty"`
	// The user's last name.  *Changed in version 2.4.0*&#58; The requirement for this to be non-empty was removed. 
	LastName *string `json:"last_name,omitempty"`
	// The username.  *Changed in version 2.2.0*&#58; A minimum character length requirement ('minLength') is added. 
	Username *string `json:"username,omitempty"`
	// The user's email.  *Changed in version 2.2.0*&#58; A minimum character length requirement ('minLength') is added. 
	Email *string `json:"email,omitempty"`
	// Whether the user is active
	Active NullableBool `json:"active,omitempty"`
	// The last user login
	LastLogin NullableString `json:"last_login,omitempty"`
	// The login count
	LoginCount NullableInt32 `json:"login_count,omitempty"`
	// The number of times the login failed
	FailedLoginCount NullableInt32 `json:"failed_login_count,omitempty"`
	// User roles.  *Changed in version 2.2.0*&#58; Field is no longer read-only. 
	Roles []UserCollectionItemRoles `json:"roles,omitempty"`
	// The date user was created
	CreatedOn NullableString `json:"created_on,omitempty"`
	// The date user was changed
	ChangedOn NullableString `json:"changed_on,omitempty"`
	Password *string `json:"password,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *User) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *User) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *User) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *User) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *User) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *User) SetLastName(v string) {
	o.LastName = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *User) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *User) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *User) SetUsername(v string) {
	o.Username = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *User) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *User) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *User) SetEmail(v string) {
	o.Email = &v
}

// GetActive returns the Active field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetActive() bool {
	if o == nil || o.Active.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Active.Get()
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetActiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Active.Get(), o.Active.IsSet()
}

// HasActive returns a boolean if a field has been set.
func (o *User) HasActive() bool {
	if o != nil && o.Active.IsSet() {
		return true
	}

	return false
}

// SetActive gets a reference to the given NullableBool and assigns it to the Active field.
func (o *User) SetActive(v bool) {
	o.Active.Set(&v)
}
// SetActiveNil sets the value for Active to be an explicit nil
func (o *User) SetActiveNil() {
	o.Active.Set(nil)
}

// UnsetActive ensures that no value is present for Active, not even an explicit nil
func (o *User) UnsetActive() {
	o.Active.Unset()
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetLastLogin() string {
	if o == nil || o.LastLogin.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastLogin.Get()
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetLastLoginOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastLogin.Get(), o.LastLogin.IsSet()
}

// HasLastLogin returns a boolean if a field has been set.
func (o *User) HasLastLogin() bool {
	if o != nil && o.LastLogin.IsSet() {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given NullableString and assigns it to the LastLogin field.
func (o *User) SetLastLogin(v string) {
	o.LastLogin.Set(&v)
}
// SetLastLoginNil sets the value for LastLogin to be an explicit nil
func (o *User) SetLastLoginNil() {
	o.LastLogin.Set(nil)
}

// UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
func (o *User) UnsetLastLogin() {
	o.LastLogin.Unset()
}

// GetLoginCount returns the LoginCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetLoginCount() int32 {
	if o == nil || o.LoginCount.Get() == nil {
		var ret int32
		return ret
	}
	return *o.LoginCount.Get()
}

// GetLoginCountOk returns a tuple with the LoginCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetLoginCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LoginCount.Get(), o.LoginCount.IsSet()
}

// HasLoginCount returns a boolean if a field has been set.
func (o *User) HasLoginCount() bool {
	if o != nil && o.LoginCount.IsSet() {
		return true
	}

	return false
}

// SetLoginCount gets a reference to the given NullableInt32 and assigns it to the LoginCount field.
func (o *User) SetLoginCount(v int32) {
	o.LoginCount.Set(&v)
}
// SetLoginCountNil sets the value for LoginCount to be an explicit nil
func (o *User) SetLoginCountNil() {
	o.LoginCount.Set(nil)
}

// UnsetLoginCount ensures that no value is present for LoginCount, not even an explicit nil
func (o *User) UnsetLoginCount() {
	o.LoginCount.Unset()
}

// GetFailedLoginCount returns the FailedLoginCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetFailedLoginCount() int32 {
	if o == nil || o.FailedLoginCount.Get() == nil {
		var ret int32
		return ret
	}
	return *o.FailedLoginCount.Get()
}

// GetFailedLoginCountOk returns a tuple with the FailedLoginCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetFailedLoginCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FailedLoginCount.Get(), o.FailedLoginCount.IsSet()
}

// HasFailedLoginCount returns a boolean if a field has been set.
func (o *User) HasFailedLoginCount() bool {
	if o != nil && o.FailedLoginCount.IsSet() {
		return true
	}

	return false
}

// SetFailedLoginCount gets a reference to the given NullableInt32 and assigns it to the FailedLoginCount field.
func (o *User) SetFailedLoginCount(v int32) {
	o.FailedLoginCount.Set(&v)
}
// SetFailedLoginCountNil sets the value for FailedLoginCount to be an explicit nil
func (o *User) SetFailedLoginCountNil() {
	o.FailedLoginCount.Set(nil)
}

// UnsetFailedLoginCount ensures that no value is present for FailedLoginCount, not even an explicit nil
func (o *User) UnsetFailedLoginCount() {
	o.FailedLoginCount.Unset()
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *User) GetRoles() []UserCollectionItemRoles {
	if o == nil || o.Roles == nil {
		var ret []UserCollectionItemRoles
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetRolesOk() ([]UserCollectionItemRoles, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *User) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []UserCollectionItemRoles and assigns it to the Roles field.
func (o *User) SetRoles(v []UserCollectionItemRoles) {
	o.Roles = v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetCreatedOn() string {
	if o == nil || o.CreatedOn.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatedOn.Get()
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetCreatedOnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedOn.Get(), o.CreatedOn.IsSet()
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *User) HasCreatedOn() bool {
	if o != nil && o.CreatedOn.IsSet() {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given NullableString and assigns it to the CreatedOn field.
func (o *User) SetCreatedOn(v string) {
	o.CreatedOn.Set(&v)
}
// SetCreatedOnNil sets the value for CreatedOn to be an explicit nil
func (o *User) SetCreatedOnNil() {
	o.CreatedOn.Set(nil)
}

// UnsetCreatedOn ensures that no value is present for CreatedOn, not even an explicit nil
func (o *User) UnsetCreatedOn() {
	o.CreatedOn.Unset()
}

// GetChangedOn returns the ChangedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetChangedOn() string {
	if o == nil || o.ChangedOn.Get() == nil {
		var ret string
		return ret
	}
	return *o.ChangedOn.Get()
}

// GetChangedOnOk returns a tuple with the ChangedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetChangedOnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ChangedOn.Get(), o.ChangedOn.IsSet()
}

// HasChangedOn returns a boolean if a field has been set.
func (o *User) HasChangedOn() bool {
	if o != nil && o.ChangedOn.IsSet() {
		return true
	}

	return false
}

// SetChangedOn gets a reference to the given NullableString and assigns it to the ChangedOn field.
func (o *User) SetChangedOn(v string) {
	o.ChangedOn.Set(&v)
}
// SetChangedOnNil sets the value for ChangedOn to be an explicit nil
func (o *User) SetChangedOnNil() {
	o.ChangedOn.Set(nil)
}

// UnsetChangedOn ensures that no value is present for ChangedOn, not even an explicit nil
func (o *User) UnsetChangedOn() {
	o.ChangedOn.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *User) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *User) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *User) SetPassword(v string) {
	o.Password = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Active.IsSet() {
		toSerialize["active"] = o.Active.Get()
	}
	if o.LastLogin.IsSet() {
		toSerialize["last_login"] = o.LastLogin.Get()
	}
	if o.LoginCount.IsSet() {
		toSerialize["login_count"] = o.LoginCount.Get()
	}
	if o.FailedLoginCount.IsSet() {
		toSerialize["failed_login_count"] = o.FailedLoginCount.Get()
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.CreatedOn.IsSet() {
		toSerialize["created_on"] = o.CreatedOn.Get()
	}
	if o.ChangedOn.IsSet() {
		toSerialize["changed_on"] = o.ChangedOn.Get()
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


