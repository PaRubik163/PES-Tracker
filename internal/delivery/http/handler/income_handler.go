package handler

import (
	"net/http"
	"tracker/internal/entity"
	"tracker/internal/usecase"

	"github.com/gin-gonic/gin"
)

type IncomeHandler struct{
	incomeUseCase *usecase.IncomeUseCase
}

func NewIncomeHandler(inUseCase *usecase.IncomeUseCase) *IncomeHandler {
	return &IncomeHandler{
		incomeUseCase: inUseCase,
	}
}

func (inH *IncomeHandler) HandlerAddIncome(c *gin.Context) {
	userID, ok := c.Get("user_id")

	if !ok{
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "user not authorized"})
		return
	}

	income := &entity.Income{}

	if err := c.ShouldBindJSON(income); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	income.UserID = userID.(int)

	if err := inH.incomeUseCase.AddIncome(income); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message" : "income added successfully"})
}