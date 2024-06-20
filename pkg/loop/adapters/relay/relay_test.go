package relay

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/libocr/offchainreporting2/reportingplugin/median"
	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

type staticMedianProvider struct {
}

var _ types.MedianProvider = staticMedianProvider{}

// ContractConfigTracker implements types.MedianProvider.
func (s staticMedianProvider) ContractConfigTracker() ocr2types.ContractConfigTracker {
	return nil
}

// ContractTransmitter implements types.MedianProvider.
func (s staticMedianProvider) ContractTransmitter() ocr2types.ContractTransmitter {
	return nil
}

// MedianContract implements types.MedianProvider.
func (s staticMedianProvider) MedianContract() median.MedianContract {
	return nil
}

// OffchainConfigDigester implements types.MedianProvider.
func (s staticMedianProvider) OffchainConfigDigester() ocr2types.OffchainConfigDigester {
	return nil
}

// OnchainConfigCodec implements types.MedianProvider.
func (s staticMedianProvider) OnchainConfigCodec() median.OnchainConfigCodec {
	return nil
}

// ReportCodec implements types.MedianProvider.
func (s staticMedianProvider) ReportCodec() median.ReportCodec {
	return nil
}

// ChainReader implements types.MedianProvider.
func (s staticMedianProvider) ChainReader() types.ContractReader {
	return nil
}

// Close implements types.MedianProvider.
func (s staticMedianProvider) Close() error {
	return nil
}

// Codec implements types.MedianProvider.
func (s staticMedianProvider) Codec() types.Codec {
	return nil
}

// HealthReport implements types.MedianProvider.
func (s staticMedianProvider) HealthReport() map[string]error {
	return nil
}

// Name implements types.MedianProvider.
func (s staticMedianProvider) Name() string {
	return ""
}

// Ready implements types.MedianProvider.
func (s staticMedianProvider) Ready() error {
	return nil
}

// Start implements types.MedianProvider.
func (s staticMedianProvider) Start(context.Context) error {
	return nil
}

type staticFunctionsProvider struct {
	types.FunctionsProvider
}

type staticMercuryProvider struct {
	types.MercuryProvider
}

type staticAutomationProvider struct {
	types.AutomationProvider
}

type staticPluginProvider struct {
	types.PluginProvider
}

type staticOCR3CapabilityProvider struct {
	types.OCR3CapabilityProvider
}

type staticCCIPCommitProvider struct {
	types.CCIPCommitProvider
}

type staticCCIPExecProvider struct {
	types.CCIPExecProvider
}

type mockRelayer struct {
	types.Relayer
}

func (m *mockRelayer) NewMedianProvider(rargs types.RelayArgs, pargs types.PluginArgs) (types.MedianProvider, error) {
	return staticMedianProvider{}, nil
}

func (m *mockRelayer) NewFunctionsProvider(rargs types.RelayArgs, pargs types.PluginArgs) (types.FunctionsProvider, error) {
	return staticFunctionsProvider{}, nil
}

func (m *mockRelayer) NewMercuryProvider(rargs types.RelayArgs, pargs types.PluginArgs) (types.MercuryProvider, error) {
	return staticMercuryProvider{}, nil
}

func (m *mockRelayer) NewAutomationProvider(rargs types.RelayArgs, pargs types.PluginArgs) (types.AutomationProvider, error) {
	return staticAutomationProvider{}, nil
}

func (m *mockRelayer) NewPluginProvider(rargs types.RelayArgs, pargs types.PluginArgs) (types.PluginProvider, error) {
	return staticPluginProvider{}, nil
}

func (m *mockRelayer) NewOCR3CapabilityProvider(rargs types.RelayArgs, pargs types.PluginArgs) (types.OCR3CapabilityProvider, error) {
	return staticOCR3CapabilityProvider{}, nil
}

func (m *mockRelayer) NewCCIPCommitProvider(rargs types.RelayArgs, pargs types.PluginArgs) (types.CCIPCommitProvider, error) {
	return staticCCIPCommitProvider{}, nil
}

func (m *mockRelayer) NewCCIPExecProvider(rargs types.RelayArgs, pargs types.PluginArgs) (types.CCIPExecProvider, error) {
	return staticCCIPExecProvider{}, nil
}

type mockRelayerExt struct {
	RelayerExt
}

func isType[T any](p any) bool {
	_, ok := p.(T)
	return ok
}

func TestRelayerServerAdapter(t *testing.T) {
	r := &mockRelayer{}
	sa := NewServerAdapter(r, mockRelayerExt{})

	testCases := []struct {
		ProviderType string
		Test         func(p any) bool
		Error        string
	}{
		{
			ProviderType: string(types.Median),
			Test:         isType[types.MedianProvider],
		},
		{
			ProviderType: string(types.Functions),
			Test:         isType[types.FunctionsProvider],
		},
		{
			ProviderType: string(types.Mercury),
			Test:         isType[types.MercuryProvider],
		},
		{
			ProviderType: string(types.CCIPCommit),
			Test:         isType[types.CCIPCommitProvider],
		},
		{
			ProviderType: string(types.CCIPExecution),
			Test:         isType[types.CCIPExecProvider],
		},
		{
			ProviderType: "unknown",
			Error:        "provider type not recognized",
		},
		{
			ProviderType: string(types.GenericPlugin),
			Test:         isType[types.PluginProvider],
		},
		{
			ProviderType: string(types.OCR3Capability),
			Test:         isType[types.OCR3CapabilityProvider],
		},
	}

	ctx := tests.Context(t)
	for _, tc := range testCases {
		pp, err := sa.NewPluginProvider(
			ctx,
			types.RelayArgs{ProviderType: tc.ProviderType},
			types.PluginArgs{},
		)

		if tc.Error != "" {
			assert.ErrorContains(t, err, tc.Error)
		} else {
			assert.NoError(t, err)
			assert.True(t, tc.Test(pp), tc.ProviderType)
		}
	}
}
