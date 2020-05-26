package service

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/nfv-aws/wcafe-api-controller/config"
	"github.com/nfv-aws/wcafe-api-controller/db"
	"github.com/nfv-aws/wcafe-api-controller/entity"
)

var (
	stores_svc       *sqs.SQS
	stores_queue_url string
)

func StoresInit() *sqs.SQS {
	config.Configure()
	aws_region = config.C.SQS.Region
	stores_queue_url = config.C.SQS.Stores_Queue_Url
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(aws_region)}))
	stores_svc := sqs.New(sess)
	return stores_svc
}

// User is alias of entity.Store struct
type Store entity.Store

// User is alias of entity.stores struct
type Stores entity.Stores

// Service procides store's behavior
type StoreService interface {
	List() ([]entity.Store, error)
	Create(c *gin.Context) (entity.Store, error)
	Get(id string) (entity.Store, error)
	Update(id string, c *gin.Context) (entity.Store, error)
	Delete(id string) (entity.Store, error)
}

func NewStoreService() StoreService {
	return &storeService{}
}

type storeService struct{}

// List is get all Store
func (s storeService) List() ([]entity.Store, error) {
	db := db.GetDB()
	var u []entity.Store

	db.Find(&u)

	return u, nil
}

// Create is create Store model
func (s storeService) Create(c *gin.Context) (entity.Store, error) {
	db := db.GetDB()
	var u entity.Store

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
	stores_svc := StoresInit()
	result, err := stores_svc.SendMessage(&sqs.SendMessageInput{
		MessageBody:  aws.String(u.Id),
		QueueUrl:     aws.String(stores_queue_url),
		DelaySeconds: aws.Int64(10),
	})
	if err != nil {
		log.Println("Store SendMessage Error")
		return u, err
	} else {
		log.Println("Store Success", *result.MessageId)
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// Get is get a Store
func (s storeService) Get(id string) (entity.Store, error) {

	db := db.GetDB()
	var u entity.Store

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// Update is update Store
func (s storeService) Update(id string, c *gin.Context) (entity.Store, error) {
	db := db.GetDB()
	var u, st entity.Store

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	//作成日・更新日を取得
	if err := db.Where("id = ?", id).First(&st).Error; err != nil {
		return u, err
	}
	u.CreatedAt = st.CreatedAt
	u.UpdatedAt = time.Now()

	if err := db.Table("stores").Where("id = ?", id).Updates(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// Delete is delete a Store
func (s storeService) Delete(id string) (entity.Store, error) {

	db := db.GetDB()
	var u entity.Store

	if err := db.Where("id = ?", id).Find(&u).Error; err != nil {
		return u, err
	}

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}
