package handler

import (
	"net/http"
	"strconv"
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

func (inH *IncomeHandler) HandlerGetAll(c *gin.Context) {
	userID, ok := c.Get("user_id")

	if !ok{
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "user not authorized"})
		return 
	}

	income, err := inH.incomeUseCase.GetAllIncome(userID.(int))

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, income)
}

func (inH *IncomeHandler) DeleteIncome(c *gin.Context) {
	incomeIDStr := c.Param("id")

	incomeID, err := strconv.Atoi(incomeIDStr)
	
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : "unknown id"})
		return
	}

	userID, ok := c.Get("user_id")

	if !ok{
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "user not authorized"})
		return
	}

	if err := inH.incomeUseCase.DeleteIncome(incomeID, userID.(int)); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"message" : "income successful deleted"})
}