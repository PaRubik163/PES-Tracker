package handler

import "tracker/internal/usecase"

type IncomeHandler struct{
	incomeUseCase *usecase.IncomeUseCase
}

func NewIncomeHandler(inUseCase *usecase.IncomeUseCase) *IncomeHandler {
	return &IncomeHandler{
		incomeUseCase: inUseCase,
	}
}