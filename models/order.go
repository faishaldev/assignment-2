package models

import "time"

type Order struct {
	ID           uint      `gorm:"primaryKey;column:order_id" json:"-" validate:"required"`
	CustomerName string    `gorm:"not null;type:VARCHAR(50)" form:"customerName" json:"customerName" validate:"required" binding:"required"`
	OrderedAt    time.Time `gorm:"not null;autoCreateTime" form:"orderedAt" json:"orderedAt" validate:"required"`
	Items        []Item    `form:"items" json:"items" validate:"required" binding:"required"`
}
