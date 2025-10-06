package entity

import (
	"time"
	"errors"
	"github.com/shopspring/decimal"
)

type Expense struct{
	ID int 				`gorm:"primaryKey;column:id" json:"id"`
	Title string 		`gorm:"notNull;column:title" json:"title"`
	Decription string   `gorm:"column:description" json:"description"`
	Category string		`gorm:"notNull;column:category" json:"category"`
	Amount decimal.Decimal `gorm:"type:numeric(12,2);notNull;column:amount" json:"amount"`
	Date time.Time	    `gorm:"type:date;notNull;column:expense_date" json:"expense_date"`
	UserID	int		    `gorm:"notNull;column:user_id" json:"user_id`	
	User 	User        `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
} 

type CategorySum struct{
	Category string       `json:"category"`
	Total decimal.Decimal `json:"total"`
}

func (exp *Expense) Validate() error {
	if exp.Title == ""{
		return errors.New("name cannot be empty")
	}

	if exp.Category == ""{
		return errors.New("category cannot be empty")
	}

	if exp.Amount.Cmp(decimal.Zero) <= 0{
		return errors.New("amount must be greater than 0")
	}
	
	if exp.Date.IsZero() {
		return errors.New("date cannot be empty")
	}

	if exp.Date.After(time.Now()) {
		return errors.New("date cannot be in the future")
	}

	return nil
}