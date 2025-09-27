package usecase

import "tracker/internal/repository"

type SubscriptionUseCase struct {
	subscriptionRepo *repository.SubscriptionRepo
}

func NewSubscriptionUseCase(sr *repository.SubscriptionRepo) *SubscriptionUseCase {
	return &SubscriptionUseCase{
		subscriptionRepo: sr,
	}
}
