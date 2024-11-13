package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("missing pokemon to catch")
	}

	_, ok := cfg.pokedex.GetPokemon(args[0])
	if ok {
		fmt.Println("Pokemon already caught")
		return nil
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a pokeball at %v...\n", args[0])

	rnd := rand.Intn(pokemon.BaseExperience)

	if rnd < (pokemon.BaseExperience / 2) {
		fmt.Printf("%v got away\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%v was caught!\n", pokemon.Name)
	cfg.pokedex.AddPokemon(pokemon)
	return nil
}
