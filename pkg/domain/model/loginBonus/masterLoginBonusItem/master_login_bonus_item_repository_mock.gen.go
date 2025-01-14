// Code generated by MockGen. DO NOT EDIT.
// Source: ./master_login_bonus_item_repository.gen.go

// Package masterLoginBonusItem is a generated GoMock package.
package masterLoginBonusItem

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockMasterLoginBonusItemRepository is a mock of MasterLoginBonusItemRepository interface.
type MockMasterLoginBonusItemRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMasterLoginBonusItemRepositoryMockRecorder
}

// MockMasterLoginBonusItemRepositoryMockRecorder is the mock recorder for MockMasterLoginBonusItemRepository.
type MockMasterLoginBonusItemRepositoryMockRecorder struct {
	mock *MockMasterLoginBonusItemRepository
}

// NewMockMasterLoginBonusItemRepository creates a new mock instance.
func NewMockMasterLoginBonusItemRepository(ctrl *gomock.Controller) *MockMasterLoginBonusItemRepository {
	mock := &MockMasterLoginBonusItemRepository{ctrl: ctrl}
	mock.recorder = &MockMasterLoginBonusItemRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMasterLoginBonusItemRepository) EXPECT() *MockMasterLoginBonusItemRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockMasterLoginBonusItemRepository) Create(ctx context.Context, tx *gorm.DB, m *MasterLoginBonusItem) (*MasterLoginBonusItem, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*MasterLoginBonusItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMasterLoginBonusItemRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMasterLoginBonusItemRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockMasterLoginBonusItemRepository) CreateList(ctx context.Context, tx *gorm.DB, ms MasterLoginBonusItems) (MasterLoginBonusItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(MasterLoginBonusItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockMasterLoginBonusItemRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockMasterLoginBonusItemRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockMasterLoginBonusItemRepository) Delete(ctx context.Context, tx *gorm.DB, m *MasterLoginBonusItem) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMasterLoginBonusItemRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMasterLoginBonusItemRepository)(nil).Delete), ctx, tx, m)
}

// Find mocks base method.
func (m *MockMasterLoginBonusItemRepository) Find(ctx context.Context, id int64) (*MasterLoginBonusItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, id)
	ret0, _ := ret[0].(*MasterLoginBonusItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMasterLoginBonusItemRepositoryMockRecorder) Find(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMasterLoginBonusItemRepository)(nil).Find), ctx, id)
}

// FindByMasterItemId mocks base method.
func (m *MockMasterLoginBonusItemRepository) FindByMasterItemId(ctx context.Context, masterItemId int64) (*MasterLoginBonusItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMasterItemId", ctx, masterItemId)
	ret0, _ := ret[0].(*MasterLoginBonusItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMasterItemId indicates an expected call of FindByMasterItemId.
func (mr *MockMasterLoginBonusItemRepositoryMockRecorder) FindByMasterItemId(ctx, masterItemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMasterItemId", reflect.TypeOf((*MockMasterLoginBonusItemRepository)(nil).FindByMasterItemId), ctx, masterItemId)
}

// FindByMasterLoginBonusScheduleId mocks base method.
func (m *MockMasterLoginBonusItemRepository) FindByMasterLoginBonusScheduleId(ctx context.Context, masterLoginBonusScheduleId int64) (*MasterLoginBonusItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMasterLoginBonusScheduleId", ctx, masterLoginBonusScheduleId)
	ret0, _ := ret[0].(*MasterLoginBonusItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMasterLoginBonusScheduleId indicates an expected call of FindByMasterLoginBonusScheduleId.
func (mr *MockMasterLoginBonusItemRepositoryMockRecorder) FindByMasterLoginBonusScheduleId(ctx, masterLoginBonusScheduleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMasterLoginBonusScheduleId", reflect.TypeOf((*MockMasterLoginBonusItemRepository)(nil).FindByMasterLoginBonusScheduleId), ctx, masterLoginBonusScheduleId)
}

// FindByMasterLoginBonusScheduleIdAndMasterItemId mocks base method.
func (m *MockMasterLoginBonusItemRepository) FindByMasterLoginBonusScheduleIdAndMasterItemId(ctx context.Context, masterLoginBonusScheduleId, masterItemId int64) (*MasterLoginBonusItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMasterLoginBonusScheduleIdAndMasterItemId", ctx, masterLoginBonusScheduleId, masterItemId)
	ret0, _ := ret[0].(*MasterLoginBonusItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMasterLoginBonusScheduleIdAndMasterItemId indicates an expected call of FindByMasterLoginBonusScheduleIdAndMasterItemId.
func (mr *MockMasterLoginBonusItemRepositoryMockRecorder) FindByMasterLoginBonusScheduleIdAndMasterItemId(ctx, masterLoginBonusScheduleId, masterItemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMasterLoginBonusScheduleIdAndMasterItemId", reflect.TypeOf((*MockMasterLoginBonusItemRepository)(nil).FindByMasterLoginBonusScheduleIdAndMasterItemId), ctx, masterLoginBonusScheduleId, masterItemId)
}

// FindList mocks base method.
func (m *MockMasterLoginBonusItemRepository) FindList(ctx context.Context) (MasterLoginBonusItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx)
	ret0, _ := ret[0].(MasterLoginBonusItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockMasterLoginBonusItemRepositoryMockRecorder) FindList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockMasterLoginBonusItemRepository)(nil).FindList), ctx)
}

// FindListByMasterItemId mocks base method.
func (m *MockMasterLoginBonusItemRepository) FindListByMasterItemId(ctx context.Context, masterItemId int64) (MasterLoginBonusItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByMasterItemId", ctx, masterItemId)
	ret0, _ := ret[0].(MasterLoginBonusItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByMasterItemId indicates an expected call of FindListByMasterItemId.
func (mr *MockMasterLoginBonusItemRepositoryMockRecorder) FindListByMasterItemId(ctx, masterItemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByMasterItemId", reflect.TypeOf((*MockMasterLoginBonusItemRepository)(nil).FindListByMasterItemId), ctx, masterItemId)
}

// FindListByMasterLoginBonusScheduleId mocks base method.
func (m *MockMasterLoginBonusItemRepository) FindListByMasterLoginBonusScheduleId(ctx context.Context, masterLoginBonusScheduleId int64) (MasterLoginBonusItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByMasterLoginBonusScheduleId", ctx, masterLoginBonusScheduleId)
	ret0, _ := ret[0].(MasterLoginBonusItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByMasterLoginBonusScheduleId indicates an expected call of FindListByMasterLoginBonusScheduleId.
func (mr *MockMasterLoginBonusItemRepositoryMockRecorder) FindListByMasterLoginBonusScheduleId(ctx, masterLoginBonusScheduleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByMasterLoginBonusScheduleId", reflect.TypeOf((*MockMasterLoginBonusItemRepository)(nil).FindListByMasterLoginBonusScheduleId), ctx, masterLoginBonusScheduleId)
}

// FindListByMasterLoginBonusScheduleIdAndMasterItemId mocks base method.
func (m *MockMasterLoginBonusItemRepository) FindListByMasterLoginBonusScheduleIdAndMasterItemId(ctx context.Context, masterLoginBonusScheduleId, masterItemId int64) (MasterLoginBonusItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByMasterLoginBonusScheduleIdAndMasterItemId", ctx, masterLoginBonusScheduleId, masterItemId)
	ret0, _ := ret[0].(MasterLoginBonusItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByMasterLoginBonusScheduleIdAndMasterItemId indicates an expected call of FindListByMasterLoginBonusScheduleIdAndMasterItemId.
func (mr *MockMasterLoginBonusItemRepositoryMockRecorder) FindListByMasterLoginBonusScheduleIdAndMasterItemId(ctx, masterLoginBonusScheduleId, masterItemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByMasterLoginBonusScheduleIdAndMasterItemId", reflect.TypeOf((*MockMasterLoginBonusItemRepository)(nil).FindListByMasterLoginBonusScheduleIdAndMasterItemId), ctx, masterLoginBonusScheduleId, masterItemId)
}

// FindOrNil mocks base method.
func (m *MockMasterLoginBonusItemRepository) FindOrNil(ctx context.Context, id int64) (*MasterLoginBonusItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, id)
	ret0, _ := ret[0].(*MasterLoginBonusItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockMasterLoginBonusItemRepositoryMockRecorder) FindOrNil(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockMasterLoginBonusItemRepository)(nil).FindOrNil), ctx, id)
}

// FindOrNilByMasterItemId mocks base method.
func (m *MockMasterLoginBonusItemRepository) FindOrNilByMasterItemId(ctx context.Context, masterItemId int64) (*MasterLoginBonusItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByMasterItemId", ctx, masterItemId)
	ret0, _ := ret[0].(*MasterLoginBonusItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByMasterItemId indicates an expected call of FindOrNilByMasterItemId.
func (mr *MockMasterLoginBonusItemRepositoryMockRecorder) FindOrNilByMasterItemId(ctx, masterItemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByMasterItemId", reflect.TypeOf((*MockMasterLoginBonusItemRepository)(nil).FindOrNilByMasterItemId), ctx, masterItemId)
}

// FindOrNilByMasterLoginBonusScheduleId mocks base method.
func (m *MockMasterLoginBonusItemRepository) FindOrNilByMasterLoginBonusScheduleId(ctx context.Context, masterLoginBonusScheduleId int64) (*MasterLoginBonusItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByMasterLoginBonusScheduleId", ctx, masterLoginBonusScheduleId)
	ret0, _ := ret[0].(*MasterLoginBonusItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByMasterLoginBonusScheduleId indicates an expected call of FindOrNilByMasterLoginBonusScheduleId.
func (mr *MockMasterLoginBonusItemRepositoryMockRecorder) FindOrNilByMasterLoginBonusScheduleId(ctx, masterLoginBonusScheduleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByMasterLoginBonusScheduleId", reflect.TypeOf((*MockMasterLoginBonusItemRepository)(nil).FindOrNilByMasterLoginBonusScheduleId), ctx, masterLoginBonusScheduleId)
}

// FindOrNilByMasterLoginBonusScheduleIdAndMasterItemId mocks base method.
func (m *MockMasterLoginBonusItemRepository) FindOrNilByMasterLoginBonusScheduleIdAndMasterItemId(ctx context.Context, masterLoginBonusScheduleId, masterItemId int64) (*MasterLoginBonusItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByMasterLoginBonusScheduleIdAndMasterItemId", ctx, masterLoginBonusScheduleId, masterItemId)
	ret0, _ := ret[0].(*MasterLoginBonusItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByMasterLoginBonusScheduleIdAndMasterItemId indicates an expected call of FindOrNilByMasterLoginBonusScheduleIdAndMasterItemId.
func (mr *MockMasterLoginBonusItemRepositoryMockRecorder) FindOrNilByMasterLoginBonusScheduleIdAndMasterItemId(ctx, masterLoginBonusScheduleId, masterItemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByMasterLoginBonusScheduleIdAndMasterItemId", reflect.TypeOf((*MockMasterLoginBonusItemRepository)(nil).FindOrNilByMasterLoginBonusScheduleIdAndMasterItemId), ctx, masterLoginBonusScheduleId, masterItemId)
}

// Update mocks base method.
func (m_2 *MockMasterLoginBonusItemRepository) Update(ctx context.Context, tx *gorm.DB, m *MasterLoginBonusItem) (*MasterLoginBonusItem, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*MasterLoginBonusItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMasterLoginBonusItemRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMasterLoginBonusItemRepository)(nil).Update), ctx, tx, m)
}
