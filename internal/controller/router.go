package controller

import (
	_ "GoLessonFifteen/docs"
	"github.com/gin-contrib/pprof"
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

	apiG := ctrl.router.Group("/api", ctrl.checkUserAuthentification)
	{
		apiG.GET("/employees", ctrl.GetAllEmployees)
		apiG.POST("/employees", ctrl.checkIsAdmid, ctrl.CreateEmployee)
		apiG.GET("/employees/:id", ctrl.GetEmployeeByID)
		apiG.PUT("/employees/:id", ctrl.checkIsAdmid, ctrl.UpdateEmployeeByID)
		apiG.DELETE("/employees/:id", ctrl.checkIsAdmid, ctrl.DeleteEmployeeByID)
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
	pprof.Register(ctrl.router)
	ctrl.RegisterEndpoints()
	if err := ctrl.router.Run(address); err != nil {
		return err
	}

	return nil
}
