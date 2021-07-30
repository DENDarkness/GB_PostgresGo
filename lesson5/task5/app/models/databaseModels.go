package models

import (
	"time"

	"gorm.io/gorm"
)

type Nodes struct {
	ID        uint   `gorm:"primaryKey"`
	Hostname  string `gorm:"unique"`
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
