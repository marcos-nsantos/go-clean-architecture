// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/port/product_usecase_port.go
//
// Generated by this command:
//
//	mockgen -source=internal/core/port/product_usecase_port.go -destination=internal/core/port/mocks/product_usecase_mock.go
//

// Package mock_port is a generated GoMock package.
package mock_port

import (
	context "context"
	reflect "reflect"
	dto "tech-challenge-2-app-example/internal/core/dto"

	gomock "go.uber.org/mock/gomock"
)

// MockListProductsUseCase is a mock of ListProductsUseCase interface.
type MockListProductsUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockListProductsUseCaseMockRecorder
	isgomock struct{}
}

// MockListProductsUseCaseMockRecorder is the mock recorder for MockListProductsUseCase.
type MockListProductsUseCaseMockRecorder struct {
	mock *MockListProductsUseCase
}

// NewMockListProductsUseCase creates a new mock instance.
func NewMockListProductsUseCase(ctrl *gomock.Controller) *MockListProductsUseCase {
	mock := &MockListProductsUseCase{ctrl: ctrl}
	mock.recorder = &MockListProductsUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListProductsUseCase) EXPECT() *MockListProductsUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockListProductsUseCase) Execute(ctx context.Context, req dto.ProductListRequest) (*dto.PaginatedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, req)
	ret0, _ := ret[0].(*dto.PaginatedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockListProductsUseCaseMockRecorder) Execute(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockListProductsUseCase)(nil).Execute), ctx, req)
}

// MockCreateProductUseCase is a mock of CreateProductUseCase interface.
type MockCreateProductUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockCreateProductUseCaseMockRecorder
	isgomock struct{}
}

// MockCreateProductUseCaseMockRecorder is the mock recorder for MockCreateProductUseCase.
type MockCreateProductUseCaseMockRecorder struct {
	mock *MockCreateProductUseCase
}

// NewMockCreateProductUseCase creates a new mock instance.
func NewMockCreateProductUseCase(ctrl *gomock.Controller) *MockCreateProductUseCase {
	mock := &MockCreateProductUseCase{ctrl: ctrl}
	mock.recorder = &MockCreateProductUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreateProductUseCase) EXPECT() *MockCreateProductUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockCreateProductUseCase) Execute(ctx context.Context, req dto.ProductRequest) (*dto.ProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, req)
	ret0, _ := ret[0].(*dto.ProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockCreateProductUseCaseMockRecorder) Execute(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockCreateProductUseCase)(nil).Execute), ctx, req)
}

// MockGetProductUseCase is a mock of GetProductUseCase interface.
type MockGetProductUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockGetProductUseCaseMockRecorder
	isgomock struct{}
}

// MockGetProductUseCaseMockRecorder is the mock recorder for MockGetProductUseCase.
type MockGetProductUseCaseMockRecorder struct {
	mock *MockGetProductUseCase
}

// NewMockGetProductUseCase creates a new mock instance.
func NewMockGetProductUseCase(ctrl *gomock.Controller) *MockGetProductUseCase {
	mock := &MockGetProductUseCase{ctrl: ctrl}
	mock.recorder = &MockGetProductUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetProductUseCase) EXPECT() *MockGetProductUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockGetProductUseCase) Execute(ctx context.Context, id uint64) (*dto.ProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, id)
	ret0, _ := ret[0].(*dto.ProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockGetProductUseCaseMockRecorder) Execute(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockGetProductUseCase)(nil).Execute), ctx, id)
}

// MockUpdateProductUseCase is a mock of UpdateProductUseCase interface.
type MockUpdateProductUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateProductUseCaseMockRecorder
	isgomock struct{}
}

// MockUpdateProductUseCaseMockRecorder is the mock recorder for MockUpdateProductUseCase.
type MockUpdateProductUseCaseMockRecorder struct {
	mock *MockUpdateProductUseCase
}

// NewMockUpdateProductUseCase creates a new mock instance.
func NewMockUpdateProductUseCase(ctrl *gomock.Controller) *MockUpdateProductUseCase {
	mock := &MockUpdateProductUseCase{ctrl: ctrl}
	mock.recorder = &MockUpdateProductUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdateProductUseCase) EXPECT() *MockUpdateProductUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockUpdateProductUseCase) Execute(ctx context.Context, id uint64, req dto.ProductRequest) (*dto.ProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, id, req)
	ret0, _ := ret[0].(*dto.ProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockUpdateProductUseCaseMockRecorder) Execute(ctx, id, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockUpdateProductUseCase)(nil).Execute), ctx, id, req)
}

// MockDeleteProductUseCase is a mock of DeleteProductUseCase interface.
type MockDeleteProductUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockDeleteProductUseCaseMockRecorder
	isgomock struct{}
}

// MockDeleteProductUseCaseMockRecorder is the mock recorder for MockDeleteProductUseCase.
type MockDeleteProductUseCaseMockRecorder struct {
	mock *MockDeleteProductUseCase
}

// NewMockDeleteProductUseCase creates a new mock instance.
func NewMockDeleteProductUseCase(ctrl *gomock.Controller) *MockDeleteProductUseCase {
	mock := &MockDeleteProductUseCase{ctrl: ctrl}
	mock.recorder = &MockDeleteProductUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeleteProductUseCase) EXPECT() *MockDeleteProductUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockDeleteProductUseCase) Execute(ctx context.Context, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockDeleteProductUseCaseMockRecorder) Execute(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockDeleteProductUseCase)(nil).Execute), ctx, id)
}
