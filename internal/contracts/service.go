package contracts

import (
	"GoLessonFifteen/internal/models"
	"context"
)

type ServiceI interface {
	GetAllUsers(ctx context.Context) (users []models.User, err error)
	CreateUser(ctx context.Context, user models.User) (err error)
	GetUserByID(ctx context.Context, id int) (user models.User, err error)
	UpdateUserByID(ctx context.Context, user models.User) (err error)
	DeleteUserByID(ctx context.Context, id int) (err error)
}
