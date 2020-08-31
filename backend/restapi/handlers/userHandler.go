package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"main/domain"
	"main/infrastructure"
	"net/http"
	"strings"
)

type UserHandler interface {
	Users(w http.ResponseWriter, r *http.Request)
	User(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userRepository infrastructure.UserRepository
}

//Constructor
func NewUserHandler(userRepository infrastructure.UserRepository) UserHandler {
	return &userHandler{
		userRepository: userRepository,
	}
}

func (cm *userHandler) User(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	case http.MethodGet:
		stringPath := r.URL.String()
		addressParts := strings.Split(stringPath, "/")
		//fmt.Fprint(w, addressParts)
		if len(strings.Split(r.URL.String(), "/")) > 2 {
			idElement := addressParts[2]
			currElement, err := cm.userRepository.GetUserByID(context.Background(), idElement)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			currElementJson, err := json.Marshal(currElement)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(currElementJson))
			return
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case http.MethodDelete:
		requestBody := json.NewDecoder(r.Body)
		var element domain.IdUserData
		err := requestBody.Decode(&element)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = cm.userRepository.DeleteUser(context.Background(), element.Id)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	case http.MethodPatch:
		requestBody := json.NewDecoder(r.Body)
		var element domain.PatchUserData
		err := requestBody.Decode(&element)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = cm.userRepository.UpdateUser(context.Background(), element.Id, element.FirstName)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (cm *userHandler) Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users, err := cm.userRepository.GetCompanyList(context.Background())
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
	case http.MethodPost:
		requestBody := json.NewDecoder(r.Body)
		var userData domain.User
		err := requestBody.Decode(&userData)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		newUser, err := cm.userRepository.CreateUser(context.Background(), &userData)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		res, err := json.Marshal(newUser)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(res))
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
