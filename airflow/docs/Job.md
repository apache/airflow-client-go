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

# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**DagId** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to **NullableString** |  | [optional] 
**JobType** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **NullableString** |  | [optional] 
**EndDate** | Pointer to **NullableString** |  | [optional] 
**LatestHeartbeat** | Pointer to **NullableString** |  | [optional] 
**ExecutorClass** | Pointer to **NullableString** |  | [optional] 
**Hostname** | Pointer to **NullableString** |  | [optional] 
**Unixname** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewJob

`func NewJob() *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Job) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Job) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Job) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Job) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDagId

`func (o *Job) GetDagId() string`

GetDagId returns the DagId field if non-nil, zero value otherwise.

### GetDagIdOk

`func (o *Job) GetDagIdOk() (*string, bool)`

GetDagIdOk returns a tuple with the DagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagId

`func (o *Job) SetDagId(v string)`

SetDagId sets DagId field to given value.

### HasDagId

`func (o *Job) HasDagId() bool`

HasDagId returns a boolean if a field has been set.

### SetDagIdNil

`func (o *Job) SetDagIdNil(b bool)`

 SetDagIdNil sets the value for DagId to be an explicit nil

### UnsetDagId
`func (o *Job) UnsetDagId()`

UnsetDagId ensures that no value is present for DagId, not even an explicit nil
### GetState

`func (o *Job) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Job) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Job) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Job) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *Job) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *Job) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetJobType

`func (o *Job) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *Job) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *Job) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *Job) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### SetJobTypeNil

`func (o *Job) SetJobTypeNil(b bool)`

 SetJobTypeNil sets the value for JobType to be an explicit nil

### UnsetJobType
`func (o *Job) UnsetJobType()`

UnsetJobType ensures that no value is present for JobType, not even an explicit nil
### GetStartDate

`func (o *Job) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Job) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Job) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Job) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *Job) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *Job) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *Job) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Job) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Job) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Job) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *Job) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *Job) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetLatestHeartbeat

`func (o *Job) GetLatestHeartbeat() string`

GetLatestHeartbeat returns the LatestHeartbeat field if non-nil, zero value otherwise.

### GetLatestHeartbeatOk

`func (o *Job) GetLatestHeartbeatOk() (*string, bool)`

GetLatestHeartbeatOk returns a tuple with the LatestHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestHeartbeat

`func (o *Job) SetLatestHeartbeat(v string)`

SetLatestHeartbeat sets LatestHeartbeat field to given value.

### HasLatestHeartbeat

`func (o *Job) HasLatestHeartbeat() bool`

HasLatestHeartbeat returns a boolean if a field has been set.

### SetLatestHeartbeatNil

`func (o *Job) SetLatestHeartbeatNil(b bool)`

 SetLatestHeartbeatNil sets the value for LatestHeartbeat to be an explicit nil

### UnsetLatestHeartbeat
`func (o *Job) UnsetLatestHeartbeat()`

UnsetLatestHeartbeat ensures that no value is present for LatestHeartbeat, not even an explicit nil
### GetExecutorClass

`func (o *Job) GetExecutorClass() string`

GetExecutorClass returns the ExecutorClass field if non-nil, zero value otherwise.

### GetExecutorClassOk

`func (o *Job) GetExecutorClassOk() (*string, bool)`

GetExecutorClassOk returns a tuple with the ExecutorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutorClass

`func (o *Job) SetExecutorClass(v string)`

SetExecutorClass sets ExecutorClass field to given value.

### HasExecutorClass

`func (o *Job) HasExecutorClass() bool`

HasExecutorClass returns a boolean if a field has been set.

### SetExecutorClassNil

`func (o *Job) SetExecutorClassNil(b bool)`

 SetExecutorClassNil sets the value for ExecutorClass to be an explicit nil

### UnsetExecutorClass
`func (o *Job) UnsetExecutorClass()`

UnsetExecutorClass ensures that no value is present for ExecutorClass, not even an explicit nil
### GetHostname

`func (o *Job) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Job) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Job) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *Job) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *Job) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *Job) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetUnixname

`func (o *Job) GetUnixname() string`

GetUnixname returns the Unixname field if non-nil, zero value otherwise.

### GetUnixnameOk

`func (o *Job) GetUnixnameOk() (*string, bool)`

GetUnixnameOk returns a tuple with the Unixname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixname

`func (o *Job) SetUnixname(v string)`

SetUnixname sets Unixname field to given value.

### HasUnixname

`func (o *Job) HasUnixname() bool`

HasUnixname returns a boolean if a field has been set.

### SetUnixnameNil

`func (o *Job) SetUnixnameNil(b bool)`

 SetUnixnameNil sets the value for Unixname to be an explicit nil

### UnsetUnixname
`func (o *Job) UnsetUnixname()`

UnsetUnixname ensures that no value is present for Unixname, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


