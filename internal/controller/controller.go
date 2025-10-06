package controller

import (
	"GoLessonFifteen/internal/contracts"
	"GoLessonFifteen/internal/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
)

type Controller struct {
	router  *gin.Engine
	service contracts.ServiceI
	logger  zerolog.Logger
}

func NewController(service contracts.ServiceI, logger zerolog.Logger) *Controller {
	return &Controller{
		service: service,
		router:  gin.Default(),
		logger:  logger,
	}
}

func (ctrl *Controller) HandleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, errs.ErrUserNotFound) ||
		errors.Is(err, errs.ErrNotfound) ||
		errors.Is(err, errs.ErrEmployeeNotFound):
		c.JSON(http.StatusNotFound, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidUserID) || errors.Is(err, errs.ErrInvalidRequestBody):
		c.JSON(http.StatusBadRequest, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrIncorrectUsernameOrPassword) ||
		errors.Is(err, errs.ErrInvalidToken):
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidFieldValuse) ||
		errors.Is(err, errs.ErrInvalidUserName) ||
		errors.Is(err, errs.ErrUsernameAlreadyExists):
		c.JSON(http.StatusUnprocessableEntity, CommonError{Error: err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
	}
}
