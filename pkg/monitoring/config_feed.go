package monitoring

import (
	"io"
	"math/big"
)

// FeedParser is the interface for deserializing feed configuration data for each chain integration.
type FeedsParser func(buf io.ReadCloser) ([]FeedConfig, error)

// FeedConfig is the interface for feed configurations extracted from the RDD.
// Implementation can add more fields as needed, but this subset is required by the framework.
type FeedConfig interface {
	// This functions as a feed identifier.
	GetID() string
	GetName() string
	GetPath() string
	GetSymbol() string
	GetHeartbeatSec() int64
	GetContractType() string
	GetContractStatus() string
	GetContractAddress() string
	GetContractAddressBytes() []byte
	// GetMultiply() returns the multiply parameter of a feed.
	// This is a misnomer kept for historical reasons. Multiply is used as divisor
	// for the big integers read from on-chain - think balances, observations,
	// etc. - into prometheus-friendly float64s.
	GetMultiply() *big.Int
	// ToMapping() is useful when encoding kafka messages.
	ToMapping() map[string]interface{}
}
