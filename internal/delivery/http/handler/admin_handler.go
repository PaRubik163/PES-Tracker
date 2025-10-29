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

func (admH *AdminHandler) GetUserByIDHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	
	if userID == 0{
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing user_id"})
		return
	}
	
	user, err := admH.userUseCase.GetUserByID(userID)

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, user)
}