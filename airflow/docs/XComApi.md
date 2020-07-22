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

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetXcomEntries**](XComApi.md#GetXcomEntries) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/xcomEntries | Get all XCom entries
[**GetXcomEntry**](XComApi.md#GetXcomEntry) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/xcomEntries/{xcom_key} | Get an XCom entry



## GetXcomEntries

> XComCollection GetXcomEntries(ctx, dagId, dagRunId, taskId, optional)

Get all XCom entries

This endpoint allows specifying `~` as the dag_id, dag_run_id, task_id to retrieve XCOM entries for for all DAGs, DAG Runs and task instances.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string**| The DAG ID. | 
**dagRunId** | **string**| The DAG Run ID. | 
**taskId** | **string**| The Task ID. | 
 **optional** | ***GetXcomEntriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetXcomEntriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **optional.Int32**| The numbers of items to return. | [default to 100]
 **offset** | **optional.Int32**| The number of items to skip before starting to collect the result set. | 

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

> XCom GetXcomEntry(ctx, dagId, dagRunId, taskId, xcomKey)

Get an XCom entry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string**| The DAG ID. | 
**dagRunId** | **string**| The DAG Run ID. | 
**taskId** | **string**| The Task ID. | 
**xcomKey** | **string**| The XCom Key. | 

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

