package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nfv-aws/wcafe-api-controller/controller"
	"github.com/nfv-aws/wcafe-api-controller/service"
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
		store_ctrl := controller.StoreController{Service: service.NewStoreService()}
		p.GET("/stores", store_ctrl.List)
		p.GET("/stores/:id", store_ctrl.Get)
		p.POST("/stores", store_ctrl.Create)

		pet_ctrl := controller.PetController{Service: service.NewPetService()}
		p.GET("/pets", pet_ctrl.List)
		p.GET("/pets/:id", pet_ctrl.Get)
		p.POST("/pets", pet_ctrl.Create)
		// p.PUT("/:id", ctrl.Update)
		// p.DELETE("/:id", ctrl.Delete)

		user_ctrl := controller.UserController{Service: service.NewUserService()}
		p.GET("/users", user_ctrl.List)
		p.GET("/users/:id", user_ctrl.Get)
		p.POST("/users", user_ctrl.Create)

	}

	return r
}
