package controller

import (
	_ "GoLessonFifteen/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func (ctrl Controller) RegisterEndpoints() {
	ctrl.router.GET("/ping", ctrl.Ping)
	ctrl.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	authG := ctrl.router.Group("/auth")
	{
		authG.POST("/sign-up", ctrl.SignUp)
		authG.POST("/sign-in", ctrl.SignIn)
		authG.GET("/refresh", ctrl.RefreshTokenPair)
	}

	apiG := ctrl.router.Group("/api", ctrl.checkEmployeeAuthentification)
	{
		apiG.GET("/users", ctrl.GetAllUsers)
		apiG.POST("/users", ctrl.checkIsAdmid, ctrl.CreateUser)
		apiG.GET("/users/:id", ctrl.GetUserByID)
		apiG.PUT("/users/:id", ctrl.checkIsAdmid, ctrl.UpdateUserByID)
		apiG.DELETE("/users/:id", ctrl.checkIsAdmid, ctrl.DeleteUserByID)
	}

}

// Ping
// @Summary Health-Check
// @Description Проверка сервиса
// @Tags Ping
// @Produce json
// @Success 200 {object} CommonResponse
// @Router /ping [get]
func (ctrl *Controller) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, CommonResponse{Message: "Server is running"})
}

func (ctrl *Controller) RunServer(address string) error {
	ctrl.RegisterEndpoints()
	if err := ctrl.router.Run(address); err != nil {
		return err
	}

	return nil
}
