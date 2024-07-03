package interfacetests

import (
	"errors"
	"reflect"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
)

type ChainReaderInterfaceTester[T TestingT[T]] interface {
	BasicTester[T]
	GetChainReader(t T) types.ContractReader

	// SetLatestValue is expected to return the same bound contract and method in the same test
	// Any setup required for this should be done in Setup.
	// The contract should take a LatestParams as the params and return the nth TestStruct set
	SetLatestValue(t T, testStruct *TestStruct)
	SetBatchLatestValues(t T, batchCallEntry BatchCallEntry)
	TriggerEvent(t T, testStruct *TestStruct)
	GetBindings(t T) []types.BoundContract
	MaxWaitTimeForEvents() time.Duration
}

const (
	AnyValueToReadWithoutAnArgument             = uint64(3)
	AnyDifferentValueToReadWithoutAnArgument    = uint64(1990)
	MethodTakingLatestParamsReturningTestStruct = "BatchGetLatestValues"
	MethodReturningUint64                       = "GetPrimitiveValue"
	MethodReturningUint64Slice                  = "GetSliceValue"
	MethodReturningSeenStruct                   = "GetSeenStruct"
	EventName                                   = "SomeEvent"
	EventWithFilterName                         = "SomeEventToFilter"
	AnyContractName                             = "TestContract"
	AnySecondContractName                       = "Not" + AnyContractName
)

var AnySliceToReadWithoutAnArgument = []uint64{3, 4}

const AnyExtraValue = 3

func RunChainReaderInterfaceTests[T TestingT[T]](t T, tester ChainReaderInterfaceTester[T]) {
	t.Run("GetLatestValue for "+tester.Name(), func(t T) { runChainReaderGetLatestValueInterfaceTests(t, tester) })
	t.Run("BatchGetLatestValues for "+tester.Name(), func(t T) { runChainReaderBatchGetLatestValuesInterfaceTests(t, tester) })
	t.Run("QueryKey for "+tester.Name(), func(t T) { runQueryKeyInterfaceTests(t, tester) })
}

func runChainReaderGetLatestValueInterfaceTests[T TestingT[T]](t T, tester ChainReaderInterfaceTester[T]) {
	tests := []testcase[T]{
		{
			name: "Gets the latest value",
			test: func(t T) {
				ctx := tests.Context(t)
				firstItem := CreateTestStruct(0, tester)
				tester.SetLatestValue(t, &firstItem)
				secondItem := CreateTestStruct(1, tester)
				tester.SetLatestValue(t, &secondItem)

				cr := tester.GetChainReader(t)
				require.NoError(t, cr.Bind(ctx, tester.GetBindings(t)))

				actual := &TestStruct{}
				params := &LatestParams{I: 1}
				require.NoError(t, cr.GetLatestValue(ctx, AnyContractName, MethodTakingLatestParamsReturningTestStruct, params, actual))
				assert.Equal(t, &firstItem, actual)

				params.I = 2
				actual = &TestStruct{}
				require.NoError(t, cr.GetLatestValue(ctx, AnyContractName, MethodTakingLatestParamsReturningTestStruct, params, actual))
				assert.Equal(t, &secondItem, actual)
			},
		},
		{
			name: "Get latest value without arguments and with primitive return",
			test: func(t T) {
				ctx := tests.Context(t)
				cr := tester.GetChainReader(t)
				require.NoError(t, cr.Bind(ctx, tester.GetBindings(t)))

				var prim uint64
				require.NoError(t, cr.GetLatestValue(ctx, AnyContractName, MethodReturningUint64, nil, &prim))

				assert.Equal(t, AnyValueToReadWithoutAnArgument, prim)
			},
		},
		{
			name: "Get latest value allows multiple contract names to have the same function Name",
			test: func(t T) {
				ctx := tests.Context(t)
				cr := tester.GetChainReader(t)
				bindings := tester.GetBindings(t)
				seenAddrs := map[string]bool{}
				for _, binding := range bindings {
					assert.False(t, seenAddrs[binding.Address])
					seenAddrs[binding.Address] = true
				}

				require.NoError(t, cr.Bind(ctx, bindings))

				var prim uint64
				require.NoError(t, cr.GetLatestValue(ctx, AnySecondContractName, MethodReturningUint64, nil, &prim))

				assert.Equal(t, AnyDifferentValueToReadWithoutAnArgument, prim)
			},
		},
		{
			name: "Get latest value without arguments and with slice return",
			test: func(t T) {
				ctx := tests.Context(t)
				cr := tester.GetChainReader(t)
				require.NoError(t, cr.Bind(ctx, tester.GetBindings(t)))

				var slice []uint64
				require.NoError(t, cr.GetLatestValue(ctx, AnyContractName, MethodReturningUint64Slice, nil, &slice))

				assert.Equal(t, AnySliceToReadWithoutAnArgument, slice)
			},
		},
		{
			name: "Get latest value wraps config with modifiers using its own mapstructure overrides",
			test: func(t T) {
				ctx := tests.Context(t)
				testStruct := CreateTestStruct(0, tester)
				testStruct.BigField = nil
				testStruct.Account = nil
				cr := tester.GetChainReader(t)
				require.NoError(t, cr.Bind(ctx, tester.GetBindings(t)))

				actual := &TestStructWithExtraField{}
				require.NoError(t, cr.GetLatestValue(ctx, AnyContractName, MethodReturningSeenStruct, testStruct, actual))

				expected := &TestStructWithExtraField{
					ExtraField: AnyExtraValue,
					TestStruct: CreateTestStruct(0, tester),
				}

				assert.Equal(t, expected, actual)
			},
		},
		{
			name: "Get latest value gets latest event",
			test: func(t T) {
				ctx := tests.Context(t)
				cr := tester.GetChainReader(t)
				require.NoError(t, cr.Bind(ctx, tester.GetBindings(t)))
				ts := CreateTestStruct[T](0, tester)
				tester.TriggerEvent(t, &ts)
				ts = CreateTestStruct[T](1, tester)
				tester.TriggerEvent(t, &ts)

				result := &TestStruct{}
				assert.Eventually(t, func() bool {
					err := cr.GetLatestValue(ctx, AnyContractName, EventName, nil, &result)
					return err == nil && reflect.DeepEqual(result, &ts)
				}, tester.MaxWaitTimeForEvents(), time.Millisecond*10)
			},
		},
		{
			name: "Get latest value returns not found if event was never triggered",
			test: func(t T) {
				ctx := tests.Context(t)
				cr := tester.GetChainReader(t)
				require.NoError(t, cr.Bind(ctx, tester.GetBindings(t)))

				result := &TestStruct{}
				err := cr.GetLatestValue(ctx, AnyContractName, EventName, nil, &result)
				assert.True(t, errors.Is(err, types.ErrNotFound))
			},
		},
		{
			name: "Get latest value gets latest event with filtering",
			test: func(t T) {
				ctx := tests.Context(t)
				cr := tester.GetChainReader(t)
				require.NoError(t, cr.Bind(ctx, tester.GetBindings(t)))
				ts0 := CreateTestStruct(0, tester)
				tester.TriggerEvent(t, &ts0)
				ts1 := CreateTestStruct(1, tester)
				tester.TriggerEvent(t, &ts1)

				filterParams := &FilterEventParams{Field: *ts0.Field}
				assert.Never(t, func() bool {
					result := &TestStruct{}
					err := cr.GetLatestValue(ctx, AnyContractName, EventWithFilterName, filterParams, &result)
					return err == nil && reflect.DeepEqual(result, &ts1)
				}, tester.MaxWaitTimeForEvents(), time.Millisecond*10)
				// get the result one more time to verify it.
				// Using the result from the Never statement by creating result outside the block is a data race
				result := &TestStruct{}
				err := cr.GetLatestValue(ctx, AnyContractName, EventWithFilterName, filterParams, &result)
				require.NoError(t, err)
				assert.Equal(t, &ts0, result)
			},
		},
	}
	runTests(t, tester, tests)
}

func runChainReaderBatchGetLatestValuesInterfaceTests[T TestingT[T]](t T, tester ChainReaderInterfaceTester[T]) {
	testCases := []testcase[T]{
		{
			name: "BatchGetLatestValues works",
			test: func(t T) {
				// setup test data
				firstItem := CreateTestStruct(1, tester)
				batchCallEntry := make(BatchCallEntry)
				batchCallEntry[AnyContractName] = ContractBatchEntry{{Name: MethodTakingLatestParamsReturningTestStruct, ReturnValue: &firstItem}}
				tester.SetBatchLatestValues(t, batchCallEntry)

				// setup call data
				params, actual := &LatestParams{I: 1}, &TestStruct{}
				batchGetLatestValueRequest := make(types.BatchGetLatestValuesRequest)
				batchGetLatestValueRequest[AnyContractName] = []types.BatchRead{{ReadName: MethodTakingLatestParamsReturningTestStruct, Params: params, ReturnVal: actual}}

				ctx := tests.Context(t)
				cr := tester.GetChainReader(t)

				require.NoError(t, cr.Bind(ctx, tester.GetBindings(t)))
				result, err := cr.BatchGetLatestValues(ctx, batchGetLatestValueRequest)
				require.NoError(t, err)

				anyContractBatch := result[AnyContractName]
				assert.Equal(t, MethodTakingLatestParamsReturningTestStruct, anyContractBatch[0].ReadName)
				assert.Equal(t, &firstItem, anyContractBatch[0].ReturnValue)
				assert.NoError(t, anyContractBatch[0].Err)
			},
		},
		{
			name: "BatchGetLatestValues works without arguments and with primitive return",
			test: func(t T) {
				// setup call data
				var primitiveReturnValue uint64
				batchGetLatestValuesRequest := make(types.BatchGetLatestValuesRequest)
				batchGetLatestValuesRequest[AnyContractName] = []types.BatchRead{{ReadName: MethodReturningUint64, Params: nil, ReturnVal: &primitiveReturnValue}}

				ctx := tests.Context(t)
				cr := tester.GetChainReader(t)
				require.NoError(t, cr.Bind(ctx, tester.GetBindings(t)))

				result, err := cr.BatchGetLatestValues(ctx, batchGetLatestValuesRequest)
				require.NoError(t, err)

				anyContractBatch := result[AnyContractName]
				require.NoError(t, anyContractBatch[0].Err)
				assert.Equal(t, MethodReturningUint64, anyContractBatch[0].ReadName)
				assert.Equal(t, AnyValueToReadWithoutAnArgument, *anyContractBatch[0].ReturnValue.(*uint64))
			},
		},
		{
			name: "BatchGetLatestValue without arguments and with slice return",
			test: func(t T) {
				// setup call data
				var sliceReturnValue []uint64
				batchGetLatestValueRequest := make(types.BatchGetLatestValuesRequest)
				batchGetLatestValueRequest[AnyContractName] = []types.BatchRead{{ReadName: MethodReturningUint64Slice, Params: nil, ReturnVal: &sliceReturnValue}}

				ctx := tests.Context(t)
				cr := tester.GetChainReader(t)
				require.NoError(t, cr.Bind(ctx, tester.GetBindings(t)))
				result, err := cr.BatchGetLatestValues(ctx, batchGetLatestValueRequest)
				require.NoError(t, err)

				anyContractBatch := result[AnyContractName]
				require.NoError(t, anyContractBatch[0].Err)
				assert.Equal(t, MethodReturningUint64Slice, anyContractBatch[0].ReadName)
				assert.Equal(t, AnySliceToReadWithoutAnArgument, *anyContractBatch[0].ReturnValue.(*[]uint64))
			},
		},
		{
			name: "BatchGetLatestValues wraps config with modifiers using its own mapstructure overrides",
			test: func(t T) {
				// setup call data
				testStruct := CreateTestStruct(0, tester)
				testStruct.BigField = nil
				testStruct.Account = nil
				actual := &TestStructWithExtraField{}
				batchGetLatestValueRequest := make(types.BatchGetLatestValuesRequest)
				batchGetLatestValueRequest[AnyContractName] = []types.BatchRead{{ReadName: MethodReturningSeenStruct, Params: testStruct, ReturnVal: actual}}

				ctx := tests.Context(t)
				cr := tester.GetChainReader(t)
				require.NoError(t, cr.Bind(ctx, tester.GetBindings(t)))
				result, err := cr.BatchGetLatestValues(ctx, batchGetLatestValueRequest)
				require.NoError(t, err)

				anyContractBatch := result[AnyContractName]
				require.NoError(t, anyContractBatch[0].Err)
				assert.Equal(t, MethodReturningSeenStruct, anyContractBatch[0].ReadName)
				assert.Equal(t,
					&TestStructWithExtraField{
						ExtraField: AnyExtraValue,
						TestStruct: CreateTestStruct(0, tester),
					},
					anyContractBatch[0].ReturnValue)
			},
		},
		{
			name: "BatchGetLatestValues supports same read with different params and results retain order from request",
			test: func(t T) {
				batchCallEntry := make(BatchCallEntry)
				batchGetLatestValueRequest := make(types.BatchGetLatestValuesRequest)
				for i := 0; i < 10; i++ {
					// setup test data
					ts := CreateTestStruct(i, tester)
					batchCallEntry[AnyContractName] = append(batchCallEntry[AnyContractName], ReadEntry{Name: MethodTakingLatestParamsReturningTestStruct, ReturnValue: &ts})
					// setup call data
					batchGetLatestValueRequest[AnyContractName] = append(batchGetLatestValueRequest[AnyContractName], types.BatchRead{ReadName: MethodTakingLatestParamsReturningTestStruct, Params: &LatestParams{I: 1 + i}, ReturnVal: &TestStruct{}})
				}
				tester.SetBatchLatestValues(t, batchCallEntry)

				ctx := tests.Context(t)
				cr := tester.GetChainReader(t)
				require.NoError(t, cr.Bind(ctx, tester.GetBindings(t)))

				result, err := cr.BatchGetLatestValues(ctx, batchGetLatestValueRequest)
				require.NoError(t, err)

				for i := 0; i < 10; i++ {
					resultAnyContract, testDataAnyContract := result[AnyContractName], batchCallEntry[AnyContractName]
					assert.Equal(t, MethodTakingLatestParamsReturningTestStruct, resultAnyContract[i].ReadName)
					assert.NoError(t, resultAnyContract[i].Err)
					assert.Equal(t, testDataAnyContract[i].ReturnValue, resultAnyContract[i].ReturnValue)
				}
			},
		},
		{
			name: "BatchGetLatestValues supports same read with different params and results retain order from request even with multiple contracts",
			test: func(t T) {
				batchCallEntry := make(BatchCallEntry)
				batchGetLatestValueRequest := make(types.BatchGetLatestValuesRequest)
				for i := 0; i < 10; i++ {
					// setup test data
					ts1, ts2 := CreateTestStruct(i, tester), CreateTestStruct(i+10, tester)
					batchCallEntry[AnyContractName] = append(batchCallEntry[AnyContractName], ReadEntry{Name: MethodTakingLatestParamsReturningTestStruct, ReturnValue: &ts1})
					batchCallEntry[AnySecondContractName] = append(batchCallEntry[AnySecondContractName], ReadEntry{Name: MethodTakingLatestParamsReturningTestStruct, ReturnValue: &ts2})
					// setup call data
					batchGetLatestValueRequest[AnyContractName] = append(batchGetLatestValueRequest[AnyContractName], types.BatchRead{ReadName: MethodTakingLatestParamsReturningTestStruct, Params: &LatestParams{I: 1 + i}, ReturnVal: &TestStruct{}})
					batchGetLatestValueRequest[AnySecondContractName] = append(batchGetLatestValueRequest[AnySecondContractName], types.BatchRead{ReadName: MethodTakingLatestParamsReturningTestStruct, Params: &LatestParams{I: 1 + i}, ReturnVal: &TestStruct{}})
				}
				tester.SetBatchLatestValues(t, batchCallEntry)

				ctx := tests.Context(t)
				cr := tester.GetChainReader(t)
				require.NoError(t, cr.Bind(ctx, tester.GetBindings(t)))

				result, err := cr.BatchGetLatestValues(ctx, batchGetLatestValueRequest)
				require.NoError(t, err)

				for i := 0; i < 10; i++ {
					resultAnyContract, testDataAnyContract := result[AnyContractName], batchCallEntry[AnyContractName]
					resultAnySecondContract, testDataAnySecondContract := result[AnySecondContractName], batchCallEntry[AnySecondContractName]
					assert.Equal(t, MethodTakingLatestParamsReturningTestStruct, resultAnyContract[i].ReadName)
					assert.NoError(t, resultAnyContract[i].Err)
					assert.NoError(t, resultAnySecondContract[i].Err)
					assert.Equal(t, MethodTakingLatestParamsReturningTestStruct, resultAnySecondContract[i].ReadName)
					assert.Equal(t, testDataAnyContract[i].ReturnValue, resultAnyContract[i].ReturnValue)
					assert.Equal(t, testDataAnySecondContract[i].ReturnValue, resultAnySecondContract[i].ReturnValue)
				}
			},
		},
		{
			name: "BatchGetLatestValues sets errors properly",
			test: func(t T) {
				batchGetLatestValueRequest := make(types.BatchGetLatestValuesRequest)
				for i := 0; i < 10; i++ {
					// setup call data and set invalid params that cause an error
					batchGetLatestValueRequest[AnyContractName] = append(batchGetLatestValueRequest[AnyContractName], types.BatchRead{ReadName: MethodTakingLatestParamsReturningTestStruct, Params: &LatestParams{I: 0}, ReturnVal: &TestStruct{}})
					batchGetLatestValueRequest[AnySecondContractName] = append(batchGetLatestValueRequest[AnySecondContractName], types.BatchRead{ReadName: MethodTakingLatestParamsReturningTestStruct, Params: &LatestParams{I: 0}, ReturnVal: &TestStruct{}})
				}

				ctx := tests.Context(t)
				cr := tester.GetChainReader(t)
				require.NoError(t, cr.Bind(ctx, tester.GetBindings(t)))

				result, err := cr.BatchGetLatestValues(ctx, batchGetLatestValueRequest)
				require.NoError(t, err)

				for i := 0; i < 10; i++ {
					resultAnyContract := result[AnyContractName]
					resultAnySecondContract := result[AnySecondContractName]
					assert.Equal(t, MethodTakingLatestParamsReturningTestStruct, resultAnyContract[i].ReadName)
					assert.Error(t, resultAnyContract[i].Err)
					assert.Error(t, resultAnySecondContract[i].Err)
					assert.Equal(t, MethodTakingLatestParamsReturningTestStruct, resultAnySecondContract[i].ReadName)
					assert.Equal(t, &TestStruct{}, resultAnyContract[i].ReturnValue)
					assert.Equal(t, &TestStruct{}, resultAnySecondContract[i].ReturnValue)
				}
			},
		},
	}

	runTests(t, tester, testCases)
}

func runQueryKeyInterfaceTests[T TestingT[T]](t T, tester ChainReaderInterfaceTester[T]) {
	tests := []testcase[T]{
		{
			name: "QueryKey returns not found if sequence never happened",
			test: func(t T) {
				ctx := tests.Context(t)
				cr := tester.GetChainReader(t)

				require.NoError(t, cr.Bind(ctx, tester.GetBindings(t)))

				logs, err := cr.QueryKey(ctx, AnyContractName, query.KeyFilter{Key: EventName}, query.LimitAndSort{}, &TestStruct{})

				require.NoError(t, err)
				assert.Len(t, logs, 0)
			},
		},
		{
			name: "QueryKey returns sequence data properly",
			test: func(t T) {
				ctx := tests.Context(t)
				cr := tester.GetChainReader(t)
				require.NoError(t, cr.Bind(ctx, tester.GetBindings(t)))
				ts1 := CreateTestStruct[T](0, tester)
				tester.TriggerEvent(t, &ts1)
				ts2 := CreateTestStruct[T](1, tester)
				tester.TriggerEvent(t, &ts2)

				ts := &TestStruct{}
				assert.Eventually(t, func() bool {
					// sequences from queryKey without limit and sort should be in descending order
					sequences, err := cr.QueryKey(ctx, AnyContractName, query.KeyFilter{Key: EventName}, query.LimitAndSort{}, ts)
					return err == nil && len(sequences) == 2 && reflect.DeepEqual(&ts1, sequences[1].Data) && reflect.DeepEqual(&ts2, sequences[0].Data)
				}, tester.MaxWaitTimeForEvents(), time.Millisecond*10)
			},
		},
	}

	runTests(t, tester, tests)
}
