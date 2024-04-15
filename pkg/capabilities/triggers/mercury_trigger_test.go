package triggers

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/mercury"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/smartcontractkit/chainlink-common/pkg/values"
)

var (
	feedOne   = "0x1111111111111111111100000000000000000000000000000000000000000000"
	feedTwo   = "0x2222222222222222222200000000000000000000000000000000000000000000"
	feedThree = "0x3333333333333333333300000000000000000000000000000000000000000000"
	feedFour  = "0x4444444444444444444400000000000000000000000000000000000000000000"
	feedFive  = "0x5555555555555555555500000000000000000000000000000000000000000000"
)

func TestMercuryTrigger(t *testing.T) {
	ts := NewMercuryTriggerService(logger.Test(t), 10)
	ctx := tests.Context(t)
	require.NotNil(t, ts)

	cm := map[string]interface{}{
		"feedIds": []string{feedOne},
	}

	configWrapped, err := values.NewMap(cm)
	require.NoError(t, err)

	im := map[string]interface{}{
		"triggerId": "test-id-1",
	}

	inputsWrapped, err := values.NewMap(im)
	require.NoError(t, err)

	cr := capabilities.TriggerRequest{
		ID:   "test-trigger-id-1",
		Type: "mercury",
		Metadata: capabilities.RequestMetadata{
			WorkflowID: "workflow-id-1",
		},
		Config: configWrapped,
		Inputs: inputsWrapped,
	}

	callback, err := ts.RegisterTrigger(ctx, cr)
	require.NoError(t, err)

	// Send events to trigger and check for them in the callback
	fr := []FeedReport{
		{
			FeedID:               mercury.FeedID(feedOne).Bytes(),
			FullReport:           []byte("0x1234"),
			BenchmarkPrice:       2,
			ObservationTimestamp: 3,
		},
	}

	mfr := []mercury.FeedReport{
		{
			FeedID:               feedOne,
			FullReport:           []byte("0x1234"),
			BenchmarkPrice:       2,
			ObservationTimestamp: 3,
		},
	}
	err = ts.ProcessReport(ctx, fr)
	assert.NoError(t, err)
	assert.Len(t, callback, 1)
	msg := <-callback
	triggerEvent, reports := upwrapTriggerEvent(t, msg)
	//	assert.Equal(t, "mercury", triggerEvent.TriggerType)
	assert.Equal(t, GenerateTriggerEventID(mfr), triggerEvent.ID)
	assert.Len(t, reports, 1)
	assert.Equal(t, mfr[0], reports[0])

	// Unregister the trigger and check that events no longer go on the callback
	require.NoError(t, ts.UnregisterTrigger(ctx, cr))
	err = ts.ProcessReport(ctx, fr)
	assert.NoError(t, err)
	assert.Len(t, callback, 0)
}

func TestMultipleMercuryTriggers(t *testing.T) {
	ts := NewMercuryTriggerService(logger.Test(t), 10)
	ctx := tests.Context(t)
	require.NotNil(t, ts)

	im1 := map[string]interface{}{
		"triggerId": "test-id-1",
	}

	iwrapped1, err := values.NewMap(im1)
	require.NoError(t, err)

	cm1 := map[string]interface{}{
		"feedIds": []string{
			feedOne,
			feedThree,
			feedFour,
		},
	}

	cwrapped1, err := values.NewMap(cm1)
	require.NoError(t, err)

	cr1 := capabilities.TriggerRequest{
		ID:   "cr1",
		Type: "mercury",
		Metadata: capabilities.RequestMetadata{
			WorkflowID: "workflow-id-1",
		},
		Inputs: iwrapped1,
		Config: cwrapped1,
	}

	im2 := map[string]interface{}{
		"triggerId": "test-id-2",
	}

	iwrapped2, err := values.NewMap(im2)
	require.NoError(t, err)

	cm2 := map[string]interface{}{
		"feedIds": []string{
			feedTwo,
			feedThree,
			feedFive,
		},
	}

	cwrapped2, err := values.NewMap(cm2)
	require.NoError(t, err)

	cr2 := capabilities.TriggerRequest{
		ID:   "cr2",
		Type: "mercury",
		Metadata: capabilities.RequestMetadata{
			WorkflowID: "workflow-id-1",
		},
		Inputs: iwrapped2,
		Config: cwrapped2,
	}

	callback1, err := ts.RegisterTrigger(ctx, cr1)
	require.NoError(t, err)

	callback2, err := ts.RegisterTrigger(ctx, cr2)
	require.NoError(t, err)

	// Send events to trigger and check for them in the callback
	fr1 := []FeedReport{
		{
			FeedID:               mercury.FeedID(feedOne).Bytes(),
			FullReport:           []byte("0x1234"),
			BenchmarkPrice:       20,
			ObservationTimestamp: 5,
		},
		{
			FeedID:               mercury.FeedID(feedThree).Bytes(),
			FullReport:           []byte("0x1234"),
			BenchmarkPrice:       25,
			ObservationTimestamp: 8,
		},
		{
			FeedID:               mercury.FeedID(feedTwo).Bytes(),
			FullReport:           []byte("0x1234"),
			BenchmarkPrice:       30,
			ObservationTimestamp: 10,
		},
		{
			FeedID:               mercury.FeedID(feedFour).Bytes(),
			FullReport:           []byte("0x1234"),
			BenchmarkPrice:       40,
			ObservationTimestamp: 15,
		},
	}

	mfr1 := []mercury.FeedReport{
		{
			FeedID:               feedOne,
			FullReport:           []byte("0x1234"),
			BenchmarkPrice:       20,
			ObservationTimestamp: 5,
		},
		{
			FeedID:               feedThree,
			FullReport:           []byte("0x1234"),
			BenchmarkPrice:       25,
			ObservationTimestamp: 8,
		},
		{
			FeedID:               feedTwo,
			FullReport:           []byte("0x1234"),
			BenchmarkPrice:       30,
			ObservationTimestamp: 10,
		},
		{
			FeedID:               feedFour,
			FullReport:           []byte("0x1234"),
			BenchmarkPrice:       40,
			ObservationTimestamp: 15,
		},
	}

	err = ts.ProcessReport(ctx, fr1)
	assert.NoError(t, err)
	assert.Len(t, callback1, 1)
	assert.Len(t, callback2, 1)

	msg := <-callback1
	triggerEvent, reports := upwrapTriggerEvent(t, msg)
	//assert.Equal(t, "mercury", triggerEvent.TriggerType)
	payload := make([]mercury.FeedReport, 0)
	payload = append(payload, mfr1[0], mfr1[1], mfr1[3])
	assert.Equal(t, GenerateTriggerEventID(payload), triggerEvent.ID)
	assert.Len(t, reports, 3)
	assert.Equal(t, mfr1[0], reports[0])
	assert.Equal(t, mfr1[1], reports[1])
	assert.Equal(t, mfr1[3], reports[2])

	msg = <-callback2
	triggerEvent, reports = upwrapTriggerEvent(t, msg)
	//assert.Equal(t, "mercury", triggerEvent.TriggerType)
	payload = make([]mercury.FeedReport, 0)
	payload = append(payload, mfr1[1], mfr1[2]) // Because GenerateTriggerEventID sorts the reports by feedID, this works
	assert.Equal(t, GenerateTriggerEventID(payload), triggerEvent.ID)
	assert.Len(t, reports, 2)
	assert.Equal(t, mfr1[1], reports[0])
	assert.Equal(t, mfr1[2], reports[1])

	require.NoError(t, ts.UnregisterTrigger(ctx, cr1))
	fr2 := []FeedReport{
		{
			FeedID:               mercury.FeedID(feedThree).Bytes(),
			FullReport:           []byte("0x1234"),
			BenchmarkPrice:       50,
			ObservationTimestamp: 20,
		},
	}
	mfr2 := []mercury.FeedReport{
		{
			FeedID:               feedThree,
			FullReport:           []byte("0x1234"),
			BenchmarkPrice:       50,
			ObservationTimestamp: 20,
		},
	}
	err = ts.ProcessReport(ctx, fr2)
	assert.NoError(t, err)
	assert.Len(t, callback1, 0)
	assert.Len(t, callback2, 1)

	msg = <-callback2
	triggerEvent, reports = upwrapTriggerEvent(t, msg)
	require.NoError(t, err)
	//assert.Equal(t, "mercury", triggerEvent.TriggerType)
	payload = make([]mercury.FeedReport, 0)
	payload = append(payload, mfr2[0])
	assert.Equal(t, GenerateTriggerEventID(payload), triggerEvent.ID)
	assert.Len(t, reports, 1)
	assert.Equal(t, mfr2[0], reports[0])

	require.NoError(t, ts.UnregisterTrigger(ctx, cr2))
	err = ts.ProcessReport(ctx, fr1)
	assert.NoError(t, err)
	assert.Len(t, callback1, 0)
	assert.Len(t, callback2, 0)
}

func TestSlowConsumer(t *testing.T) {

	ts := NewMercuryTriggerService(logger.Test(t), 1)
	ctx := tests.Context(t)
	require.NotNil(t, ts)

	im1 := map[string]interface{}{
		"triggerId": "test-id-1",
	}

	iwrapped1, err := values.NewMap(im1)
	require.NoError(t, err)

	cm1 := map[string]interface{}{
		"feedIds": []string{
			feedOne,
			feedTwo,
		},
	}

	cwrapped1, err := values.NewMap(cm1)
	require.NoError(t, err)

	cr1 := capabilities.TriggerRequest{
		ID:   "cr1",
		Type: "mercury",
		Metadata: capabilities.RequestMetadata{
			WorkflowID: "workflow-id-1",
		},
		Inputs: iwrapped1,
		Config: cwrapped1,
	}

	im2 := map[string]interface{}{
		"triggerId": "test-id-2",
	}

	iwrapped2, err := values.NewMap(im2)
	require.NoError(t, err)

	cm2 := map[string]interface{}{
		"feedIds": []string{
			feedThree,
			feedFour,
		},
	}

	cwrapped2, err := values.NewMap(cm2)
	require.NoError(t, err)

	cr2 := capabilities.TriggerRequest{
		ID:   "cr2",
		Type: "mercury",
		Metadata: capabilities.RequestMetadata{
			WorkflowID: "workflow-id-1",
		},
		Inputs: iwrapped2,
		Config: cwrapped2,
	}

	callback1, err := ts.RegisterTrigger(ctx, cr1)
	require.NoError(t, err)

	callback2, err := ts.RegisterTrigger(ctx, cr2)
	require.NoError(t, err)

	// Send events to trigger and check for them in the callback
	fr1 := []FeedReport{
		{
			FeedID:               mercury.FeedID(feedOne).Bytes(),
			FullReport:           []byte("0x1234"),
			BenchmarkPrice:       20,
			ObservationTimestamp: 5,
		},
		{
			FeedID:               mercury.FeedID(feedTwo).Bytes(),
			FullReport:           []byte("0x1234"),
			BenchmarkPrice:       25,
			ObservationTimestamp: 8,
		},
		{
			FeedID:               mercury.FeedID(feedThree).Bytes(),
			FullReport:           []byte("0x1234"),
			BenchmarkPrice:       30,
			ObservationTimestamp: 10,
		},
		{
			FeedID:               mercury.FeedID(feedFour).Bytes(),
			FullReport:           []byte("0x1234"),
			BenchmarkPrice:       40,
			ObservationTimestamp: 15,
		},
	}

	err = ts.ProcessReport(ctx, fr1)
	assert.NoError(t, err)
	assert.Len(t, callback1, 1)
	assert.Len(t, callback2, 1)

	// Consume from callback 1 but not from callback 2 to simulate a slow consumer
	<-callback1

	fr2 := []FeedReport{
		{
			FeedID:               mercury.FeedID(feedOne).Bytes(),
			FullReport:           []byte("0x1234"),
			BenchmarkPrice:       30,
			ObservationTimestamp: 15,
		},
		{
			FeedID:               mercury.FeedID(feedTwo).Bytes(),
			FullReport:           []byte("0x1234"),
			BenchmarkPrice:       35,
			ObservationTimestamp: 18,
		},
		{
			FeedID:               mercury.FeedID(feedThree).Bytes(),
			FullReport:           []byte("0x1234"),
			BenchmarkPrice:       40,
			ObservationTimestamp: 20,
		},
		{
			FeedID:               mercury.FeedID(feedFour).Bytes(),
			FullReport:           []byte("0x1234"),
			BenchmarkPrice:       50,
			ObservationTimestamp: 25,
		},
	}

	err = ts.ProcessReport(ctx, fr2)
	assert.NotNil(t, err)
	var processReportError *ProcessReportError
	errors.As(err, &processReportError)
	assert.Equal(t, processReportError.FailedTriggerIDsToReportIDs["workflow-id-1|cr2"], []int{2, 3})
	assert.Len(t, callback1, 1)
	assert.Len(t, callback2, 1)
}

func upwrapTriggerEvent(t *testing.T, event capabilities.TriggerEvent) (capabilities.TriggerEvent, []mercury.FeedReport) {
	require.NotNil(t, event.Payload)
	mercuryReports, err := mercury.Codec{}.Unwrap(event.Payload)
	require.NoError(t, err)
	return event, mercuryReports
}
