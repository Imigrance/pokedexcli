package battle

import (
	"errors"
	"math/rand"

	"github.com/imigrance/pokedexcli/pokedex"
)

func aiChooseMove(pokemon *pokedex.Pokemon) error {
	rnd := rand.Intn(len(pokemon.LearnedMoves))
	i := 1
	for _, move := range pokemon.LearnedMoves {
		if i == rnd {
			pokemon.Battling.UseMove = move
			return nil
		}
		i++
	}

	return errors.New("randomisation error")
}
