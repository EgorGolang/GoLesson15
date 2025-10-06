package controller

import (
	"GoLessonFifteen/internal/configs"
	"GoLessonFifteen/internal/models"
	"GoLessonFifteen/pkg"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func (ctrl *Controller) extractTokenFromHeader(c *gin.Context, headerKey string) (string, error) {
	header := c.GetHeader(headerKey)

	if header == "" {
		return "", errors.New("empty authorization header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return "", errors.New("invalid authorization header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("empty token")
	}
	return headerParts[1], nil
}

func (ctrl *Controller) generateNewTokenPair(employeeID int, employeeRole models.Role) (string, string, error) {
	accessToken, err := pkg.GenerateToken(employeeID,
		configs.AppSettings.AuthParams.AccessTokenTtlMinutes,
		employeeRole, false)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := pkg.GenerateToken(employeeID,
		configs.AppSettings.AuthParams.RefreshTokenTtlDays,
		employeeRole, true)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}
