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

# ListTaskInstanceForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DagIds** | Pointer to **[]string** | Return objects with specific DAG IDs. The value can be repeated to retrieve multiple matching values (OR condition). | [optional] 
**DagRunIds** | Pointer to **[]string** | Return objects with specific DAG Run IDs. The value can be repeated to retrieve multiple matching values (OR condition). *New in version 2.7.1* | [optional] 
**TaskIds** | Pointer to **[]string** | Return objects with specific task IDs. The value can be repeated to retrieve multiple matching values (OR condition). *New in version 2.7.1* | [optional] 
**ExecutionDateGte** | Pointer to **time.Time** | Returns objects greater or equal to the specified date.  This can be combined with execution_date_lte parameter to receive only the selected period.  | [optional] 
**ExecutionDateLte** | Pointer to **time.Time** | Returns objects less than or equal to the specified date.  This can be combined with execution_date_gte parameter to receive only the selected period.  | [optional] 
**StartDateGte** | Pointer to **time.Time** | Returns objects greater or equal the specified date.  This can be combined with start_date_lte parameter to receive only the selected period.  | [optional] 
**StartDateLte** | Pointer to **time.Time** | Returns objects less or equal the specified date.  This can be combined with start_date_gte parameter to receive only the selected period.  | [optional] 
**EndDateGte** | Pointer to **time.Time** | Returns objects greater or equal the specified date.  This can be combined with start_date_lte parameter to receive only the selected period.  | [optional] 
**EndDateLte** | Pointer to **time.Time** | Returns objects less than or equal to the specified date.  This can be combined with start_date_gte parameter to receive only the selected period.  | [optional] 
**DurationGte** | Pointer to **float32** | Returns objects greater than or equal to the specified values.  This can be combined with duration_lte parameter to receive only the selected period.  | [optional] 
**DurationLte** | Pointer to **float32** | Returns objects less than or equal to the specified values.  This can be combined with duration_gte parameter to receive only the selected range.  | [optional] 
**State** | Pointer to [**[]TaskState**](TaskState.md) | The value can be repeated to retrieve multiple matching values (OR condition). | [optional] 
**Pool** | Pointer to **[]string** | The value can be repeated to retrieve multiple matching values (OR condition). | [optional] 
**Queue** | Pointer to **[]string** | The value can be repeated to retrieve multiple matching values (OR condition). | [optional] 

## Methods

### NewListTaskInstanceForm

`func NewListTaskInstanceForm() *ListTaskInstanceForm`

NewListTaskInstanceForm instantiates a new ListTaskInstanceForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTaskInstanceFormWithDefaults

`func NewListTaskInstanceFormWithDefaults() *ListTaskInstanceForm`

NewListTaskInstanceFormWithDefaults instantiates a new ListTaskInstanceForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDagIds

`func (o *ListTaskInstanceForm) GetDagIds() []string`

GetDagIds returns the DagIds field if non-nil, zero value otherwise.

### GetDagIdsOk

`func (o *ListTaskInstanceForm) GetDagIdsOk() (*[]string, bool)`

GetDagIdsOk returns a tuple with the DagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagIds

`func (o *ListTaskInstanceForm) SetDagIds(v []string)`

SetDagIds sets DagIds field to given value.

### HasDagIds

`func (o *ListTaskInstanceForm) HasDagIds() bool`

HasDagIds returns a boolean if a field has been set.

### GetDagRunIds

`func (o *ListTaskInstanceForm) GetDagRunIds() []string`

GetDagRunIds returns the DagRunIds field if non-nil, zero value otherwise.

### GetDagRunIdsOk

`func (o *ListTaskInstanceForm) GetDagRunIdsOk() (*[]string, bool)`

GetDagRunIdsOk returns a tuple with the DagRunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagRunIds

`func (o *ListTaskInstanceForm) SetDagRunIds(v []string)`

SetDagRunIds sets DagRunIds field to given value.

### HasDagRunIds

`func (o *ListTaskInstanceForm) HasDagRunIds() bool`

HasDagRunIds returns a boolean if a field has been set.

### GetTaskIds

`func (o *ListTaskInstanceForm) GetTaskIds() []string`

GetTaskIds returns the TaskIds field if non-nil, zero value otherwise.

### GetTaskIdsOk

`func (o *ListTaskInstanceForm) GetTaskIdsOk() (*[]string, bool)`

GetTaskIdsOk returns a tuple with the TaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIds

`func (o *ListTaskInstanceForm) SetTaskIds(v []string)`

SetTaskIds sets TaskIds field to given value.

### HasTaskIds

`func (o *ListTaskInstanceForm) HasTaskIds() bool`

HasTaskIds returns a boolean if a field has been set.

### GetExecutionDateGte

`func (o *ListTaskInstanceForm) GetExecutionDateGte() time.Time`

GetExecutionDateGte returns the ExecutionDateGte field if non-nil, zero value otherwise.

### GetExecutionDateGteOk

`func (o *ListTaskInstanceForm) GetExecutionDateGteOk() (*time.Time, bool)`

GetExecutionDateGteOk returns a tuple with the ExecutionDateGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDateGte

`func (o *ListTaskInstanceForm) SetExecutionDateGte(v time.Time)`

SetExecutionDateGte sets ExecutionDateGte field to given value.

### HasExecutionDateGte

`func (o *ListTaskInstanceForm) HasExecutionDateGte() bool`

HasExecutionDateGte returns a boolean if a field has been set.

### GetExecutionDateLte

`func (o *ListTaskInstanceForm) GetExecutionDateLte() time.Time`

GetExecutionDateLte returns the ExecutionDateLte field if non-nil, zero value otherwise.

### GetExecutionDateLteOk

`func (o *ListTaskInstanceForm) GetExecutionDateLteOk() (*time.Time, bool)`

GetExecutionDateLteOk returns a tuple with the ExecutionDateLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDateLte

`func (o *ListTaskInstanceForm) SetExecutionDateLte(v time.Time)`

SetExecutionDateLte sets ExecutionDateLte field to given value.

### HasExecutionDateLte

`func (o *ListTaskInstanceForm) HasExecutionDateLte() bool`

HasExecutionDateLte returns a boolean if a field has been set.

### GetStartDateGte

`func (o *ListTaskInstanceForm) GetStartDateGte() time.Time`

GetStartDateGte returns the StartDateGte field if non-nil, zero value otherwise.

### GetStartDateGteOk

`func (o *ListTaskInstanceForm) GetStartDateGteOk() (*time.Time, bool)`

GetStartDateGteOk returns a tuple with the StartDateGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateGte

`func (o *ListTaskInstanceForm) SetStartDateGte(v time.Time)`

SetStartDateGte sets StartDateGte field to given value.

### HasStartDateGte

`func (o *ListTaskInstanceForm) HasStartDateGte() bool`

HasStartDateGte returns a boolean if a field has been set.

### GetStartDateLte

`func (o *ListTaskInstanceForm) GetStartDateLte() time.Time`

GetStartDateLte returns the StartDateLte field if non-nil, zero value otherwise.

### GetStartDateLteOk

`func (o *ListTaskInstanceForm) GetStartDateLteOk() (*time.Time, bool)`

GetStartDateLteOk returns a tuple with the StartDateLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateLte

`func (o *ListTaskInstanceForm) SetStartDateLte(v time.Time)`

SetStartDateLte sets StartDateLte field to given value.

### HasStartDateLte

`func (o *ListTaskInstanceForm) HasStartDateLte() bool`

HasStartDateLte returns a boolean if a field has been set.

### GetEndDateGte

`func (o *ListTaskInstanceForm) GetEndDateGte() time.Time`

GetEndDateGte returns the EndDateGte field if non-nil, zero value otherwise.

### GetEndDateGteOk

`func (o *ListTaskInstanceForm) GetEndDateGteOk() (*time.Time, bool)`

GetEndDateGteOk returns a tuple with the EndDateGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateGte

`func (o *ListTaskInstanceForm) SetEndDateGte(v time.Time)`

SetEndDateGte sets EndDateGte field to given value.

### HasEndDateGte

`func (o *ListTaskInstanceForm) HasEndDateGte() bool`

HasEndDateGte returns a boolean if a field has been set.

### GetEndDateLte

`func (o *ListTaskInstanceForm) GetEndDateLte() time.Time`

GetEndDateLte returns the EndDateLte field if non-nil, zero value otherwise.

### GetEndDateLteOk

`func (o *ListTaskInstanceForm) GetEndDateLteOk() (*time.Time, bool)`

GetEndDateLteOk returns a tuple with the EndDateLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateLte

`func (o *ListTaskInstanceForm) SetEndDateLte(v time.Time)`

SetEndDateLte sets EndDateLte field to given value.

### HasEndDateLte

`func (o *ListTaskInstanceForm) HasEndDateLte() bool`

HasEndDateLte returns a boolean if a field has been set.

### GetDurationGte

`func (o *ListTaskInstanceForm) GetDurationGte() float32`

GetDurationGte returns the DurationGte field if non-nil, zero value otherwise.

### GetDurationGteOk

`func (o *ListTaskInstanceForm) GetDurationGteOk() (*float32, bool)`

GetDurationGteOk returns a tuple with the DurationGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationGte

`func (o *ListTaskInstanceForm) SetDurationGte(v float32)`

SetDurationGte sets DurationGte field to given value.

### HasDurationGte

`func (o *ListTaskInstanceForm) HasDurationGte() bool`

HasDurationGte returns a boolean if a field has been set.

### GetDurationLte

`func (o *ListTaskInstanceForm) GetDurationLte() float32`

GetDurationLte returns the DurationLte field if non-nil, zero value otherwise.

### GetDurationLteOk

`func (o *ListTaskInstanceForm) GetDurationLteOk() (*float32, bool)`

GetDurationLteOk returns a tuple with the DurationLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationLte

`func (o *ListTaskInstanceForm) SetDurationLte(v float32)`

SetDurationLte sets DurationLte field to given value.

### HasDurationLte

`func (o *ListTaskInstanceForm) HasDurationLte() bool`

HasDurationLte returns a boolean if a field has been set.

### GetState

`func (o *ListTaskInstanceForm) GetState() []TaskState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ListTaskInstanceForm) GetStateOk() (*[]TaskState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ListTaskInstanceForm) SetState(v []TaskState)`

SetState sets State field to given value.

### HasState

`func (o *ListTaskInstanceForm) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPool

`func (o *ListTaskInstanceForm) GetPool() []string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *ListTaskInstanceForm) GetPoolOk() (*[]string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *ListTaskInstanceForm) SetPool(v []string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *ListTaskInstanceForm) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetQueue

`func (o *ListTaskInstanceForm) GetQueue() []string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *ListTaskInstanceForm) GetQueueOk() (*[]string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *ListTaskInstanceForm) SetQueue(v []string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *ListTaskInstanceForm) HasQueue() bool`

HasQueue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


