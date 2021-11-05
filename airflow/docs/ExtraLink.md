# ExtraLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassRef** | Pointer to [**ClassReference**](ClassReference.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Href** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewExtraLink

`func NewExtraLink() *ExtraLink`

NewExtraLink instantiates a new ExtraLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtraLinkWithDefaults

`func NewExtraLinkWithDefaults() *ExtraLink`

NewExtraLinkWithDefaults instantiates a new ExtraLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassRef

`func (o *ExtraLink) GetClassRef() ClassReference`

GetClassRef returns the ClassRef field if non-nil, zero value otherwise.

### GetClassRefOk

`func (o *ExtraLink) GetClassRefOk() (*ClassReference, bool)`

GetClassRefOk returns a tuple with the ClassRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassRef

`func (o *ExtraLink) SetClassRef(v ClassReference)`

SetClassRef sets ClassRef field to given value.

### HasClassRef

`func (o *ExtraLink) HasClassRef() bool`

HasClassRef returns a boolean if a field has been set.

### GetName

`func (o *ExtraLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtraLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtraLink) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExtraLink) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHref

`func (o *ExtraLink) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ExtraLink) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ExtraLink) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ExtraLink) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


