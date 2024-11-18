package main

import (
	"fmt"

	"github.com/imigrance/pokedexcli/pokedex"
)

type Battle struct {
	playerPokemon *pokedex.Pokemon
	enemyPokemon  *pokedex.Pokemon
	turn          *pokedex.Pokemon
}

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
	enemyP, err := cfg.pokeapiClient.GetPokemon(enemy)
	if err != nil {
		return err
	}

	var starts *pokedex.Pokemon
	if err := getStarter(&playerP, &enemyP, starts); err != nil {
		return err
	}

	battle := Battle{
		playerPokemon: &playerP,
		enemyPokemon:  &enemyP,
		turn:          starts,
	}
	battle.startBattle()

	return nil
}

func getStarter(player, enemy, starts *pokedex.Pokemon) error {
	if player.Stats["speed"].BaseStat > enemy.Stats["speed"].BaseStat {
		starts = player
	} else {
		starts = enemy
	}

	return nil
}

func (b *Battle) startBattle() {
	for {
		fmt.Printf("Your %v's hp: ", b.playerPokemon.Name)
		return
	}
}
