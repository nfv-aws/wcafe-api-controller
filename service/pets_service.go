package service

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/nfv-aws/wcafe-api-controller/config"
	"github.com/nfv-aws/wcafe-api-controller/entity"
	"github.com/nfv-aws/wcafe-api-controller/internal"
)

var (
	pets_svc       *sqs.SQS
	aws_region     string
	pets_queue_url string
)

// User is alias of entity.Pet struct
type Pet entity.Pet

// Service procides pet's behavior
type PetService interface {
	List() ([]entity.Pet, error)
	Create(c *gin.Context) (entity.Pet, error)
	Get(id string) (entity.Pet, error)
	Update(id string, c *gin.Context) (entity.Pet, error)
	Delete(id string) (entity.Pet, error)
}

//type PetService struct{}
type petService struct {
	petRepository entity.PetRepository
}

func NewPetService(db entity.PetRepository) PetService {
	return &petService{petRepository: db}
}

func Pets_Init() *sqs.SQS {
	config.Configure()
	aws_region = config.C.SQS.Region
	pets_queue_url = config.C.SQS.Pets_Queue_Url
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(aws_region)}))
	pets_svc := sqs.New(sess)
	return pets_svc
}

// List is get all Pet
func (s petService) List() ([]entity.Pet, error) {
	var u []entity.Pet
	pr := s.petRepository

	u, err := pr.Find()
	if err != nil {
		return u, err
	}
	return u, nil
}

// Create is create Pet model
func (s petService) Create(c *gin.Context) (entity.Pet, error) {
	pr := s.petRepository
	var u entity.Pet

	//UUID生成
	id, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		return u, err
	}
	u.Id = id.String()

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	// SQSに接続
	pets_svc := Pets_Init()
	result, err := pets_svc.SendMessage(&sqs.SendMessageInput{
		MessageBody:  aws.String(u.Id),
		QueueUrl:     aws.String(pets_queue_url),
		DelaySeconds: aws.Int64(10),
	})
	if err != nil {
		log.Println("Pet SendMessage Error")
		return u, err
	} else {
		log.Println("Pet SendMessage Success", *result.MessageId)
	}

	// DBに登録
	u.Status = "PENDING_CREATE"
	u.CreatedAt = internal.JstTime()

	u, err = pr.Create(u)
	if err != nil {
		return u, err
	}
	return u, nil
}

// GetByID is get a Pet
func (s petService) Get(id string) (entity.Pet, error) {
	pr := s.petRepository
	var u entity.Pet

	u, err := pr.Get(id)
	if err != nil {
		return u, err
	}
	return u, nil
}

// Update is modify pet
func (s petService) Update(id string, c *gin.Context) (entity.Pet, error) {
	pr := s.petRepository
	var u entity.Pet

	u, err := pr.Get(id)
	if err != nil {
		return u, err
	}

	// 取得したPet情報にUpdateする内容をBind
	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	// 更新日時を上書き
	u.UpdatedAt = internal.JstTime()

	u, err = pr.Update(id, u)
	if err != nil {
		return u, err
	}

	return u, nil
}

//  Delete is delete a pet
func (s petService) Delete(id string) (entity.Pet, error) {
	pr := s.petRepository
	var u entity.Pet

	// 指定したIDが存在するか確認
	u, err := pr.Get(id)
	if err != nil {
		return u, err
	}

	// 削除
	u, err = pr.Delete(id)
	if err != nil {
		return u, err
	}

	return u, nil
}
