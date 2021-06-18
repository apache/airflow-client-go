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

# DAGRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DagRunId** | Pointer to **NullableString** | Run ID.  The value of this field can be set only when creating the object. If you try to modify the field of an existing object, the request fails with an BAD_REQUEST error.  If not provided, a value will be generated based on execution_date.  If the specified dag_run_id is in use, the creation request fails with an ALREADY_EXISTS error.  This together with DAG_ID are a unique key.  | [optional] 
**DagId** | **string** |  | [readonly] 
**ExecutionDate** | Pointer to **time.Time** | The execution date. This is the time when the DAG run should be started according to the DAG definition. The value of this field can be set only when creating the object. If you try to modify the field of an existing object, the request fails with an BAD_REQUEST error. This together with DAG_ID are a unique key.  | [optional] 
**StartDate** | Pointer to **time.Time** | The start time. The time when DAG run was actually created.  | [optional] [readonly] 
**EndDate** | Pointer to **NullableTime** |  | [optional] [readonly] 
**State** | Pointer to [**DagState**](DagState.md) |  | [optional] 
**ExternalTrigger** | Pointer to **bool** |  | [optional] [readonly] [default to true]
**Conf** | Pointer to **map[string]interface{}** | JSON object describing additional configuration parameters.  The value of this field can be set only when creating the object. If you try to modify the field of an existing object, the request fails with an BAD_REQUEST error.  | [optional] 

## Methods

### NewDAGRun

`func NewDAGRun(dagId string, ) *DAGRun`

NewDAGRun instantiates a new DAGRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDAGRunWithDefaults

`func NewDAGRunWithDefaults() *DAGRun`

NewDAGRunWithDefaults instantiates a new DAGRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDagRunId

`func (o *DAGRun) GetDagRunId() string`

GetDagRunId returns the DagRunId field if non-nil, zero value otherwise.

### GetDagRunIdOk

`func (o *DAGRun) GetDagRunIdOk() (*string, bool)`

GetDagRunIdOk returns a tuple with the DagRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagRunId

`func (o *DAGRun) SetDagRunId(v string)`

SetDagRunId sets DagRunId field to given value.

### HasDagRunId

`func (o *DAGRun) HasDagRunId() bool`

HasDagRunId returns a boolean if a field has been set.

### SetDagRunIdNil

`func (o *DAGRun) SetDagRunIdNil(b bool)`

 SetDagRunIdNil sets the value for DagRunId to be an explicit nil

### UnsetDagRunId
`func (o *DAGRun) UnsetDagRunId()`

UnsetDagRunId ensures that no value is present for DagRunId, not even an explicit nil
### GetDagId

`func (o *DAGRun) GetDagId() string`

GetDagId returns the DagId field if non-nil, zero value otherwise.

### GetDagIdOk

`func (o *DAGRun) GetDagIdOk() (*string, bool)`

GetDagIdOk returns a tuple with the DagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagId

`func (o *DAGRun) SetDagId(v string)`

SetDagId sets DagId field to given value.


### GetExecutionDate

`func (o *DAGRun) GetExecutionDate() time.Time`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *DAGRun) GetExecutionDateOk() (*time.Time, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *DAGRun) SetExecutionDate(v time.Time)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *DAGRun) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetStartDate

`func (o *DAGRun) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DAGRun) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DAGRun) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DAGRun) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *DAGRun) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DAGRun) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DAGRun) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DAGRun) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *DAGRun) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *DAGRun) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetState

`func (o *DAGRun) GetState() DagState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DAGRun) GetStateOk() (*DagState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DAGRun) SetState(v DagState)`

SetState sets State field to given value.

### HasState

`func (o *DAGRun) HasState() bool`

HasState returns a boolean if a field has been set.

### GetExternalTrigger

`func (o *DAGRun) GetExternalTrigger() bool`

GetExternalTrigger returns the ExternalTrigger field if non-nil, zero value otherwise.

### GetExternalTriggerOk

`func (o *DAGRun) GetExternalTriggerOk() (*bool, bool)`

GetExternalTriggerOk returns a tuple with the ExternalTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTrigger

`func (o *DAGRun) SetExternalTrigger(v bool)`

SetExternalTrigger sets ExternalTrigger field to given value.

### HasExternalTrigger

`func (o *DAGRun) HasExternalTrigger() bool`

HasExternalTrigger returns a boolean if a field has been set.

### GetConf

`func (o *DAGRun) GetConf() map[string]interface{}`

GetConf returns the Conf field if non-nil, zero value otherwise.

### GetConfOk

`func (o *DAGRun) GetConfOk() (*map[string]interface{}, bool)`

GetConfOk returns a tuple with the Conf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConf

`func (o *DAGRun) SetConf(v map[string]interface{})`

SetConf sets Conf field to given value.

### HasConf

`func (o *DAGRun) HasConf() bool`

HasConf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


