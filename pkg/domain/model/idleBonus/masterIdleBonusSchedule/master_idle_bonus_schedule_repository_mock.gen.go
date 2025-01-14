// Code generated by MockGen. DO NOT EDIT.
// Source: ./master_idle_bonus_schedule_repository.gen.go

// Package masterIdleBonusSchedule is a generated GoMock package.
package masterIdleBonusSchedule

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockMasterIdleBonusScheduleRepository is a mock of MasterIdleBonusScheduleRepository interface.
type MockMasterIdleBonusScheduleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMasterIdleBonusScheduleRepositoryMockRecorder
}

// MockMasterIdleBonusScheduleRepositoryMockRecorder is the mock recorder for MockMasterIdleBonusScheduleRepository.
type MockMasterIdleBonusScheduleRepositoryMockRecorder struct {
	mock *MockMasterIdleBonusScheduleRepository
}

// NewMockMasterIdleBonusScheduleRepository creates a new mock instance.
func NewMockMasterIdleBonusScheduleRepository(ctrl *gomock.Controller) *MockMasterIdleBonusScheduleRepository {
	mock := &MockMasterIdleBonusScheduleRepository{ctrl: ctrl}
	mock.recorder = &MockMasterIdleBonusScheduleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMasterIdleBonusScheduleRepository) EXPECT() *MockMasterIdleBonusScheduleRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockMasterIdleBonusScheduleRepository) Create(ctx context.Context, tx *gorm.DB, m *MasterIdleBonusSchedule) (*MasterIdleBonusSchedule, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*MasterIdleBonusSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMasterIdleBonusScheduleRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMasterIdleBonusScheduleRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockMasterIdleBonusScheduleRepository) CreateList(ctx context.Context, tx *gorm.DB, ms MasterIdleBonusSchedules) (MasterIdleBonusSchedules, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(MasterIdleBonusSchedules)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockMasterIdleBonusScheduleRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockMasterIdleBonusScheduleRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockMasterIdleBonusScheduleRepository) Delete(ctx context.Context, tx *gorm.DB, m *MasterIdleBonusSchedule) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMasterIdleBonusScheduleRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMasterIdleBonusScheduleRepository)(nil).Delete), ctx, tx, m)
}

// Find mocks base method.
func (m *MockMasterIdleBonusScheduleRepository) Find(ctx context.Context, id int64) (*MasterIdleBonusSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, id)
	ret0, _ := ret[0].(*MasterIdleBonusSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMasterIdleBonusScheduleRepositoryMockRecorder) Find(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMasterIdleBonusScheduleRepository)(nil).Find), ctx, id)
}

// FindByMasterIdleBonusId mocks base method.
func (m *MockMasterIdleBonusScheduleRepository) FindByMasterIdleBonusId(ctx context.Context, masterIdleBonusId int64) (*MasterIdleBonusSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMasterIdleBonusId", ctx, masterIdleBonusId)
	ret0, _ := ret[0].(*MasterIdleBonusSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMasterIdleBonusId indicates an expected call of FindByMasterIdleBonusId.
func (mr *MockMasterIdleBonusScheduleRepositoryMockRecorder) FindByMasterIdleBonusId(ctx, masterIdleBonusId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMasterIdleBonusId", reflect.TypeOf((*MockMasterIdleBonusScheduleRepository)(nil).FindByMasterIdleBonusId), ctx, masterIdleBonusId)
}

// FindByMasterIdleBonusIdAndStep mocks base method.
func (m *MockMasterIdleBonusScheduleRepository) FindByMasterIdleBonusIdAndStep(ctx context.Context, masterIdleBonusId int64, step int32) (*MasterIdleBonusSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMasterIdleBonusIdAndStep", ctx, masterIdleBonusId, step)
	ret0, _ := ret[0].(*MasterIdleBonusSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMasterIdleBonusIdAndStep indicates an expected call of FindByMasterIdleBonusIdAndStep.
func (mr *MockMasterIdleBonusScheduleRepositoryMockRecorder) FindByMasterIdleBonusIdAndStep(ctx, masterIdleBonusId, step interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMasterIdleBonusIdAndStep", reflect.TypeOf((*MockMasterIdleBonusScheduleRepository)(nil).FindByMasterIdleBonusIdAndStep), ctx, masterIdleBonusId, step)
}

// FindByStep mocks base method.
func (m *MockMasterIdleBonusScheduleRepository) FindByStep(ctx context.Context, step int32) (*MasterIdleBonusSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByStep", ctx, step)
	ret0, _ := ret[0].(*MasterIdleBonusSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByStep indicates an expected call of FindByStep.
func (mr *MockMasterIdleBonusScheduleRepositoryMockRecorder) FindByStep(ctx, step interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByStep", reflect.TypeOf((*MockMasterIdleBonusScheduleRepository)(nil).FindByStep), ctx, step)
}

// FindList mocks base method.
func (m *MockMasterIdleBonusScheduleRepository) FindList(ctx context.Context) (MasterIdleBonusSchedules, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx)
	ret0, _ := ret[0].(MasterIdleBonusSchedules)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockMasterIdleBonusScheduleRepositoryMockRecorder) FindList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockMasterIdleBonusScheduleRepository)(nil).FindList), ctx)
}

// FindListByMasterIdleBonusId mocks base method.
func (m *MockMasterIdleBonusScheduleRepository) FindListByMasterIdleBonusId(ctx context.Context, masterIdleBonusId int64) (MasterIdleBonusSchedules, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByMasterIdleBonusId", ctx, masterIdleBonusId)
	ret0, _ := ret[0].(MasterIdleBonusSchedules)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByMasterIdleBonusId indicates an expected call of FindListByMasterIdleBonusId.
func (mr *MockMasterIdleBonusScheduleRepositoryMockRecorder) FindListByMasterIdleBonusId(ctx, masterIdleBonusId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByMasterIdleBonusId", reflect.TypeOf((*MockMasterIdleBonusScheduleRepository)(nil).FindListByMasterIdleBonusId), ctx, masterIdleBonusId)
}

// FindListByMasterIdleBonusIdAndStep mocks base method.
func (m *MockMasterIdleBonusScheduleRepository) FindListByMasterIdleBonusIdAndStep(ctx context.Context, masterIdleBonusId int64, step int32) (MasterIdleBonusSchedules, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByMasterIdleBonusIdAndStep", ctx, masterIdleBonusId, step)
	ret0, _ := ret[0].(MasterIdleBonusSchedules)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByMasterIdleBonusIdAndStep indicates an expected call of FindListByMasterIdleBonusIdAndStep.
func (mr *MockMasterIdleBonusScheduleRepositoryMockRecorder) FindListByMasterIdleBonusIdAndStep(ctx, masterIdleBonusId, step interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByMasterIdleBonusIdAndStep", reflect.TypeOf((*MockMasterIdleBonusScheduleRepository)(nil).FindListByMasterIdleBonusIdAndStep), ctx, masterIdleBonusId, step)
}

// FindListByStep mocks base method.
func (m *MockMasterIdleBonusScheduleRepository) FindListByStep(ctx context.Context, step int32) (MasterIdleBonusSchedules, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByStep", ctx, step)
	ret0, _ := ret[0].(MasterIdleBonusSchedules)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByStep indicates an expected call of FindListByStep.
func (mr *MockMasterIdleBonusScheduleRepositoryMockRecorder) FindListByStep(ctx, step interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByStep", reflect.TypeOf((*MockMasterIdleBonusScheduleRepository)(nil).FindListByStep), ctx, step)
}

// FindOrNil mocks base method.
func (m *MockMasterIdleBonusScheduleRepository) FindOrNil(ctx context.Context, id int64) (*MasterIdleBonusSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, id)
	ret0, _ := ret[0].(*MasterIdleBonusSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockMasterIdleBonusScheduleRepositoryMockRecorder) FindOrNil(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockMasterIdleBonusScheduleRepository)(nil).FindOrNil), ctx, id)
}

// FindOrNilByMasterIdleBonusId mocks base method.
func (m *MockMasterIdleBonusScheduleRepository) FindOrNilByMasterIdleBonusId(ctx context.Context, masterIdleBonusId int64) (*MasterIdleBonusSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByMasterIdleBonusId", ctx, masterIdleBonusId)
	ret0, _ := ret[0].(*MasterIdleBonusSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByMasterIdleBonusId indicates an expected call of FindOrNilByMasterIdleBonusId.
func (mr *MockMasterIdleBonusScheduleRepositoryMockRecorder) FindOrNilByMasterIdleBonusId(ctx, masterIdleBonusId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByMasterIdleBonusId", reflect.TypeOf((*MockMasterIdleBonusScheduleRepository)(nil).FindOrNilByMasterIdleBonusId), ctx, masterIdleBonusId)
}

// FindOrNilByMasterIdleBonusIdAndStep mocks base method.
func (m *MockMasterIdleBonusScheduleRepository) FindOrNilByMasterIdleBonusIdAndStep(ctx context.Context, masterIdleBonusId int64, step int32) (*MasterIdleBonusSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByMasterIdleBonusIdAndStep", ctx, masterIdleBonusId, step)
	ret0, _ := ret[0].(*MasterIdleBonusSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByMasterIdleBonusIdAndStep indicates an expected call of FindOrNilByMasterIdleBonusIdAndStep.
func (mr *MockMasterIdleBonusScheduleRepositoryMockRecorder) FindOrNilByMasterIdleBonusIdAndStep(ctx, masterIdleBonusId, step interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByMasterIdleBonusIdAndStep", reflect.TypeOf((*MockMasterIdleBonusScheduleRepository)(nil).FindOrNilByMasterIdleBonusIdAndStep), ctx, masterIdleBonusId, step)
}

// FindOrNilByStep mocks base method.
func (m *MockMasterIdleBonusScheduleRepository) FindOrNilByStep(ctx context.Context, step int32) (*MasterIdleBonusSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByStep", ctx, step)
	ret0, _ := ret[0].(*MasterIdleBonusSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByStep indicates an expected call of FindOrNilByStep.
func (mr *MockMasterIdleBonusScheduleRepositoryMockRecorder) FindOrNilByStep(ctx, step interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByStep", reflect.TypeOf((*MockMasterIdleBonusScheduleRepository)(nil).FindOrNilByStep), ctx, step)
}

// Update mocks base method.
func (m_2 *MockMasterIdleBonusScheduleRepository) Update(ctx context.Context, tx *gorm.DB, m *MasterIdleBonusSchedule) (*MasterIdleBonusSchedule, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*MasterIdleBonusSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMasterIdleBonusScheduleRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMasterIdleBonusScheduleRepository)(nil).Update), ctx, tx, m)
}
