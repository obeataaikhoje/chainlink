package testutils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/cli/cmd/testdata/fixtures/capabilities/basicaction"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/cli/cmd/testdata/fixtures/capabilities/basicaction/basicactiontest"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/cli/cmd/testdata/fixtures/capabilities/basictrigger"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/cli/cmd/testdata/fixtures/capabilities/basictrigger/basictriggertest"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/cli/cmd/testdata/fixtures/capabilities/referenceaction"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/cli/cmd/testdata/fixtures/capabilities/referenceaction/referenceactiontest"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/consensus/ocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/consensus/ocr3/ocr3test"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/targets/chainwriter"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/targets/chainwriter/chainwritertest"
	"github.com/smartcontractkit/chainlink-common/pkg/workflows"
	"github.com/smartcontractkit/chainlink-common/pkg/workflows/testutils"
)

func TestRunner(t *testing.T) {
	t.Parallel()
	t.Run("Run runs a workflow for a single execution", func(t *testing.T) {
		t.Skip("TODO https://smartcontract-it.atlassian.net/browse/KS-442, written earlier for a more full look and feel")
		helper := &testHelper{t: t}
		workflow := createTestWorkflow(helper.transformTrigger)

		runner := testutils.NewRunner()

		triggerMock := basictriggertest.Trigger(runner, func() (basictrigger.TriggerOutputs, error) {
			return basictrigger.TriggerOutputs{CoolOutput: "cool"}, nil
		})

		actionMock := basicactiontest.Action(runner, func(input basicaction.ActionInputs) (basicaction.ActionOutputs, error) {
			assert.True(t, input.InputThing)
			return basicaction.ActionOutputs{AdaptedThing: "it was true"}, nil
		})

		consensusMock := ocr3test.IdenticalConsensus[basicaction.ActionOutputs](runner)

		targetMock := chainwritertest.Target(runner, "chainwriter@1.0.0", func(input chainwriter.TargetInputs) error {
			return nil
		})

		err := runner.Run(workflow)
		require.NoError(t, err)

		trigger := triggerMock.GetStep()
		assert.NotNil(t, trigger)
		assert.NoError(t, trigger.Error)
		assert.Equal(t, "cool", trigger.Output.CoolOutput)

		action := actionMock.GetStep("action")
		assert.NotNil(t, action)
		assert.Equal(t, basicaction.ActionInputs{InputThing: true}, action.Input)
		assert.NoError(t, action.Error)
		assert.Equal(t, "it was true", action.Output.AdaptedThing)

		assert.True(t, helper.transformTriggerCalled)
		consensus := consensusMock.GetStepDecoded("Consensus")
		assert.Equal(t, "it was true", consensus.Output.AdaptedThing)
		require.Len(t, consensus.Input, 1)

		target := targetMock.GetAllWrites()
		require.NoError(t, err)
		assert.Empty(t, target.Errors)
		assert.Len(t, target.Inputs, 1)
		assert.Equal(t, consensus.Output, target.Inputs[0].SignedReport)
	})

	t.Run("Captures errors", func(t *testing.T) {
		t.Skip("TODO https://smartcontract-it.atlassian.net/browse/KS-442")
		assert.Fail(t, "Not implemented")
	})

	t.Run("Fails if MockCapability is not provided for a step that is run", func(t *testing.T) {
		t.Skip("TODO https://smartcontract-it.atlassian.net/browse/KS-442")
		assert.Fail(t, "Not implemented")
	})

	t.Run("Fails build if workflow spec generation fails", func(t *testing.T) {
		t.Skip("TODO https://smartcontract-it.atlassian.net/browse/KS-442")
		assert.Fail(t, "Not implemented")
	})

	t.Run("Fails build if not all leafs are targets", func(t *testing.T) {
		t.Skip("TODO https://smartcontract-it.atlassian.net/browse/KS-442")
		assert.Fail(t, "Not implemented")
	})

	t.Run("GetRegisteredMock returns the mock for a step", func(t *testing.T) {
		runner := testutils.NewRunner()
		expected := basicactiontest.ActionForStep(runner, "action", func(input basicaction.ActionInputs) (basicaction.ActionOutputs, error) {
			return basicaction.ActionOutputs{}, nil
		})
		actual := runner.GetRegisteredMock(expected.ID(), "action")
		assert.Same(t, expected, actual)

		basicactiontest.Action(runner, func(input basicaction.ActionInputs) (basicaction.ActionOutputs, error) {
			return basicaction.ActionOutputs{}, nil
		})
		actual = runner.GetRegisteredMock(expected.ID(), "action")
		assert.Same(t, expected, actual)
	})

	t.Run("GetRegisteredMock returns a default mock if step wasn't specified", func(t *testing.T) {
		runner := testutils.NewRunner()
		expected := basicactiontest.Action(runner, func(input basicaction.ActionInputs) (basicaction.ActionOutputs, error) {
			return basicaction.ActionOutputs{}, nil
		})
		actual := runner.GetRegisteredMock(expected.ID(), "action")
		assert.Same(t, expected, actual)
	})

	t.Run("GetRegisteredMock returns nil if no mock was registered", func(t *testing.T) {
		runner := testutils.NewRunner()
		referenceactiontest.Action(runner, func(input referenceaction.SomeInputs) (referenceaction.SomeOutputs, error) {
			return referenceaction.SomeOutputs{}, nil
		})
		assert.Nil(t, runner.GetRegisteredMock("basic-test-action@1.0.0", "action"))
	})

	t.Run("GetRegisteredMock returns nil if no mock was registered for a step", func(t *testing.T) {
		runner := testutils.NewRunner()
		differentStep := basicactiontest.ActionForStep(runner, "step", func(input basicaction.ActionInputs) (basicaction.ActionOutputs, error) {
			return basicaction.ActionOutputs{}, nil
		})
		actual := runner.GetRegisteredMock(differentStep.ID(), "action")
		assert.Nil(t, actual)
	})
}

type actionTransform func(sdk workflows.SDK, outputs basictrigger.TriggerOutputs) (bool, error)

func createTestWorkflow(actionTransform actionTransform) *workflows.WorkflowSpecFactory {
	workflow := workflows.NewWorkflowSpecFactory(workflows.NewWorkflowParams{Name: "tester", Owner: "ryan"})
	trigger := basictrigger.TriggerConfig{Name: "trigger", Number: 100}.New(workflow)
	tTransform := workflows.Compute1[basictrigger.TriggerOutputs, bool](
		workflow,
		"transform",
		workflows.Compute1Inputs[basictrigger.TriggerOutputs]{Arg0: trigger},
		actionTransform)

	action := basicaction.ActionConfig{CamelCaseInSchemaForTesting: "action", SnakeCaseInSchemaForTesting: 20}.
		New(workflow, "basic action", basicaction.ActionInput{InputThing: tTransform.Value()})

	consensus := ocr3.IdenticalConsensusConfig[basicaction.ActionOutputs]{
		Encoder:       "Test",
		EncoderConfig: ocr3.EncoderConfig{},
	}.New(workflow, "consensus", ocr3.IdenticalConsensusInput[basicaction.ActionOutputs]{Observations: action})

	chainwriter.TargetConfig{
		Address:    "0x123",
		DeltaStage: "2m",
		Schedule:   "oneAtATime",
	}.New(workflow, "chainwriter@1.0.0", chainwriter.TargetInput{SignedReport: consensus})
	return workflow
}

type testHelper struct {
	t                      *testing.T
	transformTriggerCalled bool
}

func (helper *testHelper) transformTrigger(sdk workflows.SDK, outputs basictrigger.TriggerOutputs) (bool, error) {
	assert.NotNil(helper.t, sdk)
	assert.Equal(helper.t, "cool", outputs.CoolOutput)
	assert.False(helper.t, helper.transformTriggerCalled)
	helper.transformTriggerCalled = true
	return true, nil
}
