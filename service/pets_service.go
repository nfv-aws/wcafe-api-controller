package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
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
	u.Status = "PENDING_CREATE"
	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	AWS_REGION := "ap-northeast-1"
	sess := session.Must(session.NewSession(&aws.Config{Region: &AWS_REGION}))

	svc := sqs.New(sess)
	QUEUE_URL := "https://sqs.ap-northeast-1.amazonaws.com/292364111734/REST-petstest"

	result, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageBody:  aws.String(u.Id),
		QueueUrl:     &QUEUE_URL,
		DelaySeconds: aws.Int64(10),
	})
	if err != nil {
		log.Println("SendMessage Error", err)
	}
	log.Println("Success", *result.MessageId)

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
