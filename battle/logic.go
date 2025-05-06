package battle

import (
	"fmt"
	"math/rand"

	"github.com/imigrance/pokedexcli/pokedex"
)

type battleConfig struct {
	battle       *Battle
	playerEffect pokedex.MoveEffect
	enemyEffect  pokedex.MoveEffect
	battleling   bool
}

func (b *Battle) StartBattle(bcfg *battleConfig) {

	b.battleREPL(bcfg)
}

func (b *Battle) setAttackOrder() {
	pokes := map[string]pokedex.Pokemon{}
	for k, v := range b.pokemons {
		pokes[k] = *v
	}

	var poke string

	for range len(b.pokemons) {
		for k, v := range pokes {
			if poke == "" {
				poke = k
				continue
			}
			if v.Battling.UseMove.Priority > b.pokemons[poke].Battling.UseMove.Priority {
				poke = k
				continue
			}
			if v.Stats["speed"].BaseStat > b.pokemons[poke].Stats["speed"].BaseStat {
				poke = k
				continue
			}
		}
		b.attackOrder = append(b.attackOrder, poke)
		delete(pokes, poke)
	}
}

func (b *Battle) attack() error {
	err := aiChooseMove(b.pokemons["enemy1"])
	if err != nil {
		return err
	}

	fmt.Printf("AI uses %v", b.pokemons["enemy1"].Battling.UseMove.Name)

	b.setAttackOrder()

	for _, attacker := range b.attackOrder {
		fmt.Printf("Attacker: %v\n", attacker)
		for _, target := range b.pokemons[attacker].Battling.Target {
			fmt.Printf("Target: %v\n", target)
			b.target = target
			if !b.canMoveHit(b.pokemons[attacker].Battling.UseMove) {
				fmt.Printf("%v cannot be hit\n", b.target.Name)
				continue
			}
			if !b.doesMoveHit(b.pokemons[attacker].Battling.UseMove) {
				fmt.Printf("%v misses\n", b.pokemons[attacker].Battling.UseMove.Name)
				continue
			}

			damage := b.calculateDamage(b.pokemons[attacker].Battling.UseMove)
			fmt.Println(damage)
			b.target.TakeDmg(damage)
		}
	}

	return nil
}

func (b *Battle) canMoveHit(move pokedex.PokeMove) bool {
	//WIP - should check if pokemon is flying or in ground
	//Specific attacks can hit if pokemon is flying or in ground
	return true
}

func (b *Battle) doesMoveHit(move pokedex.PokeMove) bool {
	rnd := rand.Intn(100)
	return rnd <= move.Accuracy
}

func (b *Battle) calculateDamage(move pokedex.PokeMove) int {
	ad := b.pokemons[b.attacking].Stats["attack"].BaseStat / b.target.Stats["defense"].BaseStat

	dmg := ((((2*b.pokemons[b.attacking].Level*1)/5)+2)*move.Power*ad)/50 + 2

	return dmg
}
