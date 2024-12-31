# TriggererStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**NullableHealthStatus**](HealthStatus.md) |  | [optional] 
**LatestTriggererHeartbeat** | Pointer to **NullableString** | The time the triggerer last did a heartbeat. | [optional] [readonly] 

## Methods

### NewTriggererStatus

`func NewTriggererStatus() *TriggererStatus`

NewTriggererStatus instantiates a new TriggererStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggererStatusWithDefaults

`func NewTriggererStatusWithDefaults() *TriggererStatus`

NewTriggererStatusWithDefaults instantiates a new TriggererStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TriggererStatus) GetStatus() HealthStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TriggererStatus) GetStatusOk() (*HealthStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TriggererStatus) SetStatus(v HealthStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TriggererStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TriggererStatus) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TriggererStatus) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetLatestTriggererHeartbeat

`func (o *TriggererStatus) GetLatestTriggererHeartbeat() string`

GetLatestTriggererHeartbeat returns the LatestTriggererHeartbeat field if non-nil, zero value otherwise.

### GetLatestTriggererHeartbeatOk

`func (o *TriggererStatus) GetLatestTriggererHeartbeatOk() (*string, bool)`

GetLatestTriggererHeartbeatOk returns a tuple with the LatestTriggererHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestTriggererHeartbeat

`func (o *TriggererStatus) SetLatestTriggererHeartbeat(v string)`

SetLatestTriggererHeartbeat sets LatestTriggererHeartbeat field to given value.

### HasLatestTriggererHeartbeat

`func (o *TriggererStatus) HasLatestTriggererHeartbeat() bool`

HasLatestTriggererHeartbeat returns a boolean if a field has been set.

### SetLatestTriggererHeartbeatNil

`func (o *TriggererStatus) SetLatestTriggererHeartbeatNil(b bool)`

 SetLatestTriggererHeartbeatNil sets the value for LatestTriggererHeartbeat to be an explicit nil

### UnsetLatestTriggererHeartbeat
`func (o *TriggererStatus) UnsetLatestTriggererHeartbeat()`

UnsetLatestTriggererHeartbeat ensures that no value is present for LatestTriggererHeartbeat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


