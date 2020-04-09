package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nfv-aws/wcafe-api-controller/db"
	"github.com/nfv-aws/wcafe-api-controller/entity"
	"log"
)

// User is alias of entity.Store struct
type Store entity.Store

// User is alias of entity.stores struct
type Stores entity.Stores

// Service procides store's behavior
type StoreService interface {
	List() (Stores, error)
	Create(c *gin.Context) (Store, error)
	Get(id string) (Store, error)
	Update(id string, c *gin.Context) (Store, error)
}

func NewStoreService() StoreService {
	return &storeService{}
}

type storeService struct {
}

// List is get all Store
func (s storeService) List() (Stores, error) {
	db := db.GetDB()
	var l Stores
	var u []entity.Store

	db.Find(&u)

	l.Stores = &u
	return l, nil
}

// Create is create Store model
func (s storeService) Create(c *gin.Context) (Store, error) {
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

// Get is get a Store
func (s storeService) Get(id string) (Store, error) {

	db := db.GetDB()
	var u Store

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (s storeService) Update(id string, c *gin.Context) (Store, error) {
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
