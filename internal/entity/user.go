package entity

import (
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct{
	ID uint					`gorm:"column:id"`
	Login string			`gorm:"unique;column:login"`
	Password string			`gorm:"column:password"`
	RegisteredAt time.Time  `gorm:"notNull;column:registered_at"`
	LastLogin time.Time 	`gorm:"notNull;column:last_login"`
}

func (u *User) CheckLoginAndPassword(login, password string) error {
	if (login == "" || password == ""){
		return errors.New("invalid login or password")
	}
	
	if (!strings.Contains(login, "@mail.ru") && !strings.Contains(login, "@gmail.com")){
		return errors.New("invalid login")
	}

	if (len(password) < 8){ 
		return errors.New("password need more then 8 letters")
	}

	return nil
}

func (u *User) HashPassword() error {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil{
		return err
	}

	u.Password = string(hashPass)
	return nil
}

func (u *User) CheckPassword(inputPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(inputPass))

	return err == nil
}