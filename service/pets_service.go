package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nfv-aws/wcafe-api-controller/db"
	"github.com/nfv-aws/wcafe-api-controller/entity"
	"log"
)

// User is alias of entity.Pet struct
type Pet entity.Pet

// User is alias of entity.Pets struct
type Pets entity.Pets

// Service procides pet's behavior
//type PetService struct{}
type PetService interface {
	List() (Pets, error)
	Create(c *gin.Context) (Pet, error)
	Get(id string) (Pet, error)
	Update(id string, c *gin.Context) (Pet, error)
}

func NewPetService() PetService {
	return &petService{}
}

type petService struct {
}

// List is get all Pet
func (s petService) List() (Pets, error) {
	db := db.GetDB()
	var l Pets
	var u []entity.Pet

	db.Find(&u)

	l.Pets = &u
	return l, nil
}

// Create is create Pet model
func (s petService) Create(c *gin.Context) (Pet, error) {
	db := db.GetDB()
	var u Pet

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

// GetByID is get a Pet
func (s petService) Get(id string) (Pet, error) {
	db := db.GetDB()
	var u Pet

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// Update is modify pet
func (s petService) Update(id string, c *gin.Context) (Pet, error) {
	db := db.GetDB()
	var u Pet

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	if err := db.Table("pets").Where("id = ?", id).Updates(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}
