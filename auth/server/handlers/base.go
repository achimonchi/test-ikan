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

func badRequestResponse(rw http.ResponseWriter, payload interface{}) {
	rw.WriteHeader(http.StatusBadRequest)
	p := errorResponse{
		Status: http.StatusBadRequest,
		Err:    payload,
	}
	writeJsonResponse(rw, p)
}
