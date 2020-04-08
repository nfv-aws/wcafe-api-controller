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

<<<<<<< HEAD
// List action: GET /pets
func (pc PetController) List(c *gin.Context) {
	p, err := pc.Service.List()
=======
// Index action: GET /pets
func (pc PetController) Index(c *gin.Context) {
	//var s service.PetService
	//s := service.NewPetService()
	p, err := pc.Service.GetAll()
>>>>>>> Add Mock Test Sample

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
		c.AbortWithStatus(400)
		log.Println(err)
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
