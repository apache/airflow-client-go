# DagScheduleDatasetReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DagId** | Pointer to **NullableString** | The DAG ID that depends on the dataset. | [optional] 
**CreatedAt** | Pointer to **string** | The dataset reference creation time | [optional] 
**UpdatedAt** | Pointer to **string** | The dataset reference update time | [optional] 

## Methods

### NewDagScheduleDatasetReference

`func NewDagScheduleDatasetReference() *DagScheduleDatasetReference`

NewDagScheduleDatasetReference instantiates a new DagScheduleDatasetReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDagScheduleDatasetReferenceWithDefaults

`func NewDagScheduleDatasetReferenceWithDefaults() *DagScheduleDatasetReference`

NewDagScheduleDatasetReferenceWithDefaults instantiates a new DagScheduleDatasetReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDagId

`func (o *DagScheduleDatasetReference) GetDagId() string`

GetDagId returns the DagId field if non-nil, zero value otherwise.

### GetDagIdOk

`func (o *DagScheduleDatasetReference) GetDagIdOk() (*string, bool)`

GetDagIdOk returns a tuple with the DagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagId

`func (o *DagScheduleDatasetReference) SetDagId(v string)`

SetDagId sets DagId field to given value.

### HasDagId

`func (o *DagScheduleDatasetReference) HasDagId() bool`

HasDagId returns a boolean if a field has been set.

### SetDagIdNil

`func (o *DagScheduleDatasetReference) SetDagIdNil(b bool)`

 SetDagIdNil sets the value for DagId to be an explicit nil

### UnsetDagId
`func (o *DagScheduleDatasetReference) UnsetDagId()`

UnsetDagId ensures that no value is present for DagId, not even an explicit nil
### GetCreatedAt

`func (o *DagScheduleDatasetReference) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DagScheduleDatasetReference) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DagScheduleDatasetReference) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DagScheduleDatasetReference) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DagScheduleDatasetReference) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DagScheduleDatasetReference) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DagScheduleDatasetReference) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DagScheduleDatasetReference) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


