module github.com/Financial-Times/cm-go-service

go 1.16

require (
	github.com/Financial-Times/api-endpoint v1.0.0
	github.com/Financial-Times/go-fthealth v0.0.0-20200609161010-4c53fbef65fa
	github.com/Financial-Times/go-logger/v2 v2.0.1
	github.com/Financial-Times/http-handlers-go/v2 v2.3.0
	github.com/Financial-Times/service-status-go v0.0.0-20200609183459-3c8b4c6d72a5
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/go-version v1.2.1 // indirect
	github.com/jawher/mow.cli v1.1.0
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.22.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.22.0
	go.opentelemetry.io/otel v1.0.0-RC2
	go.opentelemetry.io/otel/exporters/jaeger v1.0.0-RC2
	go.opentelemetry.io/otel/sdk v1.0.0-RC2
	go.opentelemetry.io/otel/trace v1.0.0-RC2
)
