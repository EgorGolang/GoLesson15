package contracts

import (
	"GoLessonFifteen/internal/models"
	"context"
)

type RepositoryI interface {
	GetAllEmployees(ctx context.Context) (employees []models.Employee, err error)
	CreateEmployee(ctx context.Context, employee models.Employee) (err error)
	GetEmployeeByID(ctx context.Context, id int) (employee models.Employee, err error)
	UpdateEmployeeByID(ctx context.Context, employee models.Employee) (err error)
	DeleteEmployeeByID(ctx context.Context, id int) (err error)
	CreateUser(ctx context.Context, user models.User) (err error)
	GetUsersByID(ctx context.Context, id int) (user models.User, err error)
	GetUsersByUsername(ctx context.Context, username string) (user models.User, err error)
}
