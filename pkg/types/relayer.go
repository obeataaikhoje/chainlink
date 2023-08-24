package types

import (
	"context"
	"math/big"

	"github.com/google/uuid"
)

// PluginArgs are the args required to create any OCR2 plugin components.
// It's possible that the plugin config might actually be different
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

type ChainStatus struct {
	ID      string
	Enabled bool
	Config  string // TOML
}

type NodeStatus struct {
	ChainID string
	Name    string
	Config  string // TOML
	State   string
}

// Deprecated: use loop.Relayer, which includes context.Context.
type Relayer interface {
	Service
	NewConfigProvider(rargs RelayArgs) (ConfigProvider, error)
	NewMedianProvider(rargs RelayArgs, pargs PluginArgs) (MedianProvider, error)
	NewMercuryProvider(rargs RelayArgs, pargs PluginArgs) (MercuryProvider, error)
	NewFunctionsProvider(rargs RelayArgs, pargs PluginArgs) (FunctionsProvider, error)
}

// Deprecated
type ChainSet[I any, C ChainService] interface {
	Service

	Chain(ctx context.Context, id I) (C, error)

	ChainStatus(ctx context.Context, id string) (ChainStatus, error)
	ChainStatuses(ctx context.Context, offset, limit int) (chains []ChainStatus, count int, err error)

	NodeStatuses(ctx context.Context, offset, limit int, chainIDs ...string) (nodes []NodeStatus, count int, err error)

	SendTx(ctx context.Context, chainID, from, to string, amount *big.Int, balanceCheck bool) error
}

// Deprecated
type ChainService interface {
	Service
	/*
	   GetChainStatus(ctx context.Context) (ChainStatus, error)
	   ListNodeStatuses(ctx context.Context, page_size int32, page_token string) (stats []NodeStatus, next_page_token string, err error)
	   // choose different name than SendTx to avoid collison during refactor.
	   Transact(ctx context.Context, from, to string, amount *big.Int, balanceCheck bool) error
	*/
}
