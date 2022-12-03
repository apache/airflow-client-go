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

# Dataset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The dataset id | [optional] 
**Uri** | Pointer to **string** | The dataset uri | [optional] 
**Extra** | Pointer to **map[string]interface{}** | The dataset extra | [optional] 
**CreatedAt** | Pointer to **string** | The dataset creation time | [optional] 
**UpdatedAt** | Pointer to **string** | The dataset update time | [optional] 
**ConsumingDags** | Pointer to [**[]DagScheduleDatasetReference**](DagScheduleDatasetReference.md) |  | [optional] 
**ProducingTasks** | Pointer to [**[]TaskOutletDatasetReference**](TaskOutletDatasetReference.md) |  | [optional] 

## Methods

### NewDataset

`func NewDataset() *Dataset`

NewDataset instantiates a new Dataset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetWithDefaults

`func NewDatasetWithDefaults() *Dataset`

NewDatasetWithDefaults instantiates a new Dataset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Dataset) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Dataset) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Dataset) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Dataset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUri

`func (o *Dataset) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Dataset) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Dataset) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Dataset) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetExtra

`func (o *Dataset) GetExtra() map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *Dataset) GetExtraOk() (*map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *Dataset) SetExtra(v map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *Dataset) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### SetExtraNil

`func (o *Dataset) SetExtraNil(b bool)`

 SetExtraNil sets the value for Extra to be an explicit nil

### UnsetExtra
`func (o *Dataset) UnsetExtra()`

UnsetExtra ensures that no value is present for Extra, not even an explicit nil
### GetCreatedAt

`func (o *Dataset) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Dataset) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Dataset) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Dataset) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Dataset) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Dataset) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Dataset) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Dataset) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetConsumingDags

`func (o *Dataset) GetConsumingDags() []DagScheduleDatasetReference`

GetConsumingDags returns the ConsumingDags field if non-nil, zero value otherwise.

### GetConsumingDagsOk

`func (o *Dataset) GetConsumingDagsOk() (*[]DagScheduleDatasetReference, bool)`

GetConsumingDagsOk returns a tuple with the ConsumingDags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumingDags

`func (o *Dataset) SetConsumingDags(v []DagScheduleDatasetReference)`

SetConsumingDags sets ConsumingDags field to given value.

### HasConsumingDags

`func (o *Dataset) HasConsumingDags() bool`

HasConsumingDags returns a boolean if a field has been set.

### GetProducingTasks

`func (o *Dataset) GetProducingTasks() []TaskOutletDatasetReference`

GetProducingTasks returns the ProducingTasks field if non-nil, zero value otherwise.

### GetProducingTasksOk

`func (o *Dataset) GetProducingTasksOk() (*[]TaskOutletDatasetReference, bool)`

GetProducingTasksOk returns a tuple with the ProducingTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducingTasks

`func (o *Dataset) SetProducingTasks(v []TaskOutletDatasetReference)`

SetProducingTasks sets ProducingTasks field to given value.

### HasProducingTasks

`func (o *Dataset) HasProducingTasks() bool`

HasProducingTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


