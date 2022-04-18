package world

type Isimulation interface {
	CreateWorld() error
	ViewWorld() string
	CreateAliens() error
	ViewAliens() error
	fight()
	burryDeadAliens(deadAliens []string)
	removeDestroyedCities(destoyedCities []string)
	deleteCityFromWorldMap(city string)
	prepareAttack()
	runNextRoundOfAttack()
	EndAndConclude() string
}
