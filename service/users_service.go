package service

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gopkg.in/go-playground/validator.v9"

	"github.com/nfv-aws/wcafe-api-controller/config"
	"github.com/nfv-aws/wcafe-api-controller/db"
	"github.com/nfv-aws/wcafe-api-controller/entity"
	"github.com/nfv-aws/wcafe-api-controller/internal"
)

//SQS用グローバル変数
var (
	users_svc       *sqs.SQS
	users_queue_url string
)

//SQS処理
func Users_Init() *sqs.SQS {
	config.Configure()
	aws_region = config.C.SQS.Region
	users_queue_url = config.C.SQS.Users_Queue_Url
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(aws_region)}))
	users_svc := sqs.New(sess)
	return users_svc
}

// User is alias of entity.user struct
type User entity.User

// User is alias of entity.users struct
type Users entity.Users

// Service procides user's behavior
type UserService interface {
	List() ([]entity.User, error)
	Create(c *gin.Context) (entity.User, error)
	Get(id string) (entity.User, error)
	Update(id string, c *gin.Context) (entity.User, error)
	Delete(id string) (User, error)
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

// List is get all user
func (s userService) List() ([]entity.User, error) {
	db := db.GetDB()
	var u []entity.User
	db.Find(&u)
	return u, nil
}

// Create is create user model
func (s userService) Create(c *gin.Context) (entity.User, error) {
	db := db.GetDB()
	var u entity.User

	//UUID生成
	id, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		return u, err
	}

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	//Email validation check
	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		return u, err
	}

	u.Id = id.String()
	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	//SQS処理呼び出し
	log.Println(u.Id)
	log.Println(users_queue_url)
	users_svc := Users_Init()
	result, err := users_svc.SendMessage(&sqs.SendMessageInput{
		MessageBody:  aws.String(u.Id),
		QueueUrl:     aws.String(users_queue_url),
		DelaySeconds: aws.Int64(10),
	})
	if err != nil {
		log.Println("User SendMessage Error", err)
	}
	log.Println("User Success", *result.MessageId)

	return u, nil
}

// Get is get a User
func (s userService) Get(id string) (entity.User, error) {
	db := db.GetDB()
	var u entity.User

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// Update is update a User
func (s userService) Update(id string, c *gin.Context) (entity.User, error) {
	db := db.GetDB()
	var u, ut entity.User

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	//作成日・更新日を取得
	if err := db.Where("id = ?", id).First(&ut).Error; err != nil {
		return u, err
	}

	//Email validation check
	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		return u, err
	}

	u.CreatedAt = ut.CreatedAt
	u.UpdatedAt = internal.JstTime()

	if err := db.Table("users").Where("id = ?", id).Updates(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

//  Delete is delete a pet
func (s userService) Delete(id string) (User, error) {
	db := db.GetDB()
	var u User
	//該当データの有無を確認
	if err := db.Where("id = ?", id).Find(&u).Error; err != nil {
		return u, err
	}
	//該当データを削除
	if err := db.Table("users").Where("id = ?", id).Delete(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}
