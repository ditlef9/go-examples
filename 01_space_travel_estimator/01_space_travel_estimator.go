package main

import "fmt"

// Define planet distances from Earth in million kilometers
var planetDistances = map[string]float64{
	"Venus":            41.4,
	"Mars":             78.3,
	"Jupiter":          628.7,
	"Saturn":           1275.0,
	"Proxima-Centauri": 40_208_000,  // 4.22 light-years
	"Kepler-22b":       587_000_000, // 587 million light-years (approx.)
}

// Define spacecraft speeds in km/h, including some Star Wars ships
var spacecraftSpeeds = map[string]float64{
	"Apollo":            39800,         // Apollo mission speed
	"Voyager":           62000,         // Voyager spacecraft speed
	"New-Horizons":      58000,         // New Horizons speed
	"Millennium-Falcon": 1_000_000_000, // Hypothetical Star Wars ship speed
	"X-Wing-Fighter":    1_200_000,     // Hypothetical Star Wars ship speed
	"TIE-Fighter":       1_050_000,     // Hypothetical Star Wars ship speed
}

func main() {
	var planetChoice, spacecraftChoice string

	// Display planet options
	fmt.Println("Select a planet or distant star system to travel to:")
	for planet := range planetDistances {
		fmt.Printf(" - %s\n", planet)
	}
	fmt.Print("Enter your choice: ")
	fmt.Scan(&planetChoice)

	// Get the distance to the selected planet or star system
	distance, planetExists := planetDistances[planetChoice]
	if !planetExists {
		fmt.Println("Invalid planet choice. Exiting program.")
		return
	}

	// Display spacecraft options
	fmt.Println("\nSelect a spacecraft:")
	for spacecraft := range spacecraftSpeeds {
		fmt.Printf(" - %s\n", spacecraft)
	}
	fmt.Print("Enter your choice: ")
	fmt.Scan(&spacecraftChoice)

	// Get the speed of the selected spacecraft
	speed, spacecraftExists := spacecraftSpeeds[spacecraftChoice]
	if !spacecraftExists {
		fmt.Println("Invalid spacecraft choice. Exiting program.")
		return
	}

	// Convert distance from million km to km
	distance *= 1_000_000

	// Calculate travel time
	travelTimeHours := distance / speed
	travelTimeDays := travelTimeHours / 24
	travelTimeYears := travelTimeDays / 365

	// Output travel times
	fmt.Printf("\nTraveling to %s on %s:\n", planetChoice, spacecraftChoice)
	fmt.Printf(" - Estimated travel time in hours: %.2f\n", travelTimeHours)
	fmt.Printf(" - Estimated travel time in days: %.2f\n", travelTimeDays)
	fmt.Printf(" - Estimated travel time in years: %.2f\n", travelTimeYears)
}
