// Code generated by MockGen. DO NOT EDIT.
// Source: service/clerks_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	entity "github.com/nfv-aws/wcafe-api-controller/entity"
	reflect "reflect"
)

// MockClerkService is a mock of ClerkService interface.
type MockClerkService struct {
	ctrl     *gomock.Controller
	recorder *MockClerkServiceMockRecorder
}

// MockClerkServiceMockRecorder is the mock recorder for MockClerkService.
type MockClerkServiceMockRecorder struct {
	mock *MockClerkService
}

// NewMockClerkService creates a new mock instance.
func NewMockClerkService(ctrl *gomock.Controller) *MockClerkService {
	mock := &MockClerkService{ctrl: ctrl}
	mock.recorder = &MockClerkServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClerkService) EXPECT() *MockClerkServiceMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockClerkService) List() ([]entity.Clerk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]entity.Clerk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockClerkServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClerkService)(nil).List))
}

// Create mocks base method.
func (m *MockClerkService) Create(c *gin.Context) (entity.Clerk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", c)
	ret0, _ := ret[0].(entity.Clerk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockClerkServiceMockRecorder) Create(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClerkService)(nil).Create), c)
}

// Get mocks base method.
func (m *MockClerkService) Get(id string) (entity.Clerk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(entity.Clerk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockClerkServiceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClerkService)(nil).Get), id)
}

// Update mocks base method.
func (m *MockClerkService) Update(id string, c *gin.Context) (entity.Clerk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, c)
	ret0, _ := ret[0].(entity.Clerk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockClerkServiceMockRecorder) Update(id, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockClerkService)(nil).Update), id, c)
}

// Delete mocks base method.
func (m *MockClerkService) Delete(id string) (entity.Clerk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(entity.Clerk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockClerkServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClerkService)(nil).Delete), id)
}
