// Code generated by MockGen. DO NOT EDIT.
// Source: meeting_app/internal/app/service (interfaces: IAppointment)

// Package service is a generated GoMock package.
package service

import (
	api_structures "meeting_app/internal/app/structures"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockIAppointment is a mock of IAppointment interface.
type MockIAppointment struct {
	ctrl     *gomock.Controller
	recorder *MockIAppointmentMockRecorder
}

// MockIAppointmentMockRecorder is the mock recorder for MockIAppointment.
type MockIAppointmentMockRecorder struct {
	mock *MockIAppointment
}

// NewMockIAppointment creates a new mock instance.
func NewMockIAppointment(ctrl *gomock.Controller) *MockIAppointment {
	mock := &MockIAppointment{ctrl: ctrl}
	mock.recorder = &MockIAppointmentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAppointment) EXPECT() *MockIAppointmentMockRecorder {
	return m.recorder
}

// CreateAppointment mocks base method.
func (m *MockIAppointment) CreateAppointment(arg0 api_structures.AppointmentForm) ([]api_structures.AppointmentForm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAppointment", arg0)
	ret0, _ := ret[0].([]api_structures.AppointmentForm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAppointment indicates an expected call of CreateAppointment.
func (mr *MockIAppointmentMockRecorder) CreateAppointment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAppointment", reflect.TypeOf((*MockIAppointment)(nil).CreateAppointment), arg0)
}

// CreateAppointmentMultiple mocks base method.
func (m *MockIAppointment) CreateAppointmentMultiple(arg0 []api_structures.AppointmentForm) ([]api_structures.AppointmentForm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAppointmentMultiple", arg0)
	ret0, _ := ret[0].([]api_structures.AppointmentForm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAppointmentMultiple indicates an expected call of CreateAppointmentMultiple.
func (mr *MockIAppointmentMockRecorder) CreateAppointmentMultiple(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAppointmentMultiple", reflect.TypeOf((*MockIAppointment)(nil).CreateAppointmentMultiple), arg0)
}

// DeleteAppointment mocks base method.
func (m *MockIAppointment) DeleteAppointment(arg0 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAppointment", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAppointment indicates an expected call of DeleteAppointment.
func (mr *MockIAppointmentMockRecorder) DeleteAppointment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAppointment", reflect.TypeOf((*MockIAppointment)(nil).DeleteAppointment), arg0)
}

// GetAppointment mocks base method.
func (m *MockIAppointment) GetAppointment(arg0 api_structures.AppointmentFilter) ([]api_structures.Appointment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppointment", arg0)
	ret0, _ := ret[0].([]api_structures.Appointment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppointment indicates an expected call of GetAppointment.
func (mr *MockIAppointmentMockRecorder) GetAppointment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppointment", reflect.TypeOf((*MockIAppointment)(nil).GetAppointment), arg0)
}

// GetAppointmentById mocks base method.
func (m *MockIAppointment) GetAppointmentById(arg0 uuid.UUID) (api_structures.Appointment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppointmentById", arg0)
	ret0, _ := ret[0].(api_structures.Appointment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppointmentById indicates an expected call of GetAppointmentById.
func (mr *MockIAppointmentMockRecorder) GetAppointmentById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppointmentById", reflect.TypeOf((*MockIAppointment)(nil).GetAppointmentById), arg0)
}

// UpdateAppointment mocks base method.
func (m *MockIAppointment) UpdateAppointment(arg0 uuid.UUID, arg1 api_structures.AppointmentEdit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAppointment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAppointment indicates an expected call of UpdateAppointment.
func (mr *MockIAppointmentMockRecorder) UpdateAppointment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAppointment", reflect.TypeOf((*MockIAppointment)(nil).UpdateAppointment), arg0, arg1)
}

// UpdateAppointmentMultiple mocks base method.
func (m *MockIAppointment) UpdateAppointmentMultiple(arg0 []api_structures.AppointmentEdit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAppointmentMultiple", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAppointmentMultiple indicates an expected call of UpdateAppointmentMultiple.
func (mr *MockIAppointmentMockRecorder) UpdateAppointmentMultiple(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAppointmentMultiple", reflect.TypeOf((*MockIAppointment)(nil).UpdateAppointmentMultiple), arg0)
}