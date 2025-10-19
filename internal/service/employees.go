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

func (s *Service) GetAllEmployees() (employees []models.Employee, err error) {
	ctx := context.Background()
	employees, err = s.repository.GetAllEmployees(ctx)
	if err != nil {
		return nil, err
	}
	return employees, nil
}
func (s *Service) CreateEmployee(employee models.Employee) (err error) {
	ctx := context.Background()
	if len(employee.Name) < 4 {
		return errs.ErrInvalidUserName
	}

	err = s.repository.CreateEmployee(ctx, employee)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetEmployeeByID(id int) (employee models.Employee, err error) {
	ctx := context.Background()
	err = s.cache.Get(ctx, fmt.Sprintf("user_%d", id), &employee)
	if err == nil {
		return employee, nil
	}
	employee, err = s.repository.GetEmployeeByID(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return models.Employee{}, errs.ErrUserNotFound
		}
		return models.Employee{}, err
	}
	if err = s.cache.Set(ctx, fmt.Sprintf("user_%d", employee.ID), employee, defaultTTL); err != nil {
		fmt.Printf("error during cache set: %v", err.Error())
	}
	return employee, nil
}

func (s *Service) UpdateEmployeeByID(employee models.Employee) (err error) {
	ctx := context.Background()
	_, err = s.repository.GetEmployeeByID(ctx, employee.ID)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return errs.ErrUserNotFound
		}
		return err
	}

	err = s.repository.UpdateEmployeeByID(ctx, employee)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteEmployeeByID(id int) (err error) {
	ctx := context.Background()
	_, err = s.repository.GetEmployeeByID(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return errs.ErrUserNotFound
		}
		return err
	}

	err = s.repository.DeleteEmployeeByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
