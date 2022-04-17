package world

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Simulation struct {
	World          map[City][]*City
	Iterations     int
	WorldFile      string
	NumberOfAliens int
	AlienNamesFile string
	Aliens         []*Alien
}

func NewSimulation(iterations, alienNumbers int, alienNames, worldFile string) (*Simulation, error) {
	simulation := Simulation{
		Iterations:     iterations,
		WorldFile:      worldFile,
		NumberOfAliens: alienNumbers,
		AlienNamesFile: alienNames,
		World:          make(map[City][]*City),
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
		sim.World[*newCity] = make([]*City, 0)
		lineSplitOnSpace = lineSplitOnSpace[1:]
		for _, cityDir := range lineSplitOnSpace {
			cityWithDirections := strings.Split(cityDir, "=")
			sim.World[*newCity] = append(sim.World[*newCity],
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
