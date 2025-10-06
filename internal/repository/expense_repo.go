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

func (expRepo *ExpenseRepository) GetExpensesByCategory(userID int) ([]*entity.CategorySum, error) {
	var result []*entity.CategorySum

	return result, expRepo.db.Table("expenses").
        Where("expense_date >= date_trunc('month', CURRENT_DATE) AND expense_date < (date_trunc('month', CURRENT_DATE) + interval '1 month') AND user_id = ?", userID).
        Select("category, COALESCE(SUM(amount::numeric), 0) AS TOTAL").
		Group("category").
        Find(&result).Error
}