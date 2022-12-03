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

# ClearTaskInstances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If set, don&#39;t actually run this operation. The response will contain a list of task instances planned to be cleaned, but not modified in any way.  | [optional] [default to true]
**TaskIds** | Pointer to **[]string** | A list of task ids to clear.  *New in version 2.1.0*  | [optional] 
**StartDate** | Pointer to **string** | The minimum execution date to clear. | [optional] 
**EndDate** | Pointer to **string** | The maximum execution date to clear. | [optional] 
**OnlyFailed** | Pointer to **bool** | Only clear failed tasks. | [optional] [default to true]
**OnlyRunning** | Pointer to **bool** | Only clear running tasks. | [optional] [default to false]
**IncludeSubdags** | Pointer to **bool** | Clear tasks in subdags and clear external tasks indicated by ExternalTaskMarker. | [optional] 
**IncludeParentdag** | Pointer to **bool** | Clear tasks in the parent dag of the subdag. | [optional] 
**ResetDagRuns** | Pointer to **bool** | Set state of DAG runs to RUNNING. | [optional] 
**DagRunId** | Pointer to **NullableString** | The DagRun ID for this task instance | [optional] 
**IncludeUpstream** | Pointer to **bool** | If set to true, upstream tasks are also affected. | [optional] [default to false]
**IncludeDownstream** | Pointer to **bool** | If set to true, downstream tasks are also affected. | [optional] [default to false]
**IncludeFuture** | Pointer to **bool** | If set to True, also tasks from future DAG Runs are affected. | [optional] [default to false]
**IncludePast** | Pointer to **bool** | If set to True, also tasks from past DAG Runs are affected. | [optional] [default to false]

## Methods

### NewClearTaskInstances

`func NewClearTaskInstances() *ClearTaskInstances`

NewClearTaskInstances instantiates a new ClearTaskInstances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearTaskInstancesWithDefaults

`func NewClearTaskInstancesWithDefaults() *ClearTaskInstances`

NewClearTaskInstancesWithDefaults instantiates a new ClearTaskInstances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ClearTaskInstances) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ClearTaskInstances) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ClearTaskInstances) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ClearTaskInstances) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetTaskIds

`func (o *ClearTaskInstances) GetTaskIds() []string`

GetTaskIds returns the TaskIds field if non-nil, zero value otherwise.

### GetTaskIdsOk

`func (o *ClearTaskInstances) GetTaskIdsOk() (*[]string, bool)`

GetTaskIdsOk returns a tuple with the TaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIds

`func (o *ClearTaskInstances) SetTaskIds(v []string)`

SetTaskIds sets TaskIds field to given value.

### HasTaskIds

`func (o *ClearTaskInstances) HasTaskIds() bool`

HasTaskIds returns a boolean if a field has been set.

### GetStartDate

`func (o *ClearTaskInstances) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ClearTaskInstances) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ClearTaskInstances) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ClearTaskInstances) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ClearTaskInstances) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ClearTaskInstances) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ClearTaskInstances) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ClearTaskInstances) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetOnlyFailed

`func (o *ClearTaskInstances) GetOnlyFailed() bool`

GetOnlyFailed returns the OnlyFailed field if non-nil, zero value otherwise.

### GetOnlyFailedOk

`func (o *ClearTaskInstances) GetOnlyFailedOk() (*bool, bool)`

GetOnlyFailedOk returns a tuple with the OnlyFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyFailed

`func (o *ClearTaskInstances) SetOnlyFailed(v bool)`

SetOnlyFailed sets OnlyFailed field to given value.

### HasOnlyFailed

`func (o *ClearTaskInstances) HasOnlyFailed() bool`

HasOnlyFailed returns a boolean if a field has been set.

### GetOnlyRunning

`func (o *ClearTaskInstances) GetOnlyRunning() bool`

GetOnlyRunning returns the OnlyRunning field if non-nil, zero value otherwise.

### GetOnlyRunningOk

`func (o *ClearTaskInstances) GetOnlyRunningOk() (*bool, bool)`

GetOnlyRunningOk returns a tuple with the OnlyRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyRunning

`func (o *ClearTaskInstances) SetOnlyRunning(v bool)`

SetOnlyRunning sets OnlyRunning field to given value.

### HasOnlyRunning

`func (o *ClearTaskInstances) HasOnlyRunning() bool`

HasOnlyRunning returns a boolean if a field has been set.

### GetIncludeSubdags

`func (o *ClearTaskInstances) GetIncludeSubdags() bool`

GetIncludeSubdags returns the IncludeSubdags field if non-nil, zero value otherwise.

### GetIncludeSubdagsOk

`func (o *ClearTaskInstances) GetIncludeSubdagsOk() (*bool, bool)`

GetIncludeSubdagsOk returns a tuple with the IncludeSubdags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubdags

`func (o *ClearTaskInstances) SetIncludeSubdags(v bool)`

SetIncludeSubdags sets IncludeSubdags field to given value.

### HasIncludeSubdags

`func (o *ClearTaskInstances) HasIncludeSubdags() bool`

HasIncludeSubdags returns a boolean if a field has been set.

### GetIncludeParentdag

`func (o *ClearTaskInstances) GetIncludeParentdag() bool`

GetIncludeParentdag returns the IncludeParentdag field if non-nil, zero value otherwise.

### GetIncludeParentdagOk

`func (o *ClearTaskInstances) GetIncludeParentdagOk() (*bool, bool)`

GetIncludeParentdagOk returns a tuple with the IncludeParentdag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeParentdag

`func (o *ClearTaskInstances) SetIncludeParentdag(v bool)`

SetIncludeParentdag sets IncludeParentdag field to given value.

### HasIncludeParentdag

`func (o *ClearTaskInstances) HasIncludeParentdag() bool`

HasIncludeParentdag returns a boolean if a field has been set.

### GetResetDagRuns

`func (o *ClearTaskInstances) GetResetDagRuns() bool`

GetResetDagRuns returns the ResetDagRuns field if non-nil, zero value otherwise.

### GetResetDagRunsOk

`func (o *ClearTaskInstances) GetResetDagRunsOk() (*bool, bool)`

GetResetDagRunsOk returns a tuple with the ResetDagRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetDagRuns

`func (o *ClearTaskInstances) SetResetDagRuns(v bool)`

SetResetDagRuns sets ResetDagRuns field to given value.

### HasResetDagRuns

`func (o *ClearTaskInstances) HasResetDagRuns() bool`

HasResetDagRuns returns a boolean if a field has been set.

### GetDagRunId

`func (o *ClearTaskInstances) GetDagRunId() string`

GetDagRunId returns the DagRunId field if non-nil, zero value otherwise.

### GetDagRunIdOk

`func (o *ClearTaskInstances) GetDagRunIdOk() (*string, bool)`

GetDagRunIdOk returns a tuple with the DagRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagRunId

`func (o *ClearTaskInstances) SetDagRunId(v string)`

SetDagRunId sets DagRunId field to given value.

### HasDagRunId

`func (o *ClearTaskInstances) HasDagRunId() bool`

HasDagRunId returns a boolean if a field has been set.

### SetDagRunIdNil

`func (o *ClearTaskInstances) SetDagRunIdNil(b bool)`

 SetDagRunIdNil sets the value for DagRunId to be an explicit nil

### UnsetDagRunId
`func (o *ClearTaskInstances) UnsetDagRunId()`

UnsetDagRunId ensures that no value is present for DagRunId, not even an explicit nil
### GetIncludeUpstream

`func (o *ClearTaskInstances) GetIncludeUpstream() bool`

GetIncludeUpstream returns the IncludeUpstream field if non-nil, zero value otherwise.

### GetIncludeUpstreamOk

`func (o *ClearTaskInstances) GetIncludeUpstreamOk() (*bool, bool)`

GetIncludeUpstreamOk returns a tuple with the IncludeUpstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeUpstream

`func (o *ClearTaskInstances) SetIncludeUpstream(v bool)`

SetIncludeUpstream sets IncludeUpstream field to given value.

### HasIncludeUpstream

`func (o *ClearTaskInstances) HasIncludeUpstream() bool`

HasIncludeUpstream returns a boolean if a field has been set.

### GetIncludeDownstream

`func (o *ClearTaskInstances) GetIncludeDownstream() bool`

GetIncludeDownstream returns the IncludeDownstream field if non-nil, zero value otherwise.

### GetIncludeDownstreamOk

`func (o *ClearTaskInstances) GetIncludeDownstreamOk() (*bool, bool)`

GetIncludeDownstreamOk returns a tuple with the IncludeDownstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDownstream

`func (o *ClearTaskInstances) SetIncludeDownstream(v bool)`

SetIncludeDownstream sets IncludeDownstream field to given value.

### HasIncludeDownstream

`func (o *ClearTaskInstances) HasIncludeDownstream() bool`

HasIncludeDownstream returns a boolean if a field has been set.

### GetIncludeFuture

`func (o *ClearTaskInstances) GetIncludeFuture() bool`

GetIncludeFuture returns the IncludeFuture field if non-nil, zero value otherwise.

### GetIncludeFutureOk

`func (o *ClearTaskInstances) GetIncludeFutureOk() (*bool, bool)`

GetIncludeFutureOk returns a tuple with the IncludeFuture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFuture

`func (o *ClearTaskInstances) SetIncludeFuture(v bool)`

SetIncludeFuture sets IncludeFuture field to given value.

### HasIncludeFuture

`func (o *ClearTaskInstances) HasIncludeFuture() bool`

HasIncludeFuture returns a boolean if a field has been set.

### GetIncludePast

`func (o *ClearTaskInstances) GetIncludePast() bool`

GetIncludePast returns the IncludePast field if non-nil, zero value otherwise.

### GetIncludePastOk

`func (o *ClearTaskInstances) GetIncludePastOk() (*bool, bool)`

GetIncludePastOk returns a tuple with the IncludePast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePast

`func (o *ClearTaskInstances) SetIncludePast(v bool)`

SetIncludePast sets IncludePast field to given value.

### HasIncludePast

`func (o *ClearTaskInstances) HasIncludePast() bool`

HasIncludePast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


