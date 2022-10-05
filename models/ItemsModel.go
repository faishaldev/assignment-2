package models

type Item struct {
	ID          string `json:"lineItemId" gorm:"primaryKey;type:VARCHAR(50);column:item_id"`
	ItemCode    string `json:"itemCode" gorm:"not null;unique;type:VARCHAR(50)"`
	Description string `json:"description" gorm:"type:TEXT"`
	Quantity    int    `json:"quantity" gorm:"not null"`
	OrderId     string `json:"-" gorm:"not null;type:VARCHAR(50)"`
	Orders      Order  `json:"-" gorm:"foreignKey:OrderId;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
}
