package controller

import (
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
		c.AbortWithStatus(404)
		log.Error().Caller().Err(err)
	} else {
		c.JSON(200, u)
	}
}

// Create action: POST /clerks
func (cc ClerkController) Create(c *gin.Context) {
	log.Debug().Caller().Msg("clerks create")
	u, err := cc.Clerkservice.Create(c)

	if err != nil {
		c.AbortWithStatus(400)
		log.Error().Caller().Err(err)
	} else {
		c.JSON(201, u)
	}
}
