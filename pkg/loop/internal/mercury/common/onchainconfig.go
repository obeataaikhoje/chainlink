package common

import (
	"context"

	"google.golang.org/grpc"

	"github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb"
	mercury_pb "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb/mercury"
	mercury_types "github.com/smartcontractkit/chainlink-common/pkg/types/mercury"
)

var _ mercury_types.OnchainConfigCodec = (*OnchainConfigCodecClient)(nil)

type OnchainConfigCodecClient struct {
	grpc mercury_pb.OnchainConfigCodecClient
}

func NewOnchainConfigCodecClient(cc grpc.ClientConnInterface) *OnchainConfigCodecClient {
	return &OnchainConfigCodecClient{grpc: mercury_pb.NewOnchainConfigCodecClient(cc)}
}

// TODO: why doesn't the interface have a context?!
func (o *OnchainConfigCodecClient) Encode(config mercury_types.OnchainConfig) ([]byte, error) {
	Response, err := o.grpc.Encode(context.TODO(), &mercury_pb.EncodeOnchainConfigRequest{
		OnchainConfig: pbOnchainConfig(config),
	})
	if err != nil {
		return nil, err
	}
	return Response.OnchainConfig, nil
}

func (o *OnchainConfigCodecClient) Decode(data []byte) (mercury_types.OnchainConfig, error) {
	Response, err := o.grpc.Decode(context.TODO(), &mercury_pb.DecodeOnchainConfigRequest{
		OnchainConfig: data,
	})
	if err != nil {
		return mercury_types.OnchainConfig{}, err
	}
	return onchainConfig(Response.OnchainConfig), nil
}

func pbOnchainConfig(config mercury_types.OnchainConfig) *mercury_pb.OnchainConfig {
	return &mercury_pb.OnchainConfig{
		Min: pb.NewBigIntFromInt(config.Min),
		Max: pb.NewBigIntFromInt(config.Max),
	}
}

func onchainConfig(config *mercury_pb.OnchainConfig) mercury_types.OnchainConfig {
	return mercury_types.OnchainConfig{
		Min: config.Min.Int(),
		Max: config.Max.Int(),
	}
}

var _ mercury_pb.OnchainConfigCodecServer = (*OnchainConfigCodecServer)(nil)

type OnchainConfigCodecServer struct {
	mercury_pb.UnimplementedOnchainConfigCodecServer

	impl mercury_types.OnchainConfigCodec
}

func NewOnchainConfigCodecServer(impl mercury_types.OnchainConfigCodec) *OnchainConfigCodecServer {
	return &OnchainConfigCodecServer{impl: impl}
}

func (o *OnchainConfigCodecServer) Encode(ctx context.Context, request *mercury_pb.EncodeOnchainConfigRequest) (*mercury_pb.EncodeOnchainConfigResponse, error) {
	val, err := o.impl.Encode(onchainConfig(request.OnchainConfig))
	if err != nil {
		return nil, err
	}
	return &mercury_pb.EncodeOnchainConfigResponse{OnchainConfig: val}, nil
}

func (o *OnchainConfigCodecServer) Decode(ctx context.Context, request *mercury_pb.DecodeOnchainConfigRequest) (*mercury_pb.DecodeOnchainConfigResponse, error) {
	val, err := o.impl.Decode(request.OnchainConfig)
	if err != nil {
		return nil, err
	}
	return &mercury_pb.DecodeOnchainConfigResponse{OnchainConfig: pbOnchainConfig(val)}, nil
}
