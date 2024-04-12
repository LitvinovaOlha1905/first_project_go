package service

import (
	"project/cmd/repository"
)

type EmployeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) CreateEmployee(empl repository.Employee) (string, error) {
	return s.repo.Create(empl)
}

func (s *EmployeeService) ListEmployee() []repository.Employee {
	return s.repo.List()
}

func (s *EmployeeService) GetEmployee(id string) (repository.Employee, error) {
	return s.repo.Get(id)
}
func (s *EmployeeService) UpdateEmployee(id, email, role string) error {
	return s.repo.Update(id, email, role)
}

func (s *EmployeeService) DeleteEmployee(id string) {
	s.repo.Delete(id)
}
