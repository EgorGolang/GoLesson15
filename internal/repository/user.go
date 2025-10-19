package repository

import (
	"GoLessonFifteen/internal/models"
	"context"
	"github.com/rs/zerolog"
	"os"
)

func (r *Repository) CreateUser(ctx context.Context, user models.User) (err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "Repository.CreateUser").Logger()
	_, err = r.db.ExecContext(ctx, `INSERT INTO users (full_name, username, password, role) VALUES ($1, $2, $3, $4)`,
		user.FullName,
		user.Username,
		user.Password,
		user.Role)
	if err != nil {
		logger.Err(err).Msg("Error inserting user")
		return r.TranslateError(err)
	}
	return nil
}

func (r *Repository) GetUsersByID(ctx context.Context, id int) (users models.User, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "Repository.GetUsersById").Logger()
	if err = r.db.GetContext(ctx, &users, `SELECT id, full_name, username, password, role, created_at, updated_at 
	FROM users WHERE id = $1`, id); err != nil {
		logger.Err(err).Msg("Error selecting user")
		return models.User{}, r.TranslateError(err)
	}
	return users, nil
}

func (r *Repository) GetUsersByUsername(ctx context.Context, username string) (user models.User, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "Repository.GetUsersById").Logger()
	if err = r.db.GetContext(ctx, &user, `SELECT id, full_name, username, password, role, created_at, updated_at 
	FROM users WHERE username = $1`, username); err != nil {
		logger.Err(err).Msg("Error selecting user")
		return models.User{}, r.TranslateError(err)
	}
	return user, nil
}
