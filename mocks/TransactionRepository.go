// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	mock "github.com/stretchr/testify/mock"

	model "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
)

// TransactionRepository is an autogenerated mock type for the TransactionRepository type
type TransactionRepository struct {
	mock.Mock
}

// FindListTransaction provides a mock function with given fields: _a0, _a1
func (_m *TransactionRepository) FindListTransaction(_a0 context.Context, _a1 dto.ListTransactionsReq) (*dto.TransactionPaginationRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *dto.TransactionPaginationRes
	if rf, ok := ret.Get(0).(func(context.Context, dto.ListTransactionsReq) *dto.TransactionPaginationRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.TransactionPaginationRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, dto.ListTransactionsReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TopUpTransaction provides a mock function with given fields: _a0, _a1
func (_m *TransactionRepository) TopUpTransaction(_a0 context.Context, _a1 model.Transaction) (*model.Transaction, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, model.Transaction) *model.Transaction); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Transaction) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransferTransaction provides a mock function with given fields: _a0, _a1
func (_m *TransactionRepository) TransferTransaction(_a0 context.Context, _a1 model.Transaction) (*model.Transaction, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, model.Transaction) *model.Transaction); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Transaction) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTransactionRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionRepository creates a new instance of TransactionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionRepository(t mockConstructorTestingTNewTransactionRepository) *TransactionRepository {
	mock := &TransactionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
