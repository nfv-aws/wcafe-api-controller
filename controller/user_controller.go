package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"

	"github.com/nfv-aws/wcafe-api-controller/service"
)

// Controller is user controller
type UserController struct {
	Userservice service.UserService
}

// List action: GET /users
func (uc UserController) List(c *gin.Context) {
	log.Debug().Caller().Msg("users list")

	u, err := uc.Userservice.List()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Error().Caller().Err(err)
	} else {
		c.JSON(http.StatusOK, u)
	}
}

// Create action: POST /users
func (uc UserController) Create(c *gin.Context) {
	log.Debug().Caller().Msg("users create")
	u, err := uc.Userservice.Create(c)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Error().Caller().Err(err)
	} else {
		c.JSON(http.StatusCreated, u)
	}
}

// Get action: GET /users/:id
func (uc UserController) Get(c *gin.Context) {
	log.Debug().Caller().Msg("users get")
	id := c.Params.ByName("id")

	u, err := uc.Userservice.Get(id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Error().Caller().Err(err)
	} else {
		c.JSON(http.StatusOK, u)
	}
}

// Update action: PATCH /users/:id
func (uc UserController) Update(c *gin.Context) {
	log.Debug().Caller().Msg("users uodate")

	id := c.Params.ByName("id")
	u, err := uc.Userservice.Update(id, c)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.AbortWithStatus(http.StatusNotFound)
			log.Error().Caller().Err(err)
		} else {
			c.AbortWithStatus(http.StatusBadRequest)
			log.Error().Caller().Err(err)
		}
	} else {
		c.JSON(http.StatusOK, u)
	}
}

// Delete action: DELETE /pets/:id
func (uc UserController) Delete(c *gin.Context) {
	log.Debug().Caller().Msg("users delete")
	id := c.Params.ByName("id")
	u, err := uc.Userservice.Delete(id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Error().Caller().Err(err)
	} else {
		c.JSON(http.StatusNoContent, u)
	}
}
