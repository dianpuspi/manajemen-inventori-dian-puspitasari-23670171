package models

import "time"

type Item struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Stock     int
	Price     float64
	CreatedAt time.Time
}
