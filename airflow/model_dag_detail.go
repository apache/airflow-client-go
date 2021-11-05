/*
Airflow API (Stable)

# Overview  To facilitate management, Apache Airflow supports a range of REST API endpoints across its objects. This section provides an overview of the API design, methods, and supported use cases.  Most of the endpoints accept `JSON` as input and return `JSON` responses. This means that you must usually add the following headers to your request: ``` Content-type: application/json Accept: application/json ```  ## Resources  The term `resource` refers to a single type of object in the Airflow metadata. An API is broken up by its endpoint's corresponding resource. The name of a resource is typically plural and expressed in camelCase. Example: `dagRuns`.  Resource names are used as part of endpoint URLs, as well as in API parameters and responses.  ## CRUD Operations  The platform supports **C**reate, **R**ead, **U**pdate, and **D**elete operations on most resources. You can review the standards for these operations and their standard parameters below.  Some endpoints have special behavior as exceptions.  ### Create  To create a resource, you typically submit an HTTP `POST` request with the resource's required metadata in the request body. The response returns a `201 Created` response code upon success with the resource's metadata, including its internal `id`, in the response body.  ### Read  The HTTP `GET` request can be used to read a resource or to list a number of resources.  A resource's `id` can be submitted in the request parameters to read a specific resource. The response usually returns a `200 OK` response code upon success, with the resource's metadata in the response body.  If a `GET` request does not include a specific resource `id`, it is treated as a list request. The response usually returns a `200 OK` response code upon success, with an object containing a list of resources' metadata in the response body.  When reading resources, some common query parameters are usually available. e.g.: ``` v1/connections?limit=25&offset=25 ```  |Query Parameter|Type|Description| |---------------|----|-----------| |limit|integer|Maximum number of objects to fetch. Usually 25 by default| |offset|integer|Offset after which to start returning objects. For use with limit query parameter.|  ### Update  Updating a resource requires the resource `id`, and is typically done using an HTTP `PATCH` request, with the fields to modify in the request body. The response usually returns a `200 OK` response code upon success, with information about the modified resource in the response body.  ### Delete  Deleting a resource requires the resource `id` and is typically executing via an HTTP `DELETE` request. The response usually returns a `204 No Content` response code upon success.  ## Conventions  - Resource names are plural and expressed in camelCase. - Names are consistent between URL parameter name and field name.  - Field names are in snake_case. ```json {     \"name\": \"string\",     \"slots\": 0,     \"occupied_slots\": 0,     \"used_slots\": 0,     \"queued_slots\": 0,     \"open_slots\": 0 } ```  ### Update Mask  Update mask is available as a query parameter in patch endpoints. It is used to notify the API which fields you want to update. Using `update_mask` makes it easier to update objects by helping the server know which fields to update in an object instead of updating all fields. The update request ignores any fields that aren't specified in the field mask, leaving them with their current values.  Example: ```   resource = request.get('/resource/my-id').json()   resource['my_field'] = 'new-value'   request.patch('/resource/my-id?update_mask=my_field', data=json.dumps(resource)) ```  ## Versioning and Endpoint Lifecycle  - API versioning is not synchronized to specific releases of the Apache Airflow. - APIs are designed to be backward compatible. - Any changes to the API will first go through a deprecation phase.  # Summary of Changes  | Airflow version | Description | |-|-| | v2.0 | Initial release | | v2.0.2    | Added /plugins endpoint | | v2.1 | New providers endpoint |  # Trying the API  You can use a third party client, such as [curl](https://curl.haxx.se/), [HTTPie](https://httpie.org/), [Postman](https://www.postman.com/) or [the Insomnia rest client](https://insomnia.rest/) to test the Apache Airflow API.  Note that you will need to pass credentials data.  For e.g., here is how to pause a DAG with [curl](https://curl.haxx.se/), when basic authorization is used: ```bash curl -X PATCH 'https://example.com/api/v1/dags/{dag_id}?update_mask=is_paused' \\ -H 'Content-Type: application/json' \\ --user \"username:password\" \\ -d '{     \"is_paused\": true }' ```  Using a graphical tool such as [Postman](https://www.postman.com/) or [Insomnia](https://insomnia.rest/), it is possible to import the API specifications directly:  1. Download the API specification by clicking the **Download** button at top of this document 2. Import the JSON specification in the graphical tool of your choice.   - In *Postman*, you can click the **import** button at the top   - With *Insomnia*, you can just drag-and-drop the file on the UI  Note that with *Postman*, you can also generate code snippets by selecting a request and clicking on the **Code** button.  ## Enabling CORS  [Cross-origin resource sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) is a browser security feature that restricts HTTP requests that are initiated from scripts running in the browser.  For details on enabling/configuring CORS, see [Enabling CORS](https://airflow.apache.org/docs/apache-airflow/stable/security/api.html).  # Authentication  To be able to meet the requirements of many organizations, Airflow supports many authentication methods, and it is even possible to add your own method.  If you want to check which auth backend is currently set, you can use `airflow config get-value api auth_backend` command as in the example below. ```bash $ airflow config get-value api auth_backend airflow.api.auth.backend.basic_auth ``` The default is to deny all requests.  For details on configuring the authentication, see [API Authorization](https://airflow.apache.org/docs/apache-airflow/stable/security/api.html).  # Errors  We follow the error response format proposed in [RFC 7807](https://tools.ietf.org/html/rfc7807) also known as Problem Details for HTTP APIs. As with our normal API responses, your client must be prepared to gracefully handle additional members of the response.  ## Unauthenticated  This indicates that the request has not been applied because it lacks valid authentication credentials for the target resource. Please check that you have valid credentials.  ## PermissionDenied  This response means that the server understood the request but refuses to authorize it because it lacks sufficient rights to the resource. It happens when you do not have the necessary permission to execute the action you performed. You need to get the appropriate permissions in other to resolve this error.  ## BadRequest  This response means that the server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing). To resolve this, please ensure that your syntax is correct.  ## NotFound  This client error response indicates that the server cannot find the requested resource.  ## MethodNotAllowed  Indicates that the request method is known by the server but is not supported by the target resource.  ## NotAcceptable  The target resource does not have a current representation that would be acceptable to the user agent, according to the proactive negotiation header fields received in the request, and the server is unwilling to supply a default representation.  ## AlreadyExists  The request could not be completed due to a conflict with the current state of the target resource, e.g. the resource it tries to create already exists.  ## Unknown  This means that the server encountered an unexpected condition that prevented it from fulfilling the request. 

API version: 1.0.0
Contact: dev@airflow.apache.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package airflow

import (
	"encoding/json"
	"time"
)

// DAGDetail DAG details.  For details see: (airflow.models.DAG)[https://airflow.apache.org/docs/apache-airflow/stable/_api/airflow/models/index.html#airflow.models.DAG] 
type DAGDetail struct {
	// The ID of the DAG.
	DagId *string `json:"dag_id,omitempty"`
	// If the DAG is SubDAG then it is the top level DAG identifier. Otherwise, null.
	RootDagId NullableString `json:"root_dag_id,omitempty"`
	// Whether the DAG is paused.
	IsPaused NullableBool `json:"is_paused,omitempty"`
	// Whether the DAG is currently seen by the scheduler(s).
	IsActive NullableBool `json:"is_active,omitempty"`
	// Whether the DAG is SubDAG.
	IsSubdag *bool `json:"is_subdag,omitempty"`
	// The absolute path to the file.
	Fileloc *string `json:"fileloc,omitempty"`
	// The key containing the encrypted path to the file. Encryption and decryption take place only on the server. This prevents the client from reading an non-DAG file. This also ensures API extensibility, because the format of encrypted data may change. 
	FileToken *string `json:"file_token,omitempty"`
	Owners *[]string `json:"owners,omitempty"`
	// User-provided DAG description, which can consist of several sentences or paragraphs that describe DAG contents. 
	Description NullableString `json:"description,omitempty"`
	ScheduleInterval *ScheduleInterval `json:"schedule_interval,omitempty"`
	// List of tags.
	Tags []Tag `json:"tags,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
	Catchup *bool `json:"catchup,omitempty"`
	Orientation *string `json:"orientation,omitempty"`
	Concurrency *float32 `json:"concurrency,omitempty"`
	StartDate NullableTime `json:"start_date,omitempty"`
	DagRunTimeout *TimeDelta `json:"dag_run_timeout,omitempty"`
	DocMd NullableString `json:"doc_md,omitempty"`
	DefaultView *string `json:"default_view,omitempty"`
	Params *map[string]interface{} `json:"params,omitempty"`
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
	return *o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGDetail) GetOwnersOk() (*[]string, bool) {
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
	o.Owners = &v
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

// GetScheduleInterval returns the ScheduleInterval field value if set, zero value otherwise.
func (o *DAGDetail) GetScheduleInterval() ScheduleInterval {
	if o == nil || o.ScheduleInterval == nil {
		var ret ScheduleInterval
		return ret
	}
	return *o.ScheduleInterval
}

// GetScheduleIntervalOk returns a tuple with the ScheduleInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGDetail) GetScheduleIntervalOk() (*ScheduleInterval, bool) {
	if o == nil || o.ScheduleInterval == nil {
		return nil, false
	}
	return o.ScheduleInterval, true
}

// HasScheduleInterval returns a boolean if a field has been set.
func (o *DAGDetail) HasScheduleInterval() bool {
	if o != nil && o.ScheduleInterval != nil {
		return true
	}

	return false
}

// SetScheduleInterval gets a reference to the given ScheduleInterval and assigns it to the ScheduleInterval field.
func (o *DAGDetail) SetScheduleInterval(v ScheduleInterval) {
	o.ScheduleInterval = &v
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
func (o *DAGDetail) GetTagsOk() (*[]Tag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
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

// GetDagRunTimeout returns the DagRunTimeout field value if set, zero value otherwise.
func (o *DAGDetail) GetDagRunTimeout() TimeDelta {
	if o == nil || o.DagRunTimeout == nil {
		var ret TimeDelta
		return ret
	}
	return *o.DagRunTimeout
}

// GetDagRunTimeoutOk returns a tuple with the DagRunTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGDetail) GetDagRunTimeoutOk() (*TimeDelta, bool) {
	if o == nil || o.DagRunTimeout == nil {
		return nil, false
	}
	return o.DagRunTimeout, true
}

// HasDagRunTimeout returns a boolean if a field has been set.
func (o *DAGDetail) HasDagRunTimeout() bool {
	if o != nil && o.DagRunTimeout != nil {
		return true
	}

	return false
}

// SetDagRunTimeout gets a reference to the given TimeDelta and assigns it to the DagRunTimeout field.
func (o *DAGDetail) SetDagRunTimeout(v TimeDelta) {
	o.DagRunTimeout = &v
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

// GetParams returns the Params field value if set, zero value otherwise.
func (o *DAGDetail) GetParams() map[string]interface{} {
	if o == nil || o.Params == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DAGDetail) GetParamsOk() (*map[string]interface{}, bool) {
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
	o.Params = &v
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
	if o.ScheduleInterval != nil {
		toSerialize["schedule_interval"] = o.ScheduleInterval
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
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
	if o.DagRunTimeout != nil {
		toSerialize["dag_run_timeout"] = o.DagRunTimeout
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


