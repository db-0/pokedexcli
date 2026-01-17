package main

import (
	"time"

	"github.com/db-0/pokedexcli/internal/pokeapi"
)

func main() {
	// HTTP Client to be used for all PokeAPI calls
	pokeClient := pokeapi.NewClient(5 * time.Second)

	// HTTP Client accessible through configuration struct
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	// Start the REPL loop
	startRepl(cfg)
}
