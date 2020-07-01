package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/nfv-aws/wcafe-api-controller/config"
	"github.com/nfv-aws/wcafe-api-controller/entity"
	"github.com/nfv-aws/wcafe-api-controller/internal"
)

//SQS用グローバル変数
var (
	users_svc       *sqs.SQS
	users_queue_url string
)

// User is alias of entity.user struct
type User entity.User

// Service procides user's behavior
type UserService interface {
	List() ([]entity.User, error)
	Create(c *gin.Context) (entity.User, error)
	Get(id string) (entity.User, error)
	Update(id string, c *gin.Context) (entity.User, error)
	Delete(id string) (entity.User, error)
}

type userService struct {
	userRepository entity.UserRepository
}

func NewUserService(db entity.UserRepository) UserService {
	return &userService{userRepository: db}
}

//SQS処理
func Users_Init() *sqs.SQS {
	log.Debug().Caller().Msg("users init")
	config.Configure()
	aws_region = config.C.SQS.Region
	users_queue_url = config.C.SQS.Users_Queue_Url
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(aws_region)}))
	users_svc := sqs.New(sess)
	return users_svc
}

// List is get all user
func (s userService) List() ([]entity.User, error) {
	log.Debug().Caller().Msg("users list")
	var u []entity.User
	ur := s.userRepository
	u, err := ur.Find()
	if err != nil {
		return u, err
	}
	return u, nil
}

// Create is create user model
func (s userService) Create(c *gin.Context) (entity.User, error) {
	log.Debug().Caller().Msg("users create")
	ur := s.userRepository
	var u entity.User

	//UUID生成
	id, err := uuid.NewRandom()
	if err != nil {
		log.Error().Caller().Err(err).Send()
		return u, err
	}

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	//validation Check
	if err := entity.UserValidator(u); err != nil {
		return u, err
	}

	u.Id = id.String()

	//SQS処理呼び出し
	log.Info().Caller().Msg(u.Id)
	log.Info().Caller().Msg(users_queue_url)
	users_svc := Users_Init()
	result, err := users_svc.SendMessage(&sqs.SendMessageInput{
		MessageBody:  aws.String(u.Id),
		QueueUrl:     aws.String(users_queue_url),
		DelaySeconds: aws.Int64(10),
	})
	if err != nil {
		log.Error().Caller().Msg("User SendMessage Error")
	} else {
		log.Info().Caller().Msg("User Success:" + string(*result.MessageId))
	}

	u.CreatedAt = internal.JstTime()
	u, err = ur.Create(u)
	if err != nil {
		return u, err
	}

	return u, nil
}

// Get is get a User
func (s userService) Get(id string) (entity.User, error) {
	log.Debug().Caller().Msg("users get")
	ur := s.userRepository
	var u entity.User

	u, err := ur.Get(id)
	if err != nil {
		return u, err
	}

	return u, nil
}

// Update is update a User
func (s userService) Update(id string, c *gin.Context) (entity.User, error) {
	log.Debug().Caller().Msg("users update")
	ur := s.userRepository

	var u entity.User

	u, err := ur.Get(id)
	if err != nil {
		return u, err
	}

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	//validation Check
	if err := entity.UserValidator(u); err != nil {
		return u, err
	}

	u.UpdatedAt = internal.JstTime()

	u, err = ur.Update(id, u)
	if err != nil {
		return u, err
	}

	return u, nil
}

//  Delete is delete a pet
func (s userService) Delete(id string) (entity.User, error) {
	log.Debug().Caller().Msg("users delete")
	ur := s.userRepository
	var u entity.User

	// 指定したIDが存在するか確認
	u, err := ur.Get(id)
	if err != nil {
		return u, err
	}

	// 削除
	u, err = ur.Delete(id)
	if err != nil {
		return u, err
	}

	return u, nil
}
