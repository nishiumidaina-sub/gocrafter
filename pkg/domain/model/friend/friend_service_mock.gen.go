// Code generated by MockGen. DO NOT EDIT.
// Source: ./friend_service.go

// Package friend is a generated GoMock package.
package friend

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockFriendService is a mock of FriendService interface.
type MockFriendService struct {
	ctrl     *gomock.Controller
	recorder *MockFriendServiceMockRecorder
}

// MockFriendServiceMockRecorder is the mock recorder for MockFriendService.
type MockFriendServiceMockRecorder struct {
	mock *MockFriendService
}

// NewMockFriendService creates a new mock instance.
func NewMockFriendService(ctrl *gomock.Controller) *MockFriendService {
	mock := &MockFriendService{ctrl: ctrl}
	mock.recorder = &MockFriendServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFriendService) EXPECT() *MockFriendServiceMockRecorder {
	return m.recorder
}

// Approve mocks base method.
func (m *MockFriendService) Approve(ctx context.Context, txs map[string]*gorm.DB, req *FriendApproveRequest) (*FriendApproveResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Approve", ctx, txs, req)
	ret0, _ := ret[0].(*FriendApproveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Approve indicates an expected call of Approve.
func (mr *MockFriendServiceMockRecorder) Approve(ctx, txs, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Approve", reflect.TypeOf((*MockFriendService)(nil).Approve), ctx, txs, req)
}

// Delete mocks base method.
func (m *MockFriendService) Delete(ctx context.Context, txs map[string]*gorm.DB, req *FriendDeleteRequest) (*FriendDeleteResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, txs, req)
	ret0, _ := ret[0].(*FriendDeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockFriendServiceMockRecorder) Delete(ctx, txs, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFriendService)(nil).Delete), ctx, txs, req)
}

// Disapprove mocks base method.
func (m *MockFriendService) Disapprove(ctx context.Context, txs map[string]*gorm.DB, req *FriendDisapproveRequest) (*FriendDisapproveResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disapprove", ctx, txs, req)
	ret0, _ := ret[0].(*FriendDisapproveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Disapprove indicates an expected call of Disapprove.
func (mr *MockFriendServiceMockRecorder) Disapprove(ctx, txs, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disapprove", reflect.TypeOf((*MockFriendService)(nil).Disapprove), ctx, txs, req)
}

// Get mocks base method.
func (m *MockFriendService) Get(ctx context.Context, req *FriendGetRequest) (*FriendGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, req)
	ret0, _ := ret[0].(*FriendGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockFriendServiceMockRecorder) Get(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockFriendService)(nil).Get), ctx, req)
}

// Send mocks base method.
func (m *MockFriendService) Send(ctx context.Context, txs map[string]*gorm.DB, req *FriendSendRequest) (*FriendSendResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", ctx, txs, req)
	ret0, _ := ret[0].(*FriendSendResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send.
func (mr *MockFriendServiceMockRecorder) Send(ctx, txs, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockFriendService)(nil).Send), ctx, txs, req)
}
