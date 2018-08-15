package dao

import (
	"time"
)

type Session struct {
	SessionID uint `gorm:"primary_key"`
	CreatedAt time.Time
	ExpireAt  time.Time
	Token     string `gorm:"not null; unique_index"`
}
