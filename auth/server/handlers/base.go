package handlers

import (
	"encoding/json"
	"net/http"
)

type successResponse struct {
	Status  int         `json:"status"`
	Payload interface{} `json:"payload"`
}
type errorResponse struct {
	Status int         `json:"status"`
	Err    interface{} `json:"error"`
}

func writeJsonResponse(rw http.ResponseWriter, payload interface{}) {
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(payload)
}

func successCreatedResponse(rw http.ResponseWriter, payload interface{}) {
	rw.WriteHeader(http.StatusCreated)
	p := successResponse{
		Status:  http.StatusCreated,
		Payload: payload,
	}
	writeJsonResponse(rw, p)
}
func successSingleResponse(rw http.ResponseWriter, payload interface{}) {
	rw.WriteHeader(http.StatusOK)
	p := successResponse{
		Status:  http.StatusOK,
		Payload: payload,
	}
	writeJsonResponse(rw, p)
}

func badRequestResponse(rw http.ResponseWriter, payload interface{}) {
	rw.WriteHeader(http.StatusBadRequest)
	p := errorResponse{
		Status: http.StatusBadRequest,
		Err:    payload,
	}
	writeJsonResponse(rw, p)
}
func notFoundResponse(rw http.ResponseWriter, payload interface{}) {
	rw.WriteHeader(http.StatusNotFound)
	p := errorResponse{
		Status: http.StatusNotFound,
		Err:    payload,
	}
	writeJsonResponse(rw, p)
}
