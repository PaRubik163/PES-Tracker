package entity

import (
	"errors"
	"time"
	"github.com/shopspring/decimal"
)

type Income struct{
	ID int			   `gorm:"primaryKey;column:id" json:"id"`
	Title string	   `gorm:"notNull;column:title" json:"title"`
	Description string  `gorm:"column:description" json:"description"`
	Amount decimal.Decimal	   `gorm:"type:numeric(12,2);notNull;column:amount" json:"amount"`
	Date time.Time	   `gorm:"type:date;notNull;column:income_date" json:"income_date"`
	UserID	int		   `gorm:"notNull;column:user_id" json:"user_id`	
	User 	User       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (in *Income) Validate() error {
	if in.Title == ""{
		return errors.New("name cannot be empty")
	}

	if in.Amount.Cmp(decimal.Zero) <= 0{
		return errors.New("amount must be greater than 0")
	}
	
	if in.Date.IsZero() {
		return errors.New("date cannot be empty")
	}

	if in.Date.After(time.Now()) {
		return errors.New("date cannot be in the future")
	}

	return nil
}