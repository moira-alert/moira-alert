// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/moira-alert/moira/metric_source (interfaces: MetricSource)

// Package mock_metric_source is a generated GoMock package.
package mock_metric_source

import (
	gomock "github.com/golang/mock/gomock"
	metric_source "github.com/moira-alert/moira/metric_source"
	reflect "reflect"
)

// MockMetricSource is a mock of MetricSource interface
type MockMetricSource struct {
	ctrl     *gomock.Controller
	recorder *MockMetricSourceMockRecorder
}

// MockMetricSourceMockRecorder is the mock recorder for MockMetricSource
type MockMetricSourceMockRecorder struct {
	mock *MockMetricSource
}

// NewMockMetricSource creates a new mock instance
func NewMockMetricSource(ctrl *gomock.Controller) *MockMetricSource {
	mock := &MockMetricSource{ctrl: ctrl}
	mock.recorder = &MockMetricSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetricSource) EXPECT() *MockMetricSourceMockRecorder {
	return m.recorder
}

// Fetch mocks base method
func (m *MockMetricSource) Fetch(arg0 string, arg1, arg2 int64, arg3 bool) (metric_source.FetchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(metric_source.FetchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch
func (mr *MockMetricSourceMockRecorder) Fetch(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockMetricSource)(nil).Fetch), arg0, arg1, arg2, arg3)
}

// IsConfigured mocks base method
func (m *MockMetricSource) IsConfigured() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsConfigured")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsConfigured indicates an expected call of IsConfigured
func (mr *MockMetricSourceMockRecorder) IsConfigured() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsConfigured", reflect.TypeOf((*MockMetricSource)(nil).IsConfigured))
}
