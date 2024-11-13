package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	pokemons := cfg.pokedex.GetAllPokemon()

	if len(pokemons) < 1 {
		return fmt.Errorf("no pokemons caught yet")
	}

	fmt.Println("Your Pokedex:")

	for k := range pokemons {
		fmt.Printf("  - %v\n", k)
	}

	return nil
}
