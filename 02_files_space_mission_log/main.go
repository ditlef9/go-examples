package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	dirName  = "space_mission_logs"
	fileName = "mission_log.txt"
)

func main() {
	// Step 1: Create directory and check if log file exists, create if not
	if err := CreateDirectory(dirName); err != nil {
		fmt.Println(err)
		return
	}

	if err := CheckAndCreateLogFile(dirName, fileName); err != nil {
		fmt.Println(err)
		return
	}

	// Step 2: Main program loop for user input
	for {
		fmt.Println("\n--- Space Mission Log ---")
		fmt.Println("1. Add a new mission log entry")
		fmt.Println("2. Read all mission log entries")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// Add a new mission log entry
			fmt.Println("\nEnter mission details below:")
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Mission Name: ")
			name, _ := reader.ReadString('\n')

			fmt.Print("Destination: ")
			destination, _ := reader.ReadString('\n')

			fmt.Print("Launch Date: ")
			launchDate, _ := reader.ReadString('\n')

			fmt.Print("Mission Duration: ")
			duration, _ := reader.ReadString('\n')

			fmt.Print("Objective: ")
			objective, _ := reader.ReadString('\n')

			// Prepare mission entry
			missionData := fmt.Sprintf("%s|%s|%s|%s|%s\n",
				strings.TrimSpace(name), strings.TrimSpace(destination), strings.TrimSpace(launchDate),
				strings.TrimSpace(duration), strings.TrimSpace(objective))

			// Write the mission entry to the log file
			if err := WriteLog(dirName, fileName, missionData); err != nil {
				fmt.Println("Error writing log:", err)
			} else {
				fmt.Println("Mission log entry added successfully!")
			}

		case 2:
			// Read and display all mission log entries
			fmt.Println("\n--- Mission Log Entries ---")
			fmt.Print("Mission Name	")
			fmt.Print("Destination	")
			fmt.Print("Launch Date	")
			fmt.Print("Mission Duration	")
			fmt.Print("Objective\n")
			if err := ReadLog(dirName, fileName); err != nil {
				fmt.Println("Error reading log:", err)
			}

		case 3:
			// Exit the program
			fmt.Println("Exiting the program.")
			return

		default:
			fmt.Println("Invalid choice. Please choose a valid option.")
		}
	}
}
