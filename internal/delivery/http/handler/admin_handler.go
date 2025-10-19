package handler

import (
	"net/http"
	"strconv"
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

	c.JSON(http.StatusOK, gin.H{"users": users,
								"total": len(users)})
}

func (admH *AdminHandler) GetUserByIDHandler(c *gin.Context) {
	userIDStr := c.Param("id")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	user, err := admH.userUseCase.GetUserByID(userID)
	
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, user)
}