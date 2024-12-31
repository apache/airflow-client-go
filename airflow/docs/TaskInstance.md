<!--
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
 -->

# TaskInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | Pointer to **string** |  | [optional] 
**DagId** | Pointer to **string** |  | [optional] 
**DagRunId** | Pointer to **string** | The DagRun ID for this task instance  *New in version 2.3.0*  | [optional] 
**ExecutionDate** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **NullableString** |  | [optional] 
**EndDate** | Pointer to **NullableString** |  | [optional] 
**Duration** | Pointer to **NullableFloat32** |  | [optional] 
**State** | Pointer to [**NullableTaskState**](TaskState.md) |  | [optional] 
**TryNumber** | Pointer to **int32** |  | [optional] 
**MapIndex** | Pointer to **int32** |  | [optional] 
**MaxTries** | Pointer to **int32** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Unixname** | Pointer to **string** |  | [optional] 
**Pool** | Pointer to **string** |  | [optional] 
**PoolSlots** | Pointer to **int32** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**PriorityWeight** | Pointer to **NullableInt32** |  | [optional] 
**Operator** | Pointer to **NullableString** | *Changed in version 2.1.1*&amp;#58; Field becomes nullable.  | [optional] 
**QueuedWhen** | Pointer to **NullableString** |  | [optional] 
**Pid** | Pointer to **NullableInt32** |  | [optional] 
**ExecutorConfig** | Pointer to **string** |  | [optional] 
**SlaMiss** | Pointer to [**NullableSLAMiss**](SLAMiss.md) |  | [optional] 
**RenderedFields** | Pointer to **map[string]interface{}** | JSON object describing rendered fields.  *New in version 2.3.0*  | [optional] 
**Trigger** | Pointer to [**NullableTrigger**](Trigger.md) |  | [optional] 
**TriggererJob** | Pointer to [**NullableJob**](Job.md) |  | [optional] 
**Note** | Pointer to **NullableString** | Contains manually entered notes by the user about the TaskInstance.  *New in version 2.5.0*  | [optional] 

## Methods

### NewTaskInstance

`func NewTaskInstance() *TaskInstance`

NewTaskInstance instantiates a new TaskInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskInstanceWithDefaults

`func NewTaskInstanceWithDefaults() *TaskInstance`

NewTaskInstanceWithDefaults instantiates a new TaskInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *TaskInstance) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TaskInstance) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TaskInstance) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *TaskInstance) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetDagId

`func (o *TaskInstance) GetDagId() string`

GetDagId returns the DagId field if non-nil, zero value otherwise.

### GetDagIdOk

`func (o *TaskInstance) GetDagIdOk() (*string, bool)`

GetDagIdOk returns a tuple with the DagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagId

`func (o *TaskInstance) SetDagId(v string)`

SetDagId sets DagId field to given value.

### HasDagId

`func (o *TaskInstance) HasDagId() bool`

HasDagId returns a boolean if a field has been set.

### GetDagRunId

`func (o *TaskInstance) GetDagRunId() string`

GetDagRunId returns the DagRunId field if non-nil, zero value otherwise.

### GetDagRunIdOk

`func (o *TaskInstance) GetDagRunIdOk() (*string, bool)`

GetDagRunIdOk returns a tuple with the DagRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagRunId

`func (o *TaskInstance) SetDagRunId(v string)`

SetDagRunId sets DagRunId field to given value.

### HasDagRunId

`func (o *TaskInstance) HasDagRunId() bool`

HasDagRunId returns a boolean if a field has been set.

### GetExecutionDate

`func (o *TaskInstance) GetExecutionDate() string`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *TaskInstance) GetExecutionDateOk() (*string, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *TaskInstance) SetExecutionDate(v string)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *TaskInstance) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetStartDate

`func (o *TaskInstance) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TaskInstance) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TaskInstance) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TaskInstance) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *TaskInstance) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *TaskInstance) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *TaskInstance) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TaskInstance) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TaskInstance) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *TaskInstance) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *TaskInstance) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *TaskInstance) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetDuration

`func (o *TaskInstance) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TaskInstance) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TaskInstance) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TaskInstance) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *TaskInstance) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *TaskInstance) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetState

`func (o *TaskInstance) GetState() TaskState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TaskInstance) GetStateOk() (*TaskState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TaskInstance) SetState(v TaskState)`

SetState sets State field to given value.

### HasState

`func (o *TaskInstance) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *TaskInstance) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *TaskInstance) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetTryNumber

`func (o *TaskInstance) GetTryNumber() int32`

GetTryNumber returns the TryNumber field if non-nil, zero value otherwise.

### GetTryNumberOk

`func (o *TaskInstance) GetTryNumberOk() (*int32, bool)`

GetTryNumberOk returns a tuple with the TryNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryNumber

`func (o *TaskInstance) SetTryNumber(v int32)`

SetTryNumber sets TryNumber field to given value.

### HasTryNumber

`func (o *TaskInstance) HasTryNumber() bool`

HasTryNumber returns a boolean if a field has been set.

### GetMapIndex

`func (o *TaskInstance) GetMapIndex() int32`

GetMapIndex returns the MapIndex field if non-nil, zero value otherwise.

### GetMapIndexOk

`func (o *TaskInstance) GetMapIndexOk() (*int32, bool)`

GetMapIndexOk returns a tuple with the MapIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapIndex

`func (o *TaskInstance) SetMapIndex(v int32)`

SetMapIndex sets MapIndex field to given value.

### HasMapIndex

`func (o *TaskInstance) HasMapIndex() bool`

HasMapIndex returns a boolean if a field has been set.

### GetMaxTries

`func (o *TaskInstance) GetMaxTries() int32`

GetMaxTries returns the MaxTries field if non-nil, zero value otherwise.

### GetMaxTriesOk

`func (o *TaskInstance) GetMaxTriesOk() (*int32, bool)`

GetMaxTriesOk returns a tuple with the MaxTries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTries

`func (o *TaskInstance) SetMaxTries(v int32)`

SetMaxTries sets MaxTries field to given value.

### HasMaxTries

`func (o *TaskInstance) HasMaxTries() bool`

HasMaxTries returns a boolean if a field has been set.

### GetHostname

`func (o *TaskInstance) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *TaskInstance) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *TaskInstance) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *TaskInstance) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetUnixname

`func (o *TaskInstance) GetUnixname() string`

GetUnixname returns the Unixname field if non-nil, zero value otherwise.

### GetUnixnameOk

`func (o *TaskInstance) GetUnixnameOk() (*string, bool)`

GetUnixnameOk returns a tuple with the Unixname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixname

`func (o *TaskInstance) SetUnixname(v string)`

SetUnixname sets Unixname field to given value.

### HasUnixname

`func (o *TaskInstance) HasUnixname() bool`

HasUnixname returns a boolean if a field has been set.

### GetPool

`func (o *TaskInstance) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *TaskInstance) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *TaskInstance) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *TaskInstance) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolSlots

`func (o *TaskInstance) GetPoolSlots() int32`

GetPoolSlots returns the PoolSlots field if non-nil, zero value otherwise.

### GetPoolSlotsOk

`func (o *TaskInstance) GetPoolSlotsOk() (*int32, bool)`

GetPoolSlotsOk returns a tuple with the PoolSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSlots

`func (o *TaskInstance) SetPoolSlots(v int32)`

SetPoolSlots sets PoolSlots field to given value.

### HasPoolSlots

`func (o *TaskInstance) HasPoolSlots() bool`

HasPoolSlots returns a boolean if a field has been set.

### GetQueue

`func (o *TaskInstance) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *TaskInstance) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *TaskInstance) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *TaskInstance) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *TaskInstance) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *TaskInstance) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetPriorityWeight

`func (o *TaskInstance) GetPriorityWeight() int32`

GetPriorityWeight returns the PriorityWeight field if non-nil, zero value otherwise.

### GetPriorityWeightOk

`func (o *TaskInstance) GetPriorityWeightOk() (*int32, bool)`

GetPriorityWeightOk returns a tuple with the PriorityWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityWeight

`func (o *TaskInstance) SetPriorityWeight(v int32)`

SetPriorityWeight sets PriorityWeight field to given value.

### HasPriorityWeight

`func (o *TaskInstance) HasPriorityWeight() bool`

HasPriorityWeight returns a boolean if a field has been set.

### SetPriorityWeightNil

`func (o *TaskInstance) SetPriorityWeightNil(b bool)`

 SetPriorityWeightNil sets the value for PriorityWeight to be an explicit nil

### UnsetPriorityWeight
`func (o *TaskInstance) UnsetPriorityWeight()`

UnsetPriorityWeight ensures that no value is present for PriorityWeight, not even an explicit nil
### GetOperator

`func (o *TaskInstance) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *TaskInstance) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *TaskInstance) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *TaskInstance) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *TaskInstance) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *TaskInstance) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetQueuedWhen

`func (o *TaskInstance) GetQueuedWhen() string`

GetQueuedWhen returns the QueuedWhen field if non-nil, zero value otherwise.

### GetQueuedWhenOk

`func (o *TaskInstance) GetQueuedWhenOk() (*string, bool)`

GetQueuedWhenOk returns a tuple with the QueuedWhen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedWhen

`func (o *TaskInstance) SetQueuedWhen(v string)`

SetQueuedWhen sets QueuedWhen field to given value.

### HasQueuedWhen

`func (o *TaskInstance) HasQueuedWhen() bool`

HasQueuedWhen returns a boolean if a field has been set.

### SetQueuedWhenNil

`func (o *TaskInstance) SetQueuedWhenNil(b bool)`

 SetQueuedWhenNil sets the value for QueuedWhen to be an explicit nil

### UnsetQueuedWhen
`func (o *TaskInstance) UnsetQueuedWhen()`

UnsetQueuedWhen ensures that no value is present for QueuedWhen, not even an explicit nil
### GetPid

`func (o *TaskInstance) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *TaskInstance) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *TaskInstance) SetPid(v int32)`

SetPid sets Pid field to given value.

### HasPid

`func (o *TaskInstance) HasPid() bool`

HasPid returns a boolean if a field has been set.

### SetPidNil

`func (o *TaskInstance) SetPidNil(b bool)`

 SetPidNil sets the value for Pid to be an explicit nil

### UnsetPid
`func (o *TaskInstance) UnsetPid()`

UnsetPid ensures that no value is present for Pid, not even an explicit nil
### GetExecutorConfig

`func (o *TaskInstance) GetExecutorConfig() string`

GetExecutorConfig returns the ExecutorConfig field if non-nil, zero value otherwise.

### GetExecutorConfigOk

`func (o *TaskInstance) GetExecutorConfigOk() (*string, bool)`

GetExecutorConfigOk returns a tuple with the ExecutorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutorConfig

`func (o *TaskInstance) SetExecutorConfig(v string)`

SetExecutorConfig sets ExecutorConfig field to given value.

### HasExecutorConfig

`func (o *TaskInstance) HasExecutorConfig() bool`

HasExecutorConfig returns a boolean if a field has been set.

### GetSlaMiss

`func (o *TaskInstance) GetSlaMiss() SLAMiss`

GetSlaMiss returns the SlaMiss field if non-nil, zero value otherwise.

### GetSlaMissOk

`func (o *TaskInstance) GetSlaMissOk() (*SLAMiss, bool)`

GetSlaMissOk returns a tuple with the SlaMiss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaMiss

`func (o *TaskInstance) SetSlaMiss(v SLAMiss)`

SetSlaMiss sets SlaMiss field to given value.

### HasSlaMiss

`func (o *TaskInstance) HasSlaMiss() bool`

HasSlaMiss returns a boolean if a field has been set.

### SetSlaMissNil

`func (o *TaskInstance) SetSlaMissNil(b bool)`

 SetSlaMissNil sets the value for SlaMiss to be an explicit nil

### UnsetSlaMiss
`func (o *TaskInstance) UnsetSlaMiss()`

UnsetSlaMiss ensures that no value is present for SlaMiss, not even an explicit nil
### GetRenderedFields

`func (o *TaskInstance) GetRenderedFields() map[string]interface{}`

GetRenderedFields returns the RenderedFields field if non-nil, zero value otherwise.

### GetRenderedFieldsOk

`func (o *TaskInstance) GetRenderedFieldsOk() (*map[string]interface{}, bool)`

GetRenderedFieldsOk returns a tuple with the RenderedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedFields

`func (o *TaskInstance) SetRenderedFields(v map[string]interface{})`

SetRenderedFields sets RenderedFields field to given value.

### HasRenderedFields

`func (o *TaskInstance) HasRenderedFields() bool`

HasRenderedFields returns a boolean if a field has been set.

### GetTrigger

`func (o *TaskInstance) GetTrigger() Trigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *TaskInstance) GetTriggerOk() (*Trigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *TaskInstance) SetTrigger(v Trigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *TaskInstance) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### SetTriggerNil

`func (o *TaskInstance) SetTriggerNil(b bool)`

 SetTriggerNil sets the value for Trigger to be an explicit nil

### UnsetTrigger
`func (o *TaskInstance) UnsetTrigger()`

UnsetTrigger ensures that no value is present for Trigger, not even an explicit nil
### GetTriggererJob

`func (o *TaskInstance) GetTriggererJob() Job`

GetTriggererJob returns the TriggererJob field if non-nil, zero value otherwise.

### GetTriggererJobOk

`func (o *TaskInstance) GetTriggererJobOk() (*Job, bool)`

GetTriggererJobOk returns a tuple with the TriggererJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggererJob

`func (o *TaskInstance) SetTriggererJob(v Job)`

SetTriggererJob sets TriggererJob field to given value.

### HasTriggererJob

`func (o *TaskInstance) HasTriggererJob() bool`

HasTriggererJob returns a boolean if a field has been set.

### SetTriggererJobNil

`func (o *TaskInstance) SetTriggererJobNil(b bool)`

 SetTriggererJobNil sets the value for TriggererJob to be an explicit nil

### UnsetTriggererJob
`func (o *TaskInstance) UnsetTriggererJob()`

UnsetTriggererJob ensures that no value is present for TriggererJob, not even an explicit nil
### GetNote

`func (o *TaskInstance) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *TaskInstance) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *TaskInstance) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *TaskInstance) HasNote() bool`

HasNote returns a boolean if a field has been set.

### SetNoteNil

`func (o *TaskInstance) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *TaskInstance) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


