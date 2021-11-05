# \DAGRunApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDagRun**](DAGRunApi.md#DeleteDagRun) | **Delete** /dags/{dag_id}/dagRuns/{dag_run_id} | Delete a DAG run
[**GetDagRun**](DAGRunApi.md#GetDagRun) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id} | Get a DAG run
[**GetDagRuns**](DAGRunApi.md#GetDagRuns) | **Get** /dags/{dag_id}/dagRuns | List DAG runs
[**GetDagRunsBatch**](DAGRunApi.md#GetDagRunsBatch) | **Post** /dags/~/dagRuns/list | List DAG runs (batch)
[**PostDagRun**](DAGRunApi.md#PostDagRun) | **Post** /dags/{dag_id}/dagRuns | Trigger a new DAG run
[**UpdateDagRunState**](DAGRunApi.md#UpdateDagRunState) | **Patch** /dags/{dag_id}/dagRuns/{dag_run_id} | Modify a DAG run



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DAGRunApi.DeleteDagRun(context.Background(), dagId, dagRunId).Execute()
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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DAGRunApi.GetDagRun(context.Background(), dagId, dagRunId).Execute()
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

> DAGRunCollection GetDagRuns(ctx, dagId).Limit(limit).Offset(offset).ExecutionDateGte(executionDateGte).ExecutionDateLte(executionDateLte).StartDateGte(startDateGte).StartDateLte(startDateLte).EndDateGte(endDateGte).EndDateLte(endDateLte).OrderBy(orderBy).Execute()

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
    orderBy := "orderBy_example" // string | The name of the field to order the results by. Prefix a field name with `-` to reverse the sort order.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DAGRunApi.GetDagRuns(context.Background(), dagId).Limit(limit).Offset(offset).ExecutionDateGte(executionDateGte).ExecutionDateLte(executionDateLte).StartDateGte(startDateGte).StartDateLte(startDateLte).EndDateGte(endDateGte).EndDateLte(endDateLte).OrderBy(orderBy).Execute()
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
 **orderBy** | **string** | The name of the field to order the results by. Prefix a field name with &#x60;-&#x60; to reverse the sort order.  | 

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DAGRunApi.GetDagRunsBatch(context.Background()).ListDagRunsForm(listDagRunsForm).Execute()
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


## PostDagRun

> DAGRun PostDagRun(ctx, dagId).DAGRun(dAGRun).Execute()

Trigger a new DAG run

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
    dAGRun := *openapiclient.NewDAGRun("DagId_example") // DAGRun | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DAGRunApi.PostDagRun(context.Background(), dagId).DAGRun(dAGRun).Execute()
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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DAGRunApi.UpdateDagRunState(context.Background(), dagId, dagRunId).UpdateDagRunState(updateDagRunState).Execute()
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

