package handler

import "tracker/internal/usecase"

type ExpenseHandler struct{
	expenseUseCase *usecase.ExpenseUseCase
}

func NewExpenseHandler(expenseUseCase *usecase.ExpenseUseCase) *ExpenseHandler {
	return &ExpenseHandler{
		expenseUseCase: expenseUseCase,
	}
}