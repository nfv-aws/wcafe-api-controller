package main

import (
	"github.com/ysdy/go_rest/db"
	"github.com/ysdy/go_rest/server"
)

func main() {
	db.Init()
	server.Init()
	db.Close()
}
