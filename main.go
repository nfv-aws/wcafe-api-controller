package main

import (
	"github.com/nfv-aws/wcafe-api-controller/db"
	"github.com/nfv-aws/wcafe-api-controller/server"
)

func main() {
	db.Init()
	server.Init()
	db.Close()
}
