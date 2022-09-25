package misc

import (
	"gorm.io/gorm"
	"math"
)
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

type Coordinate struct {
	latitude  float64
	longitude float64
}

type DistanceUnit int

const (
	kilometer DistanceUnit = iota
	mile
	nauticalMile
)

// distanceTo calculates the distance from one Coordinat to another.
// It uses the algorithm found here: https://www.geodatasource.com/developers/go
func (c Coordinate) distanceTo(coord Coordinate, unit DistanceUnit) float64 {
	radlat1 := math.Pi * c.latitude / 180
	radlat2 := math.Pi * coord.latitude / 180

	theta := c.longitude - coord.longitude
	radtheta := math.Pi * theta / 180

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515 // To radians

	switch unit {
	case kilometer:
		dist = dist * 1.609344
	case nauticalMile:
		dist = dist * 0.8684
	}

	return dist
}
