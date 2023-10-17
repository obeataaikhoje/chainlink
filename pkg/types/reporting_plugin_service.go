package types

import (
	"context"

	"google.golang.org/grpc"
)

type ReportingPluginServiceConfig struct {
	ProviderType string
	Command      string
	PluginName   string
	PluginConfig []byte
}

// ReportingPluginClient is the client interface to a plugin running
// as a generic job (job type = GenericPlugin) inside the core node.
type ReportingPluginClient interface {
	NewReportingPluginFactory(ctx context.Context, config ReportingPluginServiceConfig, grpcProvider grpc.ClientConnInterface, errorLog ErrorLog) (ReportingPluginFactory, error)
}

// ReportingPluginServer is the server interface to a plugin running
// as a generic job (job type = GenericPlugin) inside the core node,
// with the passthrough provider connection converted to the provider
// expected by the plugin.
type ReportingPluginServer[T PluginProvider] interface {
	NewReportingPluginFactory(ctx context.Context, config ReportingPluginServiceConfig, provider T, errorLog ErrorLog) (ReportingPluginFactory, error)
}
