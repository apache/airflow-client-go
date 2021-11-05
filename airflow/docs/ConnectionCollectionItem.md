# ConnectionCollectionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | Pointer to **string** | The connection ID. | [optional] 
**ConnType** | Pointer to **string** | The connection type. | [optional] 
**Host** | Pointer to **NullableString** | Host of the connection. | [optional] 
**Login** | Pointer to **NullableString** | Login of the connection. | [optional] 
**Schema** | Pointer to **NullableString** | Schema of the connection. | [optional] 
**Port** | Pointer to **NullableInt32** | Port of the connection. | [optional] 

## Methods

### NewConnectionCollectionItem

`func NewConnectionCollectionItem() *ConnectionCollectionItem`

NewConnectionCollectionItem instantiates a new ConnectionCollectionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionCollectionItemWithDefaults

`func NewConnectionCollectionItemWithDefaults() *ConnectionCollectionItem`

NewConnectionCollectionItemWithDefaults instantiates a new ConnectionCollectionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *ConnectionCollectionItem) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ConnectionCollectionItem) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ConnectionCollectionItem) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *ConnectionCollectionItem) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetConnType

`func (o *ConnectionCollectionItem) GetConnType() string`

GetConnType returns the ConnType field if non-nil, zero value otherwise.

### GetConnTypeOk

`func (o *ConnectionCollectionItem) GetConnTypeOk() (*string, bool)`

GetConnTypeOk returns a tuple with the ConnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnType

`func (o *ConnectionCollectionItem) SetConnType(v string)`

SetConnType sets ConnType field to given value.

### HasConnType

`func (o *ConnectionCollectionItem) HasConnType() bool`

HasConnType returns a boolean if a field has been set.

### GetHost

`func (o *ConnectionCollectionItem) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ConnectionCollectionItem) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ConnectionCollectionItem) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ConnectionCollectionItem) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *ConnectionCollectionItem) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *ConnectionCollectionItem) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetLogin

`func (o *ConnectionCollectionItem) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *ConnectionCollectionItem) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *ConnectionCollectionItem) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *ConnectionCollectionItem) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### SetLoginNil

`func (o *ConnectionCollectionItem) SetLoginNil(b bool)`

 SetLoginNil sets the value for Login to be an explicit nil

### UnsetLogin
`func (o *ConnectionCollectionItem) UnsetLogin()`

UnsetLogin ensures that no value is present for Login, not even an explicit nil
### GetSchema

`func (o *ConnectionCollectionItem) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ConnectionCollectionItem) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ConnectionCollectionItem) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *ConnectionCollectionItem) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *ConnectionCollectionItem) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *ConnectionCollectionItem) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetPort

`func (o *ConnectionCollectionItem) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ConnectionCollectionItem) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ConnectionCollectionItem) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ConnectionCollectionItem) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *ConnectionCollectionItem) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *ConnectionCollectionItem) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


