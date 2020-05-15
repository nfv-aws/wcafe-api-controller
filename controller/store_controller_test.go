package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/nfv-aws/wcafe-api-controller/mocks"
	"github.com/nfv-aws/wcafe-api-controller/service"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

var (
	//	ErrRecordNotFound = errors.New("record not found")
	ErrConflict = errors.New("Error 1451: Cannot delete or update a parent row: a foreign key constraint fails (`wcafe`.`pets`, CONSTRAINT `pets_store_id_stores_id_foreign` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT)")
)

func TestStoreList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().List().Return(service.Stores{}, nil)
	controller := StoreController{Service: serviceMock}

	controller.List(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestStoreGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().Get(gomock.Any()).Return(service.Store{}, nil)
	controller := StoreController{Service: serviceMock}

	controller.Get(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestStoreCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().Create(c).Return(service.Store{}, nil)
	controller := StoreController{Service: serviceMock}

	controller.Create(c)
	assert.Equal(t, 201, c.Writer.Status())
}

func TestStoreUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().Update(gomock.Any(), c).Return(service.Store{}, nil)
	controller := StoreController{Service: serviceMock}

	controller.Update(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestStoreDeleteOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().Delete(gomock.Any()).Return(service.Store{}, nil)
	controller := StoreController{Service: serviceMock}

	controller.Delete(c)
	assert.Equal(t, 204, c.Writer.Status())
}

// *** ToDo wcafe-103 or wcafe-118 ***
//func TestStoreDeleteNotFound(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()

//	c, _ := gin.CreateTestContext(httptest.NewRecorder())

//	log.Println(ErrRecordNotFound)
//	serviceMock := mocks.NewMockStoreService(ctrl)
//	serviceMock.EXPECT().Delete(gomock.Any()).Return(service.Store{}, ErrRecordNotFound)
//	controller := StoreController{Service: serviceMock}

//	controller.Delete(c)
//	assert.Equal(t, 404, c.Writer.Status())
//}

func TestStoreDeleteConflict(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().Delete(gomock.Any()).Return(service.Store{}, ErrConflict)
	controller := StoreController{Service: serviceMock}

	controller.Delete(c)
	assert.Equal(t, 409, c.Writer.Status())
}
