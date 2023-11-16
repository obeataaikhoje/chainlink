package internal

import (
	"context"
	jsonv1 "encoding/json"
	"fmt"

	jsonv2 "github.com/go-json-experiment/json"

	"github.com/fxamacker/cbor/v2"
	"google.golang.org/grpc/status"

	"github.com/smartcontractkit/chainlink-relay/pkg/loop/internal/pb"
	"github.com/smartcontractkit/chainlink-relay/pkg/types"
)

var _ types.ChainReader = (*chainReaderClient)(nil)

type chainReaderClient struct {
	*brokerExt
	grpc pb.ChainReaderClient
}

// enum of all known encoding formats for versioned data
const (
	JSONEncodingVersion1 = iota
	JSONEncodingVersion2
	CBOREncodingVersion
)

// Version to be used for encoding ( version used for decoding is determined by data received )
// These are separate constants in case we want to upgrade their data formats independently
const ParamsCurrentEncodingVersion = CBOREncodingVersion
const RetvalCurrentEncodingVersion = CBOREncodingVersion

func encodeVersionedBytes(data any, version int32) (*pb.VersionedBytes, error) {
	var bytes []byte
	var err error

	switch version {
	case JSONEncodingVersion1:
		bytes, err = jsonv1.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", types.InvalidTypeError{}, err)
		}
	case JSONEncodingVersion2:
		bytes, err = jsonv2.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", types.InvalidTypeError{}, err)
		}
	case CBOREncodingVersion:
		enco := cbor.CoreDetEncOptions()
		enco.Time = cbor.TimeRFC3339Nano
		var enc cbor.EncMode
		enc, err = enco.EncMode()
		if err != nil {
			return nil, err
		}
		bytes, err = enc.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", types.InvalidTypeError{}, err)
		}
	default:
		return nil, fmt.Errorf("unsupported encoding version %d for data %v", version, data)
	}

	return &pb.VersionedBytes{Version: uint32(version), Data: bytes}, nil
}

func decodeVersionedBytes(res any, vData *pb.VersionedBytes) error {
	var err error
	switch vData.Version {
	case JSONEncodingVersion1:
		err = jsonv1.Unmarshal(vData.Data, res)
	case JSONEncodingVersion2:
		err = jsonv2.Unmarshal(vData.Data, res)
	case CBOREncodingVersion:
		err = cbor.Unmarshal(vData.Data, res)
	default:
		return fmt.Errorf("unsupported encoding version %d for versionedData %v", vData.Version, vData.Data)
	}

	if err != nil {
		return fmt.Errorf("%w: %w", types.InvalidTypeError{}, err)
	}
	return nil
}

func (c *chainReaderClient) GetLatestValue(ctx context.Context, bc types.BoundContract, method string, params, retVal any) error {
	versionedParams, err := encodeVersionedBytes(params, ParamsCurrentEncodingVersion)
	if err != nil {
		return err
	}

	boundContract := pb.BoundContract{Name: bc.Name, Address: bc.Address, Pending: bc.Pending}

	reply, err := c.grpc.GetLatestValue(ctx, &pb.GetLatestValueRequest{Bc: &boundContract, Method: method, Params: versionedParams})
	if err != nil {
		return unwrapClientError(err)
	}

	return decodeVersionedBytes(retVal, reply.RetVal)
}

var _ pb.ChainReaderServer = (*chainReaderServer)(nil)

type chainReaderServer struct {
	pb.UnimplementedChainReaderServer
	impl types.ChainReader
}

func (c *chainReaderServer) GetLatestValue(ctx context.Context, request *pb.GetLatestValueRequest) (*pb.GetLatestValueReply, error) {
	var bc types.BoundContract
	bc.Name = request.Bc.Name[:]
	bc.Address = request.Bc.Address[:]
	bc.Pending = request.Bc.Pending

	params, err := getEncodedType(request.Method, c.impl, true)
	if err != nil {
		return nil, err
	}

	if err = decodeVersionedBytes(params, request.Params); err != nil {
		return nil, err
	}

	retVal, err := getEncodedType(request.Method, c.impl, false)
	if err != nil {
		return nil, err
	}
	err = c.impl.GetLatestValue(ctx, bc, request.Method, params, retVal)
	if err != nil {
		return nil, err
	}

	encodedRetVal, err := encodeVersionedBytes(retVal, RetvalCurrentEncodingVersion)
	if err != nil {
		return nil, err
	}

	return &pb.GetLatestValueReply{RetVal: encodedRetVal}, nil
}

func getEncodedType(itemType string, possibleTypeProvider any, forEncoding bool) (any, error) {
	if rc, ok := possibleTypeProvider.(types.TypeProvider); ok {
		return rc.CreateType(itemType, forEncoding)
	}

	return &map[string]any{}, nil
}

func unwrapClientError(err error) error {
	if s, ok := status.FromError(err); ok {
		switch s.Message() {
		case types.InvalidEncodingError{}.Error():
			return types.InvalidEncodingError{}
		case types.InvalidTypeError{}.Error():
			return types.InvalidTypeError{}
		case types.FieldNotFoundError{}.Error():
			return types.FieldNotFoundError{}
		case types.WrongNumberOfElements{}.Error():
			return types.WrongNumberOfElements{}
		case types.NotASliceError{}.Error():
			return types.NotASliceError{}
		}
	}
	return err
}
