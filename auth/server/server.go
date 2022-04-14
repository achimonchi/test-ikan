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
	auth     *middleware.Auth
	handlers *handlers.Handlers
}

func NewServer(port string, router *httprouter.Router, trace *middleware.Trace, auth *middleware.Auth, handlers *handlers.Handlers) *server {
	return &server{
		port:     port,
		router:   router,
		trace:    trace,
		auth:     auth,
		handlers: handlers,
	}
}

func (s *server) StartServer() {
	s.router.GET("/v1/ping", s.trace.Trace(s.handlers.Ping.Ping))

	s.router.POST("/v1/signup", s.trace.Trace(s.handlers.Auth.Registration))
	s.router.POST("/v1/signin", s.trace.Trace(s.handlers.Auth.Login))
	s.router.GET("/v1/profile", s.trace.Trace(s.auth.Auth(s.handlers.Auth.Profile)))

	fmt.Println("server running at port", s.port)
	http.ListenAndServe(s.port, s.router)
}
