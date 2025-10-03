package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type UserSession struct{
	ID int									`json:"id"`
	Login string 							`json:"login"`
	Token string							`json:"token"`
	SubscriptionsQuantity int64 			`json:"subscriptions_quantity"`
	ExpensesMonth decimal.Decimal			`json:"expenses_month"`
	IncomeMonth decimal.Decimal				`json:"income_month"`
	CreateSessionAt time.Time 				`json:"create_session_at"`
}