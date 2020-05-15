package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nfv-aws/wcafe-api-controller/db"
	"github.com/nfv-aws/wcafe-api-controller/entity"
	"log"
	"time"
)

// User is alias of entity.user struct
type User entity.User

// User is alias of entity.users struct
type Users entity.Users

// Service procides user's behavior
type UserService interface {
	List() ([]entity.User, error)
	Create(c *gin.Context) (entity.User, error)
	Get(id string) (entity.User, error)
	Update(id string, c *gin.Context) (entity.User, error)
	Delete(id string) (User, error)
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

// List is get all user
func (s userService) List() ([]entity.User, error) {
	db := db.GetDB()
	var u []entity.User
	db.Find(&u)
	return u, nil
}

// Create is create user model
func (s userService) Create(c *gin.Context) (entity.User, error) {
	db := db.GetDB()
	var u entity.User

	//UUID生成
	id, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		return u, err
	}

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	u.Id = id.String()
	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// Get is get a User
func (s userService) Get(id string) (entity.User, error) {
	db := db.GetDB()
	var u entity.User

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// Update is update a User
func (s userService) Update(id string, c *gin.Context) (entity.User, error) {
	db := db.GetDB()
	var u, ut entity.User

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	//作成日・更新日を取得
	if err := db.Where("id = ?", id).First(&ut).Error; err != nil {
		return u, err
	}

	u.CreatedAt = ut.CreatedAt
	u.UpdatedAt = time.Now()

	if err := db.Table("users").Where("id = ?", id).Updates(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

//  Delete is delete a pet
func (s userService) Delete(id string) (User, error) {
	db := db.GetDB()
	var u User

	if err := db.Table("users").Where("id = ?", id).Delete(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}
