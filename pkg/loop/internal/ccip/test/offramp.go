package test

import (
	"context"
	"fmt"
	"reflect"

	"github.com/stretchr/testify/assert"

	testtypes "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/test/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
)

type OffRampEvaluator interface {
	ccip.OffRampReader
	testtypes.Evaluator[ccip.OffRampReader]
}

var _ OffRampEvaluator = staticOffRamp{}

type staticOffRampConfig struct {
	addressResponse ccip.Address
	changeConfigRequest
	changeConfigResponse

	currentRateLimiterStateResponse ccip.TokenBucketRateLimit

	decodeExecutionReportResponse ccip.ExecReport

	encodeExecutionReportResponse []byte

	gasPriceEstimatorExecResponse ccip.GasPriceEstimatorExec

	getExecutionStateRequest  uint64
	getExecutionStateResponse uint8

	getExecutionStateChangesBetweenSeqNumsRequest
	getExecutionStateChangesBetweenSeqNumsResponse

	getSenderNonceRequest  ccip.Address
	getSenderNonceResponse uint64

	getSourceToDestTokensMappingResponse map[ccip.Address]ccip.Address

	getStaticConfigResponse ccip.OffRampStaticConfig

	getTokensResponse ccip.OffRampTokens

	offchainConfigResponse ccip.ExecOffchainConfig

	onchainConfigResponse ccip.ExecOnchainConfig
}

type staticOffRamp struct {
	staticOffRampConfig
}

// Address implements OffRampEvaluator.
func (s staticOffRamp) Address(ctx context.Context) (ccip.Address, error) {
	return s.addressResponse, nil
}

// ChangeConfig implements OffRampEvaluator.
func (s staticOffRamp) ChangeConfig(ctx context.Context, onchainConfig []byte, offchainConfig []byte) (ccip.Address, ccip.Address, error) {
	if !reflect.DeepEqual(onchainConfig, s.onchainConfig) {
		return ccip.Address(""), ccip.Address(""), fmt.Errorf("expected onchainConfig %v but got %v", s.onchainConfig, onchainConfig)
	}
	if !reflect.DeepEqual(offchainConfig, s.offchainConfig) {
		return ccip.Address(""), ccip.Address(""), fmt.Errorf("expected offchainConfig %v but got %v", s.offchainConfig, offchainConfig)
	}
	return s.onchainConfigDigest, s.offchainConfigDigest, nil
}

// CurrentRateLimiterState implements OffRampEvaluator.
func (s staticOffRamp) CurrentRateLimiterState(ctx context.Context) (ccip.TokenBucketRateLimit, error) {
	return s.currentRateLimiterStateResponse, nil
}

// DecodeExecutionReport implements OffRampEvaluator.
func (s staticOffRamp) DecodeExecutionReport(ctx context.Context, report []byte) (ccip.ExecReport, error) {
	if !reflect.DeepEqual(report, s.decodeExecutionReportResponse) {
		return ccip.ExecReport{}, fmt.Errorf("expected report %v but got %v", s.decodeExecutionReportResponse, report)
	}
	return s.decodeExecutionReportResponse, nil
}

// EncodeExecutionReport implements OffRampEvaluator.
func (s staticOffRamp) EncodeExecutionReport(ctx context.Context, report ccip.ExecReport) ([]byte, error) {
	if !reflect.DeepEqual(report, s.encodeExecutionReportResponse) {
		return nil, fmt.Errorf("expected report %v but got %v", s.encodeExecutionReportResponse, report)
	}
	return s.encodeExecutionReportResponse, nil
}

// GasPriceEstimator implements OffRampEvaluator.
func (s staticOffRamp) GasPriceEstimator(ctx context.Context) (ccip.GasPriceEstimatorExec, error) {
	return s.gasPriceEstimatorExecResponse, nil
}

// GetExecutionState implements OffRampEvaluator.
func (s staticOffRamp) GetExecutionState(ctx context.Context, sequenceNumber uint64) (uint8, error) {
	if sequenceNumber != s.getExecutionStateRequest {
		return 0, fmt.Errorf("expected sequenceNumber %d but got %d", s.getExecutionStateRequest, sequenceNumber)
	}
	return s.getExecutionStateResponse, nil
}

// GetExecutionStateChangesBetweenSeqNums implements OffRampEvaluator.
func (s staticOffRamp) GetExecutionStateChangesBetweenSeqNums(ctx context.Context, seqNumMin uint64, seqNumMax uint64, confirmations int) ([]ccip.ExecutionStateChangedWithTxMeta, error) {
	if seqNumMin != s.seqNumMin {
		return nil, fmt.Errorf("expected seqNumMin %d but got %d", s.seqNumMin, seqNumMin)
	}
	if seqNumMax != s.seqNumMax {
		return nil, fmt.Errorf("expected seqNumMax %d but got %d", s.seqNumMax, seqNumMax)
	}
	if confirmations != s.confirmations {
		return nil, fmt.Errorf("expected confirmations %d but got %d", s.confirmations, confirmations)
	}
	return s.executionStateChangedWithTxMeta, nil
}

// GetSenderNonce implements OffRampEvaluator.
func (s staticOffRamp) GetSenderNonce(ctx context.Context, sender ccip.Address) (uint64, error) {
	if sender != s.getSenderNonceRequest {
		return 0, fmt.Errorf("expected sender %s but got %s", s.getSenderNonceRequest, sender)
	}
	return s.getSenderNonceResponse, nil
}

// GetSourceToDestTokensMapping implements OffRampEvaluator.
func (s staticOffRamp) GetSourceToDestTokensMapping(ctx context.Context) (map[ccip.Address]ccip.Address, error) {
	return s.getSourceToDestTokensMappingResponse, nil
}

// GetStaticConfig implements OffRampEvaluator.
func (s staticOffRamp) GetStaticConfig(ctx context.Context) (ccip.OffRampStaticConfig, error) {
	return s.getStaticConfigResponse, nil
}

// GetTokens implements OffRampEvaluator.
func (s staticOffRamp) GetTokens(ctx context.Context) (ccip.OffRampTokens, error) {
	return s.getTokensResponse, nil
}

// OffchainConfig implements OffRampEvaluator.
func (s staticOffRamp) OffchainConfig(ctx context.Context) (ccip.ExecOffchainConfig, error) {
	return s.offchainConfigResponse, nil
}

// OnchainConfig implements OffRampEvaluator.
func (s staticOffRamp) OnchainConfig(ctx context.Context) (ccip.ExecOnchainConfig, error) {
	return s.onchainConfigResponse, nil
}

// Evaluate implements OffRampEvaluator.
func (s staticOffRamp) Evaluate(ctx context.Context, other ccip.OffRampReader) error {
	address, err := other.Address(ctx)
	if err != nil {
		return fmt.Errorf("failed to get address: %w", err)
	}
	if address != s.addressResponse {
		return fmt.Errorf("expected address %s but got %s", s.addressResponse, address)
	}

	currentRateLimiterState, err := other.CurrentRateLimiterState(ctx)
	if err != nil {
		return fmt.Errorf("failed to get currentRateLimiterState: %w", err)
	}
	if currentRateLimiterState != s.currentRateLimiterStateResponse {
		return fmt.Errorf("expected currentRateLimiterState %v but got %v", s.currentRateLimiterStateResponse, currentRateLimiterState)
	}

	decodeExecutionReport, err := other.DecodeExecutionReport(ctx, s.encodeExecutionReportResponse)
	if err != nil {
		return fmt.Errorf("failed to decodeExecutionReport: %w", err)
	}
	if !assert.ObjectsAreEqual(decodeExecutionReport, s.decodeExecutionReportResponse) {
		return fmt.Errorf("expected decodeExecutionReport %v but got %v", s.decodeExecutionReportResponse, decodeExecutionReport)
	}

	encodeExecutionReport, err := other.EncodeExecutionReport(ctx, s.decodeExecutionReportResponse)
	if err != nil {
		return fmt.Errorf("failed to encodeExecutionReport: %w", err)
	}
	if !reflect.DeepEqual(encodeExecutionReport, s.encodeExecutionReportResponse) {
		return fmt.Errorf("expected encodeExecutionReport %v but got %v", s.encodeExecutionReportResponse, encodeExecutionReport)
	}

	gasPriceEstimator, err := other.GasPriceEstimator(ctx)
	if err != nil {
		return fmt.Errorf("failed to get gasPriceEstimator: %w", err)
	}
	if gasPriceEstimator != s.gasPriceEstimatorExecResponse {
		return fmt.Errorf("expected gasPriceEstimator %v but got %v", s.gasPriceEstimatorExecResponse, gasPriceEstimator)
	}

	getExecutionState, err := other.GetExecutionState(ctx, s.getExecutionStateRequest)
	if err != nil {
		return fmt.Errorf("failed to get getExecutionState: %w", err)
	}
	if getExecutionState != s.getExecutionStateResponse {
		return fmt.Errorf("expected getExecutionState %d but got %d", s.getExecutionStateResponse, getExecutionState)
	}

	getExecutionStateChangesBetweenSeqNums, err := other.GetExecutionStateChangesBetweenSeqNums(ctx, s.seqNumMin, s.seqNumMax, s.confirmations)
	if err != nil {
		return fmt.Errorf("failed to get getExecutionStateChangesBetweenSeqNums: %w", err)
	}
	if !reflect.DeepEqual(getExecutionStateChangesBetweenSeqNums, s.executionStateChangedWithTxMeta) {
		return fmt.Errorf("expected getExecutionStateChangesBetweenSeqNums %v but got %v", s.executionStateChangedWithTxMeta, getExecutionStateChangesBetweenSeqNums)
	}

	getSenderNonce, err := other.GetSenderNonce(ctx, s.getSenderNonceRequest)
	if err != nil {
		return fmt.Errorf("failed to get getSenderNonce: %w", err)
	}
	if getSenderNonce != s.getSenderNonceResponse {
		return fmt.Errorf("expected getSenderNonce %d but got %d", s.getSenderNonceResponse, getSenderNonce)
	}

	getSourceToDestTokensMapping, err := other.GetSourceToDestTokensMapping(ctx)
	if err != nil {
		return fmt.Errorf("failed to get getSourceToDestTokensMapping: %w", err)
	}
	if !reflect.DeepEqual(getSourceToDestTokensMapping, s.getSourceToDestTokensMappingResponse) {
		return fmt.Errorf("expected getSourceToDestTokensMapping %v but got %v", s.getSourceToDestTokensMappingResponse, getSourceToDestTokensMapping)
	}

	getStaticConfig, err := other.GetStaticConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to get getStaticConfig: %w", err)
	}
	if getStaticConfig != s.getStaticConfigResponse {
		return fmt.Errorf("expected getStaticConfig %v but got %v", s.getStaticConfigResponse, getStaticConfig)
	}

	getTokens, err := other.GetTokens(ctx)
	if err != nil {
		return fmt.Errorf("failed to get getTokens: %w", err)
	}
	if !assert.ObjectsAreEqual(getTokens, s.getTokensResponse) {
		return fmt.Errorf("expected getTokens %v but got %v", s.getTokensResponse, getTokens)
	}

	offchainConfig, err := other.OffchainConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to get offchainConfig: %w", err)
	}
	if offchainConfig != s.offchainConfigResponse {
		return fmt.Errorf("expected offchainConfig %v but got %v", s.offchainConfigResponse, offchainConfig)
	}

	onchainConfig, err := other.OnchainConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to get onchainConfig: %w", err)
	}
	if onchainConfig != s.onchainConfigResponse {
		return fmt.Errorf("expected onchainConfig %v but got %v", s.onchainConfigResponse, onchainConfig)
	}

	return nil
}

type changeConfigRequest struct {
	onchainConfig  []byte
	offchainConfig []byte
}

type changeConfigResponse struct {
	onchainConfigDigest  ccip.Address
	offchainConfigDigest ccip.Address
}

type getExecutionStateChangesBetweenSeqNumsRequest struct {
	seqNumMin     uint64
	seqNumMax     uint64
	confirmations int
}

type getExecutionStateChangesBetweenSeqNumsResponse struct {
	executionStateChangedWithTxMeta []ccip.ExecutionStateChangedWithTxMeta
}
