// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/czarro/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/czarro/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockStore) CreateProduct(arg0 context.Context, arg1 db.CreateProductParams) (db.CzProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0, arg1)
	ret0, _ := ret[0].(db.CzProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockStoreMockRecorder) CreateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockStore)(nil).CreateProduct), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.CzUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.CzUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteProduct mocks base method.
func (m *MockStore) DeleteProduct(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockStoreMockRecorder) DeleteProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockStore)(nil).DeleteProduct), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// ExecTx mocks base method.
func (m *MockStore) ExecTx(arg0 context.Context, arg1 func(*db.Queries) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecTx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecTx indicates an expected call of ExecTx.
func (mr *MockStoreMockRecorder) ExecTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecTx", reflect.TypeOf((*MockStore)(nil).ExecTx), arg0, arg1)
}

// GetProduct mocks base method.
func (m *MockStore) GetProduct(arg0 context.Context, arg1 int32) (db.CzProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", arg0, arg1)
	ret0, _ := ret[0].(db.CzProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockStoreMockRecorder) GetProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockStore)(nil).GetProduct), arg0, arg1)
}

// GetProductForUpdate mocks base method.
func (m *MockStore) GetProductForUpdate(arg0 context.Context, arg1 int32) (db.CzProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.CzProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductForUpdate indicates an expected call of GetProductForUpdate.
func (mr *MockStoreMockRecorder) GetProductForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductForUpdate", reflect.TypeOf((*MockStore)(nil).GetProductForUpdate), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 int64) (db.CzUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.CzUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetUserForUpdate mocks base method.
func (m *MockStore) GetUserForUpdate(arg0 context.Context, arg1 int64) (db.CzUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.CzUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserForUpdate indicates an expected call of GetUserForUpdate.
func (mr *MockStoreMockRecorder) GetUserForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserForUpdate", reflect.TypeOf((*MockStore)(nil).GetUserForUpdate), arg0, arg1)
}

// ListProducts mocks base method.
func (m *MockStore) ListProducts(arg0 context.Context, arg1 db.ListProductsParams) ([]db.CzProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProducts", arg0, arg1)
	ret0, _ := ret[0].([]db.CzProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProducts indicates an expected call of ListProducts.
func (mr *MockStoreMockRecorder) ListProducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProducts", reflect.TypeOf((*MockStore)(nil).ListProducts), arg0, arg1)
}

// Listusers mocks base method.
func (m *MockStore) Listusers(arg0 context.Context, arg1 db.ListusersParams) ([]db.CzUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Listusers", arg0, arg1)
	ret0, _ := ret[0].([]db.CzUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Listusers indicates an expected call of Listusers.
func (mr *MockStoreMockRecorder) Listusers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Listusers", reflect.TypeOf((*MockStore)(nil).Listusers), arg0, arg1)
}

// UpdateProducts mocks base method.
func (m *MockStore) UpdateProducts(arg0 context.Context, arg1 db.UpdateProductsParams) (db.CzProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProducts", arg0, arg1)
	ret0, _ := ret[0].(db.CzProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProducts indicates an expected call of UpdateProducts.
func (mr *MockStoreMockRecorder) UpdateProducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProducts", reflect.TypeOf((*MockStore)(nil).UpdateProducts), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.CzUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.CzUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}