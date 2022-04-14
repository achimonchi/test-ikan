package server

import (
	"fetch/server/handlers"
	"fetch/server/middleware"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type server struct {
	port       string
	router     *httprouter.Router
	handlers   *handlers.Handlers
	middleware *middleware.Middleware
}

func NewServer(port string, router *httprouter.Router, handlers *handlers.Handlers, middleware *middleware.Middleware) *server {
	return &server{
		port:       port,
		router:     router,
		handlers:   handlers,
		middleware: middleware,
	}
}

func (s *server) StartServer() {

	s.router.GET("/v1/ping", s.middleware.Trace.Trace(s.handlers.PingHandlers.Ping))

	fmt.Println("server running at port", s.port)
	http.ListenAndServe(s.port, s.router)
}
