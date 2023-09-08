package mercury

import (
	"math/big"
	"github.com/shopspring/decimal"
)

// PriceScalingFactor indicates the multiplier applied to token prices that we expect from data source
// e.g. for a 1e8 multiplier, a LINK/USD value of 7.42 will be represented as 742000000
// The factor is decreased 1e8 -> 1e6 to comnpensate for baseUSDFee being in cents not usd
var PRICE_SCALING_FACTOR = decimal.NewFromInt(1e6)

// FeeScalingFactor indicates the multiplier applied to fees.
// e.g. for a 1e18 multiplier, a LINK fee of 7.42 will be represented as 7.42e18
// This is what will be baked into the report for use on-chain.
var FEE_SCALING_FACTOR = decimal.NewFromInt(1e18)

// CalculateFee outputs a fee in wei according to the formula: baseUSDFeeCents * scaleFactor / tokenPriceInUSD
func CalculateFee(tokenPriceInUSD *big.Int, baseUSDFeeCents uint32) (*big.Int) {
	if tokenPriceInUSD.Cmp(big.NewInt(0)) == 0 || baseUSDFeeCents == 0 {
		// zero fee if token price or base fee is zero
		return big.NewInt(0)
	}

	// scale baseFee in USD
	baseFeeScaled := decimal.NewFromInt32(int32(baseUSDFeeCents)).Mul(PRICE_SCALING_FACTOR)

	tokenPrice := decimal.NewFromBigInt(tokenPriceInUSD, 0)

	// fee denominated in token
	fee := baseFeeScaled.Div(tokenPrice)

	// scale fee to the expected format
	fee = fee.Mul(FEE_SCALING_FACTOR)

	// convert to big.Int
	return fee.BigInt()
}
