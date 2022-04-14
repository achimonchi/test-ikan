package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PingHandlers interface {
	Ping(rw http.ResponseWriter, r *http.Request, p httprouter.Params)
}

type ping struct{}

func NewPingHandlers() PingHandlers {
	return &ping{}
}

func (*ping) Ping(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	json.NewEncoder(rw).Encode(map[string]string{
		"status": "ok",
	})
}
