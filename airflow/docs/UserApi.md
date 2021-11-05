# \UserApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUser**](UserApi.md#DeleteUser) | **Delete** /users/{username} | Delete a user
[**GetUser**](UserApi.md#GetUser) | **Get** /users/{username} | Get a user
[**GetUsers**](UserApi.md#GetUsers) | **Get** /users | List users
[**PatchUser**](UserApi.md#PatchUser) | **Patch** /users/{username} | Update a user
[**PostUser**](UserApi.md#PostUser) | **Post** /users | Create a user



## DeleteUser

> DeleteUser(ctx, username).Execute()

Delete a user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    username := "username_example" // string | The username of the user

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.DeleteUser(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The username of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> UserCollectionItem GetUser(ctx, username).Execute()

Get a user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    username := "username_example" // string | The username of the user

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.GetUser(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: UserCollectionItem
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The username of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserCollectionItem**](UserCollectionItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> UserCollection GetUsers(ctx).Limit(limit).Offset(offset).OrderBy(orderBy).Execute()

List users

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := int32(56) // int32 | The numbers of items to return. (optional) (default to 100)
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set. (optional)
    orderBy := "orderBy_example" // string | The name of the field to order the results by. Prefix a field name with `-` to reverse the sort order.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.GetUsers(context.Background()).Limit(limit).Offset(offset).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsers`: UserCollection
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The numbers of items to return. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **orderBy** | **string** | The name of the field to order the results by. Prefix a field name with &#x60;-&#x60; to reverse the sort order.  | 

### Return type

[**UserCollection**](UserCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchUser

> Role PatchUser(ctx, username).User(user).UpdateMask(updateMask).Execute()

Update a user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    username := "username_example" // string | The username of the user
    user := *openapiclient.NewUser() // User | 
    updateMask := []string{"Inner_example"} // []string | The fields to update on the resource. If absent or empty, all modifiable fields are updated. A comma-separated list of fully qualified names of fields.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.PatchUser(context.Background(), username).User(user).UpdateMask(updateMask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.PatchUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchUser`: Role
    fmt.Fprintf(os.Stdout, "Response from `UserApi.PatchUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The username of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) |  | 
 **updateMask** | **[]string** | The fields to update on the resource. If absent or empty, all modifiable fields are updated. A comma-separated list of fully qualified names of fields.  | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUser

> User PostUser(ctx).User(user).Execute()

Create a user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    user := *openapiclient.NewUser() // User | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.PostUser(context.Background()).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.PostUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostUser`: User
    fmt.Fprintf(os.Stdout, "Response from `UserApi.PostUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**User**](User.md) |  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

