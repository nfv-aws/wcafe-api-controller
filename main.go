package main

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"

	"github.com/nfv-aws/wcafe-api-controller/config"
	"github.com/nfv-aws/wcafe-api-controller/db"
	"github.com/nfv-aws/wcafe-api-controller/server"
)

func main() {
	//ログレベル設定
	switch os.Getenv("LOG_LVE") {
	case "Debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "Info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "Error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	// ログ出力を指定
	config.Configure()
	log_path := config.C.LOG.File_path
	writer := (&lumberjack.Logger{
		Filename:   log_path + "gin.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})
	writers := io.MultiWriter(os.Stdout, writer)

	// 時間をJSTに設定
	jst, _ := time.LoadLocation("Asia/Tokyo")
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().In(jst)
	}

	log.Logger = zerolog.New(writers).With().Timestamp().Logger()

	log.Debug().Caller().Msg("db.init")
	db.Init()

	log.Debug().Caller().Msg("server.Init")
	server.Init()

	db.Close()
}
