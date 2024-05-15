package triggers

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/mercury"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/pb"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/values"
)

const (
	eventID    = "ev_id_1"
	timestamp  = 1000
	rawReport1 = "abcd"
	rawReport2 = "efgh"
)

type testMercuryCodec struct {
}

func (c testMercuryCodec) Unwrap(raw values.Value) ([]mercury.FeedReport, error) {
	dest := []mercury.FeedReport{}
	err := raw.UnwrapTo(&dest)
	return dest, err
}

func (c testMercuryCodec) Wrap(reports []mercury.FeedReport) (values.Value, error) {
	return values.Wrap(reports)
}

func TestMercuryRemoteAggregator(t *testing.T) {
	agg := NewMercuryRemoteAggregator(testMercuryCodec{}, logger.Nop())
	rs := [][]byte{{1, 2, 3}}
	ss := [][]byte{{4, 5, 6}}
	vs := []byte{7, 8, 9}

	feed1Old := mercury.FeedReport{
		FeedID:               feedOne,
		BenchmarkPrice:       big.NewInt(100).Bytes(),
		ObservationTimestamp: 100,
		FullReport:           []byte(rawReport1),
		Rs:                   rs,
		Ss:                   ss,
		Vs:                   vs,
	}
	feed1New := mercury.FeedReport{
		FeedID:               feedOne,
		BenchmarkPrice:       big.NewInt(200).Bytes(),
		ObservationTimestamp: 200,
		FullReport:           []byte(rawReport1),
		Rs:                   rs,
		Ss:                   ss,
		Vs:                   vs,
	}
	feed2Old := mercury.FeedReport{
		FeedID:               feedTwo,
		BenchmarkPrice:       big.NewInt(300).Bytes(),
		ObservationTimestamp: 300,
		FullReport:           []byte(rawReport2),
		Rs:                   rs,
		Ss:                   ss,
		Vs:                   vs,
	}
	feed2New := mercury.FeedReport{
		FeedID:               feedTwo,
		BenchmarkPrice:       big.NewInt(400).Bytes(),
		ObservationTimestamp: 400,
		FullReport:           []byte(rawReport2),
		Rs:                   rs,
		Ss:                   ss,
		Vs:                   vs,
	}

	node1Resp, err := wrapReports([]mercury.FeedReport{feed1Old, feed2New}, eventID, 400)
	require.NoError(t, err)
	rawNode1Resp, err := pb.MarshalCapabilityResponse(node1Resp)
	require.NoError(t, err)
	node2Resp, err := wrapReports([]mercury.FeedReport{feed1New, feed2Old}, eventID, 300)
	require.NoError(t, err)
	rawNode2Resp, err := pb.MarshalCapabilityResponse(node2Resp)
	require.NoError(t, err)

	// aggregator should return latest value for each feedID
	aggResponse, err := agg.Aggregate(eventID, [][]byte{rawNode1Resp, rawNode2Resp})
	require.NoError(t, err)
	aggEvent := capabilities.TriggerEvent{}
	require.NoError(t, aggResponse.Value.UnwrapTo(&aggEvent))
	decodedReports, err := testMercuryCodec{}.Unwrap(aggEvent.Payload)
	require.NoError(t, err)

	require.Len(t, decodedReports, 2)
	require.Equal(t, feed1New, decodedReports[0])
	require.Equal(t, feed2New, decodedReports[1])
}
