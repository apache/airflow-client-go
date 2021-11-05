# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassRef** | Pointer to [**ClassReference**](ClassReference.md) |  | [optional] 
**TaskId** | Pointer to **string** |  | [optional] [readonly] 
**Owner** | Pointer to **string** |  | [optional] [readonly] 
**StartDate** | Pointer to **time.Time** |  | [optional] [readonly] 
**EndDate** | Pointer to **NullableTime** |  | [optional] [readonly] 
**TriggerRule** | Pointer to [**TriggerRule**](TriggerRule.md) |  | [optional] 
**ExtraLinks** | Pointer to [**[]TaskExtraLinks**](TaskExtraLinks.md) |  | [optional] [readonly] 
**DependsOnPast** | Pointer to **bool** |  | [optional] [readonly] 
**WaitForDownstream** | Pointer to **bool** |  | [optional] [readonly] 
**Retries** | Pointer to **float32** |  | [optional] [readonly] 
**Queue** | Pointer to **string** |  | [optional] [readonly] 
**Pool** | Pointer to **string** |  | [optional] [readonly] 
**PoolSlots** | Pointer to **float32** |  | [optional] [readonly] 
**ExecutionTimeout** | Pointer to [**TimeDelta**](TimeDelta.md) |  | [optional] 
**RetryDelay** | Pointer to [**TimeDelta**](TimeDelta.md) |  | [optional] 
**RetryExponentialBackoff** | Pointer to **bool** |  | [optional] [readonly] 
**PriorityWeight** | Pointer to **float32** |  | [optional] [readonly] 
**WeightRule** | Pointer to [**WeightRule**](WeightRule.md) |  | [optional] 
**UiColor** | Pointer to **string** | Color in hexadecimal notation. | [optional] 
**UiFgcolor** | Pointer to **string** | Color in hexadecimal notation. | [optional] 
**TemplateFields** | Pointer to **[]string** |  | [optional] [readonly] 
**SubDag** | Pointer to [**DAG**](DAG.md) |  | [optional] 
**DownstreamTaskIds** | Pointer to **[]string** |  | [optional] [readonly] 

## Methods

### NewTask

`func NewTask() *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassRef

`func (o *Task) GetClassRef() ClassReference`

GetClassRef returns the ClassRef field if non-nil, zero value otherwise.

### GetClassRefOk

`func (o *Task) GetClassRefOk() (*ClassReference, bool)`

GetClassRefOk returns a tuple with the ClassRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassRef

`func (o *Task) SetClassRef(v ClassReference)`

SetClassRef sets ClassRef field to given value.

### HasClassRef

`func (o *Task) HasClassRef() bool`

HasClassRef returns a boolean if a field has been set.

### GetTaskId

`func (o *Task) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *Task) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *Task) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *Task) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetOwner

`func (o *Task) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Task) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Task) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Task) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetStartDate

`func (o *Task) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Task) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Task) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Task) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *Task) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Task) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Task) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Task) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *Task) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *Task) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetTriggerRule

`func (o *Task) GetTriggerRule() TriggerRule`

GetTriggerRule returns the TriggerRule field if non-nil, zero value otherwise.

### GetTriggerRuleOk

`func (o *Task) GetTriggerRuleOk() (*TriggerRule, bool)`

GetTriggerRuleOk returns a tuple with the TriggerRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerRule

`func (o *Task) SetTriggerRule(v TriggerRule)`

SetTriggerRule sets TriggerRule field to given value.

### HasTriggerRule

`func (o *Task) HasTriggerRule() bool`

HasTriggerRule returns a boolean if a field has been set.

### GetExtraLinks

`func (o *Task) GetExtraLinks() []TaskExtraLinks`

GetExtraLinks returns the ExtraLinks field if non-nil, zero value otherwise.

### GetExtraLinksOk

`func (o *Task) GetExtraLinksOk() (*[]TaskExtraLinks, bool)`

GetExtraLinksOk returns a tuple with the ExtraLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraLinks

`func (o *Task) SetExtraLinks(v []TaskExtraLinks)`

SetExtraLinks sets ExtraLinks field to given value.

### HasExtraLinks

`func (o *Task) HasExtraLinks() bool`

HasExtraLinks returns a boolean if a field has been set.

### GetDependsOnPast

`func (o *Task) GetDependsOnPast() bool`

GetDependsOnPast returns the DependsOnPast field if non-nil, zero value otherwise.

### GetDependsOnPastOk

`func (o *Task) GetDependsOnPastOk() (*bool, bool)`

GetDependsOnPastOk returns a tuple with the DependsOnPast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnPast

`func (o *Task) SetDependsOnPast(v bool)`

SetDependsOnPast sets DependsOnPast field to given value.

### HasDependsOnPast

`func (o *Task) HasDependsOnPast() bool`

HasDependsOnPast returns a boolean if a field has been set.

### GetWaitForDownstream

`func (o *Task) GetWaitForDownstream() bool`

GetWaitForDownstream returns the WaitForDownstream field if non-nil, zero value otherwise.

### GetWaitForDownstreamOk

`func (o *Task) GetWaitForDownstreamOk() (*bool, bool)`

GetWaitForDownstreamOk returns a tuple with the WaitForDownstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForDownstream

`func (o *Task) SetWaitForDownstream(v bool)`

SetWaitForDownstream sets WaitForDownstream field to given value.

### HasWaitForDownstream

`func (o *Task) HasWaitForDownstream() bool`

HasWaitForDownstream returns a boolean if a field has been set.

### GetRetries

`func (o *Task) GetRetries() float32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *Task) GetRetriesOk() (*float32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *Task) SetRetries(v float32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *Task) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetQueue

`func (o *Task) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *Task) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *Task) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *Task) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### GetPool

`func (o *Task) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *Task) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *Task) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *Task) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolSlots

`func (o *Task) GetPoolSlots() float32`

GetPoolSlots returns the PoolSlots field if non-nil, zero value otherwise.

### GetPoolSlotsOk

`func (o *Task) GetPoolSlotsOk() (*float32, bool)`

GetPoolSlotsOk returns a tuple with the PoolSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSlots

`func (o *Task) SetPoolSlots(v float32)`

SetPoolSlots sets PoolSlots field to given value.

### HasPoolSlots

`func (o *Task) HasPoolSlots() bool`

HasPoolSlots returns a boolean if a field has been set.

### GetExecutionTimeout

`func (o *Task) GetExecutionTimeout() TimeDelta`

GetExecutionTimeout returns the ExecutionTimeout field if non-nil, zero value otherwise.

### GetExecutionTimeoutOk

`func (o *Task) GetExecutionTimeoutOk() (*TimeDelta, bool)`

GetExecutionTimeoutOk returns a tuple with the ExecutionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTimeout

`func (o *Task) SetExecutionTimeout(v TimeDelta)`

SetExecutionTimeout sets ExecutionTimeout field to given value.

### HasExecutionTimeout

`func (o *Task) HasExecutionTimeout() bool`

HasExecutionTimeout returns a boolean if a field has been set.

### GetRetryDelay

`func (o *Task) GetRetryDelay() TimeDelta`

GetRetryDelay returns the RetryDelay field if non-nil, zero value otherwise.

### GetRetryDelayOk

`func (o *Task) GetRetryDelayOk() (*TimeDelta, bool)`

GetRetryDelayOk returns a tuple with the RetryDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelay

`func (o *Task) SetRetryDelay(v TimeDelta)`

SetRetryDelay sets RetryDelay field to given value.

### HasRetryDelay

`func (o *Task) HasRetryDelay() bool`

HasRetryDelay returns a boolean if a field has been set.

### GetRetryExponentialBackoff

`func (o *Task) GetRetryExponentialBackoff() bool`

GetRetryExponentialBackoff returns the RetryExponentialBackoff field if non-nil, zero value otherwise.

### GetRetryExponentialBackoffOk

`func (o *Task) GetRetryExponentialBackoffOk() (*bool, bool)`

GetRetryExponentialBackoffOk returns a tuple with the RetryExponentialBackoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryExponentialBackoff

`func (o *Task) SetRetryExponentialBackoff(v bool)`

SetRetryExponentialBackoff sets RetryExponentialBackoff field to given value.

### HasRetryExponentialBackoff

`func (o *Task) HasRetryExponentialBackoff() bool`

HasRetryExponentialBackoff returns a boolean if a field has been set.

### GetPriorityWeight

`func (o *Task) GetPriorityWeight() float32`

GetPriorityWeight returns the PriorityWeight field if non-nil, zero value otherwise.

### GetPriorityWeightOk

`func (o *Task) GetPriorityWeightOk() (*float32, bool)`

GetPriorityWeightOk returns a tuple with the PriorityWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityWeight

`func (o *Task) SetPriorityWeight(v float32)`

SetPriorityWeight sets PriorityWeight field to given value.

### HasPriorityWeight

`func (o *Task) HasPriorityWeight() bool`

HasPriorityWeight returns a boolean if a field has been set.

### GetWeightRule

`func (o *Task) GetWeightRule() WeightRule`

GetWeightRule returns the WeightRule field if non-nil, zero value otherwise.

### GetWeightRuleOk

`func (o *Task) GetWeightRuleOk() (*WeightRule, bool)`

GetWeightRuleOk returns a tuple with the WeightRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightRule

`func (o *Task) SetWeightRule(v WeightRule)`

SetWeightRule sets WeightRule field to given value.

### HasWeightRule

`func (o *Task) HasWeightRule() bool`

HasWeightRule returns a boolean if a field has been set.

### GetUiColor

`func (o *Task) GetUiColor() string`

GetUiColor returns the UiColor field if non-nil, zero value otherwise.

### GetUiColorOk

`func (o *Task) GetUiColorOk() (*string, bool)`

GetUiColorOk returns a tuple with the UiColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiColor

`func (o *Task) SetUiColor(v string)`

SetUiColor sets UiColor field to given value.

### HasUiColor

`func (o *Task) HasUiColor() bool`

HasUiColor returns a boolean if a field has been set.

### GetUiFgcolor

`func (o *Task) GetUiFgcolor() string`

GetUiFgcolor returns the UiFgcolor field if non-nil, zero value otherwise.

### GetUiFgcolorOk

`func (o *Task) GetUiFgcolorOk() (*string, bool)`

GetUiFgcolorOk returns a tuple with the UiFgcolor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiFgcolor

`func (o *Task) SetUiFgcolor(v string)`

SetUiFgcolor sets UiFgcolor field to given value.

### HasUiFgcolor

`func (o *Task) HasUiFgcolor() bool`

HasUiFgcolor returns a boolean if a field has been set.

### GetTemplateFields

`func (o *Task) GetTemplateFields() []string`

GetTemplateFields returns the TemplateFields field if non-nil, zero value otherwise.

### GetTemplateFieldsOk

`func (o *Task) GetTemplateFieldsOk() (*[]string, bool)`

GetTemplateFieldsOk returns a tuple with the TemplateFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateFields

`func (o *Task) SetTemplateFields(v []string)`

SetTemplateFields sets TemplateFields field to given value.

### HasTemplateFields

`func (o *Task) HasTemplateFields() bool`

HasTemplateFields returns a boolean if a field has been set.

### GetSubDag

`func (o *Task) GetSubDag() DAG`

GetSubDag returns the SubDag field if non-nil, zero value otherwise.

### GetSubDagOk

`func (o *Task) GetSubDagOk() (*DAG, bool)`

GetSubDagOk returns a tuple with the SubDag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubDag

`func (o *Task) SetSubDag(v DAG)`

SetSubDag sets SubDag field to given value.

### HasSubDag

`func (o *Task) HasSubDag() bool`

HasSubDag returns a boolean if a field has been set.

### GetDownstreamTaskIds

`func (o *Task) GetDownstreamTaskIds() []string`

GetDownstreamTaskIds returns the DownstreamTaskIds field if non-nil, zero value otherwise.

### GetDownstreamTaskIdsOk

`func (o *Task) GetDownstreamTaskIdsOk() (*[]string, bool)`

GetDownstreamTaskIdsOk returns a tuple with the DownstreamTaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstreamTaskIds

`func (o *Task) SetDownstreamTaskIds(v []string)`

SetDownstreamTaskIds sets DownstreamTaskIds field to given value.

### HasDownstreamTaskIds

`func (o *Task) HasDownstreamTaskIds() bool`

HasDownstreamTaskIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


