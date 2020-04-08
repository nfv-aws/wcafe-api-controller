package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/nfv-aws/wcafe-api-controller/mocks"
	"github.com/nfv-aws/wcafe-api-controller/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http/httptest"
	"testing"
)

/*
type MockServiceInterface struct {
}

func (_m *MockServiceInterface) GetAll() (service.Pets, error) {
	return service.Pets{}, nil
}

func (_m *MockServiceInterface) CreateModel(c *gin.Context) (service.Pet, error) {
	return service.Pet{
		Id:   "id",
		Name: "Name",
	}, nil
}

func (_m *MockServiceInterface) GetByID(id string) (service.Pet, error) {
	return service.Pet{Name: "test"}, nil
}
*/
type ControllerSuite struct {
	suite.Suite
	controller PetController
	service    service.PetService
}

func (s *ControllerSuite) SetupTest() {

}

func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(ControllerSuite))
}

func (s *ControllerSuite) TestIndex() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	serviceMock := mocks.NewMockPetService(ctrl)
	serviceMock.EXPECT().GetAll().Return(service.Pets{}, nil)
	s.controller = PetController{Service: serviceMock}

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	s.controller.Index(c)
	assert.Equal(s.T(), 200, c.Writer.Status())
}
