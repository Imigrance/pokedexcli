package main

import (
	"fmt"
	"time"

	"github.com/imigrance/pokedexcli/internal/pokeapi"
	"github.com/imigrance/pokedexcli/pokedex"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	pokedex := pokedex.NewPokedex()

	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       pokedex,
	}

	cfg.commands = getCommands()

	err := loadSave(cfg)
	if err != nil {
		fmt.Println(err)
	}

	repl(cfg)
}
