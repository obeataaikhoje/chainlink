// Code generated by github.com/smartcontractkit/chainlink-common/pkg/capabilities/cli, DO NOT EDIT.

package arrayaction

import (
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/workflows"
)

func (cfg ActionConfig) New(w *workflows.WorkflowSpecFactory, ref string, input ActionInput) workflows.CapDefinition[[]ActionOutputsElem] {

	def := workflows.StepDefinition{
		ID: "array-test-action@1.0.0", Ref: ref,
		Inputs: input.ToSteps(),
		Config: map[string]any{
			"details": cfg.Details,
		},
		CapabilityType: capabilities.CapabilityTypeAction,
	}

	step := workflows.Step[[]ActionOutputsElem]{Definition: def}
	return step.AddTo(w)
}

type ActionOutputsElemCap interface {
	workflows.CapDefinition[ActionOutputsElem]
	Results() ActionOutputsElemResultsCap
	private()
}

// ActionOutputsElemCapFromStep should only be called from generated code to assure type safety
func ActionOutputsElemCapFromStep(w *workflows.WorkflowSpecFactory, step workflows.Step[ActionOutputsElem]) ActionOutputsElemCap {
	raw := step.AddTo(w)
	return &actionOutputsElem{CapDefinition: raw}
}

type actionOutputsElem struct {
	workflows.CapDefinition[ActionOutputsElem]
}

func (*actionOutputsElem) private() {}
func (c *actionOutputsElem) Results() ActionOutputsElemResultsCap {
	return &actionOutputsElemResults{CapDefinition: workflows.AccessField[ActionOutputsElem, ActionOutputsElemResults](c.CapDefinition, "results,omitempty")}
}

func NewActionOutputsElemFromFields(
	results ActionOutputsElemResultsCap) ActionOutputsElemCap {
	return &simpleActionOutputsElem{
		CapDefinition: workflows.ComponentCapDefinition[ActionOutputsElem]{
			"results,omitempty": results.Ref(),
		},
		results: results,
	}
}

type simpleActionOutputsElem struct {
	workflows.CapDefinition[ActionOutputsElem]
	results ActionOutputsElemResultsCap
}

func (c *simpleActionOutputsElem) Results() ActionOutputsElemResultsCap {
	return c.results
}

func (c *simpleActionOutputsElem) private() {}

type ActionOutputsElemResultsCap interface {
	workflows.CapDefinition[ActionOutputsElemResults]
	AdaptedThing() workflows.CapDefinition[string]
	private()
}

// ActionOutputsElemResultsCapFromStep should only be called from generated code to assure type safety
func ActionOutputsElemResultsCapFromStep(w *workflows.WorkflowSpecFactory, step workflows.Step[ActionOutputsElemResults]) ActionOutputsElemResultsCap {
	raw := step.AddTo(w)
	return &actionOutputsElemResults{CapDefinition: raw}
}

type actionOutputsElemResults struct {
	workflows.CapDefinition[ActionOutputsElemResults]
}

func (*actionOutputsElemResults) private() {}
func (c *actionOutputsElemResults) AdaptedThing() workflows.CapDefinition[string] {
	return workflows.AccessField[ActionOutputsElemResults, string](c.CapDefinition, "adapted_thing")
}

func NewActionOutputsElemResultsFromFields(
	adaptedThing workflows.CapDefinition[string]) ActionOutputsElemResultsCap {
	return &simpleActionOutputsElemResults{
		CapDefinition: workflows.ComponentCapDefinition[ActionOutputsElemResults]{
			"adapted_thing": adaptedThing.Ref(),
		},
		adaptedThing: adaptedThing,
	}
}

type simpleActionOutputsElemResults struct {
	workflows.CapDefinition[ActionOutputsElemResults]
	adaptedThing workflows.CapDefinition[string]
}

func (c *simpleActionOutputsElemResults) AdaptedThing() workflows.CapDefinition[string] {
	return c.adaptedThing
}

func (c *simpleActionOutputsElemResults) private() {}

type ActionInput struct {
	Metadata workflows.CapDefinition[ActionInputsMetadata]
}

func (input ActionInput) ToSteps() workflows.StepInputs {
	return workflows.StepInputs{
		Mapping: map[string]any{
			"metadata": input.Metadata.Ref(),
		},
	}
}
