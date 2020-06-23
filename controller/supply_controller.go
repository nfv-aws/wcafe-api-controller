package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/nfv-aws/wcafe-api-controller/service"
)

// Controller is user controller
type SupplyController struct {
	Supplyservice service.SupplyService
}

// List action: GET /supplies
func (sc SupplyController) List(c *gin.Context) {
	log.Debug().Caller().Msg("supplies list")
	s, err := sc.Supplyservice.List()

	if err != nil {
		c.AbortWithStatus(404)
		log.Error().Caller().Err(err)
	} else {
		c.JSON(200, s)
	}
}

//Create
func (sc SupplyController) Create(c *gin.Context) {
	log.Debug().Caller().Msg("supplies create")
	s, err := sc.Supplyservice.Create(c)

	if err != nil {
		c.AbortWithStatus(400)
		log.Error().Caller().Err(err)
	} else {
		c.JSON(201, s)
	}
}
