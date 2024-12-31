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

# \DAGRunApi

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearDagRun**](DAGRunApi.md#ClearDagRun) | **Post** /dags/{dag_id}/dagRuns/{dag_run_id}/clear | Clear a DAG run
[**DeleteDagRun**](DAGRunApi.md#DeleteDagRun) | **Delete** /dags/{dag_id}/dagRuns/{dag_run_id} | Delete a DAG run
[**GetDagRun**](DAGRunApi.md#GetDagRun) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id} | Get a DAG run
[**GetDagRuns**](DAGRunApi.md#GetDagRuns) | **Get** /dags/{dag_id}/dagRuns | List DAG runs
[**GetDagRunsBatch**](DAGRunApi.md#GetDagRunsBatch) | **Post** /dags/~/dagRuns/list | List DAG runs (batch)
[**GetUpstreamDatasetEvents**](DAGRunApi.md#GetUpstreamDatasetEvents) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/upstreamDatasetEvents | Get dataset events for a DAG run
[**PostDagRun**](DAGRunApi.md#PostDagRun) | **Post** /dags/{dag_id}/dagRuns | Trigger a new DAG run.
[**SetDagRunNote**](DAGRunApi.md#SetDagRunNote) | **Patch** /dags/{dag_id}/dagRuns/{dag_run_id}/setNote | Update the DagRun note.
[**UpdateDagRunState**](DAGRunApi.md#UpdateDagRunState) | **Patch** /dags/{dag_id}/dagRuns/{dag_run_id} | Modify a DAG run



## ClearDagRun

> AnyOfDAGRunTaskInstanceCollection ClearDagRun(ctx, dagId, dagRunId).ClearDagRun(clearDagRun).Execute()

Clear a DAG run



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
    clearDagRun := *openapiclient.NewClearDagRun() // ClearDagRun | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DAGRunApi.ClearDagRun(context.Background(), dagId, dagRunId).ClearDagRun(clearDagRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAGRunApi.ClearDagRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClearDagRun`: AnyOfDAGRunTaskInstanceCollection
    fmt.Fprintf(os.Stdout, "Response from `DAGRunApi.ClearDagRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 
**dagRunId** | **string** | The DAG run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearDagRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clearDagRun** | [**ClearDagRun**](ClearDagRun.md) |  | 

### Return type

[**AnyOfDAGRunTaskInstanceCollection**](anyOf&lt;DAGRun,TaskInstanceCollection&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDagRun

> DeleteDagRun(ctx, dagId, dagRunId).Execute()

Delete a DAG run

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
    resp, r, err := apiClient.DAGRunApi.DeleteDagRun(context.Background(), dagId, dagRunId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAGRunApi.DeleteDagRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 
**dagRunId** | **string** | The DAG run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDagRunRequest struct via the builder pattern


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


## GetDagRun

> DAGRun GetDagRun(ctx, dagId, dagRunId).Execute()

Get a DAG run

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
    resp, r, err := apiClient.DAGRunApi.GetDagRun(context.Background(), dagId, dagRunId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAGRunApi.GetDagRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDagRun`: DAGRun
    fmt.Fprintf(os.Stdout, "Response from `DAGRunApi.GetDagRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 
**dagRunId** | **string** | The DAG run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDagRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DAGRun**](DAGRun.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDagRuns

> DAGRunCollection GetDagRuns(ctx, dagId).Limit(limit).Offset(offset).ExecutionDateGte(executionDateGte).ExecutionDateLte(executionDateLte).StartDateGte(startDateGte).StartDateLte(startDateLte).EndDateGte(endDateGte).EndDateLte(endDateLte).UpdatedAtGte(updatedAtGte).UpdatedAtLte(updatedAtLte).State(state).OrderBy(orderBy).Execute()

List DAG runs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    dagId := "dagId_example" // string | The DAG ID.
    limit := int32(56) // int32 | The numbers of items to return. (optional) (default to 100)
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set. (optional)
    executionDateGte := time.Now() // time.Time | Returns objects greater or equal to the specified date.  This can be combined with execution_date_lte parameter to receive only the selected period.  (optional)
    executionDateLte := time.Now() // time.Time | Returns objects less than or equal to the specified date.  This can be combined with execution_date_gte parameter to receive only the selected period.  (optional)
    startDateGte := time.Now() // time.Time | Returns objects greater or equal the specified date.  This can be combined with start_date_lte parameter to receive only the selected period.  (optional)
    startDateLte := time.Now() // time.Time | Returns objects less or equal the specified date.  This can be combined with start_date_gte parameter to receive only the selected period.  (optional)
    endDateGte := time.Now() // time.Time | Returns objects greater or equal the specified date.  This can be combined with start_date_lte parameter to receive only the selected period.  (optional)
    endDateLte := time.Now() // time.Time | Returns objects less than or equal to the specified date.  This can be combined with start_date_gte parameter to receive only the selected period.  (optional)
    updatedAtGte := time.Now() // time.Time | Returns objects greater or equal the specified date.  This can be combined with updated_at_lte parameter to receive only the selected period.  *New in version 2.6.0*  (optional)
    updatedAtLte := time.Now() // time.Time | Returns objects less or equal the specified date.  This can be combined with updated_at_gte parameter to receive only the selected period.  *New in version 2.6.0*  (optional)
    state := []string{"Inner_example"} // []string | The value can be repeated to retrieve multiple matching values (OR condition). (optional)
    orderBy := "orderBy_example" // string | The name of the field to order the results by. Prefix a field name with `-` to reverse the sort order.  *New in version 2.1.0*  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DAGRunApi.GetDagRuns(context.Background(), dagId).Limit(limit).Offset(offset).ExecutionDateGte(executionDateGte).ExecutionDateLte(executionDateLte).StartDateGte(startDateGte).StartDateLte(startDateLte).EndDateGte(endDateGte).EndDateLte(endDateLte).UpdatedAtGte(updatedAtGte).UpdatedAtLte(updatedAtLte).State(state).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAGRunApi.GetDagRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDagRuns`: DAGRunCollection
    fmt.Fprintf(os.Stdout, "Response from `DAGRunApi.GetDagRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDagRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The numbers of items to return. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **executionDateGte** | **time.Time** | Returns objects greater or equal to the specified date.  This can be combined with execution_date_lte parameter to receive only the selected period.  | 
 **executionDateLte** | **time.Time** | Returns objects less than or equal to the specified date.  This can be combined with execution_date_gte parameter to receive only the selected period.  | 
 **startDateGte** | **time.Time** | Returns objects greater or equal the specified date.  This can be combined with start_date_lte parameter to receive only the selected period.  | 
 **startDateLte** | **time.Time** | Returns objects less or equal the specified date.  This can be combined with start_date_gte parameter to receive only the selected period.  | 
 **endDateGte** | **time.Time** | Returns objects greater or equal the specified date.  This can be combined with start_date_lte parameter to receive only the selected period.  | 
 **endDateLte** | **time.Time** | Returns objects less than or equal to the specified date.  This can be combined with start_date_gte parameter to receive only the selected period.  | 
 **updatedAtGte** | **time.Time** | Returns objects greater or equal the specified date.  This can be combined with updated_at_lte parameter to receive only the selected period.  *New in version 2.6.0*  | 
 **updatedAtLte** | **time.Time** | Returns objects less or equal the specified date.  This can be combined with updated_at_gte parameter to receive only the selected period.  *New in version 2.6.0*  | 
 **state** | **[]string** | The value can be repeated to retrieve multiple matching values (OR condition). | 
 **orderBy** | **string** | The name of the field to order the results by. Prefix a field name with &#x60;-&#x60; to reverse the sort order.  *New in version 2.1.0*  | 

### Return type

[**DAGRunCollection**](DAGRunCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDagRunsBatch

> DAGRunCollection GetDagRunsBatch(ctx).ListDagRunsForm(listDagRunsForm).Execute()

List DAG runs (batch)



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
    listDagRunsForm := *openapiclient.NewListDagRunsForm() // ListDagRunsForm | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DAGRunApi.GetDagRunsBatch(context.Background()).ListDagRunsForm(listDagRunsForm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAGRunApi.GetDagRunsBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDagRunsBatch`: DAGRunCollection
    fmt.Fprintf(os.Stdout, "Response from `DAGRunApi.GetDagRunsBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDagRunsBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listDagRunsForm** | [**ListDagRunsForm**](ListDagRunsForm.md) |  | 

### Return type

[**DAGRunCollection**](DAGRunCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
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
    resp, r, err := apiClient.DAGRunApi.GetUpstreamDatasetEvents(context.Background(), dagId, dagRunId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAGRunApi.GetUpstreamDatasetEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUpstreamDatasetEvents`: DatasetEventCollection
    fmt.Fprintf(os.Stdout, "Response from `DAGRunApi.GetUpstreamDatasetEvents`: %v\n", resp)
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


## PostDagRun

> DAGRun PostDagRun(ctx, dagId).DAGRun(dAGRun).Execute()

Trigger a new DAG run.



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
    dAGRun := *openapiclient.NewDAGRun() // DAGRun | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DAGRunApi.PostDagRun(context.Background(), dagId).DAGRun(dAGRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAGRunApi.PostDagRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDagRun`: DAGRun
    fmt.Fprintf(os.Stdout, "Response from `DAGRunApi.PostDagRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDagRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dAGRun** | [**DAGRun**](DAGRun.md) |  | 

### Return type

[**DAGRun**](DAGRun.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDagRunNote

> DAGRun SetDagRunNote(ctx, dagId, dagRunId).SetDagRunNote(setDagRunNote).Execute()

Update the DagRun note.



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
    setDagRunNote := *openapiclient.NewSetDagRunNote() // SetDagRunNote | Parameters of set DagRun note.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DAGRunApi.SetDagRunNote(context.Background(), dagId, dagRunId).SetDagRunNote(setDagRunNote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAGRunApi.SetDagRunNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetDagRunNote`: DAGRun
    fmt.Fprintf(os.Stdout, "Response from `DAGRunApi.SetDagRunNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 
**dagRunId** | **string** | The DAG run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDagRunNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setDagRunNote** | [**SetDagRunNote**](SetDagRunNote.md) | Parameters of set DagRun note. | 

### Return type

[**DAGRun**](DAGRun.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDagRunState

> DAGRun UpdateDagRunState(ctx, dagId, dagRunId).UpdateDagRunState(updateDagRunState).Execute()

Modify a DAG run



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
    updateDagRunState := *openapiclient.NewUpdateDagRunState() // UpdateDagRunState | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DAGRunApi.UpdateDagRunState(context.Background(), dagId, dagRunId).UpdateDagRunState(updateDagRunState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAGRunApi.UpdateDagRunState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDagRunState`: DAGRun
    fmt.Fprintf(os.Stdout, "Response from `DAGRunApi.UpdateDagRunState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 
**dagRunId** | **string** | The DAG run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDagRunStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateDagRunState** | [**UpdateDagRunState**](UpdateDagRunState.md) |  | 

### Return type

[**DAGRun**](DAGRun.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

