package controller

import (
	"net/http"

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
		c.AbortWithStatus(http.StatusNotFound)
		log.Error().Caller().Err(err).Send()
	} else {
		c.JSON(http.StatusOK, s)
	}
}

// Create action: POST /supplies
func (sc SupplyController) Create(c *gin.Context) {
	log.Debug().Caller().Msg("supplies create")
	s, err := sc.Supplyservice.Create(c)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Error().Caller().Err(err).Send()
	} else {
		c.JSON(http.StatusCreated, s)
	}
}
