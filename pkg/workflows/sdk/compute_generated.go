// Code generated by github.com/smartcontractkit/chainlink-common/pkg/workflows/gen, DO NOT EDIT.

package sdk

import (
	"encoding/json"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/values"
)

type Compute1Inputs[T0 any] struct {
	Arg0 CapDefinition[T0]
}

type runtime1Inputs[T0 any] struct {
	Arg0 T0
}

func (input Compute1Inputs[I0]) ToSteps() StepInputs {
	return StepInputs{
		Mapping: map[string]any{
			"Arg0": input.Arg0.Ref(),
		},
	}
}

func Compute1[I0 any, O any](w *WorkflowSpecFactory, ref string, input Compute1Inputs[I0], compute func(Runtime, I0) (O, error)) ComputeOutputCap[O] {
	def := StepDefinition{
		ID:     "custom_compute@1.0.0",
		Ref:    ref,
		Inputs: input.ToSteps(),
		Config: map[string]any{
			"config": "$(ENV.config)",
			"binary": "$(ENV.binary)",
		},
		CapabilityType: capabilities.CapabilityTypeAction,
	}

	capFn := func(runtime Runtime, request capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error) {
		var inputs runtime1Inputs[I0]
		if err := request.Inputs.UnwrapTo(&inputs); err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		// verify against any schema by marshalling and unmarshalling
		ji, err := json.Marshal(inputs)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		// use a temp variable to unmarshal to avoid type loss if the inputs has an any in it
		var tmp runtime1Inputs[I0]
		if err := json.Unmarshal(ji, &tmp); err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		output, err := compute(runtime, inputs.Arg0)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		computeOutput := ComputeOutput[O]{Value: output}
		wrapped, err := values.CreateMapFromStruct(computeOutput)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		return capabilities.CapabilityResponse{Value: wrapped}, nil
	}

	if w.fns == nil {
		w.fns = map[string]func(runtime Runtime, request capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error){}
	}
	w.fns[ref] = capFn
	return &computeOutputCap[O]{(&Step[ComputeOutput[O]]{Definition: def}).AddTo(w)}
}

type Compute2Inputs[T0 any, T1 any] struct {
	Arg0 CapDefinition[T0]
	Arg1 CapDefinition[T1]
}

type runtime2Inputs[T0 any, T1 any] struct {
	Arg0 T0
	Arg1 T1
}

func (input Compute2Inputs[I0, I1]) ToSteps() StepInputs {
	return StepInputs{
		Mapping: map[string]any{
			"Arg0": input.Arg0.Ref(),
			"Arg1": input.Arg1.Ref(),
		},
	}
}

func Compute2[I0 any, I1 any, O any](w *WorkflowSpecFactory, ref string, input Compute2Inputs[I0, I1], compute func(Runtime, I0, I1) (O, error)) ComputeOutputCap[O] {
	def := StepDefinition{
		ID:     "custom_compute@1.0.0",
		Ref:    ref,
		Inputs: input.ToSteps(),
		Config: map[string]any{
			"config": "$(ENV.config)",
			"binary": "$(ENV.binary)",
		},
		CapabilityType: capabilities.CapabilityTypeAction,
	}

	capFn := func(runtime Runtime, request capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error) {
		var inputs runtime2Inputs[I0, I1]
		if err := request.Inputs.UnwrapTo(&inputs); err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		// verify against any schema by marshalling and unmarshalling
		ji, err := json.Marshal(inputs)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		// use a temp variable to unmarshal to avoid type loss if the inputs has an any in it
		var tmp runtime2Inputs[I0, I1]
		if err := json.Unmarshal(ji, &tmp); err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		output, err := compute(runtime, inputs.Arg0, inputs.Arg1)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		computeOutput := ComputeOutput[O]{Value: output}
		wrapped, err := values.CreateMapFromStruct(computeOutput)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		return capabilities.CapabilityResponse{Value: wrapped}, nil
	}

	if w.fns == nil {
		w.fns = map[string]func(runtime Runtime, request capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error){}
	}
	w.fns[ref] = capFn
	return &computeOutputCap[O]{(&Step[ComputeOutput[O]]{Definition: def}).AddTo(w)}
}

type Compute3Inputs[T0 any, T1 any, T2 any] struct {
	Arg0 CapDefinition[T0]
	Arg1 CapDefinition[T1]
	Arg2 CapDefinition[T2]
}

type runtime3Inputs[T0 any, T1 any, T2 any] struct {
	Arg0 T0
	Arg1 T1
	Arg2 T2
}

func (input Compute3Inputs[I0, I1, I2]) ToSteps() StepInputs {
	return StepInputs{
		Mapping: map[string]any{
			"Arg0": input.Arg0.Ref(),
			"Arg1": input.Arg1.Ref(),
			"Arg2": input.Arg2.Ref(),
		},
	}
}

func Compute3[I0 any, I1 any, I2 any, O any](w *WorkflowSpecFactory, ref string, input Compute3Inputs[I0, I1, I2], compute func(Runtime, I0, I1, I2) (O, error)) ComputeOutputCap[O] {
	def := StepDefinition{
		ID:     "custom_compute@1.0.0",
		Ref:    ref,
		Inputs: input.ToSteps(),
		Config: map[string]any{
			"config": "$(ENV.config)",
			"binary": "$(ENV.binary)",
		},
		CapabilityType: capabilities.CapabilityTypeAction,
	}

	capFn := func(runtime Runtime, request capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error) {
		var inputs runtime3Inputs[I0, I1, I2]
		if err := request.Inputs.UnwrapTo(&inputs); err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		// verify against any schema by marshalling and unmarshalling
		ji, err := json.Marshal(inputs)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		// use a temp variable to unmarshal to avoid type loss if the inputs has an any in it
		var tmp runtime3Inputs[I0, I1, I2]
		if err := json.Unmarshal(ji, &tmp); err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		output, err := compute(runtime, inputs.Arg0, inputs.Arg1, inputs.Arg2)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		computeOutput := ComputeOutput[O]{Value: output}
		wrapped, err := values.CreateMapFromStruct(computeOutput)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		return capabilities.CapabilityResponse{Value: wrapped}, nil
	}

	if w.fns == nil {
		w.fns = map[string]func(runtime Runtime, request capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error){}
	}
	w.fns[ref] = capFn
	return &computeOutputCap[O]{(&Step[ComputeOutput[O]]{Definition: def}).AddTo(w)}
}

type Compute4Inputs[T0 any, T1 any, T2 any, T3 any] struct {
	Arg0 CapDefinition[T0]
	Arg1 CapDefinition[T1]
	Arg2 CapDefinition[T2]
	Arg3 CapDefinition[T3]
}

type runtime4Inputs[T0 any, T1 any, T2 any, T3 any] struct {
	Arg0 T0
	Arg1 T1
	Arg2 T2
	Arg3 T3
}

func (input Compute4Inputs[I0, I1, I2, I3]) ToSteps() StepInputs {
	return StepInputs{
		Mapping: map[string]any{
			"Arg0": input.Arg0.Ref(),
			"Arg1": input.Arg1.Ref(),
			"Arg2": input.Arg2.Ref(),
			"Arg3": input.Arg3.Ref(),
		},
	}
}

func Compute4[I0 any, I1 any, I2 any, I3 any, O any](w *WorkflowSpecFactory, ref string, input Compute4Inputs[I0, I1, I2, I3], compute func(Runtime, I0, I1, I2, I3) (O, error)) ComputeOutputCap[O] {
	def := StepDefinition{
		ID:     "custom_compute@1.0.0",
		Ref:    ref,
		Inputs: input.ToSteps(),
		Config: map[string]any{
			"config": "$(ENV.config)",
			"binary": "$(ENV.binary)",
		},
		CapabilityType: capabilities.CapabilityTypeAction,
	}

	capFn := func(runtime Runtime, request capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error) {
		var inputs runtime4Inputs[I0, I1, I2, I3]
		if err := request.Inputs.UnwrapTo(&inputs); err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		// verify against any schema by marshalling and unmarshalling
		ji, err := json.Marshal(inputs)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		// use a temp variable to unmarshal to avoid type loss if the inputs has an any in it
		var tmp runtime4Inputs[I0, I1, I2, I3]
		if err := json.Unmarshal(ji, &tmp); err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		output, err := compute(runtime, inputs.Arg0, inputs.Arg1, inputs.Arg2, inputs.Arg3)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		computeOutput := ComputeOutput[O]{Value: output}
		wrapped, err := values.CreateMapFromStruct(computeOutput)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		return capabilities.CapabilityResponse{Value: wrapped}, nil
	}

	if w.fns == nil {
		w.fns = map[string]func(runtime Runtime, request capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error){}
	}
	w.fns[ref] = capFn
	return &computeOutputCap[O]{(&Step[ComputeOutput[O]]{Definition: def}).AddTo(w)}
}

type Compute5Inputs[T0 any, T1 any, T2 any, T3 any, T4 any] struct {
	Arg0 CapDefinition[T0]
	Arg1 CapDefinition[T1]
	Arg2 CapDefinition[T2]
	Arg3 CapDefinition[T3]
	Arg4 CapDefinition[T4]
}

type runtime5Inputs[T0 any, T1 any, T2 any, T3 any, T4 any] struct {
	Arg0 T0
	Arg1 T1
	Arg2 T2
	Arg3 T3
	Arg4 T4
}

func (input Compute5Inputs[I0, I1, I2, I3, I4]) ToSteps() StepInputs {
	return StepInputs{
		Mapping: map[string]any{
			"Arg0": input.Arg0.Ref(),
			"Arg1": input.Arg1.Ref(),
			"Arg2": input.Arg2.Ref(),
			"Arg3": input.Arg3.Ref(),
			"Arg4": input.Arg4.Ref(),
		},
	}
}

func Compute5[I0 any, I1 any, I2 any, I3 any, I4 any, O any](w *WorkflowSpecFactory, ref string, input Compute5Inputs[I0, I1, I2, I3, I4], compute func(Runtime, I0, I1, I2, I3, I4) (O, error)) ComputeOutputCap[O] {
	def := StepDefinition{
		ID:     "custom_compute@1.0.0",
		Ref:    ref,
		Inputs: input.ToSteps(),
		Config: map[string]any{
			"config": "$(ENV.config)",
			"binary": "$(ENV.binary)",
		},
		CapabilityType: capabilities.CapabilityTypeAction,
	}

	capFn := func(runtime Runtime, request capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error) {
		var inputs runtime5Inputs[I0, I1, I2, I3, I4]
		if err := request.Inputs.UnwrapTo(&inputs); err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		// verify against any schema by marshalling and unmarshalling
		ji, err := json.Marshal(inputs)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		// use a temp variable to unmarshal to avoid type loss if the inputs has an any in it
		var tmp runtime5Inputs[I0, I1, I2, I3, I4]
		if err := json.Unmarshal(ji, &tmp); err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		output, err := compute(runtime, inputs.Arg0, inputs.Arg1, inputs.Arg2, inputs.Arg3, inputs.Arg4)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		computeOutput := ComputeOutput[O]{Value: output}
		wrapped, err := values.CreateMapFromStruct(computeOutput)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		return capabilities.CapabilityResponse{Value: wrapped}, nil
	}

	if w.fns == nil {
		w.fns = map[string]func(runtime Runtime, request capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error){}
	}
	w.fns[ref] = capFn
	return &computeOutputCap[O]{(&Step[ComputeOutput[O]]{Definition: def}).AddTo(w)}
}

type Compute6Inputs[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any] struct {
	Arg0 CapDefinition[T0]
	Arg1 CapDefinition[T1]
	Arg2 CapDefinition[T2]
	Arg3 CapDefinition[T3]
	Arg4 CapDefinition[T4]
	Arg5 CapDefinition[T5]
}

type runtime6Inputs[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any] struct {
	Arg0 T0
	Arg1 T1
	Arg2 T2
	Arg3 T3
	Arg4 T4
	Arg5 T5
}

func (input Compute6Inputs[I0, I1, I2, I3, I4, I5]) ToSteps() StepInputs {
	return StepInputs{
		Mapping: map[string]any{
			"Arg0": input.Arg0.Ref(),
			"Arg1": input.Arg1.Ref(),
			"Arg2": input.Arg2.Ref(),
			"Arg3": input.Arg3.Ref(),
			"Arg4": input.Arg4.Ref(),
			"Arg5": input.Arg5.Ref(),
		},
	}
}

func Compute6[I0 any, I1 any, I2 any, I3 any, I4 any, I5 any, O any](w *WorkflowSpecFactory, ref string, input Compute6Inputs[I0, I1, I2, I3, I4, I5], compute func(Runtime, I0, I1, I2, I3, I4, I5) (O, error)) ComputeOutputCap[O] {
	def := StepDefinition{
		ID:     "custom_compute@1.0.0",
		Ref:    ref,
		Inputs: input.ToSteps(),
		Config: map[string]any{
			"config": "$(ENV.config)",
			"binary": "$(ENV.binary)",
		},
		CapabilityType: capabilities.CapabilityTypeAction,
	}

	capFn := func(runtime Runtime, request capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error) {
		var inputs runtime6Inputs[I0, I1, I2, I3, I4, I5]
		if err := request.Inputs.UnwrapTo(&inputs); err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		// verify against any schema by marshalling and unmarshalling
		ji, err := json.Marshal(inputs)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		// use a temp variable to unmarshal to avoid type loss if the inputs has an any in it
		var tmp runtime6Inputs[I0, I1, I2, I3, I4, I5]
		if err := json.Unmarshal(ji, &tmp); err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		output, err := compute(runtime, inputs.Arg0, inputs.Arg1, inputs.Arg2, inputs.Arg3, inputs.Arg4, inputs.Arg5)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		computeOutput := ComputeOutput[O]{Value: output}
		wrapped, err := values.CreateMapFromStruct(computeOutput)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		return capabilities.CapabilityResponse{Value: wrapped}, nil
	}

	if w.fns == nil {
		w.fns = map[string]func(runtime Runtime, request capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error){}
	}
	w.fns[ref] = capFn
	return &computeOutputCap[O]{(&Step[ComputeOutput[O]]{Definition: def}).AddTo(w)}
}

type Compute7Inputs[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any] struct {
	Arg0 CapDefinition[T0]
	Arg1 CapDefinition[T1]
	Arg2 CapDefinition[T2]
	Arg3 CapDefinition[T3]
	Arg4 CapDefinition[T4]
	Arg5 CapDefinition[T5]
	Arg6 CapDefinition[T6]
}

type runtime7Inputs[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any] struct {
	Arg0 T0
	Arg1 T1
	Arg2 T2
	Arg3 T3
	Arg4 T4
	Arg5 T5
	Arg6 T6
}

func (input Compute7Inputs[I0, I1, I2, I3, I4, I5, I6]) ToSteps() StepInputs {
	return StepInputs{
		Mapping: map[string]any{
			"Arg0": input.Arg0.Ref(),
			"Arg1": input.Arg1.Ref(),
			"Arg2": input.Arg2.Ref(),
			"Arg3": input.Arg3.Ref(),
			"Arg4": input.Arg4.Ref(),
			"Arg5": input.Arg5.Ref(),
			"Arg6": input.Arg6.Ref(),
		},
	}
}

func Compute7[I0 any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, O any](w *WorkflowSpecFactory, ref string, input Compute7Inputs[I0, I1, I2, I3, I4, I5, I6], compute func(Runtime, I0, I1, I2, I3, I4, I5, I6) (O, error)) ComputeOutputCap[O] {
	def := StepDefinition{
		ID:     "custom_compute@1.0.0",
		Ref:    ref,
		Inputs: input.ToSteps(),
		Config: map[string]any{
			"config": "$(ENV.config)",
			"binary": "$(ENV.binary)",
		},
		CapabilityType: capabilities.CapabilityTypeAction,
	}

	capFn := func(runtime Runtime, request capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error) {
		var inputs runtime7Inputs[I0, I1, I2, I3, I4, I5, I6]
		if err := request.Inputs.UnwrapTo(&inputs); err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		// verify against any schema by marshalling and unmarshalling
		ji, err := json.Marshal(inputs)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		// use a temp variable to unmarshal to avoid type loss if the inputs has an any in it
		var tmp runtime7Inputs[I0, I1, I2, I3, I4, I5, I6]
		if err := json.Unmarshal(ji, &tmp); err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		output, err := compute(runtime, inputs.Arg0, inputs.Arg1, inputs.Arg2, inputs.Arg3, inputs.Arg4, inputs.Arg5, inputs.Arg6)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		computeOutput := ComputeOutput[O]{Value: output}
		wrapped, err := values.CreateMapFromStruct(computeOutput)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		return capabilities.CapabilityResponse{Value: wrapped}, nil
	}

	if w.fns == nil {
		w.fns = map[string]func(runtime Runtime, request capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error){}
	}
	w.fns[ref] = capFn
	return &computeOutputCap[O]{(&Step[ComputeOutput[O]]{Definition: def}).AddTo(w)}
}

type Compute8Inputs[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any] struct {
	Arg0 CapDefinition[T0]
	Arg1 CapDefinition[T1]
	Arg2 CapDefinition[T2]
	Arg3 CapDefinition[T3]
	Arg4 CapDefinition[T4]
	Arg5 CapDefinition[T5]
	Arg6 CapDefinition[T6]
	Arg7 CapDefinition[T7]
}

type runtime8Inputs[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any] struct {
	Arg0 T0
	Arg1 T1
	Arg2 T2
	Arg3 T3
	Arg4 T4
	Arg5 T5
	Arg6 T6
	Arg7 T7
}

func (input Compute8Inputs[I0, I1, I2, I3, I4, I5, I6, I7]) ToSteps() StepInputs {
	return StepInputs{
		Mapping: map[string]any{
			"Arg0": input.Arg0.Ref(),
			"Arg1": input.Arg1.Ref(),
			"Arg2": input.Arg2.Ref(),
			"Arg3": input.Arg3.Ref(),
			"Arg4": input.Arg4.Ref(),
			"Arg5": input.Arg5.Ref(),
			"Arg6": input.Arg6.Ref(),
			"Arg7": input.Arg7.Ref(),
		},
	}
}

func Compute8[I0 any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, O any](w *WorkflowSpecFactory, ref string, input Compute8Inputs[I0, I1, I2, I3, I4, I5, I6, I7], compute func(Runtime, I0, I1, I2, I3, I4, I5, I6, I7) (O, error)) ComputeOutputCap[O] {
	def := StepDefinition{
		ID:     "custom_compute@1.0.0",
		Ref:    ref,
		Inputs: input.ToSteps(),
		Config: map[string]any{
			"config": "$(ENV.config)",
			"binary": "$(ENV.binary)",
		},
		CapabilityType: capabilities.CapabilityTypeAction,
	}

	capFn := func(runtime Runtime, request capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error) {
		var inputs runtime8Inputs[I0, I1, I2, I3, I4, I5, I6, I7]
		if err := request.Inputs.UnwrapTo(&inputs); err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		// verify against any schema by marshalling and unmarshalling
		ji, err := json.Marshal(inputs)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		// use a temp variable to unmarshal to avoid type loss if the inputs has an any in it
		var tmp runtime8Inputs[I0, I1, I2, I3, I4, I5, I6, I7]
		if err := json.Unmarshal(ji, &tmp); err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		output, err := compute(runtime, inputs.Arg0, inputs.Arg1, inputs.Arg2, inputs.Arg3, inputs.Arg4, inputs.Arg5, inputs.Arg6, inputs.Arg7)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		computeOutput := ComputeOutput[O]{Value: output}
		wrapped, err := values.CreateMapFromStruct(computeOutput)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		return capabilities.CapabilityResponse{Value: wrapped}, nil
	}

	if w.fns == nil {
		w.fns = map[string]func(runtime Runtime, request capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error){}
	}
	w.fns[ref] = capFn
	return &computeOutputCap[O]{(&Step[ComputeOutput[O]]{Definition: def}).AddTo(w)}
}

type Compute9Inputs[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any] struct {
	Arg0 CapDefinition[T0]
	Arg1 CapDefinition[T1]
	Arg2 CapDefinition[T2]
	Arg3 CapDefinition[T3]
	Arg4 CapDefinition[T4]
	Arg5 CapDefinition[T5]
	Arg6 CapDefinition[T6]
	Arg7 CapDefinition[T7]
	Arg8 CapDefinition[T8]
}

type runtime9Inputs[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any] struct {
	Arg0 T0
	Arg1 T1
	Arg2 T2
	Arg3 T3
	Arg4 T4
	Arg5 T5
	Arg6 T6
	Arg7 T7
	Arg8 T8
}

func (input Compute9Inputs[I0, I1, I2, I3, I4, I5, I6, I7, I8]) ToSteps() StepInputs {
	return StepInputs{
		Mapping: map[string]any{
			"Arg0": input.Arg0.Ref(),
			"Arg1": input.Arg1.Ref(),
			"Arg2": input.Arg2.Ref(),
			"Arg3": input.Arg3.Ref(),
			"Arg4": input.Arg4.Ref(),
			"Arg5": input.Arg5.Ref(),
			"Arg6": input.Arg6.Ref(),
			"Arg7": input.Arg7.Ref(),
			"Arg8": input.Arg8.Ref(),
		},
	}
}

func Compute9[I0 any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, O any](w *WorkflowSpecFactory, ref string, input Compute9Inputs[I0, I1, I2, I3, I4, I5, I6, I7, I8], compute func(Runtime, I0, I1, I2, I3, I4, I5, I6, I7, I8) (O, error)) ComputeOutputCap[O] {
	def := StepDefinition{
		ID:     "custom_compute@1.0.0",
		Ref:    ref,
		Inputs: input.ToSteps(),
		Config: map[string]any{
			"config": "$(ENV.config)",
			"binary": "$(ENV.binary)",
		},
		CapabilityType: capabilities.CapabilityTypeAction,
	}

	capFn := func(runtime Runtime, request capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error) {
		var inputs runtime9Inputs[I0, I1, I2, I3, I4, I5, I6, I7, I8]
		if err := request.Inputs.UnwrapTo(&inputs); err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		// verify against any schema by marshalling and unmarshalling
		ji, err := json.Marshal(inputs)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		// use a temp variable to unmarshal to avoid type loss if the inputs has an any in it
		var tmp runtime9Inputs[I0, I1, I2, I3, I4, I5, I6, I7, I8]
		if err := json.Unmarshal(ji, &tmp); err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		output, err := compute(runtime, inputs.Arg0, inputs.Arg1, inputs.Arg2, inputs.Arg3, inputs.Arg4, inputs.Arg5, inputs.Arg6, inputs.Arg7, inputs.Arg8)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		computeOutput := ComputeOutput[O]{Value: output}
		wrapped, err := values.CreateMapFromStruct(computeOutput)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		return capabilities.CapabilityResponse{Value: wrapped}, nil
	}

	if w.fns == nil {
		w.fns = map[string]func(runtime Runtime, request capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error){}
	}
	w.fns[ref] = capFn
	return &computeOutputCap[O]{(&Step[ComputeOutput[O]]{Definition: def}).AddTo(w)}
}

type Compute10Inputs[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any] struct {
	Arg0 CapDefinition[T0]
	Arg1 CapDefinition[T1]
	Arg2 CapDefinition[T2]
	Arg3 CapDefinition[T3]
	Arg4 CapDefinition[T4]
	Arg5 CapDefinition[T5]
	Arg6 CapDefinition[T6]
	Arg7 CapDefinition[T7]
	Arg8 CapDefinition[T8]
	Arg9 CapDefinition[T9]
}

type runtime10Inputs[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any] struct {
	Arg0 T0
	Arg1 T1
	Arg2 T2
	Arg3 T3
	Arg4 T4
	Arg5 T5
	Arg6 T6
	Arg7 T7
	Arg8 T8
	Arg9 T9
}

func (input Compute10Inputs[I0, I1, I2, I3, I4, I5, I6, I7, I8, I9]) ToSteps() StepInputs {
	return StepInputs{
		Mapping: map[string]any{
			"Arg0": input.Arg0.Ref(),
			"Arg1": input.Arg1.Ref(),
			"Arg2": input.Arg2.Ref(),
			"Arg3": input.Arg3.Ref(),
			"Arg4": input.Arg4.Ref(),
			"Arg5": input.Arg5.Ref(),
			"Arg6": input.Arg6.Ref(),
			"Arg7": input.Arg7.Ref(),
			"Arg8": input.Arg8.Ref(),
			"Arg9": input.Arg9.Ref(),
		},
	}
}

func Compute10[I0 any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, O any](w *WorkflowSpecFactory, ref string, input Compute10Inputs[I0, I1, I2, I3, I4, I5, I6, I7, I8, I9], compute func(Runtime, I0, I1, I2, I3, I4, I5, I6, I7, I8, I9) (O, error)) ComputeOutputCap[O] {
	def := StepDefinition{
		ID:     "custom_compute@1.0.0",
		Ref:    ref,
		Inputs: input.ToSteps(),
		Config: map[string]any{
			"config": "$(ENV.config)",
			"binary": "$(ENV.binary)",
		},
		CapabilityType: capabilities.CapabilityTypeAction,
	}

	capFn := func(runtime Runtime, request capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error) {
		var inputs runtime10Inputs[I0, I1, I2, I3, I4, I5, I6, I7, I8, I9]
		if err := request.Inputs.UnwrapTo(&inputs); err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		// verify against any schema by marshalling and unmarshalling
		ji, err := json.Marshal(inputs)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		// use a temp variable to unmarshal to avoid type loss if the inputs has an any in it
		var tmp runtime10Inputs[I0, I1, I2, I3, I4, I5, I6, I7, I8, I9]
		if err := json.Unmarshal(ji, &tmp); err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		output, err := compute(runtime, inputs.Arg0, inputs.Arg1, inputs.Arg2, inputs.Arg3, inputs.Arg4, inputs.Arg5, inputs.Arg6, inputs.Arg7, inputs.Arg8, inputs.Arg9)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		computeOutput := ComputeOutput[O]{Value: output}
		wrapped, err := values.CreateMapFromStruct(computeOutput)
		if err != nil {
			return capabilities.CapabilityResponse{}, err
		}

		return capabilities.CapabilityResponse{Value: wrapped}, nil
	}

	if w.fns == nil {
		w.fns = map[string]func(runtime Runtime, request capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error){}
	}
	w.fns[ref] = capFn
	return &computeOutputCap[O]{(&Step[ComputeOutput[O]]{Definition: def}).AddTo(w)}
}
