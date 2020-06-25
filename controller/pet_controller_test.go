package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/nfv-aws/wcafe-api-controller/entity"
	"github.com/nfv-aws/wcafe-api-controller/mocks"
)

var (
	p = entity.Pet{
		Id:        "74684838-a5d9-47d8-91a4-ff63ce802763",
		Species:   "Canine",
		Name:      "Shiba-inu",
		Age:       1,
		StoreId:   "a103c7e0-b560-4b01-9628-24553f136a6f",
		CreatedAt: ct,
		UpdatedAt: ut,
		Status:    "PENDING_CREATE",
	}
	ErrInvalidAddress = errors.New("InvalidAddress: The address https://sqs.ap-northeast-1.amazonaws.com/ is not valid for this endpoint.")
)

func TestPetList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockPetService(ctrl)
	p := []entity.Pet{
		{Id: "sa5bafac-b35c-4852-82ca-b272cd79f2f3", Name: "Sano Shinichiro"},
		{Id: "sa5bafac-b35c-4852-82ca-b272cd79f2f5", Name: "Suzuki Chihiro"},
	}

	serviceMock.EXPECT().List().Return(p, nil)
	controller := PetController{Petservice: serviceMock}

	controller.List(c)
	assert.Equal(t, http.StatusOK, c.Writer.Status())
}

func TestPetGetOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockPetService(ctrl)
	serviceMock.EXPECT().Get(gomock.Any()).Return(p, nil)
	controller := PetController{Petservice: serviceMock}

	controller.Get(c)
	assert.Equal(t, http.StatusOK, c.Writer.Status())
}

func TestPetGetNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockPetService(ctrl)

	serviceMock.EXPECT().Get(gomock.Any()).Return(entity.Pet{}, gorm.ErrRecordNotFound)
	controller := PetController{Petservice: serviceMock}

	controller.Get(c)
	assert.Equal(t, http.StatusNotFound, c.Writer.Status())
}

func TestPetCreateOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockPetService(ctrl)
	serviceMock.EXPECT().Create(c).Return(p, nil)
	controller := PetController{Petservice: serviceMock}

	controller.Create(c)
	assert.Equal(t, http.StatusCreated, c.Writer.Status())
}

func TestPetCreateInvalidAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockPetService(ctrl)
	serviceMock.EXPECT().Create(c).Return(entity.Pet{}, ErrInvalidAddress)
	controller := PetController{Petservice: serviceMock}

	controller.Create(c)
	assert.Equal(t, http.StatusNotFound, c.Writer.Status())
}

func TestPetCreateBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockPetService(ctrl)
	serviceMock.EXPECT().Create(c).Return(entity.Pet{}, ErrBadRequest)
	controller := PetController{Petservice: serviceMock}

	controller.Create(c)
	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
}

func TestPetUpdateOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockPetService(ctrl)
	serviceMock.EXPECT().Update(gomock.Any(), c).Return(p, nil)
	controller := PetController{Petservice: serviceMock}

	controller.Update(c)
	assert.Equal(t, http.StatusOK, c.Writer.Status())
}

func TestPetUpdataNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockPetService(ctrl)

	serviceMock.EXPECT().Update(gomock.Any(), c).Return(entity.Pet{}, gorm.ErrRecordNotFound)
	controller := PetController{Petservice: serviceMock}

	controller.Update(c)
	assert.Equal(t, http.StatusNotFound, c.Writer.Status())
}

func TestPetUpdateBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockPetService(ctrl)
	serviceMock.EXPECT().Update(gomock.Any(), c).Return(entity.Pet{}, ErrBadRequest)
	controller := PetController{Petservice: serviceMock}

	controller.Update(c)
	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
}

func TestPetDeleteOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockPetService(ctrl)
	serviceMock.EXPECT().Delete(gomock.Any()).Return(entity.Pet{}, nil)
	controller := PetController{Petservice: serviceMock}

	controller.Delete(c)
	assert.Equal(t, http.StatusNoContent, c.Writer.Status())
}

func TestPetDeleteNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockPetService(ctrl)
	serviceMock.EXPECT().Delete(gomock.Any()).Return(entity.Pet{}, gorm.ErrRecordNotFound)
	controller := PetController{Petservice: serviceMock}

	controller.Delete(c)
	assert.Equal(t, http.StatusNotFound, c.Writer.Status())
}
