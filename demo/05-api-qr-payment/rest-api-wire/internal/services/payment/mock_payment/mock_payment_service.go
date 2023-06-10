// Code generated by MockGen. DO NOT EDIT.
// Source: ./payment.go

// Package mock_payment_service is a generated GoMock package.
package mock_payment_service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPaymentService is a mock of PaymentService interface.
type MockPaymentService struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentServiceMockRecorder
}

// MockPaymentServiceMockRecorder is the mock recorder for MockPaymentService.
type MockPaymentServiceMockRecorder struct {
	mock *MockPaymentService
}

// NewMockPaymentService creates a new mock instance.
func NewMockPaymentService(ctrl *gomock.Controller) *MockPaymentService {
	mock := &MockPaymentService{ctrl: ctrl}
	mock.recorder = &MockPaymentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentService) EXPECT() *MockPaymentServiceMockRecorder {
	return m.recorder
}

// GenerateQr mocks base method.
func (m *MockPaymentService) GenerateQr(ctx context.Context, promptPayID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateQr", ctx, promptPayID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateQr indicates an expected call of GenerateQr.
func (mr *MockPaymentServiceMockRecorder) GenerateQr(ctx, promptPayID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateQr", reflect.TypeOf((*MockPaymentService)(nil).GenerateQr), ctx, promptPayID)
}
