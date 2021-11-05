# DAGCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dags** | Pointer to [**[]DAG**](DAG.md) |  | [optional] 
**TotalEntries** | Pointer to **int32** | Count of objects in the current result set. | [optional] 

## Methods

### NewDAGCollection

`func NewDAGCollection() *DAGCollection`

NewDAGCollection instantiates a new DAGCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDAGCollectionWithDefaults

`func NewDAGCollectionWithDefaults() *DAGCollection`

NewDAGCollectionWithDefaults instantiates a new DAGCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDags

`func (o *DAGCollection) GetDags() []DAG`

GetDags returns the Dags field if non-nil, zero value otherwise.

### GetDagsOk

`func (o *DAGCollection) GetDagsOk() (*[]DAG, bool)`

GetDagsOk returns a tuple with the Dags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDags

`func (o *DAGCollection) SetDags(v []DAG)`

SetDags sets Dags field to given value.

### HasDags

`func (o *DAGCollection) HasDags() bool`

HasDags returns a boolean if a field has been set.

### GetTotalEntries

`func (o *DAGCollection) GetTotalEntries() int32`

GetTotalEntries returns the TotalEntries field if non-nil, zero value otherwise.

### GetTotalEntriesOk

`func (o *DAGCollection) GetTotalEntriesOk() (*int32, bool)`

GetTotalEntriesOk returns a tuple with the TotalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntries

`func (o *DAGCollection) SetTotalEntries(v int32)`

SetTotalEntries sets TotalEntries field to given value.

### HasTotalEntries

`func (o *DAGCollection) HasTotalEntries() bool`

HasTotalEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


