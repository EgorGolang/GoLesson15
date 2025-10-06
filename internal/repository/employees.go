package repository

import (
	"GoLessonFifteen/internal/models"
	"context"
	"github.com/rs/zerolog"
	"os"
)

func (r *Repository) CreateEmployees(ctx context.Context, employee models.Employee) (err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "Repository.CreateEmployees").Logger()
	_, err = r.db.ExecContext(ctx, `INSERT INTO employees (full_name, username, password, role) VALUES ($1, $2, $3, $4)`,
		employee.FullName,
		employee.Username,
		employee.Password,
		employee.Role)
	if err != nil {
		logger.Err(err).Msg("Error inserting employee")
		return r.TranslateError(err)
	}
	return nil
}

func (r *Repository) GetEmployeesByID(ctx context.Context, id int) (employees models.Employee, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "Repository.GetEmployeesById").Logger()
	if err = r.db.GetContext(ctx, &employees, `SELECT id, full_name, username, password, role, created_at, updated_at 
	FROM employees WHERE id = $1`, id); err != nil {
		logger.Err(err).Msg("Error selecting employee")
		return models.Employee{}, r.TranslateError(err)
	}
	return employees, nil
}

func (r *Repository) GetEmployeesByUsername(ctx context.Context, username string) (employee models.Employee, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "Repository.GetEmployeesById").Logger()
	if err = r.db.GetContext(ctx, &employee, `SELECT id, full_name, username, password, role, created_at, updated_at 
	FROM employees WHERE username = $1`, username); err != nil {
		logger.Err(err).Msg("Error selecting employee")
		return models.Employee{}, r.TranslateError(err)
	}
	return employee, nil
}
