package repository

import (
	"github.com/mehmetcanhoroz/digital-marketplace/sdk/apperrors"
	"github.com/mehmetcanhoroz/digital-marketplace/sdk/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FetchAllParentCategories() ([]models.Category, *apperrors.AppError)
	GetCategoryWithItems(id uint64) (models.Category, *apperrors.AppError)
}

type categoryRepository struct {
	database *gorm.DB
}

func (repository categoryRepository) FetchAllParentCategories() ([]models.Category, *apperrors.AppError) {
	var categories []models.Category

	repository.database.Where("parent_category_id IS NULL").Find(&categories)

	return categories, nil
}

func (repository categoryRepository) GetCategoryWithItems(id uint64) (models.Category, *apperrors.AppError) {
	var category models.Category

	repository.database.Preload(models.CategoryModelItemsForeignObjectName).Find(&category, id)

	return category, nil
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return categoryRepository{
		database: db,
	}
}
