package battle

import (
	"github.com/imigrance/pokedexcli/pokedex"
)

type Battle struct {
	playerPokemon *pokedex.Pokemon
	enemyPokemon  *pokedex.Pokemon
	attacking     *pokedex.Pokemon
}

func NewBattle(player, enemy *pokedex.Pokemon) Battle {
	return Battle{
		playerPokemon: player,
		enemyPokemon:  enemy,
	}
}
