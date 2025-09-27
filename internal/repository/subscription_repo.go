package repository

import (
	"tracker/internal/entity"

	"gorm.io/gorm"
)

type SubscriptionRepo struct {
	db *gorm.DB
}

func NewSubscriptionRepo(db *gorm.DB) (*SubscriptionRepo) {
	db.AutoMigrate(&entity.Subscription{})
	return &SubscriptionRepo{db: db}
}

func (sr *SubscriptionRepo) Create(subscription *entity.Subscription) (error) {
	return sr.db.Create(&subscription).Error
}

//later I will do this using an array
func (sr *SubscriptionRepo) GetAll() (*entity.Subscription, error) {
	sub := &entity.Subscription{}

	err := sr.db.Find(sub).Error

	if err != nil{
		return nil, err
	}

	return sub, nil
}