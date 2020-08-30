package restapi

import (
	"main/restapi/handlers"
	"fmt"
	"log"
	"net/http"
)

type MyServer struct {
	port           string
	companyHandler	handlers.CompanyHandler
	vacancyHandler  handlers.VacancyHandler
	specialityHandler  handlers.SpecialityHandler
	//authMiddleware middleware.AuthMiddleware
}

//func NewServer(port string, usersHandler handlers.UsersHandler, schoolHandlers handlers.SchoolHandler, groupHandlers handlers.GroupsHandler, authMiddleware middleware.AuthMiddleware) *MyServer {
func NewServer(port string, companyHandler handlers.CompanyHandler, vacancyHandler handlers.VacancyHandler, specialityHandler handlers.SpecialityHandler) *MyServer {
	return &MyServer{
		port:           port,
		companyHandler: companyHandler,
		vacancyHandler: vacancyHandler,
		specialityHandler:  specialityHandler,
	}
}

func (server *MyServer) ConfigureAndRun() {

	userMux := http.NewServeMux()
	userMux.HandleFunc("/companies", server.companyHandler.Companies)
	userMux.HandleFunc("/companies/", server.companyHandler.Company)
	userMux.HandleFunc("/", indexHandler)//)server.companyHandler.Companies

	userMux.HandleFunc("/specialities", server.specialityHandler.Specialities)
	userMux.HandleFunc("/specialities/", server.specialityHandler.Speciality)
	userMux.HandleFunc("/vacancies", server.vacancyHandler.Vacancies)
	userMux.HandleFunc("/vacancies/", server.vacancyHandler.Vacancy)


	fmt.Printf("listening at %s", server.port)
	log.Fatal(http.ListenAndServe(server.port, userMux))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}