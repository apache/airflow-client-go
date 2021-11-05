# PluginCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plugins** | Pointer to [**[]PluginCollectionItem**](PluginCollectionItem.md) |  | [optional] 
**TotalEntries** | Pointer to **int32** | Count of objects in the current result set. | [optional] 

## Methods

### NewPluginCollection

`func NewPluginCollection() *PluginCollection`

NewPluginCollection instantiates a new PluginCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginCollectionWithDefaults

`func NewPluginCollectionWithDefaults() *PluginCollection`

NewPluginCollectionWithDefaults instantiates a new PluginCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlugins

`func (o *PluginCollection) GetPlugins() []PluginCollectionItem`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *PluginCollection) GetPluginsOk() (*[]PluginCollectionItem, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *PluginCollection) SetPlugins(v []PluginCollectionItem)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *PluginCollection) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### GetTotalEntries

`func (o *PluginCollection) GetTotalEntries() int32`

GetTotalEntries returns the TotalEntries field if non-nil, zero value otherwise.

### GetTotalEntriesOk

`func (o *PluginCollection) GetTotalEntriesOk() (*int32, bool)`

GetTotalEntriesOk returns a tuple with the TotalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntries

`func (o *PluginCollection) SetTotalEntries(v int32)`

SetTotalEntries sets TotalEntries field to given value.

### HasTotalEntries

`func (o *PluginCollection) HasTotalEntries() bool`

HasTotalEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


