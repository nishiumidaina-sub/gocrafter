// Code generated by MockGen. DO NOT EDIT.
// Source: ./shard_service.go

// Package shard is a generated GoMock package.
package shard

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockShardService is a mock of ShardService interface.
type MockShardService struct {
	ctrl     *gomock.Controller
	recorder *MockShardServiceMockRecorder
}

// MockShardServiceMockRecorder is the mock recorder for MockShardService.
type MockShardServiceMockRecorder struct {
	mock *MockShardService
}

// NewMockShardService creates a new mock instance.
func NewMockShardService(ctrl *gomock.Controller) *MockShardService {
	mock := &MockShardService{ctrl: ctrl}
	mock.recorder = &MockShardServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShardService) EXPECT() *MockShardServiceMockRecorder {
	return m.recorder
}

// GetShardKey mocks base method.
func (m *MockShardService) GetShardKey(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShardKey", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShardKey indicates an expected call of GetShardKey.
func (mr *MockShardServiceMockRecorder) GetShardKey(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShardKey", reflect.TypeOf((*MockShardService)(nil).GetShardKey), ctx)
}
