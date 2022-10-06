package services

import (
	. "assignment-2/models"
)

func CreateOrder(newOrder *Order) error {
	db := GetDb()

	err := db.Create(newOrder).Error

	return err
}
