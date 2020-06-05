package main

import (
	"github.com/nfv-aws/wcafe-api-controller/db"
	logger "github.com/nfv-aws/wcafe-api-controller/log"
	"github.com/nfv-aws/wcafe-api-controller/server"
)

func main() {
	logger.LoggingSettings()
	db.Init()
	server.Init()
	db.Close()
}
