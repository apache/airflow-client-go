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

# EventLogCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventLogs** | Pointer to [**[]EventLog**](EventLog.md) |  | [optional] 
**TotalEntries** | Pointer to **int32** | Count of total objects in the current result set before pagination parameters (limit, offset) are applied.  | [optional] 

## Methods

### NewEventLogCollection

`func NewEventLogCollection() *EventLogCollection`

NewEventLogCollection instantiates a new EventLogCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventLogCollectionWithDefaults

`func NewEventLogCollectionWithDefaults() *EventLogCollection`

NewEventLogCollectionWithDefaults instantiates a new EventLogCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventLogs

`func (o *EventLogCollection) GetEventLogs() []EventLog`

GetEventLogs returns the EventLogs field if non-nil, zero value otherwise.

### GetEventLogsOk

`func (o *EventLogCollection) GetEventLogsOk() (*[]EventLog, bool)`

GetEventLogsOk returns a tuple with the EventLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLogs

`func (o *EventLogCollection) SetEventLogs(v []EventLog)`

SetEventLogs sets EventLogs field to given value.

### HasEventLogs

`func (o *EventLogCollection) HasEventLogs() bool`

HasEventLogs returns a boolean if a field has been set.

### GetTotalEntries

`func (o *EventLogCollection) GetTotalEntries() int32`

GetTotalEntries returns the TotalEntries field if non-nil, zero value otherwise.

### GetTotalEntriesOk

`func (o *EventLogCollection) GetTotalEntriesOk() (*int32, bool)`

GetTotalEntriesOk returns a tuple with the TotalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntries

`func (o *EventLogCollection) SetTotalEntries(v int32)`

SetTotalEntries sets TotalEntries field to given value.

### HasTotalEntries

`func (o *EventLogCollection) HasTotalEntries() bool`

HasTotalEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


