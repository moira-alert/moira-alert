// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/moira-alert/moira (interfaces: Logger)

// Package mock_moira_alert is a generated GoMock package.
package mock_moira_alert

import (
	gomock "github.com/golang/mock/gomock"
	"github.com/moira-alert/moira"
	reflect "reflect"
)

// MockLogger is a mock of Logger interface
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Clone mocks base method
func (m *MockLogger) Clone() moira.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(moira.Logger)
	return ret0
}

// Clone indicates an expected call of Clone
func (mr *MockLoggerMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockLogger)(nil).Clone))
}

// Debug mocks base method
func (m *MockLogger) Debug(arg0 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug
func (mr *MockLoggerMockRecorder) Debug(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockLogger)(nil).Debug), arg0...)
}

// Debugf mocks base method
func (m *MockLogger) Debugf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf
func (mr *MockLoggerMockRecorder) Debugf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockLogger)(nil).Debugf), varargs...)
}

// Error mocks base method
func (m *MockLogger) Error(arg0 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error
func (mr *MockLoggerMockRecorder) Error(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogger)(nil).Error), arg0...)
}

// Errorf mocks base method
func (m *MockLogger) Errorf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf
func (mr *MockLoggerMockRecorder) Errorf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogger)(nil).Errorf), varargs...)
}

// Fatal mocks base method
func (m *MockLogger) Fatal(arg0 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatal", varargs...)
}

// Fatal indicates an expected call of Fatal
func (mr *MockLoggerMockRecorder) Fatal(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockLogger)(nil).Fatal), arg0...)
}

// Fatalf mocks base method
func (m *MockLogger) Fatalf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatalf", varargs...)
}

// Fatalf indicates an expected call of Fatalf
func (mr *MockLoggerMockRecorder) Fatalf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatalf", reflect.TypeOf((*MockLogger)(nil).Fatalf), varargs...)
}

// Fields mocks base method
func (m *MockLogger) Fields(arg0 map[string]interface{}) moira.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fields", arg0)
	ret0, _ := ret[0].(moira.Logger)
	return ret0
}

// Fields indicates an expected call of Fields
func (mr *MockLoggerMockRecorder) Fields(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fields", reflect.TypeOf((*MockLogger)(nil).Fields), arg0)
}

// Info mocks base method
func (m *MockLogger) Info(arg0 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info
func (mr *MockLoggerMockRecorder) Info(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLogger)(nil).Info), arg0...)
}

// Infof mocks base method
func (m *MockLogger) Infof(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof
func (mr *MockLoggerMockRecorder) Infof(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockLogger)(nil).Infof), varargs...)
}

// Int mocks base method
func (m *MockLogger) Int(arg0 string, arg1 int) moira.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int", arg0, arg1)
	ret0, _ := ret[0].(moira.Logger)
	return ret0
}

// Int indicates an expected call of Int
func (mr *MockLoggerMockRecorder) Int(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int", reflect.TypeOf((*MockLogger)(nil).Int), arg0, arg1)
}

// Int64 mocks base method
func (m *MockLogger) Int64(arg0 string, arg1 int64) moira.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int64", arg0, arg1)
	ret0, _ := ret[0].(moira.Logger)
	return ret0
}

// Int64 indicates an expected call of Int64
func (mr *MockLoggerMockRecorder) Int64(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int64", reflect.TypeOf((*MockLogger)(nil).Int64), arg0, arg1)
}

// Level mocks base method
func (m *MockLogger) Level(arg0 string) (moira.Logger, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Level", arg0)
	ret0, _ := ret[0].(moira.Logger)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Level indicates an expected call of Level
func (mr *MockLoggerMockRecorder) Level(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Level", reflect.TypeOf((*MockLogger)(nil).Level), arg0)
}

// String mocks base method
func (m *MockLogger) String(arg0, arg1 string) moira.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String", arg0, arg1)
	ret0, _ := ret[0].(moira.Logger)
	return ret0
}

// String indicates an expected call of String
func (mr *MockLoggerMockRecorder) String(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockLogger)(nil).String), arg0, arg1)
}

// Warning mocks base method
func (m *MockLogger) Warning(arg0 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warning", varargs...)
}

// Warning indicates an expected call of Warning
func (mr *MockLoggerMockRecorder) Warning(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warning", reflect.TypeOf((*MockLogger)(nil).Warning), arg0...)
}

// Warningf mocks base method
func (m *MockLogger) Warningf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warningf", varargs...)
}

// Warningf indicates an expected call of Warningf
func (mr *MockLoggerMockRecorder) Warningf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warningf", reflect.TypeOf((*MockLogger)(nil).Warningf), varargs...)
}
