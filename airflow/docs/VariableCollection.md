# VariableCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variables** | Pointer to [**[]VariableCollectionItem**](VariableCollectionItem.md) |  | [optional] 
**TotalEntries** | Pointer to **int32** | Count of objects in the current result set. | [optional] 

## Methods

### NewVariableCollection

`func NewVariableCollection() *VariableCollection`

NewVariableCollection instantiates a new VariableCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableCollectionWithDefaults

`func NewVariableCollectionWithDefaults() *VariableCollection`

NewVariableCollectionWithDefaults instantiates a new VariableCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariables

`func (o *VariableCollection) GetVariables() []VariableCollectionItem`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *VariableCollection) GetVariablesOk() (*[]VariableCollectionItem, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *VariableCollection) SetVariables(v []VariableCollectionItem)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *VariableCollection) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetTotalEntries

`func (o *VariableCollection) GetTotalEntries() int32`

GetTotalEntries returns the TotalEntries field if non-nil, zero value otherwise.

### GetTotalEntriesOk

`func (o *VariableCollection) GetTotalEntriesOk() (*int32, bool)`

GetTotalEntriesOk returns a tuple with the TotalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntries

`func (o *VariableCollection) SetTotalEntries(v int32)`

SetTotalEntries sets TotalEntries field to given value.

### HasTotalEntries

`func (o *VariableCollection) HasTotalEntries() bool`

HasTotalEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


