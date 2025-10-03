package repository

import (
	"time"
	"tracker/internal/entity"

	"gorm.io/gorm"
)

type UserRepository struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) (*UserRepository, error) {
	if err := db.AutoMigrate(&entity.User{}); err != nil{
		return nil, err
	}
	
	return &UserRepository{db: db}, nil
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

func (ur *UserRepository) UpdateLogin(login string) error {
	return ur.db.Model(entity.User{}).Where("login = ?", login).Update("last_login", time.Now()).Error
}

//counting quantity users subscriptions
func (ur *UserRepository) CountUsersSubscription(login string) (int64,error) {
	var count int64
	return count, ur.db.Table("subscriptions").
    			 Joins("JOIN users ON subscriptions.user_id = users.id").
    			 Where("users.login = ?", login).
    			 Count(&count).Error
}

//counting quantity users income
func (ur *UserRepository) CountUsersIncome(login string) (int64, error) {
	var amount int64
	return amount, ur.db.Table("incomes").
						Joins("JOIN users ON incomes.user_id = users.id").
						Where("users.login = ?", login).
						Select("COALESCE(SUM(incomes.amount), 0)").
						Scan(&amount).Error
						
}