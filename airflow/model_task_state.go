// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

/*
 * Airflow API (Stable)
 *
 * Apache Airflow management API.
 *
 * API version: 1.0.0
 * Contact: dev@airflow.apache.org
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package airflow
// TaskState the model 'TaskState'
type TaskState string

// List of TaskState
const (
	TASKSTATE_SUCCESS TaskState = "success"
	TASKSTATE_RUNNING TaskState = "running"
	TASKSTATE_FAILED TaskState = "failed"
	TASKSTATE_UPSTREAM_FAILED TaskState = "upstream_failed"
	TASKSTATE_SKIPPED TaskState = "skipped"
	TASKSTATE_UP_FOR_RETRY TaskState = "up_for_retry"
	TASKSTATE_UP_FOR_RESCHEDULE TaskState = "up_for_reschedule"
	TASKSTATE_QUEUED TaskState = "queued"
	TASKSTATE_NONE TaskState = "none"
	TASKSTATE_SCHEDULED TaskState = "scheduled"
)
