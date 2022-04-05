package test

import (
	"context"
	"math/big"
	"time"

	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/ocrcommon"
	"github.com/smartcontractkit/libocr/offchainreporting2/reportingplugin/median"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"
)

// constants
var (
	MockEpoch     = uint32(10)
	MockRound     = uint8(1)
	MockAns       = big.NewInt(1234567890)
	MockDigest, _ = types.BytesToConfigDigest([]byte("the placeholder is 32 bytes long"))
	MockTimestamp = time.Unix(2000000000, 0)
	MockOCRLogger = ocrcommon.NewLogger(logger.Default, true, func(string) {})
)

func MockOnchainConfig() ([]byte, error) {
	output := []byte{}
	output = append(output, byte(1))

	max, err := median.ToBytes(big.NewInt(9999999999))
	if err != nil {
		return output, err
	}
	min, err := median.ToBytes(big.NewInt(0))
	if err != nil {
		return output, err
	}
	output = append(output, min[:]...)
	return append(output, max[:]...), nil
}

type MockDataSource struct{}

func (mds *MockDataSource) Observe(ctx context.Context) (*big.Int, error) {
	return MockAns, nil
}

type MockMedianContract struct{}

func (mmc MockMedianContract) LatestTransmissionDetails(
	ctx context.Context,
) (
	configDigest types.ConfigDigest,
	epoch uint32,
	round uint8,
	latestAnswer *big.Int,
	latestTimestamp time.Time,
	err error,
) {
	return MockDigest, MockEpoch, MockRound, MockAns, MockTimestamp, err
}

func (mmc MockMedianContract) LatestRoundRequested(
	ctx context.Context,
	lookback time.Duration,
) (
	configDigest types.ConfigDigest,
	epoch uint32,
	round uint8,
	err error,
) {
	return MockDigest, MockEpoch, MockRound, err
}
