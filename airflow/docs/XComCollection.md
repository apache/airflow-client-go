# XComCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**XcomEntries** | Pointer to [**[]XComCollectionItem**](XComCollectionItem.md) |  | [optional] 
**TotalEntries** | Pointer to **int32** | Count of objects in the current result set. | [optional] 

## Methods

### NewXComCollection

`func NewXComCollection() *XComCollection`

NewXComCollection instantiates a new XComCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXComCollectionWithDefaults

`func NewXComCollectionWithDefaults() *XComCollection`

NewXComCollectionWithDefaults instantiates a new XComCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetXcomEntries

`func (o *XComCollection) GetXcomEntries() []XComCollectionItem`

GetXcomEntries returns the XcomEntries field if non-nil, zero value otherwise.

### GetXcomEntriesOk

`func (o *XComCollection) GetXcomEntriesOk() (*[]XComCollectionItem, bool)`

GetXcomEntriesOk returns a tuple with the XcomEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXcomEntries

`func (o *XComCollection) SetXcomEntries(v []XComCollectionItem)`

SetXcomEntries sets XcomEntries field to given value.

### HasXcomEntries

`func (o *XComCollection) HasXcomEntries() bool`

HasXcomEntries returns a boolean if a field has been set.

### GetTotalEntries

`func (o *XComCollection) GetTotalEntries() int32`

GetTotalEntries returns the TotalEntries field if non-nil, zero value otherwise.

### GetTotalEntriesOk

`func (o *XComCollection) GetTotalEntriesOk() (*int32, bool)`

GetTotalEntriesOk returns a tuple with the TotalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntries

`func (o *XComCollection) SetTotalEntries(v int32)`

SetTotalEntries sets TotalEntries field to given value.

### HasTotalEntries

`func (o *XComCollection) HasTotalEntries() bool`

HasTotalEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


