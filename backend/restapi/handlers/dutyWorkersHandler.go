package handlers

import (
	"dsv/infrastructure"
	"net/http"
)

// ProductHandler constructor
type DutyWorkersHandler interface {
	DutyWorkers(w http.ResponseWriter, r *http.Request)
	DutyWorker(w http.ResponseWriter, r *http.Request)
}

type dutyWorkersHandler struct {
	dutyWorkersRepository infrastructure.DutyWorkersRepository
}

// NewproductHandler constructor
func NewDutyWorkersHandler(dutyWorkersRepository infrastructure.DutyWorkersRepository) DutyWorkersHandler {
	return &dutyWorkersHandler{
		dutyWorkersRepository: dutyWorkersRepository,
	}
}

func (dh *dutyWorkersHandler) DutyWorkers(w http.ResponseWriter, r *http.Request) {

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

func (dh *dutyWorkersHandler) DutyWorker(w http.ResponseWriter, r *http.Request) {

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
