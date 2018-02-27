// Code generated by MockGen. DO NOT EDIT.
// Source: code.storageos.net/storageos/service/rdbplugin/v1 (interfaces: RdbPluginClient)

// Package mock_v1 is a generated GoMock package.
package mock_v1

import (
	v1 "code.storageos.net/storageos/service/common/v1"
	v10 "code.storageos.net/storageos/service/rdbplugin/v1"
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockRdbPluginClient is a mock of RdbPluginClient interface
type MockRdbPluginClient struct {
	ctrl     *gomock.Controller
	recorder *MockRdbPluginClientMockRecorder
}

// MockRdbPluginClientMockRecorder is the mock recorder for MockRdbPluginClient
type MockRdbPluginClientMockRecorder struct {
	mock *MockRdbPluginClient
}

// NewMockRdbPluginClient creates a new mock instance
func NewMockRdbPluginClient(ctrl *gomock.Controller) *MockRdbPluginClient {
	mock := &MockRdbPluginClient{ctrl: ctrl}
	mock.recorder = &MockRdbPluginClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRdbPluginClient) EXPECT() *MockRdbPluginClientMockRecorder {
	return m.recorder
}

// ConfigGetBool mocks base method
func (m *MockRdbPluginClient) ConfigGetBool(arg0 context.Context, arg1 *v1.ConfigKey, arg2 ...grpc.CallOption) (*v1.ConfigGetBoolReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConfigGetBool", varargs...)
	ret0, _ := ret[0].(*v1.ConfigGetBoolReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigGetBool indicates an expected call of ConfigGetBool
func (mr *MockRdbPluginClientMockRecorder) ConfigGetBool(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigGetBool", reflect.TypeOf((*MockRdbPluginClient)(nil).ConfigGetBool), varargs...)
}

// ConfigGetString mocks base method
func (m *MockRdbPluginClient) ConfigGetString(arg0 context.Context, arg1 *v1.ConfigKey, arg2 ...grpc.CallOption) (*v1.ConfigGetStringReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConfigGetString", varargs...)
	ret0, _ := ret[0].(*v1.ConfigGetStringReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigGetString indicates an expected call of ConfigGetString
func (mr *MockRdbPluginClientMockRecorder) ConfigGetString(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigGetString", reflect.TypeOf((*MockRdbPluginClient)(nil).ConfigGetString), varargs...)
}

// ConfigListBool mocks base method
func (m *MockRdbPluginClient) ConfigListBool(arg0 context.Context, arg1 *v1.ConfigListQuery, arg2 ...grpc.CallOption) (*v1.ConfigBoolList, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConfigListBool", varargs...)
	ret0, _ := ret[0].(*v1.ConfigBoolList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigListBool indicates an expected call of ConfigListBool
func (mr *MockRdbPluginClientMockRecorder) ConfigListBool(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigListBool", reflect.TypeOf((*MockRdbPluginClient)(nil).ConfigListBool), varargs...)
}

// ConfigListString mocks base method
func (m *MockRdbPluginClient) ConfigListString(arg0 context.Context, arg1 *v1.ConfigListQuery, arg2 ...grpc.CallOption) (*v1.ConfigStringList, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConfigListString", varargs...)
	ret0, _ := ret[0].(*v1.ConfigStringList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigListString indicates an expected call of ConfigListString
func (mr *MockRdbPluginClientMockRecorder) ConfigListString(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigListString", reflect.TypeOf((*MockRdbPluginClient)(nil).ConfigListString), varargs...)
}

// ConfigUpdateBool mocks base method
func (m *MockRdbPluginClient) ConfigUpdateBool(arg0 context.Context, arg1 *v1.ConfigBool, arg2 ...grpc.CallOption) (*v1.ConfigUpdateReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConfigUpdateBool", varargs...)
	ret0, _ := ret[0].(*v1.ConfigUpdateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigUpdateBool indicates an expected call of ConfigUpdateBool
func (mr *MockRdbPluginClientMockRecorder) ConfigUpdateBool(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigUpdateBool", reflect.TypeOf((*MockRdbPluginClient)(nil).ConfigUpdateBool), varargs...)
}

// ConfigUpdateString mocks base method
func (m *MockRdbPluginClient) ConfigUpdateString(arg0 context.Context, arg1 *v1.ConfigString, arg2 ...grpc.CallOption) (*v1.ConfigUpdateReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConfigUpdateString", varargs...)
	ret0, _ := ret[0].(*v1.ConfigUpdateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigUpdateString indicates an expected call of ConfigUpdateString
func (mr *MockRdbPluginClientMockRecorder) ConfigUpdateString(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigUpdateString", reflect.TypeOf((*MockRdbPluginClient)(nil).ConfigUpdateString), varargs...)
}

// Status mocks base method
func (m *MockRdbPluginClient) Status(arg0 context.Context, arg1 *v10.RdbStatusRequest, arg2 ...grpc.CallOption) (*v10.RdbStatus, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Status", varargs...)
	ret0, _ := ret[0].(*v10.RdbStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status
func (mr *MockRdbPluginClientMockRecorder) Status(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockRdbPluginClient)(nil).Status), varargs...)
}

// VolumeCreate mocks base method
func (m *MockRdbPluginClient) VolumeCreate(arg0 context.Context, arg1 *v10.RdbVolume, arg2 ...grpc.CallOption) (*v1.RpcResult, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "VolumeCreate", varargs...)
	ret0, _ := ret[0].(*v1.RpcResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeCreate indicates an expected call of VolumeCreate
func (mr *MockRdbPluginClientMockRecorder) VolumeCreate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeCreate", reflect.TypeOf((*MockRdbPluginClient)(nil).VolumeCreate), varargs...)
}

// VolumeDelete mocks base method
func (m *MockRdbPluginClient) VolumeDelete(arg0 context.Context, arg1 *v10.RdbVolume, arg2 ...grpc.CallOption) (*v1.RpcResult, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "VolumeDelete", varargs...)
	ret0, _ := ret[0].(*v1.RpcResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeDelete indicates an expected call of VolumeDelete
func (mr *MockRdbPluginClientMockRecorder) VolumeDelete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeDelete", reflect.TypeOf((*MockRdbPluginClient)(nil).VolumeDelete), varargs...)
}

// VolumeList mocks base method
func (m *MockRdbPluginClient) VolumeList(arg0 context.Context, arg1 *v10.RdbVolumeListQuery, arg2 ...grpc.CallOption) (*v10.RdbVolumeList, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "VolumeList", varargs...)
	ret0, _ := ret[0].(*v10.RdbVolumeList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeList indicates an expected call of VolumeList
func (mr *MockRdbPluginClientMockRecorder) VolumeList(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeList", reflect.TypeOf((*MockRdbPluginClient)(nil).VolumeList), varargs...)
}

// VolumeUpdate mocks base method
func (m *MockRdbPluginClient) VolumeUpdate(arg0 context.Context, arg1 *v10.RdbVolume, arg2 ...grpc.CallOption) (*v1.RpcResult, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "VolumeUpdate", varargs...)
	ret0, _ := ret[0].(*v1.RpcResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeUpdate indicates an expected call of VolumeUpdate
func (mr *MockRdbPluginClientMockRecorder) VolumeUpdate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeUpdate", reflect.TypeOf((*MockRdbPluginClient)(nil).VolumeUpdate), varargs...)
}
