# PoolCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pools** | Pointer to [**[]Pool**](Pool.md) |  | [optional] 
**TotalEntries** | Pointer to **int32** | Count of objects in the current result set. | [optional] 

## Methods

### NewPoolCollection

`func NewPoolCollection() *PoolCollection`

NewPoolCollection instantiates a new PoolCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolCollectionWithDefaults

`func NewPoolCollectionWithDefaults() *PoolCollection`

NewPoolCollectionWithDefaults instantiates a new PoolCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPools

`func (o *PoolCollection) GetPools() []Pool`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *PoolCollection) GetPoolsOk() (*[]Pool, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *PoolCollection) SetPools(v []Pool)`

SetPools sets Pools field to given value.

### HasPools

`func (o *PoolCollection) HasPools() bool`

HasPools returns a boolean if a field has been set.

### GetTotalEntries

`func (o *PoolCollection) GetTotalEntries() int32`

GetTotalEntries returns the TotalEntries field if non-nil, zero value otherwise.

### GetTotalEntriesOk

`func (o *PoolCollection) GetTotalEntriesOk() (*int32, bool)`

GetTotalEntriesOk returns a tuple with the TotalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntries

`func (o *PoolCollection) SetTotalEntries(v int32)`

SetTotalEntries sets TotalEntries field to given value.

### HasTotalEntries

`func (o *PoolCollection) HasTotalEntries() bool`

HasTotalEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


