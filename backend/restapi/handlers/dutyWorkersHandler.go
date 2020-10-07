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
type DutyWorkersHandler interface {
	AddWorker(w http.ResponseWriter, r *http.Request)
	UpdateWorker(w http.ResponseWriter, r *http.Request)
	DeleteWorker(w http.ResponseWriter, r *http.Request)
	GetWorkerById(w http.ResponseWriter, r *http.Request)
	GetWorkersList(w http.ResponseWriter, r *http.Request)
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

func (dw *dutyWorkersHandler) AddWorker(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var reg domain.DutyWorkersData
	err := decoder.Decode(&reg)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	us, err := dw.dutyWorkersRepository.CreateDutyWorker(context.Background(), domain.ConvertDutyWorkers(reg))
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	if us != nil {
		mUser, err := json.Marshal(&us)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintln(w, string(mUser))
		return
	}
}

func (dw *dutyWorkersHandler) UpdateWorker(w http.ResponseWriter, r *http.Request) {

	requestBody := json.NewDecoder(r.Body)
	var element domain.PatchDutyWorkersData
	err := requestBody.Decode(&element)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = dw.dutyWorkersRepository.UpdateDutyWorker(context.Background(), element.ID, element.FirstName, element.LastName, element.IDCategory)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (dw *dutyWorkersHandler) DeleteWorker(w http.ResponseWriter, r *http.Request) {

	requestBody := json.NewDecoder(r.Body)
	var element domain.IdDutyWorkersData
	err := requestBody.Decode(&element)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = dw.dutyWorkersRepository.DeleteDutyWorker(context.Background(), element.ID)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (dw *dutyWorkersHandler) GetWorkerById(w http.ResponseWriter, r *http.Request) {
	stringPath := r.URL.String()
	addressParts := strings.Split(stringPath, "/")
	idElement := addressParts[2]
	dutyWorker, err := dw.dutyWorkersRepository.GetDutyWorkerByID(context.Background(), idElement)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	dutyWorkerJson, err := json.Marshal(dutyWorker)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(dutyWorkerJson))
	return
}

func (dw *dutyWorkersHandler) GetWorkersList(w http.ResponseWriter, r *http.Request) {

	users, err := dw.dutyWorkersRepository.GetDutyWorkers(context.Background())
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	usersJson, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(usersJson))
	return
}