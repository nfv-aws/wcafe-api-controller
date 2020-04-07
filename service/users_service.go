package service

import (
	"github.com/gin-gonic/gin"
	"github.com/nfv-aws/wcafe-api-controller/db"
	"github.com/nfv-aws/wcafe-api-controller/entity"
)

// Service procides user's behavior
type UserService struct{}

// User is alias of entity.user struct
type User entity.User

// User is alias of entity.users struct
type Users entity.Users

// GetAll is get all user
func (s UserService) GetAll() (Users, error) {
	db := db.GetDB()
	var l Users
	var u []entity.User

	db.Find(&u)

	l.Users = &u
	return l, nil
}

// CreateModel is create user model
func (s UserService) CreateModel(c *gin.Context) (User, error) {
	db := db.GetDB()
	var u User

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// GetByID is get a User
func (s UserService) GetByID(id string) (User, error) {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}
