module github.com/signalfx/splunk-otel-go/instrumentation/github.com/graph-gophers/graphql-go/splunkgraphql

go 1.20

require (
	github.com/graph-gophers/graphql-go v1.5.0
	github.com/signalfx/splunk-otel-go/instrumentation/internal v1.10.0
	go.opentelemetry.io/otel v1.20.0
	go.opentelemetry.io/otel/trace v1.20.0
)

require (
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel/metric v1.20.0 // indirect
)

replace github.com/signalfx/splunk-otel-go/instrumentation/internal => ../../../../internal
