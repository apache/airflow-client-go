<!--
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
 -->

# ListDagRunsForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderBy** | Pointer to **string** | The name of the field to order the results by. Prefix a field name with &#x60;-&#x60; to reverse the sort order.  *New in version 2.1.0*  | [optional] 
**PageOffset** | Pointer to **int32** | The number of items to skip before starting to collect the result set. | [optional] 
**PageLimit** | Pointer to **int32** | The numbers of items to return. | [optional] [default to 100]
**DagIds** | Pointer to **[]string** | Return objects with specific DAG IDs. The value can be repeated to retrieve multiple matching values (OR condition). | [optional] 
**States** | Pointer to **[]string** | Return objects with specific states. The value can be repeated to retrieve multiple matching values (OR condition). | [optional] 
**ExecutionDateGte** | Pointer to **time.Time** | Returns objects greater or equal to the specified date.  This can be combined with execution_date_lte key to receive only the selected period.  | [optional] 
**ExecutionDateLte** | Pointer to **time.Time** | Returns objects less than or equal to the specified date.  This can be combined with execution_date_gte key to receive only the selected period.  | [optional] 
**StartDateGte** | Pointer to **time.Time** | Returns objects greater or equal the specified date.  This can be combined with start_date_lte key to receive only the selected period.  | [optional] 
**StartDateLte** | Pointer to **time.Time** | Returns objects less or equal the specified date.  This can be combined with start_date_gte parameter to receive only the selected period  | [optional] 
**EndDateGte** | Pointer to **time.Time** | Returns objects greater or equal the specified date.  This can be combined with end_date_lte parameter to receive only the selected period.  | [optional] 
**EndDateLte** | Pointer to **time.Time** | Returns objects less than or equal to the specified date.  This can be combined with end_date_gte parameter to receive only the selected period.  | [optional] 

## Methods

### NewListDagRunsForm

`func NewListDagRunsForm() *ListDagRunsForm`

NewListDagRunsForm instantiates a new ListDagRunsForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDagRunsFormWithDefaults

`func NewListDagRunsFormWithDefaults() *ListDagRunsForm`

NewListDagRunsFormWithDefaults instantiates a new ListDagRunsForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderBy

`func (o *ListDagRunsForm) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *ListDagRunsForm) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *ListDagRunsForm) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *ListDagRunsForm) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetPageOffset

`func (o *ListDagRunsForm) GetPageOffset() int32`

GetPageOffset returns the PageOffset field if non-nil, zero value otherwise.

### GetPageOffsetOk

`func (o *ListDagRunsForm) GetPageOffsetOk() (*int32, bool)`

GetPageOffsetOk returns a tuple with the PageOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageOffset

`func (o *ListDagRunsForm) SetPageOffset(v int32)`

SetPageOffset sets PageOffset field to given value.

### HasPageOffset

`func (o *ListDagRunsForm) HasPageOffset() bool`

HasPageOffset returns a boolean if a field has been set.

### GetPageLimit

`func (o *ListDagRunsForm) GetPageLimit() int32`

GetPageLimit returns the PageLimit field if non-nil, zero value otherwise.

### GetPageLimitOk

`func (o *ListDagRunsForm) GetPageLimitOk() (*int32, bool)`

GetPageLimitOk returns a tuple with the PageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageLimit

`func (o *ListDagRunsForm) SetPageLimit(v int32)`

SetPageLimit sets PageLimit field to given value.

### HasPageLimit

`func (o *ListDagRunsForm) HasPageLimit() bool`

HasPageLimit returns a boolean if a field has been set.

### GetDagIds

`func (o *ListDagRunsForm) GetDagIds() []string`

GetDagIds returns the DagIds field if non-nil, zero value otherwise.

### GetDagIdsOk

`func (o *ListDagRunsForm) GetDagIdsOk() (*[]string, bool)`

GetDagIdsOk returns a tuple with the DagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagIds

`func (o *ListDagRunsForm) SetDagIds(v []string)`

SetDagIds sets DagIds field to given value.

### HasDagIds

`func (o *ListDagRunsForm) HasDagIds() bool`

HasDagIds returns a boolean if a field has been set.

### GetStates

`func (o *ListDagRunsForm) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *ListDagRunsForm) GetStatesOk() (*[]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *ListDagRunsForm) SetStates(v []string)`

SetStates sets States field to given value.

### HasStates

`func (o *ListDagRunsForm) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetExecutionDateGte

`func (o *ListDagRunsForm) GetExecutionDateGte() time.Time`

GetExecutionDateGte returns the ExecutionDateGte field if non-nil, zero value otherwise.

### GetExecutionDateGteOk

`func (o *ListDagRunsForm) GetExecutionDateGteOk() (*time.Time, bool)`

GetExecutionDateGteOk returns a tuple with the ExecutionDateGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDateGte

`func (o *ListDagRunsForm) SetExecutionDateGte(v time.Time)`

SetExecutionDateGte sets ExecutionDateGte field to given value.

### HasExecutionDateGte

`func (o *ListDagRunsForm) HasExecutionDateGte() bool`

HasExecutionDateGte returns a boolean if a field has been set.

### GetExecutionDateLte

`func (o *ListDagRunsForm) GetExecutionDateLte() time.Time`

GetExecutionDateLte returns the ExecutionDateLte field if non-nil, zero value otherwise.

### GetExecutionDateLteOk

`func (o *ListDagRunsForm) GetExecutionDateLteOk() (*time.Time, bool)`

GetExecutionDateLteOk returns a tuple with the ExecutionDateLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDateLte

`func (o *ListDagRunsForm) SetExecutionDateLte(v time.Time)`

SetExecutionDateLte sets ExecutionDateLte field to given value.

### HasExecutionDateLte

`func (o *ListDagRunsForm) HasExecutionDateLte() bool`

HasExecutionDateLte returns a boolean if a field has been set.

### GetStartDateGte

`func (o *ListDagRunsForm) GetStartDateGte() time.Time`

GetStartDateGte returns the StartDateGte field if non-nil, zero value otherwise.

### GetStartDateGteOk

`func (o *ListDagRunsForm) GetStartDateGteOk() (*time.Time, bool)`

GetStartDateGteOk returns a tuple with the StartDateGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateGte

`func (o *ListDagRunsForm) SetStartDateGte(v time.Time)`

SetStartDateGte sets StartDateGte field to given value.

### HasStartDateGte

`func (o *ListDagRunsForm) HasStartDateGte() bool`

HasStartDateGte returns a boolean if a field has been set.

### GetStartDateLte

`func (o *ListDagRunsForm) GetStartDateLte() time.Time`

GetStartDateLte returns the StartDateLte field if non-nil, zero value otherwise.

### GetStartDateLteOk

`func (o *ListDagRunsForm) GetStartDateLteOk() (*time.Time, bool)`

GetStartDateLteOk returns a tuple with the StartDateLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateLte

`func (o *ListDagRunsForm) SetStartDateLte(v time.Time)`

SetStartDateLte sets StartDateLte field to given value.

### HasStartDateLte

`func (o *ListDagRunsForm) HasStartDateLte() bool`

HasStartDateLte returns a boolean if a field has been set.

### GetEndDateGte

`func (o *ListDagRunsForm) GetEndDateGte() time.Time`

GetEndDateGte returns the EndDateGte field if non-nil, zero value otherwise.

### GetEndDateGteOk

`func (o *ListDagRunsForm) GetEndDateGteOk() (*time.Time, bool)`

GetEndDateGteOk returns a tuple with the EndDateGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateGte

`func (o *ListDagRunsForm) SetEndDateGte(v time.Time)`

SetEndDateGte sets EndDateGte field to given value.

### HasEndDateGte

`func (o *ListDagRunsForm) HasEndDateGte() bool`

HasEndDateGte returns a boolean if a field has been set.

### GetEndDateLte

`func (o *ListDagRunsForm) GetEndDateLte() time.Time`

GetEndDateLte returns the EndDateLte field if non-nil, zero value otherwise.

### GetEndDateLteOk

`func (o *ListDagRunsForm) GetEndDateLteOk() (*time.Time, bool)`

GetEndDateLteOk returns a tuple with the EndDateLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateLte

`func (o *ListDagRunsForm) SetEndDateLte(v time.Time)`

SetEndDateLte sets EndDateLte field to given value.

### HasEndDateLte

`func (o *ListDagRunsForm) HasEndDateLte() bool`

HasEndDateLte returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


