// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/keybase/client/go/kbfs/libkey (interfaces: KeyOps,KeyServer)

// Package libkbfs is a generated GoMock package.
package libkbfs

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	kbfscrypto "github.com/keybase/client/go/kbfs/kbfscrypto"
	kbfsmd "github.com/keybase/client/go/kbfs/kbfsmd"
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
)

// MockKeyOps is a mock of KeyOps interface.
type MockKeyOps struct {
	ctrl     *gomock.Controller
	recorder *MockKeyOpsMockRecorder
}

// MockKeyOpsMockRecorder is the mock recorder for MockKeyOps.
type MockKeyOpsMockRecorder struct {
	mock *MockKeyOps
}

// NewMockKeyOps creates a new mock instance.
func NewMockKeyOps(ctrl *gomock.Controller) *MockKeyOps {
	mock := &MockKeyOps{ctrl: ctrl}
	mock.recorder = &MockKeyOpsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeyOps) EXPECT() *MockKeyOpsMockRecorder {
	return m.recorder
}

// DeleteTLFCryptKeyServerHalf mocks base method.
func (m *MockKeyOps) DeleteTLFCryptKeyServerHalf(arg0 context.Context, arg1 keybase1.UID, arg2 kbfscrypto.CryptPublicKey, arg3 kbfscrypto.TLFCryptKeyServerHalfID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTLFCryptKeyServerHalf", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTLFCryptKeyServerHalf indicates an expected call of DeleteTLFCryptKeyServerHalf.
func (mr *MockKeyOpsMockRecorder) DeleteTLFCryptKeyServerHalf(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTLFCryptKeyServerHalf", reflect.TypeOf((*MockKeyOps)(nil).DeleteTLFCryptKeyServerHalf), arg0, arg1, arg2, arg3)
}

// GetTLFCryptKeyServerHalf mocks base method.
func (m *MockKeyOps) GetTLFCryptKeyServerHalf(arg0 context.Context, arg1 kbfscrypto.TLFCryptKeyServerHalfID, arg2 kbfscrypto.CryptPublicKey) (kbfscrypto.TLFCryptKeyServerHalf, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTLFCryptKeyServerHalf", arg0, arg1, arg2)
	ret0, _ := ret[0].(kbfscrypto.TLFCryptKeyServerHalf)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTLFCryptKeyServerHalf indicates an expected call of GetTLFCryptKeyServerHalf.
func (mr *MockKeyOpsMockRecorder) GetTLFCryptKeyServerHalf(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTLFCryptKeyServerHalf", reflect.TypeOf((*MockKeyOps)(nil).GetTLFCryptKeyServerHalf), arg0, arg1, arg2)
}

// PutTLFCryptKeyServerHalves mocks base method.
func (m *MockKeyOps) PutTLFCryptKeyServerHalves(arg0 context.Context, arg1 kbfsmd.UserDeviceKeyServerHalves) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutTLFCryptKeyServerHalves", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutTLFCryptKeyServerHalves indicates an expected call of PutTLFCryptKeyServerHalves.
func (mr *MockKeyOpsMockRecorder) PutTLFCryptKeyServerHalves(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutTLFCryptKeyServerHalves", reflect.TypeOf((*MockKeyOps)(nil).PutTLFCryptKeyServerHalves), arg0, arg1)
}

// MockKeyServer is a mock of KeyServer interface.
type MockKeyServer struct {
	ctrl     *gomock.Controller
	recorder *MockKeyServerMockRecorder
}

// MockKeyServerMockRecorder is the mock recorder for MockKeyServer.
type MockKeyServerMockRecorder struct {
	mock *MockKeyServer
}

// NewMockKeyServer creates a new mock instance.
func NewMockKeyServer(ctrl *gomock.Controller) *MockKeyServer {
	mock := &MockKeyServer{ctrl: ctrl}
	mock.recorder = &MockKeyServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeyServer) EXPECT() *MockKeyServerMockRecorder {
	return m.recorder
}

// DeleteTLFCryptKeyServerHalf mocks base method.
func (m *MockKeyServer) DeleteTLFCryptKeyServerHalf(arg0 context.Context, arg1 keybase1.UID, arg2 kbfscrypto.CryptPublicKey, arg3 kbfscrypto.TLFCryptKeyServerHalfID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTLFCryptKeyServerHalf", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTLFCryptKeyServerHalf indicates an expected call of DeleteTLFCryptKeyServerHalf.
func (mr *MockKeyServerMockRecorder) DeleteTLFCryptKeyServerHalf(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTLFCryptKeyServerHalf", reflect.TypeOf((*MockKeyServer)(nil).DeleteTLFCryptKeyServerHalf), arg0, arg1, arg2, arg3)
}

// GetTLFCryptKeyServerHalf mocks base method.
func (m *MockKeyServer) GetTLFCryptKeyServerHalf(arg0 context.Context, arg1 kbfscrypto.TLFCryptKeyServerHalfID, arg2 kbfscrypto.CryptPublicKey) (kbfscrypto.TLFCryptKeyServerHalf, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTLFCryptKeyServerHalf", arg0, arg1, arg2)
	ret0, _ := ret[0].(kbfscrypto.TLFCryptKeyServerHalf)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTLFCryptKeyServerHalf indicates an expected call of GetTLFCryptKeyServerHalf.
func (mr *MockKeyServerMockRecorder) GetTLFCryptKeyServerHalf(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTLFCryptKeyServerHalf", reflect.TypeOf((*MockKeyServer)(nil).GetTLFCryptKeyServerHalf), arg0, arg1, arg2)
}

// PutTLFCryptKeyServerHalves mocks base method.
func (m *MockKeyServer) PutTLFCryptKeyServerHalves(arg0 context.Context, arg1 kbfsmd.UserDeviceKeyServerHalves) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutTLFCryptKeyServerHalves", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutTLFCryptKeyServerHalves indicates an expected call of PutTLFCryptKeyServerHalves.
func (mr *MockKeyServerMockRecorder) PutTLFCryptKeyServerHalves(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutTLFCryptKeyServerHalves", reflect.TypeOf((*MockKeyServer)(nil).PutTLFCryptKeyServerHalves), arg0, arg1)
}

// Shutdown mocks base method.
func (m *MockKeyServer) Shutdown() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Shutdown")
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockKeyServerMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockKeyServer)(nil).Shutdown))
}
