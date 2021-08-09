package model

// Company struct
type Company struct {
	ID   int    `gorm:"primary_key"`
	Name string `gorm:"NOT NULL"`
}
