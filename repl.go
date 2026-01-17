package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     string
	Previous string
}

func startRepl() {
	inputReader := bufio.NewScanner(os.Stdin)

	conf := config{
		Next:     "",
		Previous: "",
	}

	// Main REPL loop
	for {
		// Prompt
		fmt.Print("Pokedex > ")
		inputReader.Scan()

		input := cleanInput(inputReader.Text())
		if len(input) == 0 {
			fmt.Println("Please enter a command.")
			continue
		}

		commandInput := input[0]

		if command, exists := getCommands()[commandInput]; exists {
			err := command.callback(&conf)
			if err != nil {
				fmt.Println(fmt.Errorf("Unable to run command, %v: %w", command.name, err))
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Lists location areas in the Pokemon world, subsequent calls advance through the list",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous Areas page",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
