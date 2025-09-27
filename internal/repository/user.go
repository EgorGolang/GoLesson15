package repository

import (
	"GoLessonFifteen/internal/models"
)

func (r *Repository) GetAllUser() (user []models.User, err error) {
	err = r.db.Select(&user, `SELECT name, email, age, id FROM users ORDER BY id`)
	if err != nil {
		return nil, r.TranslateError(err)
	}
	return user, err
}

func (r *Repository) CreateUser(user models.User) (err error) {
	_, err = r.db.Exec(`INSERT INTO users (name, email, age) VALUES ($1, $2, $3)`,
		user.Name,
		user.Email,
		user.Age)
	if err != nil {
		return r.TranslateError(err)
	}
	return nil
}

func (r *Repository) GetUserByID(id int) (user models.User, err error) {
	if err = r.db.Get(&user, `SELECT id, name, email, age FROM users WHERE id = $1`, id); err != nil {
		return models.User{}, r.TranslateError(err)
	}
	return user, nil
}

func (r *Repository) UpdateUserByID(user models.User) (err error) {
	_, err = r.db.Exec(`UPDATE users SET name=$1, email=$2, age=$3 WHERE id = $4`,
		user.Name,
		user.Email,
		user.Age,
		user.ID)
	if err != nil {
		return r.TranslateError(err)
	}
	return nil
}

func (r *Repository) DeleteUserByID(id int) (err error) {
	_, err = r.db.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return r.TranslateError(err)
	}
	return nil
}
