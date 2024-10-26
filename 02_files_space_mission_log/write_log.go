package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// WriteLog appends mission data to the log file in the specified directory.
func WriteLog(dirName, fileName, missionData string) error {
	filePath := filepath.Join(dirName, fileName)

	// Open the file in append mode to add new mission data
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(missionData)
	if err != nil {
		return fmt.Errorf("failed to write data to file: %v", err)
	}
	return nil
}
