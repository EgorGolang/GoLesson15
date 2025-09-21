package controller

import (
	"GoLessonFifteen/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllUsers
// @Summary Получение данных пользователей
// @Description Получение данных всех пользователей
// @Tags Users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} CommonError
// @Router /users [get]
func (ctrl *Controller) GetAllUsers(c *gin.Context) {
	user, err := ctrl.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// GetCreateUser
// @Summary Добавление нового пользователя
// @Description Добавление нового пользователя
// @Tags Users
// @Consume json
// @Produce json
// @Param request_body body CreateUserRequest true "информания о новом пользователе"
// @Success 201 {array} models.User
// @Failure 400 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users [post]
func (ctrl *Controller) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := ctrl.service.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, CommonResponse{Message: "User created successfully!"})
}

// GetUserByID
// @Summary Получение данных пользователя по ID
// @Description Получение данных пользователя по ID
// @Tags Users
// @Produce json
// @Param id path int true "id продукта"
// @Success 200 {object} models.User
// @Failure 400 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [get]
func (ctrl *Controller) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommonError{Error: err.Error()})
		return
	}
	user, err := ctrl.service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUserByID
// @Summary Обновление данных пользователя по ID
// @Description Обновление данных пользователя по ID
// @Tags Users
// @Consume json
// @Produce json
// @Param id path int true "id продукта"
// @Param request_body body CreateUserRequest true "информация о пользователе"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [put]
func (ctrl *Controller) UpdateUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var user models.User
	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user.ID = id
	if err = ctrl.service.UpdateUserByID(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully!",
	})
}

// DeleteUserByID
// @Summary Удаление данных пользователя по ID
// @Description Удаление данных пользователя по ID
// @Tags Users
// @Produce json
// @Param id path int true "id продукта"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [delete]
func (ctrl *Controller) DeleteUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err = ctrl.service.DeleteUserByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, CommonResponse{Message: "User deleted successfully!"})
}
