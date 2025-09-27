package handler

import (
	"net/http"
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
	c.JSON(http.StatusOK, gin.H{"message": "good job!"})
}