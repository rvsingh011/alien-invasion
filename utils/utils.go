package utils

import (
	"math/rand"
	"strings"
)

// GetOppositeDirection returns opposite direction to the input direction
func GetOppositeDirection(direction string) string {
	switch strings.ToLower(direction) {
	case "north":
		return "south"
	case "south":
		return "north"
	case "east":
		return "west"
	case "west":
		return "east"
	}
	return ""
}

// GetRandomNumber returns a random number in range min and max using the source provided
func GetRandomNumber(min, max int, r *rand.Rand) int {
	return r.Intn(max-min+1) + min
}
