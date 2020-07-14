package controller

import (
	"encoding/json"
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
	cl = entity.Clerk{
		Id:   "cc5bafac-b35c-4852-82ca-b272cd79f2f3",
		Name: "Sano Shinichiro",
	}
)

func TestClerkList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	serviceMock := mocks.NewMockClerkService(ctrl)
	cl := []entity.Clerk{
		{Id: "cc5bafac-b35c-4852-82ca-b272cd79f2f3", Name: "Sano Shinichiro"},
		{Id: "cc5bafac-b35c-4852-82ca-b272cd79f2f5", Name: "Suzuki Chihiro"},
	}
	serviceMock.EXPECT().List().Return(cl, nil)
	controller := ClerkController{Clerkservice: serviceMock}
	controller.List(c)
	assert.Equal(t, http.StatusOK, c.Writer.Status())

	var clerks []entity.Clerk
	err := json.Unmarshal([]byte(w.Body.String()), &clerks)
	if err != nil {
		panic(err.Error())
	}
	assert.Equal(t, cl, clerks)
}

func TestClerkGetOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	serviceMock := mocks.NewMockClerkService(ctrl)
	serviceMock.EXPECT().Get(gomock.Any()).Return(cl, nil)
	controller := ClerkController{Clerkservice: serviceMock}
	controller.Get(c)
	assert.Equal(t, http.StatusOK, c.Writer.Status())
}

func TestClerkGetNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockClerkService(ctrl)

	// TODO error
	serviceMock.EXPECT().Get(gomock.Any()).Return(entity.Clerk{}, gorm.ErrRecordNotFound)
	controller := ClerkController{Clerkservice: serviceMock}

	controller.Get(c)
	assert.Equal(t, http.StatusNotFound, c.Writer.Status())
}

func TestClerkCreateOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	serviceMock := mocks.NewMockClerkService(ctrl)
	serviceMock.EXPECT().Create(c).Return(cl, nil)
	controller := ClerkController{Clerkservice: serviceMock}
	controller.Create(c)
	assert.Equal(t, http.StatusCreated, c.Writer.Status())

	var clerk entity.Clerk
	err := json.Unmarshal([]byte(w.Body.String()), &clerk)
	if err != nil {
		panic(err.Error())
	}
	assert.Equal(t, cl, clerk)
}

func TestClerkCreateBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockClerkService(ctrl)
	serviceMock.EXPECT().Create(c).Return(entity.Clerk{}, ErrBadRequest)
	controller := ClerkController{Clerkservice: serviceMock}

	controller.Create(c)
	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
}

func TestClerkDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockClerkService(ctrl)
	serviceMock.EXPECT().Delete(gomock.Any()).Return(entity.Clerk{}, nil)
	controller := ClerkController{Clerkservice: serviceMock}

	controller.Delete(c)
	assert.Equal(t, http.StatusNoContent, c.Writer.Status())
}

func TestClerkDeleteNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockClerkService(ctrl)
	serviceMock.EXPECT().Delete(gomock.Any()).Return(entity.Clerk{}, gorm.ErrRecordNotFound)
	controller := ClerkController{Clerkservice: serviceMock}

	controller.Delete(c)
	assert.Equal(t, http.StatusNotFound, c.Writer.Status())
}
