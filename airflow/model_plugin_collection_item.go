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

// PluginCollectionItem A plugin Item.  *New in version 2.1.0* 
type PluginCollectionItem struct {
	// The name of the plugin
	Name *string `json:"name,omitempty"`
	// The plugin hooks
	Hooks []*string `json:"hooks,omitempty"`
	// The plugin executors
	Executors []*string `json:"executors,omitempty"`
	// The plugin macros
	Macros []*map[string]interface{} `json:"macros,omitempty"`
	// The flask blueprints
	FlaskBlueprints []*map[string]interface{} `json:"flask_blueprints,omitempty"`
	// The appuilder views
	AppbuilderViews []*map[string]interface{} `json:"appbuilder_views,omitempty"`
	// The Flask Appbuilder menu items
	AppbuilderMenuItems []*map[string]interface{} `json:"appbuilder_menu_items,omitempty"`
	// The global operator extra links
	GlobalOperatorExtraLinks []*map[string]interface{} `json:"global_operator_extra_links,omitempty"`
	// Operator extra links
	OperatorExtraLinks []*map[string]interface{} `json:"operator_extra_links,omitempty"`
	// The plugin source
	Source NullableString `json:"source,omitempty"`
}

// NewPluginCollectionItem instantiates a new PluginCollectionItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginCollectionItem() *PluginCollectionItem {
	this := PluginCollectionItem{}
	return &this
}

// NewPluginCollectionItemWithDefaults instantiates a new PluginCollectionItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginCollectionItemWithDefaults() *PluginCollectionItem {
	this := PluginCollectionItem{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PluginCollectionItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCollectionItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PluginCollectionItem) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PluginCollectionItem) SetName(v string) {
	o.Name = &v
}

// GetHooks returns the Hooks field value if set, zero value otherwise.
func (o *PluginCollectionItem) GetHooks() []*string {
	if o == nil || o.Hooks == nil {
		var ret []*string
		return ret
	}
	return o.Hooks
}

// GetHooksOk returns a tuple with the Hooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCollectionItem) GetHooksOk() ([]*string, bool) {
	if o == nil || o.Hooks == nil {
		return nil, false
	}
	return o.Hooks, true
}

// HasHooks returns a boolean if a field has been set.
func (o *PluginCollectionItem) HasHooks() bool {
	if o != nil && o.Hooks != nil {
		return true
	}

	return false
}

// SetHooks gets a reference to the given []*string and assigns it to the Hooks field.
func (o *PluginCollectionItem) SetHooks(v []*string) {
	o.Hooks = v
}

// GetExecutors returns the Executors field value if set, zero value otherwise.
func (o *PluginCollectionItem) GetExecutors() []*string {
	if o == nil || o.Executors == nil {
		var ret []*string
		return ret
	}
	return o.Executors
}

// GetExecutorsOk returns a tuple with the Executors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCollectionItem) GetExecutorsOk() ([]*string, bool) {
	if o == nil || o.Executors == nil {
		return nil, false
	}
	return o.Executors, true
}

// HasExecutors returns a boolean if a field has been set.
func (o *PluginCollectionItem) HasExecutors() bool {
	if o != nil && o.Executors != nil {
		return true
	}

	return false
}

// SetExecutors gets a reference to the given []*string and assigns it to the Executors field.
func (o *PluginCollectionItem) SetExecutors(v []*string) {
	o.Executors = v
}

// GetMacros returns the Macros field value if set, zero value otherwise.
func (o *PluginCollectionItem) GetMacros() []*map[string]interface{} {
	if o == nil || o.Macros == nil {
		var ret []*map[string]interface{}
		return ret
	}
	return o.Macros
}

// GetMacrosOk returns a tuple with the Macros field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCollectionItem) GetMacrosOk() ([]*map[string]interface{}, bool) {
	if o == nil || o.Macros == nil {
		return nil, false
	}
	return o.Macros, true
}

// HasMacros returns a boolean if a field has been set.
func (o *PluginCollectionItem) HasMacros() bool {
	if o != nil && o.Macros != nil {
		return true
	}

	return false
}

// SetMacros gets a reference to the given []*map[string]interface{} and assigns it to the Macros field.
func (o *PluginCollectionItem) SetMacros(v []*map[string]interface{}) {
	o.Macros = v
}

// GetFlaskBlueprints returns the FlaskBlueprints field value if set, zero value otherwise.
func (o *PluginCollectionItem) GetFlaskBlueprints() []*map[string]interface{} {
	if o == nil || o.FlaskBlueprints == nil {
		var ret []*map[string]interface{}
		return ret
	}
	return o.FlaskBlueprints
}

// GetFlaskBlueprintsOk returns a tuple with the FlaskBlueprints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCollectionItem) GetFlaskBlueprintsOk() ([]*map[string]interface{}, bool) {
	if o == nil || o.FlaskBlueprints == nil {
		return nil, false
	}
	return o.FlaskBlueprints, true
}

// HasFlaskBlueprints returns a boolean if a field has been set.
func (o *PluginCollectionItem) HasFlaskBlueprints() bool {
	if o != nil && o.FlaskBlueprints != nil {
		return true
	}

	return false
}

// SetFlaskBlueprints gets a reference to the given []*map[string]interface{} and assigns it to the FlaskBlueprints field.
func (o *PluginCollectionItem) SetFlaskBlueprints(v []*map[string]interface{}) {
	o.FlaskBlueprints = v
}

// GetAppbuilderViews returns the AppbuilderViews field value if set, zero value otherwise.
func (o *PluginCollectionItem) GetAppbuilderViews() []*map[string]interface{} {
	if o == nil || o.AppbuilderViews == nil {
		var ret []*map[string]interface{}
		return ret
	}
	return o.AppbuilderViews
}

// GetAppbuilderViewsOk returns a tuple with the AppbuilderViews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCollectionItem) GetAppbuilderViewsOk() ([]*map[string]interface{}, bool) {
	if o == nil || o.AppbuilderViews == nil {
		return nil, false
	}
	return o.AppbuilderViews, true
}

// HasAppbuilderViews returns a boolean if a field has been set.
func (o *PluginCollectionItem) HasAppbuilderViews() bool {
	if o != nil && o.AppbuilderViews != nil {
		return true
	}

	return false
}

// SetAppbuilderViews gets a reference to the given []*map[string]interface{} and assigns it to the AppbuilderViews field.
func (o *PluginCollectionItem) SetAppbuilderViews(v []*map[string]interface{}) {
	o.AppbuilderViews = v
}

// GetAppbuilderMenuItems returns the AppbuilderMenuItems field value if set, zero value otherwise.
func (o *PluginCollectionItem) GetAppbuilderMenuItems() []*map[string]interface{} {
	if o == nil || o.AppbuilderMenuItems == nil {
		var ret []*map[string]interface{}
		return ret
	}
	return o.AppbuilderMenuItems
}

// GetAppbuilderMenuItemsOk returns a tuple with the AppbuilderMenuItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCollectionItem) GetAppbuilderMenuItemsOk() ([]*map[string]interface{}, bool) {
	if o == nil || o.AppbuilderMenuItems == nil {
		return nil, false
	}
	return o.AppbuilderMenuItems, true
}

// HasAppbuilderMenuItems returns a boolean if a field has been set.
func (o *PluginCollectionItem) HasAppbuilderMenuItems() bool {
	if o != nil && o.AppbuilderMenuItems != nil {
		return true
	}

	return false
}

// SetAppbuilderMenuItems gets a reference to the given []*map[string]interface{} and assigns it to the AppbuilderMenuItems field.
func (o *PluginCollectionItem) SetAppbuilderMenuItems(v []*map[string]interface{}) {
	o.AppbuilderMenuItems = v
}

// GetGlobalOperatorExtraLinks returns the GlobalOperatorExtraLinks field value if set, zero value otherwise.
func (o *PluginCollectionItem) GetGlobalOperatorExtraLinks() []*map[string]interface{} {
	if o == nil || o.GlobalOperatorExtraLinks == nil {
		var ret []*map[string]interface{}
		return ret
	}
	return o.GlobalOperatorExtraLinks
}

// GetGlobalOperatorExtraLinksOk returns a tuple with the GlobalOperatorExtraLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCollectionItem) GetGlobalOperatorExtraLinksOk() ([]*map[string]interface{}, bool) {
	if o == nil || o.GlobalOperatorExtraLinks == nil {
		return nil, false
	}
	return o.GlobalOperatorExtraLinks, true
}

// HasGlobalOperatorExtraLinks returns a boolean if a field has been set.
func (o *PluginCollectionItem) HasGlobalOperatorExtraLinks() bool {
	if o != nil && o.GlobalOperatorExtraLinks != nil {
		return true
	}

	return false
}

// SetGlobalOperatorExtraLinks gets a reference to the given []*map[string]interface{} and assigns it to the GlobalOperatorExtraLinks field.
func (o *PluginCollectionItem) SetGlobalOperatorExtraLinks(v []*map[string]interface{}) {
	o.GlobalOperatorExtraLinks = v
}

// GetOperatorExtraLinks returns the OperatorExtraLinks field value if set, zero value otherwise.
func (o *PluginCollectionItem) GetOperatorExtraLinks() []*map[string]interface{} {
	if o == nil || o.OperatorExtraLinks == nil {
		var ret []*map[string]interface{}
		return ret
	}
	return o.OperatorExtraLinks
}

// GetOperatorExtraLinksOk returns a tuple with the OperatorExtraLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCollectionItem) GetOperatorExtraLinksOk() ([]*map[string]interface{}, bool) {
	if o == nil || o.OperatorExtraLinks == nil {
		return nil, false
	}
	return o.OperatorExtraLinks, true
}

// HasOperatorExtraLinks returns a boolean if a field has been set.
func (o *PluginCollectionItem) HasOperatorExtraLinks() bool {
	if o != nil && o.OperatorExtraLinks != nil {
		return true
	}

	return false
}

// SetOperatorExtraLinks gets a reference to the given []*map[string]interface{} and assigns it to the OperatorExtraLinks field.
func (o *PluginCollectionItem) SetOperatorExtraLinks(v []*map[string]interface{}) {
	o.OperatorExtraLinks = v
}

// GetSource returns the Source field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PluginCollectionItem) GetSource() string {
	if o == nil || o.Source.Get() == nil {
		var ret string
		return ret
	}
	return *o.Source.Get()
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PluginCollectionItem) GetSourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Source.Get(), o.Source.IsSet()
}

// HasSource returns a boolean if a field has been set.
func (o *PluginCollectionItem) HasSource() bool {
	if o != nil && o.Source.IsSet() {
		return true
	}

	return false
}

// SetSource gets a reference to the given NullableString and assigns it to the Source field.
func (o *PluginCollectionItem) SetSource(v string) {
	o.Source.Set(&v)
}
// SetSourceNil sets the value for Source to be an explicit nil
func (o *PluginCollectionItem) SetSourceNil() {
	o.Source.Set(nil)
}

// UnsetSource ensures that no value is present for Source, not even an explicit nil
func (o *PluginCollectionItem) UnsetSource() {
	o.Source.Unset()
}

func (o PluginCollectionItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Hooks != nil {
		toSerialize["hooks"] = o.Hooks
	}
	if o.Executors != nil {
		toSerialize["executors"] = o.Executors
	}
	if o.Macros != nil {
		toSerialize["macros"] = o.Macros
	}
	if o.FlaskBlueprints != nil {
		toSerialize["flask_blueprints"] = o.FlaskBlueprints
	}
	if o.AppbuilderViews != nil {
		toSerialize["appbuilder_views"] = o.AppbuilderViews
	}
	if o.AppbuilderMenuItems != nil {
		toSerialize["appbuilder_menu_items"] = o.AppbuilderMenuItems
	}
	if o.GlobalOperatorExtraLinks != nil {
		toSerialize["global_operator_extra_links"] = o.GlobalOperatorExtraLinks
	}
	if o.OperatorExtraLinks != nil {
		toSerialize["operator_extra_links"] = o.OperatorExtraLinks
	}
	if o.Source.IsSet() {
		toSerialize["source"] = o.Source.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePluginCollectionItem struct {
	value *PluginCollectionItem
	isSet bool
}

func (v NullablePluginCollectionItem) Get() *PluginCollectionItem {
	return v.value
}

func (v *NullablePluginCollectionItem) Set(val *PluginCollectionItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginCollectionItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginCollectionItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginCollectionItem(val *PluginCollectionItem) *NullablePluginCollectionItem {
	return &NullablePluginCollectionItem{value: val, isSet: true}
}

func (v NullablePluginCollectionItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginCollectionItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


