package orm


import (
	"time"

	"gorm.io/gorm"

)

type Booking struct {
	gorm.Model
	UserID string
	CarID  string
	Start  time.Time
	End    time.Time
}
