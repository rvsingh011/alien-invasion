package utils

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func ValidateInput(iterations, alienNumbers int, alienNames, worldFile string) error {
	if iterations < 0 {
		return fmt.Errorf("Number of iterations cannot be negative")
	}
	if alienNumbers < 0 {
		return fmt.Errorf("Number of aliens cannot be negative")
	}
	alienNamesData, err := os.ReadFile(alienNames)
	if err != nil {
		return fmt.Errorf("Unable to read the alien names: %s", err.Error())
	}
	numberOfAlienNames, err := lineCounter(bytes.NewReader(alienNamesData))
	if err != nil {
		return fmt.Errorf("Unable to count alien names: %s", err.Error())
	}
	if numberOfAlienNames < alienNumbers {
		return fmt.Errorf("There is a 1:1 mapping between alien name and number of aliens, the number of alien names should be greater than or equal to the number of aliens specified. Number of alines specified: %d, Number of names found: %d", alienNumbers, numberOfAlienNames)
	}
	// TODO: Create a advanced validator by parsing the format
	if _, err := os.ReadFile(worldFile); err != nil {
		return fmt.Errorf("Unable to read the world file")
	}

	return nil
}
