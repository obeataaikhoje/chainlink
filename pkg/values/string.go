package values

import (
	"github.com/smartcontractkit/chainlink-common/pkg/values/pb"
)

type String struct {
	Underlying string
}

func NewString(s string) *String {
	return &String{Underlying: s}
}

func (s *String) proto() *pb.Value {
	return pb.NewStringValue(s.Underlying)
}

func (s *String) Unwrap() (any, error) {
	return s.Underlying, nil
}

func (s *String) UnwrapTo(to any) error {
	return unwrapTo(s.Underlying, to)
}

func (s *String) Copy() Value {
	return &String{Underlying: s.Underlying}
}
