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
		"message": "Welcome to the PES Tracker",
		"token": user.Token,
		"user": gin.H{
			"id": user.ID,
			"login": user.Login,
			"create_session_at": user.CreateSessionAt,
		},
	})
}

func (uh *UserHandler) HandlerLogout(c *gin.Context) {
	uuid, exists := c.Get("uuid")
	if !exists{
		c.JSON(http.StatusUnauthorized, gin.H{"error":"invalid token"})
	}

	if err := uh.userUC.Logout(uuid.(string)); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "logout successful"})
}

func (uh *UserHandler) HandlerGetMe(c *gin.Context){
	uuid, exists := c.Get("uuid")

	if !exists{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}

	userSession, err := uh.userUC.GetMe(uuid.(string))

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": userSession.ID,
		"login": userSession.Login,
		"token": userSession.Token,
		"subscriptions_quantity": userSession.SubscriptionsQuantity,
		"expenses_month": userSession.ExpensesMonth,
		"income_month": userSession.IncomeMonth,
		"created_session_at": userSession.CreateSessionAt,
	})
}