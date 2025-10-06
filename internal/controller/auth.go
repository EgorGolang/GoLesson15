package controller

import (
	"GoLessonFifteen/internal/errs"
	"GoLessonFifteen/internal/models"
	"GoLessonFifteen/pkg"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignUpRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

// SingUp
// @Summary Регистрация
// @Description Создать новый аккаунт
// @Tags Auth
// @Consume json
// @Produce json
// @Param request_body body SignUpRequest true "информация о новом аккаунте"
// @Success 201 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /auth/sign-up [post]
func (ctrl *Controller) SignUp(c *gin.Context) {
	var input SignUpRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		ctrl.HandleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}
	if err := ctrl.service.CreateEmployees(c, models.Employee{
		FullName: input.FullName,
		Password: input.Password,
		Username: input.Username,
	}); err != nil {
		ctrl.HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, CommonResponse{Message: "Employees successfully created"})
}

type SignInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenPairResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// SingIn
// @Summary Вход
// @Description Войти в аккаунт
// @Tags Auth
// @Consume json
// @Produce json
// @Param request_body body SignInRequest true "логин и пароль"
// @Success 200 {object} TokenPairResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /auth/sign-in [post]
func (ctrl *Controller) SignIn(c *gin.Context) {
	var input SignInRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		ctrl.HandleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	employeeID, employeeRole, err := ctrl.service.Authentificate(c, models.Employee{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		ctrl.HandleError(c, err)
		return
	}

	accessToken, refreshToken, err := ctrl.generateNewTokenPair(employeeID, employeeRole)
	if err != nil {
		ctrl.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, TokenPairResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

const (
	refreshTokenHeader = "X-Refresh-Token"
)

// RefreshTokenPair
// @Summary Обновить пару токенов
// @Description Обновить пару токенов
// @Tags Auth
// @Produce json
// @Param X-Refresh-Token header string true "вставьте refresh token"
// @Success 200 {object} TokenPairResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /auth/refresh [get]
func (ctrl *Controller) RefreshTokenPair(c *gin.Context) {
	token, err := ctrl.extractTokenFromHeader(c, refreshTokenHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	employeeID, isRefresh, employeeRole, err := pkg.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}
	if !isRefresh {
		c.JSON(http.StatusUnauthorized, CommonError{Error: "inappropriate token"})
		return
	}
	accessToken, refreshToken, err := ctrl.generateNewTokenPair(employeeID, employeeRole)
	if err != nil {
		ctrl.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, TokenPairResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
