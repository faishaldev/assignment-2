package controllers

import (
	. "assignment-2/models"
	"assignment-2/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostOrder(ctx *gin.Context) {
	var order Order

	if err := ctx.ShouldBind(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	newOrder := Order{CustomerName: order.CustomerName, Items: order.Items}

	services.CreateOrder(&newOrder)

	ctx.JSON(http.StatusOK, gin.H{"data": newOrder})
}

func GetOrders(ctx *gin.Context) {
	var db = GetDb()
	var orders []Order

	db.Model(&Order{}).Preload("Items").Find(&orders)

	ctx.JSON(http.StatusOK, gin.H{"data": orders})
}

func PutOrder(ctx *gin.Context) {
	var db = GetDb()
	var order Order

	if err := db.Where("order_id = ?", ctx.Param("orderId")).First(&order).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})

		return
	}

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	updatedOrder := db.Model(&order).Preload("Items").Updates(&order)

	ctx.JSON(http.StatusOK, gin.H{"data": updatedOrder})
}

func DeleteOrder(ctx *gin.Context) {
	var db = GetDb()
	var order Order

	if err := db.Where("order_id = ?", ctx.Param("orderId")).First(&order).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})

		return
	}

	services.DeleteOrder(&order)

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
