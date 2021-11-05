# EventLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventLogId** | Pointer to **int32** | The event log ID | [optional] [readonly] 
**When** | Pointer to **time.Time** | The time when these events happened. | [optional] [readonly] 
**DagId** | Pointer to **NullableString** | The DAG ID | [optional] [readonly] 
**TaskId** | Pointer to **NullableString** | The DAG ID | [optional] [readonly] 
**Event** | Pointer to **string** | A key describing the type of event. | [optional] [readonly] 
**ExecutionDate** | Pointer to **NullableTime** | When the event was dispatched for an object having execution_date, the value of this field.  | [optional] [readonly] 
**Owner** | Pointer to **string** | Name of the user who triggered these events a. | [optional] [readonly] 
**Extra** | Pointer to **NullableString** | Other information that was not included in the other fields, e.g. the complete CLI command.  | [optional] [readonly] 

## Methods

### NewEventLog

`func NewEventLog() *EventLog`

NewEventLog instantiates a new EventLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventLogWithDefaults

`func NewEventLogWithDefaults() *EventLog`

NewEventLogWithDefaults instantiates a new EventLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventLogId

`func (o *EventLog) GetEventLogId() int32`

GetEventLogId returns the EventLogId field if non-nil, zero value otherwise.

### GetEventLogIdOk

`func (o *EventLog) GetEventLogIdOk() (*int32, bool)`

GetEventLogIdOk returns a tuple with the EventLogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLogId

`func (o *EventLog) SetEventLogId(v int32)`

SetEventLogId sets EventLogId field to given value.

### HasEventLogId

`func (o *EventLog) HasEventLogId() bool`

HasEventLogId returns a boolean if a field has been set.

### GetWhen

`func (o *EventLog) GetWhen() time.Time`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *EventLog) GetWhenOk() (*time.Time, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *EventLog) SetWhen(v time.Time)`

SetWhen sets When field to given value.

### HasWhen

`func (o *EventLog) HasWhen() bool`

HasWhen returns a boolean if a field has been set.

### GetDagId

`func (o *EventLog) GetDagId() string`

GetDagId returns the DagId field if non-nil, zero value otherwise.

### GetDagIdOk

`func (o *EventLog) GetDagIdOk() (*string, bool)`

GetDagIdOk returns a tuple with the DagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagId

`func (o *EventLog) SetDagId(v string)`

SetDagId sets DagId field to given value.

### HasDagId

`func (o *EventLog) HasDagId() bool`

HasDagId returns a boolean if a field has been set.

### SetDagIdNil

`func (o *EventLog) SetDagIdNil(b bool)`

 SetDagIdNil sets the value for DagId to be an explicit nil

### UnsetDagId
`func (o *EventLog) UnsetDagId()`

UnsetDagId ensures that no value is present for DagId, not even an explicit nil
### GetTaskId

`func (o *EventLog) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *EventLog) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *EventLog) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *EventLog) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *EventLog) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *EventLog) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetEvent

`func (o *EventLog) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventLog) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventLog) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *EventLog) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetExecutionDate

`func (o *EventLog) GetExecutionDate() time.Time`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *EventLog) GetExecutionDateOk() (*time.Time, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *EventLog) SetExecutionDate(v time.Time)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *EventLog) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### SetExecutionDateNil

`func (o *EventLog) SetExecutionDateNil(b bool)`

 SetExecutionDateNil sets the value for ExecutionDate to be an explicit nil

### UnsetExecutionDate
`func (o *EventLog) UnsetExecutionDate()`

UnsetExecutionDate ensures that no value is present for ExecutionDate, not even an explicit nil
### GetOwner

`func (o *EventLog) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *EventLog) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *EventLog) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *EventLog) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetExtra

`func (o *EventLog) GetExtra() string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *EventLog) GetExtraOk() (*string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *EventLog) SetExtra(v string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *EventLog) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### SetExtraNil

`func (o *EventLog) SetExtraNil(b bool)`

 SetExtraNil sets the value for Extra to be an explicit nil

### UnsetExtra
`func (o *EventLog) UnsetExtra()`

UnsetExtra ensures that no value is present for Extra, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


