package dto

import "time"

type UserSession struct{
	ID int									`json:"id"`
	Login string 							`json:"login"`
	Token string							`json:"token"`
	SubscriptionsQuantity int64 			`json:"subscriptions_quantity"`
	ExpensesMonth float32					`json:"expenses_month"`
	IncomeMonth float32						`json:"income_month"`
	CreateSessionAt time.Time 				`json:"create_session_at"`
}