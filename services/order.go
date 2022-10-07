package services

import (
	"assignment-2/models"
)

func CreateOrder(newOrder *models.Order) error {
	db := models.GetDb()

	err := db.Create(newOrder).Error

	return err
}

func UpdateOrder(payload *models.Order) error {
	var db = models.GetDb()
	var order models.Order

	order.CustomerName = payload.CustomerName
	order.Items = payload.Items

	err := db.Model(&order).Updates(order).Error

	return err
}
