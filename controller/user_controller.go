package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/nfv-aws/wcafe-api-controller/service"
)

// Controller is user controller
type UserController struct {
	Userservice service.UserService
}

// List action: GET /users
func (uc UserController) List(c *gin.Context) {
	u, err := uc.Userservice.List()

	if err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(200, u)
	}
}

// Create action: POST /users
func (uc UserController) Create(c *gin.Context) {
	u, err := uc.Userservice.Create(c)

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

	u, err := uc.Userservice.Get(id)

	if err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(200, u)
	}
}

// Update action: PATCH /users/:id
func (uc UserController) Update(c *gin.Context) {

	id := c.Params.ByName("id")
	u, err := uc.Userservice.Update(id, c)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.AbortWithStatus(404)
			log.Println(err)
		} else {
			c.AbortWithStatus(400)
			log.Println(err)
		}
	} else {
		c.JSON(200, u)
	}
}

// Delete action: DELETE /pets/:id
func (uc UserController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	u, err := uc.Userservice.Delete(id)

	if err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(204, u)
	}
}
