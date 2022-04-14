package main

import (
	"fetch/config"
	"fetch/server"
	"fmt"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	config := config.GenerateConfig()

	fmt.Println(config.APP_PORT)

	server.NewServer(config.APP_PORT, router).StartServer()
}
