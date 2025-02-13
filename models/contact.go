package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name    string `gorm:"size:100;unique;not null" json:"name"`
	Email   string `gorm:"unique;not null" json:"email"`
	Phone   string `gorm:"unique;not null" json:"phone"`
	Address string `gorm:"size:255" json:"address"`
}
