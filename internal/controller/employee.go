package controller

import (
	"GoLessonFifteen/internal/errs"
	"GoLessonFifteen/internal/models"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllEmployees
// @Summary Получение данных пользователей
// @Description Получение данных всех пользователей
// @Tags Employees
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Employee
// @Failure 500 {object} CommonError
// @Router /api/employees [get]
func (ctrl *Controller) GetAllEmployees(c *gin.Context) {

	employee, err := ctrl.service.GetAllEmployees()
	if err != nil {
		ctrl.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, employee)
}

type CreateEmployeeRequest struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// CreateEmployee
// @Summary Добавление нового пользователя
// @Description Добавление нового пользователя
// @Tags Employees
// @Consume json
// @Produce json
// @Security BearerAuth
// @Param request_body body CreateEmployeeRequest true "информания о новом пользователе"
// @Success 201 {array} models.Employee
// @Failure 400 {object} CommonError
// @Failure 422 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /api/employees [post]
func (ctrl *Controller) CreateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		ctrl.HandleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}
	if employee.Name == "" || employee.Email == "" || employee.Age < 0 {
		ctrl.HandleError(c, errs.ErrInvalidFieldValuse)
		return
	}
	if err := ctrl.service.CreateEmployee(employee); err != nil {
		ctrl.HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, CommonResponse{Message: "Employee created successfully!"})
}

// GetEmployeeByID
// @Summary Получение данных пользователя по ID
// @Description Получение данных пользователя по ID
// @Tags Employees
// @Produce json
// @Security BearerAuth
// @Param id path int true "id продукта"
// @Success 200 {object} models.Employee
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /api/employees/{id} [get]
func (ctrl *Controller) GetEmployeeByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.HandleError(c, errs.ErrInvalidUserID)
		return
	}
	employee, err := ctrl.service.GetEmployeeByID(id)
	if err != nil {
		ctrl.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, employee)
}

// UpdateEmployeeByID
// @Summary Обновление данных пользователя по ID
// @Description Обновление данных пользователя по ID
// @Tags Employees
// @Consume json
// @Produce json
// @Security BearerAuth
// @Param id path int true "id продукта"
// @Param request_body body CreateEmployeeRequest true "информация о пользователе"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 422 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /api/employees/{id} [put]
func (ctrl *Controller) UpdateEmployeeByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.HandleError(c, errs.ErrInvalidUserID)
		return
	}
	var employee models.Employee
	if err = c.ShouldBindJSON(&employee); err != nil {
		ctrl.HandleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	if employee.Name == "" || employee.Email == "" || employee.Age < 0 {
		ctrl.HandleError(c, errs.ErrInvalidFieldValuse)
		return
	}
	employee.ID = id

	if err = ctrl.service.UpdateEmployeeByID(employee); err != nil {
		ctrl.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee updated successfully!",
	})
}

// DeleteEmployeeByID
// @Summary Удаление данных пользователя по ID
// @Description Удаление данных пользователя по ID
// @Tags Employees
// @Produce json
// @Security BearerAuth
// @Param id path int true "id продукта"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /api/employees/{id} [delete]
func (ctrl *Controller) DeleteEmployeeByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.HandleError(c, errs.ErrInvalidUserID)
		return
	}
	if err = ctrl.service.DeleteEmployeeByID(id); err != nil {
		ctrl.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, CommonResponse{Message: "Employee deleted successfully!"})
}
