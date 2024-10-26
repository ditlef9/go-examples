package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// CheckAndCreateLogFile checks if the log file exists and creates it with default data if not.
func CheckAndCreateLogFile(dirName, fileName string) error {
	filePath := filepath.Join(dirName, fileName)

	// Check if the file already exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// If the file doesn't exist, create it and add default mission data
		defaultData := `Mars Exploration Rover - Spirit|Mars|June 10, 2003|6 years, 2 months|To explore the Martian surface, study its geology, and search for signs of past water activity.
Voyager 1|Interstellar Space|05.09.1977|Ongoing (over 40 years and counting)|To study the outer planets and their moons, and to travel beyond the influence of the Sun’s gravity, reaching interstellar space for the first time.`

		// Create the file and write default data
		err = os.WriteFile(filePath, []byte(defaultData+"\n"), 0644)
		if err != nil {
			return fmt.Errorf("failed to create file with default data: %v", err)
		}
		fmt.Printf("File '%s' created with default mission data.\n", filePath)
	}
	return nil
}
