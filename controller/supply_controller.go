package controller

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/nfv-aws/wcafe-api-controller/service"
)

// Controller is user controller
type SupplyController struct {
	Supplyservice service.SupplyService
}

// List action: GET /supplies
func (sc SupplyController) List(c *gin.Context) {
	s, err := sc.Supplyservice.List()

	if err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(200, s)
	}
}
