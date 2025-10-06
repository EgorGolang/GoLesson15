package service

import (
	"GoLessonFifteen/internal/errs"
	"GoLessonFifteen/internal/models"
	"GoLessonFifteen/utils"
	"context"
	"errors"
)

func (s *Service) CreateEmployees(ctx context.Context, employees models.Employee) (err error) {
	_, err = s.repository.GetEmployeesByUsername(ctx, employees.Username)
	if err != nil {
		if !errors.Is(err, errs.ErrNotfound) {
			return err
		}
	} else {
		return errs.ErrUsernameAlreadyExists
	}
	employees.Password, err = utils.GenerateHash(employees.Password)
	if err != nil {
		return err
	}

	employees.Role = models.RoleUser

	if err = s.repository.CreateEmployees(ctx, employees); err != nil {
		return err
	}
	return nil
}

func (s *Service) Authentificate(ctx context.Context, employees models.Employee) (int, models.Role, error) {
	employeeFromDB, err := s.repository.GetEmployeesByUsername(ctx, employees.Username)
	if err != nil {
		if !errors.Is(err, errs.ErrNotfound) {
			return 0, "", errs.ErrEmployeeNotFound
		}
		return 0, "", err
	}
	employees.Password, err = utils.GenerateHash(employees.Password)
	if err != nil {
		return 0, "", err
	}
	if employees.Password != employeeFromDB.Password {
		return 0, "", errs.ErrIncorrectUsernameOrPassword
	}

	return employeeFromDB.ID, employeeFromDB.Role, nil
}
