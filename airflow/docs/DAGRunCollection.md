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

# DAGRunCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DagRuns** | Pointer to [**[]DAGRun**](DAGRun.md) |  | [optional] 
**TotalEntries** | Pointer to **int32** | Count of total objects in the current result set before pagination parameters (limit, offset) are applied.  | [optional] 

## Methods

### NewDAGRunCollection

`func NewDAGRunCollection() *DAGRunCollection`

NewDAGRunCollection instantiates a new DAGRunCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDAGRunCollectionWithDefaults

`func NewDAGRunCollectionWithDefaults() *DAGRunCollection`

NewDAGRunCollectionWithDefaults instantiates a new DAGRunCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDagRuns

`func (o *DAGRunCollection) GetDagRuns() []DAGRun`

GetDagRuns returns the DagRuns field if non-nil, zero value otherwise.

### GetDagRunsOk

`func (o *DAGRunCollection) GetDagRunsOk() (*[]DAGRun, bool)`

GetDagRunsOk returns a tuple with the DagRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagRuns

`func (o *DAGRunCollection) SetDagRuns(v []DAGRun)`

SetDagRuns sets DagRuns field to given value.

### HasDagRuns

`func (o *DAGRunCollection) HasDagRuns() bool`

HasDagRuns returns a boolean if a field has been set.

### GetTotalEntries

`func (o *DAGRunCollection) GetTotalEntries() int32`

GetTotalEntries returns the TotalEntries field if non-nil, zero value otherwise.

### GetTotalEntriesOk

`func (o *DAGRunCollection) GetTotalEntriesOk() (*int32, bool)`

GetTotalEntriesOk returns a tuple with the TotalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntries

`func (o *DAGRunCollection) SetTotalEntries(v int32)`

SetTotalEntries sets TotalEntries field to given value.

### HasTotalEntries

`func (o *DAGRunCollection) HasTotalEntries() bool`

HasTotalEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


