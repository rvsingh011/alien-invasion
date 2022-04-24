package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/rvsingh011/alien-invasion/simulation"
	"github.com/rvsingh011/alien-invasion/utils"
	"go.uber.org/zap"
)

const (
	// DefaultIterations used if number of iteration is not otherwise specified
	DefaultIterations int = 10000
	// DefaultNumberOfAliens used if number of Aliens is not otherwise specified
	DefaultNumberOfAliens int = 10
	// AlienNames default file if not specified
	AlienNames = "./data/alien_names.txt"
	// CitiesFiles used if not specified
	WorldFile = "./data/world-example-1.txt"
	// Default log level is info
	LogLevel = "info"
)

var (
	iterations, alienNumber int
	worldFile, alienNames   string
)

// init cli flags
func init() {
	flag.IntVar(&iterations, "iterations", DefaultIterations, "number of iterations")
	flag.IntVar(&alienNumber, "aliens", DefaultNumberOfAliens, "number of aliens invading")
	flag.StringVar(&alienNames, "names", AlienNames, "a file used as alien names input")
	flag.StringVar(&worldFile, "world", WorldFile, "a file used as world map input")
	// flag.StringVar(&logLevel, "loglevel", LogLevel, "log level for the program")
	flag.Parse()
}

func main() {
	// Init the logger, will only be used for debug messages
	println("=========================================")
	println("Starting the Alien Invasion Simulation")
	println("=========================================")

	// set the logger for debugging purposes
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("Unable to instintiate logger for the program: ", err.Error())
		os.Exit(1)
	}

	defer logger.Sync()

	// Validte the user input
	if err := utils.ValidateInput(iterations, alienNumber, alienNames, worldFile); err != nil {
		fmt.Println("Invalid User Input, Reason: ", err.Error())
		os.Exit(1)
	}

	// Create the Seed for the psedudo random genrator
	randomSeed := buildSeed()

	// create the simulation for the alien invasion
	simulation, err := simulation.NewSimulation(iterations, alienNumber, alienNames, worldFile, randomSeed, logger)
	if err != nil {
		fmt.Println("Error Initiating a world: ", err.Error())
		os.Exit(1)
	}
	simulation.CreateWorld()
	simulation.ViewWorld()
	simulation.CreateAliens()
	simulation.ViewAliens()
	simulation.Start()
	simulation.EndAndConclude()

}

func buildSeed() *rand.Rand {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	return rand.New(source)
}
