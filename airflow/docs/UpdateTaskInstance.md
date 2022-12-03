# UpdateTaskInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If set, don&#39;t actually run this operation. The response will contain the task instance planned to be affected, but won&#39;t be modified in any way.  | [optional] [default to false]
**NewState** | Pointer to **string** | Expected new state. | [optional] 

## Methods

### NewUpdateTaskInstance

`func NewUpdateTaskInstance() *UpdateTaskInstance`

NewUpdateTaskInstance instantiates a new UpdateTaskInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTaskInstanceWithDefaults

`func NewUpdateTaskInstanceWithDefaults() *UpdateTaskInstance`

NewUpdateTaskInstanceWithDefaults instantiates a new UpdateTaskInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *UpdateTaskInstance) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateTaskInstance) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateTaskInstance) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateTaskInstance) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetNewState

`func (o *UpdateTaskInstance) GetNewState() string`

GetNewState returns the NewState field if non-nil, zero value otherwise.

### GetNewStateOk

`func (o *UpdateTaskInstance) GetNewStateOk() (*string, bool)`

GetNewStateOk returns a tuple with the NewState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewState

`func (o *UpdateTaskInstance) SetNewState(v string)`

SetNewState sets NewState field to given value.

### HasNewState

`func (o *UpdateTaskInstance) HasNewState() bool`

HasNewState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


