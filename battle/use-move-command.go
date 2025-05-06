package battle

import (
	"errors"
	"fmt"
)

func battleCommandUseMove(cfg *battleConfig, args ...string) error {
	if len(args) < 1 {
		return errors.New("missing move")
	}

	m := args[0]
	pokemon := *cfg.battle.pokemons["player"]

	move, exist := pokemon.LearnedMoves[m]
	if !exist {
		return fmt.Errorf("%v has not learnt %v", pokemon.Name, m)
	}

	pokemon.Battling.UseMove = move
	pokemon.Battling.Target = append(pokemon.Battling.Target, cfg.battle.pokemons["enemy1"])

	move.Pp -= 1
	pokemon.LearnedMoves[m] = move

	fmt.Println(move.Target.Name)

	switch move.Target.Name {
	case "selected-pokemon":
		//hits targeted pokemon
		fmt.Printf("%v used %v\n", pokemon.Name, move.Name)
		err := cfg.battle.attack()
		if err != nil {
			return err
		}
	case "all-other-pokemon":
		//hits all other pokemon than self
	case "random-opponent":
		//hits a random opponent
		err := cfg.battle.attack()
		if err != nil {
			return err
		}
	case "user":
		//hits self
	case "opponents-field":
		//hits all opponents
	case "user-or-ally":
		//hits either user or ally´
	case "users-field":
		//hits self and allies
	case "ally":
		//hits only ally
	case "selected-pokemon-me-first":
		//hits self first then enemy
	case "specific-move":
		//dunno...
	}

	return nil
}
