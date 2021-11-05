# HealthInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadatabase** | Pointer to [**MetadatabaseStatus**](MetadatabaseStatus.md) |  | [optional] 
**Scheduler** | Pointer to [**SchedulerStatus**](SchedulerStatus.md) |  | [optional] 

## Methods

### NewHealthInfo

`func NewHealthInfo() *HealthInfo`

NewHealthInfo instantiates a new HealthInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthInfoWithDefaults

`func NewHealthInfoWithDefaults() *HealthInfo`

NewHealthInfoWithDefaults instantiates a new HealthInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadatabase

`func (o *HealthInfo) GetMetadatabase() MetadatabaseStatus`

GetMetadatabase returns the Metadatabase field if non-nil, zero value otherwise.

### GetMetadatabaseOk

`func (o *HealthInfo) GetMetadatabaseOk() (*MetadatabaseStatus, bool)`

GetMetadatabaseOk returns a tuple with the Metadatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadatabase

`func (o *HealthInfo) SetMetadatabase(v MetadatabaseStatus)`

SetMetadatabase sets Metadatabase field to given value.

### HasMetadatabase

`func (o *HealthInfo) HasMetadatabase() bool`

HasMetadatabase returns a boolean if a field has been set.

### GetScheduler

`func (o *HealthInfo) GetScheduler() SchedulerStatus`

GetScheduler returns the Scheduler field if non-nil, zero value otherwise.

### GetSchedulerOk

`func (o *HealthInfo) GetSchedulerOk() (*SchedulerStatus, bool)`

GetSchedulerOk returns a tuple with the Scheduler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduler

`func (o *HealthInfo) SetScheduler(v SchedulerStatus)`

SetScheduler sets Scheduler field to given value.

### HasScheduler

`func (o *HealthInfo) HasScheduler() bool`

HasScheduler returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


