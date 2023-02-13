package model

import (
	"time"

	"gorm.io/gorm"
)

type Credit struct {
	gorm.Model
	Uid           uint64 `gorm:"uniqueIndex"`
	No_of_credits int32
	Expiry        time.Time
}
