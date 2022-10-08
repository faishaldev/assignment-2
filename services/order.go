package services

import (
	. "assignment-2/models"

	"gorm.io/gorm/clause"
)

func CreateOrder(newOrder *Order) error {
	db := GetDb()

	err := db.Create(newOrder).Error

	return err
}

func DeleteOrder(order *Order) error {
	var db = GetDb()

	err := db.Select(clause.Associations).Delete(&order).Error

	return err
}
