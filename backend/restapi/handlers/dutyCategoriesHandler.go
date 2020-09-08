package handlers

import (
	"dsv/infrastructure"
	"net/http"
)

// ProductHandler constructor
type DutyCategoriesHandler interface {
	DutyCategory(w http.ResponseWriter, r *http.Request)
	DutyCategories(w http.ResponseWriter, r *http.Request)
}

type dutyCategoriesHandler struct {
	dutyCategoriesRepository infrastructure.DutyCategoriesRepository
}

// NewproductHandler constructor
func NewDutyCategoriesHandler(dutyCategoriesRepository infrastructure.DutyCategoriesRepository) DutyCategoriesHandler {
	return &dutyCategoriesHandler{
		dutyCategoriesRepository: dutyCategoriesRepository,
	}
}

func (dh *dutyCategoriesHandler) DutyCategory(w http.ResponseWriter, r *http.Request) {

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

func (dh *dutyCategoriesHandler) DutyCategories(w http.ResponseWriter, r *http.Request) {

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
