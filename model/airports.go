package model

type Airport struct {
	ID         int    `json:"id"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	// https://help.landbot.io/article/w2uqkq5z12-calculate-distances-with-google-maps-api
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (a Airport) TableName() string {
	return "airports"
}

func ExampleAirport() Airport {
	return Airport{
		ID:         1,
		Identifier: "KSWI",
		Name:       "Sherman Municipal Airport",
		Latitude:   -96.58609771728516,
		Longitude:  33.62419891357422,
	}
}
