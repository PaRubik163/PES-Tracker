package dto

import "time"

type UserSession struct{
	ID uint				`json:"id"`
	Login string 		`json:"login"`
	Jwt string 			`json:"jwt"`
	RegisteredAt time.Time `json:"registered_at"`
	LastLogin time.Time    `json:"last_login"`
}