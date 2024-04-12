package router

import (
	"github.com/gofiber/fiber/v3"
	"project/cmd/handler"
)

func SetupRoutes(app *fiber.App, employeeHandler *handler.EmployeeHandler) {
	app.Post("/employees", employeeHandler.CreateEmployee)
	app.Get("/employees", employeeHandler.ListEmployees)
	app.Get("/employees/:id", employeeHandler.GetEmployee)
	app.Patch("/employees/:id", employeeHandler.UpdateEmployee)
	app.Delete("employees/:id", employeeHandler.DeleteEmployee)
}
