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

# VersionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | The version of Airflow | [optional] 
**GitVersion** | Pointer to **NullableString** | The git version (including git commit hash) | [optional] 

## Methods

### NewVersionInfo

`func NewVersionInfo() *VersionInfo`

NewVersionInfo instantiates a new VersionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionInfoWithDefaults

`func NewVersionInfoWithDefaults() *VersionInfo`

NewVersionInfoWithDefaults instantiates a new VersionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *VersionInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VersionInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VersionInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VersionInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetGitVersion

`func (o *VersionInfo) GetGitVersion() string`

GetGitVersion returns the GitVersion field if non-nil, zero value otherwise.

### GetGitVersionOk

`func (o *VersionInfo) GetGitVersionOk() (*string, bool)`

GetGitVersionOk returns a tuple with the GitVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitVersion

`func (o *VersionInfo) SetGitVersion(v string)`

SetGitVersion sets GitVersion field to given value.

### HasGitVersion

`func (o *VersionInfo) HasGitVersion() bool`

HasGitVersion returns a boolean if a field has been set.

### SetGitVersionNil

`func (o *VersionInfo) SetGitVersionNil(b bool)`

 SetGitVersionNil sets the value for GitVersion to be an explicit nil

### UnsetGitVersion
`func (o *VersionInfo) UnsetGitVersion()`

UnsetGitVersion ensures that no value is present for GitVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


