package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nfv-aws/wcafe-api-controller/service"
	"log"
	"strings"
)

// Controller is pet controlller
type PetController struct {
	Service service.PetService
}

// List action: GET /pets
func (pc PetController) List(c *gin.Context) {
	p, err := pc.Service.List()

	if err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /pets
func (pc PetController) Create(c *gin.Context) {
	p, err := pc.Service.Create(c)

	if err != nil {
		if strings.Contains(err.Error(), "InvalidAddress") {
			c.AbortWithStatus(404)
			log.Println(err)
		} else {
			c.AbortWithStatus(400)
			log.Println(err)
		}
	} else {
		c.JSON(201, p)
	}
}

// Get action: GET /pets/:id
func (pc PetController) Get(c *gin.Context) {
	id := c.Params.ByName("id")

	p, err := pc.Service.Get(id)

	if err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Update aciton: PATCH /pets/:id
func (pc PetController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	p, err := pc.Service.Update(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		log.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /pets/:id
func (pc PetController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	p, err := pc.Service.Delete(id)

	if err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(204, p)
	}
}
