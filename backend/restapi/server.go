package restapi

import (
	"dsv/restapi/handlers"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type MyServer struct {
	port                  string
	//dutiesHandler         handlers.DutiesHandler
	dutyCategoriesHandler handlers.DutyCategoriesHandler
	//dutyDayHandler        handlers.DutyDayHandler
	dutyWorkersHandler    handlers.DutyWorkersHandler
	userHandler           handlers.UserHandler
	//authMiddleware middleware.AuthMiddleware
}

//func NewServer(port string, usersHandler handlers.UsersHandler, schoolHandlers handlers.SchoolHandler, groupHandlers handlers.GroupsHandler, authMiddleware middleware.AuthMiddleware) *MyServer {
func NewServer(port string,
	//dutiesHandler handlers.DutiesHandler,
	dutyCategoriesHandler handlers.DutyCategoriesHandler,
	//dutyDayHandler handlers.DutyDayHandler,
	dutyWorkersHandler handlers.DutyWorkersHandler,
	userHandler handlers.UserHandler,
) *MyServer {
	return &MyServer{
		port:                  port,
		//dutiesHandler:         dutiesHandler,
		dutyCategoriesHandler: dutyCategoriesHandler,
		//dutyDayHandler:        dutyDayHandler,
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

	siteMux.HandleFunc("/addworker", server.dutyWorkersHandler.AddWorker)
	siteMux.HandleFunc("/deleteworker", server.dutyWorkersHandler.DeleteWorker)
	siteMux.HandleFunc("/updateworker", server.dutyWorkersHandler.UpdateWorker)
	siteMux.HandleFunc("/getworker", server.dutyWorkersHandler.GetWorkerById)
	siteMux.HandleFunc("/getworkerslist", server.dutyWorkersHandler.GetWorkersList)

	//userMux := http.NewServeMux()
	//userMux.HandleFunc("/users", server.userHandler.Users)
	//userMux.HandleFunc("/user/", server.userHandler.User)
	//
	//userMux.HandleFunc("/", indexHandler) //)server.companyHandler.Companies
	//
	//userMux.HandleFunc("/dutyWorkers", server.dutyWorkersHandler.DutyWorkers)
	//userMux.HandleFunc("/dutyWorker/", server.dutyWorkersHandler.DutyWorker)
	//
	//userMux.HandleFunc("/dutyDays", server.dutyDayHandler.DutyDays)
	//userMux.HandleFunc("/dutyDay/", server.dutyDayHandler.DutyDay)
	//
	//userMux.HandleFunc("/dutyCategories", server.dutyCategoriesHandler.DutyCategories)
	//userMux.HandleFunc("/dutyCategory/", server.dutyCategoriesHandler.DutyCategory)
	//
	//userMux.HandleFunc("/duties", server.dutiesHandler.Duties)
	//userMux.HandleFunc("/duty/", server.dutiesHandler.Duty)

	fmt.Printf("listening at %s", server.port)
	log.Fatal(http.ListenAndServe(server.port, siteMux))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
