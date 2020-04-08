package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/nfv-aws/wcafe-api-controller/service"
)

// Controller is store controlller
type StoreController struct{}

// Index action: GET /stores
func (pc StoreController) Index(c *gin.Context) {
	var s service.StoreService
	p, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /stores
func (pc StoreController) Create(c *gin.Context) {
	var s service.StoreService
	p, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

// Show action: GET /stores/:id
func (pc StoreController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.StoreService
	p, err := s.GetByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}
