// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/moira-alert/moira/metric_source (interfaces: FetchResult)

// Package mock_metric_source is a generated GoMock package.
package mock_metric_source

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	metric_source "github.com/moira-alert/moira/metric_source"
)

// MockFetchResult is a mock of FetchResult interface
type MockFetchResult struct {
	ctrl     *gomock.Controller
	recorder *MockFetchResultMockRecorder
}

// MockFetchResultMockRecorder is the mock recorder for MockFetchResult
type MockFetchResultMockRecorder struct {
	mock *MockFetchResult
}

// NewMockFetchResult creates a new mock instance
func NewMockFetchResult(ctrl *gomock.Controller) *MockFetchResult {
	mock := &MockFetchResult{ctrl: ctrl}
	mock.recorder = &MockFetchResultMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFetchResult) EXPECT() *MockFetchResultMockRecorder {
	return m.recorder
}

// GetMetricsData mocks base method
func (m *MockFetchResult) GetMetricsData() []*metric_source.MetricData {
	ret := m.ctrl.Call(m, "GetMetricsData")
	ret0, _ := ret[0].([]*metric_source.MetricData)
	return ret0
}

// GetMetricsData indicates an expected call of GetMetricsData
func (mr *MockFetchResultMockRecorder) GetMetricsData() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetricsData", reflect.TypeOf((*MockFetchResult)(nil).GetMetricsData))
}

// GetPatternMetrics mocks base method
func (m *MockFetchResult) GetPatternMetrics() ([]string, error) {
	ret := m.ctrl.Call(m, "GetPatternMetrics")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPatternMetrics indicates an expected call of GetPatternMetrics
func (mr *MockFetchResultMockRecorder) GetPatternMetrics() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPatternMetrics", reflect.TypeOf((*MockFetchResult)(nil).GetPatternMetrics))
}

// GetPatterns mocks base method
func (m *MockFetchResult) GetPatterns() ([]string, error) {
	ret := m.ctrl.Call(m, "GetPatterns")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPatterns indicates an expected call of GetPatterns
func (mr *MockFetchResultMockRecorder) GetPatterns() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPatterns", reflect.TypeOf((*MockFetchResult)(nil).GetPatterns))
}
