package ocr3

import (
	"context"

	"github.com/mwitkow/grpc-proxy/proxy"
	"google.golang.org/grpc"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/loop/internal"
	"github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb"
	ocr3pb "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb/ocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

type ReportingPluginServiceClient struct {
	*internal.PluginClient
	*internal.ServiceClient

	reportingPluginService pb.ReportingPluginServiceClient
}

func NewReportingPluginServiceClient(broker internal.Broker, brokerCfg internal.BrokerConfig, conn *grpc.ClientConn) *ReportingPluginServiceClient {
	brokerCfg.Logger = logger.Named(brokerCfg.Logger, "ReportingPluginServiceClient")
	pc := internal.NewPluginClient(broker, brokerCfg, conn)
	return &ReportingPluginServiceClient{PluginClient: pc, reportingPluginService: pb.NewReportingPluginServiceClient(pc), ServiceClient: internal.NewServiceClient(pc.BrokerExt, pc)}
}

func (o *ReportingPluginServiceClient) NewReportingPluginFactory(
	ctx context.Context,
	config types.ReportingPluginServiceConfig,
	grpcProvider grpc.ClientConnInterface,
	pipelineRunner types.PipelineRunnerService,
	telemetry types.TelemetryService,
	errorLog types.ErrorLog,
	capRegistry types.CapabilitiesRegistry,
) (types.OCR3ReportingPluginFactory, error) {
	cc := o.NewClientConn("ReportingPluginServiceFactory", func(ctx context.Context) (id uint32, deps internal.Resources, err error) {
		providerID, providerRes, err := o.Serve("PluginProvider", proxy.NewProxy(grpcProvider))
		if err != nil {
			return 0, nil, err
		}
		deps.Add(providerRes)

		pipelineRunnerID, pipelineRunnerRes, err := o.ServeNew("PipelineRunner", func(s *grpc.Server) {
			pb.RegisterPipelineRunnerServiceServer(s, &internal.PipelineRunnerServiceServer{Impl: pipelineRunner})
		})
		if err != nil {
			return 0, nil, err
		}
		deps.Add(pipelineRunnerRes)

		telemetryID, telemetryRes, err := o.ServeNew("Telemetry", func(s *grpc.Server) {
			pb.RegisterTelemetryServer(s, internal.NewTelemetryServer(telemetry))
		})
		if err != nil {
			return 0, nil, err
		}
		deps.Add(telemetryRes)

		errorLogID, errorLogRes, err := o.ServeNew("ErrorLog", func(s *grpc.Server) {
			pb.RegisterErrorLogServer(s, &internal.ErrorLogServer{Impl: errorLog})
		})
		if err != nil {
			return 0, nil, err
		}
		deps.Add(errorLogRes)

		capRegistryID, capRegistryRes, err := o.ServeNew("CapRegistry", func(s *grpc.Server) {
			pb.RegisterCapabilitiesRegistryServer(s, internal.NewCapabilitiesRegistryServer(o.BrokerExt, capRegistry))
		})
		if err != nil {
			return 0, nil, err
		}
		deps.Add(capRegistryRes)

		reply, err := o.reportingPluginService.NewReportingPluginFactory(ctx, &pb.NewReportingPluginFactoryRequest{
			ReportingPluginServiceConfig: &pb.ReportingPluginServiceConfig{
				ProviderType:  config.ProviderType,
				Command:       config.Command,
				PluginName:    config.PluginName,
				TelemetryType: config.TelemetryType,
				PluginConfig:  config.PluginConfig,
			},
			ProviderID:       providerID,
			ErrorLogID:       errorLogID,
			PipelineRunnerID: pipelineRunnerID,
			TelemetryID:      telemetryID,
			CapRegistryID:    capRegistryID,
		})
		if err != nil {
			return 0, nil, err
		}
		return reply.ID, nil, nil
	})
	return newReportingPluginFactoryClient(o.PluginClient.BrokerExt, cc), nil
}

func (o *ReportingPluginServiceClient) NewValidationService(ctx context.Context) (types.ValidationService, error) {
	cc := o.NewClientConn("ReportingPluginServiceFactory", func(ctx context.Context) (id uint32, deps internal.Resources, err error) {
		reply, err := o.reportingPluginService.NewValidationService(ctx, &pb.ValidationServiceRequest{})
		if err != nil {
			return 0, nil, err
		}
		return reply.ID, nil, nil
	})
	return newValidationServiceClient(o.PluginClient.BrokerExt, cc), nil
}

var _ pb.ReportingPluginServiceServer = (*reportingPluginServiceServer)(nil)

type reportingPluginServiceServer struct {
	pb.UnimplementedReportingPluginServiceServer

	*internal.BrokerExt
	impl types.OCR3ReportingPluginClient
}

func (m reportingPluginServiceServer) NewReportingPluginFactory(ctx context.Context, request *pb.NewReportingPluginFactoryRequest) (*pb.NewReportingPluginFactoryReply, error) {
	errorLogConn, err := m.Dial(request.ErrorLogID)
	if err != nil {
		return nil, internal.ErrConnDial{Name: "ErrorLog", ID: request.ErrorLogID, Err: err}
	}
	errorLogRes := internal.Resource{Closer: errorLogConn, Name: "ErrorLog"}
	errorLog := internal.NewErrorLogClient(errorLogConn)

	providerConn, err := m.Dial(request.ProviderID)
	if err != nil {
		m.CloseAll(errorLogRes)
		return nil, internal.ErrConnDial{Name: "PluginProvider", ID: request.ProviderID, Err: err}
	}
	providerRes := internal.Resource{Closer: providerConn, Name: "PluginProvider"}

	pipelineRunnerConn, err := m.Dial(request.PipelineRunnerID)
	if err != nil {
		m.CloseAll(errorLogRes, providerRes)
		return nil, internal.ErrConnDial{Name: "PipelineRunner", ID: request.PipelineRunnerID, Err: err}
	}
	pipelineRunnerRes := internal.Resource{Closer: pipelineRunnerConn, Name: "PipelineRunner"}
	pipelineRunner := internal.NewPipelineRunnerClient(pipelineRunnerConn)

	telemetryConn, err := m.Dial(request.TelemetryID)
	if err != nil {
		m.CloseAll(errorLogRes, providerRes, pipelineRunnerRes)
		return nil, internal.ErrConnDial{Name: "Telemetry", ID: request.TelemetryID, Err: err}
	}
	telemetryRes := internal.Resource{Closer: telemetryConn, Name: "Telemetry"}
	telemetry := internal.NewTelemetryServiceClient(telemetryConn)

	capRegistryConn, err := m.Dial(request.CapRegistryID)
	if err != nil {
		m.CloseAll(errorLogRes, providerRes, pipelineRunnerRes, telemetryRes)
		return nil, internal.ErrConnDial{Name: "CapabilitiesRegistry", ID: request.CapRegistryID, Err: err}
	}
	capRegistryRes := internal.Resource{Closer: capRegistryConn, Name: "CapabilitiesRegistry"}
	capRegistry := internal.NewCapabilitiesRegistryClient(capRegistryConn, m.BrokerExt)

	config := types.ReportingPluginServiceConfig{
		ProviderType:  request.ReportingPluginServiceConfig.ProviderType,
		PluginConfig:  request.ReportingPluginServiceConfig.PluginConfig,
		PluginName:    request.ReportingPluginServiceConfig.PluginName,
		Command:       request.ReportingPluginServiceConfig.Command,
		TelemetryType: request.ReportingPluginServiceConfig.TelemetryType,
	}

	factory, err := m.impl.NewReportingPluginFactory(ctx, config, providerConn, pipelineRunner, telemetry, errorLog, capRegistry)
	if err != nil {
		m.CloseAll(providerRes, errorLogRes, pipelineRunnerRes, telemetryRes, capRegistryRes)
		return nil, err
	}

	id, _, err := m.ServeNew("ReportingPluginProvider", func(s *grpc.Server) {
		pb.RegisterServiceServer(s, &internal.ServiceServer{Srv: factory})
		ocr3pb.RegisterReportingPluginFactoryServer(s, newReportingPluginFactoryServer(factory, m.BrokerExt))
	}, providerRes, errorLogRes, pipelineRunnerRes, telemetryRes, capRegistryRes)
	if err != nil {
		return nil, err
	}

	return &pb.NewReportingPluginFactoryReply{ID: id}, nil
}

func RegisterReportingPluginServiceServer(server *grpc.Server, broker internal.Broker, brokerCfg internal.BrokerConfig, impl types.OCR3ReportingPluginClient) error {
	pb.RegisterReportingPluginServiceServer(server, newReportingPluginServiceServer(&internal.BrokerExt{Broker: broker, BrokerConfig: brokerCfg}, impl))
	return nil
}

func newReportingPluginServiceServer(b *internal.BrokerExt, gp types.OCR3ReportingPluginClient) *reportingPluginServiceServer {
	return &reportingPluginServiceServer{BrokerExt: b.WithName("OCR3ReportingPluginService"), impl: gp}
}
