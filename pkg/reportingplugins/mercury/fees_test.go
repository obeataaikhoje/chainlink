package mercury

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func scalePrice(usdPrice float64) *big.Int {
	scaledPrice := new(big.Float).Mul(big.NewFloat(usdPrice), big.NewFloat(1e8))
	scaledPriceInt, _ := scaledPrice.Int(nil)
	return scaledPriceInt
}

func Test_Fees(t *testing.T) {
	var baseUSDFeeCents uint32 = 70
	t.Run("with token price > 1", func(t *testing.T) {
		tokenPriceInUSD := scalePrice(1630)
		fee := CalculateFee(tokenPriceInUSD, baseUSDFeeCents)
		expectedFee := big.NewInt(429447852760736) // 0.000429447852760736 18 decimals
		if fee.Cmp(expectedFee) != 0 {
			t.Errorf("Expected fee to be %v, got %v", expectedFee, fee)
		}
	})

	t.Run("with token price < 1", func(t *testing.T) {
		tokenPriceInUSD := scalePrice(0.4)
		fee := CalculateFee(tokenPriceInUSD, baseUSDFeeCents)
		expectedFee := big.NewInt(1750000000000000000) // 1.75 18 decimals
		if fee.Cmp(expectedFee) != 0 {
			t.Errorf("Expected fee to be %v, got %v", expectedFee, fee)
		}
	})

	t.Run("with token price == 0", func(t *testing.T) {
		tokenPriceInUSD := scalePrice(0)
		fee := CalculateFee(tokenPriceInUSD, baseUSDFeeCents)
		assert.Equal(t, big.NewInt(0), fee)
	})

	t.Run("with base fee == 0", func(t *testing.T) {
		tokenPriceInUSD := scalePrice(123)
		baseUSDFeeCents = 0
		fee := CalculateFee(tokenPriceInUSD, baseUSDFeeCents)
		assert.Equal(t, big.NewInt(0), fee)
	})
}
