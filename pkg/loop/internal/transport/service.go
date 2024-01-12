package transport

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
)

var ErrPluginUnavailable = errors.New("plugin unavailable")

var _ services.Service = (*ServiceClient)(nil)

type ServiceClient struct {
	LoggerStopper
	//b    *brokerExt
	cc   grpc.ClientConnInterface
	grpc pb.ServiceClient
}

func NewServiceClient(b LoggerStopper, cc grpc.ClientConnInterface) *ServiceClient {
	return &ServiceClient{b, cc, pb.NewServiceClient(cc)}
}

func (s *ServiceClient) Start(ctx context.Context) error {
	return nil // no-op: server side starts automatically
}

func (s *ServiceClient) Close() error {
	ctx, cancel := s.StopCtx()
	defer cancel()

	_, err := s.grpc.Close(ctx, &emptypb.Empty{})
	return err
}

func (s *ServiceClient) Ready() error {
	ctx, cancel := s.StopCtx()
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, time.Second)
	defer cancel()

	_, err := s.grpc.Ready(ctx, &emptypb.Empty{})
	return err
}

func (s *ServiceClient) Name() string { return s.Logger().Name() }

func (s *ServiceClient) HealthReport() map[string]error {
	ctx, cancel := s.StopCtx()
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, time.Second)
	defer cancel()

	reply, err := s.grpc.HealthReport(ctx, &emptypb.Empty{})
	if err != nil {
		return map[string]error{s.Logger().Name(): err}
	}
	hr := healthReport(reply.HealthReport)
	hr[s.Logger().Name()] = nil
	return hr
}

var _ pb.ServiceServer = (*serviceServer)(nil)

type serviceServer struct {
	pb.UnimplementedServiceServer
	srv services.Service
}

func (s *serviceServer) Close(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.srv.Close()
}

func (s *serviceServer) Ready(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.srv.Ready()
}

func (s *serviceServer) HealthReport(ctx context.Context, empty *emptypb.Empty) (*pb.HealthReportReply, error) {
	var r pb.HealthReportReply
	r.HealthReport = make(map[string]string)
	for n, err := range s.srv.HealthReport() {
		var serr string
		if err != nil {
			serr = err.Error()
		}
		r.HealthReport[n] = serr
	}
	return &r, nil
}

func healthReport(s map[string]string) (hr map[string]error) {
	hr = make(map[string]error, len(s))
	for n, e := range s {
		var err error
		if e != "" {
			err = errors.New(e)
		}
		hr[n] = err
	}
	return hr
}
