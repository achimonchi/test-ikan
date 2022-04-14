package handlers

import (
	"fetch/services"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type FetchHandlers interface {
	FindAll(rw http.ResponseWriter, r *http.Request, p httprouter.Params)
}

type fetch struct {
	services *services.FetchServices
}

func NewFetchHandlers(services *services.FetchServices) FetchHandlers {
	return &fetch{
		services: services,
	}
}

func (f *fetch) FindAll(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	list, err := f.services.GetList()
	if err != nil {
		badRequestResponse(rw, err)
		return
	}
	successSingleResponse(rw, list)
}
