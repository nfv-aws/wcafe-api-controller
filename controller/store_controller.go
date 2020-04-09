package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nfv-aws/wcafe-api-controller/service"
)

// Controller is store controlller
type StoreController struct {
	Service service.StoreService
}

// List action: GET /stores
func (sc StoreController) List(c *gin.Context) {
	p, err := sc.Service.List()

	if err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /stores
func (sc StoreController) Create(c *gin.Context) {
	p, err := sc.Service.Create(c)

	if err != nil {
		c.AbortWithStatus(400)
		log.Println(err)
	} else {
		c.JSON(201, p)
	}
}

// Get action: GET /stores/:id
func (sc StoreController) Get(c *gin.Context) {
	id := c.Params.ByName("id")
	p, err := sc.Service.Get(id)

	if err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(200, p)
	}
}
