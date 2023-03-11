package pkg

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID          int    `gorm:"primaryKey"`
	Title       string `gorm:"column:Title"`
	Description string `gorm:"column:Description"`
	Cost        int    `gorm:"column:Cost"`
}
