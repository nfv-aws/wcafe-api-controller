package controller

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/nfv-aws/wcafe-api-controller/entity"
	"github.com/nfv-aws/wcafe-api-controller/mocks"
	"github.com/nfv-aws/wcafe-api-controller/service"
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
	ErrConflict = errors.New("Error 1451: Cannot delete or update a parent row: a foreign key constraint fails (`wcafe`.`pets`, CONSTRAINT `pets_store_id_stores_id_foreign` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT)")
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

	serviceMock.EXPECT().Get(gomock.Any()).Return(entity.Store{}, gorm.ErrRecordNotFound)
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

func TestStoreCreateInvalidAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().Create(c).Return(entity.Store{}, ErrInvalidAddress)
	controller := StoreController{Storeservice: serviceMock}

	controller.Create(c)
	assert.Equal(t, 404, c.Writer.Status())
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

func TestStoreUpdate(t *testing.T) {
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

func TestStoreUpdataNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)

	serviceMock.EXPECT().Update(gomock.Any(), c).Return(entity.Store{}, gorm.ErrRecordNotFound)
	controller := StoreController{Storeservice: serviceMock}

	controller.Update(c)
	assert.Equal(t, 404, c.Writer.Status())
}

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

func TestStoreDeleteOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().Delete(gomock.Any()).Return(entity.Store{}, nil)
	controller := StoreController{Storeservice: serviceMock}

	controller.Delete(c)
	assert.Equal(t, 204, c.Writer.Status())
}

func TestStoreDeleteNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().Delete(gomock.Any()).Return(entity.Store{}, gorm.ErrRecordNotFound)
	controller := StoreController{Storeservice: serviceMock}

	controller.Delete(c)
	assert.Equal(t, 404, c.Writer.Status())
}

func TestStoreDeleteConflict(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().Delete(gomock.Any()).Return(entity.Store{}, ErrConflict)
	controller := StoreController{Storeservice: serviceMock}

	controller.Delete(c)
	assert.Equal(t, 409, c.Writer.Status())
}

func TestStorePetsList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().PetsList(gomock.Any()).Return(service.Pets{}, nil)
	controller := StoreController{Storeservice: serviceMock}

	controller.PetsList(c)
	assert.Equal(t, 200, c.Writer.Status())

}

func TestStorePetsListNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)

	serviceMock.EXPECT().PetsList(gomock.Any()).Return(service.Pets{}, gorm.ErrRecordNotFound)
	controller := StoreController{Storeservice: serviceMock}

	controller.PetsList(c)
	assert.Equal(t, 404, c.Writer.Status())
}
