package mappers

import "github.com/mehmetcanhoroz/digital-marketplace/sdk/models"

func MapUserModelToGetUserSuccessfulResponse(user models.User) models.GetUserSuccessfulResponse {
	return models.GetUserSuccessfulResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Phone:    user.Phone,
		Email:    user.Email,
		Username: user.Username,
		Birthday: user.Birthday,
	}
}
