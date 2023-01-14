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

# DagWarning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DagId** | Pointer to **string** | The dag_id. | [optional] [readonly] 
**WarningType** | Pointer to **string** | The warning type for the dag warning. | [optional] [readonly] 
**Message** | Pointer to **string** | The message for the dag warning. | [optional] [readonly] 
**Timestamp** | Pointer to **string** | The time when this warning was logged. | [optional] [readonly] 

## Methods

### NewDagWarning

`func NewDagWarning() *DagWarning`

NewDagWarning instantiates a new DagWarning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDagWarningWithDefaults

`func NewDagWarningWithDefaults() *DagWarning`

NewDagWarningWithDefaults instantiates a new DagWarning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDagId

`func (o *DagWarning) GetDagId() string`

GetDagId returns the DagId field if non-nil, zero value otherwise.

### GetDagIdOk

`func (o *DagWarning) GetDagIdOk() (*string, bool)`

GetDagIdOk returns a tuple with the DagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagId

`func (o *DagWarning) SetDagId(v string)`

SetDagId sets DagId field to given value.

### HasDagId

`func (o *DagWarning) HasDagId() bool`

HasDagId returns a boolean if a field has been set.

### GetWarningType

`func (o *DagWarning) GetWarningType() string`

GetWarningType returns the WarningType field if non-nil, zero value otherwise.

### GetWarningTypeOk

`func (o *DagWarning) GetWarningTypeOk() (*string, bool)`

GetWarningTypeOk returns a tuple with the WarningType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningType

`func (o *DagWarning) SetWarningType(v string)`

SetWarningType sets WarningType field to given value.

### HasWarningType

`func (o *DagWarning) HasWarningType() bool`

HasWarningType returns a boolean if a field has been set.

### GetMessage

`func (o *DagWarning) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DagWarning) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DagWarning) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DagWarning) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTimestamp

`func (o *DagWarning) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DagWarning) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DagWarning) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DagWarning) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


