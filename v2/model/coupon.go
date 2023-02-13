package model

import (
	"time"

	"gorm.io/gorm"
)

type Coupon struct {
	gorm.Model
	Uid         uint64
	Code        string
	Expiry      time.Time
	Redeem_Time time.Time
	Redeemed    bool
}
