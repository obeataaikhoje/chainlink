package internal

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/mwitkow/grpc-proxy/proxy"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	mercury_common_internal "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/mercury/common"
	mercury_v1_internal "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/mercury/v1"
	mercury_v2_internal "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/mercury/v2"
	mercury_v3_internal "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/mercury/v3"
	"github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb"
	mercury_pb "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb/mercury"
	mercury_v1_pb "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb/mercury/v1"
	mercury_v2_pb "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb/mercury/v2"
	mercury_v3_pb "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb/mercury/v3"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/mercury"
	mercury_v1 "github.com/smartcontractkit/chainlink-common/pkg/types/mercury/v1"
	mercury_v2 "github.com/smartcontractkit/chainlink-common/pkg/types/mercury/v2"
	mercury_v3 "github.com/smartcontractkit/chainlink-common/pkg/types/mercury/v3"
)

type MercuryAdapterClient struct {
	*pluginClient
	*serviceClient

	mercury mercury_pb.MercuryAdapterClient
}

func NewMercuryAdapterClient(broker Broker, brokerCfg BrokerConfig, conn *grpc.ClientConn) *MercuryAdapterClient {
	brokerCfg.Logger = logger.Named(brokerCfg.Logger, "MercuryAdapterClient")
	pc := newPluginClient(broker, brokerCfg, conn)
	return &MercuryAdapterClient{
		pluginClient:  pc,
		serviceClient: newServiceClient(pc.brokerExt, pc),
		mercury:       mercury_pb.NewMercuryAdapterClient(pc),
	}
}

func (c *MercuryAdapterClient) NewMercuryV1Factory(ctx context.Context,
	provider types.MercuryProvider, dataSource mercury_v1.DataSource) (types.ReportingPluginFactory, error) {
	// every time a new client is created, we have to ensure that all the external dependencies are satisfied.
	// at this layer of the stack, all of those dependencies are other gRPC services.
	// some of those services are hosted in the same process as the client itself and others may be remote.
	newMercuryClientFn := func(ctx context.Context) (id uint32, deps resources, err error) {
		// the local resources for mercury are the DataSource
		dataSourceID, dsRes, err := c.serveNew("DataSource", func(s *grpc.Server) {
			mercury_v1_pb.RegisterDataSourceServer(s, mercury_v1_internal.NewDataSourceServer(dataSource))
		})
		if err != nil {
			return 0, nil, err
		}
		deps.Add(dsRes)

		// the proxyable resources for mercury are the Provider,  which may or may not be local to the client process. (legacy vs loopp)
		var (
			providerID  uint32
			providerRes resource
		)
		if grpcProvider, ok := provider.(GRPCClientConn); ok {
			providerID, providerRes, err = c.serve("MercuryProvider", proxy.NewProxy(grpcProvider.ClientConn()))
		} else {
			providerID, providerRes, err = c.serveNew("MercuryProvider", func(s *grpc.Server) {
				registerCommonServices(s, provider)

				mercury_pb.RegisterReportCodecV1Server(s, mercury_common_internal.NewReportCodecV1Server(s, provider.ReportCodecV1()))
				mercury_pb.RegisterReportCodecV2Server(s, mercury_pb.UnimplementedReportCodecV2Server{})
				mercury_pb.RegisterReportCodecV3Server(s, mercury_pb.UnimplementedReportCodecV3Server{})
			})
		}
		if err != nil {
			return 0, nil, err
		}
		deps.Add(providerRes)

		reply, err := c.mercury.NewMercuryV1Factory(ctx, &mercury_pb.NewMercuryV1FactoryRequest{
			MercuryProviderID: providerID,
			DataSourceV1ID:    dataSourceID,
		})
		if err != nil {
			return 0, nil, err
		}
		return reply.MercuryV1FactoryID, deps, nil
	}

	cc := c.newClientConn("MercuryV3Factory", newMercuryClientFn)
	return newReportingPluginFactoryClient(c.pluginClient.brokerExt, cc), nil
}

func (c *MercuryAdapterClient) NewMercuryV2Factory(ctx context.Context,
	provider types.MercuryProvider, dataSource mercury_v2.DataSource) (types.ReportingPluginFactory, error) {
	// every time a new client is created, we have to ensure that all the external dependencies are satisfied.
	// at this layer of the stack, all of those dependencies are other gRPC services.
	// some of those services are hosted in the same process as the client itself and others may be remote.
	newMercuryClientFn := func(ctx context.Context) (id uint32, deps resources, err error) {
		// the local resources for mercury are the DataSource
		dataSourceID, dsRes, err := c.serveNew("DataSource", func(s *grpc.Server) {
			mercury_v2_pb.RegisterDataSourceServer(s, mercury_v2_internal.NewDataSourceServer(dataSource))
		})
		if err != nil {
			return 0, nil, err
		}
		deps.Add(dsRes)

		// the proxyable resources for mercury are the Provider,  which may or may not be local to the client process. (legacy vs loopp)
		var (
			providerID  uint32
			providerRes resource
		)
		if grpcProvider, ok := provider.(GRPCClientConn); ok {
			providerID, providerRes, err = c.serve("MercuryProvider", proxy.NewProxy(grpcProvider.ClientConn()))
		} else {
			providerID, providerRes, err = c.serveNew("MercuryProvider", func(s *grpc.Server) {
				registerCommonServices(s, provider)

				mercury_pb.RegisterReportCodecV2Server(s, mercury_common_internal.NewReportCodecV2Server(s, provider.ReportCodecV2()))

				mercury_pb.RegisterReportCodecV1Server(s, mercury_pb.UnimplementedReportCodecV1Server{})
				mercury_pb.RegisterReportCodecV3Server(s, mercury_pb.UnimplementedReportCodecV3Server{})
			})
		}
		if err != nil {
			return 0, nil, err
		}
		deps.Add(providerRes)

		reply, err := c.mercury.NewMercuryV2Factory(ctx, &mercury_pb.NewMercuryV2FactoryRequest{
			MercuryProviderID: providerID,
			DataSourceV2ID:    dataSourceID,
		})
		if err != nil {
			return 0, nil, err
		}
		return reply.MercuryV2FactoryID, deps, nil
	}

	cc := c.newClientConn("MercuryV2Factory", newMercuryClientFn)
	return newReportingPluginFactoryClient(c.pluginClient.brokerExt, cc), nil
}

func registerCommonServices(s *grpc.Server, provider types.MercuryProvider) {
	pb.RegisterServiceServer(s, &serviceServer{srv: provider})
	pb.RegisterOffchainConfigDigesterServer(s, &offchainConfigDigesterServer{impl: provider.OffchainConfigDigester()})
	pb.RegisterContractConfigTrackerServer(s, &contractConfigTrackerServer{impl: provider.ContractConfigTracker()})
	pb.RegisterContractTransmitterServer(s, &contractTransmitterServer{impl: provider.ContractTransmitter()})
	mercury_pb.RegisterOnchainConfigCodecServer(s, mercury_common_internal.NewOnchainConfigCodecServer(provider.OnchainConfigCodec()))
	mercury_pb.RegisterServerFetcherServer(s, mercury_common_internal.NewServerFetcherServer(provider.MercuryServerFetcher()))
	mercury_pb.RegisterMercuryChainReaderServer(s, mercury_common_internal.NewChainReaderServer(provider.MercuryChainReader()))
}

func (c *MercuryAdapterClient) NewMercuryV3Factory(ctx context.Context,
	provider types.MercuryProvider, dataSource mercury_v3.DataSource,
) (types.ReportingPluginFactory, error) {
	// every time a new client is created, we have to ensure that all the external dependencies are satisfied.
	// at this layer of the stack, all of those dependencies are other gRPC services.
	// some of those services are hosted in the same process as the client itself and others may be remote.
	newMercuryClientFn := func(ctx context.Context) (id uint32, deps resources, err error) {
		// the local resources for mercury are the DataSource
		dataSourceID, dsRes, err := c.serveNew("DataSource", func(s *grpc.Server) {
			mercury_v3_pb.RegisterDataSourceServer(s, mercury_v3_internal.NewDataSourceServer(dataSource))
		})
		if err != nil {
			return 0, nil, err
		}
		deps.Add(dsRes)

		// the proxyable resources for mercury are the Provider,  which may or may not be local to the client process. (legacy vs loopp)
		var (
			providerID  uint32
			providerRes resource
		)
		if grpcProvider, ok := provider.(GRPCClientConn); ok {
			providerID, providerRes, err = c.serve("MercuryProvider", proxy.NewProxy(grpcProvider.ClientConn()))
		} else {
			providerID, providerRes, err = c.serveNew("MercuryProvider", func(s *grpc.Server) {
				registerCommonServices(s, provider)

				mercury_pb.RegisterReportCodecV3Server(s, mercury_common_internal.NewReportCodecV3Server(s, provider.ReportCodecV3()))

				mercury_pb.RegisterReportCodecV1Server(s, mercury_pb.UnimplementedReportCodecV1Server{})
				mercury_pb.RegisterReportCodecV2Server(s, mercury_pb.UnimplementedReportCodecV2Server{})
			})
		}
		if err != nil {
			return 0, nil, err
		}
		deps.Add(providerRes)

		reply, err := c.mercury.NewMercuryV3Factory(ctx, &mercury_pb.NewMercuryV3FactoryRequest{
			MercuryProviderID: providerID,
			DataSourceV3ID:    dataSourceID,
		})
		if err != nil {
			return 0, nil, err
		}
		return reply.MercuryV3FactoryID, deps, nil
	}

	cc := c.newClientConn("MercuryV3Factory", newMercuryClientFn)
	return newReportingPluginFactoryClient(c.pluginClient.brokerExt, cc), nil
}

var _ mercury_pb.MercuryAdapterServer = (*mercuryAdapterServer)(nil)

type mercuryAdapterServer struct {
	mercury_pb.UnimplementedMercuryAdapterServer

	*brokerExt
	impl types.PluginMercury
}

func RegisterMercuryAdapterServer(s *grpc.Server, broker Broker, brokerCfg BrokerConfig, impl types.PluginMercury) error {
	mercury_pb.RegisterMercuryAdapterServer(s, newMercuryAdapterServer(&brokerExt{broker, brokerCfg}, impl))
	return nil
}

func newMercuryAdapterServer(b *brokerExt, impl types.PluginMercury) *mercuryAdapterServer {
	return &mercuryAdapterServer{brokerExt: b.withName("MercuryAdapter"), impl: impl}
}

func (ms *mercuryAdapterServer) NewMercuryV1Factory(ctx context.Context, req *mercury_pb.NewMercuryV1FactoryRequest) (*mercury_pb.NewMercuryV1FactoryReply, error) {
	// declared so we can clean up open resources
	var err error
	var deps resources
	defer func() {
		if err != nil {
			ms.closeAll(deps...)
		}
	}()

	dsConn, err := ms.dial(req.DataSourceV1ID)
	if err != nil {
		return nil, ErrConnDial{Name: "DataSourceV1", ID: req.DataSourceV1ID, Err: err}
	}
	dsRes := resource{Closer: dsConn, name: "DataSourceV1"}
	deps.Add(dsRes)
	ds := mercury_v1_internal.NewDataSourceClient(dsConn)

	providerConn, err := ms.dial(req.MercuryProviderID)
	if err != nil {
		return nil, ErrConnDial{Name: "MercuryProvider", ID: req.MercuryProviderID, Err: err}
	}
	providerRes := resource{Closer: providerConn, name: "MercuryProvider"}
	deps.Add(providerRes)
	provider := newMercuryProviderClient(ms.brokerExt, providerConn)
	factory, err := ms.impl.NewMercuryV1Factory(ctx, provider, ds)
	if err != nil {
		return nil, fmt.Errorf("failed to create MercuryV1Factory: %w", err)
	}

	id, _, err := ms.serveNew("MercuryV1Factory", func(s *grpc.Server) {
		pb.RegisterServiceServer(s, &serviceServer{srv: factory})
		pb.RegisterReportingPluginFactoryServer(s, newReportingPluginFactoryServer(factory, ms.brokerExt))
	}, deps...)
	if err != nil {
		return nil, fmt.Errorf("failed to serverNew: %w", err)
	}

	return &mercury_pb.NewMercuryV1FactoryReply{MercuryV1FactoryID: id}, nil
}

func (ms *mercuryAdapterServer) NewMercuryV2Factory(ctx context.Context, req *mercury_pb.NewMercuryV2FactoryRequest) (*mercury_pb.NewMercuryV2FactoryReply, error) {
	// declared so we can clean up open resources
	var err error
	var deps resources
	defer func() {
		if err != nil {
			ms.closeAll(deps...)
		}
	}()

	dsConn, err := ms.dial(req.DataSourceV2ID)
	if err != nil {
		return nil, ErrConnDial{Name: "DataSourceV2", ID: req.DataSourceV2ID, Err: err}
	}
	dsRes := resource{Closer: dsConn, name: "DataSourceV2"}
	deps.Add(dsRes)
	ds := mercury_v2_internal.NewDataSourceClient(dsConn)

	providerConn, err := ms.dial(req.MercuryProviderID)
	if err != nil {
		return nil, ErrConnDial{Name: "MercuryProvider", ID: req.MercuryProviderID, Err: err}
	}
	providerRes := resource{Closer: providerConn, name: "MercuryProvider"}
	deps.Add(providerRes)
	provider := newMercuryProviderClient(ms.brokerExt, providerConn)
	factory, err := ms.impl.NewMercuryV2Factory(ctx, provider, ds)
	if err != nil {
		return nil, fmt.Errorf("failed to create MercuryV2Factory: %w", err)
	}

	id, _, err := ms.serveNew("MercuryV2Factory", func(s *grpc.Server) {
		pb.RegisterServiceServer(s, &serviceServer{srv: factory})
		pb.RegisterReportingPluginFactoryServer(s, newReportingPluginFactoryServer(factory, ms.brokerExt))
	}, deps...)
	if err != nil {
		return nil, fmt.Errorf("failed to serverNew: %w", err)
	}

	return &mercury_pb.NewMercuryV2FactoryReply{MercuryV2FactoryID: id}, nil
}

func (ms *mercuryAdapterServer) NewMercuryV3Factory(ctx context.Context, req *mercury_pb.NewMercuryV3FactoryRequest) (*mercury_pb.NewMercuryV3FactoryReply, error) {
	// declared so we can clean up open resources
	var err error
	var deps resources
	defer func() {
		if err != nil {
			ms.closeAll(deps...)
		}
	}()

	dsConn, err := ms.dial(req.DataSourceV3ID)
	if err != nil {
		return nil, ErrConnDial{Name: "DataSourceV3", ID: req.DataSourceV3ID, Err: err}
	}
	dsRes := resource{Closer: dsConn, name: "DataSourceV3"}
	deps.Add(dsRes)
	ds := mercury_v3_internal.NewDataSourceClient(dsConn)

	providerConn, err := ms.dial(req.MercuryProviderID)
	if err != nil {
		return nil, ErrConnDial{Name: "MercuryProvider", ID: req.MercuryProviderID, Err: err}
	}
	providerRes := resource{Closer: providerConn, name: "MercuryProvider"}
	deps.Add(providerRes)
	provider := newMercuryProviderClient(ms.brokerExt, providerConn)
	factory, err := ms.impl.NewMercuryV3Factory(ctx, provider, ds)
	if err != nil {
		return nil, fmt.Errorf("failed to create MercuryV3Factory: %w", err)
	}

	id, _, err := ms.serveNew("MercuryV3Factory", func(s *grpc.Server) {
		pb.RegisterServiceServer(s, &serviceServer{srv: factory})
		pb.RegisterReportingPluginFactoryServer(s, newReportingPluginFactoryServer(factory, ms.brokerExt))
	}, deps...)
	if err != nil {
		return nil, fmt.Errorf("failed to serverNew: %w", err)
	}

	return &mercury_pb.NewMercuryV3FactoryReply{MercuryV3FactoryID: id}, nil
}

var (
	_ types.MercuryProvider = (*mercuryProviderClient)(nil)
	_ GRPCClientConn        = (*mercuryProviderClient)(nil)
)

type mercuryProviderClient struct {
	*pluginProviderClient
	reportCodecV3      mercury_v3.ReportCodec
	reportCodecV2      mercury_v2.ReportCodec
	reportCodecV1      mercury_v1.ReportCodec
	onchainConfigCodec mercury.OnchainConfigCodec
	serverFetcher      mercury.ServerFetcher
	chainReader        types.ChainReader
	mercuryChainReader mercury.ChainReader
}

func (m *mercuryProviderClient) ClientConn() grpc.ClientConnInterface { return m.cc }

func newMercuryProviderClient(b *brokerExt, cc grpc.ClientConnInterface) *mercuryProviderClient {
	m := &mercuryProviderClient{pluginProviderClient: newPluginProviderClient(b.withName("MercuryProviderClient"), cc)}

	m.reportCodecV1 = mercury_common_internal.NewReportCodecV1Client(mercury_v1_internal.NewReportCodecClient(m.cc))
	m.reportCodecV2 = mercury_common_internal.NewReportCodecV2Client(mercury_v2_internal.NewReportCodecClient(m.cc))
	m.reportCodecV3 = mercury_common_internal.NewReportCodecV3Client(mercury_v3_internal.NewReportCodecClient(m.cc))

	m.onchainConfigCodec = mercury_common_internal.NewOnchainConfigCodecClient(m.cc)
	m.serverFetcher = mercury_common_internal.NewServerFetcherClient(m.cc)
	m.mercuryChainReader = mercury_common_internal.NewChainReaderClient(m.cc)

	m.chainReader = &chainReaderClient{b, pb.NewChainReaderClient(m.cc)}
	return m
}

func (m *mercuryProviderClient) ReportCodecV3() mercury_v3.ReportCodec {
	return m.reportCodecV3
}

func (m *mercuryProviderClient) ReportCodecV2() mercury_v2.ReportCodec {
	return m.reportCodecV2
}

func (m *mercuryProviderClient) ReportCodecV1() mercury_v1.ReportCodec {
	return m.reportCodecV1
}

func (m *mercuryProviderClient) OnchainConfigCodec() mercury.OnchainConfigCodec {
	return m.onchainConfigCodec
}

func (m *mercuryProviderClient) ChainReader() types.ChainReader {
	return m.chainReader
}

func (m *mercuryProviderClient) MercuryChainReader() mercury.ChainReader {
	return m.mercuryChainReader
}

func (m *mercuryProviderClient) MercuryServerFetcher() mercury.ServerFetcher {
	return m.serverFetcher
}
