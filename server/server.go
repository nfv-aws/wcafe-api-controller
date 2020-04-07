package server

import (
	"github.com/gin-gonic/gin"

	"github.com/nfv-aws/wcafe-api-controller/controller"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	p := r.Group("/api/v1")
	{
		ctrl := pet.Controller{}
		p.GET("/pets", ctrl.Index)
		p.GET("/pets/:id", ctrl.Show)
		p.POST("/pets", ctrl.Create)
		// p.PUT("/:id", ctrl.Update)
		// p.DELETE("/:id", ctrl.Delete)
	}

	return r
}
