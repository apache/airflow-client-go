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

# ClassReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModulePath** | Pointer to **string** |  | [optional] [readonly] 
**ClassName** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewClassReference

`func NewClassReference() *ClassReference`

NewClassReference instantiates a new ClassReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassReferenceWithDefaults

`func NewClassReferenceWithDefaults() *ClassReference`

NewClassReferenceWithDefaults instantiates a new ClassReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModulePath

`func (o *ClassReference) GetModulePath() string`

GetModulePath returns the ModulePath field if non-nil, zero value otherwise.

### GetModulePathOk

`func (o *ClassReference) GetModulePathOk() (*string, bool)`

GetModulePathOk returns a tuple with the ModulePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModulePath

`func (o *ClassReference) SetModulePath(v string)`

SetModulePath sets ModulePath field to given value.

### HasModulePath

`func (o *ClassReference) HasModulePath() bool`

HasModulePath returns a boolean if a field has been set.

### GetClassName

`func (o *ClassReference) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *ClassReference) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *ClassReference) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *ClassReference) HasClassName() bool`

HasClassName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


