package codec

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"github.com/mitchellh/mapstructure"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

// ModifiersConfig unmarshalls as a list of [ModifierConfig] by using a field called Type
// The values available for Type are case-insensitive and the config they require are below:
// - rename -> [RenameModifierConfig]
// - drop -> [DropModifierConfig]
// - hard code -> [HardCodeModifierConfig]
// - extract element -> [ElementExtractorModifierConfig]
// - epoch to time -> [EpochToTimeModifierConfig]
type ModifiersConfig []ModifierConfig

func (m *ModifiersConfig) UnmarshalJSON(data []byte) error {
	var rawDeserialized []json.RawMessage
	if err := json.Unmarshal(data, &rawDeserialized); err != nil {
		return err
	}

	*m = make([]ModifierConfig, len(rawDeserialized))

	for i, d := range rawDeserialized {
		t := typer{}
		if err := json.Unmarshal(d, &t); err != nil {
			return fmt.Errorf("%w: %w", types.ErrInvalidConfig, err)
		}

		mType := ModifierType(strings.ToLower(t.Type))
		switch mType {
		case ModifierRename:
			(*m)[i] = &RenameModifierConfig{}
		case ModifierDrop:
			(*m)[i] = &DropModifierConfig{}
		case ModifierHardCode:
			(*m)[i] = &HardCodeModifierConfig{}
		case ModifierExtractElement:
			(*m)[i] = &ElementExtractorModifierConfig{}
		case ModifierEpochToTime:
			(*m)[i] = &EpochToTimeModifierConfig{}
		default:
			return fmt.Errorf("%w: unknown modifier type: %s", types.ErrInvalidConfig, mType)
		}

		if err := json.Unmarshal(d, (*m)[i]); err != nil {
			return fmt.Errorf("%w: %w", types.ErrInvalidConfig, err)
		}
	}
	return nil
}

func (m *ModifiersConfig) ToModifier(onChainHooks ...mapstructure.DecodeHookFunc) (Modifier, error) {
	modifier := make(MultiModifier, len(*m))
	for i, c := range *m {
		mod, err := c.ToModifier(onChainHooks...)
		if err != nil {
			return nil, err
		}
		modifier[i] = mod
	}
	return modifier, nil
}

type ModifierType string

const (
	ModifierRename         ModifierType = "rename"
	ModifierDrop           ModifierType = "drop"
	ModifierHardCode       ModifierType = "hard code"
	ModifierExtractElement ModifierType = "extract element"
	ModifierEpochToTime    ModifierType = "epoch to time"
)

type ModifierConfig interface {
	ToModifier(onChainHooks ...mapstructure.DecodeHookFunc) (Modifier, error)
}

// RenameModifierConfig renames all fields in the map from the key to the value
// The casing of the first character is ignored to allow compatibility
// of go convention for public fields and on-chain names.
type RenameModifierConfig struct {
	Fields map[string]string
}

func (r *RenameModifierConfig) ToModifier(_ ...mapstructure.DecodeHookFunc) (Modifier, error) {
	for k, v := range r.Fields {
		delete(r.Fields, k)
		r.Fields[upperFirstCharacter(k)] = upperFirstCharacter(v)
	}
	return NewRenamer(r.Fields), nil
}

func (r *RenameModifierConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(&modifierMarshaller[RenameModifierConfig]{
		Type: ModifierRename,
		T:    r,
	})
}

// DropModifierConfig drops all fields listed.  The casing of the first character is ignored to allow compatibility.
// Note that unused fields are ignored by [types.Codec].
// This is only required if you want to rename a field to an already used name.
// For example, if a struct has fields A and B, and you want to rename A to B,
// then you need to either also rename B or drop it.
type DropModifierConfig struct {
	Fields []string
}

func (d *DropModifierConfig) ToModifier(_ ...mapstructure.DecodeHookFunc) (Modifier, error) {
	fields := map[string]string{}
	for i, f := range d.Fields {
		// using a private variable will make the field not serialize, essentially dropping the field
		fields[upperFirstCharacter(f)] = fmt.Sprintf("dropFieldPrivateName%d", i)
	}

	return NewRenamer(fields), nil
}

func (d *DropModifierConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(&modifierMarshaller[DropModifierConfig]{
		Type: ModifierDrop,
		T:    d,
	})
}

// ElementExtractorModifierConfig is used to extract an element from a slice or array
type ElementExtractorModifierConfig struct {
	// Key is the name of the field to extract from and the value is which element to extract.
	Extractions map[string]*ElementExtractorLocation
}

func (e *ElementExtractorModifierConfig) ToModifier(_ ...mapstructure.DecodeHookFunc) (Modifier, error) {
	mapKeyToUpperFirst(e.Extractions)
	return NewElementExtractor(e.Extractions), nil
}

func (e *ElementExtractorModifierConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(&modifierMarshaller[ElementExtractorModifierConfig]{
		Type: ModifierExtractElement,
		T:    e,
	})
}

// HardCodeModifierConfig is used to hard code values into the map.
// Note that hard-coding values will override other values.
type HardCodeModifierConfig struct {
	OnChainValues  map[string]any
	OffChainValues map[string]any
}

func (h *HardCodeModifierConfig) ToModifier(onChainHooks ...mapstructure.DecodeHookFunc) (Modifier, error) {
	mapKeyToUpperFirst(h.OnChainValues)
	mapKeyToUpperFirst(h.OffChainValues)
	return NewHardCoder(h.OnChainValues, h.OffChainValues, onChainHooks...)
}

func (h *HardCodeModifierConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(&modifierMarshaller[HardCodeModifierConfig]{
		Type: ModifierHardCode,
		T:    h,
	})
}

// EpochToTimeModifierConfig is used to convert epoch seconds as uint64 fields on-chain to time.Time
type EpochToTimeModifierConfig struct {
	Fields []string
}

func (e *EpochToTimeModifierConfig) ToModifier(_ ...mapstructure.DecodeHookFunc) (Modifier, error) {
	for i, f := range e.Fields {
		e.Fields[i] = upperFirstCharacter(f)
	}
	return NewEpochToTimeModifier(e.Fields), nil
}

func (e *EpochToTimeModifierConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(&modifierMarshaller[EpochToTimeModifierConfig]{
		Type: ModifierEpochToTime,
		T:    e,
	})
}

type typer struct {
	Type string
}

func upperFirstCharacter(s string) string {
	parts := strings.Split(s, ".")
	for i, p := range parts {
		r := []rune(p)
		r[0] = unicode.ToUpper(r[0])
		parts[i] = string(r)
	}

	return strings.Join(parts, ".")
}

func mapKeyToUpperFirst[T any](m map[string]T) {
	for k, v := range m {
		delete(m, k)
		m[upperFirstCharacter(k)] = v
	}
}

type modifierMarshaller[T any] struct {
	Type ModifierType
	T    *T
}

func (h *modifierMarshaller[T]) MarshalJSON() ([]byte, error) {
	v := reflect.Indirect(reflect.ValueOf(h.T))
	t := v.Type()

	m := map[string]interface{}{
		"Type": h.Type,
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		m[field.Name] = value
	}

	return json.Marshal(m)
}
