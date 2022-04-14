package handlers

import (
	"auth/constants"
	"auth/server/params"
	"auth/services"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthHandler interface {
	Registration(rw http.ResponseWriter, r *http.Request, p httprouter.Params)
	Login(rw http.ResponseWriter, r *http.Request, p httprouter.Params)
	Profile(rw http.ResponseWriter, r *http.Request, p httprouter.Params)
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

func (a *authHandler) Login(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var req params.Login

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		badRequestResponse(rw, err.Error())
		return
	}

	token, err := a.service.LoginByPhone(&req)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			notFoundResponse(rw, err.Error())
		} else {
			badRequestResponse(rw, err.Error())
		}

		return
	}

	successSingleResponse(rw, token)
}

func (a *authHandler) Profile(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	token := r.Context().Value(constants.TOKEN)

	tokenStr := fmt.Sprintf("%v", token)

	profile, err := a.service.Profile(tokenStr)
	if err != nil {
		badRequestResponse(rw, err)
		return
	}

	successSingleResponse(rw, profile)
}
