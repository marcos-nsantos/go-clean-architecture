// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/port/product_gateway_port.go
//
// Generated by this command:
//
//	mockgen -source=internal/core/port/product_gateway_port.go -destination=internal/core/port/mocks/product_gateway_mock.go
//

// Package mock_port is a generated GoMock package.
package mock_port

import (
	context "context"
	entity "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"


	gomock "go.uber.org/mock/gomock"
)

// MockProductGateway is a mock of ProductGateway interface.
type MockProductGateway struct {
	ctrl     *gomock.Controller
	recorder *MockProductGatewayMockRecorder
	isgomock struct{}
}

// MockProductGatewayMockRecorder is the mock recorder for MockProductGateway.
type MockProductGatewayMockRecorder struct {
	mock *MockProductGateway
}

// NewMockProductGateway creates a new mock instance.
func NewMockProductGateway(ctrl *gomock.Controller) *MockProductGateway {
	mock := &MockProductGateway{ctrl: ctrl}
	mock.recorder = &MockProductGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductGateway) EXPECT() *MockProductGatewayMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductGateway) Create(ctx context.Context, product *entity.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockProductGatewayMockRecorder) Create(ctx, product any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductGateway)(nil).Create), ctx, product)
}

// Delete mocks base method.
func (m *MockProductGateway) Delete(ctx context.Context, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockProductGatewayMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProductGateway)(nil).Delete), ctx, id)
}

// FindAll mocks base method.
func (m *MockProductGateway) FindAll(ctx context.Context, name string, categoryID uint64, page, limit int) ([]*entity.Product, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, name, categoryID, page, limit)
	ret0, _ := ret[0].([]*entity.Product)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindAll indicates an expected call of FindAll.
func (mr *MockProductGatewayMockRecorder) FindAll(ctx, name, categoryID, page, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockProductGateway)(nil).FindAll), ctx, name, categoryID, page, limit)
}

// FindByID mocks base method.
func (m *MockProductGateway) FindByID(ctx context.Context, id uint64) (*entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(*entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockProductGatewayMockRecorder) FindByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockProductGateway)(nil).FindByID), ctx, id)
}

// Update mocks base method.
func (m *MockProductGateway) Update(ctx context.Context, product *entity.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockProductGatewayMockRecorder) Update(ctx, product any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductGateway)(nil).Update), ctx, product)
}
