// Code generated by github.com/smartcontractkit/chainlink-common/pkg/capabilities/cli, DO NOT EDIT.

package basicaction

import (
	"encoding/json"
	"fmt"
)

// Basic Test Action
type Action struct {
	// Config corresponds to the JSON schema field "config".
	Config ActionConfig `json:"config" yaml:"config" mapstructure:"config"`

	// Inputs corresponds to the JSON schema field "inputs".
	Inputs *ActionInputs `json:"inputs,omitempty" yaml:"inputs,omitempty" mapstructure:"inputs,omitempty"`

	// Outputs corresponds to the JSON schema field "outputs".
	Outputs *ActionOutputs `json:"outputs,omitempty" yaml:"outputs,omitempty" mapstructure:"outputs,omitempty"`
}

type ActionConfig struct {
	// Name corresponds to the JSON schema field "name".
	Name string `json:"name" yaml:"name" mapstructure:"name"`

	// The interval in seconds after which a new trigger event is generated.
	Number int `json:"number" yaml:"number" mapstructure:"number"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ActionConfig) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["name"]; raw != nil && !ok {
		return fmt.Errorf("field name in ActionConfig: required")
	}
	if _, ok := raw["number"]; raw != nil && !ok {
		return fmt.Errorf("field number in ActionConfig: required")
	}
	type Plain ActionConfig
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if 1 > plain.Number {
		return fmt.Errorf("field %s: must be >= %v", "number", 1)
	}
	*j = ActionConfig(plain)
	return nil
}

type ActionInputs struct {
	// InputThing corresponds to the JSON schema field "input_thing".
	InputThing bool `json:"input_thing" yaml:"input_thing" mapstructure:"input_thing"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ActionInputs) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["input_thing"]; raw != nil && !ok {
		return fmt.Errorf("field input_thing in ActionInputs: required")
	}
	type Plain ActionInputs
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ActionInputs(plain)
	return nil
}

type ActionOutputs struct {
	// AdaptedThing corresponds to the JSON schema field "adapted_thing".
	AdaptedThing *string `json:"adapted_thing,omitempty" yaml:"adapted_thing,omitempty" mapstructure:"adapted_thing,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Action) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["config"]; raw != nil && !ok {
		return fmt.Errorf("field config in Action: required")
	}
	type Plain Action
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Action(plain)
	return nil
}
