package service

import (
	"github.com/mehmetcanhoroz/digital-marketplace/repository"
	"github.com/mehmetcanhoroz/digital-marketplace/sdk/apperrors"
	"github.com/mehmetcanhoroz/digital-marketplace/sdk/models"
)

type CategoryService interface {
	AllParentCategories() (interface{}, *apperrors.AppError)
	GetCategoryWithItems(id uint64) (models.Category, *apperrors.AppError)
}

type categoryService struct {
	repository repository.CategoryRepository
}

func NewCategoryService(repository repository.CategoryRepository) CategoryService {
	return categoryService{
		repository: repository,
	}
}

func (s categoryService) AllParentCategories() (interface{}, *apperrors.AppError) {
	result, err := s.repository.FetchAllParentCategories()
	return result, err
}

func (s categoryService) GetCategoryWithItems(id uint64) (models.Category, *apperrors.AppError) {
	result, err := s.repository.GetCategoryWithItems(id)
	return result, err
}
