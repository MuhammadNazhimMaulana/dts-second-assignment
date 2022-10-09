package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ItemID      uint   `sql:"type:VARCHAR(15)" gorm:"primaryKey;column:item_id"`
	ItemCode    string `sql:"type:VARCHAR(15)"`
	Description string `sql:"type:VARCHAR(100)" gorm:"column:description"`
	Quantity    int
	Price       int64 `gorm:"column:harga"`
	OrderItem   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
