package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mehmetcanhoroz/digital-marketplace/sdk/models"
	"github.com/mehmetcanhoroz/digital-marketplace/service"
)

type ItemController interface {
	GetAllItems(*gin.Context)
}

type itemController struct {
	service service.ItemService
}

func NewItemController(itemService service.ItemService) ItemController {
	return itemController{
		service: itemService,
	}
}

func (ctr itemController) GetAllItems(c *gin.Context) {
	result, _ := ctr.service.AllItems()

	c.JSON(http.StatusOK,
		models.RestResponse{
			Result: result,
		},
	)
	return
}
