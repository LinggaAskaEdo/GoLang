package model

import "github.com/jinzhu/gorm"

// User struct
type User struct {
	gorm.Model
	Name      string  `gorm:"NOT NULL"`
	CompanyID int     `gorm:"NOT NULL"`
	Company   Company `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
