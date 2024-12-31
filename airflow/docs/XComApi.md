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

# \XComApi

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetXcomEntries**](XComApi.md#GetXcomEntries) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/xcomEntries | List XCom entries
[**GetXcomEntry**](XComApi.md#GetXcomEntry) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/xcomEntries/{xcom_key} | Get an XCom entry



## GetXcomEntries

> XComCollection GetXcomEntries(ctx, dagId, dagRunId, taskId).MapIndex(mapIndex).XcomKey(xcomKey).Limit(limit).Offset(offset).Execute()

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
    mapIndex := int32(56) // int32 | Filter on map index for mapped task. (optional)
    xcomKey := "xcomKey_example" // string | Only filter the XCom records which have the provided key. (optional)
    limit := int32(56) // int32 | The numbers of items to return. (optional) (default to 100)
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.XComApi.GetXcomEntries(context.Background(), dagId, dagRunId, taskId).MapIndex(mapIndex).XcomKey(xcomKey).Limit(limit).Offset(offset).Execute()
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



 **mapIndex** | **int32** | Filter on map index for mapped task. | 
 **xcomKey** | **string** | Only filter the XCom records which have the provided key. | 
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

> XCom GetXcomEntry(ctx, dagId, dagRunId, taskId, xcomKey).MapIndex(mapIndex).Deserialize(deserialize).Execute()

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
    mapIndex := int32(56) // int32 | Filter on map index for mapped task. (optional)
    deserialize := true // bool | Whether to deserialize an XCom value when using a custom XCom backend.  The XCom API endpoint calls `orm_deserialize_value` by default since an XCom may contain value that is potentially expensive to deserialize in the web server. Setting this to true overrides the consideration, and calls `deserialize_value` instead.  This parameter is not meaningful when using the default XCom backend.  *New in version 2.4.0*  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.XComApi.GetXcomEntry(context.Background(), dagId, dagRunId, taskId, xcomKey).MapIndex(mapIndex).Deserialize(deserialize).Execute()
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




 **mapIndex** | **int32** | Filter on map index for mapped task. | 
 **deserialize** | **bool** | Whether to deserialize an XCom value when using a custom XCom backend.  The XCom API endpoint calls &#x60;orm_deserialize_value&#x60; by default since an XCom may contain value that is potentially expensive to deserialize in the web server. Setting this to true overrides the consideration, and calls &#x60;deserialize_value&#x60; instead.  This parameter is not meaningful when using the default XCom backend.  *New in version 2.4.0*  | [default to false]

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

