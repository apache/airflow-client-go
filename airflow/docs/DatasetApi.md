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

# \DatasetApi

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDataset**](DatasetApi.md#GetDataset) | **Get** /datasets/{uri} | Get a dataset
[**GetDatasetEvents**](DatasetApi.md#GetDatasetEvents) | **Get** /datasets/events | Get dataset events
[**GetDatasets**](DatasetApi.md#GetDatasets) | **Get** /datasets | List datasets
[**GetUpstreamDatasetEvents**](DatasetApi.md#GetUpstreamDatasetEvents) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/upstreamDatasetEvents | Get dataset events for a DAG run



## GetDataset

> Dataset GetDataset(ctx, uri).Execute()

Get a dataset



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
    uri := "uri_example" // string | The encoded Dataset URI

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatasetApi.GetDataset(context.Background(), uri).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatasetApi.GetDataset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataset`: Dataset
    fmt.Fprintf(os.Stdout, "Response from `DatasetApi.GetDataset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uri** | **string** | The encoded Dataset URI | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatasetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dataset**](Dataset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatasetEvents

> DatasetEventCollection GetDatasetEvents(ctx).Limit(limit).Offset(offset).OrderBy(orderBy).DatasetId(datasetId).SourceDagId(sourceDagId).SourceTaskId(sourceTaskId).SourceRunId(sourceRunId).SourceMapIndex(sourceMapIndex).Execute()

Get dataset events



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
    datasetId := int32(56) // int32 | The Dataset ID that updated the dataset. (optional)
    sourceDagId := "sourceDagId_example" // string | The DAG ID that updated the dataset. (optional)
    sourceTaskId := "sourceTaskId_example" // string | The task ID that updated the dataset. (optional)
    sourceRunId := "sourceRunId_example" // string | The DAG run ID that updated the dataset. (optional)
    sourceMapIndex := int32(56) // int32 | The map index that updated the dataset. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatasetApi.GetDatasetEvents(context.Background()).Limit(limit).Offset(offset).OrderBy(orderBy).DatasetId(datasetId).SourceDagId(sourceDagId).SourceTaskId(sourceTaskId).SourceRunId(sourceRunId).SourceMapIndex(sourceMapIndex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatasetApi.GetDatasetEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatasetEvents`: DatasetEventCollection
    fmt.Fprintf(os.Stdout, "Response from `DatasetApi.GetDatasetEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDatasetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The numbers of items to return. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **orderBy** | **string** | The name of the field to order the results by. Prefix a field name with &#x60;-&#x60; to reverse the sort order.  *New in version 2.1.0*  | 
 **datasetId** | **int32** | The Dataset ID that updated the dataset. | 
 **sourceDagId** | **string** | The DAG ID that updated the dataset. | 
 **sourceTaskId** | **string** | The task ID that updated the dataset. | 
 **sourceRunId** | **string** | The DAG run ID that updated the dataset. | 
 **sourceMapIndex** | **int32** | The map index that updated the dataset. | 

### Return type

[**DatasetEventCollection**](DatasetEventCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatasets

> DatasetCollection GetDatasets(ctx).Limit(limit).Offset(offset).OrderBy(orderBy).UriPattern(uriPattern).Execute()

List datasets

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
    uriPattern := "uriPattern_example" // string | If set, only return datasets with uris matching this pattern.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatasetApi.GetDatasets(context.Background()).Limit(limit).Offset(offset).OrderBy(orderBy).UriPattern(uriPattern).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatasetApi.GetDatasets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatasets`: DatasetCollection
    fmt.Fprintf(os.Stdout, "Response from `DatasetApi.GetDatasets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDatasetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The numbers of items to return. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **orderBy** | **string** | The name of the field to order the results by. Prefix a field name with &#x60;-&#x60; to reverse the sort order.  *New in version 2.1.0*  | 
 **uriPattern** | **string** | If set, only return datasets with uris matching this pattern.  | 

### Return type

[**DatasetCollection**](DatasetCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpstreamDatasetEvents

> DatasetEventCollection GetUpstreamDatasetEvents(ctx, dagId, dagRunId).Execute()

Get dataset events for a DAG run



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatasetApi.GetUpstreamDatasetEvents(context.Background(), dagId, dagRunId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatasetApi.GetUpstreamDatasetEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUpstreamDatasetEvents`: DatasetEventCollection
    fmt.Fprintf(os.Stdout, "Response from `DatasetApi.GetUpstreamDatasetEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 
**dagRunId** | **string** | The DAG run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUpstreamDatasetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DatasetEventCollection**](DatasetEventCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

