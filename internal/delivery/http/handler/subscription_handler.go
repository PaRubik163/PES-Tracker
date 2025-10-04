package handler

import (
	"net/http"
	"strconv"
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
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authorized"})
		return
	}
	
	sub := &entity.Subscription{}

	if err := c.ShouldBindJSON(sub); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
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

func (sh *SubscriptionHandler) HandlerDeleteSubscription(c *gin.Context) {
	subIdStr := c.Param("id")

	if subIdStr == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error" : "invalid request"})
		return
	}

	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	id, err := strconv.Atoi(subIdStr)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := sh.subscriptionUseCase.DeleteSubscription(id, userID.(int)); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "subscription successful deleted"})
}