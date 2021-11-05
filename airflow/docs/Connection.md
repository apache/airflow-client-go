# Connection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | Pointer to **string** | The connection ID. | [optional] 
**ConnType** | Pointer to **string** | The connection type. | [optional] 
**Host** | Pointer to **NullableString** | Host of the connection. | [optional] 
**Login** | Pointer to **NullableString** | Login of the connection. | [optional] 
**Schema** | Pointer to **NullableString** | Schema of the connection. | [optional] 
**Port** | Pointer to **NullableInt32** | Port of the connection. | [optional] 
**Password** | Pointer to **string** | Password of the connection. | [optional] 
**Extra** | Pointer to **NullableString** | Other values that cannot be put into another field, e.g. RSA keys. | [optional] 

## Methods

### NewConnection

`func NewConnection() *Connection`

NewConnection instantiates a new Connection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionWithDefaults

`func NewConnectionWithDefaults() *Connection`

NewConnectionWithDefaults instantiates a new Connection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *Connection) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *Connection) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *Connection) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *Connection) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetConnType

`func (o *Connection) GetConnType() string`

GetConnType returns the ConnType field if non-nil, zero value otherwise.

### GetConnTypeOk

`func (o *Connection) GetConnTypeOk() (*string, bool)`

GetConnTypeOk returns a tuple with the ConnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnType

`func (o *Connection) SetConnType(v string)`

SetConnType sets ConnType field to given value.

### HasConnType

`func (o *Connection) HasConnType() bool`

HasConnType returns a boolean if a field has been set.

### GetHost

`func (o *Connection) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Connection) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Connection) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *Connection) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *Connection) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *Connection) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetLogin

`func (o *Connection) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *Connection) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *Connection) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *Connection) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### SetLoginNil

`func (o *Connection) SetLoginNil(b bool)`

 SetLoginNil sets the value for Login to be an explicit nil

### UnsetLogin
`func (o *Connection) UnsetLogin()`

UnsetLogin ensures that no value is present for Login, not even an explicit nil
### GetSchema

`func (o *Connection) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Connection) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Connection) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *Connection) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *Connection) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *Connection) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetPort

`func (o *Connection) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Connection) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Connection) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Connection) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *Connection) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *Connection) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetPassword

`func (o *Connection) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Connection) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Connection) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Connection) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetExtra

`func (o *Connection) GetExtra() string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *Connection) GetExtraOk() (*string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *Connection) SetExtra(v string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *Connection) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### SetExtraNil

`func (o *Connection) SetExtraNil(b bool)`

 SetExtraNil sets the value for Extra to be an explicit nil

### UnsetExtra
`func (o *Connection) UnsetExtra()`

UnsetExtra ensures that no value is present for Extra, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


