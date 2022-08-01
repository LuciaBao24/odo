// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/deploy/interface.go

// Package deploy is a generated GoMock package.
package deploy

import (
	reflect "reflect"

	parser "github.com/devfile/library/pkg/devfile/parser"
	gomock "github.com/golang/mock/gomock"
	filesystem "github.com/redhat-developer/odo/pkg/testingutil/filesystem"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Deploy mocks base method.
func (m *MockClient) Deploy(fs filesystem.Filesystem, devfileObj parser.DevfileObj, path, appName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", fs, devfileObj, path, appName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deploy indicates an expected call of Deploy.
func (mr *MockClientMockRecorder) Deploy(fs, devfileObj, path, appName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockClient)(nil).Deploy), fs, devfileObj, path, appName)
}
