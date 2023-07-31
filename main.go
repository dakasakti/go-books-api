package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/fathxn/go-books-restapi/config"
	routes "github.com/fathxn/go-books-restapi/routers"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

	routes.Run()

	log.Println("Server is running on port", config.ENV.PORT)
}