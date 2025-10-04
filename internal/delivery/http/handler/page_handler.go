package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageHandler struct{}

func NewPageHandler() *PageHandler {
	return &PageHandler{}
}

func (pgh *PageHandler) Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func (pgh *PageHandler) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func (pgh *PageHandler) GetMe(c *gin.Context) {
	c.HTML(http.StatusOK, "profile.html", nil)
}

func (pgh *PageHandler) NewSubscription(c *gin.Context) {
	c.HTML(http.StatusOK, "new_subscription.html", nil)
}

func (pgh *PageHandler) GetAllSubscriptions(c *gin.Context) {
	c.HTML(http.StatusOK, "all_subscriptions.html", nil)
}

func (pgh *PageHandler) NewIncome(c *gin.Context) {
	c.HTML(http.StatusOK, "new_income.html", nil)
}

func (pgh *PageHandler) GetAllIncome(c *gin.Context) {
	c.HTML(http.StatusOK, "all_income.html", nil)
}

func (pgh *PageHandler) NewExpense(c *gin.Context) {
	c.HTML(http.StatusOK, "new_expense.html", nil)
}