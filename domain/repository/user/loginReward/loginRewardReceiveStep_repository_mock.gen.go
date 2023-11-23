// Code generated by MockGen. DO NOT EDIT.
// Source: ./loginRewardReceiveStep_repository.gen.go

// Package loginReward is a generated GoMock package.
package loginReward

import (
	reflect "reflect"

	loginReward "github.com/game-core/gocrafter/domain/entity/user/loginReward"
	gomock "github.com/golang/mock/gomock"
	gorm "github.com/jinzhu/gorm"
)

// MockLoginRewardReceiveStepRepository is a mock of LoginRewardReceiveStepRepository interface.
type MockLoginRewardReceiveStepRepository struct {
	ctrl     *gomock.Controller
	recorder *MockLoginRewardReceiveStepRepositoryMockRecorder
}

// MockLoginRewardReceiveStepRepositoryMockRecorder is the mock recorder for MockLoginRewardReceiveStepRepository.
type MockLoginRewardReceiveStepRepositoryMockRecorder struct {
	mock *MockLoginRewardReceiveStepRepository
}

// NewMockLoginRewardReceiveStepRepository creates a new mock instance.
func NewMockLoginRewardReceiveStepRepository(ctrl *gomock.Controller) *MockLoginRewardReceiveStepRepository {
	mock := &MockLoginRewardReceiveStepRepository{ctrl: ctrl}
	mock.recorder = &MockLoginRewardReceiveStepRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoginRewardReceiveStepRepository) EXPECT() *MockLoginRewardReceiveStepRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockLoginRewardReceiveStepRepository) Create(entity *loginReward.LoginRewardReceiveStep, shardKey int, tx *gorm.DB) (*loginReward.LoginRewardReceiveStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", entity, shardKey, tx)
	ret0, _ := ret[0].(*loginReward.LoginRewardReceiveStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockLoginRewardReceiveStepRepositoryMockRecorder) Create(entity, shardKey, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLoginRewardReceiveStepRepository)(nil).Create), entity, shardKey, tx)
}

// Delete mocks base method.
func (m *MockLoginRewardReceiveStepRepository) Delete(entity *loginReward.LoginRewardReceiveStep, shardKey int, tx *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", entity, shardKey, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockLoginRewardReceiveStepRepositoryMockRecorder) Delete(entity, shardKey, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockLoginRewardReceiveStepRepository)(nil).Delete), entity, shardKey, tx)
}

// FindByAccountID mocks base method.
func (m *MockLoginRewardReceiveStepRepository) FindByAccountID(AccountID int64, shardKey int) (*loginReward.LoginRewardReceiveStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByAccountID", AccountID, shardKey)
	ret0, _ := ret[0].(*loginReward.LoginRewardReceiveStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByAccountID indicates an expected call of FindByAccountID.
func (mr *MockLoginRewardReceiveStepRepositoryMockRecorder) FindByAccountID(AccountID, shardKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByAccountID", reflect.TypeOf((*MockLoginRewardReceiveStepRepository)(nil).FindByAccountID), AccountID, shardKey)
}

// FindByAccountIDAndLoginRewardStatusID mocks base method.
func (m *MockLoginRewardReceiveStepRepository) FindByAccountIDAndLoginRewardStatusID(AccountID, LoginRewardStatusID int64, shardKey int) (*loginReward.LoginRewardReceiveStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByAccountIDAndLoginRewardStatusID", AccountID, LoginRewardStatusID, shardKey)
	ret0, _ := ret[0].(*loginReward.LoginRewardReceiveStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByAccountIDAndLoginRewardStatusID indicates an expected call of FindByAccountIDAndLoginRewardStatusID.
func (mr *MockLoginRewardReceiveStepRepositoryMockRecorder) FindByAccountIDAndLoginRewardStatusID(AccountID, LoginRewardStatusID, shardKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByAccountIDAndLoginRewardStatusID", reflect.TypeOf((*MockLoginRewardReceiveStepRepository)(nil).FindByAccountIDAndLoginRewardStatusID), AccountID, LoginRewardStatusID, shardKey)
}

// FindByID mocks base method.
func (m *MockLoginRewardReceiveStepRepository) FindByID(ID int64, shardKey int) (*loginReward.LoginRewardReceiveStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ID, shardKey)
	ret0, _ := ret[0].(*loginReward.LoginRewardReceiveStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockLoginRewardReceiveStepRepositoryMockRecorder) FindByID(ID, shardKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockLoginRewardReceiveStepRepository)(nil).FindByID), ID, shardKey)
}

// FindByLoginRewardStatusID mocks base method.
func (m *MockLoginRewardReceiveStepRepository) FindByLoginRewardStatusID(LoginRewardStatusID int64, shardKey int) (*loginReward.LoginRewardReceiveStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByLoginRewardStatusID", LoginRewardStatusID, shardKey)
	ret0, _ := ret[0].(*loginReward.LoginRewardReceiveStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByLoginRewardStatusID indicates an expected call of FindByLoginRewardStatusID.
func (mr *MockLoginRewardReceiveStepRepositoryMockRecorder) FindByLoginRewardStatusID(LoginRewardStatusID, shardKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByLoginRewardStatusID", reflect.TypeOf((*MockLoginRewardReceiveStepRepository)(nil).FindByLoginRewardStatusID), LoginRewardStatusID, shardKey)
}

// FindOrNilByAccountID mocks base method.
func (m *MockLoginRewardReceiveStepRepository) FindOrNilByAccountID(AccountID int64, shardKey int) (*loginReward.LoginRewardReceiveStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByAccountID", AccountID, shardKey)
	ret0, _ := ret[0].(*loginReward.LoginRewardReceiveStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByAccountID indicates an expected call of FindOrNilByAccountID.
func (mr *MockLoginRewardReceiveStepRepositoryMockRecorder) FindOrNilByAccountID(AccountID, shardKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByAccountID", reflect.TypeOf((*MockLoginRewardReceiveStepRepository)(nil).FindOrNilByAccountID), AccountID, shardKey)
}

// FindOrNilByAccountIDAndLoginRewardStatusID mocks base method.
func (m *MockLoginRewardReceiveStepRepository) FindOrNilByAccountIDAndLoginRewardStatusID(AccountID, LoginRewardStatusID int64, shardKey int) (*loginReward.LoginRewardReceiveStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByAccountIDAndLoginRewardStatusID", AccountID, LoginRewardStatusID, shardKey)
	ret0, _ := ret[0].(*loginReward.LoginRewardReceiveStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByAccountIDAndLoginRewardStatusID indicates an expected call of FindOrNilByAccountIDAndLoginRewardStatusID.
func (mr *MockLoginRewardReceiveStepRepositoryMockRecorder) FindOrNilByAccountIDAndLoginRewardStatusID(AccountID, LoginRewardStatusID, shardKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByAccountIDAndLoginRewardStatusID", reflect.TypeOf((*MockLoginRewardReceiveStepRepository)(nil).FindOrNilByAccountIDAndLoginRewardStatusID), AccountID, LoginRewardStatusID, shardKey)
}

// FindOrNilByID mocks base method.
func (m *MockLoginRewardReceiveStepRepository) FindOrNilByID(ID int64, shardKey int) (*loginReward.LoginRewardReceiveStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByID", ID, shardKey)
	ret0, _ := ret[0].(*loginReward.LoginRewardReceiveStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByID indicates an expected call of FindOrNilByID.
func (mr *MockLoginRewardReceiveStepRepositoryMockRecorder) FindOrNilByID(ID, shardKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByID", reflect.TypeOf((*MockLoginRewardReceiveStepRepository)(nil).FindOrNilByID), ID, shardKey)
}

// FindOrNilByLoginRewardStatusID mocks base method.
func (m *MockLoginRewardReceiveStepRepository) FindOrNilByLoginRewardStatusID(LoginRewardStatusID int64, shardKey int) (*loginReward.LoginRewardReceiveStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByLoginRewardStatusID", LoginRewardStatusID, shardKey)
	ret0, _ := ret[0].(*loginReward.LoginRewardReceiveStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByLoginRewardStatusID indicates an expected call of FindOrNilByLoginRewardStatusID.
func (mr *MockLoginRewardReceiveStepRepositoryMockRecorder) FindOrNilByLoginRewardStatusID(LoginRewardStatusID, shardKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByLoginRewardStatusID", reflect.TypeOf((*MockLoginRewardReceiveStepRepository)(nil).FindOrNilByLoginRewardStatusID), LoginRewardStatusID, shardKey)
}

// List mocks base method.
func (m *MockLoginRewardReceiveStepRepository) List(limit int64, shardKey int) (*loginReward.LoginRewardReceiveSteps, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", limit, shardKey)
	ret0, _ := ret[0].(*loginReward.LoginRewardReceiveSteps)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockLoginRewardReceiveStepRepositoryMockRecorder) List(limit, shardKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockLoginRewardReceiveStepRepository)(nil).List), limit, shardKey)
}

// ListByAccountID mocks base method.
func (m *MockLoginRewardReceiveStepRepository) ListByAccountID(AccountID int64, shardKey int) (*loginReward.LoginRewardReceiveSteps, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByAccountID", AccountID, shardKey)
	ret0, _ := ret[0].(*loginReward.LoginRewardReceiveSteps)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByAccountID indicates an expected call of ListByAccountID.
func (mr *MockLoginRewardReceiveStepRepositoryMockRecorder) ListByAccountID(AccountID, shardKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByAccountID", reflect.TypeOf((*MockLoginRewardReceiveStepRepository)(nil).ListByAccountID), AccountID, shardKey)
}

// ListByAccountIDAndLoginRewardStatusID mocks base method.
func (m *MockLoginRewardReceiveStepRepository) ListByAccountIDAndLoginRewardStatusID(AccountID, LoginRewardStatusID int64, shardKey int) (*loginReward.LoginRewardReceiveSteps, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByAccountIDAndLoginRewardStatusID", AccountID, LoginRewardStatusID, shardKey)
	ret0, _ := ret[0].(*loginReward.LoginRewardReceiveSteps)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByAccountIDAndLoginRewardStatusID indicates an expected call of ListByAccountIDAndLoginRewardStatusID.
func (mr *MockLoginRewardReceiveStepRepositoryMockRecorder) ListByAccountIDAndLoginRewardStatusID(AccountID, LoginRewardStatusID, shardKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByAccountIDAndLoginRewardStatusID", reflect.TypeOf((*MockLoginRewardReceiveStepRepository)(nil).ListByAccountIDAndLoginRewardStatusID), AccountID, LoginRewardStatusID, shardKey)
}

// ListByLoginRewardStatusID mocks base method.
func (m *MockLoginRewardReceiveStepRepository) ListByLoginRewardStatusID(LoginRewardStatusID int64, shardKey int) (*loginReward.LoginRewardReceiveSteps, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByLoginRewardStatusID", LoginRewardStatusID, shardKey)
	ret0, _ := ret[0].(*loginReward.LoginRewardReceiveSteps)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByLoginRewardStatusID indicates an expected call of ListByLoginRewardStatusID.
func (mr *MockLoginRewardReceiveStepRepositoryMockRecorder) ListByLoginRewardStatusID(LoginRewardStatusID, shardKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByLoginRewardStatusID", reflect.TypeOf((*MockLoginRewardReceiveStepRepository)(nil).ListByLoginRewardStatusID), LoginRewardStatusID, shardKey)
}

// Update mocks base method.
func (m *MockLoginRewardReceiveStepRepository) Update(entity *loginReward.LoginRewardReceiveStep, shardKey int, tx *gorm.DB) (*loginReward.LoginRewardReceiveStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", entity, shardKey, tx)
	ret0, _ := ret[0].(*loginReward.LoginRewardReceiveStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockLoginRewardReceiveStepRepositoryMockRecorder) Update(entity, shardKey, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockLoginRewardReceiveStepRepository)(nil).Update), entity, shardKey, tx)
}