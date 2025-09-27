package usecase

import (
	"tracker/internal/entity"
	"tracker/internal/repository"
)

type SubscriptionUseCase struct {
	subscriptionRepo *repository.SubscriptionRepo
}

func NewSubscriptionUseCase(sr *repository.SubscriptionRepo) *SubscriptionUseCase {
	return &SubscriptionUseCase{
		subscriptionRepo: sr,
	}
}

func (su *SubscriptionUseCase) CreateSubscription(sub *entity.Subscription) error {
	err := sub.CheckNewSubscription()

	if err != nil{
		return err
	}

	err = su.subscriptionRepo.Create(sub)

	if err != nil{
		return err
	}

	return nil
}