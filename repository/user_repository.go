package repository

import (
	"github.com/mehmetcanhoroz/digital-marketplace/sdk/apperrors"
	"github.com/mehmetcanhoroz/digital-marketplace/sdk/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	SaveUser(models.User) (*models.User, *apperrors.AppError)
	FindByEmail(string) (*models.User, *apperrors.AppError)
	FindByID(uint64) (*models.User, *apperrors.AppError)
}

type userRepository struct {
	database *gorm.DB
}

func (repository userRepository) SaveUser(user models.User) (*models.User, *apperrors.AppError) {
	result := repository.database.Omit(clause.Associations).Create(&user)

	if result.Error != nil {
		return nil, apperrors.NewAppError("", result.Error)
	}

	return &user, nil
}

func (repository userRepository) FindByEmail(email string) (*models.User, *apperrors.AppError) {
	user := models.User{Email: email}
	result := repository.database.First(&user, "email = ?", user.Email)

	if result.Error != nil {
		return nil, apperrors.NewAppError("", result.Error)
	}

	return &user, nil
}

func (repository userRepository) FindByID(userID uint64) (*models.User, *apperrors.AppError) {
	user := models.User{ID: userID}
	result := repository.database.Find(&user)

	if result.Error != nil {
		return nil, apperrors.NewAppError("", result.Error)
	}
	if result.RowsAffected < 1 {
		return nil, apperrors.NewAppError("User could not be found!", nil)
	}

	return &user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		database: db,
	}
}
