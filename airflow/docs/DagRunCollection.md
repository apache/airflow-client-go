# DAGRunCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DagRuns** | Pointer to [**[]DAGRun**](DAGRun.md) |  | [optional] 
**TotalEntries** | Pointer to **int32** | Count of objects in the current result set. | [optional] 

## Methods

### NewDAGRunCollection

`func NewDAGRunCollection() *DAGRunCollection`

NewDAGRunCollection instantiates a new DAGRunCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDAGRunCollectionWithDefaults

`func NewDAGRunCollectionWithDefaults() *DAGRunCollection`

NewDAGRunCollectionWithDefaults instantiates a new DAGRunCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDagRuns

`func (o *DAGRunCollection) GetDagRuns() []DAGRun`

GetDagRuns returns the DagRuns field if non-nil, zero value otherwise.

### GetDagRunsOk

`func (o *DAGRunCollection) GetDagRunsOk() (*[]DAGRun, bool)`

GetDagRunsOk returns a tuple with the DagRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagRuns

`func (o *DAGRunCollection) SetDagRuns(v []DAGRun)`

SetDagRuns sets DagRuns field to given value.

### HasDagRuns

`func (o *DAGRunCollection) HasDagRuns() bool`

HasDagRuns returns a boolean if a field has been set.

### GetTotalEntries

`func (o *DAGRunCollection) GetTotalEntries() int32`

GetTotalEntries returns the TotalEntries field if non-nil, zero value otherwise.

### GetTotalEntriesOk

`func (o *DAGRunCollection) GetTotalEntriesOk() (*int32, bool)`

GetTotalEntriesOk returns a tuple with the TotalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntries

`func (o *DAGRunCollection) SetTotalEntries(v int32)`

SetTotalEntries sets TotalEntries field to given value.

### HasTotalEntries

`func (o *DAGRunCollection) HasTotalEntries() bool`

HasTotalEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


