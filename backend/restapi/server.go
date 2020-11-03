package restapi

import (
	"dsv/restapi/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type MyServer struct {
	port                  string
	dutiesHandler         handlers.DutiesHandler
	dutyCategoriesHandler handlers.DutyCategoriesHandler
	dutyDayHandler        handlers.DutyDayHandler
	dutyWorkersHandler    handlers.DutyWorkersHandler
	userHandler           handlers.UserHandler
	//authMiddleware middleware.AuthMiddleware
}

//func NewServer(port string, usersHandler handlers.UsersHandler, schoolHandlers handlers.SchoolHandler, groupHandlers handlers.GroupsHandler, authMiddleware middleware.AuthMiddleware) *MyServer {
func NewServer(port string,
	dutiesHandler handlers.DutiesHandler,
	dutyCategoriesHandler handlers.DutyCategoriesHandler,
	dutyDayHandler handlers.DutyDayHandler,
	dutyWorkersHandler handlers.DutyWorkersHandler,
	userHandler handlers.UserHandler,
) *MyServer {
	return &MyServer{
		port:                  port,
		dutiesHandler:         dutiesHandler,
		dutyCategoriesHandler: dutyCategoriesHandler,
		dutyDayHandler:        dutyDayHandler,
		dutyWorkersHandler:    dutyWorkersHandler,
		userHandler:           userHandler,
	}
}

func (server *MyServer) ConfigureAndRun() {

	siteMux := mux.NewRouter()

	// Users
	siteMux.HandleFunc("/login", server.userHandler.Login)
	siteMux.HandleFunc("/logout", server.userHandler.Logout)
	siteMux.HandleFunc("/register", server.userHandler.AddUser)
	siteMux.HandleFunc("/getusers", server.userHandler.GetUsersList)

	//GET byID and List
	siteMux.HandleFunc("/DutyCategories/", server.dutyCategoriesHandler.DutyCategories)
	//Post,Patch,Delete
	siteMux.HandleFunc("/DutyCategory/", server.dutyCategoriesHandler.DutyCategory)
	//
	siteMux.HandleFunc("/addworker", server.dutyWorkersHandler.AddWorker)
	siteMux.HandleFunc("/deleteworker", server.dutyWorkersHandler.DeleteWorker)
	siteMux.HandleFunc("/updateworker", server.dutyWorkersHandler.UpdateWorker)
	siteMux.HandleFunc("/getworker", server.dutyWorkersHandler.GetWorkerById)
	siteMux.HandleFunc("/getworkerslist", server.dutyWorkersHandler.GetWorkersList)
	//
	siteMux.HandleFunc("/addday", server.dutyDayHandler.DutyDay)
	siteMux.HandleFunc("/deleteday", server.dutyDayHandler.DutyDay)
	siteMux.HandleFunc("/updateday", server.dutyDayHandler.DutyDay)
	siteMux.HandleFunc("/getday", server.dutyDayHandler.DutyDay)
	siteMux.HandleFunc("/getdaylist", server.dutyDayHandler.DutyDays)
	//
	siteMux.HandleFunc("/Duty/", server.dutiesHandler.Duty)
	siteMux.HandleFunc("/Duties", server.dutiesHandler.Duties)

	fmt.Printf("listening at %s", server.port)
	log.Fatal(http.ListenAndServe(server.port, siteMux))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
