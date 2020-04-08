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
		store_ctrl := controller.StoreController{}
		p.GET("/stores", store_ctrl.Index)
		p.GET("/stores/:id", store_ctrl.Show)
		p.POST("/stores", store_ctrl.Create)

		pet_ctrl := controller.PetController{}
		p.GET("/pets", pet_ctrl.Index)
		p.GET("/pets/:id", pet_ctrl.Show)
		p.POST("/pets", pet_ctrl.Create)
		// p.PUT("/:id", ctrl.Update)
		// p.DELETE("/:id", ctrl.Delete)

		user_ctrl := controller.UserController{}
		p.GET("/users", user_ctrl.Index)
		p.GET("/users/:id", user_ctrl.Show)
		p.POST("/users", user_ctrl.Create)

	}

	return r
}
