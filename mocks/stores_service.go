// Code generated by MockGen. DO NOT EDIT.
// Source: service/stores_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	entity "github.com/nfv-aws/wcafe-api-controller/entity"
	reflect "reflect"
)

// MockStoreService is a mock of StoreService interface.
type MockStoreService struct {
	ctrl     *gomock.Controller
	recorder *MockStoreServiceMockRecorder
}

// MockStoreServiceMockRecorder is the mock recorder for MockStoreService.
type MockStoreServiceMockRecorder struct {
	mock *MockStoreService
}

// NewMockStoreService creates a new mock instance.
func NewMockStoreService(ctrl *gomock.Controller) *MockStoreService {
	mock := &MockStoreService{ctrl: ctrl}
	mock.recorder = &MockStoreServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoreService) EXPECT() *MockStoreServiceMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockStoreService) List() ([]entity.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]entity.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockStoreServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStoreService)(nil).List))
}

// Create mocks base method.
func (m *MockStoreService) Create(c *gin.Context) (entity.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", c)
	ret0, _ := ret[0].(entity.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockStoreServiceMockRecorder) Create(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStoreService)(nil).Create), c)
}

// Get mocks base method.
func (m *MockStoreService) Get(id string) (entity.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(entity.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStoreServiceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStoreService)(nil).Get), id)
}

// GetName mocks base method.
func (m *MockStoreService) GetName(name string) ([]entity.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName", name)
	ret0, _ := ret[0].([]entity.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetName indicates an expected call of GetName.
func (mr *MockStoreServiceMockRecorder) GetName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockStoreService)(nil).GetName), name)
}

// Update mocks base method.
func (m *MockStoreService) Update(id string, c *gin.Context) (entity.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, c)
	ret0, _ := ret[0].(entity.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockStoreServiceMockRecorder) Update(id, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStoreService)(nil).Update), id, c)
}

// Delete mocks base method.
func (m *MockStoreService) Delete(id string) (entity.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(entity.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockStoreServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStoreService)(nil).Delete), id)
}

// PetsList mocks base method.
func (m *MockStoreService) PetsList(id string) ([]entity.Pet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PetsList", id)
	ret0, _ := ret[0].([]entity.Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PetsList indicates an expected call of PetsList.
func (mr *MockStoreServiceMockRecorder) PetsList(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PetsList", reflect.TypeOf((*MockStoreService)(nil).PetsList), id)
}
