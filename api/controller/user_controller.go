package controller

import (
	"net/http"
	"task/domain"
	"task/usecase"

	"github.com/gin-gonic/gin"
)

type SignUpController struct {
	UserUC *usecase.RegisterUsecase
}

func NewSignUpController(userUC *usecase.RegisterUsecase) *SignUpController {
	return &SignUpController{UserUC: userUC}
}

func (uc *SignUpController) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := uc.UserUC.CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}
