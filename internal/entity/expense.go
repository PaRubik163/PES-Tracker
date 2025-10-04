package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Expense struct{
	ID int 				`gorm:"primaryKey;column:id" json:"id"`
	Title string 		`gorm:"notNull;column:title" json:"title"`
	Decription string   `gorm:"column:description" json:"description"`
	Category string		`gorm:"notNull;column:category" json:"category"`
	Amount decimal.Decimal `gorm:"type:numeric(12,2);notNull;column:amount" json:"amount"`
	Date time.Time	    `gorm:"type:date;notNull;column:expense_date; json:"expense_date"`
	UserID	int		    `gorm:"notNull;column:user_id" json:"user_id`	
	User 	User        `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
} 