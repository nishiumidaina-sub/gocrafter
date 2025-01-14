// Code generated by MockGen. DO NOT EDIT.
// Source: ./master_login_bonus_schedule_repository.gen.go

// Package masterLoginBonusSchedule is a generated GoMock package.
package masterLoginBonusSchedule

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockMasterLoginBonusScheduleRepository is a mock of MasterLoginBonusScheduleRepository interface.
type MockMasterLoginBonusScheduleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMasterLoginBonusScheduleRepositoryMockRecorder
}

// MockMasterLoginBonusScheduleRepositoryMockRecorder is the mock recorder for MockMasterLoginBonusScheduleRepository.
type MockMasterLoginBonusScheduleRepositoryMockRecorder struct {
	mock *MockMasterLoginBonusScheduleRepository
}

// NewMockMasterLoginBonusScheduleRepository creates a new mock instance.
func NewMockMasterLoginBonusScheduleRepository(ctrl *gomock.Controller) *MockMasterLoginBonusScheduleRepository {
	mock := &MockMasterLoginBonusScheduleRepository{ctrl: ctrl}
	mock.recorder = &MockMasterLoginBonusScheduleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMasterLoginBonusScheduleRepository) EXPECT() *MockMasterLoginBonusScheduleRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockMasterLoginBonusScheduleRepository) Create(ctx context.Context, tx *gorm.DB, m *MasterLoginBonusSchedule) (*MasterLoginBonusSchedule, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*MasterLoginBonusSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMasterLoginBonusScheduleRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMasterLoginBonusScheduleRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockMasterLoginBonusScheduleRepository) CreateList(ctx context.Context, tx *gorm.DB, ms MasterLoginBonusSchedules) (MasterLoginBonusSchedules, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(MasterLoginBonusSchedules)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockMasterLoginBonusScheduleRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockMasterLoginBonusScheduleRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockMasterLoginBonusScheduleRepository) Delete(ctx context.Context, tx *gorm.DB, m *MasterLoginBonusSchedule) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMasterLoginBonusScheduleRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMasterLoginBonusScheduleRepository)(nil).Delete), ctx, tx, m)
}

// Find mocks base method.
func (m *MockMasterLoginBonusScheduleRepository) Find(ctx context.Context, id int64) (*MasterLoginBonusSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, id)
	ret0, _ := ret[0].(*MasterLoginBonusSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMasterLoginBonusScheduleRepositoryMockRecorder) Find(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMasterLoginBonusScheduleRepository)(nil).Find), ctx, id)
}

// FindByMasterLoginBonusId mocks base method.
func (m *MockMasterLoginBonusScheduleRepository) FindByMasterLoginBonusId(ctx context.Context, masterLoginBonusId int64) (*MasterLoginBonusSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMasterLoginBonusId", ctx, masterLoginBonusId)
	ret0, _ := ret[0].(*MasterLoginBonusSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMasterLoginBonusId indicates an expected call of FindByMasterLoginBonusId.
func (mr *MockMasterLoginBonusScheduleRepositoryMockRecorder) FindByMasterLoginBonusId(ctx, masterLoginBonusId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMasterLoginBonusId", reflect.TypeOf((*MockMasterLoginBonusScheduleRepository)(nil).FindByMasterLoginBonusId), ctx, masterLoginBonusId)
}

// FindByMasterLoginBonusIdAndStep mocks base method.
func (m *MockMasterLoginBonusScheduleRepository) FindByMasterLoginBonusIdAndStep(ctx context.Context, masterLoginBonusId int64, step int32) (*MasterLoginBonusSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMasterLoginBonusIdAndStep", ctx, masterLoginBonusId, step)
	ret0, _ := ret[0].(*MasterLoginBonusSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMasterLoginBonusIdAndStep indicates an expected call of FindByMasterLoginBonusIdAndStep.
func (mr *MockMasterLoginBonusScheduleRepositoryMockRecorder) FindByMasterLoginBonusIdAndStep(ctx, masterLoginBonusId, step interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMasterLoginBonusIdAndStep", reflect.TypeOf((*MockMasterLoginBonusScheduleRepository)(nil).FindByMasterLoginBonusIdAndStep), ctx, masterLoginBonusId, step)
}

// FindByStep mocks base method.
func (m *MockMasterLoginBonusScheduleRepository) FindByStep(ctx context.Context, step int32) (*MasterLoginBonusSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByStep", ctx, step)
	ret0, _ := ret[0].(*MasterLoginBonusSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByStep indicates an expected call of FindByStep.
func (mr *MockMasterLoginBonusScheduleRepositoryMockRecorder) FindByStep(ctx, step interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByStep", reflect.TypeOf((*MockMasterLoginBonusScheduleRepository)(nil).FindByStep), ctx, step)
}

// FindList mocks base method.
func (m *MockMasterLoginBonusScheduleRepository) FindList(ctx context.Context) (MasterLoginBonusSchedules, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx)
	ret0, _ := ret[0].(MasterLoginBonusSchedules)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockMasterLoginBonusScheduleRepositoryMockRecorder) FindList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockMasterLoginBonusScheduleRepository)(nil).FindList), ctx)
}

// FindListByMasterLoginBonusId mocks base method.
func (m *MockMasterLoginBonusScheduleRepository) FindListByMasterLoginBonusId(ctx context.Context, masterLoginBonusId int64) (MasterLoginBonusSchedules, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByMasterLoginBonusId", ctx, masterLoginBonusId)
	ret0, _ := ret[0].(MasterLoginBonusSchedules)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByMasterLoginBonusId indicates an expected call of FindListByMasterLoginBonusId.
func (mr *MockMasterLoginBonusScheduleRepositoryMockRecorder) FindListByMasterLoginBonusId(ctx, masterLoginBonusId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByMasterLoginBonusId", reflect.TypeOf((*MockMasterLoginBonusScheduleRepository)(nil).FindListByMasterLoginBonusId), ctx, masterLoginBonusId)
}

// FindListByMasterLoginBonusIdAndStep mocks base method.
func (m *MockMasterLoginBonusScheduleRepository) FindListByMasterLoginBonusIdAndStep(ctx context.Context, masterLoginBonusId int64, step int32) (MasterLoginBonusSchedules, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByMasterLoginBonusIdAndStep", ctx, masterLoginBonusId, step)
	ret0, _ := ret[0].(MasterLoginBonusSchedules)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByMasterLoginBonusIdAndStep indicates an expected call of FindListByMasterLoginBonusIdAndStep.
func (mr *MockMasterLoginBonusScheduleRepositoryMockRecorder) FindListByMasterLoginBonusIdAndStep(ctx, masterLoginBonusId, step interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByMasterLoginBonusIdAndStep", reflect.TypeOf((*MockMasterLoginBonusScheduleRepository)(nil).FindListByMasterLoginBonusIdAndStep), ctx, masterLoginBonusId, step)
}

// FindListByStep mocks base method.
func (m *MockMasterLoginBonusScheduleRepository) FindListByStep(ctx context.Context, step int32) (MasterLoginBonusSchedules, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByStep", ctx, step)
	ret0, _ := ret[0].(MasterLoginBonusSchedules)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByStep indicates an expected call of FindListByStep.
func (mr *MockMasterLoginBonusScheduleRepositoryMockRecorder) FindListByStep(ctx, step interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByStep", reflect.TypeOf((*MockMasterLoginBonusScheduleRepository)(nil).FindListByStep), ctx, step)
}

// FindOrNil mocks base method.
func (m *MockMasterLoginBonusScheduleRepository) FindOrNil(ctx context.Context, id int64) (*MasterLoginBonusSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, id)
	ret0, _ := ret[0].(*MasterLoginBonusSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockMasterLoginBonusScheduleRepositoryMockRecorder) FindOrNil(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockMasterLoginBonusScheduleRepository)(nil).FindOrNil), ctx, id)
}

// FindOrNilByMasterLoginBonusId mocks base method.
func (m *MockMasterLoginBonusScheduleRepository) FindOrNilByMasterLoginBonusId(ctx context.Context, masterLoginBonusId int64) (*MasterLoginBonusSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByMasterLoginBonusId", ctx, masterLoginBonusId)
	ret0, _ := ret[0].(*MasterLoginBonusSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByMasterLoginBonusId indicates an expected call of FindOrNilByMasterLoginBonusId.
func (mr *MockMasterLoginBonusScheduleRepositoryMockRecorder) FindOrNilByMasterLoginBonusId(ctx, masterLoginBonusId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByMasterLoginBonusId", reflect.TypeOf((*MockMasterLoginBonusScheduleRepository)(nil).FindOrNilByMasterLoginBonusId), ctx, masterLoginBonusId)
}

// FindOrNilByMasterLoginBonusIdAndStep mocks base method.
func (m *MockMasterLoginBonusScheduleRepository) FindOrNilByMasterLoginBonusIdAndStep(ctx context.Context, masterLoginBonusId int64, step int32) (*MasterLoginBonusSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByMasterLoginBonusIdAndStep", ctx, masterLoginBonusId, step)
	ret0, _ := ret[0].(*MasterLoginBonusSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByMasterLoginBonusIdAndStep indicates an expected call of FindOrNilByMasterLoginBonusIdAndStep.
func (mr *MockMasterLoginBonusScheduleRepositoryMockRecorder) FindOrNilByMasterLoginBonusIdAndStep(ctx, masterLoginBonusId, step interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByMasterLoginBonusIdAndStep", reflect.TypeOf((*MockMasterLoginBonusScheduleRepository)(nil).FindOrNilByMasterLoginBonusIdAndStep), ctx, masterLoginBonusId, step)
}

// FindOrNilByStep mocks base method.
func (m *MockMasterLoginBonusScheduleRepository) FindOrNilByStep(ctx context.Context, step int32) (*MasterLoginBonusSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByStep", ctx, step)
	ret0, _ := ret[0].(*MasterLoginBonusSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByStep indicates an expected call of FindOrNilByStep.
func (mr *MockMasterLoginBonusScheduleRepositoryMockRecorder) FindOrNilByStep(ctx, step interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByStep", reflect.TypeOf((*MockMasterLoginBonusScheduleRepository)(nil).FindOrNilByStep), ctx, step)
}

// Update mocks base method.
func (m_2 *MockMasterLoginBonusScheduleRepository) Update(ctx context.Context, tx *gorm.DB, m *MasterLoginBonusSchedule) (*MasterLoginBonusSchedule, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*MasterLoginBonusSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMasterLoginBonusScheduleRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMasterLoginBonusScheduleRepository)(nil).Update), ctx, tx, m)
}
