# BasicDAGRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | Pointer to **string** | Run ID.  | [optional] 
**DagId** | Pointer to **string** |  | [optional] [readonly] 
**LogicalDate** | Pointer to **time.Time** | The logical date (previously called execution date). This is the time or interval covered by this DAG run, according to the DAG definition.  The value of this field can be set only when creating the object. If you try to modify the field of an existing object, the request fails with an BAD_REQUEST error.  This together with DAG_ID are a unique key.  *New in version 2.2.0*  | [optional] 
**StartDate** | Pointer to **NullableTime** | The start time. The time when DAG run was actually created.  *Changed in version 2.1.3*&amp;#58; Field becomes nullable.  | [optional] [readonly] 
**EndDate** | Pointer to **NullableTime** |  | [optional] [readonly] 
**DataIntervalStart** | Pointer to **NullableTime** |  | [optional] [readonly] 
**DataIntervalEnd** | Pointer to **NullableTime** |  | [optional] [readonly] 
**State** | Pointer to [**DagState**](DagState.md) |  | [optional] 

## Methods

### NewBasicDAGRun

`func NewBasicDAGRun() *BasicDAGRun`

NewBasicDAGRun instantiates a new BasicDAGRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicDAGRunWithDefaults

`func NewBasicDAGRunWithDefaults() *BasicDAGRun`

NewBasicDAGRunWithDefaults instantiates a new BasicDAGRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *BasicDAGRun) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *BasicDAGRun) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *BasicDAGRun) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *BasicDAGRun) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetDagId

`func (o *BasicDAGRun) GetDagId() string`

GetDagId returns the DagId field if non-nil, zero value otherwise.

### GetDagIdOk

`func (o *BasicDAGRun) GetDagIdOk() (*string, bool)`

GetDagIdOk returns a tuple with the DagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagId

`func (o *BasicDAGRun) SetDagId(v string)`

SetDagId sets DagId field to given value.

### HasDagId

`func (o *BasicDAGRun) HasDagId() bool`

HasDagId returns a boolean if a field has been set.

### GetLogicalDate

`func (o *BasicDAGRun) GetLogicalDate() time.Time`

GetLogicalDate returns the LogicalDate field if non-nil, zero value otherwise.

### GetLogicalDateOk

`func (o *BasicDAGRun) GetLogicalDateOk() (*time.Time, bool)`

GetLogicalDateOk returns a tuple with the LogicalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalDate

`func (o *BasicDAGRun) SetLogicalDate(v time.Time)`

SetLogicalDate sets LogicalDate field to given value.

### HasLogicalDate

`func (o *BasicDAGRun) HasLogicalDate() bool`

HasLogicalDate returns a boolean if a field has been set.

### GetStartDate

`func (o *BasicDAGRun) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BasicDAGRun) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BasicDAGRun) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BasicDAGRun) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *BasicDAGRun) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *BasicDAGRun) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *BasicDAGRun) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BasicDAGRun) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BasicDAGRun) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BasicDAGRun) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *BasicDAGRun) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *BasicDAGRun) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetDataIntervalStart

`func (o *BasicDAGRun) GetDataIntervalStart() time.Time`

GetDataIntervalStart returns the DataIntervalStart field if non-nil, zero value otherwise.

### GetDataIntervalStartOk

`func (o *BasicDAGRun) GetDataIntervalStartOk() (*time.Time, bool)`

GetDataIntervalStartOk returns a tuple with the DataIntervalStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataIntervalStart

`func (o *BasicDAGRun) SetDataIntervalStart(v time.Time)`

SetDataIntervalStart sets DataIntervalStart field to given value.

### HasDataIntervalStart

`func (o *BasicDAGRun) HasDataIntervalStart() bool`

HasDataIntervalStart returns a boolean if a field has been set.

### SetDataIntervalStartNil

`func (o *BasicDAGRun) SetDataIntervalStartNil(b bool)`

 SetDataIntervalStartNil sets the value for DataIntervalStart to be an explicit nil

### UnsetDataIntervalStart
`func (o *BasicDAGRun) UnsetDataIntervalStart()`

UnsetDataIntervalStart ensures that no value is present for DataIntervalStart, not even an explicit nil
### GetDataIntervalEnd

`func (o *BasicDAGRun) GetDataIntervalEnd() time.Time`

GetDataIntervalEnd returns the DataIntervalEnd field if non-nil, zero value otherwise.

### GetDataIntervalEndOk

`func (o *BasicDAGRun) GetDataIntervalEndOk() (*time.Time, bool)`

GetDataIntervalEndOk returns a tuple with the DataIntervalEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataIntervalEnd

`func (o *BasicDAGRun) SetDataIntervalEnd(v time.Time)`

SetDataIntervalEnd sets DataIntervalEnd field to given value.

### HasDataIntervalEnd

`func (o *BasicDAGRun) HasDataIntervalEnd() bool`

HasDataIntervalEnd returns a boolean if a field has been set.

### SetDataIntervalEndNil

`func (o *BasicDAGRun) SetDataIntervalEndNil(b bool)`

 SetDataIntervalEndNil sets the value for DataIntervalEnd to be an explicit nil

### UnsetDataIntervalEnd
`func (o *BasicDAGRun) UnsetDataIntervalEnd()`

UnsetDataIntervalEnd ensures that no value is present for DataIntervalEnd, not even an explicit nil
### GetState

`func (o *BasicDAGRun) GetState() DagState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BasicDAGRun) GetStateOk() (*DagState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BasicDAGRun) SetState(v DagState)`

SetState sets State field to given value.

### HasState

`func (o *BasicDAGRun) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


