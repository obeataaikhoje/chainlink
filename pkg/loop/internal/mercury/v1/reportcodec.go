package v1

import (
	"context"

	"google.golang.org/grpc"

	mercury_v1_types "github.com/smartcontractkit/chainlink-common/pkg/types/mercury/v1"

	ocr2plus_types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb"
	mercury_v1_pb "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb/mercury/v1"
)

var _ mercury_v1_types.ReportCodec = (*ReportCodecClient)(nil)

type ReportCodecClient struct {
	grpc mercury_v1_pb.ReportCodecClient
}

func NewReportCodecClient(cc grpc.ClientConnInterface) *ReportCodecClient {
	return &ReportCodecClient{grpc: mercury_v1_pb.NewReportCodecClient(cc)}
}

func (r *ReportCodecClient) BuildReport(fields mercury_v1_types.ReportFields) (ocr2plus_types.Report, error) {
	Response, err := r.grpc.BuildReport(context.TODO(), &mercury_v1_pb.BuildReportRequest{
		ReportFields: pbReportFields(fields),
	})
	if err != nil {
		return ocr2plus_types.Report{}, err
	}
	return Response.Report, nil
}

func (r *ReportCodecClient) MaxReportLength(n int) (int, error) {
	Response, err := r.grpc.MaxReportLength(context.TODO(), &mercury_v1_pb.MaxReportLengthRequest{})
	if err != nil {
		return 0, err
	}
	return int(Response.MaxReportLength), nil
}

func (r *ReportCodecClient) CurrentBlockNumFromReport(report ocr2plus_types.Report) (int64, error) {
	Response, err := r.grpc.CurrentBlockNumFromReport(context.TODO(), &mercury_v1_pb.CurrentBlockNumFromReportRequest{
		Report: report,
	})
	if err != nil {
		return 0, err
	}
	return Response.CurrentBlockNum, nil
}

func pbReportFields(fields mercury_v1_types.ReportFields) *mercury_v1_pb.ReportFields {
	return &mercury_v1_pb.ReportFields{
		Timestamp:             fields.Timestamp,
		BenchmarkPrice:        pb.NewBigIntFromInt(fields.BenchmarkPrice),
		Ask:                   pb.NewBigIntFromInt(fields.Ask),
		Bid:                   pb.NewBigIntFromInt(fields.Bid),
		CurrentBlockNum:       fields.CurrentBlockNum,
		CurrentBlockHash:      fields.CurrentBlockHash,
		ValidFromBlockNum:     fields.ValidFromBlockNum,
		CurrentBlockTimestamp: fields.CurrentBlockTimestamp,
	}
}

var _ mercury_v1_pb.ReportCodecServer = (*ReportCodecServer)(nil)

type ReportCodecServer struct {
	mercury_v1_pb.UnimplementedReportCodecServer
	impl mercury_v1_types.ReportCodec
}

func NewReportCodecServer(impl mercury_v1_types.ReportCodec) *ReportCodecServer {
	return &ReportCodecServer{impl: impl}
}

func (r *ReportCodecServer) BuildReport(ctx context.Context, request *mercury_v1_pb.BuildReportRequest) (*mercury_v1_pb.BuildReportResponse, error) {
	report, err := r.impl.BuildReport(reportFields(request.ReportFields))
	if err != nil {
		return nil, err
	}
	return &mercury_v1_pb.BuildReportResponse{Report: report}, nil
}

func (r *ReportCodecServer) MaxReportLength(ctx context.Context, request *mercury_v1_pb.MaxReportLengthRequest) (*mercury_v1_pb.MaxReportLengthResponse, error) {
	n, err := r.impl.MaxReportLength(int(request.NumOracles))
	if err != nil {
		return nil, err
	}
	return &mercury_v1_pb.MaxReportLengthResponse{MaxReportLength: uint64(n)}, nil
}

func (r *ReportCodecServer) CurrentBlockNumFromReport(ctx context.Context, request *mercury_v1_pb.CurrentBlockNumFromReportRequest) (*mercury_v1_pb.CurrentBlockNumFromReportResponse, error) {
	n, err := r.impl.CurrentBlockNumFromReport(request.Report)
	if err != nil {
		return nil, err
	}
	return &mercury_v1_pb.CurrentBlockNumFromReportResponse{CurrentBlockNum: n}, nil
}

func reportFields(fields *mercury_v1_pb.ReportFields) mercury_v1_types.ReportFields {
	return mercury_v1_types.ReportFields{
		Timestamp:             fields.Timestamp,
		BenchmarkPrice:        fields.BenchmarkPrice.Int(),
		Ask:                   fields.Ask.Int(),
		Bid:                   fields.Bid.Int(),
		CurrentBlockNum:       fields.CurrentBlockNum,
		CurrentBlockHash:      fields.CurrentBlockHash,
		ValidFromBlockNum:     fields.ValidFromBlockNum,
		CurrentBlockTimestamp: fields.CurrentBlockTimestamp,
	}
}
