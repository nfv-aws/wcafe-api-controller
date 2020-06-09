package controller

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/nfv-aws/wcafe-api-controller/service"
)

// Controller is clerk controller
type ClerkController struct {
	Clerkservice service.ClerkService
}

// List action: GET /clerks
func (cc ClerkController) List(c *gin.Context) {
	u, err := cc.Clerkservice.List()

	if err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(200, u)
	}
}
