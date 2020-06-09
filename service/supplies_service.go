package service

import (
	"context"
	"log"
	"time"

	"encoding/json"
	"google.golang.org/grpc"

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
	//	Create(c *gin.Context) (entity.Supply, error)
}

func NewSupplyService() SupplyService {
	return &supplyService{}
}

type supplyService struct{}

const (
	address = "localhost:50051"
)

// List is get all supply
func (s supplyService) List() ([]entity.Supply, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSuppliesClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	table := "supplies"
	r, err := c.SupplyList(ctx, &pb.SupplyRequest{Table: table})
	if err != nil {
		log.Fatalf("could not: %v", err)
	}
	log.Printf("%s", r.GetMessage())
	var supplies []entity.Supply
	err = json.Unmarshal([]byte(r.GetMessage()), &supplies)
	if err != nil {
		log.Fatal(err)
	}
	return supplies, nil
}
