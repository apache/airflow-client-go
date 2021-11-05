# ConnectionCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | Pointer to [**[]ConnectionCollectionItem**](ConnectionCollectionItem.md) |  | [optional] 
**TotalEntries** | Pointer to **int32** | Count of objects in the current result set. | [optional] 

## Methods

### NewConnectionCollection

`func NewConnectionCollection() *ConnectionCollection`

NewConnectionCollection instantiates a new ConnectionCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionCollectionWithDefaults

`func NewConnectionCollectionWithDefaults() *ConnectionCollection`

NewConnectionCollectionWithDefaults instantiates a new ConnectionCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *ConnectionCollection) GetConnections() []ConnectionCollectionItem`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *ConnectionCollection) GetConnectionsOk() (*[]ConnectionCollectionItem, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *ConnectionCollection) SetConnections(v []ConnectionCollectionItem)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *ConnectionCollection) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetTotalEntries

`func (o *ConnectionCollection) GetTotalEntries() int32`

GetTotalEntries returns the TotalEntries field if non-nil, zero value otherwise.

### GetTotalEntriesOk

`func (o *ConnectionCollection) GetTotalEntriesOk() (*int32, bool)`

GetTotalEntriesOk returns a tuple with the TotalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntries

`func (o *ConnectionCollection) SetTotalEntries(v int32)`

SetTotalEntries sets TotalEntries field to given value.

### HasTotalEntries

`func (o *ConnectionCollection) HasTotalEntries() bool`

HasTotalEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


