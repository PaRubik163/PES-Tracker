package repository

import (
	"time"
	"tracker/internal/entity"

	"gorm.io/gorm"
)

type UserRepository struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(user *entity.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) GetByLogin(login string) (*entity.User, error) {
	var user *entity.User

	if err := ur.db.Where("login = ?", login).First(&user).Error; err != nil{
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) UpdateLogin(login string) error{
	return ur.db.Model(entity.User{}).Where("login = ?", login).Update("last_login", time.Now()).Error
}