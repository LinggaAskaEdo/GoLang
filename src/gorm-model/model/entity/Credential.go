package model

// Credential struct
type Credential struct {
	ID       uint64   `gorm:"primary_key"`
	Email    string `gorm:"primary_key"`
	Password string `gorm:"NOT NULL"`
}
