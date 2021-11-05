# TaskInstanceCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskInstances** | Pointer to [**[]TaskInstance**](TaskInstance.md) |  | [optional] 
**TotalEntries** | Pointer to **int32** | Count of objects in the current result set. | [optional] 

## Methods

### NewTaskInstanceCollection

`func NewTaskInstanceCollection() *TaskInstanceCollection`

NewTaskInstanceCollection instantiates a new TaskInstanceCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskInstanceCollectionWithDefaults

`func NewTaskInstanceCollectionWithDefaults() *TaskInstanceCollection`

NewTaskInstanceCollectionWithDefaults instantiates a new TaskInstanceCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskInstances

`func (o *TaskInstanceCollection) GetTaskInstances() []TaskInstance`

GetTaskInstances returns the TaskInstances field if non-nil, zero value otherwise.

### GetTaskInstancesOk

`func (o *TaskInstanceCollection) GetTaskInstancesOk() (*[]TaskInstance, bool)`

GetTaskInstancesOk returns a tuple with the TaskInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskInstances

`func (o *TaskInstanceCollection) SetTaskInstances(v []TaskInstance)`

SetTaskInstances sets TaskInstances field to given value.

### HasTaskInstances

`func (o *TaskInstanceCollection) HasTaskInstances() bool`

HasTaskInstances returns a boolean if a field has been set.

### GetTotalEntries

`func (o *TaskInstanceCollection) GetTotalEntries() int32`

GetTotalEntries returns the TotalEntries field if non-nil, zero value otherwise.

### GetTotalEntriesOk

`func (o *TaskInstanceCollection) GetTotalEntriesOk() (*int32, bool)`

GetTotalEntriesOk returns a tuple with the TotalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntries

`func (o *TaskInstanceCollection) SetTotalEntries(v int32)`

SetTotalEntries sets TotalEntries field to given value.

### HasTotalEntries

`func (o *TaskInstanceCollection) HasTotalEntries() bool`

HasTotalEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


