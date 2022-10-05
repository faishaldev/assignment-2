package models

import "time"

type Order struct {
	ID           string    `gorm:"primaryKey;type:VARCHAR(50);column:order_id"`
	CustomerName string    `json:"customerName" gorm:"not null;type:VARCHAR(50)"`
	OrderedAt    time.Time `json:"orderedAt" gorm:"not null;type:TEXT"`
	Items        []Item    `json:"items" gorm:"many2many:order_items"`
}
