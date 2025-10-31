package usecase

import (
	"fmt"
	"time"
	"tracker/internal/entity"
	"tracker/internal/logger"
	"tracker/internal/repository"
	pb "tracker/pkg/logger"
)

type IncomeUseCase struct{
	incomeRepo *repository.IncomeRepository
	loggerAdapter *logger.Adapter
}

func NewIncomeUseCase(inRepo *repository.IncomeRepository, loggerAdapter *logger.Adapter) *IncomeUseCase {
	return &IncomeUseCase{
		incomeRepo: inRepo,
		loggerAdapter: loggerAdapter,
	}
}

func (inUseCase *IncomeUseCase) AddIncome(income *entity.Income) error {
	start := time.Now()

	if err := inUseCase.incomeRepo.Create(income); err != nil{
		inUseCase.loggerAdapter.Income("ERROR", &pb.LogData{
			UserId: fmt.Sprintf("%d", income.UserID),
			Token: "",
			Method: "POST",
			Path: "/new_income",
			Status: logger.Internal,
			Time: start.Format(time.RFC3339),
			Latency: fmt.Sprintf("%v", time.Since(start)),
			Error: err.Error(),
		})
		return err
	}
	
	inUseCase.loggerAdapter.Income("INFO", &pb.LogData{
		UserId: fmt.Sprintf("%d", income.UserID),
		Token: "",
		Method: "POST",
		Path: "/new_income",
		Status: logger.Created,
		Time: start.Format(time.RFC3339),
		Latency: fmt.Sprintf("%v", time.Since(start)),
		Error: "",
	})

	return nil
}

func (inUseCase *IncomeUseCase) GetAllIncome(userID int) ([]entity.Income, error) {
	start := time.Now()

	income, err := inUseCase.incomeRepo.GetAll(userID)

	if err != nil{
		inUseCase.loggerAdapter.Income("ERROR", &pb.LogData{
			UserId: fmt.Sprintf("%d", userID),
			Token: "",
			Method: "GET",
			Path: "/income",
			Status: logger.Internal,
			Time: start.Format(time.RFC3339),
			Latency: fmt.Sprintf("%v", time.Since(start)),
			Error: err.Error(),
		})
		return nil, err
	}

	inUseCase.loggerAdapter.Income("INFO", &pb.LogData{
		UserId: fmt.Sprintf("%d", userID),
		Token: "",
		Method: "GET",
		Path: "/income",
		Status: logger.Ok,
		Time: start.Format(time.RFC3339),
		Latency: fmt.Sprintf("%v", time.Since(start)),
		Error: "",
	})
	return income, nil
}

func (inUseCase *IncomeUseCase) DeleteIncome(incomeID, userID int) error {
	start := time.Now()

	if err := inUseCase.incomeRepo.DeleteByID(incomeID, userID); err != nil{
		inUseCase.loggerAdapter.Income("ERROR", &pb.LogData{
			UserId: fmt.Sprintf("%d", userID),
			Token: "",
			Method: "DELETE",
			Path: fmt.Sprintf("/income/%d", incomeID),
			Status: logger.Internal,
			Time: start.Format(time.RFC3339),
			Latency: fmt.Sprintf("%v", time.Since(start)),
			Error: err.Error(),
		})
		return err
	}

	inUseCase.loggerAdapter.Income("INFO", &pb.LogData{
		UserId: fmt.Sprintf("%d", userID),
		Token: "",
		Method: "DELETE",
		Path: fmt.Sprintf("/income/%d", incomeID),
		Status: logger.Ok,
		Time: start.Format(time.RFC3339),
		Latency: fmt.Sprintf("%v", time.Since(start)),
		Error: "",
	})
	return nil
}