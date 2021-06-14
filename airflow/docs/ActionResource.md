# ActionResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**Action**](Action.md) |  | [optional] 
**Resource** | Pointer to [**Resource**](Resource.md) |  | [optional] 

## Methods

### NewActionResource

`func NewActionResource() *ActionResource`

NewActionResource instantiates a new ActionResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionResourceWithDefaults

`func NewActionResourceWithDefaults() *ActionResource`

NewActionResourceWithDefaults instantiates a new ActionResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ActionResource) GetAction() Action`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ActionResource) GetActionOk() (*Action, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ActionResource) SetAction(v Action)`

SetAction sets Action field to given value.

### HasAction

`func (o *ActionResource) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResource

`func (o *ActionResource) GetResource() Resource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ActionResource) GetResourceOk() (*Resource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ActionResource) SetResource(v Resource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ActionResource) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


