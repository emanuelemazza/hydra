// Automatically generated by MockGen. DO NOT EDIT!
// Source: etcd_client.go

package mock_supervisor

import (
	gomock "github.com/innotech/hydra/vendors/code.google.com/p/gomock/gomock"

	. "github.com/innotech/hydra/supervisor"
)

// Mock of EtcdRequester interface
type MockEtcdRequester struct {
	ctrl     *gomock.Controller
	recorder *_MockEtcdRequesterRecorder
}

// Recorder for MockEtcdRequester (not exported)
type _MockEtcdRequesterRecorder struct {
	mock *MockEtcdRequester
}

func NewMockEtcdRequester(ctrl *gomock.Controller) *MockEtcdRequester {
	mock := &MockEtcdRequester{ctrl: ctrl}
	mock.recorder = &_MockEtcdRequesterRecorder{mock}
	return mock
}

func (_m *MockEtcdRequester) EXPECT() *_MockEtcdRequesterRecorder {
	return _m.recorder
}

func (_m *MockEtcdRequester) BaseGet(key string) (*RawResponse, error) {
	ret := _m.ctrl.Call(_m, "BaseGet", key)
	ret0, _ := ret[0].(*RawResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockEtcdRequesterRecorder) BaseGet(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BaseGet", arg0)
}

func (_m *MockEtcdRequester) CompareAndSwap(key string, value string, ttl uint64, prevValue string, prevIndex uint64, prevExist KeyExistence) (*Response, error) {
	ret := _m.ctrl.Call(_m, "CompareAndSwap", key, value, ttl, prevValue, prevIndex, prevExist)
	ret0, _ := ret[0].(*Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockEtcdRequesterRecorder) CompareAndSwap(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CompareAndSwap", arg0, arg1, arg2, arg3, arg4, arg5)
}

func (_m *MockEtcdRequester) Get(key string, sort bool, recursive bool) (*Response, error) {
	ret := _m.ctrl.Call(_m, "Get", key, sort, recursive)
	ret0, _ := ret[0].(*Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockEtcdRequesterRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0, arg1, arg2)
}

func (_m *MockEtcdRequester) Set(key string, value string, ttl uint64) (*Response, error) {
	ret := _m.ctrl.Call(_m, "Set", key, value, ttl)
	ret0, _ := ret[0].(*Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockEtcdRequesterRecorder) Set(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Set", arg0, arg1, arg2)
}

func (_m *MockEtcdRequester) WithMachineAddr(machineAddr string) EtcdRequester {
	ret := _m.ctrl.Call(_m, "WithMachineAddr", machineAddr)
	ret0, _ := ret[0].(EtcdRequester)
	return ret0
}

func (_mr *_MockEtcdRequesterRecorder) WithMachineAddr(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WithMachineAddr", arg0)
}
