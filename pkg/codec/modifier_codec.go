package codec

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

func NewModifierCodec(codec types.CodecTypeProvider, modifier Modifier) (types.Codec, error) {
	if codec == nil || modifier == nil {
		return nil, errors.New("inputs must not be nil")
	}

	return &modifierCodec{
		codec:    codec,
		modifier: modifier,
	}, nil
}

var _ types.TypeProvider = &modifierCodec{}

type modifierCodec struct {
	codec    types.CodecTypeProvider
	modifier Modifier
}

func (m *modifierCodec) CreateType(itemType string, forEncoding bool) (any, error) {
	t, err := m.codec.CreateType(itemType, forEncoding)
	if err != nil {
		return nil, err
	}

	ot := reflect.TypeOf(t)
	nt, err := m.modifier.RetypeForOffChain(ot)
	if err != nil {
		return nil, err
	}

	if nt.Kind() == reflect.Pointer {
		return reflect.New(nt.Elem()).Interface(), nil
	}
	return reflect.Zero(nt).Interface(), nil
}

func (m *modifierCodec) Encode(ctx context.Context, item any, itemType string) ([]byte, error) {
	offChainItem, err := m.CreateType(itemType, true)
	if err != nil {
		return nil, err
	}

	rItem := reflect.ValueOf(item)
	rOffChainItem := reflect.ValueOf(offChainItem)

	// If the item is not a pointer, make it one so that it can be modified by convert.
	// Eg: if rOffChainItem is a slice, it'll be nil and can't have elements set.
	if rOffChainItem.Kind() != reflect.Pointer {
		rItem = addr(rItem)
		rOffChainItem = addr(rOffChainItem)
		offChainItem = rOffChainItem.Interface()
	}

	if err = convert(rItem, rOffChainItem); err != nil {
		return nil, err
	}

	onChainItem, err := m.modifier.TransformForOnChain(offChainItem)
	if err != nil {
		return nil, err
	}
	return m.codec.Encode(ctx, onChainItem, itemType)
}

func (m *modifierCodec) GetMaxEncodingSize(ctx context.Context, n int, itemType string) (int, error) {
	return m.codec.GetMaxEncodingSize(ctx, n, itemType)
}

func (m *modifierCodec) Decode(ctx context.Context, raw []byte, into any, itemType string) error {
	rInto := reflect.ValueOf(into)
	if rInto.Kind() != reflect.Pointer {
		return fmt.Errorf("%w: into must be a pointer", types.ErrInvalidType)
	}

	onChain, err := m.codec.CreateType(itemType, false)
	if err != nil {
		return err
	}

	if err = m.codec.Decode(ctx, raw, onChain, itemType); err != nil {
		return err
	}
	offChain, err := m.modifier.TransformForOffChain(onChain)
	if err != nil {
		return err
	}

	return convert(reflect.ValueOf(offChain), rInto)
}

func (m *modifierCodec) GetMaxDecodingSize(ctx context.Context, n int, itemType string) (int, error) {
	return m.codec.GetMaxDecodingSize(ctx, n, itemType)
}

func convert(from, to reflect.Value) error {
	if from.Type() == to.Type() && from.Kind() == reflect.Pointer {
		// Types are the same, just copy the element.
		//  The variable itself may not be addressable
		to.Elem().Set(from.Elem())
		return nil
	}

	switch from.Kind() {
	case reflect.Pointer:

		iFrom := reflect.Indirect(from)
		switch iFrom.Kind() {
		// Pointers can be decoded directly with mapstructure if they are not a pointer to one of these kinds.
		// If they are, use recursion to set the pointer's elements the same.
		case reflect.Array, reflect.Slice, reflect.Pointer:
			return convert(iFrom, reflect.Indirect(to))
		default:
			return mapstructure.Decode(from.Interface(), to.Interface())
		}
	case reflect.Array, reflect.Slice:
		switch to.Kind() {
		// Arrays and slices can't be encoded to a map, so convert each element individually.
		case reflect.Array:
			if from.Len() != to.Len() {
				return types.ErrWrongNumberOfElements
			}
			return convertSliceOrArray(from, to)
		case reflect.Slice:
			// A slice may not be initialized yet, make the right number of elements to copy to
			length := from.Len()
			to.Set(reflect.MakeSlice(to.Type(), length, length))
			return convertSliceOrArray(from, to)
		default:
			return types.ErrInvalidType
		}
	default:
		return mapstructure.Decode(from.Interface(), to.Interface())
	}
}

func convertSliceOrArray(from, to reflect.Value) error {
	switch from.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < from.Len(); i++ {
			if err := convert(addr(from.Index(i)), addr(to.Index(i))); err != nil {
				return err
			}
		}
		return nil
	default:
		return types.ErrInvalidType
	}
}
