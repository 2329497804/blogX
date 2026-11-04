package models

import (
	"time"
)

type Model struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
