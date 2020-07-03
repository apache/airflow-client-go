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

# Dag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DagId** | **string** |  | [optional] [readonly] 
**RootDagId** | Pointer to **string** |  | [optional] [readonly] 
**IsPaused** | Pointer to **bool** |  | [optional] 
**IsSubdag** | **bool** |  | [optional] [readonly] 
**Fileloc** | **string** |  | [optional] [readonly] 
**FileToken** | **string** | The key containing the encrypted path to the file. Encryption and decryption take place only on the server. This prevents the client from reading an non-DAG file. This also ensures API extensibility, because the format of encrypted data may change.  | [optional] [readonly] 
**Owners** | **[]string** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] [readonly] 
**ScheduleInterval** | [**ScheduleInterval**](ScheduleInterval.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


