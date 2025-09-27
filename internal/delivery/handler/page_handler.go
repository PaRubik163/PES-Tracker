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
	c.HTML(http.StatusOK, "user_info.html", nil)
}