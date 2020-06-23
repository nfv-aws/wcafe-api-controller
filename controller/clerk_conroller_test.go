package controller

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/nfv-aws/wcafe-api-controller/entity"
	"github.com/nfv-aws/wcafe-api-controller/mocks"
)

var (
	cl = entity.Clerk{
		NameId: "cc5bafac-b35c-4852-82ca-b272cd79f2f3",
		Name:   "Sano Shinichiro",
	}
)

func TestClerkList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockClerkService(ctrl)
	cl := []entity.Clerk{
		{NameId: "cc5bafac-b35c-4852-82ca-b272cd79f2f3", Name: "Sano Shinichiro"},
		{NameId: "cc5bafac-b35c-4852-82ca-b272cd79f2f5", Name: "Suzuki Chihiro"},
	}
	serviceMock.EXPECT().List().Return(cl, nil)
	controller := ClerkController{Clerkservice: serviceMock}

	controller.List(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestClerkCreateOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockClerkService(ctrl)

	serviceMock.EXPECT().Create(c).Return(cl, nil)
	controller := ClerkController{Clerkservice: serviceMock}

	controller.Create(c)
	assert.Equal(t, 201, c.Writer.Status())
}

func TestClerkCreateBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockClerkService(ctrl)
	serviceMock.EXPECT().Create(c).Return(entity.Clerk{}, ErrBadRequest)
	controller := ClerkController{Clerkservice: serviceMock}

	controller.Create(c)
	assert.Equal(t, 400, c.Writer.Status())
}
