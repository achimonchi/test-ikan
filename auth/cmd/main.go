package main

import (
	"auth/config"
	"auth/pkg/database"
	"auth/server"
	"auth/server/handlers"
	"auth/server/middleware"
	"fmt"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	config := config.GenerateConfig()

	trace := middleware.NewTraceMiddleware()

	pingHandler := handlers.NewPingHandler()

	handler := handlers.NewHandlers(pingHandler)

	db := database.NewPostgres(config)
	if db != nil {
		fmt.Println("db success")
	}

	server.NewServer(
		config.APP_PORT,
		router,
		trace,
		handler,
	).StartServer()
}
