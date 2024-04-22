package errorlog

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb"
	"github.com/smartcontractkit/chainlink-common/pkg/types/core"
)

var _ core.ErrorLog = (*errorLogClient)(nil)

type errorLogClient struct {
	grpc pb.ErrorLogClient
}

func (e errorLogClient) SaveError(ctx context.Context, msg string) error {
	_, err := e.grpc.SaveError(ctx, &pb.SaveErrorRequest{Message: msg})
	return err
}

func NewClient(cc grpc.ClientConnInterface) *errorLogClient {
	return &errorLogClient{pb.NewErrorLogClient(cc)}
}

var _ pb.ErrorLogServer = (*Server)(nil)

type Server struct {
	pb.UnimplementedErrorLogServer

	impl core.ErrorLog
}

func NewServer(impl core.ErrorLog) *Server {
	return &Server{impl: impl}
}

func (e *Server) SaveError(ctx context.Context, request *pb.SaveErrorRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, e.impl.SaveError(ctx, request.Message)
}
