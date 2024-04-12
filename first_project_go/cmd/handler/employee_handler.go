package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"project/cmd/repository"
	"project/cmd/service"
)

type EmployeeHandler struct {
	service *service.EmployeeService
}

type CreateEmployeeRequest struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

type CreateEmployeeResponse struct {
	ID string `json:"id"`
}

type EmployeePayload struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type ListEmployeesResponse struct {
	Employees []EmployeePayload `json:"employees"`
}

type GetEmployeeResponse struct {
	EmployeePayload
}

type UpdateEmployeeRequest struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

func NewEmployeeHandler(service *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

func (h *EmployeeHandler) CreateEmployee(ctx *fiber.Ctx) error {
	var req CreateEmployeeRequest
	if err := ctx.BodyParser(&req); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}

	id, err := h.service.CreateEmployee(repository.Employee{
		Email: req.Email,
		Role:  req.Role,
	})

	if err != nil {
		return fmt.Errorf("create in storage: %w", err)
	}
	return ctx.JSON(CreateEmployeeResponse{ID: id})
}

func (h *EmployeeHandler) ListEmployees(ctx *fiber.Ctx) error {
	employees := h.service.ListEmployee()

	resp := ListEmployeesResponse{
		Employees: make([]EmployeePayload, len(employees)),
	}
	for i, empl := range employees {
		resp.Employees[i] = EmployeePayload(empl)
	}
	return ctx.JSON(resp)
}

func (h *EmployeeHandler) GetEmployee(ctx *fiber.Ctx) error {
	empl, err := h.service.GetEmployee(ctx.Params("id"))
	if err != nil {
		return fiber.ErrNotFound
	}
	return ctx.JSON(GetEmployeeResponse{EmployeePayload(empl)})
}

func (h *EmployeeHandler) UpdateEmployee(ctx *fiber.Ctx) error {
	var req UpdateEmployeeRequest
	if err := ctx.BodyParser(&req); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}
	err := h.service.UpdateEmployee(ctx.Params("id"), req.Email, req.Role)
	if err != nil {
		return fmt.Errorf("update: %w", err)
	}
	return nil
}

func (h *EmployeeHandler) DeleteEmployee(ctx *fiber.Ctx) error {
	h.service.DeleteEmployee(ctx.Params("id"))
	return ctx.SendStatus(fiber.StatusNoContent)
}
