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

# Pool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of pool. | [optional] 
**Slots** | Pointer to **int32** | The maximum number of slots that can be assigned to tasks. One job may occupy one or more slots.  | [optional] 
**OccupiedSlots** | Pointer to **int32** | The number of slots used by running/queued tasks at the moment. | [optional] [readonly] 
**UsedSlots** | Pointer to **int32** | The number of slots used by running tasks at the moment. | [optional] [readonly] 
**QueuedSlots** | Pointer to **int32** | The number of slots used by queued tasks at the moment. | [optional] [readonly] 
**OpenSlots** | Pointer to **int32** | The number of free slots at the moment. | [optional] [readonly] 
**Description** | Pointer to **NullableString** | The description of the pool.  *New in version 2.3.0*  | [optional] 

## Methods

### NewPool

`func NewPool() *Pool`

NewPool instantiates a new Pool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolWithDefaults

`func NewPoolWithDefaults() *Pool`

NewPoolWithDefaults instantiates a new Pool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Pool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Pool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlots

`func (o *Pool) GetSlots() int32`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *Pool) GetSlotsOk() (*int32, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *Pool) SetSlots(v int32)`

SetSlots sets Slots field to given value.

### HasSlots

`func (o *Pool) HasSlots() bool`

HasSlots returns a boolean if a field has been set.

### GetOccupiedSlots

`func (o *Pool) GetOccupiedSlots() int32`

GetOccupiedSlots returns the OccupiedSlots field if non-nil, zero value otherwise.

### GetOccupiedSlotsOk

`func (o *Pool) GetOccupiedSlotsOk() (*int32, bool)`

GetOccupiedSlotsOk returns a tuple with the OccupiedSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupiedSlots

`func (o *Pool) SetOccupiedSlots(v int32)`

SetOccupiedSlots sets OccupiedSlots field to given value.

### HasOccupiedSlots

`func (o *Pool) HasOccupiedSlots() bool`

HasOccupiedSlots returns a boolean if a field has been set.

### GetUsedSlots

`func (o *Pool) GetUsedSlots() int32`

GetUsedSlots returns the UsedSlots field if non-nil, zero value otherwise.

### GetUsedSlotsOk

`func (o *Pool) GetUsedSlotsOk() (*int32, bool)`

GetUsedSlotsOk returns a tuple with the UsedSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSlots

`func (o *Pool) SetUsedSlots(v int32)`

SetUsedSlots sets UsedSlots field to given value.

### HasUsedSlots

`func (o *Pool) HasUsedSlots() bool`

HasUsedSlots returns a boolean if a field has been set.

### GetQueuedSlots

`func (o *Pool) GetQueuedSlots() int32`

GetQueuedSlots returns the QueuedSlots field if non-nil, zero value otherwise.

### GetQueuedSlotsOk

`func (o *Pool) GetQueuedSlotsOk() (*int32, bool)`

GetQueuedSlotsOk returns a tuple with the QueuedSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedSlots

`func (o *Pool) SetQueuedSlots(v int32)`

SetQueuedSlots sets QueuedSlots field to given value.

### HasQueuedSlots

`func (o *Pool) HasQueuedSlots() bool`

HasQueuedSlots returns a boolean if a field has been set.

### GetOpenSlots

`func (o *Pool) GetOpenSlots() int32`

GetOpenSlots returns the OpenSlots field if non-nil, zero value otherwise.

### GetOpenSlotsOk

`func (o *Pool) GetOpenSlotsOk() (*int32, bool)`

GetOpenSlotsOk returns a tuple with the OpenSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenSlots

`func (o *Pool) SetOpenSlots(v int32)`

SetOpenSlots sets OpenSlots field to given value.

### HasOpenSlots

`func (o *Pool) HasOpenSlots() bool`

HasOpenSlots returns a boolean if a field has been set.

### GetDescription

`func (o *Pool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Pool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Pool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Pool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Pool) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Pool) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


