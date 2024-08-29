// Code generated by github.com/smartcontractkit/chainlink-common/pkg/capabilities/cli, DO NOT EDIT.

package ocr3

import (
	"encoding/json"
	"fmt"

	"reflect"

	streams "github.com/smartcontractkit/chainlink-common/pkg/capabilities/triggers/streams"
)

// OCR3 consensus exposed as a capability.
type DataFeedsConsensus struct {
	// Config corresponds to the JSON schema field "config".
	Config DataFeedsConsensusConfig `json:"config" yaml:"config" mapstructure:"config"`

	// Inputs corresponds to the JSON schema field "inputs".
	Inputs DataFeedsConsensusInputs `json:"inputs" yaml:"inputs" mapstructure:"inputs"`

	// Outputs corresponds to the JSON schema field "outputs".
	Outputs SignedReport `json:"outputs" yaml:"outputs" mapstructure:"outputs"`
}

type DataFeedsConsensusConfig struct {
	// AggregationConfig corresponds to the JSON schema field "aggregation_config".
	AggregationConfig DataFeedsConsensusConfigAggregationConfig `json:"aggregation_config" yaml:"aggregation_config" mapstructure:"aggregation_config"`

	// AggregationMethod corresponds to the JSON schema field "aggregation_method".
	AggregationMethod DataFeedsConsensusConfigAggregationMethod `json:"aggregation_method" yaml:"aggregation_method" mapstructure:"aggregation_method"`

	// Encoder corresponds to the JSON schema field "encoder".
	Encoder Encoder `json:"encoder" yaml:"encoder" mapstructure:"encoder"`

	// EncoderConfig corresponds to the JSON schema field "encoder_config".
	EncoderConfig EncoderConfig `json:"encoder_config" yaml:"encoder_config" mapstructure:"encoder_config"`

	// ReportId corresponds to the JSON schema field "report_id".
	ReportId ReportId `json:"report_id" yaml:"report_id" mapstructure:"report_id"`
}

type DataFeedsConsensusConfigAggregationConfig struct {
	// Allowed partial staleness as a number between 0 and 1.
	AllowedPartialStaleness string `json:"allowedPartialStaleness" yaml:"allowedPartialStaleness" mapstructure:"allowedPartialStaleness"`

	// Feeds corresponds to the JSON schema field "feeds".
	Feeds DataFeedsConsensusConfigAggregationConfigFeeds `json:"feeds" yaml:"feeds" mapstructure:"feeds"`
}

type DataFeedsConsensusConfigAggregationConfigFeeds map[string]FeedValue

// UnmarshalJSON implements json.Unmarshaler.
func (j *DataFeedsConsensusConfigAggregationConfig) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["allowedPartialStaleness"]; raw != nil && !ok {
		return fmt.Errorf("field allowedPartialStaleness in DataFeedsConsensusConfigAggregationConfig: required")
	}
	if _, ok := raw["feeds"]; raw != nil && !ok {
		return fmt.Errorf("field feeds in DataFeedsConsensusConfigAggregationConfig: required")
	}
	type Plain DataFeedsConsensusConfigAggregationConfig
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DataFeedsConsensusConfigAggregationConfig(plain)
	return nil
}

type DataFeedsConsensusConfigAggregationMethod string

const DataFeedsConsensusConfigAggregationMethodDataFeeds DataFeedsConsensusConfigAggregationMethod = "data_feeds"

var enumValues_DataFeedsConsensusConfigAggregationMethod = []interface{}{
	"data_feeds",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DataFeedsConsensusConfigAggregationMethod) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_DataFeedsConsensusConfigAggregationMethod {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_DataFeedsConsensusConfigAggregationMethod, v)
	}
	*j = DataFeedsConsensusConfigAggregationMethod(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DataFeedsConsensusConfig) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["aggregation_config"]; raw != nil && !ok {
		return fmt.Errorf("field aggregation_config in DataFeedsConsensusConfig: required")
	}
	if _, ok := raw["aggregation_method"]; raw != nil && !ok {
		return fmt.Errorf("field aggregation_method in DataFeedsConsensusConfig: required")
	}
	if _, ok := raw["encoder"]; raw != nil && !ok {
		return fmt.Errorf("field encoder in DataFeedsConsensusConfig: required")
	}
	if _, ok := raw["encoder_config"]; raw != nil && !ok {
		return fmt.Errorf("field encoder_config in DataFeedsConsensusConfig: required")
	}
	if _, ok := raw["report_id"]; raw != nil && !ok {
		return fmt.Errorf("field report_id in DataFeedsConsensusConfig: required")
	}
	type Plain DataFeedsConsensusConfig
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DataFeedsConsensusConfig(plain)
	return nil
}

type DataFeedsConsensusInputs struct {
	// Observations corresponds to the JSON schema field "observations".
	Observations []streams.Feed `json:"observations" yaml:"observations" mapstructure:"observations"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DataFeedsConsensusInputs) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["observations"]; raw != nil && !ok {
		return fmt.Errorf("field observations in DataFeedsConsensusInputs: required")
	}
	type Plain DataFeedsConsensusInputs
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DataFeedsConsensusInputs(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DataFeedsConsensus) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["config"]; raw != nil && !ok {
		return fmt.Errorf("field config in DataFeedsConsensus: required")
	}
	if _, ok := raw["inputs"]; raw != nil && !ok {
		return fmt.Errorf("field inputs in DataFeedsConsensus: required")
	}
	if _, ok := raw["outputs"]; raw != nil && !ok {
		return fmt.Errorf("field outputs in DataFeedsConsensus: required")
	}
	type Plain DataFeedsConsensus
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DataFeedsConsensus(plain)
	return nil
}

type FeedValue struct {
	// The deviation that is required to generate a new report. Expressed as a
	// percentage. For example, 0.01 is 1% deviation.
	Deviation string `json:"deviation" yaml:"deviation" mapstructure:"deviation"`

	// The interval in seconds after which a new report is generated, regardless of
	// whether any deviations have occurred. New reports reset the timer.
	Heartbeat uint64 `json:"heartbeat" yaml:"heartbeat" mapstructure:"heartbeat"`

	// An optional remapped ID for the feed.
	RemappedID *string `json:"remappedID,omitempty" yaml:"remappedID,omitempty" mapstructure:"remappedID,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *FeedValue) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["deviation"]; raw != nil && !ok {
		return fmt.Errorf("field deviation in FeedValue: required")
	}
	if _, ok := raw["heartbeat"]; raw != nil && !ok {
		return fmt.Errorf("field heartbeat in FeedValue: required")
	}
	type Plain FeedValue
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if 1 > plain.Heartbeat {
		return fmt.Errorf("field %s: must be >= %v", "heartbeat", 1)
	}
	*j = FeedValue(plain)
	return nil
}
