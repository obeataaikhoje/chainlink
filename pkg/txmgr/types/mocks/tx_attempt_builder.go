// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	context "context"

	feetypes "github.com/smartcontractkit/chainlink-relay/pkg/fee/types"
	logger "github.com/smartcontractkit/chainlink-relay/pkg/logger"

	mock "github.com/stretchr/testify/mock"

	txmgrtypes "github.com/smartcontractkit/chainlink-relay/pkg/txmgr/types"

	types "github.com/smartcontractkit/chainlink-relay/pkg/types"
)

// TxAttemptBuilder is an autogenerated mock type for the TxAttemptBuilder type
type TxAttemptBuilder[CHAIN_ID types.ID, HEAD types.Head[BLOCK_HASH], ADDR types.Hashable, TX_HASH types.Hashable, BLOCK_HASH types.Hashable, SEQ types.Sequence, FEE feetypes.Fee] struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HealthReport provides a mock function with given fields:
func (_m *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) HealthReport() map[string]error {
	ret := _m.Called()

	var r0 map[string]error
	if rf, ok := ret.Get(0).(func() map[string]error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]error)
		}
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewBumpTxAttempt provides a mock function with given fields: ctx, tx, previousAttempt, priorAttempts, lggr
func (_m *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) NewBumpTxAttempt(ctx context.Context, tx txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], previousAttempt txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], priorAttempts []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], lggr logger.Logger) (txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], FEE, uint32, bool, error) {
	ret := _m.Called(ctx, tx, previousAttempt, priorAttempts, lggr)

	var r0 txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 FEE
	var r2 uint32
	var r3 bool
	var r4 error
	if rf, ok := ret.Get(0).(func(context.Context, txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], logger.Logger) (txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], FEE, uint32, bool, error)); ok {
		return rf(ctx, tx, previousAttempt, priorAttempts, lggr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], logger.Logger) txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(ctx, tx, previousAttempt, priorAttempts, lggr)
	} else {
		r0 = ret.Get(0).(txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
	}

	if rf, ok := ret.Get(1).(func(context.Context, txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], logger.Logger) FEE); ok {
		r1 = rf(ctx, tx, previousAttempt, priorAttempts, lggr)
	} else {
		r1 = ret.Get(1).(FEE)
	}

	if rf, ok := ret.Get(2).(func(context.Context, txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], logger.Logger) uint32); ok {
		r2 = rf(ctx, tx, previousAttempt, priorAttempts, lggr)
	} else {
		r2 = ret.Get(2).(uint32)
	}

	if rf, ok := ret.Get(3).(func(context.Context, txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], logger.Logger) bool); ok {
		r3 = rf(ctx, tx, previousAttempt, priorAttempts, lggr)
	} else {
		r3 = ret.Get(3).(bool)
	}

	if rf, ok := ret.Get(4).(func(context.Context, txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], logger.Logger) error); ok {
		r4 = rf(ctx, tx, previousAttempt, priorAttempts, lggr)
	} else {
		r4 = ret.Error(4)
	}

	return r0, r1, r2, r3, r4
}

// NewCustomTxAttempt provides a mock function with given fields: tx, fee, gasLimit, txType, lggr
func (_m *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) NewCustomTxAttempt(tx txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], fee FEE, gasLimit uint32, txType int, lggr logger.Logger) (txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], bool, error) {
	ret := _m.Called(tx, fee, gasLimit, txType, lggr)

	var r0 txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], FEE, uint32, int, logger.Logger) (txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], bool, error)); ok {
		return rf(tx, fee, gasLimit, txType, lggr)
	}
	if rf, ok := ret.Get(0).(func(txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], FEE, uint32, int, logger.Logger) txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(tx, fee, gasLimit, txType, lggr)
	} else {
		r0 = ret.Get(0).(txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
	}

	if rf, ok := ret.Get(1).(func(txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], FEE, uint32, int, logger.Logger) bool); ok {
		r1 = rf(tx, fee, gasLimit, txType, lggr)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], FEE, uint32, int, logger.Logger) error); ok {
		r2 = rf(tx, fee, gasLimit, txType, lggr)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewEmptyTxAttempt provides a mock function with given fields: seq, feeLimit, fee, fromAddress
func (_m *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) NewEmptyTxAttempt(seq SEQ, feeLimit uint32, fee FEE, fromAddress ADDR) (txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(seq, feeLimit, fee, fromAddress)

	var r0 txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(SEQ, uint32, FEE, ADDR) (txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(seq, feeLimit, fee, fromAddress)
	}
	if rf, ok := ret.Get(0).(func(SEQ, uint32, FEE, ADDR) txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(seq, feeLimit, fee, fromAddress)
	} else {
		r0 = ret.Get(0).(txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
	}

	if rf, ok := ret.Get(1).(func(SEQ, uint32, FEE, ADDR) error); ok {
		r1 = rf(seq, feeLimit, fee, fromAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTxAttempt provides a mock function with given fields: ctx, tx, lggr, opts
func (_m *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) NewTxAttempt(ctx context.Context, tx txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], lggr logger.Logger, opts ...feetypes.Opt) (txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], FEE, uint32, bool, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, tx, lggr)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 FEE
	var r2 uint32
	var r3 bool
	var r4 error
	if rf, ok := ret.Get(0).(func(context.Context, txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], logger.Logger, ...feetypes.Opt) (txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], FEE, uint32, bool, error)); ok {
		return rf(ctx, tx, lggr, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], logger.Logger, ...feetypes.Opt) txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(ctx, tx, lggr, opts...)
	} else {
		r0 = ret.Get(0).(txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
	}

	if rf, ok := ret.Get(1).(func(context.Context, txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], logger.Logger, ...feetypes.Opt) FEE); ok {
		r1 = rf(ctx, tx, lggr, opts...)
	} else {
		r1 = ret.Get(1).(FEE)
	}

	if rf, ok := ret.Get(2).(func(context.Context, txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], logger.Logger, ...feetypes.Opt) uint32); ok {
		r2 = rf(ctx, tx, lggr, opts...)
	} else {
		r2 = ret.Get(2).(uint32)
	}

	if rf, ok := ret.Get(3).(func(context.Context, txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], logger.Logger, ...feetypes.Opt) bool); ok {
		r3 = rf(ctx, tx, lggr, opts...)
	} else {
		r3 = ret.Get(3).(bool)
	}

	if rf, ok := ret.Get(4).(func(context.Context, txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], logger.Logger, ...feetypes.Opt) error); ok {
		r4 = rf(ctx, tx, lggr, opts...)
	} else {
		r4 = ret.Error(4)
	}

	return r0, r1, r2, r3, r4
}

// NewTxAttemptWithType provides a mock function with given fields: ctx, tx, lggr, txType, opts
func (_m *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) NewTxAttemptWithType(ctx context.Context, tx txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], lggr logger.Logger, txType int, opts ...feetypes.Opt) (txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], FEE, uint32, bool, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, tx, lggr, txType)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 FEE
	var r2 uint32
	var r3 bool
	var r4 error
	if rf, ok := ret.Get(0).(func(context.Context, txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], logger.Logger, int, ...feetypes.Opt) (txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], FEE, uint32, bool, error)); ok {
		return rf(ctx, tx, lggr, txType, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], logger.Logger, int, ...feetypes.Opt) txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(ctx, tx, lggr, txType, opts...)
	} else {
		r0 = ret.Get(0).(txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
	}

	if rf, ok := ret.Get(1).(func(context.Context, txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], logger.Logger, int, ...feetypes.Opt) FEE); ok {
		r1 = rf(ctx, tx, lggr, txType, opts...)
	} else {
		r1 = ret.Get(1).(FEE)
	}

	if rf, ok := ret.Get(2).(func(context.Context, txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], logger.Logger, int, ...feetypes.Opt) uint32); ok {
		r2 = rf(ctx, tx, lggr, txType, opts...)
	} else {
		r2 = ret.Get(2).(uint32)
	}

	if rf, ok := ret.Get(3).(func(context.Context, txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], logger.Logger, int, ...feetypes.Opt) bool); ok {
		r3 = rf(ctx, tx, lggr, txType, opts...)
	} else {
		r3 = ret.Get(3).(bool)
	}

	if rf, ok := ret.Get(4).(func(context.Context, txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], logger.Logger, int, ...feetypes.Opt) error); ok {
		r4 = rf(ctx, tx, lggr, txType, opts...)
	} else {
		r4 = ret.Error(4)
	}

	return r0, r1, r2, r3, r4
}

// OnNewLongestChain provides a mock function with given fields: ctx, head
func (_m *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) OnNewLongestChain(ctx context.Context, head HEAD) {
	_m.Called(ctx, head)
}

// Ready provides a mock function with given fields:
func (_m *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) Ready() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTxAttemptBuilder interface {
	mock.TestingT
	Cleanup(func())
}

// NewTxAttemptBuilder creates a new instance of TxAttemptBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTxAttemptBuilder[CHAIN_ID types.ID, HEAD types.Head[BLOCK_HASH], ADDR types.Hashable, TX_HASH types.Hashable, BLOCK_HASH types.Hashable, SEQ types.Sequence, FEE feetypes.Fee](t mockConstructorTestingTNewTxAttemptBuilder) *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE] {
	mock := &TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
