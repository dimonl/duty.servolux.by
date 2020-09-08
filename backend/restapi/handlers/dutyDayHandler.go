package handlers

import (
	"dsv/infrastructure"
	"net/http"
)

// ProductHandler constructor
type DutyDayHandler interface {
	DutyDays(w http.ResponseWriter, r *http.Request)
	DutyDay(w http.ResponseWriter, r *http.Request)
}

type dutyDayHandler struct {
	dutyDayRepository infrastructure.DutyDayRepository
}

// NewproductHandler constructor
func NewDutyDayHandler(dutyDayRepository infrastructure.DutyDayRepository) DutyDayHandler {
	return &dutyDayHandler{
		dutyDayRepository: dutyDayRepository,
	}
}

func (dh *dutyDayHandler) DutyDays(w http.ResponseWriter, r *http.Request) {

}

func (dh *dutyDayHandler) DutyDay(w http.ResponseWriter, r *http.Request) {

}

func (dh *dutyDayHandler) DutyCategory(w http.ResponseWriter, r *http.Request) {

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

func (dh *dutyDayHandler) DutyCategories(w http.ResponseWriter, r *http.Request) {

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
