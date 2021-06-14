# UserCollectionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | The user firstname | [optional] 
**LastName** | Pointer to **string** | The user lastname | [optional] 
**Username** | Pointer to **string** | The username | [optional] 
**Email** | Pointer to **string** | The user&#39;s email | [optional] 
**Active** | Pointer to **NullableBool** | Whether the user is active | [optional] [readonly] 
**LastLogin** | Pointer to **NullableString** | The last user login | [optional] [readonly] 
**LoginCount** | Pointer to **NullableInt32** | The login count | [optional] [readonly] 
**FailedLoginCount** | Pointer to **NullableInt32** | The number of times the login failed | [optional] [readonly] 
**Roles** | Pointer to [**[]UserCollectionItemRoles**](UserCollectionItemRoles.md) | User roles | [optional] [readonly] 
**CreatedOn** | Pointer to **NullableString** | The date user was created | [optional] [readonly] 
**ChangedOn** | Pointer to **NullableString** | The date user was changed | [optional] [readonly] 

## Methods

### NewUserCollectionItem

`func NewUserCollectionItem() *UserCollectionItem`

NewUserCollectionItem instantiates a new UserCollectionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCollectionItemWithDefaults

`func NewUserCollectionItemWithDefaults() *UserCollectionItem`

NewUserCollectionItemWithDefaults instantiates a new UserCollectionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *UserCollectionItem) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserCollectionItem) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserCollectionItem) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserCollectionItem) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserCollectionItem) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserCollectionItem) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserCollectionItem) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserCollectionItem) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetUsername

`func (o *UserCollectionItem) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserCollectionItem) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserCollectionItem) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserCollectionItem) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *UserCollectionItem) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserCollectionItem) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserCollectionItem) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserCollectionItem) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetActive

`func (o *UserCollectionItem) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UserCollectionItem) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UserCollectionItem) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UserCollectionItem) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *UserCollectionItem) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *UserCollectionItem) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetLastLogin

`func (o *UserCollectionItem) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *UserCollectionItem) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *UserCollectionItem) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *UserCollectionItem) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### SetLastLoginNil

`func (o *UserCollectionItem) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *UserCollectionItem) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetLoginCount

`func (o *UserCollectionItem) GetLoginCount() int32`

GetLoginCount returns the LoginCount field if non-nil, zero value otherwise.

### GetLoginCountOk

`func (o *UserCollectionItem) GetLoginCountOk() (*int32, bool)`

GetLoginCountOk returns a tuple with the LoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCount

`func (o *UserCollectionItem) SetLoginCount(v int32)`

SetLoginCount sets LoginCount field to given value.

### HasLoginCount

`func (o *UserCollectionItem) HasLoginCount() bool`

HasLoginCount returns a boolean if a field has been set.

### SetLoginCountNil

`func (o *UserCollectionItem) SetLoginCountNil(b bool)`

 SetLoginCountNil sets the value for LoginCount to be an explicit nil

### UnsetLoginCount
`func (o *UserCollectionItem) UnsetLoginCount()`

UnsetLoginCount ensures that no value is present for LoginCount, not even an explicit nil
### GetFailedLoginCount

`func (o *UserCollectionItem) GetFailedLoginCount() int32`

GetFailedLoginCount returns the FailedLoginCount field if non-nil, zero value otherwise.

### GetFailedLoginCountOk

`func (o *UserCollectionItem) GetFailedLoginCountOk() (*int32, bool)`

GetFailedLoginCountOk returns a tuple with the FailedLoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedLoginCount

`func (o *UserCollectionItem) SetFailedLoginCount(v int32)`

SetFailedLoginCount sets FailedLoginCount field to given value.

### HasFailedLoginCount

`func (o *UserCollectionItem) HasFailedLoginCount() bool`

HasFailedLoginCount returns a boolean if a field has been set.

### SetFailedLoginCountNil

`func (o *UserCollectionItem) SetFailedLoginCountNil(b bool)`

 SetFailedLoginCountNil sets the value for FailedLoginCount to be an explicit nil

### UnsetFailedLoginCount
`func (o *UserCollectionItem) UnsetFailedLoginCount()`

UnsetFailedLoginCount ensures that no value is present for FailedLoginCount, not even an explicit nil
### GetRoles

`func (o *UserCollectionItem) GetRoles() []UserCollectionItemRoles`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserCollectionItem) GetRolesOk() (*[]UserCollectionItemRoles, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserCollectionItem) SetRoles(v []UserCollectionItemRoles)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserCollectionItem) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetCreatedOn

`func (o *UserCollectionItem) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *UserCollectionItem) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *UserCollectionItem) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *UserCollectionItem) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### SetCreatedOnNil

`func (o *UserCollectionItem) SetCreatedOnNil(b bool)`

 SetCreatedOnNil sets the value for CreatedOn to be an explicit nil

### UnsetCreatedOn
`func (o *UserCollectionItem) UnsetCreatedOn()`

UnsetCreatedOn ensures that no value is present for CreatedOn, not even an explicit nil
### GetChangedOn

`func (o *UserCollectionItem) GetChangedOn() string`

GetChangedOn returns the ChangedOn field if non-nil, zero value otherwise.

### GetChangedOnOk

`func (o *UserCollectionItem) GetChangedOnOk() (*string, bool)`

GetChangedOnOk returns a tuple with the ChangedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedOn

`func (o *UserCollectionItem) SetChangedOn(v string)`

SetChangedOn sets ChangedOn field to given value.

### HasChangedOn

`func (o *UserCollectionItem) HasChangedOn() bool`

HasChangedOn returns a boolean if a field has been set.

### SetChangedOnNil

`func (o *UserCollectionItem) SetChangedOnNil(b bool)`

 SetChangedOnNil sets the value for ChangedOn to be an explicit nil

### UnsetChangedOn
`func (o *UserCollectionItem) UnsetChangedOn()`

UnsetChangedOn ensures that no value is present for ChangedOn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


