package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/nfv-aws/wcafe-api-controller/mocks"
	"github.com/nfv-aws/wcafe-api-controller/service"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestUserList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockUserService(ctrl)
	serviceMock.EXPECT().List().Return(service.Users{}, nil)
	controller := UserController{Service: serviceMock}

	controller.List(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestUserGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockUserService(ctrl)
	serviceMock.EXPECT().Get(gomock.Any()).Return(service.User{}, nil)
	controller := UserController{Service: serviceMock}

	controller.Get(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestUserCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockUserService(ctrl)
	serviceMock.EXPECT().Create(c).Return(service.User{}, nil)
	controller := UserController{Service: serviceMock}

	controller.Create(c)
	assert.Equal(t, 201, c.Writer.Status())
}
