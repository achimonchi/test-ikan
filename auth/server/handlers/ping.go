package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PingHandler interface {
	Ping(rw http.ResponseWriter, r *http.Request, _ httprouter.Params)
}

type pingHandler struct{}

func NewPingHandler() PingHandler {
	return &pingHandler{}
}

func (p *pingHandler) Ping(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	json.NewEncoder(rw).Encode(map[string]string{
		"status": "ok",
	})
}
