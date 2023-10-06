module github.com/apache/airflow-client-go

go 1.21

replace github.com/apache/airflow-client-go/airflow => ./airflow

require (
	github.com/apache/airflow-client-go/airflow v0.0.0-20200725194829-7af9875e7d4c
	github.com/stretchr/testify v1.6.1
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/oauth2 v0.0.0-20210218202405-ba52d332ba99 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)
