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

# DAGDetailAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | Pointer to **string** |  | [optional] 
**Catchup** | Pointer to **bool** |  | [optional] [readonly] 
**Orientation** | Pointer to **string** |  | [optional] [readonly] 
**Concurrency** | Pointer to **float32** |  | [optional] [readonly] 
**StartDate** | Pointer to **NullableTime** |  | [optional] [readonly] 
**DagRunTimeout** | Pointer to [**TimeDelta**](TimeDelta.md) |  | [optional] 
**DocMd** | Pointer to **NullableString** |  | [optional] [readonly] 
**DefaultView** | Pointer to **string** |  | [optional] [readonly] 
**Params** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewDAGDetailAllOf

`func NewDAGDetailAllOf() *DAGDetailAllOf`

NewDAGDetailAllOf instantiates a new DAGDetailAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDAGDetailAllOfWithDefaults

`func NewDAGDetailAllOfWithDefaults() *DAGDetailAllOf`

NewDAGDetailAllOfWithDefaults instantiates a new DAGDetailAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezone

`func (o *DAGDetailAllOf) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *DAGDetailAllOf) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *DAGDetailAllOf) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *DAGDetailAllOf) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetCatchup

`func (o *DAGDetailAllOf) GetCatchup() bool`

GetCatchup returns the Catchup field if non-nil, zero value otherwise.

### GetCatchupOk

`func (o *DAGDetailAllOf) GetCatchupOk() (*bool, bool)`

GetCatchupOk returns a tuple with the Catchup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatchup

`func (o *DAGDetailAllOf) SetCatchup(v bool)`

SetCatchup sets Catchup field to given value.

### HasCatchup

`func (o *DAGDetailAllOf) HasCatchup() bool`

HasCatchup returns a boolean if a field has been set.

### GetOrientation

`func (o *DAGDetailAllOf) GetOrientation() string`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *DAGDetailAllOf) GetOrientationOk() (*string, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *DAGDetailAllOf) SetOrientation(v string)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *DAGDetailAllOf) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetConcurrency

`func (o *DAGDetailAllOf) GetConcurrency() float32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *DAGDetailAllOf) GetConcurrencyOk() (*float32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *DAGDetailAllOf) SetConcurrency(v float32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *DAGDetailAllOf) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### GetStartDate

`func (o *DAGDetailAllOf) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DAGDetailAllOf) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DAGDetailAllOf) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DAGDetailAllOf) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *DAGDetailAllOf) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *DAGDetailAllOf) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetDagRunTimeout

`func (o *DAGDetailAllOf) GetDagRunTimeout() TimeDelta`

GetDagRunTimeout returns the DagRunTimeout field if non-nil, zero value otherwise.

### GetDagRunTimeoutOk

`func (o *DAGDetailAllOf) GetDagRunTimeoutOk() (*TimeDelta, bool)`

GetDagRunTimeoutOk returns a tuple with the DagRunTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagRunTimeout

`func (o *DAGDetailAllOf) SetDagRunTimeout(v TimeDelta)`

SetDagRunTimeout sets DagRunTimeout field to given value.

### HasDagRunTimeout

`func (o *DAGDetailAllOf) HasDagRunTimeout() bool`

HasDagRunTimeout returns a boolean if a field has been set.

### GetDocMd

`func (o *DAGDetailAllOf) GetDocMd() string`

GetDocMd returns the DocMd field if non-nil, zero value otherwise.

### GetDocMdOk

`func (o *DAGDetailAllOf) GetDocMdOk() (*string, bool)`

GetDocMdOk returns a tuple with the DocMd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocMd

`func (o *DAGDetailAllOf) SetDocMd(v string)`

SetDocMd sets DocMd field to given value.

### HasDocMd

`func (o *DAGDetailAllOf) HasDocMd() bool`

HasDocMd returns a boolean if a field has been set.

### SetDocMdNil

`func (o *DAGDetailAllOf) SetDocMdNil(b bool)`

 SetDocMdNil sets the value for DocMd to be an explicit nil

### UnsetDocMd
`func (o *DAGDetailAllOf) UnsetDocMd()`

UnsetDocMd ensures that no value is present for DocMd, not even an explicit nil
### GetDefaultView

`func (o *DAGDetailAllOf) GetDefaultView() string`

GetDefaultView returns the DefaultView field if non-nil, zero value otherwise.

### GetDefaultViewOk

`func (o *DAGDetailAllOf) GetDefaultViewOk() (*string, bool)`

GetDefaultViewOk returns a tuple with the DefaultView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultView

`func (o *DAGDetailAllOf) SetDefaultView(v string)`

SetDefaultView sets DefaultView field to given value.

### HasDefaultView

`func (o *DAGDetailAllOf) HasDefaultView() bool`

HasDefaultView returns a boolean if a field has been set.

### GetParams

`func (o *DAGDetailAllOf) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *DAGDetailAllOf) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *DAGDetailAllOf) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *DAGDetailAllOf) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


