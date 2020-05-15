package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nfv-aws/wcafe-api-controller/config"
	"github.com/nfv-aws/wcafe-api-controller/db"
	"github.com/nfv-aws/wcafe-api-controller/entity"
	"log"
	"time"
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
	List() (Stores, error)
	Create(c *gin.Context) (Store, error)
	Get(id string) (Store, error)
	Update(id string, c *gin.Context) (Store, error)
	Delete(id string) (Store, error)
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

	stores_svc := StoresInit()
	result, err := stores_svc.SendMessage(&sqs.SendMessageInput{
		MessageBody:  aws.String(u.Id),
		QueueUrl:     aws.String(stores_queue_url),
		DelaySeconds: aws.Int64(10),
	})
	if err != nil {
		log.Println("Store SendMessage Error", err)
	}
	log.Println("Store Success", *result.MessageId)

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

// Update is update Store
func (s storeService) Update(id string, c *gin.Context) (Store, error) {
	db := db.GetDB()
	var u, st Store

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
func (s storeService) Delete(id string) (Store, error) {

	db := db.GetDB()
	var u Store

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}
