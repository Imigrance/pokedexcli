package battle

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (b *Battle) battleREPL(bcfg *battleConfig) {

	player := b.pokemons["player"]
	enemy := b.pokemons["enemy1"]

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
		fmt.Print("Battle > ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		input := cleanInput(scanner.Text())
		cmdName := input[0]

		args := []string{}

		if len(input) > 1 {
			args = input[1:]
		}

		command, exists := getBattleCommands()[cmdName]
		if exists {
			err := command.callback(bcfg, args...)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Println("command not found, try 'Help'")
			continue
		}

		if player.Stats["hp"].BaseStat < 1 {
			fmt.Printf("%v wins!\n", enemy.Name)
			return
		}
		if enemy.Stats["hp"].BaseStat < 1 {
			fmt.Printf("%v wins!\n", enemy.Name)
			return
		}
		if !bcfg.battleling {
			return
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Split(lower, " ")
	return words
}
