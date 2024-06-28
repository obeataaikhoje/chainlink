package llo

import (
	"fmt"
	"math"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/chainlink-common/pkg/services"
)

// Chose uint32 to represent StreamID and ChannelID for the following reasons:
// - 4 billion is more than enough to cover our conceivable channel/stream requirements
// - It is the most compatible, supported everywhere, can be serialized into JSON and used in Javascript without problems
// - It is the smallest reasonable data type that balances between a large set of possible identifiers and not using too much space
// - If randomly chosen, low chance of off-by-one ids being valid
// - Is not specific to any chain, e.g. [32]byte is not fully supported on starknet etc
// - Avoids any possible encoding/copypasta issues e.g. UUIDs which can convert to [32]byte in multiple different ways
type StreamID = uint32

type LifeCycleStage string

// ReportFormat represents different formats for different targets e.g. EVM,
// Solana, JSON, kalechain etc
type ReportFormat uint32

const (
	_ ReportFormat = 0 // reserved

	ReportFormatEVM      = 1
	ReportFormatJSON     = 2
	ReportFormatSolana   = 3
	ReportFormatCosmos   = 4
	ReportFormatStarknet = 5
	ReportFormatAptos    = 6

	_ ReportFormat = math.MaxUint32 // reserved
)

var ReportFormats = []ReportFormat{
	ReportFormatEVM,
	ReportFormatJSON,
	ReportFormatSolana,
	ReportFormatCosmos,
	ReportFormatStarknet,
	ReportFormatAptos,
}

func (rf ReportFormat) String() string {
	switch rf {
	case ReportFormatEVM:
		return "evm"
	case ReportFormatJSON:
		return "json"
	case ReportFormatSolana:
		return "solana"
	case ReportFormatCosmos:
		return "cosmos"
	case ReportFormatStarknet:
		return "starknet"
	case ReportFormatAptos:
		return "aptos"
	default:
		return fmt.Sprintf("unknown(%d)", rf)
	}
}

func ReportFormatFromString(s string) (ReportFormat, error) {
	switch s {
	case "evm":
		return ReportFormatEVM, nil
	case "json":
		return ReportFormatJSON, nil
	case "solana":
		return ReportFormatSolana, nil
	case "cosmos":
		return ReportFormatCosmos, nil
	case "starknet":
		return ReportFormatStarknet, nil
	case "aptos":
		return ReportFormatAptos, nil
	default:
		return 0, fmt.Errorf("unknown report format: %s", s)
	}
}

type ReportInfo struct {
	LifeCycleStage LifeCycleStage
	ReportFormat   ReportFormat
}

type Transmitter ocr3types.ContractTransmitter[ReportInfo]

// QUESTION: Do we also want to include an (optional) designated verifier
// address, i.e. the only address allowed to verify reports from this channel
// https://smartcontract-it.atlassian.net/browse/MERC-3492
type ChannelDefinition struct {
	ReportFormat ReportFormat
	// Specifies the chain on which this channel can be verified. Currently uses
	// CCIP chain selectors.
	ChainSelector uint64
	// We assume that StreamIDs is always non-empty and that the 0-th stream
	// contains the verification price in LINK and the 1-st stream contains the
	// verification price in the native coin.
	StreamIDs []StreamID
}

type ChannelDefinitions map[ChannelID]ChannelDefinition

type ChannelID = uint32

type ChannelDefinitionCache interface {
	Definitions() ChannelDefinitions
	services.Service
}
