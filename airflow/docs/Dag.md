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
**LastParsedTime** | Pointer to **NullableTime** | The last time the DAG was parsed.  *New in version 2.3.0*  | [optional] [readonly] 
**LastPickled** | Pointer to **NullableTime** | The last time the DAG was pickled.  *New in version 2.3.0*  | [optional] [readonly] 
**LastExpired** | Pointer to **NullableTime** | Time when the DAG last received a refresh signal (e.g. the DAG&#39;s \&quot;refresh\&quot; button was clicked in the web UI)  *New in version 2.3.0*  | [optional] [readonly] 
**SchedulerLock** | Pointer to **NullableBool** | Whether (one of) the scheduler is scheduling this DAG at the moment  *New in version 2.3.0*  | [optional] [readonly] 
**PickleId** | Pointer to **NullableString** | Foreign key to the latest pickle_id  *New in version 2.3.0*  | [optional] [readonly] 
**DefaultView** | Pointer to **NullableString** | Default view of the DAG inside the webserver  *New in version 2.3.0*  | [optional] [readonly] 
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

### GetLastParsedTime

`func (o *DAG) GetLastParsedTime() time.Time`

GetLastParsedTime returns the LastParsedTime field if non-nil, zero value otherwise.

### GetLastParsedTimeOk

`func (o *DAG) GetLastParsedTimeOk() (*time.Time, bool)`

GetLastParsedTimeOk returns a tuple with the LastParsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastParsedTime

`func (o *DAG) SetLastParsedTime(v time.Time)`

SetLastParsedTime sets LastParsedTime field to given value.

### HasLastParsedTime

`func (o *DAG) HasLastParsedTime() bool`

HasLastParsedTime returns a boolean if a field has been set.

### SetLastParsedTimeNil

`func (o *DAG) SetLastParsedTimeNil(b bool)`

 SetLastParsedTimeNil sets the value for LastParsedTime to be an explicit nil

### UnsetLastParsedTime
`func (o *DAG) UnsetLastParsedTime()`

UnsetLastParsedTime ensures that no value is present for LastParsedTime, not even an explicit nil
### GetLastPickled

`func (o *DAG) GetLastPickled() time.Time`

GetLastPickled returns the LastPickled field if non-nil, zero value otherwise.

### GetLastPickledOk

`func (o *DAG) GetLastPickledOk() (*time.Time, bool)`

GetLastPickledOk returns a tuple with the LastPickled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPickled

`func (o *DAG) SetLastPickled(v time.Time)`

SetLastPickled sets LastPickled field to given value.

### HasLastPickled

`func (o *DAG) HasLastPickled() bool`

HasLastPickled returns a boolean if a field has been set.

### SetLastPickledNil

`func (o *DAG) SetLastPickledNil(b bool)`

 SetLastPickledNil sets the value for LastPickled to be an explicit nil

### UnsetLastPickled
`func (o *DAG) UnsetLastPickled()`

UnsetLastPickled ensures that no value is present for LastPickled, not even an explicit nil
### GetLastExpired

`func (o *DAG) GetLastExpired() time.Time`

GetLastExpired returns the LastExpired field if non-nil, zero value otherwise.

### GetLastExpiredOk

`func (o *DAG) GetLastExpiredOk() (*time.Time, bool)`

GetLastExpiredOk returns a tuple with the LastExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExpired

`func (o *DAG) SetLastExpired(v time.Time)`

SetLastExpired sets LastExpired field to given value.

### HasLastExpired

`func (o *DAG) HasLastExpired() bool`

HasLastExpired returns a boolean if a field has been set.

### SetLastExpiredNil

`func (o *DAG) SetLastExpiredNil(b bool)`

 SetLastExpiredNil sets the value for LastExpired to be an explicit nil

### UnsetLastExpired
`func (o *DAG) UnsetLastExpired()`

UnsetLastExpired ensures that no value is present for LastExpired, not even an explicit nil
### GetSchedulerLock

`func (o *DAG) GetSchedulerLock() bool`

GetSchedulerLock returns the SchedulerLock field if non-nil, zero value otherwise.

### GetSchedulerLockOk

`func (o *DAG) GetSchedulerLockOk() (*bool, bool)`

GetSchedulerLockOk returns a tuple with the SchedulerLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerLock

`func (o *DAG) SetSchedulerLock(v bool)`

SetSchedulerLock sets SchedulerLock field to given value.

### HasSchedulerLock

`func (o *DAG) HasSchedulerLock() bool`

HasSchedulerLock returns a boolean if a field has been set.

### SetSchedulerLockNil

`func (o *DAG) SetSchedulerLockNil(b bool)`

 SetSchedulerLockNil sets the value for SchedulerLock to be an explicit nil

### UnsetSchedulerLock
`func (o *DAG) UnsetSchedulerLock()`

UnsetSchedulerLock ensures that no value is present for SchedulerLock, not even an explicit nil
### GetPickleId

`func (o *DAG) GetPickleId() string`

GetPickleId returns the PickleId field if non-nil, zero value otherwise.

### GetPickleIdOk

`func (o *DAG) GetPickleIdOk() (*string, bool)`

GetPickleIdOk returns a tuple with the PickleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickleId

`func (o *DAG) SetPickleId(v string)`

SetPickleId sets PickleId field to given value.

### HasPickleId

`func (o *DAG) HasPickleId() bool`

HasPickleId returns a boolean if a field has been set.

### SetPickleIdNil

`func (o *DAG) SetPickleIdNil(b bool)`

 SetPickleIdNil sets the value for PickleId to be an explicit nil

### UnsetPickleId
`func (o *DAG) UnsetPickleId()`

UnsetPickleId ensures that no value is present for PickleId, not even an explicit nil
### GetDefaultView

`func (o *DAG) GetDefaultView() string`

GetDefaultView returns the DefaultView field if non-nil, zero value otherwise.

### GetDefaultViewOk

`func (o *DAG) GetDefaultViewOk() (*string, bool)`

GetDefaultViewOk returns a tuple with the DefaultView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultView

`func (o *DAG) SetDefaultView(v string)`

SetDefaultView sets DefaultView field to given value.

### HasDefaultView

`func (o *DAG) HasDefaultView() bool`

HasDefaultView returns a boolean if a field has been set.

### SetDefaultViewNil

`func (o *DAG) SetDefaultViewNil(b bool)`

 SetDefaultViewNil sets the value for DefaultView to be an explicit nil

### UnsetDefaultView
`func (o *DAG) UnsetDefaultView()`

UnsetDefaultView ensures that no value is present for DefaultView, not even an explicit nil
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

### SetScheduleIntervalNil

`func (o *DAG) SetScheduleIntervalNil(b bool)`

 SetScheduleIntervalNil sets the value for ScheduleInterval to be an explicit nil

### UnsetScheduleInterval
`func (o *DAG) UnsetScheduleInterval()`

UnsetScheduleInterval ensures that no value is present for ScheduleInterval, not even an explicit nil
### GetTimetableDescription

`func (o *DAG) GetTimetableDescription() string`

GetTimetableDescription returns the TimetableDescription field if non-nil, zero value otherwise.

### GetTimetableDescriptionOk

`func (o *DAG) GetTimetableDescriptionOk() (*string, bool)`

GetTimetableDescriptionOk returns a tuple with the TimetableDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimetableDescription

`func (o *DAG) SetTimetableDescription(v string)`

SetTimetableDescription sets TimetableDescription field to given value.

### HasTimetableDescription

`func (o *DAG) HasTimetableDescription() bool`

HasTimetableDescription returns a boolean if a field has been set.

### SetTimetableDescriptionNil

`func (o *DAG) SetTimetableDescriptionNil(b bool)`

 SetTimetableDescriptionNil sets the value for TimetableDescription to be an explicit nil

### UnsetTimetableDescription
`func (o *DAG) UnsetTimetableDescription()`

UnsetTimetableDescription ensures that no value is present for TimetableDescription, not even an explicit nil
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
### GetMaxActiveTasks

`func (o *DAG) GetMaxActiveTasks() int32`

GetMaxActiveTasks returns the MaxActiveTasks field if non-nil, zero value otherwise.

### GetMaxActiveTasksOk

`func (o *DAG) GetMaxActiveTasksOk() (*int32, bool)`

GetMaxActiveTasksOk returns a tuple with the MaxActiveTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxActiveTasks

`func (o *DAG) SetMaxActiveTasks(v int32)`

SetMaxActiveTasks sets MaxActiveTasks field to given value.

### HasMaxActiveTasks

`func (o *DAG) HasMaxActiveTasks() bool`

HasMaxActiveTasks returns a boolean if a field has been set.

### SetMaxActiveTasksNil

`func (o *DAG) SetMaxActiveTasksNil(b bool)`

 SetMaxActiveTasksNil sets the value for MaxActiveTasks to be an explicit nil

### UnsetMaxActiveTasks
`func (o *DAG) UnsetMaxActiveTasks()`

UnsetMaxActiveTasks ensures that no value is present for MaxActiveTasks, not even an explicit nil
### GetMaxActiveRuns

`func (o *DAG) GetMaxActiveRuns() int32`

GetMaxActiveRuns returns the MaxActiveRuns field if non-nil, zero value otherwise.

### GetMaxActiveRunsOk

`func (o *DAG) GetMaxActiveRunsOk() (*int32, bool)`

GetMaxActiveRunsOk returns a tuple with the MaxActiveRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxActiveRuns

`func (o *DAG) SetMaxActiveRuns(v int32)`

SetMaxActiveRuns sets MaxActiveRuns field to given value.

### HasMaxActiveRuns

`func (o *DAG) HasMaxActiveRuns() bool`

HasMaxActiveRuns returns a boolean if a field has been set.

### SetMaxActiveRunsNil

`func (o *DAG) SetMaxActiveRunsNil(b bool)`

 SetMaxActiveRunsNil sets the value for MaxActiveRuns to be an explicit nil

### UnsetMaxActiveRuns
`func (o *DAG) UnsetMaxActiveRuns()`

UnsetMaxActiveRuns ensures that no value is present for MaxActiveRuns, not even an explicit nil
### GetHasTaskConcurrencyLimits

`func (o *DAG) GetHasTaskConcurrencyLimits() bool`

GetHasTaskConcurrencyLimits returns the HasTaskConcurrencyLimits field if non-nil, zero value otherwise.

### GetHasTaskConcurrencyLimitsOk

`func (o *DAG) GetHasTaskConcurrencyLimitsOk() (*bool, bool)`

GetHasTaskConcurrencyLimitsOk returns a tuple with the HasTaskConcurrencyLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTaskConcurrencyLimits

`func (o *DAG) SetHasTaskConcurrencyLimits(v bool)`

SetHasTaskConcurrencyLimits sets HasTaskConcurrencyLimits field to given value.

### HasHasTaskConcurrencyLimits

`func (o *DAG) HasHasTaskConcurrencyLimits() bool`

HasHasTaskConcurrencyLimits returns a boolean if a field has been set.

### SetHasTaskConcurrencyLimitsNil

`func (o *DAG) SetHasTaskConcurrencyLimitsNil(b bool)`

 SetHasTaskConcurrencyLimitsNil sets the value for HasTaskConcurrencyLimits to be an explicit nil

### UnsetHasTaskConcurrencyLimits
`func (o *DAG) UnsetHasTaskConcurrencyLimits()`

UnsetHasTaskConcurrencyLimits ensures that no value is present for HasTaskConcurrencyLimits, not even an explicit nil
### GetHasImportErrors

`func (o *DAG) GetHasImportErrors() bool`

GetHasImportErrors returns the HasImportErrors field if non-nil, zero value otherwise.

### GetHasImportErrorsOk

`func (o *DAG) GetHasImportErrorsOk() (*bool, bool)`

GetHasImportErrorsOk returns a tuple with the HasImportErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasImportErrors

`func (o *DAG) SetHasImportErrors(v bool)`

SetHasImportErrors sets HasImportErrors field to given value.

### HasHasImportErrors

`func (o *DAG) HasHasImportErrors() bool`

HasHasImportErrors returns a boolean if a field has been set.

### SetHasImportErrorsNil

`func (o *DAG) SetHasImportErrorsNil(b bool)`

 SetHasImportErrorsNil sets the value for HasImportErrors to be an explicit nil

### UnsetHasImportErrors
`func (o *DAG) UnsetHasImportErrors()`

UnsetHasImportErrors ensures that no value is present for HasImportErrors, not even an explicit nil
### GetNextDagrun

`func (o *DAG) GetNextDagrun() time.Time`

GetNextDagrun returns the NextDagrun field if non-nil, zero value otherwise.

### GetNextDagrunOk

`func (o *DAG) GetNextDagrunOk() (*time.Time, bool)`

GetNextDagrunOk returns a tuple with the NextDagrun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDagrun

`func (o *DAG) SetNextDagrun(v time.Time)`

SetNextDagrun sets NextDagrun field to given value.

### HasNextDagrun

`func (o *DAG) HasNextDagrun() bool`

HasNextDagrun returns a boolean if a field has been set.

### SetNextDagrunNil

`func (o *DAG) SetNextDagrunNil(b bool)`

 SetNextDagrunNil sets the value for NextDagrun to be an explicit nil

### UnsetNextDagrun
`func (o *DAG) UnsetNextDagrun()`

UnsetNextDagrun ensures that no value is present for NextDagrun, not even an explicit nil
### GetNextDagrunDataIntervalStart

`func (o *DAG) GetNextDagrunDataIntervalStart() time.Time`

GetNextDagrunDataIntervalStart returns the NextDagrunDataIntervalStart field if non-nil, zero value otherwise.

### GetNextDagrunDataIntervalStartOk

`func (o *DAG) GetNextDagrunDataIntervalStartOk() (*time.Time, bool)`

GetNextDagrunDataIntervalStartOk returns a tuple with the NextDagrunDataIntervalStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDagrunDataIntervalStart

`func (o *DAG) SetNextDagrunDataIntervalStart(v time.Time)`

SetNextDagrunDataIntervalStart sets NextDagrunDataIntervalStart field to given value.

### HasNextDagrunDataIntervalStart

`func (o *DAG) HasNextDagrunDataIntervalStart() bool`

HasNextDagrunDataIntervalStart returns a boolean if a field has been set.

### SetNextDagrunDataIntervalStartNil

`func (o *DAG) SetNextDagrunDataIntervalStartNil(b bool)`

 SetNextDagrunDataIntervalStartNil sets the value for NextDagrunDataIntervalStart to be an explicit nil

### UnsetNextDagrunDataIntervalStart
`func (o *DAG) UnsetNextDagrunDataIntervalStart()`

UnsetNextDagrunDataIntervalStart ensures that no value is present for NextDagrunDataIntervalStart, not even an explicit nil
### GetNextDagrunDataIntervalEnd

`func (o *DAG) GetNextDagrunDataIntervalEnd() time.Time`

GetNextDagrunDataIntervalEnd returns the NextDagrunDataIntervalEnd field if non-nil, zero value otherwise.

### GetNextDagrunDataIntervalEndOk

`func (o *DAG) GetNextDagrunDataIntervalEndOk() (*time.Time, bool)`

GetNextDagrunDataIntervalEndOk returns a tuple with the NextDagrunDataIntervalEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDagrunDataIntervalEnd

`func (o *DAG) SetNextDagrunDataIntervalEnd(v time.Time)`

SetNextDagrunDataIntervalEnd sets NextDagrunDataIntervalEnd field to given value.

### HasNextDagrunDataIntervalEnd

`func (o *DAG) HasNextDagrunDataIntervalEnd() bool`

HasNextDagrunDataIntervalEnd returns a boolean if a field has been set.

### SetNextDagrunDataIntervalEndNil

`func (o *DAG) SetNextDagrunDataIntervalEndNil(b bool)`

 SetNextDagrunDataIntervalEndNil sets the value for NextDagrunDataIntervalEnd to be an explicit nil

### UnsetNextDagrunDataIntervalEnd
`func (o *DAG) UnsetNextDagrunDataIntervalEnd()`

UnsetNextDagrunDataIntervalEnd ensures that no value is present for NextDagrunDataIntervalEnd, not even an explicit nil
### GetNextDagrunCreateAfter

`func (o *DAG) GetNextDagrunCreateAfter() time.Time`

GetNextDagrunCreateAfter returns the NextDagrunCreateAfter field if non-nil, zero value otherwise.

### GetNextDagrunCreateAfterOk

`func (o *DAG) GetNextDagrunCreateAfterOk() (*time.Time, bool)`

GetNextDagrunCreateAfterOk returns a tuple with the NextDagrunCreateAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDagrunCreateAfter

`func (o *DAG) SetNextDagrunCreateAfter(v time.Time)`

SetNextDagrunCreateAfter sets NextDagrunCreateAfter field to given value.

### HasNextDagrunCreateAfter

`func (o *DAG) HasNextDagrunCreateAfter() bool`

HasNextDagrunCreateAfter returns a boolean if a field has been set.

### SetNextDagrunCreateAfterNil

`func (o *DAG) SetNextDagrunCreateAfterNil(b bool)`

 SetNextDagrunCreateAfterNil sets the value for NextDagrunCreateAfter to be an explicit nil

### UnsetNextDagrunCreateAfter
`func (o *DAG) UnsetNextDagrunCreateAfter()`

UnsetNextDagrunCreateAfter ensures that no value is present for NextDagrunCreateAfter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


