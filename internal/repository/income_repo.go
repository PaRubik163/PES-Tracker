package repository

import (
	"tracker/internal/entity"

	"gorm.io/gorm"
)

type IncomeRepository struct {
	db *gorm.DB
}

func NewIncomeRepository(db *gorm.DB) (*IncomeRepository, error) {
	if err := db.AutoMigrate(&entity.Income{}); err != nil{
		return nil, err
	}
	return &IncomeRepository{
		db: db,
	}, nil
}

func (inRepo *IncomeRepository) Create(income *entity.Income) error{
	return inRepo.db.Create(income).Error
}