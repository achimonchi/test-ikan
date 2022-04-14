package main

import (
	"fetch/config"
	"fetch/pkg/httpclient"
	"fetch/pkg/utils"
	"fetch/server"
	"fetch/server/handlers"
	"fetch/server/middleware"
	"fetch/services"
	"fmt"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	config := config.GenerateConfig()
	token := utils.NewToken(config)

	client := httpclient.NewHttpClient(
		config.CLIENT_HOST,
		config.CLIENT_PORT,
		config.CLIENT_TIMEOUT,
	)

	converter := httpclient.NewHttpClient(
		config.CONVERTER_HOST,
		config.CONVERTER_PORT,
		config.CONVERTER_TIMEOUT,
	)

	fetchServices := services.NewFetchServices(client, converter, config)

	fetchHandlers := handlers.NewFetchHandlers(fetchServices)

	pingHandlers := handlers.NewPingHandlers()

	handlers := handlers.NewHandlers(pingHandlers, fetchHandlers)

	traceMiddleware := middleware.NewTraceMiddleware()
	authMiddleware := middleware.NewAuthMiddleware(token)
	middleware := middleware.NewMiddleware(traceMiddleware, authMiddleware)

	fmt.Println(config.APP_PORT)

	server.NewServer(config.APP_PORT, router, handlers, middleware).StartServer()
}
