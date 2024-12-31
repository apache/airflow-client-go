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

# SchedulerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**NullableHealthStatus**](HealthStatus.md) |  | [optional] 
**LatestSchedulerHeartbeat** | Pointer to **NullableString** | The time the scheduler last did a heartbeat. | [optional] [readonly] 

## Methods

### NewSchedulerStatus

`func NewSchedulerStatus() *SchedulerStatus`

NewSchedulerStatus instantiates a new SchedulerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerStatusWithDefaults

`func NewSchedulerStatusWithDefaults() *SchedulerStatus`

NewSchedulerStatusWithDefaults instantiates a new SchedulerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SchedulerStatus) GetStatus() HealthStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SchedulerStatus) GetStatusOk() (*HealthStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SchedulerStatus) SetStatus(v HealthStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SchedulerStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SchedulerStatus) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SchedulerStatus) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetLatestSchedulerHeartbeat

`func (o *SchedulerStatus) GetLatestSchedulerHeartbeat() string`

GetLatestSchedulerHeartbeat returns the LatestSchedulerHeartbeat field if non-nil, zero value otherwise.

### GetLatestSchedulerHeartbeatOk

`func (o *SchedulerStatus) GetLatestSchedulerHeartbeatOk() (*string, bool)`

GetLatestSchedulerHeartbeatOk returns a tuple with the LatestSchedulerHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSchedulerHeartbeat

`func (o *SchedulerStatus) SetLatestSchedulerHeartbeat(v string)`

SetLatestSchedulerHeartbeat sets LatestSchedulerHeartbeat field to given value.

### HasLatestSchedulerHeartbeat

`func (o *SchedulerStatus) HasLatestSchedulerHeartbeat() bool`

HasLatestSchedulerHeartbeat returns a boolean if a field has been set.

### SetLatestSchedulerHeartbeatNil

`func (o *SchedulerStatus) SetLatestSchedulerHeartbeatNil(b bool)`

 SetLatestSchedulerHeartbeatNil sets the value for LatestSchedulerHeartbeat to be an explicit nil

### UnsetLatestSchedulerHeartbeat
`func (o *SchedulerStatus) UnsetLatestSchedulerHeartbeat()`

UnsetLatestSchedulerHeartbeat ensures that no value is present for LatestSchedulerHeartbeat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


