package models

type City struct {
	Name  string
	TempC float64
}

// NewYork creates a new city instance with the given name
func NewCity(n string) *City {
	return &City{
		Name: n,
	}
}

// TempF returns the city tempeture converted to Fareheit
func (c *City) TempF() float64 {
	return (c.TempC * 9.0 / 5.0) + 32.0
}
