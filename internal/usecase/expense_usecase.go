package usecase

import (
	"tracker/internal/entity"
	"tracker/internal/repository"
)

type ExpenseUseCase struct{
	expenseRepository *repository.ExpenseRepository
}

func NewExpenseUseCase(expenseRepo *repository.ExpenseRepository) *ExpenseUseCase {
	return &ExpenseUseCase{
		expenseRepository: expenseRepo,
	}
} 

func (expUseCase *ExpenseUseCase) AddExpense(expense *entity.Expense) error {
	if err := expUseCase.expenseRepository.Create(expense); err != nil{
		return err
	}

	return nil
}

func (expUseCase *ExpenseUseCase) GetAllExpenses(userID int) ([]entity.Expense,error) {
	expenses, err := expUseCase.expenseRepository.GetAll(userID)

	if err != nil{
		return nil, err
	}

	return expenses, nil
}

func (expUseCase *ExpenseUseCase) DeleteExpense(expenseID, userID int) error {
	if err := expUseCase.expenseRepository.DeleteByID(expenseID, userID); err != nil{
		return err
	}

	return nil
}

func (expUseCase *ExpenseUseCase) GetExpensesByCategory(userID int) ([]entity.CategorySum, error) {
	return expUseCase.expenseRepository.GetExpensesByCategory(userID)
}