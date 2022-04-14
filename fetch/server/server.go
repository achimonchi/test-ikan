package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type server struct {
	port   string
	router *httprouter.Router
}

func NewServer(port string, router *httprouter.Router) *server {
	return &server{
		port:   port,
		router: router,
	}
}

func (s *server) StartServer() {

	fmt.Println("server running at port", s.port)
	http.ListenAndServe(s.port, s.router)
}
