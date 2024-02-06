package internal_test

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/smartcontractkit/chainlink-common/pkg/loop/internal/test"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	. "github.com/smartcontractkit/chainlink-common/pkg/types/interfacetests"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
)

func TestCodecClient(t *testing.T) {
	interfaceTester := test.WrapCodecTesterForLoop(&fakeCodecInterfaceTester{impl: &fakeCodec{}})
	RunCodecInterfaceTests(t, interfaceTester)

	es := &errCodec{}
	esTester := test.WrapCodecTesterForLoop(&fakeCodecInterfaceTester{impl: es})
	esTester.Setup(t)
	esCodec := esTester.GetCodec(t)

	anyObj := &TestStruct{}
	for _, errorType := range errorTypes {
		es.err = errorType
		t.Run("Encode unwraps errors from server "+errorType.Error(), func(t *testing.T) {
			_, err := esCodec.Encode(tests.Context(t), anyObj, "doesnotmatter")
			assert.True(t, errors.Is(err, errorType))
		})

		t.Run("Decode unwraps errors from server "+errorType.Error(), func(t *testing.T) {
			_, err := esCodec.Encode(tests.Context(t), anyObj, "doesnotmatter")
			assert.True(t, errors.Is(err, errorType))
		})

		t.Run("GetMaxEncodingSize unwraps errors from server "+errorType.Error(), func(t *testing.T) {
			_, err := esCodec.GetMaxEncodingSize(tests.Context(t), 1, "anything")
			assert.True(t, errors.Is(err, errorType))
		})

		t.Run("GetMaxDecodingSize unwraps errors from server "+errorType.Error(), func(t *testing.T) {
			_, err := esCodec.GetMaxDecodingSize(tests.Context(t), 1, "anything")
			assert.True(t, errors.Is(err, errorType))
		})
	}

	// make sure that errors come from client directly
	es.err = nil
	t.Run("Encode returns error if type cannot be encoded in the wire format", func(t *testing.T) {
		interfaceTester.Setup(t)
		c := interfaceTester.GetCodec(t)
		_, err := c.Encode(tests.Context(t), &cannotEncode{}, "doesnotmatter")
		assert.True(t, errors.Is(err, types.ErrInvalidType))
	})

	t.Run("Decode returns error if type cannot be decoded in the wire format", func(t *testing.T) {
		interfaceTester.Setup(t)
		c := interfaceTester.GetCodec(t)
		fv := int32(1)
		toDecode, err := c.Encode(tests.Context(t), &TestStruct{Field: &fv}, TestItemType)
		require.NoError(t, err)
		err = c.Decode(tests.Context(t), toDecode, &cannotEncode{}, TestItemType)
		assert.True(t, errors.Is(err, types.ErrInvalidType))
	})

	t.Run("Nil esCodec returns unimplemented", func(t *testing.T) {
		ctx := tests.Context(t)
		nilTester := test.WrapCodecTesterForLoop(&fakeCodecInterfaceTester{impl: nil})
		nilTester.Setup(t)
		nilCodec := nilTester.GetCodec(t)

		item := &TestStruct{}

		_, err := nilCodec.Encode(ctx, item, TestItemType)
		assert.Equal(t, codes.Unimplemented, status.Convert(err).Code())

		err = nilCodec.Decode(ctx, []byte("does not matter"), &item, TestItemType)
		assert.Equal(t, codes.Unimplemented, status.Convert(err).Code())

		_, err = nilCodec.GetMaxEncodingSize(ctx, 1, TestItemType)
		assert.Equal(t, codes.Unimplemented, status.Convert(err).Code())

		_, err = nilCodec.GetMaxDecodingSize(ctx, 1, TestItemType)
		assert.Equal(t, codes.Unimplemented, status.Convert(err).Code())
	})
}

type fakeCodecInterfaceTester struct {
	interfaceTesterBase
	impl types.Codec
}

func (it *fakeCodecInterfaceTester) Setup(_ *testing.T) {}

func (it *fakeCodecInterfaceTester) GetCodec(_ *testing.T) types.Codec {
	return it.impl
}

type fakeCodec struct {
	fakeTypeProvider
	lastItem any
}

func (f *fakeCodec) GetMaxDecodingSize(ctx context.Context, n int, itemType string) (int, error) {
	return f.GetMaxEncodingSize(ctx, n, itemType)
}

func (f *fakeCodec) GetMaxEncodingSize(_ context.Context, _ int, itemType string) (int, error) {
	switch itemType {
	case TestItemType, TestItemSliceType, TestItemArray2Type, TestItemArray1Type:
		return 1, nil
	}
	return 0, types.ErrInvalidType
}

func (it *fakeCodecInterfaceTester) EncodeFields(t *testing.T, request *EncodeRequest) []byte {
	if request.TestOn == TestItemType {
		bytes, err := encoder.Marshal(request.TestStructs[0])
		require.NoError(t, err)
		return bytes
	}

	bytes, err := encoder.Marshal(request.TestStructs)
	require.NoError(t, err)
	return bytes
}

func (it *fakeCodecInterfaceTester) IncludeArrayEncodingSizeEnforcement() bool {
	return false
}

func (f *fakeCodec) Encode(_ context.Context, item any, itemType string) ([]byte, error) {
	f.lastItem = item
	switch itemType {
	case NilType:
		return []byte{}, nil
	case TestItemWithConfigExtra:
		ts := item.(*TestStruct)
		ts.Account = anyAccountBytes
		ts.BigField = big.NewInt(2)
		return encoder.Marshal(ts)
	case TestItemType, TestItemSliceType, TestItemArray2Type, TestItemArray1Type:
		switch i := item.(type) {
		case *TestStruct:
			if i.Field == nil {
				return nil, types.ErrInvalidType
			}
		case *TestStructMissingField:
			return nil, types.ErrInvalidType
		}
		return encoder.Marshal(item)
	}
	return nil, types.ErrInvalidType
}

func (f *fakeCodec) Decode(_ context.Context, _ []byte, into any, itemType string) error {
	if f.lastItem == nil {
		return fmt.Errorf("%w: no item to decode", types.ErrInvalidEncoding)
	}
	switch itemType {
	case TestItemWithConfigExtra:
		decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{Squash: true, Result: into})
		if err != nil {
			return err
		}

		if err = decoder.Decode(f.lastItem); err != nil {
			return err
		}
		extra := into.(*TestStructWithExtraField)
		extra.ExtraField = AnyExtraValue
		return nil
	case TestItemType, TestItemSliceType, TestItemArray2Type, TestItemArray1Type:
		return mapstructure.Decode(f.lastItem, into)
	}
	return types.ErrInvalidType
}

type errCodec struct {
	err error
}

func (e *errCodec) Encode(_ context.Context, _ any, _ string) ([]byte, error) {
	return nil, e.err
}

func (e *errCodec) GetMaxEncodingSize(_ context.Context, _ int, _ string) (int, error) {
	return 0, e.err
}

func (e *errCodec) Decode(_ context.Context, _ []byte, _ any, _ string) error {
	return e.err
}

func (e *errCodec) GetMaxDecodingSize(_ context.Context, _ int, _ string) (int, error) {
	return 0, e.err
}
