// Code generated by MockGen. DO NOT EDIT.
// Source: code.storageos.net/storageos/service/director/v1 (interfaces: DirectorClient)

// Package mock_v1 is a generated GoMock package.
package mock_v1

import (
	v1 "code.storageos.net/storageos/service/common/v1"
	v10 "code.storageos.net/storageos/service/director/v1"
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockDirectorClient is a mock of DirectorClient interface
type MockDirectorClient struct {
	ctrl     *gomock.Controller
	recorder *MockDirectorClientMockRecorder
}

// MockDirectorClientMockRecorder is the mock recorder for MockDirectorClient
type MockDirectorClientMockRecorder struct {
	mock *MockDirectorClient
}

// NewMockDirectorClient creates a new mock instance
func NewMockDirectorClient(ctrl *gomock.Controller) *MockDirectorClient {
	mock := &MockDirectorClient{ctrl: ctrl}
	mock.recorder = &MockDirectorClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDirectorClient) EXPECT() *MockDirectorClientMockRecorder {
	return m.recorder
}

// ConfigGetBool mocks base method
func (m *MockDirectorClient) ConfigGetBool(arg0 context.Context, arg1 *v1.ConfigKey, arg2 ...grpc.CallOption) (*v1.ConfigGetBoolReply, error) {
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
func (mr *MockDirectorClientMockRecorder) ConfigGetBool(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigGetBool", reflect.TypeOf((*MockDirectorClient)(nil).ConfigGetBool), varargs...)
}

// ConfigGetString mocks base method
func (m *MockDirectorClient) ConfigGetString(arg0 context.Context, arg1 *v1.ConfigKey, arg2 ...grpc.CallOption) (*v1.ConfigGetStringReply, error) {
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
func (mr *MockDirectorClientMockRecorder) ConfigGetString(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigGetString", reflect.TypeOf((*MockDirectorClient)(nil).ConfigGetString), varargs...)
}

// ConfigGetUint32 mocks base method
func (m *MockDirectorClient) ConfigGetUint32(arg0 context.Context, arg1 *v1.ConfigKey, arg2 ...grpc.CallOption) (*v1.ConfigGetUint32Reply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConfigGetUint32", varargs...)
	ret0, _ := ret[0].(*v1.ConfigGetUint32Reply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigGetUint32 indicates an expected call of ConfigGetUint32
func (mr *MockDirectorClientMockRecorder) ConfigGetUint32(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigGetUint32", reflect.TypeOf((*MockDirectorClient)(nil).ConfigGetUint32), varargs...)
}

// ConfigListBool mocks base method
func (m *MockDirectorClient) ConfigListBool(arg0 context.Context, arg1 *v1.ConfigListQuery, arg2 ...grpc.CallOption) (*v1.ConfigBoolList, error) {
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
func (mr *MockDirectorClientMockRecorder) ConfigListBool(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigListBool", reflect.TypeOf((*MockDirectorClient)(nil).ConfigListBool), varargs...)
}

// ConfigListString mocks base method
func (m *MockDirectorClient) ConfigListString(arg0 context.Context, arg1 *v1.ConfigListQuery, arg2 ...grpc.CallOption) (*v1.ConfigStringList, error) {
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
func (mr *MockDirectorClientMockRecorder) ConfigListString(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigListString", reflect.TypeOf((*MockDirectorClient)(nil).ConfigListString), varargs...)
}

// ConfigListUint32 mocks base method
func (m *MockDirectorClient) ConfigListUint32(arg0 context.Context, arg1 *v1.ConfigListQuery, arg2 ...grpc.CallOption) (*v1.ConfigUint32List, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConfigListUint32", varargs...)
	ret0, _ := ret[0].(*v1.ConfigUint32List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigListUint32 indicates an expected call of ConfigListUint32
func (mr *MockDirectorClientMockRecorder) ConfigListUint32(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigListUint32", reflect.TypeOf((*MockDirectorClient)(nil).ConfigListUint32), varargs...)
}

// ConfigUpdateBool mocks base method
func (m *MockDirectorClient) ConfigUpdateBool(arg0 context.Context, arg1 *v1.ConfigBool, arg2 ...grpc.CallOption) (*v1.ConfigUpdateReply, error) {
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
func (mr *MockDirectorClientMockRecorder) ConfigUpdateBool(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigUpdateBool", reflect.TypeOf((*MockDirectorClient)(nil).ConfigUpdateBool), varargs...)
}

// ConfigUpdateString mocks base method
func (m *MockDirectorClient) ConfigUpdateString(arg0 context.Context, arg1 *v1.ConfigString, arg2 ...grpc.CallOption) (*v1.ConfigUpdateReply, error) {
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
func (mr *MockDirectorClientMockRecorder) ConfigUpdateString(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigUpdateString", reflect.TypeOf((*MockDirectorClient)(nil).ConfigUpdateString), varargs...)
}

// ConfigUpdateUint32 mocks base method
func (m *MockDirectorClient) ConfigUpdateUint32(arg0 context.Context, arg1 *v1.ConfigUint32, arg2 ...grpc.CallOption) (*v1.ConfigUpdateReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConfigUpdateUint32", varargs...)
	ret0, _ := ret[0].(*v1.ConfigUpdateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigUpdateUint32 indicates an expected call of ConfigUpdateUint32
func (mr *MockDirectorClientMockRecorder) ConfigUpdateUint32(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigUpdateUint32", reflect.TypeOf((*MockDirectorClient)(nil).ConfigUpdateUint32), varargs...)
}

// PresentationCreate mocks base method
func (m *MockDirectorClient) PresentationCreate(arg0 context.Context, arg1 *v10.DirectorPresentation, arg2 ...grpc.CallOption) (*v1.RpcResult, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PresentationCreate", varargs...)
	ret0, _ := ret[0].(*v1.RpcResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PresentationCreate indicates an expected call of PresentationCreate
func (mr *MockDirectorClientMockRecorder) PresentationCreate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PresentationCreate", reflect.TypeOf((*MockDirectorClient)(nil).PresentationCreate), varargs...)
}

// PresentationDelete mocks base method
func (m *MockDirectorClient) PresentationDelete(arg0 context.Context, arg1 *v10.DirectorPresentation, arg2 ...grpc.CallOption) (*v1.RpcResult, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PresentationDelete", varargs...)
	ret0, _ := ret[0].(*v1.RpcResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PresentationDelete indicates an expected call of PresentationDelete
func (mr *MockDirectorClientMockRecorder) PresentationDelete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PresentationDelete", reflect.TypeOf((*MockDirectorClient)(nil).PresentationDelete), varargs...)
}

// PresentationList mocks base method
func (m *MockDirectorClient) PresentationList(arg0 context.Context, arg1 *v10.DirectorPresentationListQuery, arg2 ...grpc.CallOption) (*v10.DirectorPresentationList, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PresentationList", varargs...)
	ret0, _ := ret[0].(*v10.DirectorPresentationList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PresentationList indicates an expected call of PresentationList
func (mr *MockDirectorClientMockRecorder) PresentationList(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PresentationList", reflect.TypeOf((*MockDirectorClient)(nil).PresentationList), varargs...)
}

// PresentationUpdate mocks base method
func (m *MockDirectorClient) PresentationUpdate(arg0 context.Context, arg1 *v10.DirectorPresentation, arg2 ...grpc.CallOption) (*v1.RpcResult, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PresentationUpdate", varargs...)
	ret0, _ := ret[0].(*v1.RpcResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PresentationUpdate indicates an expected call of PresentationUpdate
func (mr *MockDirectorClientMockRecorder) PresentationUpdate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PresentationUpdate", reflect.TypeOf((*MockDirectorClient)(nil).PresentationUpdate), varargs...)
}

// Status mocks base method
func (m *MockDirectorClient) Status(arg0 context.Context, arg1 *v10.DirectorStatusRequest, arg2 ...grpc.CallOption) (*v10.DirectorStatus, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Status", varargs...)
	ret0, _ := ret[0].(*v10.DirectorStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status
func (mr *MockDirectorClientMockRecorder) Status(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockDirectorClient)(nil).Status), varargs...)
}

// VolumeCreate mocks base method
func (m *MockDirectorClient) VolumeCreate(arg0 context.Context, arg1 *v10.DirectorVolume, arg2 ...grpc.CallOption) (*v1.RpcResult, error) {
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
func (mr *MockDirectorClientMockRecorder) VolumeCreate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeCreate", reflect.TypeOf((*MockDirectorClient)(nil).VolumeCreate), varargs...)
}

// VolumeDelete mocks base method
func (m *MockDirectorClient) VolumeDelete(arg0 context.Context, arg1 *v10.DirectorVolume, arg2 ...grpc.CallOption) (*v1.RpcResult, error) {
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
func (mr *MockDirectorClientMockRecorder) VolumeDelete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeDelete", reflect.TypeOf((*MockDirectorClient)(nil).VolumeDelete), varargs...)
}

// VolumeList mocks base method
func (m *MockDirectorClient) VolumeList(arg0 context.Context, arg1 *v10.DirectorVolumeListQuery, arg2 ...grpc.CallOption) (*v10.DirectorVolumeList, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "VolumeList", varargs...)
	ret0, _ := ret[0].(*v10.DirectorVolumeList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeList indicates an expected call of VolumeList
func (mr *MockDirectorClientMockRecorder) VolumeList(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeList", reflect.TypeOf((*MockDirectorClient)(nil).VolumeList), varargs...)
}

// VolumeUpdate mocks base method
func (m *MockDirectorClient) VolumeUpdate(arg0 context.Context, arg1 *v10.DirectorVolume, arg2 ...grpc.CallOption) (*v1.RpcResult, error) {
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
func (mr *MockDirectorClientMockRecorder) VolumeUpdate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeUpdate", reflect.TypeOf((*MockDirectorClient)(nil).VolumeUpdate), varargs...)
}
