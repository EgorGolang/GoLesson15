package repository

import (
	"GoLessonFifteen/internal/errs"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	_ "github.com/rs/zerolog/log"
)

type Repository struct {
	db     *sqlx.DB
	logger zerolog.Logger
}

func NewRepository(db *sqlx.DB, logger zerolog.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger}
}

func (r *Repository) TranslateError(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return errs.ErrNotfound
	default:
		return err
	}
}
