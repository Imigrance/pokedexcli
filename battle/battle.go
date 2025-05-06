package battle

import (
	"github.com/imigrance/pokedexcli/pokedex"
)

type Battle struct {
	battleType  string
	pokemons    map[string]*pokedex.Pokemon
	attackOrder []string
	attacking   string
	target      *pokedex.Pokemon
}

func NewBattle(player, enemy *pokedex.Pokemon) {

	pokes := map[string]*pokedex.Pokemon{}
	pokes["player"] = player
	pokes["enemy1"] = enemy

	newBattle := Battle{
		pokemons:    pokes,
		attackOrder: []string{},
	}

	bcfg := battleConfig{
		battle:     &newBattle,
		battleling: true,
	}

	newBattle.StartBattle(&bcfg)
}
