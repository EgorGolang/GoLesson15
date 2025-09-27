package contracts

import "GoLessonFifteen/internal/models"

type RepositoryI interface {
	GetAllUser() (user []models.User, err error)
	CreateUser(user models.User) (err error)
	GetUserByID(id int) (user models.User, err error)
	UpdateUserByID(user models.User) (err error)
	DeleteUserByID(id int) (err error)
}
