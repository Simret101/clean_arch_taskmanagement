package router

import (
	"task/api/controller"

	"github.com/gin-gonic/gin"
)

func RegisterLoginRoutes(r *gin.Engine, loginController *controller.LoginController) {
	r.POST("/login", loginController.Login)
}
