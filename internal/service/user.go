package service

import (
	"GoLessonFifteen/internal/errs"
	"GoLessonFifteen/internal/models"
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	defaultTTL = time.Minute * 5
)

func (s *Service) GetAllUsers() (users []models.User, err error) {
	ctx := context.Background()
	users, err = s.repository.GetAllUser(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (s *Service) CreateUser(user models.User) (err error) {
	ctx := context.Background()
	if len(user.Name) < 4 {
		return errs.ErrInvalidUserName
	}

	err = s.repository.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetUserByID(id int) (user models.User, err error) {
	ctx := context.Background()
	err = s.cache.Get(ctx, fmt.Sprintf("user_%d", id), &user)
	if err == nil {
		return user, nil
	}
	user, err = s.repository.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return models.User{}, errs.ErrUserNotFound
		}
		return models.User{}, err
	}
	if err = s.cache.Set(ctx, fmt.Sprintf("user_%d", user.ID), user, defaultTTL); err != nil {
		fmt.Printf("error during cache set: %v", err.Error())
	}
	return user, nil
}

func (s *Service) UpdateUserByID(user models.User) (err error) {
	ctx := context.Background()
	_, err = s.repository.GetUserByID(ctx, user.ID)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return errs.ErrUserNotFound
		}
		return err
	}

	err = s.repository.UpdateUserByID(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteUserByID(id int) (err error) {
	ctx := context.Background()
	_, err = s.repository.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return errs.ErrUserNotFound
		}
		return err
	}

	err = s.repository.DeleteUserByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
