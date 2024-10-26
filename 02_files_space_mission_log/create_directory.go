package main

import (
	"fmt"
	"os"
)

// CreateDirectory checks if a directory exists and creates it if not.
func CreateDirectory(dirName string) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.Mkdir(dirName, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
		fmt.Printf("Directory '%s' created successfully.\n", dirName)
	} else {
		fmt.Printf("Directory '%s' already exists.\n", dirName)
	}
	return nil
}
