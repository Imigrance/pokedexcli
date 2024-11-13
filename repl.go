package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/imigrance/pokedexcli/internal/pokeapi"
	"github.com/imigrance/pokedexcli/pokedex"
)

type config struct {
	commands        map[string]CliCommand
	pokeapiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
	pokedex         pokedex.Pokedex
}

func repl(cfg *config) {
	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		input := cleanInput(scanner.Text())
		cmdName := input[0]

		args := []string{}

		if len(input) > 1 {
			args = input[1:]
		}

		if command, exists := cfg.commands[cmdName]; exists {

			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}

			continue
		} else {
			fmt.Println("command not found, try 'Help'")
			continue
		}

	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Split(lower, " ")
	return words
}
