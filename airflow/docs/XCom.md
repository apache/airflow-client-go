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

# XCom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**ExecutionDate** | Pointer to **string** |  | [optional] 
**MapIndex** | Pointer to **int32** |  | [optional] 
**TaskId** | Pointer to **string** |  | [optional] 
**DagId** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** | The value | [optional] 

## Methods

### NewXCom

`func NewXCom() *XCom`

NewXCom instantiates a new XCom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXComWithDefaults

`func NewXComWithDefaults() *XCom`

NewXComWithDefaults instantiates a new XCom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *XCom) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *XCom) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *XCom) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *XCom) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetTimestamp

`func (o *XCom) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *XCom) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *XCom) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *XCom) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetExecutionDate

`func (o *XCom) GetExecutionDate() string`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *XCom) GetExecutionDateOk() (*string, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *XCom) SetExecutionDate(v string)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *XCom) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetMapIndex

`func (o *XCom) GetMapIndex() int32`

GetMapIndex returns the MapIndex field if non-nil, zero value otherwise.

### GetMapIndexOk

`func (o *XCom) GetMapIndexOk() (*int32, bool)`

GetMapIndexOk returns a tuple with the MapIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapIndex

`func (o *XCom) SetMapIndex(v int32)`

SetMapIndex sets MapIndex field to given value.

### HasMapIndex

`func (o *XCom) HasMapIndex() bool`

HasMapIndex returns a boolean if a field has been set.

### GetTaskId

`func (o *XCom) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *XCom) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *XCom) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *XCom) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetDagId

`func (o *XCom) GetDagId() string`

GetDagId returns the DagId field if non-nil, zero value otherwise.

### GetDagIdOk

`func (o *XCom) GetDagIdOk() (*string, bool)`

GetDagIdOk returns a tuple with the DagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagId

`func (o *XCom) SetDagId(v string)`

SetDagId sets DagId field to given value.

### HasDagId

`func (o *XCom) HasDagId() bool`

HasDagId returns a boolean if a field has been set.

### GetValue

`func (o *XCom) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *XCom) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *XCom) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *XCom) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


