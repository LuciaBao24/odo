// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/watch/interface.go

// Package watch is a generated GoMock package.
package watch

import (
	context "context"
	io "io"
	reflect "reflect"

	parser "github.com/devfile/library/pkg/devfile/parser"
	gomock "github.com/golang/mock/gomock"
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

// CleanupDevResources mocks base method.
func (m *MockClient) CleanupDevResources(devfileObj parser.DevfileObj, out io.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanupDevResources", devfileObj, out)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanupDevResources indicates an expected call of CleanupDevResources.
func (mr *MockClientMockRecorder) CleanupDevResources(devfileObj, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanupDevResources", reflect.TypeOf((*MockClient)(nil).CleanupDevResources), devfileObj, out)
}

// WatchAndPush mocks base method.
func (m *MockClient) WatchAndPush(out io.Writer, parameters WatchParameters, ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchAndPush", out, parameters, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// WatchAndPush indicates an expected call of WatchAndPush.
func (mr *MockClientMockRecorder) WatchAndPush(out, parameters, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchAndPush", reflect.TypeOf((*MockClient)(nil).WatchAndPush), out, parameters, ctx)
}
