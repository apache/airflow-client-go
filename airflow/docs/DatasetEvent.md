# DatasetEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetId** | Pointer to **int32** | The dataset id | [optional] 
**DatasetUri** | Pointer to **string** | The URI of the dataset | [optional] 
**Extra** | Pointer to **map[string]interface{}** | The dataset event extra | [optional] 
**SourceDagId** | Pointer to **NullableString** | The DAG ID that updated the dataset. | [optional] 
**SourceTaskId** | Pointer to **NullableString** | The task ID that updated the dataset. | [optional] 
**SourceRunId** | Pointer to **NullableString** | The DAG run ID that updated the dataset. | [optional] 
**SourceMapIndex** | Pointer to **NullableInt32** | The task map index that updated the dataset. | [optional] 
**CreatedDagruns** | Pointer to [**[]BasicDAGRun**](BasicDAGRun.md) |  | [optional] 
**Timestamp** | Pointer to **string** | The dataset event creation time | [optional] 

## Methods

### NewDatasetEvent

`func NewDatasetEvent() *DatasetEvent`

NewDatasetEvent instantiates a new DatasetEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetEventWithDefaults

`func NewDatasetEventWithDefaults() *DatasetEvent`

NewDatasetEventWithDefaults instantiates a new DatasetEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasetId

`func (o *DatasetEvent) GetDatasetId() int32`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *DatasetEvent) GetDatasetIdOk() (*int32, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetId

`func (o *DatasetEvent) SetDatasetId(v int32)`

SetDatasetId sets DatasetId field to given value.

### HasDatasetId

`func (o *DatasetEvent) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### GetDatasetUri

`func (o *DatasetEvent) GetDatasetUri() string`

GetDatasetUri returns the DatasetUri field if non-nil, zero value otherwise.

### GetDatasetUriOk

`func (o *DatasetEvent) GetDatasetUriOk() (*string, bool)`

GetDatasetUriOk returns a tuple with the DatasetUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetUri

`func (o *DatasetEvent) SetDatasetUri(v string)`

SetDatasetUri sets DatasetUri field to given value.

### HasDatasetUri

`func (o *DatasetEvent) HasDatasetUri() bool`

HasDatasetUri returns a boolean if a field has been set.

### GetExtra

`func (o *DatasetEvent) GetExtra() map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *DatasetEvent) GetExtraOk() (*map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *DatasetEvent) SetExtra(v map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *DatasetEvent) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### SetExtraNil

`func (o *DatasetEvent) SetExtraNil(b bool)`

 SetExtraNil sets the value for Extra to be an explicit nil

### UnsetExtra
`func (o *DatasetEvent) UnsetExtra()`

UnsetExtra ensures that no value is present for Extra, not even an explicit nil
### GetSourceDagId

`func (o *DatasetEvent) GetSourceDagId() string`

GetSourceDagId returns the SourceDagId field if non-nil, zero value otherwise.

### GetSourceDagIdOk

`func (o *DatasetEvent) GetSourceDagIdOk() (*string, bool)`

GetSourceDagIdOk returns a tuple with the SourceDagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDagId

`func (o *DatasetEvent) SetSourceDagId(v string)`

SetSourceDagId sets SourceDagId field to given value.

### HasSourceDagId

`func (o *DatasetEvent) HasSourceDagId() bool`

HasSourceDagId returns a boolean if a field has been set.

### SetSourceDagIdNil

`func (o *DatasetEvent) SetSourceDagIdNil(b bool)`

 SetSourceDagIdNil sets the value for SourceDagId to be an explicit nil

### UnsetSourceDagId
`func (o *DatasetEvent) UnsetSourceDagId()`

UnsetSourceDagId ensures that no value is present for SourceDagId, not even an explicit nil
### GetSourceTaskId

`func (o *DatasetEvent) GetSourceTaskId() string`

GetSourceTaskId returns the SourceTaskId field if non-nil, zero value otherwise.

### GetSourceTaskIdOk

`func (o *DatasetEvent) GetSourceTaskIdOk() (*string, bool)`

GetSourceTaskIdOk returns a tuple with the SourceTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTaskId

`func (o *DatasetEvent) SetSourceTaskId(v string)`

SetSourceTaskId sets SourceTaskId field to given value.

### HasSourceTaskId

`func (o *DatasetEvent) HasSourceTaskId() bool`

HasSourceTaskId returns a boolean if a field has been set.

### SetSourceTaskIdNil

`func (o *DatasetEvent) SetSourceTaskIdNil(b bool)`

 SetSourceTaskIdNil sets the value for SourceTaskId to be an explicit nil

### UnsetSourceTaskId
`func (o *DatasetEvent) UnsetSourceTaskId()`

UnsetSourceTaskId ensures that no value is present for SourceTaskId, not even an explicit nil
### GetSourceRunId

`func (o *DatasetEvent) GetSourceRunId() string`

GetSourceRunId returns the SourceRunId field if non-nil, zero value otherwise.

### GetSourceRunIdOk

`func (o *DatasetEvent) GetSourceRunIdOk() (*string, bool)`

GetSourceRunIdOk returns a tuple with the SourceRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRunId

`func (o *DatasetEvent) SetSourceRunId(v string)`

SetSourceRunId sets SourceRunId field to given value.

### HasSourceRunId

`func (o *DatasetEvent) HasSourceRunId() bool`

HasSourceRunId returns a boolean if a field has been set.

### SetSourceRunIdNil

`func (o *DatasetEvent) SetSourceRunIdNil(b bool)`

 SetSourceRunIdNil sets the value for SourceRunId to be an explicit nil

### UnsetSourceRunId
`func (o *DatasetEvent) UnsetSourceRunId()`

UnsetSourceRunId ensures that no value is present for SourceRunId, not even an explicit nil
### GetSourceMapIndex

`func (o *DatasetEvent) GetSourceMapIndex() int32`

GetSourceMapIndex returns the SourceMapIndex field if non-nil, zero value otherwise.

### GetSourceMapIndexOk

`func (o *DatasetEvent) GetSourceMapIndexOk() (*int32, bool)`

GetSourceMapIndexOk returns a tuple with the SourceMapIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMapIndex

`func (o *DatasetEvent) SetSourceMapIndex(v int32)`

SetSourceMapIndex sets SourceMapIndex field to given value.

### HasSourceMapIndex

`func (o *DatasetEvent) HasSourceMapIndex() bool`

HasSourceMapIndex returns a boolean if a field has been set.

### SetSourceMapIndexNil

`func (o *DatasetEvent) SetSourceMapIndexNil(b bool)`

 SetSourceMapIndexNil sets the value for SourceMapIndex to be an explicit nil

### UnsetSourceMapIndex
`func (o *DatasetEvent) UnsetSourceMapIndex()`

UnsetSourceMapIndex ensures that no value is present for SourceMapIndex, not even an explicit nil
### GetCreatedDagruns

`func (o *DatasetEvent) GetCreatedDagruns() []BasicDAGRun`

GetCreatedDagruns returns the CreatedDagruns field if non-nil, zero value otherwise.

### GetCreatedDagrunsOk

`func (o *DatasetEvent) GetCreatedDagrunsOk() (*[]BasicDAGRun, bool)`

GetCreatedDagrunsOk returns a tuple with the CreatedDagruns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDagruns

`func (o *DatasetEvent) SetCreatedDagruns(v []BasicDAGRun)`

SetCreatedDagruns sets CreatedDagruns field to given value.

### HasCreatedDagruns

`func (o *DatasetEvent) HasCreatedDagruns() bool`

HasCreatedDagruns returns a boolean if a field has been set.

### GetTimestamp

`func (o *DatasetEvent) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DatasetEvent) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DatasetEvent) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DatasetEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


