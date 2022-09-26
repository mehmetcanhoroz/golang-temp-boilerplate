package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mehmetcanhoroz/digital-marketplace/sdk/models"
	"github.com/mehmetcanhoroz/digital-marketplace/service"
)

type CategoryController interface {
	GetAllCategories(*gin.Context)
	GetCategoryWithItems(*gin.Context)
}

type categoryController struct {
	service service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return categoryController{
		service: categoryService,
	}
}

func (ctr categoryController) GetAllCategories(c *gin.Context) {
	result, _ := ctr.service.AllCategories()

	c.JSON(http.StatusOK,
		models.RestResponse{
			Result: result,
		},
	)
	return
}

func (ctr categoryController) GetCategoryWithItems(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.RestError{
			Message: "ID is invalid!",
		})
		return
	}

	result, _ := ctr.service.GetCategoryWithItems(id)

	c.JSON(http.StatusOK,
		models.RestResponse{
			Result: result,
		},
	)
	return
}
