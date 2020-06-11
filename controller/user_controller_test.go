package controller

import (
	"errors"
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
	u = entity.User{
		Id:        "ba5bafac-b35c-4852-82ca-b272cd79f2f3",
		Number:    1,
		Name:      "Sano Shinichiro",
		Address:   "Shinagawa",
		Email:     "test@example.com",
		CreatedAt: ct,
		UpdatedAt: ut,
	}
	ct, ut        = time.Now(), time.Now()
	ErrBadRequest = errors.New("json: cannot unmarshal number into Go struct field User.email of type string")
)

func TestUserList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockUserService(ctrl)
	u := []entity.User{
		{Id: "ba5bafac-b35c-4852-82ca-b272cd79f2f3", Name: "Sano Shinichiro"},
		{Id: "ba5bafac-b35c-4852-82ca-b272cd79f2f5", Name: "Suzuki Chihiro"},
	}
	serviceMock.EXPECT().List().Return(u, nil)
	controller := UserController{Userservice: serviceMock}

	controller.List(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestUserGetOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockUserService(ctrl)

	serviceMock.EXPECT().Get(gomock.Any()).Return(u, nil)
	controller := UserController{Userservice: serviceMock}

	controller.Get(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestUserGetNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockUserService(ctrl)

	// TODO error
	serviceMock.EXPECT().Get(gomock.Any()).Return(entity.User{}, gorm.ErrRecordNotFound)
	controller := UserController{Userservice: serviceMock}

	controller.Get(c)
	assert.Equal(t, 404, c.Writer.Status())
}

func TestUserCreateOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockUserService(ctrl)

	serviceMock.EXPECT().Create(c).Return(u, nil)
	controller := UserController{Userservice: serviceMock}

	controller.Create(c)
	assert.Equal(t, 201, c.Writer.Status())
}

func TestUserCreateBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockUserService(ctrl)

	serviceMock.EXPECT().Create(c).Return(entity.User{}, ErrBadRequest)
	controller := UserController{Userservice: serviceMock}

	controller.Create(c)
	assert.Equal(t, 400, c.Writer.Status())
}

func TestUserUpdateOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockUserService(ctrl)

	serviceMock.EXPECT().Update(gomock.Any(), c).Return(u, nil)
	controller := UserController{Userservice: serviceMock}

	controller.Update(c)
	assert.Equal(t, 200, c.Writer.Status())
}

// ToDo WCAF-118
func TestUserUpdateNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockUserService(ctrl)
	serviceMock.EXPECT().Update(gomock.Any(), c).Return(entity.User{}, gorm.ErrRecordNotFound)
	controller := UserController{Userservice: serviceMock}

	controller.Update(c)
	assert.Equal(t, 404, c.Writer.Status())
}

func TestUserUpdateBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockUserService(ctrl)
	serviceMock.EXPECT().Update(gomock.Any(), c).Return(entity.User{}, ErrBadRequest)
	controller := UserController{Userservice: serviceMock}

	controller.Update(c)
	assert.Equal(t, 400, c.Writer.Status())
}

func TestUserDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockUserService(ctrl)
	serviceMock.EXPECT().Delete(gomock.Any()).Return(entity.User{}, nil)
	controller := UserController{Userservice: serviceMock}

	controller.Delete(c)
	assert.Equal(t, 204, c.Writer.Status())
}

func TestUserDeleteNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockUserService(ctrl)
	serviceMock.EXPECT().Delete(gomock.Any()).Return(entity.User{}, gorm.ErrRecordNotFound)
	controller := UserController{Userservice: serviceMock}

	controller.Delete(c)
	assert.Equal(t, 404, c.Writer.Status())
}
