# \XComApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetXcomEntries**](XComApi.md#GetXcomEntries) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/xcomEntries | List XCom entries
[**GetXcomEntry**](XComApi.md#GetXcomEntry) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/xcomEntries/{xcom_key} | Get an XCom entry



## GetXcomEntries

> XComCollection GetXcomEntries(ctx, dagId, dagRunId, taskId).Limit(limit).Offset(offset).Execute()

List XCom entries



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
    dagRunId := "dagRunId_example" // string | The DAG run ID.
    taskId := "taskId_example" // string | The task ID.
    limit := int32(56) // int32 | The numbers of items to return. (optional) (default to 100)
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XComApi.GetXcomEntries(context.Background(), dagId, dagRunId, taskId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XComApi.GetXcomEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetXcomEntries`: XComCollection
    fmt.Fprintf(os.Stdout, "Response from `XComApi.GetXcomEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 
**dagRunId** | **string** | The DAG run ID. | 
**taskId** | **string** | The task ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetXcomEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The numbers of items to return. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 

### Return type

[**XComCollection**](XComCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetXcomEntry

> XCom GetXcomEntry(ctx, dagId, dagRunId, taskId, xcomKey).Execute()

Get an XCom entry

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
    dagRunId := "dagRunId_example" // string | The DAG run ID.
    taskId := "taskId_example" // string | The task ID.
    xcomKey := "xcomKey_example" // string | The XCom key.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XComApi.GetXcomEntry(context.Background(), dagId, dagRunId, taskId, xcomKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XComApi.GetXcomEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetXcomEntry`: XCom
    fmt.Fprintf(os.Stdout, "Response from `XComApi.GetXcomEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 
**dagRunId** | **string** | The DAG run ID. | 
**taskId** | **string** | The task ID. | 
**xcomKey** | **string** | The XCom key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetXcomEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**XCom**](XCom.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

