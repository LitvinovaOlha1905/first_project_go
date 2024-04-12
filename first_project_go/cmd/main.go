package main

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"project/cmd/handler"
	"project/cmd/repository"
	"project/cmd/router"
	"project/cmd/service"
)

func main() {
	//fmt.Println("hi")

	logegr, _ := zap.NewProduction()
	defer logegr.Sync()

	webApp := fiber.New()
	repo := repository.NewMemoryEmployeeRepository()
	employeeService := service.NewEmployeeService(repo)
	employeeHandler := handler.NewEmployeeHandler(employeeService)
	router.SetupRoutes(webApp, employeeHandler)

	err := webApp.Listen(":3000")
	if err != nil {
		logegr.Fatal("failed to start server", zap.Error(err))
	}
}
