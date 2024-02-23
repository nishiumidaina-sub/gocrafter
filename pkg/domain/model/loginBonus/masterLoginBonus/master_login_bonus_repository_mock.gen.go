// Code generated by MockGen. DO NOT EDIT.
// Source: ./master_login_bonus_repository.gen.go

// Package masterLoginBonus is a generated GoMock package.
package masterLoginBonus

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockMasterLoginBonusRepository is a mock of MasterLoginBonusRepository interface.
type MockMasterLoginBonusRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMasterLoginBonusRepositoryMockRecorder
}

// MockMasterLoginBonusRepositoryMockRecorder is the mock recorder for MockMasterLoginBonusRepository.
type MockMasterLoginBonusRepositoryMockRecorder struct {
	mock *MockMasterLoginBonusRepository
}

// NewMockMasterLoginBonusRepository creates a new mock instance.
func NewMockMasterLoginBonusRepository(ctrl *gomock.Controller) *MockMasterLoginBonusRepository {
	mock := &MockMasterLoginBonusRepository{ctrl: ctrl}
	mock.recorder = &MockMasterLoginBonusRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMasterLoginBonusRepository) EXPECT() *MockMasterLoginBonusRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockMasterLoginBonusRepository) Create(ctx context.Context, tx *gorm.DB, m *MasterLoginBonus) (*MasterLoginBonus, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*MasterLoginBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMasterLoginBonusRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMasterLoginBonusRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockMasterLoginBonusRepository) CreateList(ctx context.Context, tx *gorm.DB, ms MasterLoginBonuses) (MasterLoginBonuses, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(MasterLoginBonuses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockMasterLoginBonusRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockMasterLoginBonusRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockMasterLoginBonusRepository) Delete(ctx context.Context, tx *gorm.DB, m *MasterLoginBonus) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMasterLoginBonusRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMasterLoginBonusRepository)(nil).Delete), ctx, tx, m)
}

// FinOrNilByMasterEventId mocks base method.
func (m *MockMasterLoginBonusRepository) FinOrNilByMasterEventId(ctx context.Context, masterEventId int64) (*MasterLoginBonus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinOrNilByMasterEventId", ctx, masterEventId)
	ret0, _ := ret[0].(*MasterLoginBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FinOrNilByMasterEventId indicates an expected call of FinOrNilByMasterEventId.
func (mr *MockMasterLoginBonusRepositoryMockRecorder) FinOrNilByMasterEventId(ctx, masterEventId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinOrNilByMasterEventId", reflect.TypeOf((*MockMasterLoginBonusRepository)(nil).FinOrNilByMasterEventId), ctx, masterEventId)
}

// Find mocks base method.
func (m *MockMasterLoginBonusRepository) Find(ctx context.Context, id int64) (*MasterLoginBonus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, id)
	ret0, _ := ret[0].(*MasterLoginBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMasterLoginBonusRepositoryMockRecorder) Find(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMasterLoginBonusRepository)(nil).Find), ctx, id)
}

// FindByMasterEventId mocks base method.
func (m *MockMasterLoginBonusRepository) FindByMasterEventId(ctx context.Context, masterEventId int64) (*MasterLoginBonus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMasterEventId", ctx, masterEventId)
	ret0, _ := ret[0].(*MasterLoginBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMasterEventId indicates an expected call of FindByMasterEventId.
func (mr *MockMasterLoginBonusRepositoryMockRecorder) FindByMasterEventId(ctx, masterEventId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMasterEventId", reflect.TypeOf((*MockMasterLoginBonusRepository)(nil).FindByMasterEventId), ctx, masterEventId)
}

// FindList mocks base method.
func (m *MockMasterLoginBonusRepository) FindList(ctx context.Context) (MasterLoginBonuses, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx)
	ret0, _ := ret[0].(MasterLoginBonuses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockMasterLoginBonusRepositoryMockRecorder) FindList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockMasterLoginBonusRepository)(nil).FindList), ctx)
}

// FindListByMasterEventId mocks base method.
func (m *MockMasterLoginBonusRepository) FindListByMasterEventId(ctx context.Context, masterEventId int64) (MasterLoginBonuses, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByMasterEventId", ctx, masterEventId)
	ret0, _ := ret[0].(MasterLoginBonuses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByMasterEventId indicates an expected call of FindListByMasterEventId.
func (mr *MockMasterLoginBonusRepositoryMockRecorder) FindListByMasterEventId(ctx, masterEventId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByMasterEventId", reflect.TypeOf((*MockMasterLoginBonusRepository)(nil).FindListByMasterEventId), ctx, masterEventId)
}

// FindOrNil mocks base method.
func (m *MockMasterLoginBonusRepository) FindOrNil(ctx context.Context, id int64) (*MasterLoginBonus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, id)
	ret0, _ := ret[0].(*MasterLoginBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockMasterLoginBonusRepositoryMockRecorder) FindOrNil(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockMasterLoginBonusRepository)(nil).FindOrNil), ctx, id)
}

// Update mocks base method.
func (m_2 *MockMasterLoginBonusRepository) Update(ctx context.Context, tx *gorm.DB, m *MasterLoginBonus) (*MasterLoginBonus, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*MasterLoginBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMasterLoginBonusRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMasterLoginBonusRepository)(nil).Update), ctx, tx, m)
}