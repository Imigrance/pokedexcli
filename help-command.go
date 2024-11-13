package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("\nWelcome to the Pokedex! \nUsage:")
	fmt.Println()

	for _, cmd := range cfg.commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()
	return nil
}
