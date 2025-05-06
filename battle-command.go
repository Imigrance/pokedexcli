package main

import (
	"fmt"

	"github.com/imigrance/pokedexcli/battle"
	"github.com/imigrance/pokedexcli/pokedex"
)

func commandBattle(cfg *config, args ...string) error {
	if len(args) != 2 {
		return fmt.Errorf("missing player and/or enemy pokemon")
	}

	player := args[0]
	playerP, ok := cfg.pokedex.GetPokemon(player)
	if !ok {
		return fmt.Errorf("you have not caught a %v", player)
	}

	enemy := args[1]
	enemyP, err := initEnemy(cfg, enemy)
	if err != nil {
		return err
	}

	battle.NewBattle(&playerP, &enemyP)

	return nil
}

func initEnemy(cfg *config, pokemon string) (pokedex.Pokemon, error) {
	enemyPokemon, err := cfg.pokeapiClient.GetPokemon(pokemon)
	if err != nil {
		return pokedex.Pokemon{}, err
	}

	move_count := 0
	for _, move := range enemyPokemon.Moves {
		if move_count > 3 {
			break
		}
		if move.LevelLearnedAt == 1 {
			enemyPokemon.LearnedMoves[move.Name], err = cfg.pokeapiClient.GetMove(move.URL)
			if err != nil {
				return pokedex.Pokemon{}, err
			}
			move_count++
		}
	}
	return enemyPokemon, err
}
