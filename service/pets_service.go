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
	GetAll() (Pets, error)
	CreateModel(c *gin.Context) (Pet, error)
	GetByID(id string) (Pet, error)
}

func NewPetService() PetService {
	return &petService{}
}

type petService struct {
}

// GetAll is get all Pet
func (s petService) GetAll() (Pets, error) {
	db := db.GetDB()
	var l Pets
	var u []entity.Pet

	db.Find(&u)

	l.Pets = &u
	return l, nil
}

// CreateModel is create Pet model
func (s petService) CreateModel(c *gin.Context) (Pet, error) {
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
func (s petService) GetByID(id string) (Pet, error) {
	db := db.GetDB()
	var u Pet

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}
