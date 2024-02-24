package values

import (
	"github.com/mitchellh/mapstructure"

	"github.com/smartcontractkit/chainlink-common/pkg/values/pb"
)

type Map struct {
	Underlying map[string]Value
}

func NewMap(m map[string]any) (*Map, error) {
	mv := map[string]Value{}
	for k, v := range m {
		val, err := Wrap(v)
		if err != nil {
			return nil, err
		}

		mv[k] = val
	}

	return &Map{
		Underlying: mv,
	}, nil
}

func MustNewMap(m map[string]any) *Map {
	value, err := NewMap(m)
	if err != nil {
		panic(err)
	}
	return value
}

func (m *Map) Proto() (*pb.Value, error) {
	pm := map[string]*pb.Value{}
	for k, v := range m.Underlying {
		pv, err := v.Proto()
		if err != nil {
			return nil, err
		}

		pm[k] = pv
	}

	return pb.NewMapValue(pm)
}

func (m *Map) Unwrap() (any, error) {
	nm := map[string]any{}
	for k, v := range m.Underlying {
		uv, err := v.Unwrap()
		if err != nil {
			return nil, err
		}

		nm[k] = uv
	}

	return nm, nil
}

func (m *Map) UnwrapTo(toStruct any) error {
	c := &mapstructure.DecoderConfig{
		Result: toStruct,
	}

	d, err := mapstructure.NewDecoder(c)
	if err != nil {
		return err
	}

	return d.Decode(m.Underlying)
}
