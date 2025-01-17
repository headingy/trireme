// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/headingy/trireme/monitor (interfaces: Monitor,ProcessingUnitsHandler,SynchronizationHandler)
// nolint
package mockmonitor

import (
	gomock "github.com/golang/mock/gomock"
	monitor "github.com/headingy/trireme/monitor"
	policy "github.com/headingy/trireme/policy"
)

// MockMonitor is a mock of Monitor interface
type MockMonitor struct {
	ctrl     *gomock.Controller
	recorder *MockMonitorMockRecorder
}

// MockMonitorMockRecorder is the mock recorder for MockMonitor
type MockMonitorMockRecorder struct {
	mock *MockMonitor
}

// NewMockMonitor creates a new mock instance
func NewMockMonitor(ctrl *gomock.Controller) *MockMonitor {
	mock := &MockMonitor{ctrl: ctrl}
	mock.recorder = &MockMonitorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockMonitor) EXPECT() *MockMonitorMockRecorder {
	return _m.recorder
}

// Start mocks base method
func (_m *MockMonitor) Start() error {
	ret := _m.ctrl.Call(_m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (_mr *MockMonitorMockRecorder) Start() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start")
}

// Stop mocks base method
func (_m *MockMonitor) Stop() error {
	ret := _m.ctrl.Call(_m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (_mr *MockMonitorMockRecorder) Stop() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stop")
}

// MockProcessingUnitsHandler is a mock of ProcessingUnitsHandler interface
type MockProcessingUnitsHandler struct {
	ctrl     *gomock.Controller
	recorder *MockProcessingUnitsHandlerMockRecorder
}

// MockProcessingUnitsHandlerMockRecorder is the mock recorder for MockProcessingUnitsHandler
type MockProcessingUnitsHandlerMockRecorder struct {
	mock *MockProcessingUnitsHandler
}

// NewMockProcessingUnitsHandler creates a new mock instance
func NewMockProcessingUnitsHandler(ctrl *gomock.Controller) *MockProcessingUnitsHandler {
	mock := &MockProcessingUnitsHandler{ctrl: ctrl}
	mock.recorder = &MockProcessingUnitsHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockProcessingUnitsHandler) EXPECT() *MockProcessingUnitsHandlerMockRecorder {
	return _m.recorder
}

// HandlePUEvent mocks base method
func (_m *MockProcessingUnitsHandler) HandlePUEvent(_param0 string, _param1 monitor.Event) error {
	ret := _m.ctrl.Call(_m, "HandlePUEvent", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandlePUEvent indicates an expected call of HandlePUEvent
func (_mr *MockProcessingUnitsHandlerMockRecorder) HandlePUEvent(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HandlePUEvent", arg0, arg1)
}

// SetPURuntime mocks base method
func (_m *MockProcessingUnitsHandler) SetPURuntime(_param0 string, _param1 *policy.PURuntime) error {
	ret := _m.ctrl.Call(_m, "SetPURuntime", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPURuntime indicates an expected call of SetPURuntime
func (_mr *MockProcessingUnitsHandlerMockRecorder) SetPURuntime(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetPURuntime", arg0, arg1)
}

// MockSynchronizationHandler is a mock of SynchronizationHandler interface
type MockSynchronizationHandler struct {
	ctrl     *gomock.Controller
	recorder *MockSynchronizationHandlerMockRecorder
}

// MockSynchronizationHandlerMockRecorder is the mock recorder for MockSynchronizationHandler
type MockSynchronizationHandlerMockRecorder struct {
	mock *MockSynchronizationHandler
}

// NewMockSynchronizationHandler creates a new mock instance
func NewMockSynchronizationHandler(ctrl *gomock.Controller) *MockSynchronizationHandler {
	mock := &MockSynchronizationHandler{ctrl: ctrl}
	mock.recorder = &MockSynchronizationHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockSynchronizationHandler) EXPECT() *MockSynchronizationHandlerMockRecorder {
	return _m.recorder
}

// HandleSynchronization mocks base method
func (_m *MockSynchronizationHandler) HandleSynchronization(_param0 string, _param1 monitor.State, _param2 policy.RuntimeReader, _param3 monitor.SynchronizationType) error {
	ret := _m.ctrl.Call(_m, "HandleSynchronization", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleSynchronization indicates an expected call of HandleSynchronization
func (_mr *MockSynchronizationHandlerMockRecorder) HandleSynchronization(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HandleSynchronization", arg0, arg1, arg2, arg3)
}

// HandleSynchronizationComplete mocks base method
func (_m *MockSynchronizationHandler) HandleSynchronizationComplete(_param0 monitor.SynchronizationType) {
	_m.ctrl.Call(_m, "HandleSynchronizationComplete", _param0)
}

// HandleSynchronizationComplete indicates an expected call of HandleSynchronizationComplete
func (_mr *MockSynchronizationHandlerMockRecorder) HandleSynchronizationComplete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HandleSynchronizationComplete", arg0)
}
