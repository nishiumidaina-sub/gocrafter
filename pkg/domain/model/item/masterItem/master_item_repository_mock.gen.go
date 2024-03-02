// Code generated by MockGen. DO NOT EDIT.
// Source: ./master_item_repository.gen.go

// Package masterItem is a generated GoMock package.
package masterItem

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockMasterItemRepository is a mock of MasterItemRepository interface.
type MockMasterItemRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMasterItemRepositoryMockRecorder
}

// MockMasterItemRepositoryMockRecorder is the mock recorder for MockMasterItemRepository.
type MockMasterItemRepositoryMockRecorder struct {
	mock *MockMasterItemRepository
}

// NewMockMasterItemRepository creates a new mock instance.
func NewMockMasterItemRepository(ctrl *gomock.Controller) *MockMasterItemRepository {
	mock := &MockMasterItemRepository{ctrl: ctrl}
	mock.recorder = &MockMasterItemRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMasterItemRepository) EXPECT() *MockMasterItemRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockMasterItemRepository) Create(ctx context.Context, tx *gorm.DB, m *MasterItem) (*MasterItem, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*MasterItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMasterItemRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMasterItemRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockMasterItemRepository) CreateList(ctx context.Context, tx *gorm.DB, ms MasterItems) (MasterItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(MasterItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockMasterItemRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockMasterItemRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockMasterItemRepository) Delete(ctx context.Context, tx *gorm.DB, m *MasterItem) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMasterItemRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMasterItemRepository)(nil).Delete), ctx, tx, m)
}

// FinOrNilByName mocks base method.
func (m *MockMasterItemRepository) FinOrNilByName(ctx context.Context, name string) (*MasterItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinOrNilByName", ctx, name)
	ret0, _ := ret[0].(*MasterItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FinOrNilByName indicates an expected call of FinOrNilByName.
func (mr *MockMasterItemRepositoryMockRecorder) FinOrNilByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinOrNilByName", reflect.TypeOf((*MockMasterItemRepository)(nil).FinOrNilByName), ctx, name)
}

// Find mocks base method.
func (m *MockMasterItemRepository) Find(ctx context.Context, id int64) (*MasterItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, id)
	ret0, _ := ret[0].(*MasterItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMasterItemRepositoryMockRecorder) Find(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMasterItemRepository)(nil).Find), ctx, id)
}

// FindByName mocks base method.
func (m *MockMasterItemRepository) FindByName(ctx context.Context, name string) (*MasterItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", ctx, name)
	ret0, _ := ret[0].(*MasterItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockMasterItemRepositoryMockRecorder) FindByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockMasterItemRepository)(nil).FindByName), ctx, name)
}

// FindList mocks base method.
func (m *MockMasterItemRepository) FindList(ctx context.Context) (MasterItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx)
	ret0, _ := ret[0].(MasterItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockMasterItemRepositoryMockRecorder) FindList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockMasterItemRepository)(nil).FindList), ctx)
}

// FindListByName mocks base method.
func (m *MockMasterItemRepository) FindListByName(ctx context.Context, name string) (MasterItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByName", ctx, name)
	ret0, _ := ret[0].(MasterItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByName indicates an expected call of FindListByName.
func (mr *MockMasterItemRepositoryMockRecorder) FindListByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByName", reflect.TypeOf((*MockMasterItemRepository)(nil).FindListByName), ctx, name)
}

// FindOrNil mocks base method.
func (m *MockMasterItemRepository) FindOrNil(ctx context.Context, id int64) (*MasterItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, id)
	ret0, _ := ret[0].(*MasterItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockMasterItemRepositoryMockRecorder) FindOrNil(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockMasterItemRepository)(nil).FindOrNil), ctx, id)
}

// Update mocks base method.
func (m_2 *MockMasterItemRepository) Update(ctx context.Context, tx *gorm.DB, m *MasterItem) (*MasterItem, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*MasterItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMasterItemRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMasterItemRepository)(nil).Update), ctx, tx, m)
}