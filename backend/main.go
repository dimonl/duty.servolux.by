package main

import (
	"os"

	"dsv/infrastructure"
	"dsv/restapi"
	"dsv/restapi/handlers"
)

func main() {
	//dutiesRepo := infrastructure.NewDutiesRepository()
	dutyCategoriesRepo := infrastructure.NewDutyCategoriesRepository()
	//dutyDayRepo := infrastructure.NewDutyDayRepository()
	//dutyWorkersRepo := infrastructure.NewDutyWorkersRepository()
	userRepo := infrastructure.NewUserRepository()

	//dutiesHandler := handlers.NewDutiesHandler(dutiesRepo)
	dutyCategoriesHandler := handlers.NewDutyCategoriesHandler(dutyCategoriesRepo)
	//dutyDayHandler := handlers.NewDutyDayHandler(dutyDayRepo)
	//dutyWorkersHandler := handlers.NewDutyWorkersHandler(dutyWorkersRepo)
	userHandler := handlers.NewUserHandler(userRepo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
//	srv := restapi.NewServer(":"+port, dutiesHandler, dutyCategoriesHandler, dutyDayHandler, dutyWorkersHandler, userHandler) //, adminsHandler, authMiddleware)
	srv := restapi.NewServer(":"+port, dutyCategoriesHandler, userHandler) //, adminsHandler, authMiddleware)
	srv.ConfigureAndRun()
}
