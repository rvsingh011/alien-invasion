package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rvsingh011/Alien-invasion/utils"
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
	WorldFile = "./data/world.txt"
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
	println("This is an Alien Invasion (Do Not Panic!)")
	println("=========================================")

	// set the logger for debugging purposes
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("Unable to instintiate logger for the program: ", err.Error())
		os.Exit(1)
	}
	defer logger.Sync()
	logger.Debug("Starting a new alien invasion")

	// Validte the user input
	if err := utils.ValidateInput(iterations, alienNumber, alienNames, worldFile); err != nil {
		fmt.Println("Invalid input given by user", err.Error())
		os.Exit(1)
	}

}
