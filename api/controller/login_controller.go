package controller

import (
	"net/http"
	"task/domain"
	"task/usecase"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	UserUC *usecase.LoginUsecase
}

func NewLoginController(userUC *usecase.LoginUsecase) *LoginController {
	return &LoginController{UserUC: userUC}
}

func (uc *LoginController) Login(c *gin.Context) {
	var credentials domain.Credentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := uc.UserUC.AuthenticateUser(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
