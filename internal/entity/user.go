package entity

import (
	"errors"
	"strings"
	"time"
	"unicode"
	"golang.org/x/crypto/bcrypt"
)

type User struct{
	ID uint					`gorm:"primaryKey;column:id"`
	Login string			`gorm:"notNull;unique;column:login"`
	Password string			`gorm:"notNull;column:password"`
	RegisteredAt time.Time  `gorm:"notNull;column:registered_at"`
	LastLogin time.Time 	`gorm:"notNull;column:last_login"`
}

func (u *User) CheckLogin(login string) error {
	if (login == ""){
		return errors.New("invalid login")
	}
	
	if (!strings.Contains(login, "@mail.ru") && !strings.Contains(login, "@gmail.com") && !strings.Contains(login, "@icloud.com")){
		return errors.New("invalid login")
	}

	return nil
}

func (u *User) CheckPassword(pass string) error {
	if len(pass) < 8 {
		return errors.New("invalid password: must be at least 8 characters")
	}

	var (
		hasUpper  bool
		hasLower  bool
		hasDigit  bool
		hasSymbol bool
	)

	for _, v := range pass {
		switch {
		case unicode.IsUpper(v):
			hasUpper = true
		case unicode.IsLower(v):
			hasLower = true
		case unicode.IsDigit(v):
			hasDigit = true
		case strings.ContainsRune("!@#$%^&*()_+[]{}<>?/.,;:'\"|\\~`-", v):
			hasSymbol = true
		}
	}

	if !hasUpper {
		return errors.New("invalid password: must contain at least one uppercase letter")
	}
	if !hasLower {
		return errors.New("invalid password: must contain at least one lowercase letter")
	}
	if !hasDigit {
		return errors.New("invalid password: must contain at least one digit")
	}
	if !hasSymbol {
		return errors.New("invalid password: must contain at least one special character")
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

func (u *User) CheckHashedPassword(inputPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(inputPass))

	return err == nil
}