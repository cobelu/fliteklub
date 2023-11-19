package misc

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestDistanceTo(t *testing.T) {
	t.Run("Same", func(t *testing.T) {
		c1 := Coordinate{Latitude: 33.628249292336456, Longitude: -96.5855631117168}
		c2 := Coordinate{Latitude: 37.35511482814887, Longitude: -121.95725336355912}
		expect := 0
		actual := c1.DistanceTo(c2, mile)
		assert.Equal(t, expect, actual, "The two words should be the same.")
	})

	t.Run("Different", func(t *testing.T) {
		c1 := Coordinate{Latitude: 33.628249292336456, Longitude: -96.5855631117168}
		c2 := Coordinate{Latitude: 37.35511482814887, Longitude: -121.95725336355912}
		miDistance := c1.DistanceTo(c2, mile)
		kmDistance := c1.DistanceTo(c2, kilometer)
		nmDistance := c1.DistanceTo(c2, nauticalMile)
		fmt.Printf("miDistance: %f mi\n", miDistance)
		fmt.Printf("kmDistance: %f km\n", kmDistance)
		fmt.Printf("nmDistance: %f nm\n", nmDistance)
	})
}
