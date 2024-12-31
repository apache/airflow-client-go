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
**IsActive** | Pointer to **NullableBool** | Whether the DAG is currently seen by the scheduler(s).  *New in version 2.1.1*  *Changed in version 2.2.0*&amp;#58; Field is read-only.  | [optional] [readonly] 
**IsSubdag** | Pointer to **bool** | Whether the DAG is SubDAG. | [optional] [readonly] 
**LastParsedTime** | Pointer to **NullableTime** | The last time the DAG was parsed.  *New in version 2.3.0*  | [optional] [readonly] 
**LastPickled** | Pointer to **NullableTime** | The last time the DAG was pickled.  *New in version 2.3.0*  | [optional] [readonly] 
**LastExpired** | Pointer to **NullableTime** | Time when the DAG last received a refresh signal (e.g. the DAG&#39;s \&quot;refresh\&quot; button was clicked in the web UI)  *New in version 2.3.0*  | [optional] [readonly] 
**SchedulerLock** | Pointer to **NullableBool** | Whether (one of) the scheduler is scheduling this DAG at the moment  *New in version 2.3.0*  | [optional] [readonly] 
**PickleId** | Pointer to **NullableString** | Foreign key to the latest pickle_id  *New in version 2.3.0*  | [optional] [readonly] 
**DefaultView** | Pointer to **string** |  | [optional] [readonly] 
**Fileloc** | Pointer to **string** | The absolute path to the file. | [optional] [readonly] 
**FileToken** | Pointer to **string** | The key containing the encrypted path to the file. Encryption and decryption take place only on the server. This prevents the client from reading an non-DAG file. This also ensures API extensibility, because the format of encrypted data may change.  | [optional] [readonly] 
**Owners** | Pointer to **[]string** |  | [optional] [readonly] 
**Description** | Pointer to **NullableString** | User-provided DAG description, which can consist of several sentences or paragraphs that describe DAG contents.  | [optional] [readonly] 
**ScheduleInterval** | Pointer to [**NullableScheduleInterval**](ScheduleInterval.md) |  | [optional] 
**TimetableDescription** | Pointer to **NullableString** | Timetable/Schedule Interval description.  *New in version 2.3.0*  | [optional] [readonly] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | List of tags. | [optional] [readonly] 
**MaxActiveTasks** | Pointer to **NullableInt32** | Maximum number of active tasks that can be run on the DAG  *New in version 2.3.0*  | [optional] [readonly] 
**MaxActiveRuns** | Pointer to **NullableInt32** | Maximum number of active DAG runs for the DAG  *New in version 2.3.0*  | [optional] [readonly] 
**HasTaskConcurrencyLimits** | Pointer to **NullableBool** | Whether the DAG has task concurrency limits  *New in version 2.3.0*  | [optional] [readonly] 
**HasImportErrors** | Pointer to **NullableBool** | Whether the DAG has import errors  *New in version 2.3.0*  | [optional] [readonly] 
**NextDagrun** | Pointer to **NullableTime** | The logical date of the next dag run.  *New in version 2.3.0*  | [optional] [readonly] 
**NextDagrunDataIntervalStart** | Pointer to **NullableTime** | The start of the interval of the next dag run.  *New in version 2.3.0*  | [optional] [readonly] 
**NextDagrunDataIntervalEnd** | Pointer to **NullableTime** | The end of the interval of the next dag run.  *New in version 2.3.0*  | [optional] [readonly] 
**NextDagrunCreateAfter** | Pointer to **NullableTime** | Earliest time at which this &#x60;&#x60;next_dagrun&#x60;&#x60; can be created.  *New in version 2.3.0*  | [optional] [readonly] 
**Timezone** | Pointer to **string** |  | [optional] 
**Catchup** | Pointer to **bool** |  | [optional] [readonly] 
**Orientation** | Pointer to **string** |  | [optional] [readonly] 
**Concurrency** | Pointer to **float32** |  | [optional] [readonly] 
**StartDate** | Pointer to **NullableTime** | The DAG&#39;s start date.  *Changed in version 2.0.1*&amp;#58; Field becomes nullable.  | [optional] [readonly] 
**DagRunTimeout** | Pointer to [**NullableTimeDelta**](TimeDelta.md) |  | [optional] 
**DocMd** | Pointer to **NullableString** |  | [optional] [readonly] 
**Params** | Pointer to **map[string]interface{}** | User-specified DAG params.  *New in version 2.0.1*  | [optional] [readonly] 
**EndDate** | Pointer to **NullableTime** | The DAG&#39;s end date.  *New in version 2.3.0*.  | [optional] [readonly] 
**IsPausedUponCreation** | Pointer to **NullableBool** | Whether the DAG is paused upon creation.  *New in version 2.3.0*  | [optional] [readonly] 
**LastParsed** | Pointer to **NullableTime** | The last time the DAG was parsed.  *New in version 2.3.0*  | [optional] [readonly] 
**TemplateSearchPath** | Pointer to **[]string** | The template search path.  *New in version 2.3.0*  | [optional] 
**RenderTemplateAsNativeObj** | Pointer to **NullableBool** | Whether to render templates as native Python objects.  *New in version 2.3.0*  | [optional] [readonly] 

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

### GetLastParsedTime

`func (o *DAGDetail) GetLastParsedTime() time.Time`

GetLastParsedTime returns the LastParsedTime field if non-nil, zero value otherwise.

### GetLastParsedTimeOk

`func (o *DAGDetail) GetLastParsedTimeOk() (*time.Time, bool)`

GetLastParsedTimeOk returns a tuple with the LastParsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastParsedTime

`func (o *DAGDetail) SetLastParsedTime(v time.Time)`

SetLastParsedTime sets LastParsedTime field to given value.

### HasLastParsedTime

`func (o *DAGDetail) HasLastParsedTime() bool`

HasLastParsedTime returns a boolean if a field has been set.

### SetLastParsedTimeNil

`func (o *DAGDetail) SetLastParsedTimeNil(b bool)`

 SetLastParsedTimeNil sets the value for LastParsedTime to be an explicit nil

### UnsetLastParsedTime
`func (o *DAGDetail) UnsetLastParsedTime()`

UnsetLastParsedTime ensures that no value is present for LastParsedTime, not even an explicit nil
### GetLastPickled

`func (o *DAGDetail) GetLastPickled() time.Time`

GetLastPickled returns the LastPickled field if non-nil, zero value otherwise.

### GetLastPickledOk

`func (o *DAGDetail) GetLastPickledOk() (*time.Time, bool)`

GetLastPickledOk returns a tuple with the LastPickled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPickled

`func (o *DAGDetail) SetLastPickled(v time.Time)`

SetLastPickled sets LastPickled field to given value.

### HasLastPickled

`func (o *DAGDetail) HasLastPickled() bool`

HasLastPickled returns a boolean if a field has been set.

### SetLastPickledNil

`func (o *DAGDetail) SetLastPickledNil(b bool)`

 SetLastPickledNil sets the value for LastPickled to be an explicit nil

### UnsetLastPickled
`func (o *DAGDetail) UnsetLastPickled()`

UnsetLastPickled ensures that no value is present for LastPickled, not even an explicit nil
### GetLastExpired

`func (o *DAGDetail) GetLastExpired() time.Time`

GetLastExpired returns the LastExpired field if non-nil, zero value otherwise.

### GetLastExpiredOk

`func (o *DAGDetail) GetLastExpiredOk() (*time.Time, bool)`

GetLastExpiredOk returns a tuple with the LastExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExpired

`func (o *DAGDetail) SetLastExpired(v time.Time)`

SetLastExpired sets LastExpired field to given value.

### HasLastExpired

`func (o *DAGDetail) HasLastExpired() bool`

HasLastExpired returns a boolean if a field has been set.

### SetLastExpiredNil

`func (o *DAGDetail) SetLastExpiredNil(b bool)`

 SetLastExpiredNil sets the value for LastExpired to be an explicit nil

### UnsetLastExpired
`func (o *DAGDetail) UnsetLastExpired()`

UnsetLastExpired ensures that no value is present for LastExpired, not even an explicit nil
### GetSchedulerLock

`func (o *DAGDetail) GetSchedulerLock() bool`

GetSchedulerLock returns the SchedulerLock field if non-nil, zero value otherwise.

### GetSchedulerLockOk

`func (o *DAGDetail) GetSchedulerLockOk() (*bool, bool)`

GetSchedulerLockOk returns a tuple with the SchedulerLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerLock

`func (o *DAGDetail) SetSchedulerLock(v bool)`

SetSchedulerLock sets SchedulerLock field to given value.

### HasSchedulerLock

`func (o *DAGDetail) HasSchedulerLock() bool`

HasSchedulerLock returns a boolean if a field has been set.

### SetSchedulerLockNil

`func (o *DAGDetail) SetSchedulerLockNil(b bool)`

 SetSchedulerLockNil sets the value for SchedulerLock to be an explicit nil

### UnsetSchedulerLock
`func (o *DAGDetail) UnsetSchedulerLock()`

UnsetSchedulerLock ensures that no value is present for SchedulerLock, not even an explicit nil
### GetPickleId

`func (o *DAGDetail) GetPickleId() string`

GetPickleId returns the PickleId field if non-nil, zero value otherwise.

### GetPickleIdOk

`func (o *DAGDetail) GetPickleIdOk() (*string, bool)`

GetPickleIdOk returns a tuple with the PickleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickleId

`func (o *DAGDetail) SetPickleId(v string)`

SetPickleId sets PickleId field to given value.

### HasPickleId

`func (o *DAGDetail) HasPickleId() bool`

HasPickleId returns a boolean if a field has been set.

### SetPickleIdNil

`func (o *DAGDetail) SetPickleIdNil(b bool)`

 SetPickleIdNil sets the value for PickleId to be an explicit nil

### UnsetPickleId
`func (o *DAGDetail) UnsetPickleId()`

UnsetPickleId ensures that no value is present for PickleId, not even an explicit nil
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

### SetScheduleIntervalNil

`func (o *DAGDetail) SetScheduleIntervalNil(b bool)`

 SetScheduleIntervalNil sets the value for ScheduleInterval to be an explicit nil

### UnsetScheduleInterval
`func (o *DAGDetail) UnsetScheduleInterval()`

UnsetScheduleInterval ensures that no value is present for ScheduleInterval, not even an explicit nil
### GetTimetableDescription

`func (o *DAGDetail) GetTimetableDescription() string`

GetTimetableDescription returns the TimetableDescription field if non-nil, zero value otherwise.

### GetTimetableDescriptionOk

`func (o *DAGDetail) GetTimetableDescriptionOk() (*string, bool)`

GetTimetableDescriptionOk returns a tuple with the TimetableDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimetableDescription

`func (o *DAGDetail) SetTimetableDescription(v string)`

SetTimetableDescription sets TimetableDescription field to given value.

### HasTimetableDescription

`func (o *DAGDetail) HasTimetableDescription() bool`

HasTimetableDescription returns a boolean if a field has been set.

### SetTimetableDescriptionNil

`func (o *DAGDetail) SetTimetableDescriptionNil(b bool)`

 SetTimetableDescriptionNil sets the value for TimetableDescription to be an explicit nil

### UnsetTimetableDescription
`func (o *DAGDetail) UnsetTimetableDescription()`

UnsetTimetableDescription ensures that no value is present for TimetableDescription, not even an explicit nil
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
### GetMaxActiveTasks

`func (o *DAGDetail) GetMaxActiveTasks() int32`

GetMaxActiveTasks returns the MaxActiveTasks field if non-nil, zero value otherwise.

### GetMaxActiveTasksOk

`func (o *DAGDetail) GetMaxActiveTasksOk() (*int32, bool)`

GetMaxActiveTasksOk returns a tuple with the MaxActiveTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxActiveTasks

`func (o *DAGDetail) SetMaxActiveTasks(v int32)`

SetMaxActiveTasks sets MaxActiveTasks field to given value.

### HasMaxActiveTasks

`func (o *DAGDetail) HasMaxActiveTasks() bool`

HasMaxActiveTasks returns a boolean if a field has been set.

### SetMaxActiveTasksNil

`func (o *DAGDetail) SetMaxActiveTasksNil(b bool)`

 SetMaxActiveTasksNil sets the value for MaxActiveTasks to be an explicit nil

### UnsetMaxActiveTasks
`func (o *DAGDetail) UnsetMaxActiveTasks()`

UnsetMaxActiveTasks ensures that no value is present for MaxActiveTasks, not even an explicit nil
### GetMaxActiveRuns

`func (o *DAGDetail) GetMaxActiveRuns() int32`

GetMaxActiveRuns returns the MaxActiveRuns field if non-nil, zero value otherwise.

### GetMaxActiveRunsOk

`func (o *DAGDetail) GetMaxActiveRunsOk() (*int32, bool)`

GetMaxActiveRunsOk returns a tuple with the MaxActiveRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxActiveRuns

`func (o *DAGDetail) SetMaxActiveRuns(v int32)`

SetMaxActiveRuns sets MaxActiveRuns field to given value.

### HasMaxActiveRuns

`func (o *DAGDetail) HasMaxActiveRuns() bool`

HasMaxActiveRuns returns a boolean if a field has been set.

### SetMaxActiveRunsNil

`func (o *DAGDetail) SetMaxActiveRunsNil(b bool)`

 SetMaxActiveRunsNil sets the value for MaxActiveRuns to be an explicit nil

### UnsetMaxActiveRuns
`func (o *DAGDetail) UnsetMaxActiveRuns()`

UnsetMaxActiveRuns ensures that no value is present for MaxActiveRuns, not even an explicit nil
### GetHasTaskConcurrencyLimits

`func (o *DAGDetail) GetHasTaskConcurrencyLimits() bool`

GetHasTaskConcurrencyLimits returns the HasTaskConcurrencyLimits field if non-nil, zero value otherwise.

### GetHasTaskConcurrencyLimitsOk

`func (o *DAGDetail) GetHasTaskConcurrencyLimitsOk() (*bool, bool)`

GetHasTaskConcurrencyLimitsOk returns a tuple with the HasTaskConcurrencyLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTaskConcurrencyLimits

`func (o *DAGDetail) SetHasTaskConcurrencyLimits(v bool)`

SetHasTaskConcurrencyLimits sets HasTaskConcurrencyLimits field to given value.

### HasHasTaskConcurrencyLimits

`func (o *DAGDetail) HasHasTaskConcurrencyLimits() bool`

HasHasTaskConcurrencyLimits returns a boolean if a field has been set.

### SetHasTaskConcurrencyLimitsNil

`func (o *DAGDetail) SetHasTaskConcurrencyLimitsNil(b bool)`

 SetHasTaskConcurrencyLimitsNil sets the value for HasTaskConcurrencyLimits to be an explicit nil

### UnsetHasTaskConcurrencyLimits
`func (o *DAGDetail) UnsetHasTaskConcurrencyLimits()`

UnsetHasTaskConcurrencyLimits ensures that no value is present for HasTaskConcurrencyLimits, not even an explicit nil
### GetHasImportErrors

`func (o *DAGDetail) GetHasImportErrors() bool`

GetHasImportErrors returns the HasImportErrors field if non-nil, zero value otherwise.

### GetHasImportErrorsOk

`func (o *DAGDetail) GetHasImportErrorsOk() (*bool, bool)`

GetHasImportErrorsOk returns a tuple with the HasImportErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasImportErrors

`func (o *DAGDetail) SetHasImportErrors(v bool)`

SetHasImportErrors sets HasImportErrors field to given value.

### HasHasImportErrors

`func (o *DAGDetail) HasHasImportErrors() bool`

HasHasImportErrors returns a boolean if a field has been set.

### SetHasImportErrorsNil

`func (o *DAGDetail) SetHasImportErrorsNil(b bool)`

 SetHasImportErrorsNil sets the value for HasImportErrors to be an explicit nil

### UnsetHasImportErrors
`func (o *DAGDetail) UnsetHasImportErrors()`

UnsetHasImportErrors ensures that no value is present for HasImportErrors, not even an explicit nil
### GetNextDagrun

`func (o *DAGDetail) GetNextDagrun() time.Time`

GetNextDagrun returns the NextDagrun field if non-nil, zero value otherwise.

### GetNextDagrunOk

`func (o *DAGDetail) GetNextDagrunOk() (*time.Time, bool)`

GetNextDagrunOk returns a tuple with the NextDagrun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDagrun

`func (o *DAGDetail) SetNextDagrun(v time.Time)`

SetNextDagrun sets NextDagrun field to given value.

### HasNextDagrun

`func (o *DAGDetail) HasNextDagrun() bool`

HasNextDagrun returns a boolean if a field has been set.

### SetNextDagrunNil

`func (o *DAGDetail) SetNextDagrunNil(b bool)`

 SetNextDagrunNil sets the value for NextDagrun to be an explicit nil

### UnsetNextDagrun
`func (o *DAGDetail) UnsetNextDagrun()`

UnsetNextDagrun ensures that no value is present for NextDagrun, not even an explicit nil
### GetNextDagrunDataIntervalStart

`func (o *DAGDetail) GetNextDagrunDataIntervalStart() time.Time`

GetNextDagrunDataIntervalStart returns the NextDagrunDataIntervalStart field if non-nil, zero value otherwise.

### GetNextDagrunDataIntervalStartOk

`func (o *DAGDetail) GetNextDagrunDataIntervalStartOk() (*time.Time, bool)`

GetNextDagrunDataIntervalStartOk returns a tuple with the NextDagrunDataIntervalStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDagrunDataIntervalStart

`func (o *DAGDetail) SetNextDagrunDataIntervalStart(v time.Time)`

SetNextDagrunDataIntervalStart sets NextDagrunDataIntervalStart field to given value.

### HasNextDagrunDataIntervalStart

`func (o *DAGDetail) HasNextDagrunDataIntervalStart() bool`

HasNextDagrunDataIntervalStart returns a boolean if a field has been set.

### SetNextDagrunDataIntervalStartNil

`func (o *DAGDetail) SetNextDagrunDataIntervalStartNil(b bool)`

 SetNextDagrunDataIntervalStartNil sets the value for NextDagrunDataIntervalStart to be an explicit nil

### UnsetNextDagrunDataIntervalStart
`func (o *DAGDetail) UnsetNextDagrunDataIntervalStart()`

UnsetNextDagrunDataIntervalStart ensures that no value is present for NextDagrunDataIntervalStart, not even an explicit nil
### GetNextDagrunDataIntervalEnd

`func (o *DAGDetail) GetNextDagrunDataIntervalEnd() time.Time`

GetNextDagrunDataIntervalEnd returns the NextDagrunDataIntervalEnd field if non-nil, zero value otherwise.

### GetNextDagrunDataIntervalEndOk

`func (o *DAGDetail) GetNextDagrunDataIntervalEndOk() (*time.Time, bool)`

GetNextDagrunDataIntervalEndOk returns a tuple with the NextDagrunDataIntervalEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDagrunDataIntervalEnd

`func (o *DAGDetail) SetNextDagrunDataIntervalEnd(v time.Time)`

SetNextDagrunDataIntervalEnd sets NextDagrunDataIntervalEnd field to given value.

### HasNextDagrunDataIntervalEnd

`func (o *DAGDetail) HasNextDagrunDataIntervalEnd() bool`

HasNextDagrunDataIntervalEnd returns a boolean if a field has been set.

### SetNextDagrunDataIntervalEndNil

`func (o *DAGDetail) SetNextDagrunDataIntervalEndNil(b bool)`

 SetNextDagrunDataIntervalEndNil sets the value for NextDagrunDataIntervalEnd to be an explicit nil

### UnsetNextDagrunDataIntervalEnd
`func (o *DAGDetail) UnsetNextDagrunDataIntervalEnd()`

UnsetNextDagrunDataIntervalEnd ensures that no value is present for NextDagrunDataIntervalEnd, not even an explicit nil
### GetNextDagrunCreateAfter

`func (o *DAGDetail) GetNextDagrunCreateAfter() time.Time`

GetNextDagrunCreateAfter returns the NextDagrunCreateAfter field if non-nil, zero value otherwise.

### GetNextDagrunCreateAfterOk

`func (o *DAGDetail) GetNextDagrunCreateAfterOk() (*time.Time, bool)`

GetNextDagrunCreateAfterOk returns a tuple with the NextDagrunCreateAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDagrunCreateAfter

`func (o *DAGDetail) SetNextDagrunCreateAfter(v time.Time)`

SetNextDagrunCreateAfter sets NextDagrunCreateAfter field to given value.

### HasNextDagrunCreateAfter

`func (o *DAGDetail) HasNextDagrunCreateAfter() bool`

HasNextDagrunCreateAfter returns a boolean if a field has been set.

### SetNextDagrunCreateAfterNil

`func (o *DAGDetail) SetNextDagrunCreateAfterNil(b bool)`

 SetNextDagrunCreateAfterNil sets the value for NextDagrunCreateAfter to be an explicit nil

### UnsetNextDagrunCreateAfter
`func (o *DAGDetail) UnsetNextDagrunCreateAfter()`

UnsetNextDagrunCreateAfter ensures that no value is present for NextDagrunCreateAfter, not even an explicit nil
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

### SetDagRunTimeoutNil

`func (o *DAGDetail) SetDagRunTimeoutNil(b bool)`

 SetDagRunTimeoutNil sets the value for DagRunTimeout to be an explicit nil

### UnsetDagRunTimeout
`func (o *DAGDetail) UnsetDagRunTimeout()`

UnsetDagRunTimeout ensures that no value is present for DagRunTimeout, not even an explicit nil
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

### GetEndDate

`func (o *DAGDetail) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DAGDetail) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DAGDetail) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DAGDetail) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *DAGDetail) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *DAGDetail) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetIsPausedUponCreation

`func (o *DAGDetail) GetIsPausedUponCreation() bool`

GetIsPausedUponCreation returns the IsPausedUponCreation field if non-nil, zero value otherwise.

### GetIsPausedUponCreationOk

`func (o *DAGDetail) GetIsPausedUponCreationOk() (*bool, bool)`

GetIsPausedUponCreationOk returns a tuple with the IsPausedUponCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPausedUponCreation

`func (o *DAGDetail) SetIsPausedUponCreation(v bool)`

SetIsPausedUponCreation sets IsPausedUponCreation field to given value.

### HasIsPausedUponCreation

`func (o *DAGDetail) HasIsPausedUponCreation() bool`

HasIsPausedUponCreation returns a boolean if a field has been set.

### SetIsPausedUponCreationNil

`func (o *DAGDetail) SetIsPausedUponCreationNil(b bool)`

 SetIsPausedUponCreationNil sets the value for IsPausedUponCreation to be an explicit nil

### UnsetIsPausedUponCreation
`func (o *DAGDetail) UnsetIsPausedUponCreation()`

UnsetIsPausedUponCreation ensures that no value is present for IsPausedUponCreation, not even an explicit nil
### GetLastParsed

`func (o *DAGDetail) GetLastParsed() time.Time`

GetLastParsed returns the LastParsed field if non-nil, zero value otherwise.

### GetLastParsedOk

`func (o *DAGDetail) GetLastParsedOk() (*time.Time, bool)`

GetLastParsedOk returns a tuple with the LastParsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastParsed

`func (o *DAGDetail) SetLastParsed(v time.Time)`

SetLastParsed sets LastParsed field to given value.

### HasLastParsed

`func (o *DAGDetail) HasLastParsed() bool`

HasLastParsed returns a boolean if a field has been set.

### SetLastParsedNil

`func (o *DAGDetail) SetLastParsedNil(b bool)`

 SetLastParsedNil sets the value for LastParsed to be an explicit nil

### UnsetLastParsed
`func (o *DAGDetail) UnsetLastParsed()`

UnsetLastParsed ensures that no value is present for LastParsed, not even an explicit nil
### GetTemplateSearchPath

`func (o *DAGDetail) GetTemplateSearchPath() []string`

GetTemplateSearchPath returns the TemplateSearchPath field if non-nil, zero value otherwise.

### GetTemplateSearchPathOk

`func (o *DAGDetail) GetTemplateSearchPathOk() (*[]string, bool)`

GetTemplateSearchPathOk returns a tuple with the TemplateSearchPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSearchPath

`func (o *DAGDetail) SetTemplateSearchPath(v []string)`

SetTemplateSearchPath sets TemplateSearchPath field to given value.

### HasTemplateSearchPath

`func (o *DAGDetail) HasTemplateSearchPath() bool`

HasTemplateSearchPath returns a boolean if a field has been set.

### SetTemplateSearchPathNil

`func (o *DAGDetail) SetTemplateSearchPathNil(b bool)`

 SetTemplateSearchPathNil sets the value for TemplateSearchPath to be an explicit nil

### UnsetTemplateSearchPath
`func (o *DAGDetail) UnsetTemplateSearchPath()`

UnsetTemplateSearchPath ensures that no value is present for TemplateSearchPath, not even an explicit nil
### GetRenderTemplateAsNativeObj

`func (o *DAGDetail) GetRenderTemplateAsNativeObj() bool`

GetRenderTemplateAsNativeObj returns the RenderTemplateAsNativeObj field if non-nil, zero value otherwise.

### GetRenderTemplateAsNativeObjOk

`func (o *DAGDetail) GetRenderTemplateAsNativeObjOk() (*bool, bool)`

GetRenderTemplateAsNativeObjOk returns a tuple with the RenderTemplateAsNativeObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderTemplateAsNativeObj

`func (o *DAGDetail) SetRenderTemplateAsNativeObj(v bool)`

SetRenderTemplateAsNativeObj sets RenderTemplateAsNativeObj field to given value.

### HasRenderTemplateAsNativeObj

`func (o *DAGDetail) HasRenderTemplateAsNativeObj() bool`

HasRenderTemplateAsNativeObj returns a boolean if a field has been set.

### SetRenderTemplateAsNativeObjNil

`func (o *DAGDetail) SetRenderTemplateAsNativeObjNil(b bool)`

 SetRenderTemplateAsNativeObjNil sets the value for RenderTemplateAsNativeObj to be an explicit nil

### UnsetRenderTemplateAsNativeObj
`func (o *DAGDetail) UnsetRenderTemplateAsNativeObj()`

UnsetRenderTemplateAsNativeObj ensures that no value is present for RenderTemplateAsNativeObj, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


