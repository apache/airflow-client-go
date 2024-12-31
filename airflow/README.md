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

# Go API client for airflow

# Overview

To facilitate management, Apache Airflow supports a range of REST API endpoints across its
objects.
This section provides an overview of the API design, methods, and supported use cases.

Most of the endpoints accept `JSON` as input and return `JSON` responses.
This means that you must usually add the following headers to your request:
```
Content-type: application/json
Accept: application/json
```

## Resources

The term `resource` refers to a single type of object in the Airflow metadata. An API is broken up by its
endpoint's corresponding resource.
The name of a resource is typically plural and expressed in camelCase. Example: `dagRuns`.

Resource names are used as part of endpoint URLs, as well as in API parameters and responses.

## CRUD Operations

The platform supports **C**reate, **R**ead, **U**pdate, and **D**elete operations on most resources.
You can review the standards for these operations and their standard parameters below.

Some endpoints have special behavior as exceptions.

### Create

To create a resource, you typically submit an HTTP `POST` request with the resource's required metadata
in the request body.
The response returns a `201 Created` response code upon success with the resource's metadata, including
its internal `id`, in the response body.

### Read

The HTTP `GET` request can be used to read a resource or to list a number of resources.

A resource's `id` can be submitted in the request parameters to read a specific resource.
The response usually returns a `200 OK` response code upon success, with the resource's metadata in
the response body.

If a `GET` request does not include a specific resource `id`, it is treated as a list request.
The response usually returns a `200 OK` response code upon success, with an object containing a list
of resources' metadata in the response body.

When reading resources, some common query parameters are usually available. e.g.:
```
v1/connections?limit=25&offset=25
```

|Query Parameter|Type|Description|
|---------------|----|-----------|
|limit|integer|Maximum number of objects to fetch. Usually 25 by default|
|offset|integer|Offset after which to start returning objects. For use with limit query parameter.|

### Update

Updating a resource requires the resource `id`, and is typically done using an HTTP `PATCH` request,
with the fields to modify in the request body.
The response usually returns a `200 OK` response code upon success, with information about the modified
resource in the response body.

### Delete

Deleting a resource requires the resource `id` and is typically executed via an HTTP `DELETE` request.
The response usually returns a `204 No Content` response code upon success.

## Conventions

- Resource names are plural and expressed in camelCase.
- Names are consistent between URL parameter name and field name.

- Field names are in snake_case.
```json
{
    \"description\": \"string\",
    \"name\": \"string\",
    \"occupied_slots\": 0,
    \"open_slots\": 0
    \"queued_slots\": 0,
    \"running_slots\": 0,
    \"scheduled_slots\": 0,
    \"slots\": 0,
}
```

### Update Mask

Update mask is available as a query parameter in patch endpoints. It is used to notify the
API which fields you want to update. Using `update_mask` makes it easier to update objects
by helping the server know which fields to update in an object instead of updating all fields.
The update request ignores any fields that aren't specified in the field mask, leaving them with
their current values.

Example:
```
  resource = request.get('/resource/my-id').json()
  resource['my_field'] = 'new-value'
  request.patch('/resource/my-id?update_mask=my_field', data=json.dumps(resource))
```

## Versioning and Endpoint Lifecycle

- API versioning is not synchronized to specific releases of the Apache Airflow.
- APIs are designed to be backward compatible.
- Any changes to the API will first go through a deprecation phase.

# Trying the API

You can use a third party client, such as [curl](https://curl.haxx.se/), [HTTPie](https://httpie.org/),
[Postman](https://www.postman.com/) or [the Insomnia rest client](https://insomnia.rest/) to test
the Apache Airflow API.

Note that you will need to pass credentials data.

For e.g., here is how to pause a DAG with [curl](https://curl.haxx.se/), when basic authorization is used:
```bash
curl -X PATCH 'https://example.com/api/v1/dags/{dag_id}?update_mask=is_paused' \\
-H 'Content-Type: application/json' \\
--user \"username:password\" \\
-d '{
    \"is_paused\": true
}'
```

Using a graphical tool such as [Postman](https://www.postman.com/) or [Insomnia](https://insomnia.rest/),
it is possible to import the API specifications directly:

1. Download the API specification by clicking the **Download** button at the top of this document
2. Import the JSON specification in the graphical tool of your choice.
  - In *Postman*, you can click the **import** button at the top
  - With *Insomnia*, you can just drag-and-drop the file on the UI

Note that with *Postman*, you can also generate code snippets by selecting a request and clicking on
the **Code** button.

## Enabling CORS

[Cross-origin resource sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS)
is a browser security feature that restricts HTTP requests that are
initiated from scripts running in the browser.

For details on enabling/configuring CORS, see
[Enabling CORS](https://airflow.apache.org/docs/apache-airflow/stable/security/api.html).

# Authentication

To be able to meet the requirements of many organizations, Airflow supports many authentication methods,
and it is even possible to add your own method.

If you want to check which auth backend is currently set, you can use
`airflow config get-value api auth_backends` command as in the example below.
```bash
$ airflow config get-value api auth_backends
airflow.api.auth.backend.basic_auth
```
The default is to deny all requests.

For details on configuring the authentication, see
[API Authorization](https://airflow.apache.org/docs/apache-airflow/stable/security/api.html).

# Errors

We follow the error response format proposed in [RFC 7807](https://tools.ietf.org/html/rfc7807)
also known as Problem Details for HTTP APIs. As with our normal API responses,
your client must be prepared to gracefully handle additional members of the response.

## Unauthenticated

This indicates that the request has not been applied because it lacks valid authentication
credentials for the target resource. Please check that you have valid credentials.

## PermissionDenied

This response means that the server understood the request but refuses to authorize
it because it lacks sufficient rights to the resource. It happens when you do not have the
necessary permission to execute the action you performed. You need to get the appropriate
permissions in other to resolve this error.

## BadRequest

This response means that the server cannot or will not process the request due to something
that is perceived to be a client error (e.g., malformed request syntax, invalid request message
framing, or deceptive request routing). To resolve this, please ensure that your syntax is correct.

## NotFound

This client error response indicates that the server cannot find the requested resource.

## MethodNotAllowed

Indicates that the request method is known by the server but is not supported by the target resource.

## NotAcceptable

The target resource does not have a current representation that would be acceptable to the user
agent, according to the proactive negotiation header fields received in the request, and the
server is unwilling to supply a default representation.

## AlreadyExists

The request could not be completed due to a conflict with the current state of the target
resource, e.g. the resource it tries to create already exists.

## Unknown

This means that the server encountered an unexpected condition that prevented it from
fulfilling the request.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.7.0
- Package version: 2.5.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://airflow.apache.org](https://airflow.apache.org)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import airflow "github.com/apache/airflow-client-go/airflow"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), airflow.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), airflow.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), airflow.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), airflow.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to */api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ConfigApi* | [**GetConfig**](docs/ConfigApi.md#getconfig) | **Get** /config | Get current configuration
*ConfigApi* | [**GetValue**](docs/ConfigApi.md#getvalue) | **Get** /config/section/{section}/option/{option} | Get a option from configuration
*ConnectionApi* | [**DeleteConnection**](docs/ConnectionApi.md#deleteconnection) | **Delete** /connections/{connection_id} | Delete a connection
*ConnectionApi* | [**GetConnection**](docs/ConnectionApi.md#getconnection) | **Get** /connections/{connection_id} | Get a connection
*ConnectionApi* | [**GetConnections**](docs/ConnectionApi.md#getconnections) | **Get** /connections | List connections
*ConnectionApi* | [**PatchConnection**](docs/ConnectionApi.md#patchconnection) | **Patch** /connections/{connection_id} | Update a connection
*ConnectionApi* | [**PostConnection**](docs/ConnectionApi.md#postconnection) | **Post** /connections | Create a connection
*ConnectionApi* | [**TestConnection**](docs/ConnectionApi.md#testconnection) | **Post** /connections/test | Test a connection
*DAGApi* | [**DeleteDag**](docs/DAGApi.md#deletedag) | **Delete** /dags/{dag_id} | Delete a DAG
*DAGApi* | [**GetDag**](docs/DAGApi.md#getdag) | **Get** /dags/{dag_id} | Get basic information about a DAG
*DAGApi* | [**GetDagDetails**](docs/DAGApi.md#getdagdetails) | **Get** /dags/{dag_id}/details | Get a simplified representation of DAG
*DAGApi* | [**GetDagSource**](docs/DAGApi.md#getdagsource) | **Get** /dagSources/{file_token} | Get a source code
*DAGApi* | [**GetDags**](docs/DAGApi.md#getdags) | **Get** /dags | List DAGs
*DAGApi* | [**GetTask**](docs/DAGApi.md#gettask) | **Get** /dags/{dag_id}/tasks/{task_id} | Get simplified representation of a task
*DAGApi* | [**GetTasks**](docs/DAGApi.md#gettasks) | **Get** /dags/{dag_id}/tasks | Get tasks for DAG
*DAGApi* | [**PatchDag**](docs/DAGApi.md#patchdag) | **Patch** /dags/{dag_id} | Update a DAG
*DAGApi* | [**PatchDags**](docs/DAGApi.md#patchdags) | **Patch** /dags | Update DAGs
*DAGApi* | [**PostClearTaskInstances**](docs/DAGApi.md#postcleartaskinstances) | **Post** /dags/{dag_id}/clearTaskInstances | Clear a set of task instances
*DAGApi* | [**PostSetTaskInstancesState**](docs/DAGApi.md#postsettaskinstancesstate) | **Post** /dags/{dag_id}/updateTaskInstancesState | Set a state of task instances
*DAGRunApi* | [**ClearDagRun**](docs/DAGRunApi.md#cleardagrun) | **Post** /dags/{dag_id}/dagRuns/{dag_run_id}/clear | Clear a DAG run
*DAGRunApi* | [**DeleteDagRun**](docs/DAGRunApi.md#deletedagrun) | **Delete** /dags/{dag_id}/dagRuns/{dag_run_id} | Delete a DAG run
*DAGRunApi* | [**GetDagRun**](docs/DAGRunApi.md#getdagrun) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id} | Get a DAG run
*DAGRunApi* | [**GetDagRuns**](docs/DAGRunApi.md#getdagruns) | **Get** /dags/{dag_id}/dagRuns | List DAG runs
*DAGRunApi* | [**GetDagRunsBatch**](docs/DAGRunApi.md#getdagrunsbatch) | **Post** /dags/~/dagRuns/list | List DAG runs (batch)
*DAGRunApi* | [**GetUpstreamDatasetEvents**](docs/DAGRunApi.md#getupstreamdatasetevents) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/upstreamDatasetEvents | Get dataset events for a DAG run
*DAGRunApi* | [**PostDagRun**](docs/DAGRunApi.md#postdagrun) | **Post** /dags/{dag_id}/dagRuns | Trigger a new DAG run.
*DAGRunApi* | [**SetDagRunNote**](docs/DAGRunApi.md#setdagrunnote) | **Patch** /dags/{dag_id}/dagRuns/{dag_run_id}/setNote | Update the DagRun note.
*DAGRunApi* | [**UpdateDagRunState**](docs/DAGRunApi.md#updatedagrunstate) | **Patch** /dags/{dag_id}/dagRuns/{dag_run_id} | Modify a DAG run
*DagWarningApi* | [**GetDagWarnings**](docs/DagWarningApi.md#getdagwarnings) | **Get** /dagWarnings | List dag warnings
*DatasetApi* | [**GetDataset**](docs/DatasetApi.md#getdataset) | **Get** /datasets/{uri} | Get a dataset
*DatasetApi* | [**GetDatasetEvents**](docs/DatasetApi.md#getdatasetevents) | **Get** /datasets/events | Get dataset events
*DatasetApi* | [**GetDatasets**](docs/DatasetApi.md#getdatasets) | **Get** /datasets | List datasets
*DatasetApi* | [**GetUpstreamDatasetEvents**](docs/DatasetApi.md#getupstreamdatasetevents) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/upstreamDatasetEvents | Get dataset events for a DAG run
*EventLogApi* | [**GetEventLog**](docs/EventLogApi.md#geteventlog) | **Get** /eventLogs/{event_log_id} | Get a log entry
*EventLogApi* | [**GetEventLogs**](docs/EventLogApi.md#geteventlogs) | **Get** /eventLogs | List log entries
*ImportErrorApi* | [**GetImportError**](docs/ImportErrorApi.md#getimporterror) | **Get** /importErrors/{import_error_id} | Get an import error
*ImportErrorApi* | [**GetImportErrors**](docs/ImportErrorApi.md#getimporterrors) | **Get** /importErrors | List import errors
*MonitoringApi* | [**GetHealth**](docs/MonitoringApi.md#gethealth) | **Get** /health | Get instance status
*MonitoringApi* | [**GetVersion**](docs/MonitoringApi.md#getversion) | **Get** /version | Get version information
*PermissionApi* | [**GetPermissions**](docs/PermissionApi.md#getpermissions) | **Get** /permissions | List permissions
*PluginApi* | [**GetPlugins**](docs/PluginApi.md#getplugins) | **Get** /plugins | Get a list of loaded plugins
*PoolApi* | [**DeletePool**](docs/PoolApi.md#deletepool) | **Delete** /pools/{pool_name} | Delete a pool
*PoolApi* | [**GetPool**](docs/PoolApi.md#getpool) | **Get** /pools/{pool_name} | Get a pool
*PoolApi* | [**GetPools**](docs/PoolApi.md#getpools) | **Get** /pools | List pools
*PoolApi* | [**PatchPool**](docs/PoolApi.md#patchpool) | **Patch** /pools/{pool_name} | Update a pool
*PoolApi* | [**PostPool**](docs/PoolApi.md#postpool) | **Post** /pools | Create a pool
*ProviderApi* | [**GetProviders**](docs/ProviderApi.md#getproviders) | **Get** /providers | List providers
*RoleApi* | [**DeleteRole**](docs/RoleApi.md#deleterole) | **Delete** /roles/{role_name} | Delete a role
*RoleApi* | [**GetRole**](docs/RoleApi.md#getrole) | **Get** /roles/{role_name} | Get a role
*RoleApi* | [**GetRoles**](docs/RoleApi.md#getroles) | **Get** /roles | List roles
*RoleApi* | [**PatchRole**](docs/RoleApi.md#patchrole) | **Patch** /roles/{role_name} | Update a role
*RoleApi* | [**PostRole**](docs/RoleApi.md#postrole) | **Post** /roles | Create a role
*TaskInstanceApi* | [**GetExtraLinks**](docs/TaskInstanceApi.md#getextralinks) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/links | List extra links
*TaskInstanceApi* | [**GetLog**](docs/TaskInstanceApi.md#getlog) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/logs/{task_try_number} | Get logs
*TaskInstanceApi* | [**GetMappedTaskInstance**](docs/TaskInstanceApi.md#getmappedtaskinstance) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/{map_index} | Get a mapped task instance
*TaskInstanceApi* | [**GetMappedTaskInstances**](docs/TaskInstanceApi.md#getmappedtaskinstances) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/listMapped | List mapped task instances
*TaskInstanceApi* | [**GetTaskInstance**](docs/TaskInstanceApi.md#gettaskinstance) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id} | Get a task instance
*TaskInstanceApi* | [**GetTaskInstances**](docs/TaskInstanceApi.md#gettaskinstances) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances | List task instances
*TaskInstanceApi* | [**GetTaskInstancesBatch**](docs/TaskInstanceApi.md#gettaskinstancesbatch) | **Post** /dags/~/dagRuns/~/taskInstances/list | List task instances (batch)
*TaskInstanceApi* | [**PatchMappedTaskInstance**](docs/TaskInstanceApi.md#patchmappedtaskinstance) | **Patch** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/{map_index} | Updates the state of a mapped task instance
*TaskInstanceApi* | [**PatchTaskInstance**](docs/TaskInstanceApi.md#patchtaskinstance) | **Patch** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id} | Updates the state of a task instance
*TaskInstanceApi* | [**SetMappedTaskInstanceNote**](docs/TaskInstanceApi.md#setmappedtaskinstancenote) | **Patch** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/{map_index}/setNote | Update the TaskInstance note.
*TaskInstanceApi* | [**SetTaskInstanceNote**](docs/TaskInstanceApi.md#settaskinstancenote) | **Patch** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/setNote | Update the TaskInstance note.
*UserApi* | [**DeleteUser**](docs/UserApi.md#deleteuser) | **Delete** /users/{username} | Delete a user
*UserApi* | [**GetUser**](docs/UserApi.md#getuser) | **Get** /users/{username} | Get a user
*UserApi* | [**GetUsers**](docs/UserApi.md#getusers) | **Get** /users | List users
*UserApi* | [**PatchUser**](docs/UserApi.md#patchuser) | **Patch** /users/{username} | Update a user
*UserApi* | [**PostUser**](docs/UserApi.md#postuser) | **Post** /users | Create a user
*VariableApi* | [**DeleteVariable**](docs/VariableApi.md#deletevariable) | **Delete** /variables/{variable_key} | Delete a variable
*VariableApi* | [**GetVariable**](docs/VariableApi.md#getvariable) | **Get** /variables/{variable_key} | Get a variable
*VariableApi* | [**GetVariables**](docs/VariableApi.md#getvariables) | **Get** /variables | List variables
*VariableApi* | [**PatchVariable**](docs/VariableApi.md#patchvariable) | **Patch** /variables/{variable_key} | Update a variable
*VariableApi* | [**PostVariables**](docs/VariableApi.md#postvariables) | **Post** /variables | Create a variable
*XComApi* | [**GetXcomEntries**](docs/XComApi.md#getxcomentries) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/xcomEntries | List XCom entries
*XComApi* | [**GetXcomEntry**](docs/XComApi.md#getxcomentry) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/xcomEntries/{xcom_key} | Get an XCom entry


## Documentation For Models

 - [Action](docs/Action.md)
 - [ActionCollection](docs/ActionCollection.md)
 - [ActionCollectionAllOf](docs/ActionCollectionAllOf.md)
 - [ActionResource](docs/ActionResource.md)
 - [BasicDAGRun](docs/BasicDAGRun.md)
 - [ClassReference](docs/ClassReference.md)
 - [ClearDagRun](docs/ClearDagRun.md)
 - [ClearTaskInstances](docs/ClearTaskInstances.md)
 - [CollectionInfo](docs/CollectionInfo.md)
 - [Config](docs/Config.md)
 - [ConfigOption](docs/ConfigOption.md)
 - [ConfigSection](docs/ConfigSection.md)
 - [Connection](docs/Connection.md)
 - [ConnectionAllOf](docs/ConnectionAllOf.md)
 - [ConnectionCollection](docs/ConnectionCollection.md)
 - [ConnectionCollectionAllOf](docs/ConnectionCollectionAllOf.md)
 - [ConnectionCollectionItem](docs/ConnectionCollectionItem.md)
 - [ConnectionTest](docs/ConnectionTest.md)
 - [CronExpression](docs/CronExpression.md)
 - [DAG](docs/DAG.md)
 - [DAGCollection](docs/DAGCollection.md)
 - [DAGCollectionAllOf](docs/DAGCollectionAllOf.md)
 - [DAGDetail](docs/DAGDetail.md)
 - [DAGDetailAllOf](docs/DAGDetailAllOf.md)
 - [DAGRun](docs/DAGRun.md)
 - [DAGRunCollection](docs/DAGRunCollection.md)
 - [DAGRunCollectionAllOf](docs/DAGRunCollectionAllOf.md)
 - [DagProcessorStatus](docs/DagProcessorStatus.md)
 - [DagScheduleDatasetReference](docs/DagScheduleDatasetReference.md)
 - [DagState](docs/DagState.md)
 - [DagWarning](docs/DagWarning.md)
 - [DagWarningCollection](docs/DagWarningCollection.md)
 - [DagWarningCollectionAllOf](docs/DagWarningCollectionAllOf.md)
 - [Dataset](docs/Dataset.md)
 - [DatasetCollection](docs/DatasetCollection.md)
 - [DatasetCollectionAllOf](docs/DatasetCollectionAllOf.md)
 - [DatasetEvent](docs/DatasetEvent.md)
 - [DatasetEventCollection](docs/DatasetEventCollection.md)
 - [DatasetEventCollectionAllOf](docs/DatasetEventCollectionAllOf.md)
 - [Error](docs/Error.md)
 - [EventLog](docs/EventLog.md)
 - [EventLogCollection](docs/EventLogCollection.md)
 - [EventLogCollectionAllOf](docs/EventLogCollectionAllOf.md)
 - [ExtraLink](docs/ExtraLink.md)
 - [ExtraLinkCollection](docs/ExtraLinkCollection.md)
 - [HealthInfo](docs/HealthInfo.md)
 - [HealthStatus](docs/HealthStatus.md)
 - [ImportError](docs/ImportError.md)
 - [ImportErrorCollection](docs/ImportErrorCollection.md)
 - [ImportErrorCollectionAllOf](docs/ImportErrorCollectionAllOf.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [Job](docs/Job.md)
 - [ListDagRunsForm](docs/ListDagRunsForm.md)
 - [ListTaskInstanceForm](docs/ListTaskInstanceForm.md)
 - [MetadatabaseStatus](docs/MetadatabaseStatus.md)
 - [PluginCollection](docs/PluginCollection.md)
 - [PluginCollectionAllOf](docs/PluginCollectionAllOf.md)
 - [PluginCollectionItem](docs/PluginCollectionItem.md)
 - [Pool](docs/Pool.md)
 - [PoolCollection](docs/PoolCollection.md)
 - [PoolCollectionAllOf](docs/PoolCollectionAllOf.md)
 - [Provider](docs/Provider.md)
 - [ProviderCollection](docs/ProviderCollection.md)
 - [RelativeDelta](docs/RelativeDelta.md)
 - [Resource](docs/Resource.md)
 - [Role](docs/Role.md)
 - [RoleCollection](docs/RoleCollection.md)
 - [RoleCollectionAllOf](docs/RoleCollectionAllOf.md)
 - [SLAMiss](docs/SLAMiss.md)
 - [ScheduleInterval](docs/ScheduleInterval.md)
 - [SchedulerStatus](docs/SchedulerStatus.md)
 - [SetDagRunNote](docs/SetDagRunNote.md)
 - [SetTaskInstanceNote](docs/SetTaskInstanceNote.md)
 - [Tag](docs/Tag.md)
 - [Task](docs/Task.md)
 - [TaskCollection](docs/TaskCollection.md)
 - [TaskExtraLinks](docs/TaskExtraLinks.md)
 - [TaskInstance](docs/TaskInstance.md)
 - [TaskInstanceCollection](docs/TaskInstanceCollection.md)
 - [TaskInstanceCollectionAllOf](docs/TaskInstanceCollectionAllOf.md)
 - [TaskInstanceReference](docs/TaskInstanceReference.md)
 - [TaskInstanceReferenceCollection](docs/TaskInstanceReferenceCollection.md)
 - [TaskOutletDatasetReference](docs/TaskOutletDatasetReference.md)
 - [TaskState](docs/TaskState.md)
 - [TimeDelta](docs/TimeDelta.md)
 - [Trigger](docs/Trigger.md)
 - [TriggerRule](docs/TriggerRule.md)
 - [TriggererStatus](docs/TriggererStatus.md)
 - [UpdateDagRunState](docs/UpdateDagRunState.md)
 - [UpdateTaskInstance](docs/UpdateTaskInstance.md)
 - [UpdateTaskInstancesState](docs/UpdateTaskInstancesState.md)
 - [User](docs/User.md)
 - [UserAllOf](docs/UserAllOf.md)
 - [UserCollection](docs/UserCollection.md)
 - [UserCollectionAllOf](docs/UserCollectionAllOf.md)
 - [UserCollectionItem](docs/UserCollectionItem.md)
 - [UserCollectionItemRoles](docs/UserCollectionItemRoles.md)
 - [Variable](docs/Variable.md)
 - [VariableAllOf](docs/VariableAllOf.md)
 - [VariableCollection](docs/VariableCollection.md)
 - [VariableCollectionAllOf](docs/VariableCollectionAllOf.md)
 - [VariableCollectionItem](docs/VariableCollectionItem.md)
 - [VersionInfo](docs/VersionInfo.md)
 - [WeightRule](docs/WeightRule.md)
 - [XCom](docs/XCom.md)
 - [XComAllOf](docs/XComAllOf.md)
 - [XComCollection](docs/XComCollection.md)
 - [XComCollectionAllOf](docs/XComCollectionAllOf.md)
 - [XComCollectionItem](docs/XComCollectionItem.md)


## Documentation For Authorization



### Basic

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


### Kerberos


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

dev@airflow.apache.org

