package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	OrderID      uint   `sql:"unique;type:VARCHAR(68);not null" gorm:"primaryKey;column:order_id"`
	CustomerName string `sql:"type:VARCHAR(68);not null" gorm:"column:customer_name"`
	OrderAt      int64
	DetailItem   []Item `gorm:"foreignKey:OrderItem"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
