# TaskInstanceReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | Pointer to **string** | The task ID. | [optional] [readonly] 
**DagId** | Pointer to **string** | The DAG ID. | [optional] [readonly] 
**ExecutionDate** | Pointer to **string** |  | [optional] [readonly] 
**DagRunId** | Pointer to **string** | The DAG run ID. | [optional] [readonly] 

## Methods

### NewTaskInstanceReference

`func NewTaskInstanceReference() *TaskInstanceReference`

NewTaskInstanceReference instantiates a new TaskInstanceReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskInstanceReferenceWithDefaults

`func NewTaskInstanceReferenceWithDefaults() *TaskInstanceReference`

NewTaskInstanceReferenceWithDefaults instantiates a new TaskInstanceReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *TaskInstanceReference) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TaskInstanceReference) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TaskInstanceReference) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *TaskInstanceReference) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetDagId

`func (o *TaskInstanceReference) GetDagId() string`

GetDagId returns the DagId field if non-nil, zero value otherwise.

### GetDagIdOk

`func (o *TaskInstanceReference) GetDagIdOk() (*string, bool)`

GetDagIdOk returns a tuple with the DagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagId

`func (o *TaskInstanceReference) SetDagId(v string)`

SetDagId sets DagId field to given value.

### HasDagId

`func (o *TaskInstanceReference) HasDagId() bool`

HasDagId returns a boolean if a field has been set.

### GetExecutionDate

`func (o *TaskInstanceReference) GetExecutionDate() string`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *TaskInstanceReference) GetExecutionDateOk() (*string, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *TaskInstanceReference) SetExecutionDate(v string)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *TaskInstanceReference) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetDagRunId

`func (o *TaskInstanceReference) GetDagRunId() string`

GetDagRunId returns the DagRunId field if non-nil, zero value otherwise.

### GetDagRunIdOk

`func (o *TaskInstanceReference) GetDagRunIdOk() (*string, bool)`

GetDagRunIdOk returns a tuple with the DagRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagRunId

`func (o *TaskInstanceReference) SetDagRunId(v string)`

SetDagRunId sets DagRunId field to given value.

### HasDagRunId

`func (o *TaskInstanceReference) HasDagRunId() bool`

HasDagRunId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


