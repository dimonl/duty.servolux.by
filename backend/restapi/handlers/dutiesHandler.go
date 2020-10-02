package handlers

import (
	"dsv/infrastructure"
	"net/http"
)

// ProductHandler constructor
type DutiesHandler interface {
	Duties(w http.ResponseWriter, r *http.Request)
	Duty(w http.ResponseWriter, r *http.Request)
}

type dutiesHandler struct {
	dutiesRepository infrastructure.DutiesRepository
}

// NewproductHandler constructor to check my github
// NewproductHandler constructor to check my github
// NewproductHandler constructor to check my github
// NewproductHandler constructor
func NewDutiesHandler(dutiesRepository infrastructure.DutiesRepository) DutiesHandler {
	return &dutiesHandler{
		dutiesRepository: dutiesRepository,
	}
}

func (dh *dutiesHandler) Duties(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Custom-Header, Quantity")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func (dh *dutiesHandler) Duty(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Custom-Header, Quantity")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
