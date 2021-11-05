# SchedulerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**HealthStatus**](HealthStatus.md) |  | [optional] 
**LatestSchedulerHeartbeat** | Pointer to **NullableString** | The time the scheduler last do a heartbeat. | [optional] [readonly] 

## Methods

### NewSchedulerStatus

`func NewSchedulerStatus() *SchedulerStatus`

NewSchedulerStatus instantiates a new SchedulerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerStatusWithDefaults

`func NewSchedulerStatusWithDefaults() *SchedulerStatus`

NewSchedulerStatusWithDefaults instantiates a new SchedulerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SchedulerStatus) GetStatus() HealthStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SchedulerStatus) GetStatusOk() (*HealthStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SchedulerStatus) SetStatus(v HealthStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SchedulerStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLatestSchedulerHeartbeat

`func (o *SchedulerStatus) GetLatestSchedulerHeartbeat() string`

GetLatestSchedulerHeartbeat returns the LatestSchedulerHeartbeat field if non-nil, zero value otherwise.

### GetLatestSchedulerHeartbeatOk

`func (o *SchedulerStatus) GetLatestSchedulerHeartbeatOk() (*string, bool)`

GetLatestSchedulerHeartbeatOk returns a tuple with the LatestSchedulerHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSchedulerHeartbeat

`func (o *SchedulerStatus) SetLatestSchedulerHeartbeat(v string)`

SetLatestSchedulerHeartbeat sets LatestSchedulerHeartbeat field to given value.

### HasLatestSchedulerHeartbeat

`func (o *SchedulerStatus) HasLatestSchedulerHeartbeat() bool`

HasLatestSchedulerHeartbeat returns a boolean if a field has been set.

### SetLatestSchedulerHeartbeatNil

`func (o *SchedulerStatus) SetLatestSchedulerHeartbeatNil(b bool)`

 SetLatestSchedulerHeartbeatNil sets the value for LatestSchedulerHeartbeat to be an explicit nil

### UnsetLatestSchedulerHeartbeat
`func (o *SchedulerStatus) UnsetLatestSchedulerHeartbeat()`

UnsetLatestSchedulerHeartbeat ensures that no value is present for LatestSchedulerHeartbeat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


