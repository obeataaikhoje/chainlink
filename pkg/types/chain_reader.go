package types

import (
	"context"

	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
)

// Errors exposed to product plugins
const (
	ErrInvalidType              = InvalidArgumentError("invalid type")
	ErrInvalidConfig            = InvalidArgumentError("invalid configuration")
	ErrChainReaderConfigMissing = UnimplementedError("ChainReader entry missing from RelayConfig")
	ErrInternal                 = InternalError("internal error")
	ErrNotFound                 = NotFoundError("not found")
)

type ChainReader interface {
	// GetLatestValue gets the latest value....
	// The params argument can be any object which maps a set of generic parameters into chain specific parameters defined in RelayConfig.
	// It must encode as an object via [json.Marshal] and [github.com/fxamacker/cbor/v2.Marshal].
	// Typically, would be either a struct with field names mapping to arguments, or anonymous map such as `map[string]any{"baz": 42, "test": true}}`
	//
	// returnVal must [json.Unmarshal] and and [github.com/fxamacker/cbor/v2.Marshal] as an object.
	//
	// Example use:
	//  type ProductParams struct {
	// 		ID int `json:"id"`
	//  }
	//  type ProductReturn struct {
	// 		Foo string `json:"foo"`
	// 		Bar *big.Int `json:"bar"`
	//  }
	//  func do(ctx context.Context, cr ChainReader) (resp ProductReturn, err error) {
	// 		err = cr.GetLatestValue(ctx, "FooContract", "GetProduct", ProductParams{ID:1}, &resp)
	// 		return
	//  }
	//
	// Note that implementations should ignore extra fields in params that are not expected in the call to allow easier
	// use across chains and contract versions.
	// Similarly, when using a struct for returnVal, fields in the return value that are not on-chain will not be set.
	GetLatestValue(ctx context.Context, contractName string, method string, params, returnVal any) error

	// Bind will override current bindings for the same contract, if one has been set and will return an error if the
	// contract is not known by the ChainReader, or if the Address is invalid
	Bind(ctx context.Context, bindings []BoundContract) error
	// TODO UnBind?, doesn't seem like its needed?

	//TODO ReBindByKey, UnBindByKey, this is needed in some places where addresses change but the contract stays the same
	// ReBindByKey(ctx context.Context, key, address string)
	// UnBindByKey() or UnBindByKey(ctx context.Context, key, address string)

	QueryKey(ctx context.Context, keys string, queryFilter query.Filter, limitAndSort query.LimitAndSort, sequenceDataType any) ([]Sequence, error)
	QueryKeys(ctx context.Context, keys []string, queryFilter query.Filter, limitAndSort query.LimitAndSort, sequenceDataTypes []any) ([][]Sequence, error)
	QueryByKeyValuesComparison(ctx context.Context, keyValuesComparator query.KeyValuesComparator, queryFilter query.Filter, limitAndSort query.LimitAndSort, sequenceDataType any) ([]Sequence, error)
	QueryByKeysValuesComparison(ctx context.Context, keysValuesComparators []query.KeyValuesComparator, queryFilter query.Filter, limitAndSort query.LimitAndSort, sequenceDataType []any) ([][]Sequence, error)
}

// Head TODO this is a general chain agnostic Head, its copied from mercury Chain Reader.
// Mercury Chain Reader should be merged with this Chain Reader, so we won't have duplicate Head definitions.
type Head struct {
	Number    uint64
	Hash      []byte
	Timestamp uint64
}

type Sequence struct {
	// TODO Cursor, this should be a unique sequence identifier that chain reader impl. understands.
	// This way we can retrieve past/future sequences (EVM log events) very granularly, but still hide the chain detail.
	Cursor string
	Head
	Data any
}

type BoundContract struct {
	Address string
	Name    string
	Pending bool
}
