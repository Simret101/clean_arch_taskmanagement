package router

import (
	"task/api/controller"
	"task/api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	r *gin.Engine,
	signUpController *controller.SignUpController,
	loginController *controller.LoginController,
	taskController *controller.TaskController,
	authMiddleware *middleware.AuthMiddleware,
) {

	RegisterSignupRoutes(r, signUpController)
	RegisterLoginRoutes(r, loginController)
	RegisterTaskRoutes(r, taskController, authMiddleware)
}
