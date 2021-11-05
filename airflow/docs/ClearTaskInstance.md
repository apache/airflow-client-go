# ClearTaskInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If set, don&#39;t actually run this operation. The response will contain a list of task instances planned to be cleaned, but not modified in any way.  | [optional] [default to true]
**TaskIds** | Pointer to **[]string** | A list of task ids to clear. | [optional] 
**StartDate** | Pointer to **string** | The minimum execution date to clear. | [optional] 
**EndDate** | Pointer to **string** | The maximum execution date to clear. | [optional] 
**OnlyFailed** | Pointer to **bool** | Only clear failed tasks. | [optional] [default to true]
**OnlyRunning** | Pointer to **bool** | Only clear running tasks. | [optional] [default to false]
**IncludeSubdags** | Pointer to **bool** | Clear tasks in subdags and clear external tasks indicated by ExternalTaskMarker. | [optional] 
**IncludeParentdag** | Pointer to **bool** | Clear tasks in the parent dag of the subdag. | [optional] 
**ResetDagRuns** | Pointer to **bool** | Set state of DAG runs to RUNNING. | [optional] 

## Methods

### NewClearTaskInstance

`func NewClearTaskInstance() *ClearTaskInstance`

NewClearTaskInstance instantiates a new ClearTaskInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearTaskInstanceWithDefaults

`func NewClearTaskInstanceWithDefaults() *ClearTaskInstance`

NewClearTaskInstanceWithDefaults instantiates a new ClearTaskInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ClearTaskInstance) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ClearTaskInstance) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ClearTaskInstance) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ClearTaskInstance) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetTaskIds

`func (o *ClearTaskInstance) GetTaskIds() []string`

GetTaskIds returns the TaskIds field if non-nil, zero value otherwise.

### GetTaskIdsOk

`func (o *ClearTaskInstance) GetTaskIdsOk() (*[]string, bool)`

GetTaskIdsOk returns a tuple with the TaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIds

`func (o *ClearTaskInstance) SetTaskIds(v []string)`

SetTaskIds sets TaskIds field to given value.

### HasTaskIds

`func (o *ClearTaskInstance) HasTaskIds() bool`

HasTaskIds returns a boolean if a field has been set.

### GetStartDate

`func (o *ClearTaskInstance) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ClearTaskInstance) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ClearTaskInstance) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ClearTaskInstance) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ClearTaskInstance) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ClearTaskInstance) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ClearTaskInstance) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ClearTaskInstance) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetOnlyFailed

`func (o *ClearTaskInstance) GetOnlyFailed() bool`

GetOnlyFailed returns the OnlyFailed field if non-nil, zero value otherwise.

### GetOnlyFailedOk

`func (o *ClearTaskInstance) GetOnlyFailedOk() (*bool, bool)`

GetOnlyFailedOk returns a tuple with the OnlyFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyFailed

`func (o *ClearTaskInstance) SetOnlyFailed(v bool)`

SetOnlyFailed sets OnlyFailed field to given value.

### HasOnlyFailed

`func (o *ClearTaskInstance) HasOnlyFailed() bool`

HasOnlyFailed returns a boolean if a field has been set.

### GetOnlyRunning

`func (o *ClearTaskInstance) GetOnlyRunning() bool`

GetOnlyRunning returns the OnlyRunning field if non-nil, zero value otherwise.

### GetOnlyRunningOk

`func (o *ClearTaskInstance) GetOnlyRunningOk() (*bool, bool)`

GetOnlyRunningOk returns a tuple with the OnlyRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyRunning

`func (o *ClearTaskInstance) SetOnlyRunning(v bool)`

SetOnlyRunning sets OnlyRunning field to given value.

### HasOnlyRunning

`func (o *ClearTaskInstance) HasOnlyRunning() bool`

HasOnlyRunning returns a boolean if a field has been set.

### GetIncludeSubdags

`func (o *ClearTaskInstance) GetIncludeSubdags() bool`

GetIncludeSubdags returns the IncludeSubdags field if non-nil, zero value otherwise.

### GetIncludeSubdagsOk

`func (o *ClearTaskInstance) GetIncludeSubdagsOk() (*bool, bool)`

GetIncludeSubdagsOk returns a tuple with the IncludeSubdags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubdags

`func (o *ClearTaskInstance) SetIncludeSubdags(v bool)`

SetIncludeSubdags sets IncludeSubdags field to given value.

### HasIncludeSubdags

`func (o *ClearTaskInstance) HasIncludeSubdags() bool`

HasIncludeSubdags returns a boolean if a field has been set.

### GetIncludeParentdag

`func (o *ClearTaskInstance) GetIncludeParentdag() bool`

GetIncludeParentdag returns the IncludeParentdag field if non-nil, zero value otherwise.

### GetIncludeParentdagOk

`func (o *ClearTaskInstance) GetIncludeParentdagOk() (*bool, bool)`

GetIncludeParentdagOk returns a tuple with the IncludeParentdag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeParentdag

`func (o *ClearTaskInstance) SetIncludeParentdag(v bool)`

SetIncludeParentdag sets IncludeParentdag field to given value.

### HasIncludeParentdag

`func (o *ClearTaskInstance) HasIncludeParentdag() bool`

HasIncludeParentdag returns a boolean if a field has been set.

### GetResetDagRuns

`func (o *ClearTaskInstance) GetResetDagRuns() bool`

GetResetDagRuns returns the ResetDagRuns field if non-nil, zero value otherwise.

### GetResetDagRunsOk

`func (o *ClearTaskInstance) GetResetDagRunsOk() (*bool, bool)`

GetResetDagRunsOk returns a tuple with the ResetDagRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetDagRuns

`func (o *ClearTaskInstance) SetResetDagRuns(v bool)`

SetResetDagRuns sets ResetDagRuns field to given value.

### HasResetDagRuns

`func (o *ClearTaskInstance) HasResetDagRuns() bool`

HasResetDagRuns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


