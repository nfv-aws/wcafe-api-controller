package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nfv-aws/wcafe-api-controller/db"
	"github.com/nfv-aws/wcafe-api-controller/entity"
	"log"
)

// User is alias of entity.user struct
type User entity.User

// User is alias of entity.users struct
type Users entity.Users

// Service procides user's behavior
type UserService interface {
	List() (Users, error)
	Create(c *gin.Context) (User, error)
	Get(id string) (User, error)
	Update(id string, c *gin.Context) (User, error)
}

func NewUserService() UserService {
	return &userService{}
}

type userService struct {
}

// List is get all user
func (s userService) List() (Users, error) {
	db := db.GetDB()
	var l Users
	var u []entity.User

	db.Find(&u)

	l.Users = &u
	return l, nil
}

// Create is create user model
func (s userService) Create(c *gin.Context) (User, error) {
	db := db.GetDB()
	var u User

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
func (s userService) Get(id string) (User, error) {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// Update is update a User
func (s userService) Update(id string, c *gin.Context) (User, error) {
	db := db.GetDB()
	tmp := User{} //id格納用のUser
	tmp.Id = id
	u := tmp
	db.First(&u)
	c.BindJSON(&u)
	db.Model(&tmp).Update(&u)
	return u, nil
}
