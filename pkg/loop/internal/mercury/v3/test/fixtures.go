package mercury_v3_test

import (
	"math/big"

	"github.com/smartcontractkit/chainlink-common/pkg/types/mercury"
	mercury_v3_types "github.com/smartcontractkit/chainlink-common/pkg/types/mercury/v3"
	ocr2plus_types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

type Parameters struct {
	// ReportCodec
	Report               ocr2plus_types.Report
	ReportFields         mercury_v3_types.ReportFields
	MaxReportLength      int
	ObservationTimestamp uint32

	// DataSource
	ReportTimestamp ocr2plus_types.ReportTimestamp
	Observation     mercury_v3_types.Observation
}

var Fixtures = Parameters{
	// ReportCodec
	Report: ocr2plus_types.Report([]byte("mercury v3 report")),
	ReportFields: mercury_v3_types.ReportFields{
		ValidFromTimestamp: 0,
		Timestamp:          1,
		NativeFee:          big.NewInt(2),
		LinkFee:            big.NewInt(3),
		ExpiresAt:          4,
		BenchmarkPrice:     big.NewInt(5),
		Ask:                big.NewInt(6),
		Bid:                big.NewInt(7),
	},
	MaxReportLength:      20,
	ObservationTimestamp: 23,

	// DataSource
	ReportTimestamp: ocr2plus_types.ReportTimestamp{
		ConfigDigest: [32]byte([]byte("mercury v3 config digest")),
		Epoch:        0,
		Round:        1,
	},
	Observation: mercury_v3_types.Observation{
		BenchmarkPrice:        mercury.ObsResult[*big.Int]{Val: big.NewInt(50)},
		Ask:                   mercury.ObsResult[*big.Int]{Val: big.NewInt(60)},
		Bid:                   mercury.ObsResult[*big.Int]{Val: big.NewInt(70)},
		MaxFinalizedTimestamp: mercury.ObsResult[int64]{Val: 79},
		LinkPrice:             mercury.ObsResult[*big.Int]{Val: big.NewInt(30)},
		NativePrice:           mercury.ObsResult[*big.Int]{Val: big.NewInt(20)},
	},
}
