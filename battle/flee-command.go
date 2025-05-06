package battle

import "fmt"

func battleCommandFlee(bcfg *battleConfig, args ...string) error {
	fmt.Println("Fleeing...")
	bcfg.battleling = false
	return nil
}
