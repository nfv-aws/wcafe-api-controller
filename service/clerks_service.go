package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"

	"github.com/nfv-aws/wcafe-api-controller/config"
	"github.com/nfv-aws/wcafe-api-controller/entity"
)

//DynamoDB用グローバル変数
var (
	dynamodb *dynamo.DB
)

// DynamoDBに接続
func Dynamo_Init() *dynamo.DB {
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
}

type clerkService struct{}

func NewClerkService() ClerkService {
	return &clerkService{}
}

// List is get all clerk
func (s clerkService) List() ([]entity.Clerk, error) {
	dynamodb := Dynamo_Init()
	table := dynamodb.Table("clerks_name")

	var c []entity.Clerk
	if err := table.Scan().All(&c); err != nil {
		return c, err
	}
	return c, nil
}
