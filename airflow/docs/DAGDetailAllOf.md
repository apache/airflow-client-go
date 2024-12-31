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

# DAGDetailAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | Pointer to **string** |  | [optional] 
**Catchup** | Pointer to **bool** |  | [optional] [readonly] 
**Orientation** | Pointer to **string** |  | [optional] [readonly] 
**Concurrency** | Pointer to **float32** |  | [optional] [readonly] 
**StartDate** | Pointer to **NullableTime** | The DAG&#39;s start date.  *Changed in version 2.0.1*&amp;#58; Field becomes nullable.  | [optional] [readonly] 
**DagRunTimeout** | Pointer to [**NullableTimeDelta**](TimeDelta.md) |  | [optional] 
**DocMd** | Pointer to **NullableString** |  | [optional] [readonly] 
**DefaultView** | Pointer to **string** |  | [optional] [readonly] 
**Params** | Pointer to **map[string]interface{}** | User-specified DAG params.  *New in version 2.0.1*  | [optional] [readonly] 
**EndDate** | Pointer to **NullableTime** | The DAG&#39;s end date.  *New in version 2.3.0*.  | [optional] [readonly] 
**IsPausedUponCreation** | Pointer to **NullableBool** | Whether the DAG is paused upon creation.  *New in version 2.3.0*  | [optional] [readonly] 
**LastParsed** | Pointer to **NullableTime** | The last time the DAG was parsed.  *New in version 2.3.0*  | [optional] [readonly] 
**TemplateSearchPath** | Pointer to **[]string** | The template search path.  *New in version 2.3.0*  | [optional] 
**RenderTemplateAsNativeObj** | Pointer to **NullableBool** | Whether to render templates as native Python objects.  *New in version 2.3.0*  | [optional] [readonly] 

## Methods

### NewDAGDetailAllOf

`func NewDAGDetailAllOf() *DAGDetailAllOf`

NewDAGDetailAllOf instantiates a new DAGDetailAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDAGDetailAllOfWithDefaults

`func NewDAGDetailAllOfWithDefaults() *DAGDetailAllOf`

NewDAGDetailAllOfWithDefaults instantiates a new DAGDetailAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezone

`func (o *DAGDetailAllOf) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *DAGDetailAllOf) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *DAGDetailAllOf) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *DAGDetailAllOf) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetCatchup

`func (o *DAGDetailAllOf) GetCatchup() bool`

GetCatchup returns the Catchup field if non-nil, zero value otherwise.

### GetCatchupOk

`func (o *DAGDetailAllOf) GetCatchupOk() (*bool, bool)`

GetCatchupOk returns a tuple with the Catchup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatchup

`func (o *DAGDetailAllOf) SetCatchup(v bool)`

SetCatchup sets Catchup field to given value.

### HasCatchup

`func (o *DAGDetailAllOf) HasCatchup() bool`

HasCatchup returns a boolean if a field has been set.

### GetOrientation

`func (o *DAGDetailAllOf) GetOrientation() string`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *DAGDetailAllOf) GetOrientationOk() (*string, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *DAGDetailAllOf) SetOrientation(v string)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *DAGDetailAllOf) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetConcurrency

`func (o *DAGDetailAllOf) GetConcurrency() float32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *DAGDetailAllOf) GetConcurrencyOk() (*float32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *DAGDetailAllOf) SetConcurrency(v float32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *DAGDetailAllOf) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### GetStartDate

`func (o *DAGDetailAllOf) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DAGDetailAllOf) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DAGDetailAllOf) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DAGDetailAllOf) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *DAGDetailAllOf) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *DAGDetailAllOf) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetDagRunTimeout

`func (o *DAGDetailAllOf) GetDagRunTimeout() TimeDelta`

GetDagRunTimeout returns the DagRunTimeout field if non-nil, zero value otherwise.

### GetDagRunTimeoutOk

`func (o *DAGDetailAllOf) GetDagRunTimeoutOk() (*TimeDelta, bool)`

GetDagRunTimeoutOk returns a tuple with the DagRunTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagRunTimeout

`func (o *DAGDetailAllOf) SetDagRunTimeout(v TimeDelta)`

SetDagRunTimeout sets DagRunTimeout field to given value.

### HasDagRunTimeout

`func (o *DAGDetailAllOf) HasDagRunTimeout() bool`

HasDagRunTimeout returns a boolean if a field has been set.

### SetDagRunTimeoutNil

`func (o *DAGDetailAllOf) SetDagRunTimeoutNil(b bool)`

 SetDagRunTimeoutNil sets the value for DagRunTimeout to be an explicit nil

### UnsetDagRunTimeout
`func (o *DAGDetailAllOf) UnsetDagRunTimeout()`

UnsetDagRunTimeout ensures that no value is present for DagRunTimeout, not even an explicit nil
### GetDocMd

`func (o *DAGDetailAllOf) GetDocMd() string`

GetDocMd returns the DocMd field if non-nil, zero value otherwise.

### GetDocMdOk

`func (o *DAGDetailAllOf) GetDocMdOk() (*string, bool)`

GetDocMdOk returns a tuple with the DocMd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocMd

`func (o *DAGDetailAllOf) SetDocMd(v string)`

SetDocMd sets DocMd field to given value.

### HasDocMd

`func (o *DAGDetailAllOf) HasDocMd() bool`

HasDocMd returns a boolean if a field has been set.

### SetDocMdNil

`func (o *DAGDetailAllOf) SetDocMdNil(b bool)`

 SetDocMdNil sets the value for DocMd to be an explicit nil

### UnsetDocMd
`func (o *DAGDetailAllOf) UnsetDocMd()`

UnsetDocMd ensures that no value is present for DocMd, not even an explicit nil
### GetDefaultView

`func (o *DAGDetailAllOf) GetDefaultView() string`

GetDefaultView returns the DefaultView field if non-nil, zero value otherwise.

### GetDefaultViewOk

`func (o *DAGDetailAllOf) GetDefaultViewOk() (*string, bool)`

GetDefaultViewOk returns a tuple with the DefaultView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultView

`func (o *DAGDetailAllOf) SetDefaultView(v string)`

SetDefaultView sets DefaultView field to given value.

### HasDefaultView

`func (o *DAGDetailAllOf) HasDefaultView() bool`

HasDefaultView returns a boolean if a field has been set.

### GetParams

`func (o *DAGDetailAllOf) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *DAGDetailAllOf) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *DAGDetailAllOf) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *DAGDetailAllOf) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetEndDate

`func (o *DAGDetailAllOf) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DAGDetailAllOf) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DAGDetailAllOf) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DAGDetailAllOf) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *DAGDetailAllOf) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *DAGDetailAllOf) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetIsPausedUponCreation

`func (o *DAGDetailAllOf) GetIsPausedUponCreation() bool`

GetIsPausedUponCreation returns the IsPausedUponCreation field if non-nil, zero value otherwise.

### GetIsPausedUponCreationOk

`func (o *DAGDetailAllOf) GetIsPausedUponCreationOk() (*bool, bool)`

GetIsPausedUponCreationOk returns a tuple with the IsPausedUponCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPausedUponCreation

`func (o *DAGDetailAllOf) SetIsPausedUponCreation(v bool)`

SetIsPausedUponCreation sets IsPausedUponCreation field to given value.

### HasIsPausedUponCreation

`func (o *DAGDetailAllOf) HasIsPausedUponCreation() bool`

HasIsPausedUponCreation returns a boolean if a field has been set.

### SetIsPausedUponCreationNil

`func (o *DAGDetailAllOf) SetIsPausedUponCreationNil(b bool)`

 SetIsPausedUponCreationNil sets the value for IsPausedUponCreation to be an explicit nil

### UnsetIsPausedUponCreation
`func (o *DAGDetailAllOf) UnsetIsPausedUponCreation()`

UnsetIsPausedUponCreation ensures that no value is present for IsPausedUponCreation, not even an explicit nil
### GetLastParsed

`func (o *DAGDetailAllOf) GetLastParsed() time.Time`

GetLastParsed returns the LastParsed field if non-nil, zero value otherwise.

### GetLastParsedOk

`func (o *DAGDetailAllOf) GetLastParsedOk() (*time.Time, bool)`

GetLastParsedOk returns a tuple with the LastParsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastParsed

`func (o *DAGDetailAllOf) SetLastParsed(v time.Time)`

SetLastParsed sets LastParsed field to given value.

### HasLastParsed

`func (o *DAGDetailAllOf) HasLastParsed() bool`

HasLastParsed returns a boolean if a field has been set.

### SetLastParsedNil

`func (o *DAGDetailAllOf) SetLastParsedNil(b bool)`

 SetLastParsedNil sets the value for LastParsed to be an explicit nil

### UnsetLastParsed
`func (o *DAGDetailAllOf) UnsetLastParsed()`

UnsetLastParsed ensures that no value is present for LastParsed, not even an explicit nil
### GetTemplateSearchPath

`func (o *DAGDetailAllOf) GetTemplateSearchPath() []string`

GetTemplateSearchPath returns the TemplateSearchPath field if non-nil, zero value otherwise.

### GetTemplateSearchPathOk

`func (o *DAGDetailAllOf) GetTemplateSearchPathOk() (*[]string, bool)`

GetTemplateSearchPathOk returns a tuple with the TemplateSearchPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSearchPath

`func (o *DAGDetailAllOf) SetTemplateSearchPath(v []string)`

SetTemplateSearchPath sets TemplateSearchPath field to given value.

### HasTemplateSearchPath

`func (o *DAGDetailAllOf) HasTemplateSearchPath() bool`

HasTemplateSearchPath returns a boolean if a field has been set.

### SetTemplateSearchPathNil

`func (o *DAGDetailAllOf) SetTemplateSearchPathNil(b bool)`

 SetTemplateSearchPathNil sets the value for TemplateSearchPath to be an explicit nil

### UnsetTemplateSearchPath
`func (o *DAGDetailAllOf) UnsetTemplateSearchPath()`

UnsetTemplateSearchPath ensures that no value is present for TemplateSearchPath, not even an explicit nil
### GetRenderTemplateAsNativeObj

`func (o *DAGDetailAllOf) GetRenderTemplateAsNativeObj() bool`

GetRenderTemplateAsNativeObj returns the RenderTemplateAsNativeObj field if non-nil, zero value otherwise.

### GetRenderTemplateAsNativeObjOk

`func (o *DAGDetailAllOf) GetRenderTemplateAsNativeObjOk() (*bool, bool)`

GetRenderTemplateAsNativeObjOk returns a tuple with the RenderTemplateAsNativeObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderTemplateAsNativeObj

`func (o *DAGDetailAllOf) SetRenderTemplateAsNativeObj(v bool)`

SetRenderTemplateAsNativeObj sets RenderTemplateAsNativeObj field to given value.

### HasRenderTemplateAsNativeObj

`func (o *DAGDetailAllOf) HasRenderTemplateAsNativeObj() bool`

HasRenderTemplateAsNativeObj returns a boolean if a field has been set.

### SetRenderTemplateAsNativeObjNil

`func (o *DAGDetailAllOf) SetRenderTemplateAsNativeObjNil(b bool)`

 SetRenderTemplateAsNativeObjNil sets the value for RenderTemplateAsNativeObj to be an explicit nil

### UnsetRenderTemplateAsNativeObj
`func (o *DAGDetailAllOf) UnsetRenderTemplateAsNativeObj()`

UnsetRenderTemplateAsNativeObj ensures that no value is present for RenderTemplateAsNativeObj, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


