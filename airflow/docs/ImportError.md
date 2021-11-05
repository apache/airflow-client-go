# ImportError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportErrorId** | Pointer to **int32** | The import error ID. | [optional] [readonly] 
**Timestamp** | Pointer to **string** | The time when this error was created. | [optional] [readonly] 
**Filename** | Pointer to **string** | The filename | [optional] [readonly] 
**StackTrace** | Pointer to **string** | The full stackstrace.. | [optional] [readonly] 

## Methods

### NewImportError

`func NewImportError() *ImportError`

NewImportError instantiates a new ImportError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportErrorWithDefaults

`func NewImportErrorWithDefaults() *ImportError`

NewImportErrorWithDefaults instantiates a new ImportError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImportErrorId

`func (o *ImportError) GetImportErrorId() int32`

GetImportErrorId returns the ImportErrorId field if non-nil, zero value otherwise.

### GetImportErrorIdOk

`func (o *ImportError) GetImportErrorIdOk() (*int32, bool)`

GetImportErrorIdOk returns a tuple with the ImportErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportErrorId

`func (o *ImportError) SetImportErrorId(v int32)`

SetImportErrorId sets ImportErrorId field to given value.

### HasImportErrorId

`func (o *ImportError) HasImportErrorId() bool`

HasImportErrorId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ImportError) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ImportError) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ImportError) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ImportError) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetFilename

`func (o *ImportError) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ImportError) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ImportError) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *ImportError) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetStackTrace

`func (o *ImportError) GetStackTrace() string`

GetStackTrace returns the StackTrace field if non-nil, zero value otherwise.

### GetStackTraceOk

`func (o *ImportError) GetStackTraceOk() (*string, bool)`

GetStackTraceOk returns a tuple with the StackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTrace

`func (o *ImportError) SetStackTrace(v string)`

SetStackTrace sets StackTrace field to given value.

### HasStackTrace

`func (o *ImportError) HasStackTrace() bool`

HasStackTrace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


