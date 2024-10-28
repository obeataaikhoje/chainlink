package beholder_test

import (
	"fmt"
	"time"

	otelattr "go.opentelemetry.io/otel/attribute"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
)

const (
	packageName = "beholder"
)

func ExampleConfig() {
	config := beholder.Config{
		InsecureConnection:       true,
		CACertFile:               "",
		OtelExporterGRPCEndpoint: "localhost:4317",
		OtelExporterHTTPEndpoint: "localhost:4318",
		// Resource
		ResourceAttributes: []otelattr.KeyValue{
			otelattr.String("package_name", packageName),
			otelattr.String("sender", "beholderclient"),
		},
		// Message Emitter
		EmitterExportTimeout:  1 * time.Second,
		EmitterBatchProcessor: true,
		// OTel message log exporter retry config
		EmitterExporterRetryConfig: nil,
		// Trace
		TraceSampleRatio:  1,
		TraceBatchTimeout: 1 * time.Second,
		// OTel trace exporter retry config
		TraceExporterRetryConfig: nil,
		// Metric
		MetricReaderInterval: 1 * time.Second,
		// OTel metric exporter retry config
		MetricExporterRetryConfig: nil,
		// Log
		LogExportTimeout:  1 * time.Second,
		LogBatchProcessor: true,
	}
	fmt.Printf("%+v\n", config)
	config.EmitterExporterRetryConfig = &beholder.RetryConfig{
		InitialInterval: 5 * time.Second,
		MaxInterval:     30 * time.Second,
		MaxElapsedTime:  1 * time.Minute, // Set to zero to disable retry
	}
	fmt.Printf("%+v\n", *config.EmitterExporterRetryConfig)
	// Output:
	// {InsecureConnection:true CACertFile: OtelExporterGRPCEndpoint:localhost:4317 OtelExporterHTTPEndpoint:localhost:4318 ResourceAttributes:[{Key:package_name Value:{vtype:4 numeric:0 stringly:beholder slice:<nil>}} {Key:sender Value:{vtype:4 numeric:0 stringly:beholderclient slice:<nil>}}] EmitterExportTimeout:1s EmitterBatchProcessor:true EmitterExporterRetryConfig:<nil> TraceSampleRatio:1 TraceBatchTimeout:1s TraceSpanExporter:<nil> TraceExporterRetryConfig:<nil> MetricReaderInterval:1s MetricExporterRetryConfig:<nil> LogExportTimeout:1s LogBatchProcessor:true}
	// {InitialInterval:5s MaxInterval:30s MaxElapsedTime:1m0s}

}
