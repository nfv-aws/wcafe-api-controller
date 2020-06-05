package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nfv-aws/wcafe-api-controller/config"
)

func GinLog() *gin.Engine {
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

func LoggingSettings() {
	config.Configure()
	file_path := config.C.LOG.File_path
	logfile, _ := os.OpenFile(file_path+"main.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	gin.DefaultWriter = io.MultiWriter(logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(gin.DefaultWriter)
}

/*
func Log() {
	logfile, err := os.OpenFile("./log/gin.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(os.Stdout, logfile)
	//	r := gin.Default()
	log.SetOutput(gin.DefaultWriter)
}
*/
