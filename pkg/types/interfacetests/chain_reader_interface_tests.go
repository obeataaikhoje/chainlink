package interfacetests

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
)

type ChainReaderInterfaceTester interface {
	BasicTester
	GetChainReader(t *testing.T) types.ChainReader

	// SetLatestValue is expected to return the same bound contract and method in the same test
	// Any setup required for this should be done in Setup.
	// The contract should take a LatestParams as the params and return the nth TestStruct set
	SetLatestValue(t *testing.T, testStruct *TestStruct) string
	GetPrimitiveContract(t *testing.T) string
	GetDifferentPrimitiveContract(t *testing.T) string
	GetSliceContract(t *testing.T) string
	GetReturnSeenContract(t *testing.T) string
	TriggerEvent(t *testing.T, testStruct *TestStruct) string
}

const (
	AnyValueToReadWithoutAnArgument             = uint64(3)
	AnyDifferentValueToReadWithoutAnArgument    = uint64(1990)
	MethodTakingLatestParamsReturningTestStruct = "GetLatestValues"
	MethodReturningUint64                       = "GetPrimitiveValue"
	DifferentMethodReturningUint64              = "GetDifferentPrimitiveValue"
	MethodReturningUint64Slice                  = "GetSliceValue"
	MethodReturningSeenStruct                   = "GetSeenStruct"
	EventName                                   = "SomeEvent"
	AnyContractName                             = "TestContract"
	AnySecondContractName                       = "Not" + AnyContractName
)

var AnySliceToReadWithoutAnArgument = []uint64{3, 4}

const AnyExtraValue = 3

func RunChainReaderInterfaceTests(t *testing.T, tester ChainReaderInterfaceTester) {
	tests := []testcase{
		{
			name: "Gets the latest value",
			test: func(t *testing.T) {
				ctx := tests.Context(t)
				firstItem := CreateTestStruct(0, tester)
				address := tester.SetLatestValue(t, &firstItem)
				secondItem := CreateTestStruct(1, tester)
				tester.SetLatestValue(t, &secondItem)

				cr := tester.GetChainReader(t)
				actual := &TestStruct{}
				params := &LatestParams{I: 1}
				bc := types.BoundContract{Name: AnyContractName, Address: address}

				require.NoError(t, cr.GetLatestValue(ctx, bc, MethodTakingLatestParamsReturningTestStruct, params, actual))
				assert.Equal(t, &firstItem, actual)

				params.I = 2
				actual = &TestStruct{}
				require.NoError(t, cr.GetLatestValue(ctx, bc, MethodTakingLatestParamsReturningTestStruct, params, actual))
				assert.Equal(t, &secondItem, actual)
			},
		},
		{
			name: "Get latest value without arguments and with primitive return",
			test: func(t *testing.T) {
				ctx := tests.Context(t)
				address := tester.GetPrimitiveContract(t)
				bc := types.BoundContract{Name: AnyContractName, Address: address}

				cr := tester.GetChainReader(t)

				var prim uint64
				require.NoError(t, cr.GetLatestValue(ctx, bc, MethodReturningUint64, nil, &prim))

				assert.Equal(t, AnyValueToReadWithoutAnArgument, prim)
			},
		},
		{
			name: "Get latest value allows a contract name to resolve different contracts internally",
			test: func(t *testing.T) {
				ctx := tests.Context(t)
				address := tester.GetDifferentPrimitiveContract(t)
				bc := types.BoundContract{Name: AnyContractName, Address: address}

				cr := tester.GetChainReader(t)

				var prim uint64
				require.NoError(t, cr.GetLatestValue(ctx, bc, DifferentMethodReturningUint64, nil, &prim))

				assert.Equal(t, AnyDifferentValueToReadWithoutAnArgument, prim)
			},
		},
		{
			name: "Get latest value allows multiple constract names to have the same function name",
			test: func(t *testing.T) {
				ctx := tests.Context(t)
				address := tester.GetDifferentPrimitiveContract(t)
				bc := types.BoundContract{Name: AnySecondContractName, Address: address}

				cr := tester.GetChainReader(t)

				var prim uint64
				require.NoError(t, cr.GetLatestValue(ctx, bc, MethodReturningUint64, nil, &prim))

				assert.Equal(t, AnyDifferentValueToReadWithoutAnArgument, prim)
			},
		},
		{
			name: "Get latest value without arguments and with slice return",
			test: func(t *testing.T) {
				ctx := tests.Context(t)
				address := tester.GetSliceContract(t)

				cr := tester.GetChainReader(t)

				var slice []uint64
				bc := types.BoundContract{Name: AnyContractName, Address: address}
				require.NoError(t, cr.GetLatestValue(ctx, bc, MethodReturningUint64Slice, nil, &slice))

				assert.Equal(t, AnySliceToReadWithoutAnArgument, slice)
			},
		},
		{
			name: "Get latest value wraps config with modifiers using its own mapstructure overrides",
			test: func(t *testing.T) {
				ctx := tests.Context(t)
				testStruct := CreateTestStruct(0, tester)
				testStruct.BigField = nil
				testStruct.Account = nil

				cr := tester.GetChainReader(t)
				address := tester.GetReturnSeenContract(t)

				actual := &TestStructWithExtraField{}
				bc := types.BoundContract{Name: AnyContractName, Address: address}
				require.NoError(t, cr.GetLatestValue(ctx, bc, MethodReturningSeenStruct, testStruct, actual))

				expected := &TestStructWithExtraField{
					ExtraField: AnyExtraValue,
					TestStruct: CreateTestStruct(0, tester),
				}

				assert.Equal(t, expected, actual)
			},
		},
		{
			name: "Get latest value gets latest event",
			test: func(t *testing.T) {
				ctx := tests.Context(t)
				cr := tester.GetChainReader(t)
				ts := CreateTestStruct(0, tester)
				address := tester.TriggerEvent(t, &ts)
				ts = CreateTestStruct(1, tester)
				tester.TriggerEvent(t, &ts)
				bc := types.BoundContract{Name: AnyContractName, Address: address}

				result := &TestStruct{}
				assert.Eventually(t, func() bool {
					err := cr.GetLatestValue(ctx, bc, EventName, nil, &result)
					return err == nil && reflect.DeepEqual(result, &ts)
				}, time.Second*20, time.Millisecond*10)

				assert.Equal(t, &ts, result)
			},
		},
	}
	runTests(t, tester, tests)
}
