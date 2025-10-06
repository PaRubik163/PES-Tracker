package repository

import (
	"time"
	"tracker/internal/entity"

	"github.com/shopspring/decimal"
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
func (ur *UserRepository) CountUsersSubscription(id int) (int64,error) {
	var count int64
	return count, ur.db.Table("subscriptions").	
    			 Where("subscriptions.user_id = ?", id).
    			 Count(&count).Error
}

//counting quantity users income
func (ur *UserRepository) CountUsersIncome(id int) (decimal.Decimal, error) {
	var amount decimal.Decimal
	return amount, ur.db.Table("incomes").
						Where("income_date >= date_trunc('month', CURRENT_DATE) AND income_date < (date_trunc('month', CURRENT_DATE) + interval '1 month') AND user_id = ?", id).
						Select("COALESCE(SUM(amount), 0)").
						Scan(&amount).Error
						
}

//counting quantity user expenses
func (ur *UserRepository) CountUserExpenses(id int) (decimal.Decimal, error) {
	var amount decimal.Decimal
	return amount, ur.db.Table("expenses").
						 Where("expense_date >= date_trunc('month', CURRENT_DATE) AND expense_date < (date_trunc('month', CURRENT_DATE) + interval '1 month') AND user_id = ?", id).
						 Select("COALESCE(SUM(amount), 0)").
						 Scan(&amount).Error
}