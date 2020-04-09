package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nfv-aws/wcafe-api-controller/service"
)

// Controller is user controlller
type UserController struct {
	Service service.UserService
}

// List action: GET /users
func (uc UserController) List(c *gin.Context) {
	u, err := uc.Service.List()

	if err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(200, u)
	}
}

// Create action: POST /users
func (uc UserController) Create(c *gin.Context) {
	u, err := uc.Service.Create(c)

	if err != nil {
		c.AbortWithStatus(400)
		log.Println(err)
	} else {
		c.JSON(201, u)
	}
}

// Get action: GET /users/:id
func (uc UserController) Get(c *gin.Context) {
	id := c.Params.ByName("id")

	u, err := uc.Service.Get(id)

	if err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(200, u)
	}
}
