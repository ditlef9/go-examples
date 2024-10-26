package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// ReadLog reads and displays all mission data from the log file in the specified directory.
func ReadLog(dirName, fileName string) error {
	filePath := filepath.Join(dirName, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	return nil
}
