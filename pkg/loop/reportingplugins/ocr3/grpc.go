package ocr3

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	"github.com/smartcontractkit/chainlink-common/pkg/loop"
	"github.com/smartcontractkit/chainlink-common/pkg/loop/internal/core/services/reportingplugin/ocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/loop/internal/core/services/telemetry"
	"github.com/smartcontractkit/chainlink-common/pkg/loop/internal/net"
	"github.com/smartcontractkit/chainlink-common/pkg/loop/reportingplugins"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

const PluginServiceName = "ocr3-plugin-service"

type ProviderServer[T types.PluginProvider] interface {
	types.OCR3ReportingPluginServer[T]
	ConnToProvider(conn grpc.ClientConnInterface, broker net.Broker, brokerConfig loop.BrokerConfig) T
}

type GRPCService[T types.PluginProvider] struct {
	plugin.NetRPCUnsupportedPlugin

	loop.BrokerConfig

	PluginServer ProviderServer[T]

	pluginClient *ocr3.ReportingPluginServiceClient
}

type serverAdapter struct {
	NewReportingPluginFactoryFn func(
		context.Context,
		types.ReportingPluginServiceConfig,
		grpc.ClientConnInterface,
		types.PipelineRunnerService,
		types.TelemetryService,
		types.ErrorLog,
		types.CapabilitiesRegistry,
		types.KeyValueStore,
	) (types.OCR3ReportingPluginFactory, error)

	ValidateConfigService
}

type ValidateConfigService interface {
	NewValidationService(ctx context.Context) (types.ValidationService, error)
}

func (s serverAdapter) NewValidationService(ctx context.Context) (types.ValidationService, error) {
	return s.ValidateConfigService.NewValidationService(ctx)
}

func (s serverAdapter) NewReportingPluginFactory(
	ctx context.Context,
	config types.ReportingPluginServiceConfig,
	conn grpc.ClientConnInterface,
	pr types.PipelineRunnerService,
	ts types.TelemetryService,
	errorLog types.ErrorLog,
	capRegistry types.CapabilitiesRegistry,
	kv types.KeyValueStore,
) (types.OCR3ReportingPluginFactory, error) {
	return s.NewReportingPluginFactoryFn(ctx, config, conn, pr, ts, errorLog, capRegistry, kv)
}

func (g *GRPCService[T]) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	newReportingPluginFactoryFn := func(
		ctx context.Context,
		cfg types.ReportingPluginServiceConfig,
		conn grpc.ClientConnInterface,
		pr types.PipelineRunnerService,
		ts types.TelemetryService,
		el types.ErrorLog,
		capRegistry types.CapabilitiesRegistry,
		kv types.KeyValueStore,
	) (types.OCR3ReportingPluginFactory, error) {
		provider := g.PluginServer.ConnToProvider(conn, broker, g.BrokerConfig)
		tc := telemetry.NewTelemetryClient(ts)
		return g.PluginServer.NewReportingPluginFactory(ctx, cfg, provider, pr, tc, el, capRegistry, kv)
	}

	return ocr3.RegisterReportingPluginServiceServer(server, broker, g.BrokerConfig, serverAdapter{
		NewReportingPluginFactoryFn: newReportingPluginFactoryFn,
		ValidateConfigService:       g.PluginServer,
	})
}

func (g *GRPCService[T]) GRPCClient(_ context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	if g.pluginClient == nil {
		g.pluginClient = ocr3.NewReportingPluginServiceClient(broker, g.BrokerConfig, conn)
	} else {
		g.pluginClient.Refresh(broker, conn)
	}

	return types.OCR3ReportingPluginClient(g.pluginClient), nil
}

func (g *GRPCService[T]) ClientConfig() *plugin.ClientConfig {
	c := &plugin.ClientConfig{
		HandshakeConfig: reportingplugins.ReportingPluginHandshakeConfig(),
		Plugins:         map[string]plugin.Plugin{reportingplugins.PluginServiceName: g},
	}
	return loop.ManagedGRPCClientConfig(c, g.BrokerConfig)
}
