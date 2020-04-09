package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nfv-aws/wcafe-api-controller/db"
	"github.com/nfv-aws/wcafe-api-controller/entity"
	"log"
)

// Service procides store's behavior
type StoreService struct{}

// User is alias of entity.Store struct
type Store entity.Store

// User is alias of entity.stores struct
type Stores entity.Stores

// GetAll is get all Store
func (s StoreService) GetAll() (Stores, error) {
	db := db.GetDB()
	var l Stores
	var u []entity.Store

	db.Find(&u)

	l.Stores = &u
	return l, nil
}

// CreateModel is create Store model
func (s StoreService) CreateModel(c *gin.Context) (Store, error) {
	db := db.GetDB()
	var u Store

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

//Update is update a Store
func (s StoreService) UpdateByID(id string, c *gin.Context) (Store, error) {
	db := db.GetDB()
	var u Store

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}
	if err := c.BindJSON(&u); err != nil {
		return u, err
	}
	if err := db.Save(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// GetByID is get a Store
func (s StoreService) GetByID(id string) (Store, error) {
	db := db.GetDB()
	var u Store

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}
