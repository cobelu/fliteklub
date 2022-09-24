package misc

import "gorm.io/gorm"
import "time"

type GormModel gorm.Model

type Date struct {
	year  uint8
	month time.Month
	day   uint8
}

func Now() time.Time {
	return time.Now().UTC()
}

func OneHourAgo() time.Time {
	return Now().Add(time.Duration(-1) * time.Hour)
}

func OneHourLater() time.Time {
	return Now().Add(time.Duration(1) * time.Hour)
}

func Today() Date {
	year, month, day := time.Now().Date()
	return Date{uint8(year), month, uint8(day)}
}
