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

// Get List of duty days
func (dd *dutyDayHandler) DutyDays(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Custom-Header, Quantity")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")

	switch r.Method {
	case http.MethodGet:
		dutyDays, err := dd.dutyDayRepository.GetDutyDays(context.Background())
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		dutyDaysJsonList, err := json.Marshal(dutyDays)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(dutyDaysJsonList))
		return
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (dd *dutyDayHandler) DutyDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Custom-Header, Quantity")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")

	switch r.Method {
	case http.MethodPost:
		reqBody := json.NewDecoder(r.Body)
		var newDayData domain.DutyDay
		err := reqBody.Decode(&newDayData)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		newCategory, err := dd.dutyDayRepository.CreateDutyDay(context.Background(), &newDayData)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			return
		}

		res, err := json.Marshal(newCategory)
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
		var element domain.IdDutyDayData
		err := requestBody.Decode(&element)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = dd.dutyDayRepository.DeleteDutyDay(context.Background(), element.ID)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	case http.MethodPatch:
		requestBody := json.NewDecoder(r.Body)
		var element domain.PatchDutyDayData
		err := requestBody.Decode(&element)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = dd.dutyDayRepository.UpdateDutyDay(context.Background(), element.ID, element.Day, element.IsDayOff)
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
			dutyDay, err := dd.dutyDayRepository.GetDutyDayByID(context.Background(), idElement)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			dutyDayJson, err := json.Marshal(dutyDay)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(dutyDayJson))
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
