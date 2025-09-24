package handler

import (
	"net/http"
	"tracker/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{
	userUC *usecase.UserUseCase
}

type request struct{
	Login string  	 `json:"login" binding:"required"`
	Password string  `json:"password" binding:"required"`
}

func NewUserHandler(us *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUC: us,
	}
}

func (uh *UserHandler) HandlerRegister(c *gin.Context) {
	var req request

	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	}

	err := uh.userUC.Register(req.Login, req.Password)
	if err != nil{
		switch err.Error(){
		case "user already exists":
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 
		}
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user registered"})
}

func (uh *UserHandler) HandlerLogin(c *gin.Context) {
	var req request

	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uh.userUC.Login(req.Login, req.Password)
	
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messgae": "login successful",
		"user": gin.H{
			"id": user.ID,
			"login": user.Login,
		},
		"token": "JWT token",
	})
}