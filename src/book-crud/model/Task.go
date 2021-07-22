package model

import "time"

// Task struct
type Task struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	CreatedAt   time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
	AssignedTo  string    `json:"assignedTo" gorm:"NOT NULL"`
	Description string    `json:"description" gorm:"NOT NULL"`
	Deadline    time.Time `json:"deadline" gorm:"NOT NULL"`
}
