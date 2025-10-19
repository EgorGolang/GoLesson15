package service

import (
	"GoLessonFifteen/internal/errs"
	"GoLessonFifteen/internal/models"
	"GoLessonFifteen/utils"
	"context"
	"errors"
)

func (s *Service) CreateUsers(ctx context.Context, users models.User) (err error) {
	_, err = s.repository.GetUsersByUsername(ctx, users.Username)
	if err != nil {
		if !errors.Is(err, errs.ErrNotfound) {
			return err
		}
	} else {
		return errs.ErrUsernameAlreadyExists
	}
	users.Password, err = utils.GenerateHash(users.Password)
	if err != nil {
		return err
	}

	users.Role = models.RoleUser

	if err = s.repository.CreateUser(ctx, users); err != nil {
		return err
	}
	return nil
}

func (s *Service) Authentificate(ctx context.Context, users models.User) (int, models.Role, error) {
	userFromDB, err := s.repository.GetUsersByUsername(ctx, users.Username)
	if err != nil {
		if !errors.Is(err, errs.ErrNotfound) {
			return 0, "", errs.ErrEmployeeNotFound
		}
		return 0, "", err
	}
	users.Password, err = utils.GenerateHash(users.Password)
	if err != nil {
		return 0, "", err
	}
	if users.Password != userFromDB.Password {
		return 0, "", errs.ErrIncorrectUsernameOrPassword
	}

	return userFromDB.ID, userFromDB.Role, nil
}
