# ConnectionTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** | The status of the request. | [optional] 
**Message** | Pointer to **string** | The success or failure message of the request. | [optional] 

## Methods

### NewConnectionTest

`func NewConnectionTest() *ConnectionTest`

NewConnectionTest instantiates a new ConnectionTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionTestWithDefaults

`func NewConnectionTestWithDefaults() *ConnectionTest`

NewConnectionTestWithDefaults instantiates a new ConnectionTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ConnectionTest) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectionTest) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectionTest) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConnectionTest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *ConnectionTest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConnectionTest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConnectionTest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConnectionTest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


