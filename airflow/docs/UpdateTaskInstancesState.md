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

# UpdateTaskInstancesState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If set, don&#39;t actually run this operation. The response will contain a list of task instances planned to be affected, but won&#39;t be modified in any way.  | [optional] [default to true]
**TaskId** | Pointer to **string** | The task ID. | [optional] 
**ExecutionDate** | Pointer to **string** | The execution date. Either set this or dag_run_id but not both. | [optional] 
**DagRunId** | Pointer to **string** | The task instance&#39;s DAG run ID. Either set this or execution_date but not both.  *New in version 2.3.0*  | [optional] 
**IncludeUpstream** | Pointer to **bool** | If set to true, upstream tasks are also affected. | [optional] 
**IncludeDownstream** | Pointer to **bool** | If set to true, downstream tasks are also affected. | [optional] 
**IncludeFuture** | Pointer to **bool** | If set to True, also tasks from future DAG Runs are affected. | [optional] 
**IncludePast** | Pointer to **bool** | If set to True, also tasks from past DAG Runs are affected. | [optional] 
**NewState** | Pointer to **string** | Expected new state. | [optional] 

## Methods

### NewUpdateTaskInstancesState

`func NewUpdateTaskInstancesState() *UpdateTaskInstancesState`

NewUpdateTaskInstancesState instantiates a new UpdateTaskInstancesState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTaskInstancesStateWithDefaults

`func NewUpdateTaskInstancesStateWithDefaults() *UpdateTaskInstancesState`

NewUpdateTaskInstancesStateWithDefaults instantiates a new UpdateTaskInstancesState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *UpdateTaskInstancesState) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateTaskInstancesState) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateTaskInstancesState) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateTaskInstancesState) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetTaskId

`func (o *UpdateTaskInstancesState) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *UpdateTaskInstancesState) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *UpdateTaskInstancesState) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *UpdateTaskInstancesState) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetExecutionDate

`func (o *UpdateTaskInstancesState) GetExecutionDate() string`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *UpdateTaskInstancesState) GetExecutionDateOk() (*string, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *UpdateTaskInstancesState) SetExecutionDate(v string)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *UpdateTaskInstancesState) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetDagRunId

`func (o *UpdateTaskInstancesState) GetDagRunId() string`

GetDagRunId returns the DagRunId field if non-nil, zero value otherwise.

### GetDagRunIdOk

`func (o *UpdateTaskInstancesState) GetDagRunIdOk() (*string, bool)`

GetDagRunIdOk returns a tuple with the DagRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagRunId

`func (o *UpdateTaskInstancesState) SetDagRunId(v string)`

SetDagRunId sets DagRunId field to given value.

### HasDagRunId

`func (o *UpdateTaskInstancesState) HasDagRunId() bool`

HasDagRunId returns a boolean if a field has been set.

### GetIncludeUpstream

`func (o *UpdateTaskInstancesState) GetIncludeUpstream() bool`

GetIncludeUpstream returns the IncludeUpstream field if non-nil, zero value otherwise.

### GetIncludeUpstreamOk

`func (o *UpdateTaskInstancesState) GetIncludeUpstreamOk() (*bool, bool)`

GetIncludeUpstreamOk returns a tuple with the IncludeUpstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeUpstream

`func (o *UpdateTaskInstancesState) SetIncludeUpstream(v bool)`

SetIncludeUpstream sets IncludeUpstream field to given value.

### HasIncludeUpstream

`func (o *UpdateTaskInstancesState) HasIncludeUpstream() bool`

HasIncludeUpstream returns a boolean if a field has been set.

### GetIncludeDownstream

`func (o *UpdateTaskInstancesState) GetIncludeDownstream() bool`

GetIncludeDownstream returns the IncludeDownstream field if non-nil, zero value otherwise.

### GetIncludeDownstreamOk

`func (o *UpdateTaskInstancesState) GetIncludeDownstreamOk() (*bool, bool)`

GetIncludeDownstreamOk returns a tuple with the IncludeDownstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDownstream

`func (o *UpdateTaskInstancesState) SetIncludeDownstream(v bool)`

SetIncludeDownstream sets IncludeDownstream field to given value.

### HasIncludeDownstream

`func (o *UpdateTaskInstancesState) HasIncludeDownstream() bool`

HasIncludeDownstream returns a boolean if a field has been set.

### GetIncludeFuture

`func (o *UpdateTaskInstancesState) GetIncludeFuture() bool`

GetIncludeFuture returns the IncludeFuture field if non-nil, zero value otherwise.

### GetIncludeFutureOk

`func (o *UpdateTaskInstancesState) GetIncludeFutureOk() (*bool, bool)`

GetIncludeFutureOk returns a tuple with the IncludeFuture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFuture

`func (o *UpdateTaskInstancesState) SetIncludeFuture(v bool)`

SetIncludeFuture sets IncludeFuture field to given value.

### HasIncludeFuture

`func (o *UpdateTaskInstancesState) HasIncludeFuture() bool`

HasIncludeFuture returns a boolean if a field has been set.

### GetIncludePast

`func (o *UpdateTaskInstancesState) GetIncludePast() bool`

GetIncludePast returns the IncludePast field if non-nil, zero value otherwise.

### GetIncludePastOk

`func (o *UpdateTaskInstancesState) GetIncludePastOk() (*bool, bool)`

GetIncludePastOk returns a tuple with the IncludePast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePast

`func (o *UpdateTaskInstancesState) SetIncludePast(v bool)`

SetIncludePast sets IncludePast field to given value.

### HasIncludePast

`func (o *UpdateTaskInstancesState) HasIncludePast() bool`

HasIncludePast returns a boolean if a field has been set.

### GetNewState

`func (o *UpdateTaskInstancesState) GetNewState() string`

GetNewState returns the NewState field if non-nil, zero value otherwise.

### GetNewStateOk

`func (o *UpdateTaskInstancesState) GetNewStateOk() (*string, bool)`

GetNewStateOk returns a tuple with the NewState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewState

`func (o *UpdateTaskInstancesState) SetNewState(v string)`

SetNewState sets NewState field to given value.

### HasNewState

`func (o *UpdateTaskInstancesState) HasNewState() bool`

HasNewState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


