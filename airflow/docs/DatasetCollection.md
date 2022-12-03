# DatasetCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datasets** | Pointer to [**[]Dataset**](Dataset.md) |  | [optional] 
**TotalEntries** | Pointer to **int32** | Count of objects in the current result set. | [optional] 

## Methods

### NewDatasetCollection

`func NewDatasetCollection() *DatasetCollection`

NewDatasetCollection instantiates a new DatasetCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetCollectionWithDefaults

`func NewDatasetCollectionWithDefaults() *DatasetCollection`

NewDatasetCollectionWithDefaults instantiates a new DatasetCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasets

`func (o *DatasetCollection) GetDatasets() []Dataset`

GetDatasets returns the Datasets field if non-nil, zero value otherwise.

### GetDatasetsOk

`func (o *DatasetCollection) GetDatasetsOk() (*[]Dataset, bool)`

GetDatasetsOk returns a tuple with the Datasets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasets

`func (o *DatasetCollection) SetDatasets(v []Dataset)`

SetDatasets sets Datasets field to given value.

### HasDatasets

`func (o *DatasetCollection) HasDatasets() bool`

HasDatasets returns a boolean if a field has been set.

### GetTotalEntries

`func (o *DatasetCollection) GetTotalEntries() int32`

GetTotalEntries returns the TotalEntries field if non-nil, zero value otherwise.

### GetTotalEntriesOk

`func (o *DatasetCollection) GetTotalEntriesOk() (*int32, bool)`

GetTotalEntriesOk returns a tuple with the TotalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntries

`func (o *DatasetCollection) SetTotalEntries(v int32)`

SetTotalEntries sets TotalEntries field to given value.

### HasTotalEntries

`func (o *DatasetCollection) HasTotalEntries() bool`

HasTotalEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


