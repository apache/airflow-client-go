# DatasetEventCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetEvents** | Pointer to [**[]DatasetEvent**](DatasetEvent.md) |  | [optional] 
**TotalEntries** | Pointer to **int32** | Count of objects in the current result set. | [optional] 

## Methods

### NewDatasetEventCollection

`func NewDatasetEventCollection() *DatasetEventCollection`

NewDatasetEventCollection instantiates a new DatasetEventCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetEventCollectionWithDefaults

`func NewDatasetEventCollectionWithDefaults() *DatasetEventCollection`

NewDatasetEventCollectionWithDefaults instantiates a new DatasetEventCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasetEvents

`func (o *DatasetEventCollection) GetDatasetEvents() []DatasetEvent`

GetDatasetEvents returns the DatasetEvents field if non-nil, zero value otherwise.

### GetDatasetEventsOk

`func (o *DatasetEventCollection) GetDatasetEventsOk() (*[]DatasetEvent, bool)`

GetDatasetEventsOk returns a tuple with the DatasetEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetEvents

`func (o *DatasetEventCollection) SetDatasetEvents(v []DatasetEvent)`

SetDatasetEvents sets DatasetEvents field to given value.

### HasDatasetEvents

`func (o *DatasetEventCollection) HasDatasetEvents() bool`

HasDatasetEvents returns a boolean if a field has been set.

### GetTotalEntries

`func (o *DatasetEventCollection) GetTotalEntries() int32`

GetTotalEntries returns the TotalEntries field if non-nil, zero value otherwise.

### GetTotalEntriesOk

`func (o *DatasetEventCollection) GetTotalEntriesOk() (*int32, bool)`

GetTotalEntriesOk returns a tuple with the TotalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntries

`func (o *DatasetEventCollection) SetTotalEntries(v int32)`

SetTotalEntries sets TotalEntries field to given value.

### HasTotalEntries

`func (o *DatasetEventCollection) HasTotalEntries() bool`

HasTotalEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


