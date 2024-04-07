package ocr3

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/jonboulle/clockwork"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/consensus/ocr3/types"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/smartcontractkit/chainlink-common/pkg/values"
)

const workflowTestID = "consensus-workflow-test-id-1"
const workflowExecutionTestID = "consensus-workflow-execution-test-id-1"

var transformJSON = cmp.FilterValues(func(x, y []byte) bool {
	return json.Valid(x) && json.Valid(y)
}, cmp.Transformer("ParseJSON", func(in []byte) (out interface{}) {
	if err := json.Unmarshal(in, &out); err != nil {
		panic(err) // should never occur given previous filter to ensure valid JSON
	}
	return out
}))

type encoder struct {
	types.Encoder
}

func mockEncoderFactory(_ *values.Map) (types.Encoder, error) {
	return &encoder{}, nil
}

func TestOCR3Capability_GenerateConfigSchema(t *testing.T) {
		n := time.Now()
	fc := clockwork.NewFakeClockAt(n)
	lggr := logger.Test(t)

	s := newStore()
	s.evictedCh = make(chan *request)

	cp := newCapability(s, fc, 1*time.Second, mockEncoderFactory, lggr)
	resp := cp.GetRequestConfigJSONSchema()
	require.NoError(t, resp.Err)
	schemaStr, ok := resp.Value.(*values.String)
	require.True(t, ok)

	var shouldUpdate = false 
	if shouldUpdate {
		err := os.WriteFile("./testdata/fixtures/capability/config_schema.json", []byte(schemaStr.Underlying), 0600)
		require.NoError(t, err)
	}

	fixture, err := os.ReadFile("./testdata/fixtures/capability/config_schema.json")
	require.NoError(t, err)
	
	if diff := cmp.Diff(fixture, []byte(schemaStr.Underlying), transformJSON); diff != "" {
		t.Errorf("TestMercuryTrigger_GenerateConfigSchema() mismatch (-want +got):\n%s", diff)
		t.FailNow()
	}
}

func TestOCR3Capability(t *testing.T) {
	n := time.Now()
	fc := clockwork.NewFakeClockAt(n)
	lggr := logger.Test(t)

	ctx := tests.Context(t)

	s := newStore()
	s.evictedCh = make(chan *request)

	cp := newCapability(s, fc, 1*time.Second, mockEncoderFactory, lggr)
	require.NoError(t, cp.Start(ctx))

	callback := make(chan capabilities.CapabilityResponse, 10)
	config, err := values.NewMap(map[string]any{"aggregation_method": "data_feeds_2_0"})
	require.NoError(t, err)

	ethUsdValue, err := decimal.NewFromString("1.123456")

	require.NoError(t, err)

	obs := map[string]any{"ETH_USD": ethUsdValue}
	inputs, err := values.NewMap(map[string]any{"observations": obs})
	require.NoError(t, err)

	executeReq := capabilities.CapabilityRequest{
		Metadata: capabilities.RequestMetadata{
			WorkflowID:          workflowTestID,
			WorkflowExecutionID: workflowExecutionTestID,
		},
		Config: config,
		Inputs: inputs,
	}
	err = cp.Execute(ctx, callback, executeReq)
	require.NoError(t, err)

	obsv, err := values.NewMap(obs)
	require.NoError(t, err)

	// Mock the oracle returning a response
	err = cp.transmitResponse(ctx, &response{
		CapabilityResponse: capabilities.CapabilityResponse{
			Value: obsv,
		},
		WorkflowExecutionID: workflowExecutionTestID,
	})
	require.NoError(t, err)

	expectedCapabilityResponse := capabilities.CapabilityResponse{
		Value: obsv,
	}

	gotR := <-s.evictedCh
	assert.Equal(t, workflowExecutionTestID, gotR.WorkflowExecutionID)

	assert.Equal(t, expectedCapabilityResponse, <-callback)
}

func TestOCR3Capability_Eviction(t *testing.T) {
	n := time.Now()
	fc := clockwork.NewFakeClockAt(n)
	lggr := logger.Test(t)

	ctx := tests.Context(t)
	rea := time.Second
	s := newStore()
	cp := newCapability(s, fc, rea, mockEncoderFactory, lggr)
	require.NoError(t, cp.Start(ctx))

	config, err := values.NewMap(map[string]any{"aggregation_method": "data_feeds_2_0"})
	require.NoError(t, err)

	ethUsdValue, err := decimal.NewFromString("1.123456")
	require.NoError(t, err)
	inputs, err := values.NewMap(map[string]any{"observations": map[string]any{"ETH_USD": ethUsdValue}})
	require.NoError(t, err)

	rid := uuid.New().String()
	executeReq := capabilities.CapabilityRequest{
		Metadata: capabilities.RequestMetadata{
			WorkflowID:          workflowTestID,
			WorkflowExecutionID: rid,
		},
		Config: config,
		Inputs: inputs,
	}
	callback := make(chan capabilities.CapabilityResponse, 10)
	err = cp.Execute(ctx, callback, executeReq)
	require.NoError(t, err)

	fc.Advance(1 * time.Hour)
	resp := <-callback
	assert.ErrorContains(t, resp.Err, "timeout exceeded: could not process request before expiry")

	_, ok := s.requests[rid]
	assert.False(t, ok)
}

func TestOCR3Capability_Registration(t *testing.T) {
	n := time.Now()
	fc := clockwork.NewFakeClockAt(n)
	lggr := logger.Test(t)

	ctx := tests.Context(t)
	s := newStore()
	cp := newCapability(s, fc, 1*time.Second, mockEncoderFactory, lggr)
	require.NoError(t, cp.Start(ctx))

	config, err := values.NewMap(map[string]any{"aggregation_method": "data_feeds_2_0"})
	require.NoError(t, err)

	registerReq := capabilities.RegisterToWorkflowRequest{
		Metadata: capabilities.RegistrationMetadata{
			WorkflowID: workflowTestID,
		},
		Config: config,
	}

	err = cp.RegisterToWorkflow(ctx, registerReq)
	require.NoError(t, err)

	agg, err := cp.getAggregator(workflowTestID)
	require.NoError(t, err)
	assert.NotNil(t, agg)

	unregisterReq := capabilities.UnregisterFromWorkflowRequest{
		Metadata: capabilities.RegistrationMetadata{
			WorkflowID: workflowTestID,
		},
	}

	err = cp.UnregisterFromWorkflow(ctx, unregisterReq)
	require.NoError(t, err)

	_, err = cp.getAggregator(workflowTestID)
	assert.ErrorContains(t, err, "no aggregator found for")
}
