package types

import (
	"context"

	"github.com/google/uuid"

	"github.com/smartcontractkit/libocr/offchainreporting2/reportingplugin/median"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-relay/pkg/reportingplugins/mercury"
	mercuryv1 "github.com/smartcontractkit/chainlink-relay/pkg/reportingplugins/mercury/v1"
	mercuryv2 "github.com/smartcontractkit/chainlink-relay/pkg/reportingplugins/mercury/v2"
)

type Service interface {
	Name() string
	Start(context.Context) error
	Close() error
	Ready() error
	HealthReport() map[string]error
}

// PluginArgs are the args required to create any OCR2 plugin components.
// Its possible that the plugin config might actually be different
// per relay type, so we pass the config directly through.
type PluginArgs struct {
	TransmitterID string
	PluginConfig  []byte
}

type RelayArgs struct {
	ExternalJobID uuid.UUID
	JobID         int32
	ContractID    string
	New           bool // Whether this is a first time job add.
	RelayConfig   []byte
}

type Relayer interface {
	Service
	NewConfigProvider(rargs RelayArgs) (ConfigProvider, error)
	NewMedianProvider(rargs RelayArgs, pargs PluginArgs) (MedianProvider, error)
	NewMercuryProvider(rargs RelayArgs, pargs PluginArgs) (MercuryProvider, error)
}

// The bootstrap jobs only watch config.
type ConfigProvider interface {
	Service
	OffchainConfigDigester() ocrtypes.OffchainConfigDigester
	ContractConfigTracker() ocrtypes.ContractConfigTracker
}

// Plugin is an alias for PluginProvider, for compatibility.
// Deprecated
type Plugin = PluginProvider

// PluginProvider provides common components for any OCR2 plugin.
// It watches config and is able to transmit.
type PluginProvider interface {
	ConfigProvider
	ContractTransmitter() ocrtypes.ContractTransmitter
}

// MedianProvider provides all components needed for a median OCR2 plugin.
type MedianProvider interface {
	PluginProvider
	ReportCodec() median.ReportCodec
	MedianContract() median.MedianContract
	OnchainConfigCodec() median.OnchainConfigCodec
}

// MercuryProvider provides components needed for a mercury OCR2 plugin.
// Mercury requires config tracking but does not transmit on-chain.
type MercuryProvider interface {
	ConfigProvider
	ReportCodecV1() mercuryv1.ReportCodec
	ReportCodecV2() mercuryv2.ReportCodec
	OnchainConfigCodec() mercury.OnchainConfigCodec
	ContractTransmitter() mercury.Transmitter
}
