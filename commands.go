package main

import (
	"fmt"
	"os"
	"github.com/db-0/pokedexcli/internal/pokeapi"
)

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Unable to exit the program")
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()

	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	URL := "https://pokeapi.co/api/v2/location-area/"
	if c.Next != "" {
		URL = c.Next
	}

	areas, err := pokeapi.GetAreas(URL)
	if err != nil {
		return fmt.Errorf("Error getting areas from PokeAPI: %w", err)
	}

	// Next / Previous come through as null values and need to be cast as strings for
	// the logic to work properly
	if areas.Next != nil {
		c.Next = areas.Next.(string)
	} else {
		c.Next = ""
	}
	if areas.Previous != nil {
		c.Previous = areas.Previous.(string)
	} else {
		c.Previous = ""
	}

	// Display batch of Location Areas returned from the API
	for _, a := range areas.Results {
		fmt.Println(a.Name)
	}

	return nil
}

func commandMapb(c *config) error {
	if c.Previous == "" {
		fmt.Println("You're on the first page!")
		return nil
	}

	areas, err := pokeapi.GetAreas(c.Previous)
	if err != nil {
		return fmt.Errorf("Error getting areas from PokeAPI: %w", err)
	}

	// Next / Previous come through as null values and need to be cast as strings for
	// the logic to work properly
	if areas.Next != nil {
		c.Next = areas.Next.(string)
	} else {
		c.Next = ""
	}
	if areas.Previous != nil {
		c.Previous = areas.Previous.(string)
	} else {
		c.Previous = ""
	}

	// Display batch of Location Areas returned from the API
	for _, a := range areas.Results {
		fmt.Println(a.Name)
	}

	return nil
}
