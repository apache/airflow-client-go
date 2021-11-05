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
)

// ClearTaskInstance struct for ClearTaskInstance
type ClearTaskInstance struct {
	// If set, don't actually run this operation. The response will contain a list of task instances planned to be cleaned, but not modified in any way. 
	DryRun *bool `json:"dry_run,omitempty"`
	// A list of task ids to clear.
	TaskIds *[]string `json:"task_ids,omitempty"`
	// The minimum execution date to clear.
	StartDate *string `json:"start_date,omitempty"`
	// The maximum execution date to clear.
	EndDate *string `json:"end_date,omitempty"`
	// Only clear failed tasks.
	OnlyFailed *bool `json:"only_failed,omitempty"`
	// Only clear running tasks.
	OnlyRunning *bool `json:"only_running,omitempty"`
	// Clear tasks in subdags and clear external tasks indicated by ExternalTaskMarker.
	IncludeSubdags *bool `json:"include_subdags,omitempty"`
	// Clear tasks in the parent dag of the subdag.
	IncludeParentdag *bool `json:"include_parentdag,omitempty"`
	// Set state of DAG runs to RUNNING.
	ResetDagRuns *bool `json:"reset_dag_runs,omitempty"`
}

// NewClearTaskInstance instantiates a new ClearTaskInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClearTaskInstance() *ClearTaskInstance {
	this := ClearTaskInstance{}
	var dryRun bool = true
	this.DryRun = &dryRun
	var onlyFailed bool = true
	this.OnlyFailed = &onlyFailed
	var onlyRunning bool = false
	this.OnlyRunning = &onlyRunning
	return &this
}

// NewClearTaskInstanceWithDefaults instantiates a new ClearTaskInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClearTaskInstanceWithDefaults() *ClearTaskInstance {
	this := ClearTaskInstance{}
	var dryRun bool = true
	this.DryRun = &dryRun
	var onlyFailed bool = true
	this.OnlyFailed = &onlyFailed
	var onlyRunning bool = false
	this.OnlyRunning = &onlyRunning
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ClearTaskInstance) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClearTaskInstance) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ClearTaskInstance) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ClearTaskInstance) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetTaskIds returns the TaskIds field value if set, zero value otherwise.
func (o *ClearTaskInstance) GetTaskIds() []string {
	if o == nil || o.TaskIds == nil {
		var ret []string
		return ret
	}
	return *o.TaskIds
}

// GetTaskIdsOk returns a tuple with the TaskIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClearTaskInstance) GetTaskIdsOk() (*[]string, bool) {
	if o == nil || o.TaskIds == nil {
		return nil, false
	}
	return o.TaskIds, true
}

// HasTaskIds returns a boolean if a field has been set.
func (o *ClearTaskInstance) HasTaskIds() bool {
	if o != nil && o.TaskIds != nil {
		return true
	}

	return false
}

// SetTaskIds gets a reference to the given []string and assigns it to the TaskIds field.
func (o *ClearTaskInstance) SetTaskIds(v []string) {
	o.TaskIds = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ClearTaskInstance) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClearTaskInstance) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ClearTaskInstance) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *ClearTaskInstance) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ClearTaskInstance) GetEndDate() string {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClearTaskInstance) GetEndDateOk() (*string, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ClearTaskInstance) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *ClearTaskInstance) SetEndDate(v string) {
	o.EndDate = &v
}

// GetOnlyFailed returns the OnlyFailed field value if set, zero value otherwise.
func (o *ClearTaskInstance) GetOnlyFailed() bool {
	if o == nil || o.OnlyFailed == nil {
		var ret bool
		return ret
	}
	return *o.OnlyFailed
}

// GetOnlyFailedOk returns a tuple with the OnlyFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClearTaskInstance) GetOnlyFailedOk() (*bool, bool) {
	if o == nil || o.OnlyFailed == nil {
		return nil, false
	}
	return o.OnlyFailed, true
}

// HasOnlyFailed returns a boolean if a field has been set.
func (o *ClearTaskInstance) HasOnlyFailed() bool {
	if o != nil && o.OnlyFailed != nil {
		return true
	}

	return false
}

// SetOnlyFailed gets a reference to the given bool and assigns it to the OnlyFailed field.
func (o *ClearTaskInstance) SetOnlyFailed(v bool) {
	o.OnlyFailed = &v
}

// GetOnlyRunning returns the OnlyRunning field value if set, zero value otherwise.
func (o *ClearTaskInstance) GetOnlyRunning() bool {
	if o == nil || o.OnlyRunning == nil {
		var ret bool
		return ret
	}
	return *o.OnlyRunning
}

// GetOnlyRunningOk returns a tuple with the OnlyRunning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClearTaskInstance) GetOnlyRunningOk() (*bool, bool) {
	if o == nil || o.OnlyRunning == nil {
		return nil, false
	}
	return o.OnlyRunning, true
}

// HasOnlyRunning returns a boolean if a field has been set.
func (o *ClearTaskInstance) HasOnlyRunning() bool {
	if o != nil && o.OnlyRunning != nil {
		return true
	}

	return false
}

// SetOnlyRunning gets a reference to the given bool and assigns it to the OnlyRunning field.
func (o *ClearTaskInstance) SetOnlyRunning(v bool) {
	o.OnlyRunning = &v
}

// GetIncludeSubdags returns the IncludeSubdags field value if set, zero value otherwise.
func (o *ClearTaskInstance) GetIncludeSubdags() bool {
	if o == nil || o.IncludeSubdags == nil {
		var ret bool
		return ret
	}
	return *o.IncludeSubdags
}

// GetIncludeSubdagsOk returns a tuple with the IncludeSubdags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClearTaskInstance) GetIncludeSubdagsOk() (*bool, bool) {
	if o == nil || o.IncludeSubdags == nil {
		return nil, false
	}
	return o.IncludeSubdags, true
}

// HasIncludeSubdags returns a boolean if a field has been set.
func (o *ClearTaskInstance) HasIncludeSubdags() bool {
	if o != nil && o.IncludeSubdags != nil {
		return true
	}

	return false
}

// SetIncludeSubdags gets a reference to the given bool and assigns it to the IncludeSubdags field.
func (o *ClearTaskInstance) SetIncludeSubdags(v bool) {
	o.IncludeSubdags = &v
}

// GetIncludeParentdag returns the IncludeParentdag field value if set, zero value otherwise.
func (o *ClearTaskInstance) GetIncludeParentdag() bool {
	if o == nil || o.IncludeParentdag == nil {
		var ret bool
		return ret
	}
	return *o.IncludeParentdag
}

// GetIncludeParentdagOk returns a tuple with the IncludeParentdag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClearTaskInstance) GetIncludeParentdagOk() (*bool, bool) {
	if o == nil || o.IncludeParentdag == nil {
		return nil, false
	}
	return o.IncludeParentdag, true
}

// HasIncludeParentdag returns a boolean if a field has been set.
func (o *ClearTaskInstance) HasIncludeParentdag() bool {
	if o != nil && o.IncludeParentdag != nil {
		return true
	}

	return false
}

// SetIncludeParentdag gets a reference to the given bool and assigns it to the IncludeParentdag field.
func (o *ClearTaskInstance) SetIncludeParentdag(v bool) {
	o.IncludeParentdag = &v
}

// GetResetDagRuns returns the ResetDagRuns field value if set, zero value otherwise.
func (o *ClearTaskInstance) GetResetDagRuns() bool {
	if o == nil || o.ResetDagRuns == nil {
		var ret bool
		return ret
	}
	return *o.ResetDagRuns
}

// GetResetDagRunsOk returns a tuple with the ResetDagRuns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClearTaskInstance) GetResetDagRunsOk() (*bool, bool) {
	if o == nil || o.ResetDagRuns == nil {
		return nil, false
	}
	return o.ResetDagRuns, true
}

// HasResetDagRuns returns a boolean if a field has been set.
func (o *ClearTaskInstance) HasResetDagRuns() bool {
	if o != nil && o.ResetDagRuns != nil {
		return true
	}

	return false
}

// SetResetDagRuns gets a reference to the given bool and assigns it to the ResetDagRuns field.
func (o *ClearTaskInstance) SetResetDagRuns(v bool) {
	o.ResetDagRuns = &v
}

func (o ClearTaskInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["dry_run"] = o.DryRun
	}
	if o.TaskIds != nil {
		toSerialize["task_ids"] = o.TaskIds
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	if o.OnlyFailed != nil {
		toSerialize["only_failed"] = o.OnlyFailed
	}
	if o.OnlyRunning != nil {
		toSerialize["only_running"] = o.OnlyRunning
	}
	if o.IncludeSubdags != nil {
		toSerialize["include_subdags"] = o.IncludeSubdags
	}
	if o.IncludeParentdag != nil {
		toSerialize["include_parentdag"] = o.IncludeParentdag
	}
	if o.ResetDagRuns != nil {
		toSerialize["reset_dag_runs"] = o.ResetDagRuns
	}
	return json.Marshal(toSerialize)
}

type NullableClearTaskInstance struct {
	value *ClearTaskInstance
	isSet bool
}

func (v NullableClearTaskInstance) Get() *ClearTaskInstance {
	return v.value
}

func (v *NullableClearTaskInstance) Set(val *ClearTaskInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableClearTaskInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableClearTaskInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClearTaskInstance(val *ClearTaskInstance) *NullableClearTaskInstance {
	return &NullableClearTaskInstance{value: val, isSet: true}
}

func (v NullableClearTaskInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClearTaskInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


