package contracts

import (
	"GoLessonFifteen/internal/models"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type ServiceI interface {
	GetAllUsers() (users []models.User, err error)
	CreateUser(user models.User) (err error)
	GetUserByID(id int) (user models.User, err error)
	UpdateUserByID(user models.User) (err error)
	DeleteUserByID(id int) (err error)
}
