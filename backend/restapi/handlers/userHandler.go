package handlers

import (
	"context"
	"dsv/domain"
	"dsv/infrastructure"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserHandler interface {
	AddUser(w http.ResponseWriter, r *http.Request)
	GetUsersList(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
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

func (uh *userHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var reg domain.RegistrationData
	decoder.Decode(&reg)

	us, err := uh.userRepository.CreateUser(context.Background(), reg.Login, reg.Password)
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

func (uh *userHandler) GetUsersList(w http.ResponseWriter, r *http.Request) {
	users, err := uh.userRepository.GetUsers(context.Background())
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

func (uh *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

	fmt.Println("request: ", r)
	fmt.Println("request body: ", r.Body)

	decoder := json.NewDecoder(r.Body)
	var reg domain.RegistrationData
	err := decoder.Decode(&reg)
	if err != nil {
		fmt.Println("cannot decode login data due to:", err)
		fmt.Println(r.Body)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("login data from frontend: ", reg)

	user, err := uh.userRepository.UserLogin(context.Background(), reg.Login, reg.Password)
	if err != nil {
		if err.Error() == "user not found" {
			fmt.Println("cannot get data from database due to:", err)
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, err.Error())
			return
		}
		if err.Error() == "empty login" {
			fmt.Println("cannot get data from database due to:", err)
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, err.Error())
			return
		}

		fmt.Println("cannot get data from database due to:", err)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, err.Error())
		return
	}

	if user != nil {
		fmt.Println("token generation")
		expiration := time.Now().Add(time.Hour)
		//userID, _ := strconv.ParseUint(user.ID, 10, 64)
		userID := user.ID.Hex()

		token, err := CreateToken(userID, expiration)
		if err != nil {
			fmt.Println("cannot create token due to: ", err)
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		fmt.Println("token created")

		cookie := http.Cookie{
			Name:    "jwt",
			Value:   token,
			Expires: expiration,
		}
		outData := domain.OutputLoginData{
			User:    user.FirstName,
			Token:   token,
			Expires: strconv.FormatInt(time.Hour.Milliseconds(), 10),
		}
		res, err := json.Marshal(outData)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &cookie)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(res))
	} else {
		fmt.Println("user is nil")
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func (uh *userHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	return
}

// CreateToken creates jwt token for user
func CreateToken(userid string, exp time.Time) (string, error) {
	var err error
	//Creating Access Token
	// os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = exp.Unix() //time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	//token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	token, err := at.SignedString([]byte("stoomgeesneecrraettekjewyt"))
	if err != nil {
		return "", err
	}
	return token, nil
}