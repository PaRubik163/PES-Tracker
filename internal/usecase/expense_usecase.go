package usecase

import (
	"fmt"
	"time"
	"tracker/internal/entity"
	"tracker/internal/logger"
	"tracker/internal/repository"
	pb "tracker/pkg/logger"
)

type ExpenseUseCase struct{
	expenseRepository *repository.ExpenseRepository
	loggerAdapter *logger.Adapter
}

func NewExpenseUseCase(expenseRepo *repository.ExpenseRepository, loggerAdapter *logger.Adapter) *ExpenseUseCase {
	return &ExpenseUseCase{
		expenseRepository: expenseRepo,
		loggerAdapter: loggerAdapter,
	}
} 

func (expUseCase *ExpenseUseCase) AddExpense(expense *entity.Expense) error {
	start := time.Now()

	if err := expUseCase.expenseRepository.Create(expense); err != nil{
		expUseCase.loggerAdapter.Expense("ERROR", &pb.LogData{
			UserId: fmt.Sprintf("%d", expense.UserID),
			Token: "",
			Method: "POST",
			Path: "/new_expense",
			Status: "500 INTERNAL",
			Time: start.Format(time.RFC3339),
			Latency: fmt.Sprintf("%v", time.Since(start)),
			Error: err.Error(),
		})

		return err
	}

	expUseCase.loggerAdapter.Expense("INFO", &pb.LogData{
		UserId: fmt.Sprintf("%d", expense.UserID),
		Token: "",
		Method: "POST",
		Path: "/new_expense",
		Status: "201 CREATED",
		Time: start.Format(time.RFC3339),
		Latency: fmt.Sprintf("%v", time.Since(start)),
		Error: "",
	})
	return nil
}

func (expUseCase *ExpenseUseCase) GetAllExpenses(userID int) ([]entity.Expense,error) {
	start := time.Now()

	expenses, err := expUseCase.expenseRepository.GetAll(userID)

	if err != nil{
		expUseCase.loggerAdapter.Expense("ERROR", &pb.LogData{
			UserId: fmt.Sprintf("%d", userID),
			Token: "",
			Method: "GET",
			Path: "/expenses",
			Status: "500 INTERNAL",
			Time: start.Format(time.RFC3339),
			Latency: fmt.Sprintf("%v", time.Since(start)),
			Error: err.Error(),
		})

		return nil, err
	}

	expUseCase.loggerAdapter.Expense("INFO", &pb.LogData{
		UserId: fmt.Sprintf("%d", userID),
		Token: "",
		Method: "GET",
		Path: "/expenses",
		Status: "200 OK",
		Time: start.Format(time.RFC3339),
		Latency: fmt.Sprintf("%v", time.Since(start)),
		Error: "",
	})
	return expenses, nil
}

func (expUseCase *ExpenseUseCase) DeleteExpense(expenseID, userID int) error {
	start := time.Now()

	if err := expUseCase.expenseRepository.DeleteByID(expenseID, userID); err != nil{
		expUseCase.loggerAdapter.Expense("ERROR", &pb.LogData{
			UserId: fmt.Sprintf("%d", userID),
			Token: "",
			Method: "DELETE",
			Path: fmt.Sprintf("/expense/%d", expenseID),
			Status: "500 INTERNAL",
			Time: start.Format(time.RFC3339),
			Latency: fmt.Sprintf("%v", time.Since(start)),
			Error: err.Error(),
		})

		return err
	}
	
	expUseCase.loggerAdapter.Expense("INFO", &pb.LogData{
		UserId: fmt.Sprintf("%d", userID),
		Token: "",
		Method: "DELETE",
		Path: fmt.Sprintf("/expense/%d", expenseID),
		Status: "200 OK",
		Time: start.Format(time.RFC3339),
		Latency: fmt.Sprintf("%v", time.Since(start)),
		Error: "",
	})

	return nil
}

func (expUseCase *ExpenseUseCase) GetExpensesByCategory(userID int) ([]entity.CategorySum, error) {
	start := time.Now()

	expenses, err := expUseCase.expenseRepository.GetExpensesByCategory(userID)

	if err != nil{
		expUseCase.loggerAdapter.Expense("ERROR", &pb.LogData{
			UserId: fmt.Sprintf("%d", userID),
			Token: "",
			Method: "GET",
			Path: "/expenses/category",
			Status: "500 INTERNAL",
			Time: start.Format(time.RFC3339),
			Latency: fmt.Sprintf("%v", time.Since(start)),
			Error: err.Error(),
		})

		return nil, err
	}

	expUseCase.loggerAdapter.Expense("INFO", &pb.LogData{
		UserId: fmt.Sprintf("%d", userID),
		Token: "",
		Method: "GET",
		Path: "/expenses/category",
		Status: "200 OK",
		Time: start.Format(time.RFC3339),
		Latency: fmt.Sprintf("%v", time.Since(start)),
		Error: "",
	})

	return expenses, nil
}