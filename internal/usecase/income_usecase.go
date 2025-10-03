package usecase

import (
	"tracker/internal/entity"
	"tracker/internal/repository"
)

type IncomeUseCase struct{
	incomeRepo *repository.IncomeRepository
}

func NewIncomeUseCase(inRepo *repository.IncomeRepository) *IncomeUseCase {
	return &IncomeUseCase{
		incomeRepo: inRepo,
	}
}

func (inUseCase *IncomeUseCase) AddIncome(income *entity.Income) error {
	if err := inUseCase.incomeRepo.Create(income); err != nil{
		return err
	}

	return nil
}

func (inUseCase *IncomeUseCase) GetAllIncome(userID int) ([]*entity.Income, error) {
	income, err := inUseCase.incomeRepo.GetAll(userID)

	if err != nil{
		return nil, err
	}

	return income, nil
}

func (inUseCase *IncomeUseCase) DeleteIncome(incomeID, userID int) error {
	if err := inUseCase.incomeRepo.DeleteByID(incomeID, userID); err != nil{
		return err
	}

	return nil
}