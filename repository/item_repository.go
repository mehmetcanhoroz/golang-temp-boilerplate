package repository

import (
	"github.com/mehmetcanhoroz/digital-marketplace/sdk/apperrors"
	"github.com/mehmetcanhoroz/digital-marketplace/sdk/models"
	"gorm.io/gorm"
)

type ItemRepository interface {
	FetchAllItems() ([]models.Item, *apperrors.AppError)
}

type itemRepository struct {
	database *gorm.DB
}

func (repository itemRepository) FetchAllItems() ([]models.Item, *apperrors.AppError) {
	var items []models.Item

	repository.database.Find(&items)

	return items, nil
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return itemRepository{
		database: db,
	}
}
