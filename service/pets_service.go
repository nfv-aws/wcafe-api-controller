package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nfv-aws/wcafe-api-controller/db"
	"github.com/nfv-aws/wcafe-api-controller/entity"
)

// Service procides pet's behavior
type PetService struct{}

// User is alias of entity.Pet struct
type Pet entity.Pet

// User is alias of entity.Pets struct
type Pets entity.Pets

// GetAll is get all Pet
func (s PetService) GetAll() (Pets, error) {
	db := db.GetDB()
	var l Pets
	var u []entity.Pet

	db.Find(&u)

	l.Pets = &u
	return l, nil
}

// CreateModel is create Pet model
func (s PetService) CreateModel(c *gin.Context) (Pet, error) {
	db := db.GetDB()
	var u Pet

	//UUID生成
	id, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
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

// GetByID is get a Pet
func (s PetService) GetByID(id string) (Pet, error) {
	db := db.GetDB()
	var u Pet

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}
