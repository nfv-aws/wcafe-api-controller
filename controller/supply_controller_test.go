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
	su = entity.Supply{
		Id:    "cc5bafac-b35c-4852-82ca-b272cd79f2f3",
		Name:  "dog food",
		Type:  "food",
		Price: 500,
	}
)

func TestSupplyList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockSupplyService(ctrl)
	su := []entity.Supply{
		{Id: "cc5bafac-b35c-4852-82ca-b272cd79f2f3", Name: "dog food", Type: "food", Price: 500},
		{Id: "cs6befhj-b93c-5672-13ka-b272bh46f2f3", Name: "ball", Type: "toy", Price: 300},
	}
	serviceMock.EXPECT().List().Return(su, nil)
	controller := SupplyController{Supplyservice: serviceMock}

	controller.List(c)
	assert.Equal(t, 200, c.Writer.Status())
}
