package entity

import "time"

type ProductDetails struct {
	ID           int       `gorm:"column:id;primaryKey;autoIncrement"`
	ItemPID      string    `gorm:"column:item_pid;not null;"`
	ItemName     string    `gorm:"column:item_name"`
	ItemCategory string    `gorm:"column:item_category"`
	ItemPrice    string    `gorm:"column:item_price"`
	ItemQuantity string    `gorm:"column:item_quantity"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}
