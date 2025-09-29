package dto

import "time"

type UserSession struct{
	ID int				`json:"id"`
	Login string 		`json:"login"`
	Jwt string 			`json:"jwt"`
	CreateSessionAt time.Time `json:"create_session_at"`
}