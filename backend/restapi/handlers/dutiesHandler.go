package handlers

import (
	"context"
	"dsv/domain"
	"dsv/infrastructure"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// ProductHandler constructor
type DutiesHandler interface {
	Duties(w http.ResponseWriter, r *http.Request)
	Duty(w http.ResponseWriter, r *http.Request)
}

type dutiesHandler struct {
	dutiesRepository infrastructure.DutiesRepository
}

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

	switch r.Method {
	case http.MethodGet:
		duties, err := dh.dutiesRepository.GetDuties(context.Background())
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		dutiesJsonList, err := json.Marshal(duties)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(dutiesJsonList))
		return
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (dh *dutiesHandler) Duty(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Custom-Header, Quantity")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")

	switch r.Method {
	case http.MethodPost:
		reqBody := json.NewDecoder(r.Body)
		var element domain.Duties
		err := reqBody.Decode(&element)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		newDuty, err := dh.dutiesRepository.CreateDuty(context.Background(), &element)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			return
		}

		res, err := json.Marshal(newDuty)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(res))
		return
	case http.MethodDelete:
		requestBody := json.NewDecoder(r.Body)
		var element domain.IdDuties
		err := requestBody.Decode(&element)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = dh.dutiesRepository.DeleteDuty(context.Background(), element.ID)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	case http.MethodPatch:
		requestBody := json.NewDecoder(r.Body)
		var element domain.PatchDutiesData
		err := requestBody.Decode(&element)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = dh.dutiesRepository.UpdateDuty(context.Background(), element.ID, element.IDDutyDay, element.IDDutyWorker)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		stringPath := r.URL.String()
		addressParts := strings.Split(stringPath, "/")

		if len(addressParts) > 2 && addressParts[2] != "" {
			idElement := addressParts[2]
			duties, err := dh.dutiesRepository.GetDutyByID(context.Background(), idElement)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			dutyJson, err := json.Marshal(duties)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(dutyJson))
			return
		}
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
