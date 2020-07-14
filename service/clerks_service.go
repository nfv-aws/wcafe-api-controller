package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/guregu/dynamo"
	"github.com/rs/zerolog/log"
	"gopkg.in/go-playground/validator.v9"

	"github.com/nfv-aws/wcafe-api-controller/config"
	"github.com/nfv-aws/wcafe-api-controller/entity"
)

//DynamoDB用グローバル変数
var (
	dynamodb *dynamo.DB
)

// DynamoDBに接続
func Dynamo_Init() *dynamo.DB {
	log.Debug().Caller().Msg("Dynamo_Init")
	config.Configure()
	aws_region = config.C.DynamoDB.Region
	dynamodb := dynamo.New(session.New(), &aws.Config{
		Region: aws.String(aws_region),
	})
	return dynamodb
}

// Clerks is alias of entity.clerk struct
type Clerk entity.Clerk

// Clerks is alias of entity.clerks struct
type Clerks entity.Clerks

// Service procides clerk's behavior
type ClerkService interface {
	List() ([]entity.Clerk, error)
	Create(c *gin.Context) (entity.Clerk, error)
	Get(id string) (entity.Clerk, error)
	Delete(id string) (entity.Clerk, error)
}

type clerkService struct{}

func NewClerkService() ClerkService {
	return &clerkService{}
}

// List is get all clerk
func (s clerkService) List() ([]entity.Clerk, error) {
	log.Debug().Caller().Msg("clerks list")
	dynamodb := Dynamo_Init()
	table := dynamodb.Table("clerks")

	var cl []entity.Clerk
	if err := table.Scan().All(&cl); err != nil {
		return cl, err
	}
	return cl, nil
}

// Create is create clerk model
func (s clerkService) Create(c *gin.Context) (entity.Clerk, error) {
	log.Debug().Caller().Msg("clerks create")
	dynamodb := Dynamo_Init()
	table := dynamodb.Table("clerks")
	var cl entity.Clerk

	//UUID生成
	id, err := uuid.NewRandom()
	if err != nil {
		log.Error().Caller().Err(err).Send()
		return cl, err
	}
	cl.Id = id.String()
	if err := c.BindJSON(&cl); err != nil {
		return cl, err
	}

	validate := validator.New()
	if err := validate.Struct(cl); err != nil {
		return cl, err
	}
	clerk := Clerk{Id: cl.Id, Name: cl.Name}
	if err := table.Put(clerk).Run(); err != nil {
		return cl, err
	}
	return cl, nil

}

// Get is get a clerk
func (s clerkService) Get(id string) (entity.Clerk, error) {
	log.Debug().Caller().Msg("clerks get")
	dynamodb := Dynamo_Init()
	table := dynamodb.Table("clerks")
	var cl entity.Clerk
	if err := table.Get("id", id).One(&cl); err != nil {
		return cl, err
	}
	return cl, nil
}

//  Delete is delete a clerk
func (s clerkService) Delete(id string) (entity.Clerk, error) {
	log.Debug().Caller().Msg("clerks delete")
	dynamodb := Dynamo_Init()
	table := dynamodb.Table("clerks")
	var cl entity.Clerk

	// 指定したIDが存在するか確認
	if err := table.Get("id", id).One(&cl); err != nil {
		return cl, err
	}

	// 削除
	if err := table.Delete("id", id).Run(); err != nil {
		panic(err.Error())
	}

	return cl, nil
}
