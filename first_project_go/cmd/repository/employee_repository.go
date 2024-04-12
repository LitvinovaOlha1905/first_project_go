package repository

import (
	"errors"
	"github.com/google/uuid"
)

type Employee struct {
	ID    string
	Email string
	Role  string
}

type EmployeeRepository interface {
	Create(Employee) (string, error)
	List() []Employee
	Get(id string) (Employee, error)
	Update(id, email, role string) error
	Delete(id string)
	GetByEmail(email string) (Employee, error)
}

type memoryEmployeeRepository struct {
	employees map[string]Employee
}

func NewMemoryEmployeeRepository() *memoryEmployeeRepository {
	return &memoryEmployeeRepository{
		employees: make(map[string]Employee),
	}
}

func (r *memoryEmployeeRepository) Create(empl Employee) (string, error) {
	empl.ID = uuid.New().String()
	r.employees[empl.ID] = empl
	return empl.ID, nil
}

func (m *memoryEmployeeRepository) GetByEmail(email string) (Employee, error) {
	for _, e := range m.employees {
		if e.Email == email {
			return e, nil
		}
	}
	return Employee{}, errors.New("employee not found")
}

func (r *memoryEmployeeRepository) List() []Employee {
	employees := make([]Employee, 0, len(r.employees))
	for _, empl := range r.employees {
		employees = append(employees, empl)
	}
	return employees
}

func (r *memoryEmployeeRepository) Get(id string) (Employee, error) {
	empl, ok := r.employees[id]
	if !ok {
		return Employee{}, errors.New("employee not found")
	}
	return empl, nil
}

func (r *memoryEmployeeRepository) Update(id, email, role string) error {
	empl, ok := r.employees[id]
	if !ok {
		return errors.New("employee not found")
	}
	if email != "" {
		empl.Email = email
	}
	if role != "" {
		empl.Role = role
	}
	r.employees[empl.ID] = empl
	return nil
}

func (r *memoryEmployeeRepository) Delete(id string) {
	delete(r.employees, id)
}
