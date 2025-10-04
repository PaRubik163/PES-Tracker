package repository

import (
	"tracker/internal/entity"

	"gorm.io/gorm"
)

type ExpenseRepository struct{
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) (*ExpenseRepository, error) {
	if err := db.AutoMigrate(&entity.Expense{}); err != nil{
		return nil, err
	}
	
	return &ExpenseRepository{
		db: db,
	}, nil
}

func (expRepo *ExpenseRepository) Create(expense *entity.Expense) error {
	return expRepo.db.Create(&expense).Error
}

func (expRepo *ExpenseRepository) GetAll(userID int) ([]*entity.Expense, error) {
	var expenses []*entity.Expense

	return expenses, expRepo.db.Where("user_id = ?", userID).Find(&expenses).Error
}

func (expRepo *ExpenseRepository) DeleteByID(expenseID, userID int) error {
	return expRepo.db.Where("id = ? AND user_id=?", expenseID, userID).Delete(&entity.Expense{}).Error
}