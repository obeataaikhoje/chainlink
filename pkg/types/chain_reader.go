package types

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ChainReaderError string

func (e ChainReaderError) Error() string { return string(e) }

const (
	ErrInvalidType   = ChainReaderError("invalid type")
	ErrInvalidConfig = ChainReaderError("invalid configuration")
)

func UnwrapClientError(err error) error {
	if s, ok := status.FromError(err); ok {
		if s.Code() == codes.Unknown { // Only unwrap custom errors we return, leave alone any other gRPC generated error codes
			return ChainReaderError(s.String())
		}
	}
	return err
}

type ChainReader interface {
	// GetLatestValue gets the latest value....
	// The params argument can be any object which maps a set of generic parameters into chain specific parameters defined in RelayConfig. It must encode as an object via [json.Marshal].
	// Typically, would be either an anonymous map such as `map[string]any{"baz": 42, "test": true}}`, a struct with `json` tags, or something which implements [json.Marshaller].
	//
	// returnVal must [json.Unmarshal] as an object, and so should be a map, struct, or implement the [json.Unmarshaler] interface.
	//
	// Example use:
	//  type ProductParams struct {
	// 		Arg int `json:"arg"`
	//  }
	//  type ProductReturn struct {
	// 		Foo string `json:"foo"`
	// 		Bar *big.Int `json:"bar"`
	//  }
	//  func do(ctx context.Context, cr ChainReader) (resp ProductReturn, err error) {
	// 		err = cr.GetLatestValue(ctx, bc, "method", ProductParams{Arg:1}, &resp)
	// 		return
	//  }
	//
	// returnVal should be a pointer which can be passed to json.Marshal()
	GetLatestValue(ctx context.Context, bc BoundContract, method string, params, returnVal any) error
}

type BoundContract struct {
	Address string
	Name    string
	Pending bool
}

type Event struct {
	ChainID           string
	EventIndexInBlock string
	Address           string
	TxHash            string
	BlockHash         string
	BlockNumber       int64
	BlockTimestamp    time.Time
	CreatedAt         time.Time
	Keys              []string
	Data              []byte
}

type EventFilter struct {
	AddressList []string   // contract address
	KeysList    [][]string // 2D list of indexed search keys, outer dim = AND, inner dim = OR.  Params[0] is the name of the event (or "event type"), rest are any narrowing parameters that may help further specify the event
	Retention   time.Duration
}
