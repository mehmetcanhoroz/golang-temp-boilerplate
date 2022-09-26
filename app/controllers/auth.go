package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mehmetcanhoroz/digital-marketplace/sdk/models"
	"github.com/mehmetcanhoroz/digital-marketplace/service"
)

type AuthController interface {
	Register(*gin.Context)
	Login(*gin.Context)
}

type authController struct {
	service service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return authController{
		service: authService,
	}
}

func (ctr authController) Register(c *gin.Context) {
	var input models.AuthRegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.RestError{
			Message: "Register payload is invalid!",
		})
		return
	}

	a, _ := ctr.service.Register(input)

	c.JSON(http.StatusOK,
		models.RestResponse{
			Result: a,
		},
	)
	return
}

func (ctr authController) Login(c *gin.Context) {
	var input models.AuthLoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.RestError{
			Message: "Login payload is invalid!",
		})
		return
	}

	a, err := ctr.service.Login(input)

	if err != nil {
		c.JSON(http.StatusForbidden,
			models.RestError{
				Message: "Credentials are invalid!",
			},
		)
		return
	}

	c.JSON(http.StatusOK,
		models.RestResponse{
			Result: a,
		},
	)
	return
}
