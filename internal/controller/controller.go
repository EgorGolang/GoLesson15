package controller

import (
	"GoLessonFifteen/internal/contracts"
	"GoLessonFifteen/internal/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	router  *gin.Engine
	service contracts.ServiceI
}

func NewController(service contracts.ServiceI) *Controller {
	return &Controller{
		service: service,
		router:  gin.Default(),
	}
}

func (ctrl *Controller) HandleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, errs.ErrUserNotFound) || errors.Is(err, errs.ErrNotfound):
		c.JSON(http.StatusNotFound, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidUserID) || errors.Is(err, errs.ErrInvalidRequestBody):
		c.JSON(http.StatusBadRequest, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidFieldValuse) || errors.Is(err, errs.ErrInvalidUserName):
		c.JSON(http.StatusUnprocessableEntity, CommonError{Error: err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
	}
}
