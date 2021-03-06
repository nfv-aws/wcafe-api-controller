// Code generated by MockGen. DO NOT EDIT.
// Source: service/pets_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	entity "github.com/nfv-aws/wcafe-api-controller/entity"
	reflect "reflect"
)

// MockPetService is a mock of PetService interface.
type MockPetService struct {
	ctrl     *gomock.Controller
	recorder *MockPetServiceMockRecorder
}

// MockPetServiceMockRecorder is the mock recorder for MockPetService.
type MockPetServiceMockRecorder struct {
	mock *MockPetService
}

// NewMockPetService creates a new mock instance.
func NewMockPetService(ctrl *gomock.Controller) *MockPetService {
	mock := &MockPetService{ctrl: ctrl}
	mock.recorder = &MockPetServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPetService) EXPECT() *MockPetServiceMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockPetService) List() ([]entity.Pet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]entity.Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPetServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPetService)(nil).List))
}

// Create mocks base method.
func (m *MockPetService) Create(c *gin.Context) (entity.Pet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", c)
	ret0, _ := ret[0].(entity.Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPetServiceMockRecorder) Create(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPetService)(nil).Create), c)
}

// Get mocks base method.
func (m *MockPetService) Get(id string) (entity.Pet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(entity.Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPetServiceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPetService)(nil).Get), id)
}

// Update mocks base method.
func (m *MockPetService) Update(id string, c *gin.Context) (entity.Pet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, c)
	ret0, _ := ret[0].(entity.Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPetServiceMockRecorder) Update(id, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPetService)(nil).Update), id, c)
}

// Delete mocks base method.
func (m *MockPetService) Delete(id string) (entity.Pet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(entity.Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockPetServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPetService)(nil).Delete), id)
}
