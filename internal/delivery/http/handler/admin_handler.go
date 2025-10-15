package handler

import (
	"net/http"
	"tracker/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct{
	userUseCase *usecase.UserUseCase
}

func NewAdminHandler(userUseCase *usecase.UserUseCase) *AdminHandler {
	return &AdminHandler{
		userUseCase: userUseCase,
	}
}

func (admH *AdminHandler) GetAllUsersHandler(c *gin.Context) {
	users, err := admH.userUseCase.GettAllUsers()

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}