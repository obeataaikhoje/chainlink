package ccip

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb"
	ccippb "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb/ccip"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
)

var _ cciptypes.OffRampReader = (*OffRampReaderClient)(nil)

type OffRampReaderClient struct {
	grpc ccippb.OffRampReaderClient
}

func NewOffRampReaderClient(grpc ccippb.OffRampReaderClient) *OffRampReaderClient {
	return &OffRampReaderClient{grpc: grpc}
}

// Address i[github.com/smartcontractkit/chainlink-common/pkg/types/ccip.OffRampReader]
func (o *OffRampReaderClient) Address(ctx context.Context) (cciptypes.Address, error) {
	resp, err := o.grpc.Address(context.TODO(), &emptypb.Empty{})
	if err != nil {
		return cciptypes.Address(""), err
	}
	return cciptypes.Address(resp.Address), nil
}

// ChangeConfig implements [github.com/smartcontractkit/chainlink-common/pkg/types/ccip.OffRampReader]
func (o *OffRampReaderClient) ChangeConfig(ctx context.Context, onchainConfig []byte, offchainConfig []byte) (cciptypes.Address, cciptypes.Address, error) {
	resp, err := o.grpc.ChangeConfig(ctx, &ccippb.ChangeConfigRequest{
		OnchainConfig:  onchainConfig,
		OffchainConfig: offchainConfig,
	})
	if err != nil {
		return cciptypes.Address(""), cciptypes.Address(""), err
	}

	return cciptypes.Address(resp.OnchainConfigAddress), cciptypes.Address(resp.OffchainConfigAddress), nil
}

// CurrentRateLimiterState i[github.com/smartcontractkit/chainlink-common/pkg/types/ccip.OffRampReader]
func (o *OffRampReaderClient) CurrentRateLimiterState(ctx context.Context) (cciptypes.TokenBucketRateLimit, error) {
	resp, err := o.grpc.CurrentRateLimiterState(ctx, &emptypb.Empty{})
	if err != nil {
		return cciptypes.TokenBucketRateLimit{}, err
	}
	return tokenBucketRateLimit(resp.RateLimiter), nil
}

// DecodeExecutionReport i[github.com/smartcontractkit/chainlink-common/pkg/types/ccip.OffRampReader]
func (o *OffRampReaderClient) DecodeExecutionReport(ctx context.Context, report []byte) (cciptypes.ExecReport, error) {
	resp, err := o.grpc.DecodeExecutionReport(ctx, &ccippb.DecodeExecutionReportRequest{
		Report: report,
	})
	if err != nil {
		return cciptypes.ExecReport{}, err
	}

	return execReport(resp.Report)
}

// EncodeExecutionReport i[github.com/smartcontractkit/chainlink-common/pkg/types/ccip.OffRampReader]
func (o *OffRampReaderClient) EncodeExecutionReport(ctx context.Context, report cciptypes.ExecReport) ([]byte, error) {
	reportPB := executionReportPB(report)

	resp, err := o.grpc.EncodeExecutionReport(ctx, &ccippb.EncodeExecutionReportRequest{
		Report: reportPB,
	})
	if err != nil {
		return nil, err
	}
	return resp.Report, nil
}

// GasPriceEstimator i[github.com/smartcontractkit/chainlink-common/pkg/types/ccip.OffRampReader]
func (o *OffRampReaderClient) GasPriceEstimator(ctx context.Context) (cciptypes.GasPriceEstimatorExec, error) {
	panic("BCF-2991 implement gas price estimator grpc service")
}

// GetExecutionState i[github.com/smartcontractkit/chainlink-common/pkg/types/ccip.OffRampReader]
func (o *OffRampReaderClient) GetExecutionState(ctx context.Context, sequenceNumber uint64) (uint8, error) {
	resp, err := o.grpc.GetExecutionState(ctx, &ccippb.GetExecutionStateRequest{
		SeqNum: sequenceNumber,
	})
	if err != nil {
		return 0, err
	}
	return uint8(resp.ExecutionState), nil
}

// GetExecutionStateChangesBetweenSeqNums i[github.com/smartcontractkit/chainlink-common/pkg/types/ccip.OffRampReader]
func (o *OffRampReaderClient) GetExecutionStateChangesBetweenSeqNums(ctx context.Context, seqNumMin uint64, seqNumMax uint64, confirmations int) ([]cciptypes.ExecutionStateChangedWithTxMeta, error) {
	resp, err := o.grpc.GetExecutionStateChanges(ctx, &ccippb.GetExecutionStateChangesRequest{
		MinSeqNum:     seqNumMin,
		MaxSeqNum:     seqNumMax,
		Confirmations: int64(confirmations),
	})
	if err != nil {
		return nil, err
	}
	return executionStateChangedWithTxMetaSlice(resp.ExecutionStateChanges), nil
}

// GetSenderNonce i[github.com/smartcontractkit/chainlink-common/pkg/types/ccip.OffRampReader]
func (o *OffRampReaderClient) GetSenderNonce(ctx context.Context, sender cciptypes.Address) (uint64, error) {
	resp, err := o.grpc.GetSenderNonce(ctx, &ccippb.GetSenderNonceRequest{
		Sender: string(sender),
	})
	if err != nil {
		return 0, err
	}
	return resp.Nonce, nil
}

// GetSourceToDestTokensMapping i[github.com/smartcontractkit/chainlink-common/pkg/types/ccip.OffRampReader]
func (o *OffRampReaderClient) GetSourceToDestTokensMapping(ctx context.Context) (map[cciptypes.Address]cciptypes.Address, error) {
	resp, err := o.grpc.GetSourceToDestTokensMapping(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	return sourceToDestTokensMapping(resp.TokenMappings), nil
}

// GetStaticConfig i[github.com/smartcontractkit/chainlink-common/pkg/types/ccip.OffRampReader]
func (o *OffRampReaderClient) GetStaticConfig(ctx context.Context) (cciptypes.OffRampStaticConfig, error) {
	resp, err := o.grpc.GetStaticConfig(ctx, &emptypb.Empty{})
	if err != nil {
		return cciptypes.OffRampStaticConfig{}, err
	}
	return cciptypes.OffRampStaticConfig{
		CommitStore:         cciptypes.Address(resp.Config.CommitStore),
		ChainSelector:       resp.Config.ChainSelector,
		SourceChainSelector: resp.Config.SourceChainSelector,
		OnRamp:              cciptypes.Address(resp.Config.OnRamp),
		PrevOffRamp:         cciptypes.Address(resp.Config.PrevOffRamp),
		ArmProxy:            cciptypes.Address(resp.Config.ArmProxy),
	}, nil
}

// GetTokens i[github.com/smartcontractkit/chainlink-common/pkg/types/ccip.OffRampReader]
func (o *OffRampReaderClient) GetTokens(ctx context.Context) (cciptypes.OffRampTokens, error) {
	resp, err := o.grpc.GetTokens(ctx, &emptypb.Empty{})
	if err != nil {
		return cciptypes.OffRampTokens{}, err
	}
	return offRampTokens(resp.Tokens), nil
}

// OffchainConfig i[github.com/smartcontractkit/chainlink-common/pkg/types/ccip.OffRampReader]
func (o *OffRampReaderClient) OffchainConfig(ctx context.Context) (cciptypes.ExecOffchainConfig, error) {
	resp, err := o.grpc.OffchainConfig(ctx, &emptypb.Empty{})
	if err != nil {
		return cciptypes.ExecOffchainConfig{}, err
	}
	return offChainConfig(resp.Config)
}

// OnchainConfig i[github.com/smartcontractkit/chainlink-common/pkg/types/ccip.OffRampReader]
func (o *OffRampReaderClient) OnchainConfig(ctx context.Context) (cciptypes.ExecOnchainConfig, error) {
	resp, err := o.grpc.OnchainConfig(ctx, &emptypb.Empty{})
	if err != nil {
		return cciptypes.ExecOnchainConfig{}, err
	}
	return cciptypes.ExecOnchainConfig{
		PermissionLessExecutionThresholdSeconds: resp.Config.PermissionlessExecThresholdSeconds.AsDuration(),
	}, nil
}

// Server implementation of OffRampReader

type OffRampReaderServer struct {
	ccippb.UnimplementedOffRampReaderServer

	impl cciptypes.OffRampReader
}

var _ ccippb.OffRampReaderServer = (*OffRampReaderServer)(nil)

func NewOffRampReaderServer(impl cciptypes.OffRampReader) *OffRampReaderServer {
	return &OffRampReaderServer{impl: impl}
}

// Address implements ccippb.OffRampReaderServer.
func (o *OffRampReaderServer) Address(ctx context.Context, req *emptypb.Empty) (*ccippb.OffRampAddressResponse, error) {
	addr, err := o.impl.Address(ctx)
	if err != nil {
		return nil, err
	}
	return &ccippb.OffRampAddressResponse{Address: string(addr)}, nil
}

// ChangeConfig implements ccippb.OffRampReaderServer.
func (o *OffRampReaderServer) ChangeConfig(ctx context.Context, req *ccippb.ChangeConfigRequest) (*ccippb.ChangeConfigResponse, error) {
	onchainAddr, offchainAddr, err := o.impl.ChangeConfig(ctx, req.OnchainConfig, req.OffchainConfig)
	if err != nil {
		return nil, err
	}
	return &ccippb.ChangeConfigResponse{
		OnchainConfigAddress:  string(onchainAddr),
		OffchainConfigAddress: string(offchainAddr),
	}, nil
}

// CurrentRateLimiterState implements ccippb.OffRampReaderServer.
func (o *OffRampReaderServer) CurrentRateLimiterState(ctx context.Context, req *emptypb.Empty) (*ccippb.CurrentRateLimiterStateResponse, error) {
	state, err := o.impl.CurrentRateLimiterState(ctx)
	if err != nil {
		return nil, err
	}
	return &ccippb.CurrentRateLimiterStateResponse{RateLimiter: tokenBucketRateLimitPB(state)}, nil
}

// DecodeExecutionReport implements ccippb.OffRampReaderServer.
func (o *OffRampReaderServer) DecodeExecutionReport(ctx context.Context, req *ccippb.DecodeExecutionReportRequest) (*ccippb.DecodeExecutionReportResponse, error) {
	report, err := o.impl.DecodeExecutionReport(ctx, req.Report)
	if err != nil {
		return nil, err
	}
	return &ccippb.DecodeExecutionReportResponse{Report: executionReportPB(report)}, nil
}

// EncodeExecutionReport implements ccippb.OffRampReaderServer.
func (o *OffRampReaderServer) EncodeExecutionReport(ctx context.Context, req *ccippb.EncodeExecutionReportRequest) (*ccippb.EncodeExecutionReportResponse, error) {
	report, err := execReport(req.Report)
	if err != nil {
		return nil, err
	}

	encoded, err := o.impl.EncodeExecutionReport(ctx, report)
	if err != nil {
		return nil, err
	}
	return &ccippb.EncodeExecutionReportResponse{Report: encoded}, nil
}

// GasPriceEstimator implements ccippb.OffRampReaderServer.
func (o *OffRampReaderServer) GasPriceEstimator(ctx context.Context, req *emptypb.Empty) (*ccippb.GasPriceEstimatorResponse, error) {
	panic("BCF-2991 implement gas price estimator grpc service")
}

// GetExecutionState implements ccippb.OffRampReaderServer.
func (o *OffRampReaderServer) GetExecutionState(ctx context.Context, req *ccippb.GetExecutionStateRequest) (*ccippb.GetExecutionStateResponse, error) {
	state, err := o.impl.GetExecutionState(ctx, req.SeqNum)
	if err != nil {
		return nil, err
	}
	return &ccippb.GetExecutionStateResponse{ExecutionState: uint32(state)}, nil
}

// GetExecutionStateChanges implements ccippb.OffRampReaderServer.
func (o *OffRampReaderServer) GetExecutionStateChanges(ctx context.Context, req *ccippb.GetExecutionStateChangesRequest) (*ccippb.GetExecutionStateChangesResponse, error) {
	changes, err := o.impl.GetExecutionStateChangesBetweenSeqNums(ctx, req.MinSeqNum, req.MaxSeqNum, int(req.Confirmations))
	if err != nil {
		return nil, err
	}
	return &ccippb.GetExecutionStateChangesResponse{ExecutionStateChanges: executionStateChangedWithTxMetaSliceToPB(changes)}, nil
}

// GetSenderNonce implements ccippb.OffRampReaderServer.
func (o *OffRampReaderServer) GetSenderNonce(ctx context.Context, req *ccippb.GetSenderNonceRequest) (*ccippb.GetSenderNonceResponse, error) {
	nonce, err := o.impl.GetSenderNonce(ctx, cciptypes.Address(req.Sender))
	if err != nil {
		return nil, err
	}
	return &ccippb.GetSenderNonceResponse{Nonce: nonce}, nil
}

// GetSourceToDestTokensMapping implements ccippb.OffRampReaderServer.
func (o *OffRampReaderServer) GetSourceToDestTokensMapping(ctx context.Context, req *emptypb.Empty) (*ccippb.GetSourceToDestTokensMappingResponse, error) {
	mapping, err := o.impl.GetSourceToDestTokensMapping(ctx)
	if err != nil {
		return nil, err
	}
	return &ccippb.GetSourceToDestTokensMappingResponse{TokenMappings: sourceDestTokenMappingToPB(mapping)}, nil
}

// GetStaticConfig implements ccippb.OffRampReaderServer.
func (o *OffRampReaderServer) GetStaticConfig(ctx context.Context, req *emptypb.Empty) (*ccippb.GetStaticConfigResponse, error) {
	config, err := o.impl.GetStaticConfig(ctx)
	if err != nil {
		return nil, err
	}

	pbConfig := ccippb.OffRampStaticConfig{
		CommitStore:         string(config.CommitStore),
		ChainSelector:       config.ChainSelector,
		SourceChainSelector: config.SourceChainSelector,
		OnRamp:              string(config.OnRamp),
		PrevOffRamp:         string(config.PrevOffRamp),
		ArmProxy:            string(config.ArmProxy),
	}
	return &ccippb.GetStaticConfigResponse{Config: &pbConfig}, nil
}

// GetTokens implements ccippb.OffRampReaderServer.
func (o *OffRampReaderServer) GetTokens(ctx context.Context, req *emptypb.Empty) (*ccippb.GetTokensResponse, error) {
	tokens, err := o.impl.GetTokens(ctx)
	if err != nil {
		return nil, err
	}
	return &ccippb.GetTokensResponse{Tokens: offRampTokensToPB(tokens)}, nil
}

// OffchainConfig implements ccippb.OffRampReaderServer.
func (o *OffRampReaderServer) OffchainConfig(ctx context.Context, req *emptypb.Empty) (*ccippb.OffchainConfigResponse, error) {
	config, err := o.impl.OffchainConfig(ctx)
	if err != nil {
		return nil, err
	}
	return &ccippb.OffchainConfigResponse{Config: offChainConfigToPB(config)}, nil
}

// OnchainConfig implements ccippb.OffRampReaderServer.
func (o *OffRampReaderServer) OnchainConfig(ctx context.Context, req *emptypb.Empty) (*ccippb.OnchainConfigResponse, error) {
	config, err := o.impl.OnchainConfig(ctx)
	if err != nil {
		return nil, err
	}
	pbConfig := ccippb.ExecOnchainConfig{
		PermissionlessExecThresholdSeconds: durationpb.New(config.PermissionLessExecutionThresholdSeconds),
	}
	return &ccippb.OnchainConfigResponse{Config: &pbConfig}, nil
}

// Conversion functions and helpers

func tokenBucketRateLimit(pb *ccippb.TokenPoolRateLimit) cciptypes.TokenBucketRateLimit {
	return cciptypes.TokenBucketRateLimit{
		Tokens:      pb.Tokens.Int(),
		LastUpdated: pb.LastUpdated,
		IsEnabled:   pb.IsEnabled,
		Capacity:    pb.Capacity.Int(),
		Rate:        pb.Rate.Int(),
	}
}

func tokenBucketRateLimitPB(state cciptypes.TokenBucketRateLimit) *ccippb.TokenPoolRateLimit {
	return &ccippb.TokenPoolRateLimit{
		Tokens:      pb.NewBigIntFromInt(state.Tokens),
		LastUpdated: state.LastUpdated,
		IsEnabled:   state.IsEnabled,
		Capacity:    pb.NewBigIntFromInt(state.Capacity),
		Rate:        pb.NewBigIntFromInt(state.Rate),
	}
}

func execReport(pb *ccippb.ExecutionReport) (cciptypes.ExecReport, error) {
	proofs, err := byte32Slice(pb.Proofs)
	if err != nil {
		return cciptypes.ExecReport{}, fmt.Errorf("execReport: invalid proofs: %w", err)
	}
	msgs, err := evm2EVMMessageSlice(pb.EvmToEvmMessages)
	if err != nil {
		return cciptypes.ExecReport{}, fmt.Errorf("execReport: invalid messages: %w", err)
	}

	return cciptypes.ExecReport{
		Messages:          msgs,
		OffchainTokenData: offchainTokenData(pb.OffchainTokenData),
		Proofs:            proofs,
		ProofFlagBits:     pb.ProofFlagBits.Int(),
	}, nil
}

func evm2EVMMessageSlice(in []*ccippb.EVM2EVMMessage) ([]cciptypes.EVM2EVMMessage, error) {
	out := make([]cciptypes.EVM2EVMMessage, len(in))
	for i, m := range in {
		decodedMsg, err := evm2EVMMessage(m)
		if err != nil {
			return nil, err
		}
		out[i] = decodedMsg
	}
	return out, nil
}

func offchainTokenData(in []*ccippb.TokenData) [][][]byte {
	out := make([][][]byte, len(in))
	for i, b := range in {
		out[i] = b.Data
	}
	return out
}

func byte32Slice(in [][]byte) ([][32]byte, error) {
	out := make([][32]byte, len(in))
	for i, b := range in {
		if len(b) != 32 {
			return nil, fmt.Errorf("byte32Slice: invalid length %d", len(b))
		}
		copy(out[i][:], b)
	}
	return out, nil
}

func executionReportPB(report cciptypes.ExecReport) *ccippb.ExecutionReport {
	return &ccippb.ExecutionReport{
		EvmToEvmMessages:  evm2EVMMessageSliceToPB(report.Messages),
		OffchainTokenData: offchainTokenDataToPB(report.OffchainTokenData),
		Proofs:            byte32SliceToPB(report.Proofs),
		ProofFlagBits:     pb.NewBigIntFromInt(report.ProofFlagBits),
	}
}

func evm2EVMMessageSliceToPB(in []cciptypes.EVM2EVMMessage) []*ccippb.EVM2EVMMessage {
	out := make([]*ccippb.EVM2EVMMessage, len(in))
	for i, m := range in {
		out[i] = evm2EVMMessageToPB(m)
	}
	return out
}

func offchainTokenDataToPB(in [][][]byte) []*ccippb.TokenData {
	out := make([]*ccippb.TokenData, len(in))
	for i, b := range in {
		out[i] = &ccippb.TokenData{Data: b}
	}
	return out
}

func byte32SliceToPB(in [][32]byte) [][]byte {
	out := make([][]byte, len(in))
	for i, b := range in {
		out[i] = b[:]
	}
	return out
}

func evm2EVMMessageToPB(m cciptypes.EVM2EVMMessage) *ccippb.EVM2EVMMessage {
	return &ccippb.EVM2EVMMessage{
		SequenceNumber:      m.SequenceNumber,
		GasLimit:            pb.NewBigIntFromInt(m.GasLimit),
		Nonce:               m.Nonce,
		MessageId:           m.MessageID[:],
		SourceChainSelector: m.SourceChainSelector,
		Sender:              string(m.Sender),
		Receiver:            string(m.Receiver),
		Strict:              m.Strict,
		FeeToken:            string(m.FeeToken),
		FeeTokenAmount:      pb.NewBigIntFromInt(m.FeeTokenAmount),
		Data:                m.Data,
		TokenAmounts:        tokenAmountSliceToPB(m.TokenAmounts),
		SourceTokenData:     m.SourceTokenData,
	}
}

func tokenAmountSliceToPB(tokenAmounts []cciptypes.TokenAmount) []*ccippb.TokenAmount {
	res := make([]*ccippb.TokenAmount, len(tokenAmounts))
	for i, t := range tokenAmounts {
		res[i] = &ccippb.TokenAmount{
			Token:  string(t.Token),
			Amount: pb.NewBigIntFromInt(t.Amount),
		}
	}
	return res
}

func executionStateChangedWithTxMetaSlice(in []*ccippb.ExecutionStateChangeWithTxMeta) []cciptypes.ExecutionStateChangedWithTxMeta {
	out := make([]cciptypes.ExecutionStateChangedWithTxMeta, len(in))
	for i, m := range in {
		out[i] = executionStateChangedWithTxMeta(m)
	}
	return out
}

func executionStateChangedWithTxMetaSliceToPB(in []cciptypes.ExecutionStateChangedWithTxMeta) []*ccippb.ExecutionStateChangeWithTxMeta {
	out := make([]*ccippb.ExecutionStateChangeWithTxMeta, len(in))
	for i, m := range in {
		out[i] = executionStateChangedWithTxMetaToPB(m)
	}
	return out
}

func executionStateChangedWithTxMeta(in *ccippb.ExecutionStateChangeWithTxMeta) cciptypes.ExecutionStateChangedWithTxMeta {
	return cciptypes.ExecutionStateChangedWithTxMeta{
		TxMeta: txMeta(in.TxMeta),
		ExecutionStateChanged: cciptypes.ExecutionStateChanged{
			SequenceNumber: in.ExecutionStateChange.SeqNum,
			Finalized:      in.ExecutionStateChange.Finalized,
		},
	}
}

func executionStateChangedWithTxMetaToPB(in cciptypes.ExecutionStateChangedWithTxMeta) *ccippb.ExecutionStateChangeWithTxMeta {
	return &ccippb.ExecutionStateChangeWithTxMeta{
		TxMeta: txMetaToPB(in.TxMeta),
		ExecutionStateChange: &ccippb.ExecutionStateChange{
			SeqNum:    in.ExecutionStateChanged.SequenceNumber,
			Finalized: in.ExecutionStateChanged.Finalized,
		},
	}
}
func txMetaToPB(in cciptypes.TxMeta) *ccippb.TxMeta {
	return &ccippb.TxMeta{
		BlockTimestampUnixMilli: in.BlockTimestampUnixMilli,
		BlockNumber:             in.BlockNumber,
		TxHash:                  in.TxHash,
		LogIndex:                in.LogIndex,
	}
}

func offRampTokens(in *ccippb.OffRampTokens) cciptypes.OffRampTokens {
	source := make([]cciptypes.Address, len(in.SourceTokens))
	for i, t := range in.SourceTokens {
		source[i] = cciptypes.Address(t)
	}
	dest := make([]cciptypes.Address, len(in.DestinationTokens))
	for i, t := range in.DestinationTokens {
		dest[i] = cciptypes.Address(t)
	}
	destPool := make(map[cciptypes.Address]cciptypes.Address)
	for k, v := range in.DestinationPool {
		destPool[cciptypes.Address(k)] = cciptypes.Address(v)
	}

	return cciptypes.OffRampTokens{
		SourceTokens:      source,
		DestinationTokens: dest,
		DestinationPool:   destPool,
	}
}

func offRampTokensToPB(in cciptypes.OffRampTokens) *ccippb.OffRampTokens {
	source := make([]string, len(in.SourceTokens))
	for i, t := range in.SourceTokens {
		source[i] = string(t)
	}
	dest := make([]string, len(in.DestinationTokens))
	for i, t := range in.DestinationTokens {
		dest[i] = string(t)
	}
	destPool := make(map[string]string)
	for k, v := range in.DestinationPool {
		destPool[string(k)] = string(v)
	}

	return &ccippb.OffRampTokens{
		SourceTokens:      source,
		DestinationTokens: dest,
		DestinationPool:   destPool,
	}
}

func offChainConfig(in *ccippb.ExecOffchainConfig) (cciptypes.ExecOffchainConfig, error) {
	cachedExpiry, err := config.NewDuration(in.InflightCacheExpiry.AsDuration())
	if err != nil {
		return cciptypes.ExecOffchainConfig{}, fmt.Errorf("offChainConfig: invalid InflightCacheExpiry: %w", err)
	}
	rootSnoozeTime, err := config.NewDuration(in.RootSnoozeTime.AsDuration())
	if err != nil {
		return cciptypes.ExecOffchainConfig{}, fmt.Errorf("offChainConfig: invalid RootSnoozeTime: %w", err)
	}

	return cciptypes.ExecOffchainConfig{
		DestOptimisticConfirmations: in.DestOptimisticConfirmations,
		BatchGasLimit:               in.BatchGasLimit,
		RelativeBoostPerWaitHour:    in.RelativeBoostPerWaitHour,
		InflightCacheExpiry:         cachedExpiry,
		RootSnoozeTime:              rootSnoozeTime,
	}, nil
}

func offChainConfigToPB(in cciptypes.ExecOffchainConfig) *ccippb.ExecOffchainConfig {
	return &ccippb.ExecOffchainConfig{
		DestOptimisticConfirmations: in.DestOptimisticConfirmations,
		BatchGasLimit:               in.BatchGasLimit,
		RelativeBoostPerWaitHour:    in.RelativeBoostPerWaitHour,
		InflightCacheExpiry:         durationpb.New(in.InflightCacheExpiry.Duration()),
		RootSnoozeTime:              durationpb.New(in.RootSnoozeTime.Duration()),
	}
}

func sourceToDestTokensMapping(in map[string]string) map[cciptypes.Address]cciptypes.Address {
	out := make(map[cciptypes.Address]cciptypes.Address)
	for k, v := range in {
		out[cciptypes.Address(k)] = cciptypes.Address(v)
	}
	return out
}

func sourceDestTokenMappingToPB(in map[cciptypes.Address]cciptypes.Address) map[string]string {
	out := make(map[string]string)
	for k, v := range in {
		out[string(k)] = string(v)
	}
	return out
}
