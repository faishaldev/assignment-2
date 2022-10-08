package controllers

import (
	"assignment-2/models"
	database "assignment-2/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func PostOrder(ctx *gin.Context) {
	var db = database.GetDb()
	var order models.Order

	if err := ctx.ShouldBind(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	newOrder := models.Order{CustomerName: order.CustomerName, Items: order.Items}

	db.Create(&newOrder)

	ctx.JSON(http.StatusOK, gin.H{"data": newOrder})
}

func GetOrders(ctx *gin.Context) {
	var db = database.GetDb()
	var orders []models.Order

	db.Model(&models.Order{}).Preload("Items").Find(&orders)

	ctx.JSON(http.StatusOK, gin.H{"data": orders})
}

func PutOrder(ctx *gin.Context) {
	var db = database.GetDb()
	var order models.Order

	if err := db.Where("order_id = ?", ctx.Param("orderId")).First(&order).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})

		return
	}

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	db.Model(&order).Updates(models.Order{CustomerName: order.CustomerName})

	ctx.JSON(http.StatusOK, gin.H{"message": "update success"})
}

func DeleteOrder(ctx *gin.Context) {
	var db = database.GetDb()
	var order models.Order

	if err := db.Where("order_id = ?", ctx.Param("orderId")).First(&order).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})

		return
	}

	db.Select(clause.Associations).Delete(&order)

	ctx.JSON(http.StatusOK, gin.H{"message": "delete success"})
}
