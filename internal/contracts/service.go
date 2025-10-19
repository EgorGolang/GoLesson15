package contracts

import (
	"GoLessonFifteen/internal/models"
	"context"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type ServiceI interface {
	GetAllEmployees() (employees []models.Employee, err error)
	CreateEmployee(employee models.Employee) (err error)
	GetEmployeeByID(id int) (employee models.Employee, err error)
	UpdateEmployeeByID(employee models.Employee) (err error)
	DeleteEmployeeByID(id int) (err error)
	CreateUsers(ctx context.Context, users models.User) (err error)
	Authentificate(ctx context.Context, users models.User) (int, models.Role, error)
}
