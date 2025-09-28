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

	err := sh.subscriptionUseCase.CreateSubscription(sub)

	if err != nil{
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "subscription successful created"})
}

func (sh *SubscriptionHandler) HandlerGetAll(c *gin.Context) {
	sub, err := sh.subscriptionUseCase.GetAllSubscriptions()

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": sub.ID,
		"name": sub.Name,
		"amount": sub.Amount,
		"started_at": sub.StartDate,
		"period": sub.BillingPeriod,
		"next_pay": sub.NextBillingDate,
	})
}