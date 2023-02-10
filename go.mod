module github.com/apache/airflow-client-go

go 1.13

replace github.com/apache/airflow-client-go/airflow => ./airflow

require (
	github.com/apache/airflow-client-go/airflow v0.0.0-20200725194829-7af9875e7d4c
	github.com/stretchr/testify v1.6.1
)
