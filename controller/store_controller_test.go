package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/nfv-aws/wcafe-api-controller/entity"
	"github.com/nfv-aws/wcafe-api-controller/mocks"
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

func resetStoreTimeFields(store *entity.Store) *entity.Store {
	// 元の値は変えず、CreatedAtとUpdatedAtだけゼロ値にしたコピーを作る
	after := *store
	after.CreatedAt = time.Time{}
	after.UpdatedAt = time.Time{}
	return &after
}

func TestStoreList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// パラメータ生成
	li := gin.Param{"limit", "100"}
	of := gin.Param{"offset", "0"}
	params := gin.Params{li, of}
	// リクエスト生成
	req, _ := http.NewRequest("GET", "/stores", nil)
	// Contextセット
	var c *gin.Context
	c = &gin.Context{Request: req, Params: params}
	// w := httptest.NewRecorder()
	// c, _ := gin.CreateTestContext(w)

	serviceMock := mocks.NewMockStoreService(ctrl)
	s := []entity.Store{
		{Id: "sa5bafac-b35c-4852-82ca-b272cd79f2f3", Name: "Sano Shinichiro"},
		{Id: "sa5bafac-b35c-4852-82ca-b272cd79f2f5", Name: "Suzuki Chihiro"},
	}
	serviceMock.EXPECT().List(100, 0).Return(s, nil)
	controller := StoreController{Storeservice: serviceMock}
	log.Println("ここまでOk")
	controller.List(c)
	// log.Println(err)
	log.Println("ここはまだ")
	assert.Equal(t, http.StatusOK, c.Writer.Status())

	// var stores []entity.Store
	// err := json.Unmarshal([]byte(w.Body.String()), &stores)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// assert.Equal(t, s, stores)
}

func TestStoreGetOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().Get(gomock.Any()).Return(s, nil)
	controller := StoreController{Storeservice: serviceMock}
	controller.Get(c)
	assert.Equal(t, http.StatusOK, c.Writer.Status())

	var store entity.Store
	err := json.Unmarshal([]byte(w.Body.String()), &store)
	if err != nil {
		panic(err.Error())
	}
	assert.Equal(t, resetStoreTimeFields(&s), resetStoreTimeFields(&store))
}

func TestStoreGetNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)

	serviceMock.EXPECT().Get(gomock.Any()).Return(entity.Store{}, gorm.ErrRecordNotFound)
	controller := StoreController{Storeservice: serviceMock}

	controller.Get(c)
	assert.Equal(t, http.StatusNotFound, c.Writer.Status())
}

func TestStoreCreateOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().Create(c).Return(s, nil)
	controller := StoreController{Storeservice: serviceMock}
	controller.Create(c)
	assert.Equal(t, http.StatusCreated, c.Writer.Status())

	var store entity.Store
	err := json.Unmarshal([]byte(w.Body.String()), &store)
	if err != nil {
		panic(err.Error())
	}
	assert.Equal(t, resetStoreTimeFields(&s), resetStoreTimeFields(&store))
}

func TestStoreCreateInvalidAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().Create(c).Return(entity.Store{}, ErrInvalidAddress)
	controller := StoreController{Storeservice: serviceMock}

	controller.Create(c)
	assert.Equal(t, http.StatusNotFound, c.Writer.Status())
}

func TestStoreCreateBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().Create(c).Return(entity.Store{}, ErrBadRequest)
	controller := StoreController{Storeservice: serviceMock}

	controller.Create(c)
	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
}

func TestStoreUpdateOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().Update(gomock.Any(), c).Return(s, nil)
	controller := StoreController{Storeservice: serviceMock}
	controller.Update(c)
	assert.Equal(t, http.StatusOK, c.Writer.Status())

	var store entity.Store
	err := json.Unmarshal([]byte(w.Body.String()), &store)
	if err != nil {
		panic(err.Error())
	}
	assert.Equal(t, resetStoreTimeFields(&s), resetStoreTimeFields(&store))
}

func TestStoreUpdataNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)

	serviceMock.EXPECT().Update(gomock.Any(), c).Return(entity.Store{}, gorm.ErrRecordNotFound)
	controller := StoreController{Storeservice: serviceMock}

	controller.Update(c)
	assert.Equal(t, http.StatusNotFound, c.Writer.Status())
}

func TestStoreUpdateBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().Update(gomock.Any(), c).Return(entity.Store{}, ErrBadRequest)
	controller := StoreController{Storeservice: serviceMock}

	controller.Update(c)
	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
}

func TestStoreDeleteOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().Delete(gomock.Any()).Return(entity.Store{}, nil)
	controller := StoreController{Storeservice: serviceMock}

	controller.Delete(c)
	assert.Equal(t, http.StatusNoContent, c.Writer.Status())
}

func TestStoreDeleteNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().Delete(gomock.Any()).Return(entity.Store{}, gorm.ErrRecordNotFound)
	controller := StoreController{Storeservice: serviceMock}

	controller.Delete(c)
	assert.Equal(t, http.StatusNotFound, c.Writer.Status())
}

func TestStoreDeleteConflict(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().Delete(gomock.Any()).Return(entity.Store{}, ErrConflict)
	controller := StoreController{Storeservice: serviceMock}

	controller.Delete(c)
	assert.Equal(t, http.StatusConflict, c.Writer.Status())
}

func TestStorePetsList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	p := []entity.Pet{
		{Id: "pa5bafac-b35c-4852-82ca-b272cd79f2f3", Species: "Dog"},
		{Id: "pa5bafac-b35c-4852-82ca-b272cd79f2f5", Species: "Cat"},
	}

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().PetsList(gomock.Any()).Return(p, nil)
	controller := StoreController{Storeservice: serviceMock}
	controller.PetsList(c)
	assert.Equal(t, http.StatusOK, c.Writer.Status())

	var pets []entity.Pet
	err := json.Unmarshal([]byte(w.Body.String()), &pets)
	if err != nil {
		panic(err.Error())
	}
	assert.Equal(t, p, pets)
}

func TestStorePetsListEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	serviceMock := mocks.NewMockStoreService(ctrl)
	serviceMock.EXPECT().PetsList(gomock.Any()).Return([]entity.Pet{}, nil)
	controller := StoreController{Storeservice: serviceMock}
	controller.PetsList(c)
	assert.Equal(t, http.StatusOK, c.Writer.Status())

	var pets []entity.Pet
	err := json.Unmarshal([]byte(w.Body.String()), &pets)
	if err != nil {
		panic(err.Error())
	}
	assert.Equal(t, []entity.Pet{}, pets)
}

func TestStorePetsListNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockStoreService(ctrl)

	serviceMock.EXPECT().PetsList(gomock.Any()).Return([]entity.Pet{}, gorm.ErrRecordNotFound)
	controller := StoreController{Storeservice: serviceMock}

	controller.PetsList(c)
	assert.Equal(t, http.StatusNotFound, c.Writer.Status())
}
