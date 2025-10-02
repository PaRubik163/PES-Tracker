package usecase

import "tracker/internal/repository"

type IncomeUseCase struct{
	incomeRepo *repository.IncomeRepository
}

func NewIncomeUseCase(inRepo *repository.IncomeRepository) *IncomeUseCase {
	return &IncomeUseCase{
		incomeRepo: inRepo,
	}
}