package world

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Simulation struct {
	World          map[*City][]*City
	Iterations     int
	WorldFile      string
	NumberOfAliens int
	AlienNamesFile string
	Aliens         []*Alien
	Cities         []*City
	RandSeed       *rand.Rand
}

func NewSimulation(iterations, alienNumbers int, alienNames, worldFile string, randomSeed *rand.Rand) (*Simulation, error) {
	simulation := Simulation{
		Iterations:     iterations,
		WorldFile:      worldFile,
		NumberOfAliens: alienNumbers,
		AlienNamesFile: alienNames,
		World:          make(map[*City][]*City),
		RandSeed:       randomSeed,
	}
	simulation.CreateWorld()
	return &simulation, nil
}

func (sim *Simulation) CreateWorld() error {
	worldFile, err := os.Open(sim.WorldFile)
	if err != nil {
		return fmt.Errorf("Error Reading the world file : %s, Error: %s", sim.WorldFile, err.Error())
	}
	defer worldFile.Close()

	scanner := bufio.NewScanner(worldFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lineSplitOnSpace := strings.Split(scanner.Text(), " ")
		newCity := NewCity(lineSplitOnSpace[0])
		sim.World[newCity] = make([]*City, 0)
		sim.Cities = append(sim.Cities, newCity)
		lineSplitOnSpace = lineSplitOnSpace[1:]
		for _, cityDir := range lineSplitOnSpace {
			cityWithDirections := strings.Split(cityDir, "=")
			sim.World[newCity] = append(sim.World[newCity],
				NewCityWithDirections(cityWithDirections[1], cityWithDirections[0]))
		}
	}
	return nil
}

func (sim *Simulation) ViewWorld() {

	println("=========================================")
	println("The World before attack is")
	println("=========================================")

	for key, value := range sim.World {
		fmt.Printf("The City %s is connected to below cities\n", key.Name)
		for _, city := range value {
			fmt.Printf("\tThe City %s is %s to the %s\n", city.Name, city.Direction, key.Name)
		}
	}
}

func (sim *Simulation) CreateAliens() error {
	alienNames, err := os.Open(sim.AlienNamesFile)
	if err != nil {
		return fmt.Errorf("Error Reading the alien name file : %s, Error: %s", sim.AlienNamesFile, err.Error())
	}
	defer alienNames.Close()
	scanner := bufio.NewScanner(alienNames)
	scanner.Split(bufio.ScanLines)

	for aliens := 0; aliens < sim.NumberOfAliens && scanner.Scan(); aliens++ {
		sim.Aliens = append(sim.Aliens, NewAlien(scanner.Text()))
	}
	return nil
}

func (sim *Simulation) ViewAliens() error {

	println("=========================================")
	println("Alien Profiles")
	println("=========================================")

	for idx, alien := range sim.Aliens {
		fmt.Printf("The alien %d has a name %s\n", idx, alien.Name)
	}
	return nil
}

func (sim *Simulation) Start() error {
	alienExists := true
	for iteration := 0; iteration < sim.Iterations && alienExists; iteration++ {
		println("=========================================")
		fmt.Printf("Iteration %d\n", iteration)
		println("=========================================")
		println(len(sim.Cities))
		// do only first time
		if iteration == 0 {
			for idx := range sim.Aliens {
				min := 0
				max := len(sim.Cities) - 1
				randomCityIdx := sim.RandSeed.Intn(max-min+1) + min
				fmt.Printf("Choosing city %d\n", randomCityIdx)
				sim.Aliens[idx].CurrentCity = sim.Cities[randomCityIdx]
				sim.Cities[randomCityIdx].CurrentAlienIndex = append(sim.Cities[randomCityIdx].CurrentAlienIndex, sim.Aliens[idx].Name)
			}
		} else {

		}

		// find if more than two aliens are in the same city
		for _, city := range sim.Cities {
			if len(city.CurrentAlienIndex) > 1 {
				sim.DestoryAlien(city)
			}
		}

	}
	return nil
}

func (sim Simulation) DestoryAlien(city *City) error {
	var message strings.Builder
	message.WriteString("The city ")
	message.WriteString(city.Name)
	message.WriteString("is destoryed by the aliens ")
	for idx := 0; idx < len(city.CurrentAlienIndex)-1; idx++ {
		message.WriteString(city.CurrentAlienIndex[idx])
		message.WriteString(" and ")
		sim.DeleteAlien(city.CurrentAlienIndex[idx])
	}
	message.WriteString(city.CurrentAlienIndex[len(city.CurrentAlienIndex)-1])
	fmt.Println(message.String())
	return nil
}

func (sim *Simulation) DeleteAlien(alienName string) {
	for idx, alien := range sim.Aliens {
		if alien.Name == alienName {
			fmt.Println("Found the alien at index", idx)
			sim.Aliens = append(sim.Aliens[:idx], sim.Aliens[idx+1:]...)
		}
	}
}
