module {{cookiecutter.project_name}}

go 1.20

require (
	github.com/Ryanair/gofrlib v1.0.25
	github.com/Ryanair/gofrlib-test v1.0.5
	github.com/aws/aws-sdk-go-v2 v1.21.0
	github.com/aws/aws-sdk-go-v2/config v1.18.37
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.8.4
	go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws v0.42.0
)
