# ClearDagRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If set, don&#39;t actually run this operation. The response will contain a list of task instances planned to be cleaned, but not modified in any way.  | [optional] [default to true]

## Methods

### NewClearDagRun

`func NewClearDagRun() *ClearDagRun`

NewClearDagRun instantiates a new ClearDagRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearDagRunWithDefaults

`func NewClearDagRunWithDefaults() *ClearDagRun`

NewClearDagRunWithDefaults instantiates a new ClearDagRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ClearDagRun) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ClearDagRun) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ClearDagRun) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ClearDagRun) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


