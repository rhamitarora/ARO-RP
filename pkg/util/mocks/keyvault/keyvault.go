// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/util/keyvault (interfaces: Manager)

// Package mock_keyvault is a generated GoMock package.
package mock_keyvault

import (
	context "context"
	rsa "crypto/rsa"
	x509 "crypto/x509"
	reflect "reflect"

	keyvault0 "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
	gomock "github.com/golang/mock/gomock"

	keyvault "github.com/Azure/ARO-RP/pkg/util/keyvault"
)

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// CreateSignedCertificate mocks base method
func (m *MockManager) CreateSignedCertificate(arg0 context.Context, arg1 keyvault.Issuer, arg2, arg3 string, arg4 keyvault.Eku) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSignedCertificate", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSignedCertificate indicates an expected call of CreateSignedCertificate
func (mr *MockManagerMockRecorder) CreateSignedCertificate(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSignedCertificate", reflect.TypeOf((*MockManager)(nil).CreateSignedCertificate), arg0, arg1, arg2, arg3, arg4)
}

// EnsureCertificateDeleted mocks base method
func (m *MockManager) EnsureCertificateDeleted(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureCertificateDeleted", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureCertificateDeleted indicates an expected call of EnsureCertificateDeleted
func (mr *MockManagerMockRecorder) EnsureCertificateDeleted(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureCertificateDeleted", reflect.TypeOf((*MockManager)(nil).EnsureCertificateDeleted), arg0, arg1)
}

// GetBase64Secret mocks base method
func (m *MockManager) GetBase64Secret(arg0 context.Context, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBase64Secret", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBase64Secret indicates an expected call of GetBase64Secret
func (mr *MockManagerMockRecorder) GetBase64Secret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBase64Secret", reflect.TypeOf((*MockManager)(nil).GetBase64Secret), arg0, arg1)
}

// GetCertificateSecret mocks base method
func (m *MockManager) GetCertificateSecret(arg0 context.Context, arg1 string) (*rsa.PrivateKey, []*x509.Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificateSecret", arg0, arg1)
	ret0, _ := ret[0].(*rsa.PrivateKey)
	ret1, _ := ret[1].([]*x509.Certificate)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCertificateSecret indicates an expected call of GetCertificateSecret
func (mr *MockManagerMockRecorder) GetCertificateSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificateSecret", reflect.TypeOf((*MockManager)(nil).GetCertificateSecret), arg0, arg1)
}

// GetSecret mocks base method
func (m *MockManager) GetSecret(arg0 context.Context, arg1 string) (keyvault0.SecretBundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", arg0, arg1)
	ret0, _ := ret[0].(keyvault0.SecretBundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret
func (mr *MockManagerMockRecorder) GetSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockManager)(nil).GetSecret), arg0, arg1)
}

// GetSecrets mocks base method
func (m *MockManager) GetSecrets(arg0 context.Context) ([]keyvault0.SecretItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecrets", arg0)
	ret0, _ := ret[0].([]keyvault0.SecretItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecrets indicates an expected call of GetSecrets
func (mr *MockManagerMockRecorder) GetSecrets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecrets", reflect.TypeOf((*MockManager)(nil).GetSecrets), arg0)
}

// SetSecret mocks base method
func (m *MockManager) SetSecret(arg0 context.Context, arg1 string, arg2 keyvault0.SecretSetParameters) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSecret indicates an expected call of SetSecret
func (mr *MockManagerMockRecorder) SetSecret(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSecret", reflect.TypeOf((*MockManager)(nil).SetSecret), arg0, arg1, arg2)
}

// UpgradeCertificatePolicy mocks base method
func (m *MockManager) UpgradeCertificatePolicy(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeCertificatePolicy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpgradeCertificatePolicy indicates an expected call of UpgradeCertificatePolicy
func (mr *MockManagerMockRecorder) UpgradeCertificatePolicy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeCertificatePolicy", reflect.TypeOf((*MockManager)(nil).UpgradeCertificatePolicy), arg0, arg1)
}

// WaitForCertificateOperation mocks base method
func (m *MockManager) WaitForCertificateOperation(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForCertificateOperation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForCertificateOperation indicates an expected call of WaitForCertificateOperation
func (mr *MockManagerMockRecorder) WaitForCertificateOperation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForCertificateOperation", reflect.TypeOf((*MockManager)(nil).WaitForCertificateOperation), arg0, arg1)
}
