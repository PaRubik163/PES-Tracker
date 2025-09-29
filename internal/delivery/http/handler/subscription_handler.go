package handler

import (
	"net/http"
	"tracker/internal/entity"
	"tracker/internal/usecase"

	"github.com/gin-gonic/gin"
)

type SubscriptionHandler struct {
	subscriptionUseCase *usecase.SubscriptionUseCase
}

func NewSubscriptionHandler(susecase *usecase.SubscriptionUseCase) *SubscriptionHandler{
	return &SubscriptionHandler{
		subscriptionUseCase: susecase,
	}
}

func (sh *SubscriptionHandler) HandlerAdd(c *gin.Context) {
	sub := &entity.Subscription{}

	if err := c.ShouldBindJSON(sub); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid requset"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authorized"})
		return
	}

	sub.UserID = userID.(int)

	err := sh.subscriptionUseCase.CreateSubscription(sub)

	if err != nil{
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "subscription successful created"})
}

func (sh *SubscriptionHandler) HandlerGetAll(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authorized"})
		return
	}

	subs, err := sh.subscriptionUseCase.GetAllSubscriptions(userID.(int))

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, subs)
}