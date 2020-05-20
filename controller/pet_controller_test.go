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
	ErrInvalidAddress = errors.New("InvalidAddress: The address https://sqs.ap-northeast-1.amazonaws.com/ is not valid for this endpoint.")
)

func TestPetList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockPetService(ctrl)
	serviceMock.EXPECT().List().Return(service.Pets{}, nil)
	controller := PetController{Service: serviceMock}

	controller.List(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestPetGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockPetService(ctrl)
	serviceMock.EXPECT().Get(gomock.Any()).Return(service.Pet{}, nil)
	controller := PetController{Service: serviceMock}

	controller.Get(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestPetCreateOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockPetService(ctrl)
	serviceMock.EXPECT().Create(c).Return(service.Pet{}, nil)
	controller := PetController{Service: serviceMock}

	controller.Create(c)
	assert.Equal(t, 201, c.Writer.Status())
}

func TestPetCreateInvalidAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockPetService(ctrl)
	serviceMock.EXPECT().Create(c).Return(service.Pet{}, ErrInvalidAddress)
	controller := PetController{Service: serviceMock}

	controller.Create(c)
	assert.Equal(t, 404, c.Writer.Status())
}

func TestPetCreateBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockPetService(ctrl)
	serviceMock.EXPECT().Create(c).Return(service.Pet{}, ErrBadRequest)
	controller := PetController{Service: serviceMock}

	controller.Create(c)
	assert.Equal(t, 400, c.Writer.Status())
}

func TestPetUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockPetService(ctrl)
	serviceMock.EXPECT().Update(gomock.Any(), c).Return(service.Pet{}, nil)
	controller := PetController{Service: serviceMock}

	controller.Update(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestPetDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockPetService(ctrl)
	serviceMock.EXPECT().Delete(gomock.Any()).Return(service.Pet{}, nil)
	controller := PetController{Service: serviceMock}

	controller.Delete(c)
	assert.Equal(t, 204, c.Writer.Status())
}
