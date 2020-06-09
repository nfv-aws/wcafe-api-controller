package server

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/nfv-aws/wcafe-api-controller/controller"
	"github.com/nfv-aws/wcafe-api-controller/db"
	"github.com/nfv-aws/wcafe-api-controller/entity"
	"github.com/nfv-aws/wcafe-api-controller/service"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	f, _ := os.Create("./log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	log.SetOutput(gin.DefaultWriter)
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())

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
		p.GET("/stores/:id/pets", store_ctrl.PetsList)

		pet_repo := entity.PetRepository{DB: db.GetDB()}
		pet_ctrl := controller.PetController{Petservice: service.NewPetService(pet_repo)}
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
		p.DELETE("/users/:id", user_ctrl.Delete)

		sup_ctrl := controller.SupplyController{Supplyservice: service.NewSupplyService()}
		p.GET("/supplies", sup_ctrl.List)

		clerk_ctrl := controller.ClerkController{Clerkservice: service.NewClerkService()}
		p.GET("/clerks", clerk_ctrl.List)

	}

	return r
}
