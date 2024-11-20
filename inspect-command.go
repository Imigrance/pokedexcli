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
	fmt.Printf("ID: %d\nName: %v\nWeight: %v \nHeight: %v\n", pokemon.ID, pokemon.Name, pokemon.Weight, pokemon.Height)
	fmt.Println("Moves:")
	for _, move := range pokemon.LearnedMoves {
		fmt.Printf("  -%v | Pp: %v\n", move.Name, move.Pp)
	}
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %v\n", t.Type)
	}

	fmt.Println()

	return nil
}
