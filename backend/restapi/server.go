package restapi

import (
	"fmt"
	"log"
	"net/http"
)

type MyServer struct {
	port                  string
	dutiesHandler         handlers.dutiesHandler
	dutyCategoriesHandler handlers.dutyCategoriesHandler
	dutyDayHandler        handlers.dutyDayHandler
	dutyWorkersHandler    handlers.dutyWorkersHandler
	userHandler           handlers.userHandler
	//authMiddleware middleware.AuthMiddleware
}

//func NewServer(port string, usersHandler handlers.UsersHandler, schoolHandlers handlers.SchoolHandler, groupHandlers handlers.GroupsHandler, authMiddleware middleware.AuthMiddleware) *MyServer {
func NewServer(port string,
	dutiesHandler handlers.dutiesHandler,
	dutyCategoriesHandler handlers.dutyCategoriesHandler,
	dutyDayHandler handlers.dutyDayHandler,
	dutyWorkersHandler handlers.dutyWorkersHandler,
	userHandler handlers.userHandler,
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

	userMux := http.NewServeMux()
	userMux.HandleFunc("/users", server.userHandler.Users)
	userMux.HandleFunc("/user/", server.userHandler.User)

	userMux.HandleFunc("/", indexHandler) //)server.companyHandler.Companies

	userMux.HandleFunc("/dutyWorkers", server.dutyWorkersHandler.DutyWorkers)
	userMux.HandleFunc("/dutyWorker/", server.dutyWorkersHandler.DutyWorker)

	userMux.HandleFunc("/dutyDays", server.dutyDayHandler.DutyDays)
	userMux.HandleFunc("/dutyDay/", server.dutyDayHandler.DutyDay)

	userMux.HandleFunc("/dutyCategories", server.dutyCategoriesHandler.DutyCategories)
	userMux.HandleFunc("/dutyCategory/", server.dutyCategoriesHandler.DutyCategory)

	userMux.HandleFunc("/duties", server.dutiesHandler.Duties)
	userMux.HandleFunc("/duty/", server.dutiesHandler.Duty)

	fmt.Printf("listening at %s", server.port)
	log.Fatal(http.ListenAndServe(server.port, userMux))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
