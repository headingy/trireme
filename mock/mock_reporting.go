// Automatically generated by MockGen. DO NOT EDIT!
// Source: collector/interfaces.go
// nolint
package mock_trireme

import (
	gomock "github.com/golang/mock/gomock"
	collector "github.com/headingy/trireme/collector"
)

// Mock of EventCollector interface
type MockEventCollector struct {
	ctrl     *gomock.Controller
	recorder *_MockEventCollectorRecorder
}

// Recorder for MockEventCollector (not exported)
type _MockEventCollectorRecorder struct {
	mock *MockEventCollector
}

func NewMockEventCollector(ctrl *gomock.Controller) *MockEventCollector {
	mock := &MockEventCollector{ctrl: ctrl}
	mock.recorder = &_MockEventCollectorRecorder{mock}
	return mock
}

func (_m *MockEventCollector) EXPECT() *_MockEventCollectorRecorder {
	return _m.recorder
}

func (_m *MockEventCollector) CollectFlowEvent(record *collector.FlowRecord) {
	_m.ctrl.Call(_m, "CollectFlowEvent", record)
}

func (_mr *_MockEventCollectorRecorder) CollectFlowEvent(arg0 interface{}) *gomock.Call {

	return _mr.mock.ctrl.RecordCall(_mr.mock, "CollectFlowEvent", arg0)
}

func (_m *MockEventCollector) CollectContainerEvent(record *collector.ContainerRecord) {
	_m.ctrl.Call(_m, "CollectContainerEvent", record)
}

func (_mr *_MockEventCollectorRecorder) CollectContainerEvent(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CollectContainerEvent", arg0)
}
