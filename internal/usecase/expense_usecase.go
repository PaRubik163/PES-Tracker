package usecase

import "tracker/internal/repository"

type ExpenseUseCase struct{
	expenseRepository *repository.ExpenseRepository
}

func NewExpenseUseCase(expenseRepo *repository.ExpenseRepository) *ExpenseUseCase {
	return &ExpenseUseCase{
		expenseRepository: expenseRepo,
	}
} 