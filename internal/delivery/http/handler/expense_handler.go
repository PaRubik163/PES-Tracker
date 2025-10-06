package handler

import (
	"net/http"
	"strconv"
	"tracker/internal/entity"
	"tracker/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ExpenseHandler struct{
	expenseUseCase *usecase.ExpenseUseCase
}

func NewExpenseHandler(expenseUseCase *usecase.ExpenseUseCase) *ExpenseHandler {
	return &ExpenseHandler{
		expenseUseCase: expenseUseCase,
	}
}

func (expH *ExpenseHandler) HandlerAdd(c *gin.Context) {
	userID, ok := c.Get("user_id")

	if !ok{
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "user not authorized"})
		return
	}

	expense := &entity.Expense{}

	if err := c.ShouldBindJSON(expense); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : "invalid request"})
		return
	}

	expense.UserID = userID.(int)

	if err := expH.expenseUseCase.AddExpense(expense); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return 
	}

	c.JSON(http.StatusCreated, gin.H{"message" : "expense successful created"})
}

func (expH *ExpenseHandler) HandlerGetAll(c *gin.Context) {
	userID, ok := c.Get("user_id")

	if !ok{
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "user not authorized"})
		return
	}

	expenses, err := expH.expenseUseCase.GetAllExpenses(userID.(int))

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, expenses)
}

func (expH *ExpenseHandler) HandlerDeleteExpense(c *gin.Context) {
	expenseIDStr := c.Param("id")

	expenseID, err := strconv.Atoi(expenseIDStr)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	userID, ok := c.Get("user_id")

	if !ok{
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "user not authorized"})
		return
	}

	if err := expH.expenseUseCase.DeleteExpense(expenseID, userID.(int)); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "expense successful deleted"})
}

func (expH *ExpenseHandler) HandlerGetExpensesByCategory(c *gin.Context) {
	userID, ok := c.Get("user_id")

	if !ok{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authorized"})
		return
	}

	result, err := expH.expenseUseCase.GetExpensesByCategory(userID.(int))

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}