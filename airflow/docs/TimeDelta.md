# TimeDelta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Days** | **int32** |  | 
**Seconds** | **int32** |  | 
**Microseconds** | **int32** |  | 

## Methods

### NewTimeDelta

`func NewTimeDelta(type_ string, days int32, seconds int32, microseconds int32, ) *TimeDelta`

NewTimeDelta instantiates a new TimeDelta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeDeltaWithDefaults

`func NewTimeDeltaWithDefaults() *TimeDelta`

NewTimeDeltaWithDefaults instantiates a new TimeDelta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TimeDelta) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimeDelta) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimeDelta) SetType(v string)`

SetType sets Type field to given value.


### GetDays

`func (o *TimeDelta) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *TimeDelta) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *TimeDelta) SetDays(v int32)`

SetDays sets Days field to given value.


### GetSeconds

`func (o *TimeDelta) GetSeconds() int32`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *TimeDelta) GetSecondsOk() (*int32, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *TimeDelta) SetSeconds(v int32)`

SetSeconds sets Seconds field to given value.


### GetMicroseconds

`func (o *TimeDelta) GetMicroseconds() int32`

GetMicroseconds returns the Microseconds field if non-nil, zero value otherwise.

### GetMicrosecondsOk

`func (o *TimeDelta) GetMicrosecondsOk() (*int32, bool)`

GetMicrosecondsOk returns a tuple with the Microseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroseconds

`func (o *TimeDelta) SetMicroseconds(v int32)`

SetMicroseconds sets Microseconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


