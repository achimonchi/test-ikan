package server

import (
	"auth/server/handlers"
	"auth/server/middleware"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type server struct {
	port     string
	router   *httprouter.Router
	trace    *middleware.Trace
	handlers *handlers.Handlers
}

func NewServer(port string, router *httprouter.Router, trace *middleware.Trace, handlers *handlers.Handlers) *server {
	return &server{
		port:     port,
		router:   router,
		trace:    trace,
		handlers: handlers,
	}
}

func (s *server) StartServer() {
	s.router.GET("/v1/ping", s.trace.Trace(s.handlers.Ping.Ping))

	fmt.Println("server running at port", s.port)
	http.ListenAndServe(s.port, s.router)
}
