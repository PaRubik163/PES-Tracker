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

func (sr *SubscriptionRepo) GetAll(userID int) ([]*entity.Subscription, error) {
	var subs []*entity.Subscription

	err := sr.db.Where("user_id = ?", userID).Find(&subs).Error

	if err != nil{
		return nil, err
	}

	return subs, nil
}

func (sr *SubscriptionRepo) DeleteByID(subID, userID int) error {
	return sr.db.Where("id = ? AND user_id = ?", subID, userID).Delete(&entity.Subscription{}).Error
}