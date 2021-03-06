package controller

import (
	"net/http"

	"errors"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/nfv-aws/wcafe-api-controller/service"
)

// Controller is clerk controller
type ClerkController struct {
	Clerkservice service.ClerkService
}

// List action: GET /clerks
func (cc ClerkController) List(c *gin.Context) {
	log.Debug().Caller().Msg("clerks list")
	u, err := cc.Clerkservice.List()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Error().Caller().Err(err).Send()
	} else {
		c.JSON(http.StatusOK, u)
	}
}

// Create action: POST /clerks
func (cc ClerkController) Create(c *gin.Context) {
	log.Debug().Caller().Msg("clerks create")
	u, err := cc.Clerkservice.Create(c)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Error().Caller().Err(err).Send()
	} else {
		c.JSON(http.StatusCreated, u)
	}
}

// Get action: GET /clerks/:id
func (cc ClerkController) Get(c *gin.Context) {
	log.Debug().Caller().Msg("clerks get")
	id := c.Params.ByName("id")

	u, err := cc.Clerkservice.Get(id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Error().Caller().Err(err).Send()
	} else {
		c.JSON(http.StatusOK, u)
	}
}

// Update action: PATCH /clerks/:id
func (cc ClerkController) Update(c *gin.Context) {
	log.Debug().Caller().Msg("clerks update")
	id := c.Params.ByName("id")

	u, err := cc.Clerkservice.Update(id, c)
	var ErrNotFound = errors.New("dynamo: no item found")

	if err != nil {
		if reflect.DeepEqual(err, ErrNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			log.Error().Caller().Err(err).Send()
		} else {
			c.AbortWithStatus(http.StatusBadRequest)
			log.Error().Caller().Err(err).Send()
		}
	} else {
		c.JSON(http.StatusOK, u)
	}
}

// Delete action: DELETE /clerks/:id
func (cc ClerkController) Delete(c *gin.Context) {
	log.Debug().Caller().Msg("clerks delete")
	id := c.Params.ByName("id")
	u, err := cc.Clerkservice.Delete(id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Error().Caller().Err(err).Send()
	} else {
		c.JSON(http.StatusNoContent, u)
	}
}
