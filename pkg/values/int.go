package values

import (
	"errors"
	"fmt"
	"math"
	"reflect"

	"github.com/smartcontractkit/chainlink-common/pkg/values/pb"
)

type Int64 struct {
	Underlying int64
}

func NewInt64(i int64) *Int64 {
	return &Int64{Underlying: i}
}

func (i *Int64) proto() *pb.Value {
	return pb.NewInt64Value(i.Underlying)
}

func (i *Int64) Unwrap() (any, error) {
	var u int64
	return u, i.UnwrapTo(&u)
}

func (i *Int64) Copy() Value {
	if i == nil {
		return nil
	}
	return &Int64{Underlying: i.Underlying}
}

func (i *Int64) UnwrapTo(to any) error {
	if i == nil {
		return errors.New("cannot unwrap nil values.Int64")
	}

	if to == nil {
		return fmt.Errorf("cannot unwrap to nil pointer: %+v", to)
	}

	switch tv := to.(type) {
	case *int64:
		*tv = i.Underlying
		return nil
	case *int:
		if i.Underlying > math.MaxInt {
			return fmt.Errorf("cannot unwrap int64 to int: number would overflow %d", i)
		}

		if i.Underlying < math.MinInt {
			return fmt.Errorf("cannot unwrap int64 to int: number would underflow %d", i)
		}

		*tv = int(i.Underlying)
		return nil
	case *uint:
		if i.Underlying > math.MaxInt {
			return fmt.Errorf("cannot unwrap int64 to int: number would overflow %d", i)
		}

		if i.Underlying < 0 {
			return fmt.Errorf("cannot unwrap int64 to uint: number would underflow %d", i)
		}

		*tv = uint(i.Underlying)
		return nil
	case *uint32:
		if i.Underlying > math.MaxInt {
			return fmt.Errorf("cannot unwrap int64 to uint32: number would overflow %d", i)
		}

		if i.Underlying < 0 {
			return fmt.Errorf("cannot unwrap int64 to uint32: number would underflow %d", i)
		}

		*tv = uint32(i.Underlying)
		return nil
	case *uint64:
		if i.Underlying < 0 {
			return fmt.Errorf("cannot unwrap int64 to uint: number would underflow %d", i)
		}

		*tv = uint64(i.Underlying)
		return nil
	case *any:
		*tv = i.Underlying
		return nil
	}

	rv := reflect.ValueOf(to)
	if rv.Kind() == reflect.Ptr {
		switch rv.Elem().Kind() {
		case reflect.Int64:
			return i.UnwrapTo(rv.Convert(reflect.PointerTo(reflect.TypeOf(int64(0)))).Interface())
		case reflect.Int32:
			return i.UnwrapTo(rv.Convert(reflect.PointerTo(reflect.TypeOf(int32(0)))).Interface())
		case reflect.Int:
			return i.UnwrapTo(rv.Convert(reflect.PointerTo(reflect.TypeOf(int(0)))).Interface())
		case reflect.Uint64:
			return i.UnwrapTo(rv.Convert(reflect.PointerTo(reflect.TypeOf(uint64(0)))).Interface())
		case reflect.Uint32:
			return i.UnwrapTo(rv.Convert(reflect.PointerTo(reflect.TypeOf(uint32(0)))).Interface())
		case reflect.Uint:
			return i.UnwrapTo(rv.Convert(reflect.PointerTo(reflect.TypeOf(uint(0)))).Interface())
		}
	}

	return fmt.Errorf("cannot unwrap to type %T", to)
}
