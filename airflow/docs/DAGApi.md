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

# \DAGApi

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDag**](DAGApi.md#DeleteDag) | **Delete** /dags/{dag_id} | Delete a DAG
[**GetDag**](DAGApi.md#GetDag) | **Get** /dags/{dag_id} | Get basic information about a DAG
[**GetDagDetails**](DAGApi.md#GetDagDetails) | **Get** /dags/{dag_id}/details | Get a simplified representation of DAG
[**GetDagSource**](DAGApi.md#GetDagSource) | **Get** /dagSources/{file_token} | Get a source code
[**GetDags**](DAGApi.md#GetDags) | **Get** /dags | List DAGs
[**GetTask**](DAGApi.md#GetTask) | **Get** /dags/{dag_id}/tasks/{task_id} | Get simplified representation of a task
[**GetTasks**](DAGApi.md#GetTasks) | **Get** /dags/{dag_id}/tasks | Get tasks for DAG
[**PatchDag**](DAGApi.md#PatchDag) | **Patch** /dags/{dag_id} | Update a DAG
[**PatchDags**](DAGApi.md#PatchDags) | **Patch** /dags | Update DAGs
[**PostClearTaskInstances**](DAGApi.md#PostClearTaskInstances) | **Post** /dags/{dag_id}/clearTaskInstances | Clear a set of task instances
[**PostSetTaskInstancesState**](DAGApi.md#PostSetTaskInstancesState) | **Post** /dags/{dag_id}/updateTaskInstancesState | Set a state of task instances



## DeleteDag

> DeleteDag(ctx, dagId).Execute()

Delete a DAG



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
    dagId := "dagId_example" // string | The DAG ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DAGApi.DeleteDag(context.Background(), dagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAGApi.DeleteDag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDagRequest struct via the builder pattern


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


## GetDag

> DAG GetDag(ctx, dagId).Execute()

Get basic information about a DAG



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
    dagId := "dagId_example" // string | The DAG ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DAGApi.GetDag(context.Background(), dagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAGApi.GetDag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDag`: DAG
    fmt.Fprintf(os.Stdout, "Response from `DAGApi.GetDag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DAG**](DAG.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDagDetails

> DAGDetail GetDagDetails(ctx, dagId).Execute()

Get a simplified representation of DAG



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
    dagId := "dagId_example" // string | The DAG ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DAGApi.GetDagDetails(context.Background(), dagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAGApi.GetDagDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDagDetails`: DAGDetail
    fmt.Fprintf(os.Stdout, "Response from `DAGApi.GetDagDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDagDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DAGDetail**](DAGDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDagSource

> InlineResponse2001 GetDagSource(ctx, fileToken).Execute()

Get a source code



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
    fileToken := "fileToken_example" // string | The key containing the encrypted path to the file. Encryption and decryption take place only on the server. This prevents the client from reading an non-DAG file. This also ensures API extensibility, because the format of encrypted data may change. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DAGApi.GetDagSource(context.Background(), fileToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAGApi.GetDagSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDagSource`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `DAGApi.GetDagSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileToken** | **string** | The key containing the encrypted path to the file. Encryption and decryption take place only on the server. This prevents the client from reading an non-DAG file. This also ensures API extensibility, because the format of encrypted data may change.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDagSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, plain/text

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDags

> DAGCollection GetDags(ctx).Limit(limit).Offset(offset).OrderBy(orderBy).Tags(tags).OnlyActive(onlyActive).Paused(paused).DagIdPattern(dagIdPattern).Execute()

List DAGs



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
    orderBy := "orderBy_example" // string | The name of the field to order the results by. Prefix a field name with `-` to reverse the sort order.  *New in version 2.1.0*  (optional)
    tags := []string{"Inner_example"} // []string | List of tags to filter results.  *New in version 2.2.0*  (optional)
    onlyActive := true // bool | Only filter active DAGs.  *New in version 2.1.1*  (optional) (default to true)
    paused := true // bool | Only filter paused/unpaused DAGs. If absent or null, it returns paused and unpaused DAGs.  *New in version 2.6.0*  (optional)
    dagIdPattern := "dagIdPattern_example" // string | If set, only return DAGs with dag_ids matching this pattern.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DAGApi.GetDags(context.Background()).Limit(limit).Offset(offset).OrderBy(orderBy).Tags(tags).OnlyActive(onlyActive).Paused(paused).DagIdPattern(dagIdPattern).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAGApi.GetDags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDags`: DAGCollection
    fmt.Fprintf(os.Stdout, "Response from `DAGApi.GetDags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The numbers of items to return. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **orderBy** | **string** | The name of the field to order the results by. Prefix a field name with &#x60;-&#x60; to reverse the sort order.  *New in version 2.1.0*  | 
 **tags** | **[]string** | List of tags to filter results.  *New in version 2.2.0*  | 
 **onlyActive** | **bool** | Only filter active DAGs.  *New in version 2.1.1*  | [default to true]
 **paused** | **bool** | Only filter paused/unpaused DAGs. If absent or null, it returns paused and unpaused DAGs.  *New in version 2.6.0*  | 
 **dagIdPattern** | **string** | If set, only return DAGs with dag_ids matching this pattern.  | 

### Return type

[**DAGCollection**](DAGCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTask

> Task GetTask(ctx, dagId, taskId).Execute()

Get simplified representation of a task

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
    dagId := "dagId_example" // string | The DAG ID.
    taskId := "taskId_example" // string | The task ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DAGApi.GetTask(context.Background(), dagId, taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAGApi.GetTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTask`: Task
    fmt.Fprintf(os.Stdout, "Response from `DAGApi.GetTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 
**taskId** | **string** | The task ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasks

> TaskCollection GetTasks(ctx, dagId).OrderBy(orderBy).Execute()

Get tasks for DAG

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
    dagId := "dagId_example" // string | The DAG ID.
    orderBy := "orderBy_example" // string | The name of the field to order the results by. Prefix a field name with `-` to reverse the sort order.  *New in version 2.1.0*  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DAGApi.GetTasks(context.Background(), dagId).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAGApi.GetTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTasks`: TaskCollection
    fmt.Fprintf(os.Stdout, "Response from `DAGApi.GetTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderBy** | **string** | The name of the field to order the results by. Prefix a field name with &#x60;-&#x60; to reverse the sort order.  *New in version 2.1.0*  | 

### Return type

[**TaskCollection**](TaskCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDag

> DAG PatchDag(ctx, dagId).DAG(dAG).UpdateMask(updateMask).Execute()

Update a DAG

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
    dagId := "dagId_example" // string | The DAG ID.
    dAG := *openapiclient.NewDAG() // DAG | 
    updateMask := []string{"Inner_example"} // []string | The fields to update on the resource. If absent or empty, all modifiable fields are updated. A comma-separated list of fully qualified names of fields.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DAGApi.PatchDag(context.Background(), dagId).DAG(dAG).UpdateMask(updateMask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAGApi.PatchDag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchDag`: DAG
    fmt.Fprintf(os.Stdout, "Response from `DAGApi.PatchDag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dAG** | [**DAG**](DAG.md) |  | 
 **updateMask** | **[]string** | The fields to update on the resource. If absent or empty, all modifiable fields are updated. A comma-separated list of fully qualified names of fields.  | 

### Return type

[**DAG**](DAG.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDags

> DAGCollection PatchDags(ctx).DagIdPattern(dagIdPattern).DAG(dAG).Limit(limit).Offset(offset).Tags(tags).UpdateMask(updateMask).OnlyActive(onlyActive).Execute()

Update DAGs



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
    dagIdPattern := "dagIdPattern_example" // string | If set, only update DAGs with dag_ids matching this pattern. 
    dAG := *openapiclient.NewDAG() // DAG | 
    limit := int32(56) // int32 | The numbers of items to return. (optional) (default to 100)
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set. (optional)
    tags := []string{"Inner_example"} // []string | List of tags to filter results.  *New in version 2.2.0*  (optional)
    updateMask := []string{"Inner_example"} // []string | The fields to update on the resource. If absent or empty, all modifiable fields are updated. A comma-separated list of fully qualified names of fields.  (optional)
    onlyActive := true // bool | Only filter active DAGs.  *New in version 2.1.1*  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DAGApi.PatchDags(context.Background()).DagIdPattern(dagIdPattern).DAG(dAG).Limit(limit).Offset(offset).Tags(tags).UpdateMask(updateMask).OnlyActive(onlyActive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAGApi.PatchDags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchDags`: DAGCollection
    fmt.Fprintf(os.Stdout, "Response from `DAGApi.PatchDags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchDagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dagIdPattern** | **string** | If set, only update DAGs with dag_ids matching this pattern.  | 
 **dAG** | [**DAG**](DAG.md) |  | 
 **limit** | **int32** | The numbers of items to return. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **tags** | **[]string** | List of tags to filter results.  *New in version 2.2.0*  | 
 **updateMask** | **[]string** | The fields to update on the resource. If absent or empty, all modifiable fields are updated. A comma-separated list of fully qualified names of fields.  | 
 **onlyActive** | **bool** | Only filter active DAGs.  *New in version 2.1.1*  | [default to true]

### Return type

[**DAGCollection**](DAGCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClearTaskInstances

> TaskInstanceReferenceCollection PostClearTaskInstances(ctx, dagId).ClearTaskInstances(clearTaskInstances).Execute()

Clear a set of task instances



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
    dagId := "dagId_example" // string | The DAG ID.
    clearTaskInstances := *openapiclient.NewClearTaskInstances() // ClearTaskInstances | Parameters of action

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DAGApi.PostClearTaskInstances(context.Background(), dagId).ClearTaskInstances(clearTaskInstances).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAGApi.PostClearTaskInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostClearTaskInstances`: TaskInstanceReferenceCollection
    fmt.Fprintf(os.Stdout, "Response from `DAGApi.PostClearTaskInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClearTaskInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clearTaskInstances** | [**ClearTaskInstances**](ClearTaskInstances.md) | Parameters of action | 

### Return type

[**TaskInstanceReferenceCollection**](TaskInstanceReferenceCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSetTaskInstancesState

> TaskInstanceReferenceCollection PostSetTaskInstancesState(ctx, dagId).UpdateTaskInstancesState(updateTaskInstancesState).Execute()

Set a state of task instances



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
    dagId := "dagId_example" // string | The DAG ID.
    updateTaskInstancesState := *openapiclient.NewUpdateTaskInstancesState() // UpdateTaskInstancesState | Parameters of action

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DAGApi.PostSetTaskInstancesState(context.Background(), dagId).UpdateTaskInstancesState(updateTaskInstancesState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAGApi.PostSetTaskInstancesState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSetTaskInstancesState`: TaskInstanceReferenceCollection
    fmt.Fprintf(os.Stdout, "Response from `DAGApi.PostSetTaskInstancesState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSetTaskInstancesStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTaskInstancesState** | [**UpdateTaskInstancesState**](UpdateTaskInstancesState.md) | Parameters of action | 

### Return type

[**TaskInstanceReferenceCollection**](TaskInstanceReferenceCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

