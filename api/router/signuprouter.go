package router

import (
	"task/api/controller"

	"github.com/gin-gonic/gin"
)

func RegisterSignupRoutes(r *gin.Engine, signUpController *controller.SignUpController) {
	r.POST("/register", signUpController.Register)
}
