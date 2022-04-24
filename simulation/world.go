package simulation

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/rvsingh011/alien-invasion/utils"
	"go.uber.org/zap"
)

/*
	Simulation hold the data for the simualtion.
*/
type Simulation struct {
	Iterations     int
	WorldFile      string
	NumberOfAliens int
	AlienNames     string

	// All cities detailed map for the planet, this includes all cities and its roads
	World map[string][]*City

	// List of all selected aliens who are selected for the mission "Death"
	Aliens []*Alien

	// List of all cities on the target planet, this does not include the roads
	Cities []*City

	// Alien Commandar record for deployed aliens, represent which alien is currently in which city
	AlienCityMapping map[string]string

	// united nations defense record, tracks and records which city is under attack by which alien
	CityAlienMapping map[string][]string

	// Record the attack vector for future generation or run simulations
	RandSeed *rand.Rand

	// communication messages for future generation to read and learn (not in use), but can be used for only listning a particular type of messages
	logger *zap.Logger
}

/*
	CreateWorld generate the world from the provided file
*/
func NewSimulation(iterations, alienNumbers int, alienNames, worldFile string, randomSeed *rand.Rand, logger *zap.Logger) (*Simulation, error) {
	simulation := Simulation{
		Iterations:       iterations,
		WorldFile:        worldFile,
		NumberOfAliens:   alienNumbers,
		AlienNames:       alienNames,
		World:            make(map[string][]*City),
		AlienCityMapping: make(map[string]string),
		CityAlienMapping: make(map[string][]string),
		RandSeed:         randomSeed,
		logger:           logger,
	}
	return &simulation, nil
}

/*
	CreateWorld simulates the world from the provided world file.
*/
func (sim *Simulation) CreateWorld() error {
	worldFile, err := os.Open(sim.WorldFile)
	if err != nil {
		return fmt.Errorf("Error Reading the world file : %s, Error: %s", sim.WorldFile, err.Error())
	}
	defer worldFile.Close()

	scanner := bufio.NewScanner(worldFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lineSplitOnSpace := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		newCity := lineSplitOnSpace[0]

		if _, ok := sim.World[newCity]; !ok {
			sim.World[newCity] = make([]*City, 0)
			sim.Cities = append(sim.Cities, NewCity(newCity))
		}

		lineSplitOnSpace = lineSplitOnSpace[1:]
		for _, cityDir := range lineSplitOnSpace {
			cityWithDirections := strings.Split(strings.TrimSpace(cityDir), "=")
			// if city already exists
			isCityExist := false
			for _, city := range sim.World[newCity] {
				if city.Name == cityWithDirections[1] {
					isCityExist = true
				}
			}

			if !isCityExist {
				sim.World[newCity] = append(
					sim.World[newCity],
					NewCityWithDirections(cityWithDirections[1], cityWithDirections[0]),
				)
				// if the income city is not there, Add it to the city.
				if _, ok := sim.World[cityWithDirections[1]]; !ok {
					sim.World[cityWithDirections[1]] = make([]*City, 0)
					sim.World[cityWithDirections[1]] = append(
						sim.World[cityWithDirections[1]],
						NewCityWithDirections(newCity, utils.GetOppositeDirection(cityWithDirections[0])),
					)
					sim.Cities = append(sim.Cities, NewCity(cityWithDirections[1]))
				}
			}

		}
	}
	return nil
}

/*
	ViewWorld print the current status of the world, print out the city layout in human readable format
*/
func (sim *Simulation) ViewWorld() string {

	println("=========================================")
	println("The World before attack is")
	println("=========================================")
	var world strings.Builder
	for key, value := range sim.World {
		fmt.Printf("The City %s is connected to below cities\n", key)
		for _, city := range value {
			fmt.Printf("\tThe City %s is %s to the %s\n", city.Name, city.Direction, key)
			world.WriteString(fmt.Sprintf("\tThe City %s is %s to the %s\n", city.Name, city.Direction, key))
		}
	}
	return world.String()
}

/*
	CreateAliens simulated aliens which will be attacking the cities.
*/
func (sim *Simulation) CreateAliens() error {
	alienNames, err := os.Open(sim.AlienNames)
	if err != nil {
		return fmt.Errorf("Error Reading the alien name file : %s, Error: %s", sim.AlienNames, err.Error())
	}
	defer alienNames.Close()
	scanner := bufio.NewScanner(alienNames)
	scanner.Split(bufio.ScanLines)

	for aliens := 0; aliens < sim.NumberOfAliens && scanner.Scan(); aliens++ {
		sim.Aliens = append(sim.Aliens, NewAlien(scanner.Text()))
	}
	return nil
}

/*
	ViewAliens prints all the aliens ready to run atatck the cities.
*/
func (sim *Simulation) ViewAliens() error {

	println("=========================================")
	println("Alien Profiles")
	println("=========================================")

	for idx, alien := range sim.Aliens {
		fmt.Printf("The alien %d has a name %s\n", idx, alien.Name)
	}
	return nil
}

/*
	Start starts the simulations.
	1. In first step each attack will randomly choose a city for attack.
	2. After step 1 alien can only move to the left links to different city.
	3. An alien can be trapped in a city in that case he stays in the same city.
	4. An alien can also decide not to move and stay in the same city.
*/
func (sim *Simulation) Start() error {
	for iteration := 1; iteration <= sim.Iterations; iteration++ {
		// should the next iteration run ?
		if sim.isNextIterationRequired() == false {
			break
		}
		println("=========================================")
		fmt.Printf("Running %d iteration of Attack\n", iteration)
		println("=========================================")

		// if aliens just arrrived they need to prepare weapons and initiate the attack
		if iteration == 1 {
			sim.prepareAttack()
		} else {
			sim.runNextRoundOfAttack()
		}
		sim.fight()
	}
	return nil
}

/*
	fight simualtes the fight between aliens which arrived in the same city.
	1. If more than one alien comes to same city, all aliens are destoyed with the city and its link.
*/
func (sim *Simulation) fight() {
	deadAliens := make([]string, 0)
	destoyedCities := make([]string, 0)
	for city, aliensInCity := range sim.CityAlienMapping {
		if len(aliensInCity) > 1 {
			var report strings.Builder
			report.WriteString(fmt.Sprintf("The %s was destroyed by ", city))
			destoyedCities = append(destoyedCities, city)
			for idx := range aliensInCity {
				report.WriteString(fmt.Sprintf("%s, ", aliensInCity[idx]))
				deadAliens = append(deadAliens, aliensInCity[idx])
			}
			fmt.Println(report.String())
		}
	}

	sim.burryDeadAliens(deadAliens)
	sim.removeDestroyedCities(destoyedCities)
}

/*
	burryDeadAliens simualtes the death of a alien.
*/
func (sim *Simulation) burryDeadAliens(deadAliens []string) {
	for _, deadAlien := range deadAliens {
		delete(sim.AlienCityMapping, deadAlien)

		for i := len(sim.Aliens) - 1; i >= 0; i-- {
			if sim.Aliens[i].Name == deadAlien {
				sim.Aliens = append(sim.Aliens[:i], sim.Aliens[i+1:]...)
			}
		}
	}
}

/*
	removeDestroyedCities simualtes a destroyed city.
*/
func (sim *Simulation) removeDestroyedCities(destoyedCities []string) {
	for _, destroyedCity := range destoyedCities {
		delete(sim.CityAlienMapping, destroyedCity)
		for i := len(sim.Cities) - 1; i >= 0; i-- {
			if sim.Cities[i].Name == destroyedCity {
				sim.Cities = append(sim.Cities[:i], sim.Cities[i+1:]...)
			}
		}
		// remove the destoryed city from the world map
		sim.deleteCityFromWorldMap(destroyedCity)
	}
}

/*
	deleteCityFromWorldMap simualtes a destroyed city removal from world map.
*/
func (sim *Simulation) deleteCityFromWorldMap(city string) {
	for _, eachLinkedCity := range sim.World[city] {
		for idx, eachLink := range sim.World[eachLinkedCity.Name] {
			if eachLink.Name == city {
				sim.World[eachLinkedCity.Name] = append(sim.World[eachLinkedCity.Name][:idx], sim.World[eachLinkedCity.Name][idx+1:]...)
			}
		}
	}
	delete(sim.World, city)
}

/*
	isNextIterationRequired checks if next iteraton of simulations is required.
	1. if all cities are destoyed, stop the simulation.
	2. if all aliens are dead, stop the simulations.
*/
func (sim *Simulation) isNextIterationRequired() bool {
	if len(sim.Aliens) < 1 || len(sim.Cities) < 1 {
		return false
	}
	return true
}

/*
	prepareAttack simulates the arrival of alien.
	1. In this step all aliens choose a city to attack.
*/
func (sim *Simulation) prepareAttack() {

	// intialize the city Command Center record
	for idx := range sim.Cities {
		sim.CityAlienMapping[sim.Cities[idx].Name] = make([]string, 0)
	}

	// all aliens will first choose a city of there choice to attack
	for _, alien := range sim.Aliens {
		cityIndex := utils.GetRandomNumber(0, len(sim.Cities)-1, sim.RandSeed)
		fmt.Printf("Alien %s choose %s city\n", alien.Name, sim.Cities[cityIndex].Name)
		sim.AlienCityMapping[alien.Name] = sim.Cities[cityIndex].Name

		// city command center intercepted target cities and who will be visiting
		sim.CityAlienMapping[sim.Cities[cityIndex].Name] = append(sim.CityAlienMapping[sim.Cities[cityIndex].Name], alien.Name)
	}
}

/*
	runNextRoundOfAttack simulates the attack of alien after the intial round.
	1. In this step all aliens either moves to a new city or stay in the same city, or get trapped in the city.
*/
func (sim *Simulation) runNextRoundOfAttack() {
	// Aliens will choose a city conneted to the exsiting city
	for idx := range sim.Aliens {
		alienCurrentCity := sim.AlienCityMapping[sim.Aliens[idx].Name]

		maxIndex := len(sim.World[alienCurrentCity])
		if maxIndex == 0 {
			fmt.Printf("The alien %s is trapped in the %s city\n", sim.Aliens[idx].Name, alienCurrentCity)
			continue
		}

		newCityIndex := utils.GetRandomNumber(0, maxIndex, sim.RandSeed)

		// maxIndex == len(sim.World[alienCurrentCity]) denoted no move by the alien.
		if newCityIndex == maxIndex {
			fmt.Printf("The alien %s decided to stay in %s city\n", sim.Aliens[idx].Name, alienCurrentCity)
			continue
		}

		// remove the alien from current city
		for index, alien := range sim.CityAlienMapping[alienCurrentCity] {
			if alien == sim.Aliens[idx].Name {
				sim.CityAlienMapping[alienCurrentCity] = append(sim.CityAlienMapping[alienCurrentCity][:index], sim.CityAlienMapping[alienCurrentCity][index+1:]...)
			}
		}

		fmt.Printf("The alien %s will now move to %s\n", sim.Aliens[idx].Name, sim.World[alienCurrentCity][newCityIndex].Name)
		sim.AlienCityMapping[sim.Aliens[idx].Name] = sim.World[alienCurrentCity][newCityIndex].Name
		sim.CityAlienMapping[sim.World[alienCurrentCity][newCityIndex].Name] = append(sim.CityAlienMapping[sim.World[alienCurrentCity][newCityIndex].Name], sim.Aliens[idx].Name)
	}
}

/*
	EndAndConclude ends the simulations and print what is left of the world.
*/
func (sim *Simulation) EndAndConclude() string {
	println("=========================================")
	println("The Bloody war ended, these are the remins of the world")
	println("=========================================")

	var leftWorld strings.Builder
	for city, linkedCities := range sim.World {
		// if len(linkedCities) == 0 {
		// 	continue
		// }
		var cityInfo strings.Builder
		cityInfo.WriteString(city)
		for _, eachCity := range linkedCities {
			cityInfo.WriteString(fmt.Sprintf(" %s=%s", eachCity.Direction, eachCity.Name))
		}
		fmt.Println(cityInfo.String())
		leftWorld.WriteString(cityInfo.String())
		leftWorld.WriteString("\n")
	}

	return leftWorld.String()
}
