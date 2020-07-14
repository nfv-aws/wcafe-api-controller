package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gopkg.in/go-playground/validator.v9"

	"github.com/nfv-aws/wcafe-api-controller/config"
	"github.com/nfv-aws/wcafe-api-controller/entity"
	"github.com/nfv-aws/wcafe-api-controller/internal"
)

var (
	stores_svc       *sqs.SQS
	stores_queue_url string
)

// User is alias of entity.Store struct
type Store entity.Store

// Service procides store's behavior
type StoreService interface {
	List(limit int, offset int) ([]entity.Store, error)
	Create(c *gin.Context) (entity.Store, error)
	Get(id string) (entity.Store, error)
	Update(id string, c *gin.Context) (entity.Store, error)
	Delete(id string) (entity.Store, error)
	PetsList(id string) ([]entity.Pet, error)
}

type storeService struct {
	storeRepository entity.StoreRepository
}

func NewStoreService(db entity.StoreRepository) StoreService {
	return &storeService{storeRepository: db}
}

func StoresInit() *sqs.SQS {
	log.Debug().Caller().Msg("stores init")
	config.Configure()
	aws_region = config.C.SQS.Region
	stores_queue_url = config.C.SQS.Stores_Queue_Url
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(aws_region)}))
	stores_svc := sqs.New(sess)
	return stores_svc
}

// List is get all Store
func (s storeService) List(limit int, offset int) ([]entity.Store, error) {
	log.Debug().Caller().Msg("stores list")
	sr := s.storeRepository
	var u []entity.Store

	u, err := sr.Find(limit, offset)
	if err != nil {
		return u, err
	}
	return u, nil
}

// Create is create Store model
func (s storeService) Create(c *gin.Context) (entity.Store, error) {
	log.Debug().Caller().Msg("stores create")
	sr := s.storeRepository
	var u entity.Store

	//UUID生成
	id, err := uuid.NewRandom()
	if err != nil {
		log.Error().Caller().Err(err).Send()
		return u, err
	}
	u.Id = id.String()

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	validate := validator.New()
	if err := validate.Struct(u); err != nil {
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
		log.Error().Caller().Msg("Store SendMessage Error")
		return u, err
	} else {
		log.Info().Caller().Msg("Store Success:" + string(*result.MessageId))
	}

	u.Status = "PENDING_CREATE"
	u.CreatedAt = internal.JstTime()

	u, err = sr.Create(u)
	if err != nil {
		return u, err
	}

	return u, nil
}

// Get is get a Store
func (s storeService) Get(id string) (entity.Store, error) {
	log.Debug().Caller().Msg("stores get")
	sr := s.storeRepository
	var u entity.Store

	u, err := sr.Get(id)
	if err != nil {
		return u, err
	}

	return u, nil
}

// Update is update Store
func (s storeService) Update(id string, c *gin.Context) (entity.Store, error) {
	log.Debug().Caller().Msg("stores update")
	sr := s.storeRepository
	var u entity.Store

	u, err := sr.Get(id)
	if err != nil {
		return u, err
	}

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		return u, err
	}

	u.UpdatedAt = internal.JstTime()

	u, err = sr.Update(id, u)
	if err != nil {
		return u, err
	}

	return u, nil
}

// Delete is delete a Store
func (s storeService) Delete(id string) (entity.Store, error) {
	log.Debug().Caller().Msg("stores delete")
	sr := s.storeRepository
	var u entity.Store

	// 指定したIDが存在するか確認
	u, err := sr.Get(id)
	if err != nil {
		return u, err
	}

	// 削除
	u, err = sr.Delete(id)
	if err != nil {
		return u, err
	}

	return u, nil
}

// Get is get a Store & List is get all Pets
func (s storeService) PetsList(id string) ([]entity.Pet, error) {
	log.Debug().Caller().Msg("stores pets list")
	sr := s.storeRepository
	var p []entity.Pet

	p, err := sr.PetsList(id)
	if err != nil {
		return p, err
	}

	return p, nil
}
