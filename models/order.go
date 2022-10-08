package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderID      string `sql:"unique;type:VARCHAR(68);not null" gorm:"column:order_id"`
	CustomerName string `sql:"type:VARCHAR(68);not null" gorm:"column:customer_name"`
	OrderAt      int64
	DetailItem   []Item `gorm:"foreignKey:OrderItem"`
}
