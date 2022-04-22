package simulation

/*
	City simulates a city in world.
*/
type City struct {
	Name      string
	Direction string
}

func NewCity(cityName string) *City {
	return &City{Name: cityName}
}

func NewCityWithDirections(cityName, direction string) *City {
	return &City{Name: cityName, Direction: direction}
}
