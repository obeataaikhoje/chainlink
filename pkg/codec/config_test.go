package codec_test

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/codec"
)

func TestModifiersConfig(t *testing.T) {
	type testStruct struct {
		A int
		C int
		T int64
	}

	jsonConfig := `[
    {
        "type": "Extract Element",
        "Extractions": {
            "A": "First"
        }
    },
    {
        "Type": "Rename",
        "Fields": {
            "A": "Z"
        }
    },
    {
        "Type": "Drop",
        "Fields": ["C"]
    },
    {
        "Type": "Hard Code",
        "OffChainValues": {
            "B": 2
        }
    },
	{
		"Type": "Epoch To time",
		"Fields": ["T"]
	}
]`

	lowerJSONConfig := `[
    {
        "type": "extract element",
        "extractions": {
            "a": "first"
        }
    },
    {
        "type": "rename",
        "fields": {
            "a": "z"
        }
    },
    {
        "type": "drop",
        "fields": ["c"]
    },
    {
        "type": "hard code",
        "offChainValues": {
            "b": 2
        }
    },
	{
		"type": "epoch to time",
		"fields": ["t"]
	}
]`

	for _, test := range []struct{ name, json string }{
		{"exact", jsonConfig},
		// Used to allow config to match on-chain names/convention
		{"lowercase", lowerJSONConfig},
	} {
		t.Run(test.name, func(t *testing.T) {
			conf := &codec.ModifiersConfig{}
			err := conf.UnmarshalJSON([]byte(test.json))
			require.NoError(t, err)
			modifier, err := conf.ToModifier()
			require.NoError(t, err)

			_, err = modifier.RetypeForOffChain(reflect.TypeOf(testStruct{}), "")
			require.NoError(t, err)

			onChain := testStruct{
				A: 1,
				C: 100,
				T: 631515600,
			}

			offChain, err := modifier.TransformForOffChain(onChain, "")
			require.NoError(t, err)

			b, err := json.Marshal(offChain)
			require.NoError(t, err)
			actualMap := map[string]any{}
			err = json.Unmarshal(b, &actualMap)
			require.NoError(t, err)

			// when decoding to actualMap, the types are lost
			// the tests for the actual modifiers verify the types are correct
			// json is also encoded differently depending on the timezone
			j, err := json.Marshal(time.Unix(onChain.T, 0))
			require.NoError(t, err)
			expectedMap := map[string]any{
				"Z": []any{float64(1)},
				"B": float64(2),
				// drop the quotes around the string
				"T": string(j)[1 : len(j)-1],
			}

			assert.Equal(t, expectedMap, actualMap)
		})
	}
}
