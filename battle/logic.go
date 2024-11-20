package battle

import (
	"fmt"
	"strings"

	"github.com/imigrance/pokedexcli/pokedex"
)

func (b *Battle) StartBattle() {
	b.battleREPL()
}

func (b *Battle) battleREPL() {
	player := *b.playerPokemon
	enemy := *b.enemyPokemon

	for {
		fmt.Println()
		fmt.Printf("### %v VS %v ###\n", strings.ToUpper(player.Name), strings.ToUpper(enemy.Name))
		fmt.Printf("Your %v's hp: %d\n", player.Name, player.Stats["hp"].BaseStat)
		fmt.Printf("%v's hp: %d\n", enemy.Name, enemy.Stats["hp"].BaseStat)
		fmt.Println()
		fmt.Println("Your moves:")
		for _, move := range player.LearnedMoves {
			fmt.Printf("Name: %v PP: %v\nPower: %d Type: %v\n\n", move.Name, move.Pp, move.Power, move.DamageClass.Name)
		}
		return
	}
}

func (b *Battle) canMoveHit(move pokedex.PokeMove) bool {
	// TODO
	return true
}

func (b *Battle) setAttacker(playerMove, enemyMove pokedex.PokeMove) {
	if playerMove.Priority != enemyMove.Priority {
		if playerMove.Priority > enemyMove.Priority {
			b.attacking = b.playerPokemon
		} else {
			b.attacking = b.enemyPokemon
		}
	} else if b.playerPokemon.Stats["speed"].BaseStat != b.enemyPokemon.Stats["speed"].BaseStat {
		if b.playerPokemon.Stats["speed"].BaseStat > b.enemyPokemon.Stats["speed"].BaseStat {
			b.attacking = b.playerPokemon
		} else {
			b.attacking = b.enemyPokemon
		}
	} else {
		b.attacking = nil
	}
}

func (b *Battle) calculateDamage() int {
	// TODO
	return 0
}
