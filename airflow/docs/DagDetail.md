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

# DAGDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DagId** | Pointer to **string** | The ID of the DAG. | [optional] [readonly] 
**RootDagId** | Pointer to **NullableString** | If the DAG is SubDAG then it is the top level DAG identifier. Otherwise, null. | [optional] [readonly] 
**IsPaused** | Pointer to **NullableBool** | Whether the DAG is paused. | [optional] 
**IsActive** | Pointer to **NullableBool** | Whether the DAG is currently seen by the scheduler(s). | [optional] [readonly] 
**IsSubdag** | Pointer to **bool** | Whether the DAG is SubDAG. | [optional] [readonly] 
**Fileloc** | Pointer to **string** | The absolute path to the file. | [optional] [readonly] 
**FileToken** | Pointer to **string** | The key containing the encrypted path to the file. Encryption and decryption take place only on the server. This prevents the client from reading an non-DAG file. This also ensures API extensibility, because the format of encrypted data may change.  | [optional] [readonly] 
**Owners** | Pointer to **[]string** |  | [optional] [readonly] 
**Description** | Pointer to **NullableString** | User-provided DAG description, which can consist of several sentences or paragraphs that describe DAG contents.  | [optional] [readonly] 
**ScheduleInterval** | Pointer to [**ScheduleInterval**](ScheduleInterval.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | List of tags. | [optional] [readonly] 
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

### NewDAGDetail

`func NewDAGDetail() *DAGDetail`

NewDAGDetail instantiates a new DAGDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDAGDetailWithDefaults

`func NewDAGDetailWithDefaults() *DAGDetail`

NewDAGDetailWithDefaults instantiates a new DAGDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDagId

`func (o *DAGDetail) GetDagId() string`

GetDagId returns the DagId field if non-nil, zero value otherwise.

### GetDagIdOk

`func (o *DAGDetail) GetDagIdOk() (*string, bool)`

GetDagIdOk returns a tuple with the DagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagId

`func (o *DAGDetail) SetDagId(v string)`

SetDagId sets DagId field to given value.

### HasDagId

`func (o *DAGDetail) HasDagId() bool`

HasDagId returns a boolean if a field has been set.

### GetRootDagId

`func (o *DAGDetail) GetRootDagId() string`

GetRootDagId returns the RootDagId field if non-nil, zero value otherwise.

### GetRootDagIdOk

`func (o *DAGDetail) GetRootDagIdOk() (*string, bool)`

GetRootDagIdOk returns a tuple with the RootDagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDagId

`func (o *DAGDetail) SetRootDagId(v string)`

SetRootDagId sets RootDagId field to given value.

### HasRootDagId

`func (o *DAGDetail) HasRootDagId() bool`

HasRootDagId returns a boolean if a field has been set.

### SetRootDagIdNil

`func (o *DAGDetail) SetRootDagIdNil(b bool)`

 SetRootDagIdNil sets the value for RootDagId to be an explicit nil

### UnsetRootDagId
`func (o *DAGDetail) UnsetRootDagId()`

UnsetRootDagId ensures that no value is present for RootDagId, not even an explicit nil
### GetIsPaused

`func (o *DAGDetail) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *DAGDetail) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *DAGDetail) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.

### HasIsPaused

`func (o *DAGDetail) HasIsPaused() bool`

HasIsPaused returns a boolean if a field has been set.

### SetIsPausedNil

`func (o *DAGDetail) SetIsPausedNil(b bool)`

 SetIsPausedNil sets the value for IsPaused to be an explicit nil

### UnsetIsPaused
`func (o *DAGDetail) UnsetIsPaused()`

UnsetIsPaused ensures that no value is present for IsPaused, not even an explicit nil
### GetIsActive

`func (o *DAGDetail) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DAGDetail) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DAGDetail) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *DAGDetail) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *DAGDetail) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *DAGDetail) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetIsSubdag

`func (o *DAGDetail) GetIsSubdag() bool`

GetIsSubdag returns the IsSubdag field if non-nil, zero value otherwise.

### GetIsSubdagOk

`func (o *DAGDetail) GetIsSubdagOk() (*bool, bool)`

GetIsSubdagOk returns a tuple with the IsSubdag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubdag

`func (o *DAGDetail) SetIsSubdag(v bool)`

SetIsSubdag sets IsSubdag field to given value.

### HasIsSubdag

`func (o *DAGDetail) HasIsSubdag() bool`

HasIsSubdag returns a boolean if a field has been set.

### GetFileloc

`func (o *DAGDetail) GetFileloc() string`

GetFileloc returns the Fileloc field if non-nil, zero value otherwise.

### GetFilelocOk

`func (o *DAGDetail) GetFilelocOk() (*string, bool)`

GetFilelocOk returns a tuple with the Fileloc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileloc

`func (o *DAGDetail) SetFileloc(v string)`

SetFileloc sets Fileloc field to given value.

### HasFileloc

`func (o *DAGDetail) HasFileloc() bool`

HasFileloc returns a boolean if a field has been set.

### GetFileToken

`func (o *DAGDetail) GetFileToken() string`

GetFileToken returns the FileToken field if non-nil, zero value otherwise.

### GetFileTokenOk

`func (o *DAGDetail) GetFileTokenOk() (*string, bool)`

GetFileTokenOk returns a tuple with the FileToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileToken

`func (o *DAGDetail) SetFileToken(v string)`

SetFileToken sets FileToken field to given value.

### HasFileToken

`func (o *DAGDetail) HasFileToken() bool`

HasFileToken returns a boolean if a field has been set.

### GetOwners

`func (o *DAGDetail) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *DAGDetail) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *DAGDetail) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *DAGDetail) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetDescription

`func (o *DAGDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DAGDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DAGDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DAGDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DAGDetail) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DAGDetail) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetScheduleInterval

`func (o *DAGDetail) GetScheduleInterval() ScheduleInterval`

GetScheduleInterval returns the ScheduleInterval field if non-nil, zero value otherwise.

### GetScheduleIntervalOk

`func (o *DAGDetail) GetScheduleIntervalOk() (*ScheduleInterval, bool)`

GetScheduleIntervalOk returns a tuple with the ScheduleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleInterval

`func (o *DAGDetail) SetScheduleInterval(v ScheduleInterval)`

SetScheduleInterval sets ScheduleInterval field to given value.

### HasScheduleInterval

`func (o *DAGDetail) HasScheduleInterval() bool`

HasScheduleInterval returns a boolean if a field has been set.

### GetTags

`func (o *DAGDetail) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DAGDetail) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DAGDetail) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DAGDetail) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *DAGDetail) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *DAGDetail) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTimezone

`func (o *DAGDetail) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *DAGDetail) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *DAGDetail) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *DAGDetail) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetCatchup

`func (o *DAGDetail) GetCatchup() bool`

GetCatchup returns the Catchup field if non-nil, zero value otherwise.

### GetCatchupOk

`func (o *DAGDetail) GetCatchupOk() (*bool, bool)`

GetCatchupOk returns a tuple with the Catchup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatchup

`func (o *DAGDetail) SetCatchup(v bool)`

SetCatchup sets Catchup field to given value.

### HasCatchup

`func (o *DAGDetail) HasCatchup() bool`

HasCatchup returns a boolean if a field has been set.

### GetOrientation

`func (o *DAGDetail) GetOrientation() string`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *DAGDetail) GetOrientationOk() (*string, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *DAGDetail) SetOrientation(v string)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *DAGDetail) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetConcurrency

`func (o *DAGDetail) GetConcurrency() float32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *DAGDetail) GetConcurrencyOk() (*float32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *DAGDetail) SetConcurrency(v float32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *DAGDetail) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### GetStartDate

`func (o *DAGDetail) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DAGDetail) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DAGDetail) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DAGDetail) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *DAGDetail) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *DAGDetail) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetDagRunTimeout

`func (o *DAGDetail) GetDagRunTimeout() TimeDelta`

GetDagRunTimeout returns the DagRunTimeout field if non-nil, zero value otherwise.

### GetDagRunTimeoutOk

`func (o *DAGDetail) GetDagRunTimeoutOk() (*TimeDelta, bool)`

GetDagRunTimeoutOk returns a tuple with the DagRunTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagRunTimeout

`func (o *DAGDetail) SetDagRunTimeout(v TimeDelta)`

SetDagRunTimeout sets DagRunTimeout field to given value.

### HasDagRunTimeout

`func (o *DAGDetail) HasDagRunTimeout() bool`

HasDagRunTimeout returns a boolean if a field has been set.

### GetDocMd

`func (o *DAGDetail) GetDocMd() string`

GetDocMd returns the DocMd field if non-nil, zero value otherwise.

### GetDocMdOk

`func (o *DAGDetail) GetDocMdOk() (*string, bool)`

GetDocMdOk returns a tuple with the DocMd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocMd

`func (o *DAGDetail) SetDocMd(v string)`

SetDocMd sets DocMd field to given value.

### HasDocMd

`func (o *DAGDetail) HasDocMd() bool`

HasDocMd returns a boolean if a field has been set.

### SetDocMdNil

`func (o *DAGDetail) SetDocMdNil(b bool)`

 SetDocMdNil sets the value for DocMd to be an explicit nil

### UnsetDocMd
`func (o *DAGDetail) UnsetDocMd()`

UnsetDocMd ensures that no value is present for DocMd, not even an explicit nil
### GetDefaultView

`func (o *DAGDetail) GetDefaultView() string`

GetDefaultView returns the DefaultView field if non-nil, zero value otherwise.

### GetDefaultViewOk

`func (o *DAGDetail) GetDefaultViewOk() (*string, bool)`

GetDefaultViewOk returns a tuple with the DefaultView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultView

`func (o *DAGDetail) SetDefaultView(v string)`

SetDefaultView sets DefaultView field to given value.

### HasDefaultView

`func (o *DAGDetail) HasDefaultView() bool`

HasDefaultView returns a boolean if a field has been set.

### GetParams

`func (o *DAGDetail) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *DAGDetail) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *DAGDetail) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *DAGDetail) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


