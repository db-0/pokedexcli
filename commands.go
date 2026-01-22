package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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

	pokemon, err := cfg.pokeapiClient.ListPokemon(area)
	if err != nil {
		return fmt.Errorf("Unable to list Pokemon in %s: %w", area, err)
	}

	fmt.Printf("Exploring %s...\n", area)
	fmt.Printf("Found Pokemon:\n")
	for _, p := range pokemon.PokemonEncounters {
		fmt.Printf(" - %s\n", p.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, mon string) error {
	if mon == "" {
		return errors.New("Use 'catch <pokemon>' to catch a pokemon!")
	}

	pokemon, err := cfg.pokeapiClient.CatchPokemon(mon)
	if err != nil {
		return fmt.Errorf("Error catching Pokemon: %w", err)
	}

	// Initialize random seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	// Chance of catching is based on the Pokemon's base experience
	// Exp: 10 (100%) to Exp: 400 (0% - legendary monsters cap at 350)
	catchChance := 100.0 - ((float32(pokemon.BaseExperience) - 10.0) / 400.0 * 100.0)
	roll := r.Float32() * 100.0
	isCaught := roll < catchChance

	if isCaught {
		cfg.userPokemon[pokemon.Name] = pokemon
		fmt.Printf("You caught %s!\n", pokemon.Name)
	} else {
		fmt.Printf("Darn it! %s got away...\n", pokemon.Name)
	}

	return nil
}

func commandInspect(cfg *config, mon string) error {
	// Validate argument
	if mon == "" {
		return errors.New("Use 'inspect <pokemon>' to inspect a captured Pokemon!")
	}

	// Check to ensure this pokemon has been captured
	pokemon, ok := cfg.userPokemon[strings.ToLower(mon)]
	if !ok {
		return errors.New("That Pokemon was not found in your Pokedex!")
	}

	// Print the Pokemon statistics:
	fmt.Printf("\nName: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, poketype := range pokemon.Types {
		fmt.Printf("  -%s\n", poketype.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *config, arg string) error {
	fmt.Println("Your Pokedex:")
	if len(cfg.userPokemon) == 0 {
		fmt.Println("You haven't caught any Pokemon!")
		fmt.Println("(Use 'catch <pokemon>' to try and catch one!)")
		return nil
	}
	for _, pokemon := range cfg.userPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
