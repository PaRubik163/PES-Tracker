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

func (inRepo *IncomeRepository) GetAll(userID int) ([]*entity.Income, error) {
	var income []*entity.Income

	if err := inRepo.db.Where("user_id = ?", userID).Find(&income).Error; err != nil{
		return nil, err
	}

	return income, nil
}

func (inRepo *IncomeRepository) DeleteByID(incomeID, userID int) error {
	return inRepo.db.Where("id = ? AND user_id = ?", incomeID, userID).Delete(&entity.Income{}).Error
}