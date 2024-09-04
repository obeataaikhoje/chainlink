package core

import (
	"context"

	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

type RelayerSet interface {
	Get(ctx context.Context, relayID types.RelayID) (Relayer, error)

	// List lists the relayers corresponding to `...types.RelayID`
	// returning all relayers if len(...types.RelayID) == 0.
	List(ctx context.Context, relayIDs ...types.RelayID) (map[types.RelayID]Relayer, error)
}

type PluginArgs struct {
	TransmitterID string
	PluginConfig  []byte
}

type RelayArgs struct {
	ContractID         string
	RelayConfig        []byte
	ProviderType       string
	MercuryCredentials *types.MercuryCredentials
}

//go:generate mockery --quiet --name Relayer --output ./mocks/ --case=underscore
type Relayer interface {
	services.Service
	NewPluginProvider(context.Context, RelayArgs, PluginArgs) (types.PluginProvider, error)
	NewContractReader(_ context.Context, contractReaderConfig []byte) (types.ContractReader, error)
	NewChainWriter(_ context.Context, chainWriterConfig []byte) (types.ChainWriter, error)
}
