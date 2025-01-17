// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/headingy/trireme/cmd/remoteenforcer (interfaces: Stats)
// nolint
package mockstats

import (
	gomock "github.com/golang/mock/gomock"
)

// MockStats is a mock of Stats interface
type MockStats struct {
	ctrl     *gomock.Controller
	recorder *MockStatsMockRecorder
}

// MockStatsMockRecorder is the mock recorder for MockStats
type MockStatsMockRecorder struct {
	mock *MockStats
}

// NewMockStats creates a new mock instance
func NewMockStats(ctrl *gomock.Controller) *MockStats {
	mock := &MockStats{ctrl: ctrl}
	mock.recorder = &MockStatsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockStats) EXPECT() *MockStatsMockRecorder {
	return _m.recorder
}

// ConnectStatsClient mocks base method
func (_m *MockStats) ConnectStatsClient() error {
	ret := _m.ctrl.Call(_m, "ConnectStatsClient")
	ret0, _ := ret[0].(error)
	return ret0
}

// ConnectStatsClient indicates an expected call of ConnectStatsClient
func (_mr *MockStatsMockRecorder) ConnectStatsClient() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ConnectStatsClient")
}

// SendStats mocks base method
func (_m *MockStats) SendStats() {
	_m.ctrl.Call(_m, "SendStats")
}

// SendStats indicates an expected call of SendStats
func (_mr *MockStatsMockRecorder) SendStats() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SendStats")
}

// Stop mocks base method
func (_m *MockStats) Stop() {
	_m.ctrl.Call(_m, "Stop")
}

// Stop indicates an expected call of Stop
func (_mr *MockStatsMockRecorder) Stop() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stop")
}
