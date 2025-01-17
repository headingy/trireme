// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/headingy/trireme/enforcer (interfaces: PolicyEnforcer,PublicKeyAdder,PacketProcessor)

package mockenforcer

import (
	gomock "github.com/golang/mock/gomock"
	enforcer "github.com/headingy/trireme/enforcer"
	fqconfig "github.com/headingy/trireme/enforcer/utils/fqconfig"
	packet "github.com/headingy/trireme/enforcer/utils/packet"
	secrets "github.com/headingy/trireme/enforcer/utils/secrets"
	tokens "github.com/headingy/trireme/enforcer/utils/tokens"
	policy "github.com/headingy/trireme/policy"
)

// MockPolicyEnforcer is a mock of PolicyEnforcer interface
type MockPolicyEnforcer struct {
	ctrl     *gomock.Controller
	recorder *MockPolicyEnforcerMockRecorder
}

// MockPolicyEnforcerMockRecorder is the mock recorder for MockPolicyEnforcer
type MockPolicyEnforcerMockRecorder struct {
	mock *MockPolicyEnforcer
}

// NewMockPolicyEnforcer creates a new mock instance
func NewMockPolicyEnforcer(ctrl *gomock.Controller) *MockPolicyEnforcer {
	mock := &MockPolicyEnforcer{ctrl: ctrl}
	mock.recorder = &MockPolicyEnforcerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockPolicyEnforcer) EXPECT() *MockPolicyEnforcerMockRecorder {
	return _m.recorder
}

// Enforce mocks base method
func (_m *MockPolicyEnforcer) Enforce(_param0 string, _param1 *policy.PUInfo) error {
	ret := _m.ctrl.Call(_m, "Enforce", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Enforce indicates an expected call of Enforce
func (_mr *MockPolicyEnforcerMockRecorder) Enforce(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Enforce", arg0, arg1)
}

// GetFilterQueue mocks base method
func (_m *MockPolicyEnforcer) GetFilterQueue() *fqconfig.FilterQueue {
	ret := _m.ctrl.Call(_m, "GetFilterQueue")
	ret0, _ := ret[0].(*fqconfig.FilterQueue)
	return ret0
}

// GetFilterQueue indicates an expected call of GetFilterQueue
func (_mr *MockPolicyEnforcerMockRecorder) GetFilterQueue() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFilterQueue")
}

// Start mocks base method
func (_m *MockPolicyEnforcer) Start() error {
	ret := _m.ctrl.Call(_m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (_mr *MockPolicyEnforcerMockRecorder) Start() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start")
}

// Stop mocks base method
func (_m *MockPolicyEnforcer) Stop() error {
	ret := _m.ctrl.Call(_m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (_mr *MockPolicyEnforcerMockRecorder) Stop() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stop")
}

// Unenforce mocks base method
func (_m *MockPolicyEnforcer) Unenforce(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Unenforce", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unenforce indicates an expected call of Unenforce
func (_mr *MockPolicyEnforcerMockRecorder) Unenforce(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Unenforce", arg0)
}

// MockPublicKeyAdder is a mock of PublicKeyAdder interface
type MockPublicKeyAdder struct {
	ctrl     *gomock.Controller
	recorder *MockPublicKeyAdderMockRecorder
}

// MockPublicKeyAdderMockRecorder is the mock recorder for MockPublicKeyAdder
type MockPublicKeyAdderMockRecorder struct {
	mock *MockPublicKeyAdder
}

// NewMockPublicKeyAdder creates a new mock instance
func NewMockPublicKeyAdder(ctrl *gomock.Controller) *MockPublicKeyAdder {
	mock := &MockPublicKeyAdder{ctrl: ctrl}
	mock.recorder = &MockPublicKeyAdderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockPublicKeyAdder) EXPECT() *MockPublicKeyAdderMockRecorder {
	return _m.recorder
}

// PublicKeyAdd mocks base method
func (_m *MockPublicKeyAdder) PublicKeyAdd(_param0 string, _param1 []byte) error {
	ret := _m.ctrl.Call(_m, "PublicKeyAdd", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublicKeyAdd indicates an expected call of PublicKeyAdd
func (_mr *MockPublicKeyAdderMockRecorder) PublicKeyAdd(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PublicKeyAdd", arg0, arg1)
}

// MockPacketProcessor is a mock of PacketProcessor interface
type MockPacketProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockPacketProcessorMockRecorder
}

// MockPacketProcessorMockRecorder is the mock recorder for MockPacketProcessor
type MockPacketProcessorMockRecorder struct {
	mock *MockPacketProcessor
}

// NewMockPacketProcessor creates a new mock instance
func NewMockPacketProcessor(ctrl *gomock.Controller) *MockPacketProcessor {
	mock := &MockPacketProcessor{ctrl: ctrl}
	mock.recorder = &MockPacketProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockPacketProcessor) EXPECT() *MockPacketProcessorMockRecorder {
	return _m.recorder
}

// Initialize mocks base method
func (_m *MockPacketProcessor) Initialize(_param0 secrets.Secrets, _param1 *fqconfig.FilterQueue) {
	_m.ctrl.Call(_m, "Initialize", _param0, _param1)
}

// Initialize indicates an expected call of Initialize
func (_mr *MockPacketProcessorMockRecorder) Initialize(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Initialize", arg0, arg1)
}

// PostProcessTCPAppPacket mocks base method
func (_m *MockPacketProcessor) PostProcessTCPAppPacket(_param0 *packet.Packet, _param1 interface{}, _param2 *enforcer.PUContext, _param3 *enforcer.TCPConnection) bool {
	ret := _m.ctrl.Call(_m, "PostProcessTCPAppPacket", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(bool)
	return ret0
}

// PostProcessTCPAppPacket indicates an expected call of PostProcessTCPAppPacket
func (_mr *MockPacketProcessorMockRecorder) PostProcessTCPAppPacket(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PostProcessTCPAppPacket", arg0, arg1, arg2, arg3)
}

// PostProcessTCPNetPacket mocks base method
func (_m *MockPacketProcessor) PostProcessTCPNetPacket(_param0 *packet.Packet, _param1 interface{}, _param2 *tokens.ConnectionClaims, _param3 *enforcer.PUContext, _param4 *enforcer.TCPConnection) bool {
	ret := _m.ctrl.Call(_m, "PostProcessTCPNetPacket", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].(bool)
	return ret0
}

// PostProcessTCPNetPacket indicates an expected call of PostProcessTCPNetPacket
func (_mr *MockPacketProcessorMockRecorder) PostProcessTCPNetPacket(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PostProcessTCPNetPacket", arg0, arg1, arg2, arg3, arg4)
}

// PreProcessTCPAppPacket mocks base method
func (_m *MockPacketProcessor) PreProcessTCPAppPacket(_param0 *packet.Packet, _param1 *enforcer.PUContext, _param2 *enforcer.TCPConnection) bool {
	ret := _m.ctrl.Call(_m, "PreProcessTCPAppPacket", _param0, _param1, _param2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// PreProcessTCPAppPacket indicates an expected call of PreProcessTCPAppPacket
func (_mr *MockPacketProcessorMockRecorder) PreProcessTCPAppPacket(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PreProcessTCPAppPacket", arg0, arg1, arg2)
}

// PreProcessTCPNetPacket mocks base method
func (_m *MockPacketProcessor) PreProcessTCPNetPacket(_param0 *packet.Packet, _param1 *enforcer.PUContext, _param2 *enforcer.TCPConnection) bool {
	ret := _m.ctrl.Call(_m, "PreProcessTCPNetPacket", _param0, _param1, _param2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// PreProcessTCPNetPacket indicates an expected call of PreProcessTCPNetPacket
func (_mr *MockPacketProcessorMockRecorder) PreProcessTCPNetPacket(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PreProcessTCPNetPacket", arg0, arg1, arg2)
}
