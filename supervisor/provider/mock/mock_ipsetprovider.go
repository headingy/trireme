// Automatically generated by MockGen. DO NOT EDIT!
// Source: supervisor/provider/ipsetprovider.go

// nolint
package mockprovider

import (
	gomock "github.com/golang/mock/gomock"
	ipset "github.com/bvandewalle/go-ipset/ipset"
	provider "github.com/headingy/trireme/supervisor/provider"
)

// Mock of IpsetProvider interface
type MockIpsetProvider struct {
	ctrl     *gomock.Controller
	recorder *_MockIpsetProviderRecorder
}

// Recorder for MockIpsetProvider (not exported)
type _MockIpsetProviderRecorder struct {
	mock *MockIpsetProvider
}

func NewMockIpsetProvider(ctrl *gomock.Controller) *MockIpsetProvider {
	mock := &MockIpsetProvider{ctrl: ctrl}
	mock.recorder = &_MockIpsetProviderRecorder{mock}
	return mock
}

func (_m *MockIpsetProvider) EXPECT() *_MockIpsetProviderRecorder {
	return _m.recorder
}

func (_m *MockIpsetProvider) NewIpset(name string, hasht string, p *ipset.Params) (provider.Ipset, error) {
	ret := _m.ctrl.Call(_m, "NewIpset", name, hasht, p)
	ret0, _ := ret[0].(provider.Ipset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockIpsetProviderRecorder) NewIpset(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewIpset", arg0, arg1, arg2)
}

func (_m *MockIpsetProvider) DestroyAll() error {
	ret := _m.ctrl.Call(_m, "DestroyAll")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockIpsetProviderRecorder) DestroyAll() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DestroyAll")
}

// Mock of Ipset interface
type MockIpset struct {
	ctrl     *gomock.Controller
	recorder *_MockIpsetRecorder
}

// Recorder for MockIpset (not exported)
type _MockIpsetRecorder struct {
	mock *MockIpset
}

func NewMockIpset(ctrl *gomock.Controller) *MockIpset {
	mock := &MockIpset{ctrl: ctrl}
	mock.recorder = &_MockIpsetRecorder{mock}
	return mock
}

func (_m *MockIpset) EXPECT() *_MockIpsetRecorder {
	return _m.recorder
}

func (_m *MockIpset) Add(entry string, timeout int) error {
	ret := _m.ctrl.Call(_m, "Add", entry, timeout)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockIpsetRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Add", arg0, arg1)
}

func (_m *MockIpset) AddOption(entry string, option string, timeout int) error {
	ret := _m.ctrl.Call(_m, "AddOption", entry, option, timeout)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockIpsetRecorder) AddOption(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddOption", arg0, arg1, arg2)
}

func (_m *MockIpset) Del(entry string) error {
	ret := _m.ctrl.Call(_m, "Del", entry)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockIpsetRecorder) Del(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Del", arg0)
}

func (_m *MockIpset) Destroy() error {
	ret := _m.ctrl.Call(_m, "Destroy")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockIpsetRecorder) Destroy() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Destroy")
}

func (_m *MockIpset) Flush() error {
	ret := _m.ctrl.Call(_m, "Flush")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockIpsetRecorder) Flush() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Flush")
}

func (_m *MockIpset) Test(entry string) (bool, error) {
	ret := _m.ctrl.Call(_m, "Test", entry)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockIpsetRecorder) Test(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Test", arg0)
}
