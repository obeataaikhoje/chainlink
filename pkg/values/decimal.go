package values

import (
	"github.com/shopspring/decimal"

	"github.com/smartcontractkit/chainlink-common/pkg/values/pb"
)

type Decimal struct {
	Underlying decimal.Decimal
}

func NewDecimal(d decimal.Decimal) *Decimal {
	return &Decimal{Underlying: d}
}

func (d *Decimal) proto() *pb.Value {
	return pb.NewDecimalValue(d.Underlying)
}

func (d *Decimal) Unwrap() (any, error) {
	return d.Underlying, nil
}

func (d *Decimal) UnwrapTo(to any) error {
	return unwrapTo(d.Underlying, to)
}

func (d *Decimal) Copy() Value {
	return &Decimal{Underlying: d.Underlying.Copy()}
}
