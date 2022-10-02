package app

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mehmetcanhoroz/digital-marketplace/app/controllers"
	"github.com/mehmetcanhoroz/digital-marketplace/db"
	"github.com/mehmetcanhoroz/digital-marketplace/middlewares"
	"github.com/mehmetcanhoroz/digital-marketplace/repository"
	"github.com/mehmetcanhoroz/digital-marketplace/service"
)

type Application struct {
}

func closeDBConnection(databaseI *sql.DB) {

	err := databaseI.Close()
	if err != nil {
		panic(err)
	}

}

func (a Application) Start() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	database := db.NewDatabase()
	dbCloser, err := database.DB()
	if err != nil {
		panic(err)
	}
	defer closeDBConnection(dbCloser)

	r := gin.Default()
	r.Use(gin.Logger())

	publicRoute := r.Group("/api/v1")

	//Repositories
	itemRepository := repository.NewItemRepository(database)
	categoryRepository := repository.NewCategoryRepository(database)
	userRepository := repository.NewUserRepository(database)

	//Services
	itemService := service.NewItemService(itemRepository)
	categoryService := service.NewCategoryService(categoryRepository)
	_ = service.NewUserService(userRepository)
	authService := service.NewAuthService(userRepository)
	middlewareService := middlewares.NewMiddlewareService(itemService, authService, categoryService)

	//Controllers
	itemController := controllers.NewItemController(itemService)
	itemRoute := publicRoute.Group("/items")
	itemRoute.GET("", itemController.GetAllItems)
	itemRoute.GET("/user/:username", itemController.GetAllItems)

	categoryController := controllers.NewCategoryController(categoryService)
	categoryRoute := publicRoute.Group("/categories")
	categoryRoute.GET("", categoryController.GetAllCategories)
	categoryRoute.GET(":id/items", categoryController.GetCategoryWithItems)

	authController := controllers.NewAuthController(authService)
	authRoute := publicRoute.Group("/auth")
	authRoute.POST("/register", authController.Register)
	authRoute.POST("/login", authController.Login)

	accountRoute := publicRoute.Group("/account")
	accountRoute.Use(middlewareService.ShouldBeAuthenticated()).GET("/me", authController.GetLoggedInUser)

	err = r.Run(":8000")
	if err != nil {
		panic(err)
	}

}
