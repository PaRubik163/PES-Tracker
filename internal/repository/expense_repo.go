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