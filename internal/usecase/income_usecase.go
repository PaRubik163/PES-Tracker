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