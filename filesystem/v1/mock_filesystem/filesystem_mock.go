// Code generated by MockGen. DO NOT EDIT.
// Source: code.storageos.net/scm/storageos/service/filesystem/v1 (interfaces: FsClient)

// Package mock_v1 is a generated GoMock package.
package mock_v1

import (
	v1 "code.storageos.net/scm/storageos/service/common/v1"
	v10 "code.storageos.net/scm/storageos/service/filesystem/v1"
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockFsClient is a mock of FsClient interface
type MockFsClient struct {
	ctrl     *gomock.Controller
	recorder *MockFsClientMockRecorder
}

// MockFsClientMockRecorder is the mock recorder for MockFsClient
type MockFsClientMockRecorder struct {
	mock *MockFsClient
}

// NewMockFsClient creates a new mock instance
func NewMockFsClient(ctrl *gomock.Controller) *MockFsClient {
	mock := &MockFsClient{ctrl: ctrl}
	mock.recorder = &MockFsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFsClient) EXPECT() *MockFsClientMockRecorder {
	return m.recorder
}

// PresentationCreate mocks base method
func (m *MockFsClient) PresentationCreate(arg0 context.Context, arg1 *v10.FsPresentation, arg2 ...grpc.CallOption) (*v1.RpcResult, error) {
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
func (mr *MockFsClientMockRecorder) PresentationCreate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PresentationCreate", reflect.TypeOf((*MockFsClient)(nil).PresentationCreate), varargs...)
}

// PresentationDelete mocks base method
func (m *MockFsClient) PresentationDelete(arg0 context.Context, arg1 *v10.FsPresentation, arg2 ...grpc.CallOption) (*v1.RpcResult, error) {
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
func (mr *MockFsClientMockRecorder) PresentationDelete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PresentationDelete", reflect.TypeOf((*MockFsClient)(nil).PresentationDelete), varargs...)
}

// PresentationList mocks base method
func (m *MockFsClient) PresentationList(arg0 context.Context, arg1 *v10.FsPresentationListQuery, arg2 ...grpc.CallOption) (*v10.FsPresentationList, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PresentationList", varargs...)
	ret0, _ := ret[0].(*v10.FsPresentationList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PresentationList indicates an expected call of PresentationList
func (mr *MockFsClientMockRecorder) PresentationList(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PresentationList", reflect.TypeOf((*MockFsClient)(nil).PresentationList), varargs...)
}

// PresentationUpdate mocks base method
func (m *MockFsClient) PresentationUpdate(arg0 context.Context, arg1 *v10.FsPresentation, arg2 ...grpc.CallOption) (*v1.RpcResult, error) {
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
func (mr *MockFsClientMockRecorder) PresentationUpdate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PresentationUpdate", reflect.TypeOf((*MockFsClient)(nil).PresentationUpdate), varargs...)
}

// Status mocks base method
func (m *MockFsClient) Status(arg0 context.Context, arg1 *v10.FsStatusRequest, arg2 ...grpc.CallOption) (*v10.FsStatus, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Status", varargs...)
	ret0, _ := ret[0].(*v10.FsStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status
func (mr *MockFsClientMockRecorder) Status(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockFsClient)(nil).Status), varargs...)
}

// VolumeCreate mocks base method
func (m *MockFsClient) VolumeCreate(arg0 context.Context, arg1 *v10.FsVolume, arg2 ...grpc.CallOption) (*v1.RpcResult, error) {
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
func (mr *MockFsClientMockRecorder) VolumeCreate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeCreate", reflect.TypeOf((*MockFsClient)(nil).VolumeCreate), varargs...)
}

// VolumeDelete mocks base method
func (m *MockFsClient) VolumeDelete(arg0 context.Context, arg1 *v10.FsVolume, arg2 ...grpc.CallOption) (*v1.RpcResult, error) {
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
func (mr *MockFsClientMockRecorder) VolumeDelete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeDelete", reflect.TypeOf((*MockFsClient)(nil).VolumeDelete), varargs...)
}

// VolumeList mocks base method
func (m *MockFsClient) VolumeList(arg0 context.Context, arg1 *v10.FsVolumeListQuery, arg2 ...grpc.CallOption) (*v10.FsVolumeList, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "VolumeList", varargs...)
	ret0, _ := ret[0].(*v10.FsVolumeList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeList indicates an expected call of VolumeList
func (mr *MockFsClientMockRecorder) VolumeList(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeList", reflect.TypeOf((*MockFsClient)(nil).VolumeList), varargs...)
}

// VolumeUpdate mocks base method
func (m *MockFsClient) VolumeUpdate(arg0 context.Context, arg1 *v10.FsVolume, arg2 ...grpc.CallOption) (*v1.RpcResult, error) {
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
func (mr *MockFsClientMockRecorder) VolumeUpdate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeUpdate", reflect.TypeOf((*MockFsClient)(nil).VolumeUpdate), varargs...)
}