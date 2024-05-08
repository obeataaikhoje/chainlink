package ocr3

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jonboulle/clockwork"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	pbtypes "github.com/smartcontractkit/chainlink-common/pkg/capabilities/consensus/ocr3/types"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types/mocks"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/smartcontractkit/chainlink-common/pkg/values"
)

func TestTransmitter(t *testing.T) {
	ctx := tests.Context(t)
	lggr := logger.Test(t)
	s := newStore()

	weid := uuid.New().String()

	cp := newCapability(
		s,
		clockwork.NewFakeClock(),
		10*time.Second,
		func(config *values.Map) (pbtypes.Encoder, error) { return &encoder{}, nil },
		lggr,
		10,
	)
	cp.Start(ctx)
	defer cp.Close()

	payload, err := values.NewMap(map[string]any{"observations": []string{"something happened"}})
	require.NoError(t, err)
	gotCh, err := cp.Execute(ctx, capabilities.CapabilityRequest{
		Metadata: capabilities.RequestMetadata{
			WorkflowExecutionID: weid,
			WorkflowID:          workflowTestID,
		},
		Inputs: payload,
	})
	require.NoError(t, err)

	r := mocks.NewCapabilitiesRegistry(t)
	r.On("Get", mock.Anything, ocrCapabilityID).Return(cp, nil)

	info := &pbtypes.ReportInfo{
		Id: &pbtypes.Id{
			WorkflowExecutionId: weid,
			WorkflowId:          workflowTestID,
		},
		ShouldReport: true,
	}
	infob, err := proto.Marshal(info)
	require.NoError(t, err)

	sp := values.Proto(values.NewString("hello"))
	spb, err := proto.Marshal(sp)
	require.NoError(t, err)
	rep := ocr3types.ReportWithInfo[[]byte]{
		Info:   infob,
		Report: spb,
	}

	transmitter := NewContractTransmitter(lggr, r, "fromAccountString")

	var sqNr uint64
	sigs := []types.AttributedOnchainSignature{
		{Signature: []byte("a-signature")},
	}
	err = transmitter.Transmit(ctx, types.ConfigDigest{}, sqNr, rep, sigs)
	require.NoError(t, err)

	resp := <-gotCh
	assert.Nil(t, resp.Err)

	unwrapped, err := values.Unwrap(resp.Value)
	um := unwrapped.(map[string]any)
	require.NoError(t, err)
	assert.Equal(t, um["report"].([]byte), spb)
	assert.Len(t, um["signatures"], 1)
}

func TestTransmitter_ShouldReportFalse(t *testing.T) {
	ctx := tests.Context(t)
	lggr := logger.Test(t)
	s := newStore()

	weid := uuid.New().String()

	cp := newCapability(
		s,
		clockwork.NewFakeClock(),
		10*time.Second,
		func(config *values.Map) (pbtypes.Encoder, error) { return &encoder{}, nil },
		lggr,
		10,
	)
	cp.Start(ctx)
	defer cp.Close()

	payload, err := values.NewMap(map[string]any{"observations": []string{"something happened"}})
	require.NoError(t, err)
	gotCh, err := cp.Execute(ctx, capabilities.CapabilityRequest{
		Metadata: capabilities.RequestMetadata{
			WorkflowExecutionID: weid,
			WorkflowID:          workflowTestID,
		},
		Inputs: payload,
	})
	require.NoError(t, err)

	r := mocks.NewCapabilitiesRegistry(t)
	r.On("Get", mock.Anything, ocrCapabilityID).Return(cp, nil)

	info := &pbtypes.ReportInfo{
		Id: &pbtypes.Id{
			WorkflowExecutionId: weid,
			WorkflowId:          workflowTestID,
		},
		ShouldReport: false,
	}
	infob, err := proto.Marshal(info)
	require.NoError(t, err)

	sp := values.Proto(values.NewString("hello"))
	spb, err := proto.Marshal(sp)
	require.NoError(t, err)
	rep := ocr3types.ReportWithInfo[[]byte]{
		Info:   infob,
		Report: spb,
	}

	transmitter := NewContractTransmitter(lggr, r, "fromAccountString")

	var sqNr uint64
	sigs := []types.AttributedOnchainSignature{
		{Signature: []byte("a-signature")},
	}
	err = transmitter.Transmit(ctx, types.ConfigDigest{}, sqNr, rep, sigs)
	require.NoError(t, err)

	resp := <-gotCh
	assert.Nil(t, resp.Err)

	unwrapped, err := values.Unwrap(resp.Value)
	um := unwrapped.(map[string]any)
	require.NoError(t, err)
	assert.Nil(t, um["report"])
	assert.Len(t, um["signatures"], 0)
}
