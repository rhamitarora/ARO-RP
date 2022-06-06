// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/util/deployer (interfaces: Deployer)

// Package mock_deployer is a generated GoMock package.
package mock_deployer

import (
	context "context"
	fs "io/fs"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	runtime "k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/Azure/ARO-RP/pkg/operator/apis/aro.openshift.io/v1alpha1"
)

// MockDeployer is a mock of Deployer interface.
type MockDeployer struct {
	ctrl     *gomock.Controller
	recorder *MockDeployerMockRecorder
}

// MockDeployerMockRecorder is the mock recorder for MockDeployer.
type MockDeployerMockRecorder struct {
	mock *MockDeployer
}

// NewMockDeployer creates a new mock instance.
func NewMockDeployer(ctrl *gomock.Controller) *MockDeployer {
	mock := &MockDeployer{ctrl: ctrl}
	mock.recorder = &MockDeployerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeployer) EXPECT() *MockDeployerMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method.
func (m *MockDeployer) CreateOrUpdate(arg0 context.Context, arg1 *v1alpha1.Cluster, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockDeployerMockRecorder) CreateOrUpdate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockDeployer)(nil).CreateOrUpdate), arg0, arg1, arg2)
}

// IsReady mocks base method.
func (m *MockDeployer) IsReady(arg0 context.Context, arg1, arg2 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsReady", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsReady indicates an expected call of IsReady.
func (mr *MockDeployerMockRecorder) IsReady(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsReady", reflect.TypeOf((*MockDeployer)(nil).IsReady), arg0, arg1, arg2)
}

// Remove mocks base method.
func (m *MockDeployer) Remove(arg0 context.Context, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockDeployerMockRecorder) Remove(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockDeployer)(nil).Remove), arg0, arg1)
}

// Template mocks base method.
func (m *MockDeployer) Template(arg0 interface{}, arg1 fs.FS) ([]runtime.Object, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Template", arg0, arg1)
	ret0, _ := ret[0].([]runtime.Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Template indicates an expected call of Template.
func (mr *MockDeployerMockRecorder) Template(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Template", reflect.TypeOf((*MockDeployer)(nil).Template), arg0, arg1)
}