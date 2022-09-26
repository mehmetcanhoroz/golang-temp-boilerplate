package service

import (
	"github.com/mehmetcanhoroz/digital-marketplace/repository"
	"github.com/mehmetcanhoroz/digital-marketplace/sdk/apperrors"
)

type ItemService interface {
	AllItems() (interface{}, *apperrors.AppError)
}

type itemService struct {
	repository repository.ItemRepository
}

func NewItemService(repository repository.ItemRepository) ItemService {
	return itemService{
		repository: repository,
	}
}

func (s itemService) AllItems() (interface{}, *apperrors.AppError) {
	result, err := s.repository.FetchAllItems()
	return result, err
}
