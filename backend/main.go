package main

import (
	"main/infrastructure"
	"main/restapi"
	"main/restapi/handlers"
)

func main() {
	dutiesRepo := infrastructure.NewDutiesRepository()
	dutyCategoriesRepo := infrastructure.NewDutyCategoriesRepository()
	dutyDayRepo := infrastructure.NewDutyDayRepository()
	dutyWorkersRepo := infrastructure.NewDutyWorkersRepository()
	userRepo := infrastructure.NewUserRepository()

	dutiesHandler := handlers.NewDutiesHandler(dutiesRepo)
	dutyCategoriesHandler := handlers.NewDutyCategoriesHandler(dutyCategoriesRepo)
	dutyDayHandler := handlers.NewDutyDayHandler(dutyDayRepo)
	dutyWorkersHandler := handlers.NewDutyWorkersHandler(dutyWorkersRepo)
	userHandler := handlers.NewUserHandler(userRepo)

	srv := restapi.NewServer(":8080", dutiesHandler, dutyCategoriesHandler, dutyDayHandler, dutyWorkersHandler, userHandler) //, adminsHandler, authMiddleware)
	srv.ConfigureAndRun()

}
