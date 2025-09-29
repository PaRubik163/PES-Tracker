package entity

import (
	"errors"
	"time"
)

type Subscription struct {
	ID              uint      `gorm:"primaryKey;column:id" json:"id"`
	Name            string    `gorm:"notNull;column:subscription_name" json:"name"`
	Amount          float32   `gorm:"column:amount" json:"amount"`
	Url             string    `gorm:"column:url" json:"url"`
	StartDate       time.Time `gorm:"notNull;column:start_date" json:"start_date"`
	BillingPeriod   string    `gorm:"notNull;column:billing_period" json:"billing_period"`
	NextBillingDate time.Time `gorm:"notNull;column:next_billing_date" json:"next_billing_date"`
}


func (s *Subscription) CheckNewSubscription() error {
	if s.Name == ""{
		return errors.New("name can't be empty")
	}
	
	if s.StartDate.IsZero(){
		return errors.New("start day can't be empty")
	}

	if s.BillingPeriod == ""{
		return errors.New("billing period can't be empty")
	}

	if s.NextBillingDate.IsZero(){
		return errors.New("next billing date can't be empty")
	}

	return nil
}