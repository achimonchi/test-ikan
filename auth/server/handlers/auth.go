package handlers

import (
	"auth/server/params"
	"auth/services"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthHandler interface {
	Registration(rw http.ResponseWriter, r *http.Request, p httprouter.Params)
}

type authHandler struct {
	service *services.AuthServices
}

func NewAuthHandler(service *services.AuthServices) AuthHandler {
	return &authHandler{
		service: service,
	}
}

func (a *authHandler) Registration(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var req params.CreateAuth

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		badRequestResponse(rw, err.Error())
		return
	}

	auth, err := a.service.CreateAuth(r.Context(), &req)
	if err != nil {
		badRequestResponse(rw, err.Error())
		return
	}

	successCreatedResponse(rw, auth)
}
