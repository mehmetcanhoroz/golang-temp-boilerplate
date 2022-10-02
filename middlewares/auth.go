package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mehmetcanhoroz/digital-marketplace/sdk/models"
	"github.com/mehmetcanhoroz/digital-marketplace/service"
)

type MiddlewareService interface {
	ShouldBeAuthenticated() gin.HandlerFunc
}

type middlewareService struct {
	itemService     service.ItemService
	authService     service.AuthService
	categoryService service.CategoryService
}

func (s middlewareService) ShouldBeAuthenticated() gin.HandlerFunc {
	//func verifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return func(c *gin.Context) {
		//err := token.TokenValid(c)
		tokenString := s.authService.ExtractJWTToken(c)

		err := s.authService.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.RestError{
				Message: "Unauthorized!",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func NewMiddlewareService(itemService service.ItemService, authService service.AuthService, categoryService service.CategoryService) MiddlewareService {
	return middlewareService{
		itemService:     itemService,
		authService:     authService,
		categoryService: categoryService,
	}
}
