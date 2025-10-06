package controller

import (
	"GoLessonFifteen/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	authorizationHeder = "Authorization"
	employeeIDCtx      = "employeeID"
	employeeRoleCtx    = "employeeRole"
)

func (ctrl *Controller) checkEmployeeAuthentification(c *gin.Context) {
	token, err := ctrl.extractTokenFromHeader(c, authorizationHeder)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}
	employeeId, isRefresh, employeeRole, err := pkg.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}
	if isRefresh {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: "inappropriate token"})
		return
	}
	c.Set(employeeIDCtx, employeeId)
	c.Set(employeeRoleCtx, string(employeeRole))
}

func (ctrl *Controller) checkIsAdmid(c *gin.Context) {
	role := c.GetString(employeeRoleCtx)
	if role == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: "role is not in context"})
		return
	}

	if role != "ADMIN" {
		c.AbortWithStatusJSON(http.StatusForbidden, CommonError{Error: "permission denied"})
		return
	}

	c.Next()
}
