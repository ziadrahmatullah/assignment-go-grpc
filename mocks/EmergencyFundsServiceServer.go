// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	pb "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/pb"
	mock "github.com/stretchr/testify/mock"
)

// EmergencyFundsServiceServer is an autogenerated mock type for the EmergencyFundsServiceServer type
type EmergencyFundsServiceServer struct {
	mock.Mock
}

// CalculateEmergencyFunds provides a mock function with given fields: _a0, _a1
func (_m *EmergencyFundsServiceServer) CalculateEmergencyFunds(_a0 context.Context, _a1 *pb.EmergencyFundsReq) (*pb.EmergencyFundsRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.EmergencyFundsRes
	if rf, ok := ret.Get(0).(func(context.Context, *pb.EmergencyFundsReq) *pb.EmergencyFundsRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.EmergencyFundsRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.EmergencyFundsReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedEmergencyFundsServiceServer provides a mock function with given fields:
func (_m *EmergencyFundsServiceServer) mustEmbedUnimplementedEmergencyFundsServiceServer() {
	_m.Called()
}

type mockConstructorTestingTNewEmergencyFundsServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewEmergencyFundsServiceServer creates a new instance of EmergencyFundsServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEmergencyFundsServiceServer(t mockConstructorTestingTNewEmergencyFundsServiceServer) *EmergencyFundsServiceServer {
	mock := &EmergencyFundsServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
