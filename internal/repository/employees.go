package repository

import (
	"GoLessonFifteen/internal/models"
	"context"
)

func (r *Repository) GetAllEmployees(ctx context.Context) (employee []models.Employee, err error) {
	err = r.db.SelectContext(ctx, &employee, `SELECT name, email, age, id FROM employees ORDER BY id`)
	if err != nil {
		return nil, r.TranslateError(err)
	}
	return employee, err
}

func (r *Repository) CreateEmployee(ctx context.Context, employee models.Employee) (err error) {
	_, err = r.db.ExecContext(ctx, `INSERT INTO employees (name, email, age) VALUES ($1, $2, $3)`,
		employee.Name,
		employee.Email,
		employee.Age)
	if err != nil {
		return r.TranslateError(err)
	}
	return nil
}

func (r *Repository) GetEmployeeByID(ctx context.Context, id int) (employee models.Employee, err error) {
	if err = r.db.GetContext(ctx, &employee, `SELECT id, name, email, age FROM employees WHERE id = $1`, id); err != nil {
		return models.Employee{}, r.TranslateError(err)
	}
	return employee, nil
}

func (r *Repository) UpdateEmployeeByID(ctx context.Context, employee models.Employee) (err error) {
	_, err = r.db.ExecContext(ctx, `UPDATE employees SET name=$1, email=$2, age=$3 WHERE id = $4`,
		employee.Name,
		employee.Email,
		employee.Age,
		employee.ID)
	if err != nil {
		return r.TranslateError(err)
	}
	return nil
}

func (r *Repository) DeleteEmployeeByID(ctx context.Context, id int) (err error) {
	_, err = r.db.ExecContext(ctx, `DELETE FROM employees WHERE id = $1`, id)
	if err != nil {
		return r.TranslateError(err)
	}
	return nil
}
