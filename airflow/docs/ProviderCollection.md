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

# ProviderCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Providers** | Pointer to [**[]Provider**](Provider.md) |  | [optional] 

## Methods

### NewProviderCollection

`func NewProviderCollection() *ProviderCollection`

NewProviderCollection instantiates a new ProviderCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderCollectionWithDefaults

`func NewProviderCollectionWithDefaults() *ProviderCollection`

NewProviderCollectionWithDefaults instantiates a new ProviderCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviders

`func (o *ProviderCollection) GetProviders() []Provider`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *ProviderCollection) GetProvidersOk() (*[]Provider, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *ProviderCollection) SetProviders(v []Provider)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *ProviderCollection) HasProviders() bool`

HasProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


