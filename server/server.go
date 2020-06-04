package server

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/nfv-aws/wcafe-api-controller/config"
	"github.com/nfv-aws/wcafe-api-controller/controller"
	"github.com/nfv-aws/wcafe-api-controller/service"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()
}

func logger() *gin.Engine {
	config.Configure()
	file_path := config.C.LOG.File_path

	// ログファイルに出力する際の色設定
	gin.ForceConsoleColor()

	// Logging to a file.
	f, _ := os.Create(file_path + "gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
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
	router.Use(gin.Recovery())
	return router
}

func router() *gin.Engine {
	r := logger()
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
		p.GET("/stores/:id/pets", store_ctrl.PetsList)

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
		p.DELETE("/users/:id", user_ctrl.Delete)

	}

	return r
}
