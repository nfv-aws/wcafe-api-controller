package server

import (
	"github.com/gin-contrib/cors"
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
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	// health check
	r.GET("/health", func(c *gin.Context) {
		c.String(200, "%s", "ok")
	})

	p := r.Group("/api/v1")
	{
		store_ctrl := controller.StoreController{Storeservice: service.NewStoreService()}
		p.GET("/stores", store_ctrl.List)
		p.GET("/stores/:id", store_ctrl.Get)
		p.POST("/stores", store_ctrl.Create)
		p.PATCH("/stores/:id", store_ctrl.Update)
		p.DELETE("stores/:id", store_ctrl.Delete)

		pet_ctrl := controller.PetController{Service: service.NewPetService()}
		p.GET("/pets", pet_ctrl.List)
		p.GET("/pets/:id", pet_ctrl.Get)
		p.POST("/pets", pet_ctrl.Create)
		p.PATCH("/pets/:id", pet_ctrl.Update)
		p.DELETE("/pets/:id", pet_ctrl.Delete)

		user_ctrl := controller.UserController{Userservice: service.NewUserService()}
		p.GET("/users", user_ctrl.List)
		p.GET("/users/:id", user_ctrl.Get)
		p.POST("/users", user_ctrl.Create)
		p.PATCH("/users/:id", user_ctrl.Update)

	}

	return r
}
