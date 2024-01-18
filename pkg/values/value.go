package values

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/smartcontractkit/chainlink-common/pkg/values/pb"
)

type Unwrappable interface {
	Unwrap() (any, error)
}

type Value interface {
	Proto() (*pb.Value, error)

	Unwrappable
}

func Wrap(v any) (Value, error) {
	switch tv := v.(type) {
	case map[string]any:
		return NewMap(tv)
	case string:
		return NewString(tv)
	case bool:
		return NewBool(tv)
	case []byte:
		return NewBytes(tv)
	case []any:
		return NewList(tv)
	case decimal.Decimal:
		return NewDecimal(tv)
	case int64:
		return NewInt64(tv)
	case int:
		return NewInt64(int64(tv))
	}

	return nil, fmt.Errorf("could not wrap into value: %+v", v)
}

func FromProto(val *pb.Value) (Value, error) {
	switch val.Value.(type) {
	case *pb.Value_StringValue:
		return NewString(val.GetStringValue())
	case *pb.Value_BoolValue:
		return NewBool(val.GetBoolValue())
	case *pb.Value_DecimalValue:
		return fromDecimalValueProto(val.GetDecimalValue())
	case *pb.Value_Int64Value:
		return NewInt64(val.GetInt64Value())
	case *pb.Value_BytesValue:
		return fromBytesValueProto(val.GetBytesValue())
	case *pb.Value_ListValue:
		return fromListValueProto(val.GetListValue())
	case *pb.Value_MapValue:
		return fromMapValueProto(val.GetMapValue())
	}

	return nil, fmt.Errorf("unsupported type %T: %+v", val, val)
}

func fromBytesValueProto(bv string) (Value, error) {
	p, err := base64.StdEncoding.DecodeString(bv)
	if err != nil {
		return nil, err
	}
	return NewBytes(p)
}

func fromMapValueProto(mv *pb.Map) (Value, error) {
	nm := map[string]Value{}
	for k, v := range mv.Fields {
		val, err := FromProto(v)
		if err != nil {
			return nil, err
		}

		nm[k] = val
	}
	return &Map{Value: nm}, nil
}

func fromListValueProto(lv *pb.List) (Value, error) {
	nl := []Value{}
	for _, el := range lv.Fields {
		elv, err := FromProto(el)
		if err != nil {
			return nil, err
		}

		nl = append(nl, elv)
	}
	return &List{Value: nl}, nil
}

func fromDecimalValueProto(decStr string) (Value, error) {
	dec := decimal.Decimal{}
	err := json.Unmarshal([]byte(decStr), &dec)
	if err != nil {
		return nil, err
	}
	return NewDecimal(dec)
}
