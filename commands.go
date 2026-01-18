package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *config, arg string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Unable to exit the program")
}

func commandHelp(cfg *config, arg string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()

	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *config, arg string) error {
	areas, err := cfg.pokeapiClient.GetAreas(cfg.nextLocURL)
	if err != nil {
		return fmt.Errorf("Error getting area list: %w", err)
	}

	cfg.nextLocURL = areas.Next
	cfg.prevLocURL = areas.Previous

	// Display batch of Location Areas returned from the API or cache
	for _, a := range areas.Results {
		fmt.Println(a.Name)
	}

	return nil
}

func commandMapb(cfg *config, arg string) error {
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

func commandExplore(cfg *config, area string) error {
	if area == "" {
		return errors.New("Use 'explore <area_name>' to explore an area!")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(area)
	if err != nil {
		return fmt.Errorf("Unable to list Pokemon in %s: %w", area, err)
	}

	fmt.Println("Exploring %s...", area)
	fmt.Println("Found Pokemon:")
	for _, p := range pokemon {
		fmt.Println(" - %s", p)
	}

	return nil
}
