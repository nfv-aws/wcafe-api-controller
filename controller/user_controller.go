package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nfv-aws/wcafe-api-controller/service"
	"log"
)

// Controller is user controlller
type UserController struct{}

// Index action: GET /users
func (uc UserController) Index(c *gin.Context) {
	var s service.UserService
	u, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(200, u)
	}
}

// Create action: POST /users
func (uc UserController) Create(c *gin.Context) {
	var s service.UserService
	u, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		log.Println(err)
	} else {
		c.JSON(201, u)
	}
}

// Show action: GET /users/:id
func (uc UserController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.UserService
	u, err := s.GetByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(200, u)
	}
}
