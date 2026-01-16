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
	callback    func() error
}

func startRepl() {
	inputReader := bufio.NewScanner(os.Stdin)

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
			err := command.callback()
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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Unable to exit the program")
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()

	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
