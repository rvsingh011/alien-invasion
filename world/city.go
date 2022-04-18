package world

type City struct {
	Name              string
	Direction         string
	CurrentAlienIndex []string
}

func NewCity(cityName string) *City {
	return &City{Name: cityName}
}

func NewCityWithDirections(cityName, direction string) *City {
	return &City{Name: cityName, Direction: direction}
}
