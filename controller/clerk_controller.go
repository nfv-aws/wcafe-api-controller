package controller

import (
	"net/http"

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
