package models

import (
	"time"
)

type Guitar struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Model            string    `gorm:"size:100;not null" json:"model"`
	Brand            string    `gorm:"size:100" json:"brand"`
	Category         string    `gorm:"size:50" json:"category"`
	Price            float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	StringCount      int       `gorm:"default:6" json:"string_count"`
	Color            string    `gorm:"size:50" json:"color"`
	SerialNumber     string    `gorm:"size:100;uniqueIndex" json:"serial_number"`
	InStock          bool      `gorm:"default:true" json:"in_stock"`
	StockQuantity    int       `gorm:"default:0" json:"stock_quantity"`
	YearManufactured int       `json:"year_manufactured"`
	Description      string    `gorm:"type:text" json:"description"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
