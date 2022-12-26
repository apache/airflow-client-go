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

# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | The user&#39;s first name.  *Changed in version 2.4.0*&amp;#58; The requirement for this to be non-empty was removed.  | [optional] 
**LastName** | Pointer to **string** | The user&#39;s last name.  *Changed in version 2.4.0*&amp;#58; The requirement for this to be non-empty was removed.  | [optional] 
**Username** | Pointer to **string** | The username.  *Changed in version 2.2.0*&amp;#58; A minimum character length requirement (&#39;minLength&#39;) is added.  | [optional] 
**Email** | Pointer to **string** | The user&#39;s email.  *Changed in version 2.2.0*&amp;#58; A minimum character length requirement (&#39;minLength&#39;) is added.  | [optional] 
**Active** | Pointer to **NullableBool** | Whether the user is active | [optional] [readonly] 
**LastLogin** | Pointer to **NullableString** | The last user login | [optional] [readonly] 
**LoginCount** | Pointer to **NullableInt32** | The login count | [optional] [readonly] 
**FailedLoginCount** | Pointer to **NullableInt32** | The number of times the login failed | [optional] [readonly] 
**Roles** | Pointer to [**[]UserCollectionItemRoles**](UserCollectionItemRoles.md) | User roles.  *Changed in version 2.2.0*&amp;#58; Field is no longer read-only.  | [optional] 
**CreatedOn** | Pointer to **NullableString** | The date user was created | [optional] [readonly] 
**ChangedOn** | Pointer to **NullableString** | The date user was changed | [optional] [readonly] 
**Password** | Pointer to **string** |  | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *User) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *User) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *User) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *User) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *User) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *User) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *User) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *User) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *User) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetActive

`func (o *User) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *User) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *User) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *User) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *User) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *User) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetLastLogin

`func (o *User) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *User) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *User) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *User) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### SetLastLoginNil

`func (o *User) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *User) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetLoginCount

`func (o *User) GetLoginCount() int32`

GetLoginCount returns the LoginCount field if non-nil, zero value otherwise.

### GetLoginCountOk

`func (o *User) GetLoginCountOk() (*int32, bool)`

GetLoginCountOk returns a tuple with the LoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCount

`func (o *User) SetLoginCount(v int32)`

SetLoginCount sets LoginCount field to given value.

### HasLoginCount

`func (o *User) HasLoginCount() bool`

HasLoginCount returns a boolean if a field has been set.

### SetLoginCountNil

`func (o *User) SetLoginCountNil(b bool)`

 SetLoginCountNil sets the value for LoginCount to be an explicit nil

### UnsetLoginCount
`func (o *User) UnsetLoginCount()`

UnsetLoginCount ensures that no value is present for LoginCount, not even an explicit nil
### GetFailedLoginCount

`func (o *User) GetFailedLoginCount() int32`

GetFailedLoginCount returns the FailedLoginCount field if non-nil, zero value otherwise.

### GetFailedLoginCountOk

`func (o *User) GetFailedLoginCountOk() (*int32, bool)`

GetFailedLoginCountOk returns a tuple with the FailedLoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedLoginCount

`func (o *User) SetFailedLoginCount(v int32)`

SetFailedLoginCount sets FailedLoginCount field to given value.

### HasFailedLoginCount

`func (o *User) HasFailedLoginCount() bool`

HasFailedLoginCount returns a boolean if a field has been set.

### SetFailedLoginCountNil

`func (o *User) SetFailedLoginCountNil(b bool)`

 SetFailedLoginCountNil sets the value for FailedLoginCount to be an explicit nil

### UnsetFailedLoginCount
`func (o *User) UnsetFailedLoginCount()`

UnsetFailedLoginCount ensures that no value is present for FailedLoginCount, not even an explicit nil
### GetRoles

`func (o *User) GetRoles() []UserCollectionItemRoles`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *User) GetRolesOk() (*[]UserCollectionItemRoles, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *User) SetRoles(v []UserCollectionItemRoles)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *User) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetCreatedOn

`func (o *User) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *User) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *User) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *User) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### SetCreatedOnNil

`func (o *User) SetCreatedOnNil(b bool)`

 SetCreatedOnNil sets the value for CreatedOn to be an explicit nil

### UnsetCreatedOn
`func (o *User) UnsetCreatedOn()`

UnsetCreatedOn ensures that no value is present for CreatedOn, not even an explicit nil
### GetChangedOn

`func (o *User) GetChangedOn() string`

GetChangedOn returns the ChangedOn field if non-nil, zero value otherwise.

### GetChangedOnOk

`func (o *User) GetChangedOnOk() (*string, bool)`

GetChangedOnOk returns a tuple with the ChangedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedOn

`func (o *User) SetChangedOn(v string)`

SetChangedOn sets ChangedOn field to given value.

### HasChangedOn

`func (o *User) HasChangedOn() bool`

HasChangedOn returns a boolean if a field has been set.

### SetChangedOnNil

`func (o *User) SetChangedOnNil(b bool)`

 SetChangedOnNil sets the value for ChangedOn to be an explicit nil

### UnsetChangedOn
`func (o *User) UnsetChangedOn()`

UnsetChangedOn ensures that no value is present for ChangedOn, not even an explicit nil
### GetPassword

`func (o *User) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *User) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *User) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *User) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


