// Automatically generated by MockGen. DO NOT EDIT!
// Source: supervisor/interfaces.go

// nolint
package mockinterfaces

import (
	gomock "github.com/aporeto-inc/mock/gomock"
	policy "github.com/headingy/trireme/policy"
)

// Mock of Supervisor interface
type MockSupervisor struct {
	ctrl     *gomock.Controller
	recorder *_MockSupervisorRecorder
}

// Recorder for MockSupervisor (not exported)
type _MockSupervisorRecorder struct {
	mock *MockSupervisor
}

func NewMockSupervisor(ctrl *gomock.Controller) *MockSupervisor {
	mock := &MockSupervisor{ctrl: ctrl}
	mock.recorder = &_MockSupervisorRecorder{mock}
	return mock
}

func (_m *MockSupervisor) EXPECT() *_MockSupervisorRecorder {
	return _m.recorder
}

func (_m *MockSupervisor) Supervise(contextID string, puInfo *policy.PUInfo) error {
	ret := _m.ctrl.Call(_m, "Supervise", contextID, puInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSupervisorRecorder) Supervise(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Supervise", arg0, arg1)
}

func (_m *MockSupervisor) Unsupervise(contextID string) error {
	ret := _m.ctrl.Call(_m, "Unsupervise", contextID)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSupervisorRecorder) Unsupervise(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Unsupervise", arg0)
}

func (_m *MockSupervisor) Start() error {
	ret := _m.ctrl.Call(_m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSupervisorRecorder) Start() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start")
}

func (_m *MockSupervisor) Stop() error {
	ret := _m.ctrl.Call(_m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSupervisorRecorder) Stop() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stop")
}

func (_m *MockSupervisor) SetTargetNetworks(_param0 []string) error {
	ret := _m.ctrl.Call(_m, "SetTargetNetworks", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSupervisorRecorder) SetTargetNetworks(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetTargetNetworks", arg0)
}

// Mock of Implementor interface
type MockImplementor struct {
	ctrl     *gomock.Controller
	recorder *_MockImplementorRecorder
}

// Recorder for MockImplementor (not exported)
type _MockImplementorRecorder struct {
	mock *MockImplementor
}

func NewMockImplementor(ctrl *gomock.Controller) *MockImplementor {
	mock := &MockImplementor{ctrl: ctrl}
	mock.recorder = &_MockImplementorRecorder{mock}
	return mock
}

func (_m *MockImplementor) EXPECT() *_MockImplementorRecorder {
	return _m.recorder
}

func (_m *MockImplementor) ConfigureRules(version int, contextID string, containerInfo *policy.PUInfo) error {
	ret := _m.ctrl.Call(_m, "ConfigureRules", version, contextID, containerInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImplementorRecorder) ConfigureRules(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ConfigureRules", arg0, arg1, arg2)
}

func (_m *MockImplementor) UpdateRules(version int, contextID string, containerInfo *policy.PUInfo) error {
	ret := _m.ctrl.Call(_m, "UpdateRules", version, contextID, containerInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImplementorRecorder) UpdateRules(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateRules", arg0, arg1, arg2)
}

func (_m *MockImplementor) DeleteRules(version int, context string, ipAddresses policy.ExtendedMap, port string, mark string) error {
	ret := _m.ctrl.Call(_m, "DeleteRules", version, context, ipAddresses, port, mark)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImplementorRecorder) DeleteRules(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteRules", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockImplementor) SetTargetNetworks(_param0 []string, _param1 []string) error {
	ret := _m.ctrl.Call(_m, "SetTargetNetworks", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImplementorRecorder) SetTargetNetworks(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetTargetNetworks", arg0, arg1)
}

func (_m *MockImplementor) Start() error {
	ret := _m.ctrl.Call(_m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImplementorRecorder) Start() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start")
}

func (_m *MockImplementor) Stop() error {
	ret := _m.ctrl.Call(_m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImplementorRecorder) Stop() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stop")
}
