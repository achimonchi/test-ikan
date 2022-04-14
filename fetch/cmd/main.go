package main

import (
	"fetch/config"
	"fetch/pkg/utils"
	"fetch/server"
	"fetch/server/handlers"
	"fetch/server/middleware"
	"fmt"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	config := config.GenerateConfig()
	token := utils.NewToken(config)

	pingHandlers := handlers.NewPingHandlers()

	handlers := handlers.NewHandlers(pingHandlers)

	traceMiddleware := middleware.NewTraceMiddleware()
	authMiddleware := middleware.NewAuthMiddleware(token)
	middleware := middleware.NewMiddleware(traceMiddleware, authMiddleware)

	fmt.Println(config.APP_PORT)

	server.NewServer(config.APP_PORT, router, handlers, middleware).StartServer()
}
