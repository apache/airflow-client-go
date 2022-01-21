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

# DAG

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DagId** | Pointer to **string** | The ID of the DAG. | [optional] [readonly] 
**RootDagId** | Pointer to **NullableString** | If the DAG is SubDAG then it is the top level DAG identifier. Otherwise, null. | [optional] [readonly] 
**IsPaused** | Pointer to **NullableBool** | Whether the DAG is paused. | [optional] 
**IsActive** | Pointer to **NullableBool** | Whether the DAG is currently seen by the scheduler(s).  *New in version 2.1.1*  *Changed in version 2.2.0*&amp;#58; Field is read-only.  | [optional] [readonly] 
**IsSubdag** | Pointer to **bool** | Whether the DAG is SubDAG. | [optional] [readonly] 
**Fileloc** | Pointer to **string** | The absolute path to the file. | [optional] [readonly] 
**FileToken** | Pointer to **string** | The key containing the encrypted path to the file. Encryption and decryption take place only on the server. This prevents the client from reading an non-DAG file. This also ensures API extensibility, because the format of encrypted data may change.  | [optional] [readonly] 
**Owners** | Pointer to **[]string** |  | [optional] [readonly] 
**Description** | Pointer to **NullableString** | User-provided DAG description, which can consist of several sentences or paragraphs that describe DAG contents.  | [optional] [readonly] 
**ScheduleInterval** | Pointer to [**ScheduleInterval**](ScheduleInterval.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | List of tags. | [optional] [readonly] 

## Methods

### NewDAG

`func NewDAG() *DAG`

NewDAG instantiates a new DAG object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDAGWithDefaults

`func NewDAGWithDefaults() *DAG`

NewDAGWithDefaults instantiates a new DAG object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDagId

`func (o *DAG) GetDagId() string`

GetDagId returns the DagId field if non-nil, zero value otherwise.

### GetDagIdOk

`func (o *DAG) GetDagIdOk() (*string, bool)`

GetDagIdOk returns a tuple with the DagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagId

`func (o *DAG) SetDagId(v string)`

SetDagId sets DagId field to given value.

### HasDagId

`func (o *DAG) HasDagId() bool`

HasDagId returns a boolean if a field has been set.

### GetRootDagId

`func (o *DAG) GetRootDagId() string`

GetRootDagId returns the RootDagId field if non-nil, zero value otherwise.

### GetRootDagIdOk

`func (o *DAG) GetRootDagIdOk() (*string, bool)`

GetRootDagIdOk returns a tuple with the RootDagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDagId

`func (o *DAG) SetRootDagId(v string)`

SetRootDagId sets RootDagId field to given value.

### HasRootDagId

`func (o *DAG) HasRootDagId() bool`

HasRootDagId returns a boolean if a field has been set.

### SetRootDagIdNil

`func (o *DAG) SetRootDagIdNil(b bool)`

 SetRootDagIdNil sets the value for RootDagId to be an explicit nil

### UnsetRootDagId
`func (o *DAG) UnsetRootDagId()`

UnsetRootDagId ensures that no value is present for RootDagId, not even an explicit nil
### GetIsPaused

`func (o *DAG) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *DAG) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *DAG) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.

### HasIsPaused

`func (o *DAG) HasIsPaused() bool`

HasIsPaused returns a boolean if a field has been set.

### SetIsPausedNil

`func (o *DAG) SetIsPausedNil(b bool)`

 SetIsPausedNil sets the value for IsPaused to be an explicit nil

### UnsetIsPaused
`func (o *DAG) UnsetIsPaused()`

UnsetIsPaused ensures that no value is present for IsPaused, not even an explicit nil
### GetIsActive

`func (o *DAG) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DAG) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DAG) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *DAG) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *DAG) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *DAG) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetIsSubdag

`func (o *DAG) GetIsSubdag() bool`

GetIsSubdag returns the IsSubdag field if non-nil, zero value otherwise.

### GetIsSubdagOk

`func (o *DAG) GetIsSubdagOk() (*bool, bool)`

GetIsSubdagOk returns a tuple with the IsSubdag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubdag

`func (o *DAG) SetIsSubdag(v bool)`

SetIsSubdag sets IsSubdag field to given value.

### HasIsSubdag

`func (o *DAG) HasIsSubdag() bool`

HasIsSubdag returns a boolean if a field has been set.

### GetFileloc

`func (o *DAG) GetFileloc() string`

GetFileloc returns the Fileloc field if non-nil, zero value otherwise.

### GetFilelocOk

`func (o *DAG) GetFilelocOk() (*string, bool)`

GetFilelocOk returns a tuple with the Fileloc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileloc

`func (o *DAG) SetFileloc(v string)`

SetFileloc sets Fileloc field to given value.

### HasFileloc

`func (o *DAG) HasFileloc() bool`

HasFileloc returns a boolean if a field has been set.

### GetFileToken

`func (o *DAG) GetFileToken() string`

GetFileToken returns the FileToken field if non-nil, zero value otherwise.

### GetFileTokenOk

`func (o *DAG) GetFileTokenOk() (*string, bool)`

GetFileTokenOk returns a tuple with the FileToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileToken

`func (o *DAG) SetFileToken(v string)`

SetFileToken sets FileToken field to given value.

### HasFileToken

`func (o *DAG) HasFileToken() bool`

HasFileToken returns a boolean if a field has been set.

### GetOwners

`func (o *DAG) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *DAG) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *DAG) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *DAG) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetDescription

`func (o *DAG) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DAG) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DAG) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DAG) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DAG) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DAG) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetScheduleInterval

`func (o *DAG) GetScheduleInterval() ScheduleInterval`

GetScheduleInterval returns the ScheduleInterval field if non-nil, zero value otherwise.

### GetScheduleIntervalOk

`func (o *DAG) GetScheduleIntervalOk() (*ScheduleInterval, bool)`

GetScheduleIntervalOk returns a tuple with the ScheduleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleInterval

`func (o *DAG) SetScheduleInterval(v ScheduleInterval)`

SetScheduleInterval sets ScheduleInterval field to given value.

### HasScheduleInterval

`func (o *DAG) HasScheduleInterval() bool`

HasScheduleInterval returns a boolean if a field has been set.

### GetTags

`func (o *DAG) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DAG) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DAG) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DAG) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *DAG) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *DAG) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


