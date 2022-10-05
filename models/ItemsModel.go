package models

type Item struct {
	ID          string `json:"lineItemId" gorm:"primaryKey;type:VARCHAR(50);column:item_id"`
	ItemCode    string `json:"itemCode" gorm:"not null;unique;type:VARCHAR(50)"`
	Description string `json:"description" gorm:"type:TEXT"`
	Quantity    int    `json:"quantity" gorm:"not null"`
	OrderId     string `gorm:"not null;unique;type:VARCHAR(50)"`
	Orders      Order  `json:"-" gorm:"foreignKey:ID;constraint:onUpdate:CASCADE,onDelete:CASCADE;many2many:order_items"`
}
