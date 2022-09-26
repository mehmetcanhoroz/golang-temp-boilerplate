package service

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mehmetcanhoroz/digital-marketplace/repository"
	"github.com/mehmetcanhoroz/digital-marketplace/sdk/apperrors"
	"github.com/mehmetcanhoroz/digital-marketplace/sdk/models"
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 13 //Default cost 10: bcrypt.DefaultCost

type AuthService interface {
	Register(models.AuthRegisterRequest) (*models.AuthSuccessfulResponse, *apperrors.AppError)
	Login(request models.AuthLoginRequest) (*models.AuthSuccessfulResponse, *apperrors.AppError)
	VerifyToken(tokenString string) *apperrors.AppError
}

type authService struct {
	repository repository.UserRepository
}

func NewAuthService(repository repository.UserRepository) AuthService {
	return authService{
		repository: repository,
	}
}

func (s authService) Register(registerRequest models.AuthRegisterRequest) (*models.AuthSuccessfulResponse, *apperrors.AppError) {
	// validate email etc

	user := &models.User{
		FullName: "Test",
		Email:    registerRequest.Email,
		Phone:    registerRequest.Phone,
	}

	hashedPassword, err := s.HashPassword(registerRequest.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	user, err = s.repository.SaveUser(*user)

	if err != nil {
		return nil, err
	}

	tokenStr, err := s.GenerateToken(*user)
	if err != nil {
		return nil, err
	}

	token := models.AuthSuccessfulResponse{Token: tokenStr}

	return &token, nil
}

func (s authService) Login(loginRequest models.AuthLoginRequest) (*models.AuthSuccessfulResponse, *apperrors.AppError) {
	// validate email etc

	user := &models.User{
		Email: loginRequest.Email,
	}

	user, err := s.repository.FindByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	err = s.VerifyPassword(loginRequest.Password, user.Password)
	if err != nil {
		return nil, apperrors.NewAppError("Wrong password", err.Error)
	}

	tokenStr, err := s.GenerateToken(*user)
	if err != nil {
		return nil, err
	}

	token := models.AuthSuccessfulResponse{Token: tokenStr}

	return &token, nil
}

func (s authService) HashPassword(password string) (string, *apperrors.AppError) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost) //Default cost 10: bcrypt.DefaultCost
	if err != nil {
		return "", apperrors.NewAppError("", err)
	}

	return string(hashedPassword), nil
}

func (s authService) VerifyPassword(password, hashedPassword string) *apperrors.AppError {
	return apperrors.NewAppError("", bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)))
}

func (s authService) GenerateToken(user models.User) (string, *apperrors.AppError) {

	tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		return "", apperrors.NewAppError("", err)
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["u"] = user.ID
	claims["m"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	result, err := token.SignedString([]byte(os.Getenv("API_SECRET")))
	return result, apperrors.NewAppError("", err)
}

func (s authService) VerifyToken(tokenString string) *apperrors.AppError {
	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("API_SECRET")), nil
		})
	if err != nil {
		return apperrors.NewAppError("", err)
	}
	if !token.Valid {
		return apperrors.NewAppError("Token is invalid.", err)
	}
	return nil
}
