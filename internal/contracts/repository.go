package contracts

import (
	"GoLessonFifteen/internal/models"
	"context"
)

type RepositoryI interface {
	GetAllUser(ctx context.Context) (user []models.User, err error)
	CreateUser(ctx context.Context, user models.User) (err error)
	GetUserByID(ctx context.Context, id int) (user models.User, err error)
	UpdateUserByID(ctx context.Context, user models.User) (err error)
	DeleteUserByID(ctx context.Context, id int) (err error)
	CreateEmployees(ctx context.Context, employees models.Employee) (err error)
	GetEmployeesByID(ctx context.Context, id int) (employees models.Employee, err error)
	GetEmployeesByUsername(ctx context.Context, username string) (employee models.Employee, err error)
}
