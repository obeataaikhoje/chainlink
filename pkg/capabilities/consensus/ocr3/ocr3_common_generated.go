// Code generated by github.com/smartcontractkit/chainlink-common/pkg/capabilities/cli, DO NOT EDIT.

package ocr3

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
)

type Encoder string

type EncoderConfig map[string]interface{}

const EncoderEVM Encoder = "EVM"

var enumValues_Encoder = []interface{}{
	"EVM",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Encoder) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_Encoder {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_Encoder, v)
	}
	*j = Encoder(v)
	return nil
}

type ReportId string

// UnmarshalJSON implements json.Unmarshaler.
func (j *ReportId) UnmarshalJSON(b []byte) error {
	type Plain ReportId
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if matched, _ := regexp.MatchString("^[a-f0-9]{4}$", string(plain)); !matched {
		return fmt.Errorf("field %s pattern match: must match %s", "^[a-f0-9]{4}$", "")
	}
	*j = ReportId(plain)
	return nil
}

type SignedReport struct {
	// Context corresponds to the JSON schema field "Context".
	Context []uint8 `json:"Context" yaml:"Context" mapstructure:"Context"`

	// ID corresponds to the JSON schema field "ID".
	ID []uint8 `json:"ID" yaml:"ID" mapstructure:"ID"`

	// Report corresponds to the JSON schema field "Report".
	Report []uint8 `json:"Report" yaml:"Report" mapstructure:"Report"`

	// Signatures corresponds to the JSON schema field "Signatures".
	Signatures [][]uint8 `json:"Signatures" yaml:"Signatures" mapstructure:"Signatures"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SignedReport) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["Context"]; raw != nil && !ok {
		return fmt.Errorf("field Context in SignedReport: required")
	}
	if _, ok := raw["ID"]; raw != nil && !ok {
		return fmt.Errorf("field ID in SignedReport: required")
	}
	if _, ok := raw["Report"]; raw != nil && !ok {
		return fmt.Errorf("field Report in SignedReport: required")
	}
	if _, ok := raw["Signatures"]; raw != nil && !ok {
		return fmt.Errorf("field Signatures in SignedReport: required")
	}
	type Plain SignedReport
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SignedReport(plain)
	return nil
}
