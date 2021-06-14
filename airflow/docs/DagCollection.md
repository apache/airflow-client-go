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

# DAGCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dags** | Pointer to [**[]DAG**](DAG.md) |  | [optional] 
**TotalEntries** | Pointer to **int32** | Count of objects in the current result set. | [optional] 

## Methods

### NewDAGCollection

`func NewDAGCollection() *DAGCollection`

NewDAGCollection instantiates a new DAGCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDAGCollectionWithDefaults

`func NewDAGCollectionWithDefaults() *DAGCollection`

NewDAGCollectionWithDefaults instantiates a new DAGCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDags

`func (o *DAGCollection) GetDags() []DAG`

GetDags returns the Dags field if non-nil, zero value otherwise.

### GetDagsOk

`func (o *DAGCollection) GetDagsOk() (*[]DAG, bool)`

GetDagsOk returns a tuple with the Dags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDags

`func (o *DAGCollection) SetDags(v []DAG)`

SetDags sets Dags field to given value.

### HasDags

`func (o *DAGCollection) HasDags() bool`

HasDags returns a boolean if a field has been set.

### GetTotalEntries

`func (o *DAGCollection) GetTotalEntries() int32`

GetTotalEntries returns the TotalEntries field if non-nil, zero value otherwise.

### GetTotalEntriesOk

`func (o *DAGCollection) GetTotalEntriesOk() (*int32, bool)`

GetTotalEntriesOk returns a tuple with the TotalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntries

`func (o *DAGCollection) SetTotalEntries(v int32)`

SetTotalEntries sets TotalEntries field to given value.

### HasTotalEntries

`func (o *DAGCollection) HasTotalEntries() bool`

HasTotalEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


