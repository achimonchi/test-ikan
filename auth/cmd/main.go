package main

import (
	"auth/config"
	"auth/pkg/database"
	"auth/pkg/utils"
	"auth/repositories"
	"auth/server"
	"auth/server/handlers"
	"auth/server/middleware"
	"auth/services"
	"fmt"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	config := config.GenerateConfig()

	db := database.NewPostgres(config)
	if db != nil {
		fmt.Println("db success")
	}

	token := utils.NewToken(config)

	trace := middleware.NewTraceMiddleware()

	authRepo := repositories.NewAuthRepo(db.DB)

	authServices := services.NewAuthServices(authRepo, token)

	pingHandler := handlers.NewPingHandler()
	authHandler := handlers.NewAuthHandler(authServices)

	handler := handlers.NewHandlers(pingHandler, authHandler)

	server.NewServer(
		config.APP_PORT,
		router,
		trace,
		handler,
	).StartServer()
}
