package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Unable to exit the program")
}

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()

	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *config) error {
	areas, err := cfg.pokeapiClient.GetAreas(cfg.nextLocURL)
	if err != nil {
		return fmt.Errorf("Error getting area list from PokeAPI: %w", err)
	}

	cfg.nextLocURL = areas.Next
	cfg.prevLocURL = areas.Previous

	// Display batch of Location Areas returned from the API
	for _, a := range areas.Results {
		fmt.Println(a.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocURL == nil {
		return errors.New("You are on the first page")
	}

	areas, err := cfg.pokeapiClient.GetAreas(cfg.prevLocURL)
	if err != nil {
		return fmt.Errorf("Error getting areas from PokeAPI: %w", err)
	}

	cfg.nextLocURL = areas.Next
	cfg.prevLocURL = areas.Previous

	// Display batch of Location Areas returned from the API
	for _, a := range areas.Results {
		fmt.Println(a.Name)
	}

	return nil
}
