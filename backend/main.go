package main

import (
	"main/infrastructure"
	"main/infrastructure/database"
	"main/restapi"
	"main/restapi/handlers"
)

func main(){
companyRepo := infrastructure.NewCompanyRepository(database.NewMongoClient)
specialityRepo := infrastructure.NewSpecialityRepository(database.NewMongoClient)
vacancyRepo := infrastructure.NewVacancyRepository(database.NewMongoClient)

	companyHandler := handlers.NewCompanyHandler(companyRepo)
	specialityHandler := handlers.NewSpecialityHandler(specialityRepo)
	vacancyHandler := handlers.NewVacancyHandler(vacancyRepo)
	srv := restapi.NewServer(":8080", companyHandler, vacancyHandler, specialityHandler) //, adminsHandler, authMiddleware)
	srv.ConfigureAndRun()

}
