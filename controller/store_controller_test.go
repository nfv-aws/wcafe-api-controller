package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/nfv-aws/wcafe-api-controller/entity"
	"github.com/nfv-aws/wcafe-api-controller/mocks"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

var (
	s = entity.Store{
		Id:          "sa5bafac-b35c-4852-82ca-b272cd79f2f3",
		Name:        "store_controller_test",
		Tag:         "Board game",
		Address:     "Shinagawa",
		StrongPoint: "helpful",
		CreatedAt:   ct,
		UpdatedAt:   ut,
	}
)

func TestStoreList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)
	s := []entity.Store{
		{Id: "sa5bafac-b35c-4852-82ca-b272cd79f2f3", Name: "Sano Shinichiro"},
		{Id: "sa5bafac-b35c-4852-82ca-b272cd79f2f5", Name: "Suzuki Chihiro"},
	}
	serviceMock.EXPECT().List().Return(s, nil)
	controller := StoreController{Storeservice: serviceMock}

	controller.List(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestStoreGetOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)

	serviceMock.EXPECT().Get(gomock.Any()).Return(s, nil)
	controller := StoreController{Storeservice: serviceMock}

	controller.Get(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestStoreGetNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)

	serviceMock.EXPECT().Get(gomock.Any()).Return(entity.Store{}, ErrRecordNotFound)
	controller := StoreController{Storeservice: serviceMock}

	controller.Get(c)
	assert.Equal(t, 404, c.Writer.Status())
}

func TestStoreCreateOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)

	serviceMock.EXPECT().Create(c).Return(s, nil)
	controller := StoreController{Storeservice: serviceMock}

	controller.Create(c)
	assert.Equal(t, 201, c.Writer.Status())
}

func TestStoreCreateBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)

	serviceMock.EXPECT().Create(c).Return(entity.Store{}, ErrBadRequest)
	controller := StoreController{Storeservice: serviceMock}

	controller.Create(c)
	assert.Equal(t, 400, c.Writer.Status())
}

func TestStoreUpdateOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().Update(gomock.Any(), c).Return(s, nil)
	controller := StoreController{Storeservice: serviceMock}

	controller.Update(c)
	assert.Equal(t, 200, c.Writer.Status())
}

// func TestStoreUpdataNotFound(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	c, _ := gin.CreateTestContext(httptest.NewRecorder())

// 	serviceMock := mocks.NewMockStoreService(ctrl)

// 	serviceMock.EXPECT().Get(gomock.Any()).Return(entity.Store{}, ErrRecordNotFound)
// 	controller := StoreController{Storeservice: serviceMock}

// 	controller.Update(c)
// 	assert.Equal(t, 404, c.Writer.Status())
// }

func TestStoreUpdateBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)

	serviceMock.EXPECT().Update(gomock.Any(), c).Return(entity.Store{}, ErrBadRequest)
	controller := StoreController{Storeservice: serviceMock}

	controller.Update(c)
	assert.Equal(t, 400, c.Writer.Status())
}
