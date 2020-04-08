package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nfv-aws/wcafe-api-controller/service"
)

// Controller is pet controlller
type PetController struct {
	Service service.PetService
}

// Index action: GET /pets
func (pc PetController) Index(c *gin.Context) {
	//var s service.PetService
	//s := service.NewPetService()
	p, err := pc.Service.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /pets
func (pc PetController) Create(c *gin.Context) {
	var s service.PetService
	p, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		log.Println(err)
	} else {
		c.JSON(201, p)
	}
}

// Show action: GET /pets/:id
func (pc PetController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.PetService
	p, err := s.GetByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(200, p)
	}
}
