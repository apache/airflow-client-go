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

# UserCollectionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]UserCollectionItem**](UserCollectionItem.md) |  | [optional] 

## Methods

### NewUserCollectionAllOf

`func NewUserCollectionAllOf() *UserCollectionAllOf`

NewUserCollectionAllOf instantiates a new UserCollectionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCollectionAllOfWithDefaults

`func NewUserCollectionAllOfWithDefaults() *UserCollectionAllOf`

NewUserCollectionAllOfWithDefaults instantiates a new UserCollectionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *UserCollectionAllOf) GetUsers() []UserCollectionItem`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UserCollectionAllOf) GetUsersOk() (*[]UserCollectionItem, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UserCollectionAllOf) SetUsers(v []UserCollectionItem)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UserCollectionAllOf) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


