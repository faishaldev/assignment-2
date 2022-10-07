package controllers

import (
	"assignment-2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostOrder(ctx *gin.Context) {
	var order models.Order

	if err := ctx.ShouldBind(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})

		return
	}

	newOrder := models.Order{OrderedAt: order.OrderedAt, CustomerName: order.CustomerName, Items: order.Items}

	models.GetDb().Create(&newOrder)

	ctx.JSON(http.StatusOK, gin.H{"data": newOrder})
}

func GetOrders(ctx *gin.Context) {
	var orders []models.Order

	models.GetDb().Model(&models.Order{}).Preload("Items").Find(&orders)

	ctx.JSON(http.StatusOK, gin.H{"data": orders})
}

// func PutOrder(ctx *gin.Context) {
// 	var order models.Order

// 	if err := models.GetDb().Where("id = ?", ctx.Param("id")).First(&order).Error; err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H("error": "Record not found"))

// 		return
// 	}

// 	var
// }
