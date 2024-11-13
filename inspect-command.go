package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("missing pokemon to inspect")
	}

	pokemon, exists := cfg.pokedex.GetPokemon(args[0])
	if !exists {
		return fmt.Errorf("you have not caught a %v", args[0])
	}

	fmt.Println()
	fmt.Printf("Name: %v\nWeight: %v \nHeight: %v\n", pokemon.Name, pokemon.Weight, pokemon.Height)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %v\n", t.Type.Name)
	}

	fmt.Println()

	return nil
}
