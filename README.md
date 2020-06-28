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
Airflow Go client
=================

Go Airflow OpenAPI client generated from [openapi spec](https://github.com/apache/airflow/tree/master/clients).


Usage
-----

```go
import (
	"github.com/apache/airflow-client-go/airflow"
)

func main() {
	cli := airflow.NewAPIClient(&airflow.Configuration{
		Scheme:   "http",
		Host:     "localhost:28080",
		BasePath: "/api/v1",
	})
	ctx := context.TODO()
	variable, _, err := client.VariableApi.GetVariable(ctx, "foo")
	if err != nil {
		fmt.Println(variable)
	} else {
		fmt.Println(err)
	}
}
```

See [README](./airflow/README.md#documentation-for-api-endpoints) for full client API documentation.


Release process
---------------

Go client is versioned using [semantic import
versioning](https://blog.golang.org/versioning-proposal).

To release a new version `1.x.y`, simply push a new tag to this repo named
`airflow/v1.x.y`.

Major version upgrade requires changing package import path based on semantic
import versioning, which needs to be done manually.
