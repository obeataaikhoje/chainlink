package fee

import (
	"math"
	"math/big"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func ApplyMultiplier(feeLimit uint32, multiplier float32) (uint32, error) {
	result := decimal.NewFromBigInt(big.NewInt(0).SetUint64(uint64(feeLimit)), 0).Mul(decimal.NewFromFloat32(multiplier)).IntPart()

	if result > math.MaxUint32 {
		return 0, errors.Errorf("integer overflow when applying multiplier of %f to fee limit of %d", multiplier, feeLimit)
	}
	return uint32(result), nil
}

// Returns the fee in its chain specific unit.
type feeUnitToChainUnit func(fee *big.Int) string
