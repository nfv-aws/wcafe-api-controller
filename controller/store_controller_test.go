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

func TestList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().List().Return(service.Stores{}, nil)
	controller := StoreController{Service: serviceMock}

	controller.List(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().Get(gomock.Any()).Return(service.Store{}, nil)
	controller := StoreController{Service: serviceMock}

	controller.Get(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().Create(c).Return(service.Store{}, nil)
	controller := StoreController{Service: serviceMock}

	controller.Create(c)
	assert.Equal(t, 201, c.Writer.Status())
}
