package pet

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nfv-aws/wcafe-api-controller/service"
)

// Controller is pet controlller
type Controller struct{}

// Index action: GET /pets
func (pc Controller) Index(c *gin.Context) {
	var s pet.Service
	p, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /pets
func (pc Controller) Create(c *gin.Context) {
	//UUID生成
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return
	}

	//POST処理
	var s pet.Service
	p, err := s.CreateModel(c, u.String())

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

// Show action: GET /pets/:id
func (pc Controller) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s pet.Service
	p, err := s.GetByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}
