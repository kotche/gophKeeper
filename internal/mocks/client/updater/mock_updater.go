// Code generated by MockGen. DO NOT EDIT.
// Source: internal/client/updater/updater.go

// Package mock_updater is a generated GoMock package.
package mock_updater

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/kotche/gophKeeper/internal/client/domain"
)

// MockISender is a mock of ISender interface.
type MockISender struct {
	ctrl     *gomock.Controller
	recorder *MockISenderMockRecorder
}

// MockISenderMockRecorder is the mock recorder for MockISender.
type MockISenderMockRecorder struct {
	mock *MockISender
}

// NewMockISender creates a new mock instance.
func NewMockISender(ctrl *gomock.Controller) *MockISender {
	mock := &MockISender{ctrl: ctrl}
	mock.recorder = &MockISenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISender) EXPECT() *MockISenderMockRecorder {
	return m.recorder
}

// GetAllBankCard mocks base method.
func (m *MockISender) GetAllBankCard(ctx context.Context) ([]*domain.BankCard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBankCard", ctx)
	ret0, _ := ret[0].([]*domain.BankCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllBankCard indicates an expected call of GetAllBankCard.
func (mr *MockISenderMockRecorder) GetAllBankCard(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBankCard", reflect.TypeOf((*MockISender)(nil).GetAllBankCard), ctx)
}

// GetAllBinary mocks base method.
func (m *MockISender) GetAllBinary(ctx context.Context) ([]*domain.Binary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBinary", ctx)
	ret0, _ := ret[0].([]*domain.Binary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllBinary indicates an expected call of GetAllBinary.
func (mr *MockISenderMockRecorder) GetAllBinary(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBinary", reflect.TypeOf((*MockISender)(nil).GetAllBinary), ctx)
}

// GetAllLoginPass mocks base method.
func (m *MockISender) GetAllLoginPass(ctx context.Context) ([]*domain.LoginPass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllLoginPass", ctx)
	ret0, _ := ret[0].([]*domain.LoginPass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllLoginPass indicates an expected call of GetAllLoginPass.
func (mr *MockISenderMockRecorder) GetAllLoginPass(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllLoginPass", reflect.TypeOf((*MockISender)(nil).GetAllLoginPass), ctx)
}

// GetAllText mocks base method.
func (m *MockISender) GetAllText(ctx context.Context) ([]*domain.Text, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllText", ctx)
	ret0, _ := ret[0].([]*domain.Text)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllText indicates an expected call of GetAllText.
func (mr *MockISenderMockRecorder) GetAllText(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllText", reflect.TypeOf((*MockISender)(nil).GetAllText), ctx)
}

// GetVersionServer mocks base method.
func (m *MockISender) GetVersionServer(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersionServer", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersionServer indicates an expected call of GetVersionServer.
func (mr *MockISenderMockRecorder) GetVersionServer(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersionServer", reflect.TypeOf((*MockISender)(nil).GetVersionServer), ctx)
}

// MockIService is a mock of IService interface.
type MockIService struct {
	ctrl     *gomock.Controller
	recorder *MockIServiceMockRecorder
}

// MockIServiceMockRecorder is the mock recorder for MockIService.
type MockIServiceMockRecorder struct {
	mock *MockIService
}

// NewMockIService creates a new mock instance.
func NewMockIService(ctrl *gomock.Controller) *MockIService {
	mock := &MockIService{ctrl: ctrl}
	mock.recorder = &MockIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIService) EXPECT() *MockIServiceMockRecorder {
	return m.recorder
}

// GetVersionCache mocks base method.
func (m *MockIService) GetVersionCache() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersionCache")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetVersionCache indicates an expected call of GetVersionCache.
func (mr *MockIServiceMockRecorder) GetVersionCache() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersionCache", reflect.TypeOf((*MockIService)(nil).GetVersionCache))
}

// SetVersionCache mocks base method.
func (m *MockIService) SetVersionCache(version int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetVersionCache", version)
}

// SetVersionCache indicates an expected call of SetVersionCache.
func (mr *MockIServiceMockRecorder) SetVersionCache(version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVersionCache", reflect.TypeOf((*MockIService)(nil).SetVersionCache), version)
}

// UpdateAll mocks base method.
func (m *MockIService) UpdateAll(data any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockIServiceMockRecorder) UpdateAll(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockIService)(nil).UpdateAll), data)
}
