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

	switch r.Method {
	case http.MethodPost:
		reqBody := json.NewDecoder(r.Body)
		var newCategoryData domain.DutyCategories
		err := reqBody.Decode(&newCategoryData)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		newCategory, err := dh.dutyCategoriesRepository.CreateDutyCategory(context.Background(), &newCategoryData)
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
		var element domain.IdDutyCategoriesData
		err := requestBody.Decode(&element)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = dh.dutyCategoriesRepository.DeleteDutyCategory(context.Background(), element.ID)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	case http.MethodPatch:
		requestBody := json.NewDecoder(r.Body)
		var element domain.PatchDutyCategoriesData
		err := requestBody.Decode(&element)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = dh.dutyCategoriesRepository.UpdateDutyCategory(context.Background(), element.ID, element.Category)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
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


//allow to get category by id and get category list
func (dh *dutyCategoriesHandler) DutyCategories(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Custom-Header, Quantity")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")


	switch r.Method {
	case http.MethodGet:
		stringPath := r.URL.String()
		addressParts := strings.Split(stringPath, "/")

		if len(addressParts) > 2 && addressParts[2] != "" {
			idElement := addressParts[2]
			dutyCategory, err := dh.dutyCategoriesRepository.GetDutyCategoryByID(context.Background(), idElement)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			dutyCategoryJson, err := json.Marshal(dutyCategory)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(dutyCategoryJson))
			return
		} else if len(addressParts) > 2 && addressParts[2] == "" {
			dutyCategoriesList, err := dh.dutyCategoriesRepository.GetDutyCategories(context.Background())
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			dutyCategoriesListJson, err := json.Marshal(dutyCategoriesList)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(dutyCategoriesListJson))
			return
		} else {
			w.WriteHeader(http.StatusBadRequest)
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
