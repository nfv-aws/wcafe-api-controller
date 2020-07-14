package service

import (
	"context"
	"time"

	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"gopkg.in/go-playground/validator.v9"

	"github.com/nfv-aws/wcafe-api-controller/config"
	"github.com/nfv-aws/wcafe-api-controller/entity"
	pb "github.com/nfv-aws/wcafe-api-controller/protoc"
)

// Supply is alias of entity.Supply struct
type Supply entity.Supply

// Supply is alias of entity.Supplies struct
type Supplies entity.Supplies

// Service procides Supply's behavior
type SupplyService interface {
	List() ([]entity.Supply, error)
	Create(c *gin.Context) (entity.Supply, error)
	Delete(id string) (entity.Supply, error)
}

func NewSupplyService() SupplyService {
	return &supplyService{}
}

type supplyService struct{}

// List is get all supply
func (s supplyService) List() ([]entity.Supply, error) {
	log.Debug().Caller().Msg("supplies list")
	// Set up a connection to the server.
	config.Configure()
	var address = config.C.Conductor.Ip + ":" + config.C.Conductor.Port
	log.Info().Caller().Msg(address)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Error().Caller().Err(err).Send()
	}
	defer conn.Close()
	c := pb.NewSuppliesClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SupplyList(ctx, &pb.SupplyListRequest{Table: "supplies"})
	if err != nil {
		log.Error().Caller().Err(err).Send()
	}
	log.Info().Caller().Msg(r.GetMessage())
	var supplies []entity.Supply
	err = json.Unmarshal([]byte(r.GetMessage()), &supplies)
	if err != nil {
		log.Error().Caller().Err(err).Send()
	}
	return supplies, nil
}

// Create new supply
func (s supplyService) Create(c *gin.Context) (entity.Supply, error) {
	log.Debug().Caller().Msg("supply create")
	// Set up a connection to the server.
	config.Configure()
	var address = config.C.Conductor.Ip + ":" + config.C.Conductor.Port
	log.Info().Caller().Msg(address)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Error().Caller().Err(err).Send()
	}
	defer conn.Close()
	connect := pb.NewSuppliesClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var supply entity.Supply
	//UUID生成
	id, err := uuid.NewRandom()
	if err != nil {
		log.Error().Caller().Err(err).Send()
		return supply, err
	}
	if err := c.BindJSON(&supply); err != nil {
		return supply, err
	}

	validate := validator.New()
	if err := validate.Struct(supply); err != nil {
		return supply, err
	}

	supply.Id = id.String()
	req, err := json.Marshal(supply)
	if err != nil {
		log.Error().Caller().Err(err).Send()
	}

	r, err := connect.SupplyCreate(ctx, &pb.SupplyCreateRequest{Table: "supplies", Body: string(req)})
	if err != nil {
		log.Error().Caller().Err(err).Send()
	}
	log.Info().Caller().Msg(r.GetMessage())
	err = json.Unmarshal([]byte(r.GetMessage()), &supply)
	if err != nil {
		log.Error().Caller().Err(err).Send()
	}
	return supply, nil
}

// Delete is delete a supply
func (s supplyService) Delete(id string) (entity.Supply, error) {
	log.Debug().Caller().Msg("supply delete")
	// Set up a connection to the server.
	config.Configure()
	var address = config.C.Conductor.Ip + ":" + config.C.Conductor.Port
	log.Info().Caller().Msg(address)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Error().Caller().Err(err).Send()
	}
	defer conn.Close()
	connect := pb.NewSuppliesClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var supply entity.Supply
	r, err := connect.SupplyDelete(ctx, &pb.SupplyDeleteRequest{Table: "supplies", Id: id})
	if err != nil {
		return supply, err
	}
	log.Info().Caller().Msg(r.GetMessage())
	err = json.Unmarshal([]byte(r.GetMessage()), &supply)
	if err != nil {
		log.Error().Caller().Err(err).Send()
		return supply, err
	}
	return supply, nil
}
