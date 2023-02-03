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

# Trigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Classpath** | Pointer to **string** |  | [optional] 
**Kwargs** | Pointer to **string** |  | [optional] 
**CreatedDate** | Pointer to **string** |  | [optional] 
**TriggererId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewTrigger

`func NewTrigger() *Trigger`

NewTrigger instantiates a new Trigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerWithDefaults

`func NewTriggerWithDefaults() *Trigger`

NewTriggerWithDefaults instantiates a new Trigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Trigger) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Trigger) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Trigger) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Trigger) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClasspath

`func (o *Trigger) GetClasspath() string`

GetClasspath returns the Classpath field if non-nil, zero value otherwise.

### GetClasspathOk

`func (o *Trigger) GetClasspathOk() (*string, bool)`

GetClasspathOk returns a tuple with the Classpath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasspath

`func (o *Trigger) SetClasspath(v string)`

SetClasspath sets Classpath field to given value.

### HasClasspath

`func (o *Trigger) HasClasspath() bool`

HasClasspath returns a boolean if a field has been set.

### GetKwargs

`func (o *Trigger) GetKwargs() string`

GetKwargs returns the Kwargs field if non-nil, zero value otherwise.

### GetKwargsOk

`func (o *Trigger) GetKwargsOk() (*string, bool)`

GetKwargsOk returns a tuple with the Kwargs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKwargs

`func (o *Trigger) SetKwargs(v string)`

SetKwargs sets Kwargs field to given value.

### HasKwargs

`func (o *Trigger) HasKwargs() bool`

HasKwargs returns a boolean if a field has been set.

### GetCreatedDate

`func (o *Trigger) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *Trigger) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *Trigger) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *Trigger) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetTriggererId

`func (o *Trigger) GetTriggererId() int32`

GetTriggererId returns the TriggererId field if non-nil, zero value otherwise.

### GetTriggererIdOk

`func (o *Trigger) GetTriggererIdOk() (*int32, bool)`

GetTriggererIdOk returns a tuple with the TriggererId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggererId

`func (o *Trigger) SetTriggererId(v int32)`

SetTriggererId sets TriggererId field to given value.

### HasTriggererId

`func (o *Trigger) HasTriggererId() bool`

HasTriggererId returns a boolean if a field has been set.

### SetTriggererIdNil

`func (o *Trigger) SetTriggererIdNil(b bool)`

 SetTriggererIdNil sets the value for TriggererId to be an explicit nil

### UnsetTriggererId
`func (o *Trigger) UnsetTriggererId()`

UnsetTriggererId ensures that no value is present for TriggererId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


