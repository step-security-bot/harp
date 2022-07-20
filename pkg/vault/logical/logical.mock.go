// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/zntrio/harp/v1/pkg/vault/logical (interfaces: Logical)

// Package logical is a generated GoMock package.
package logical

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/hashicorp/vault/api"
)

// MockLogical is a mock of Logical interface.
type MockLogical struct {
	ctrl     *gomock.Controller
	recorder *MockLogicalMockRecorder
}

// MockLogicalMockRecorder is the mock recorder for MockLogical.
type MockLogicalMockRecorder struct {
	mock *MockLogical
}

// NewMockLogical creates a new mock instance.
func NewMockLogical(ctrl *gomock.Controller) *MockLogical {
	mock := &MockLogical{ctrl: ctrl}
	mock.recorder = &MockLogicalMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogical) EXPECT() *MockLogicalMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockLogical) Delete(arg0 string) (*api.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(*api.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockLogicalMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockLogical)(nil).Delete), arg0)
}

// DeleteWithData mocks base method.
func (m *MockLogical) DeleteWithData(arg0 string, arg1 map[string][]string) (*api.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWithData", arg0, arg1)
	ret0, _ := ret[0].(*api.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteWithData indicates an expected call of DeleteWithData.
func (mr *MockLogicalMockRecorder) DeleteWithData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWithData", reflect.TypeOf((*MockLogical)(nil).DeleteWithData), arg0, arg1)
}

// List mocks base method.
func (m *MockLogical) List(arg0 string) (*api.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(*api.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockLogicalMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockLogical)(nil).List), arg0)
}

// Read mocks base method.
func (m *MockLogical) Read(arg0 string) (*api.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(*api.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockLogicalMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockLogical)(nil).Read), arg0)
}

// ReadWithData mocks base method.
func (m *MockLogical) ReadWithData(arg0 string, arg1 map[string][]string) (*api.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadWithData", arg0, arg1)
	ret0, _ := ret[0].(*api.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadWithData indicates an expected call of ReadWithData.
func (mr *MockLogicalMockRecorder) ReadWithData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadWithData", reflect.TypeOf((*MockLogical)(nil).ReadWithData), arg0, arg1)
}

// Unwrap mocks base method.
func (m *MockLogical) Unwrap(arg0 string) (*api.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unwrap", arg0)
	ret0, _ := ret[0].(*api.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unwrap indicates an expected call of Unwrap.
func (mr *MockLogicalMockRecorder) Unwrap(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unwrap", reflect.TypeOf((*MockLogical)(nil).Unwrap), arg0)
}

// Write mocks base method.
func (m *MockLogical) Write(arg0 string, arg1 map[string]interface{}) (*api.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(*api.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockLogicalMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockLogical)(nil).Write), arg0, arg1)
}

// WriteBytes mocks base method.
func (m *MockLogical) WriteBytes(arg0 string, arg1 []byte) (*api.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteBytes", arg0, arg1)
	ret0, _ := ret[0].(*api.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteBytes indicates an expected call of WriteBytes.
func (mr *MockLogicalMockRecorder) WriteBytes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteBytes", reflect.TypeOf((*MockLogical)(nil).WriteBytes), arg0, arg1)
}
